package statistics

import (
	"encoding/json"
	"fmt"
	"github.com/startopsz/rule/pkg/os/filesystem"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	
	rtimestamp "github.com/startopsz/rule/pkg/timestamp"
)

type LogFormat struct {
	Timestamp            float64         `json:"timestamp"`
	RemoteAddr           string          `json:"remote_addr"`
	BodyBytesSent        int64           `json:"body_bytes_sent"`
	Status               string          `json:"status"`
	RequestTime          float64         `json:"request_time"`
	UpstreamResponseTime string          `json:"upstream_response_time"`
}

type SecondReport struct {
	Timestamp         int64
	RequestCount      int64
	TotalRequestTime  float64
	TotalResponseTime float64
	TotalBodyByteSize int64
}


func SecondStatistics(path string, tail bool, printRemoteAddrCount bool) {
	
	if !filesystem.IsFile(path) {
		log.Fatal("path is not file.")
		return
	}
	
	buf := make([]byte, 1024)
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("open path err:", err)
		return
	}
	
	defer f.Close()
	
	var last string
	var lastTimestamp int64
	var logFormat LogFormat
	var secondReport SecondReport
	
	secondIPMap := make(map[string]int64)
	
	// tail 未实现读取之前 N 行
	if tail {
		ff, err := f.Stat()
		if err == nil {
			fsize := ff.Size()
			_, _ =f.Seek(fsize, 0)
		}
	}

	//_, _ = f.ReadAt(buf, fsize)
	
	for {
		n, err := f.Read(buf)
		// 实现 tail
		if n == 0 && err == io.EOF && tail == true {
			continue
		}
		
		if err == io.EOF {
			//fmt.Println("读取完毕 EOF")
			break
		} else if n == 0 {
			//fmt.Println("内容为 0")
			break
		} else if err != nil {
			log.Fatal("read err: ", err.Error())
			return
		}
		
		readString := string(buf[:n])
		newReadString := last + readString
		
		readStrings := strings.Split(newReadString, "\n")
		// 读取每一行信息
		for _, content := range readStrings[:len(readStrings)-1] {
			err := json.Unmarshal([]byte(content), &logFormat)
			if err != nil {
				fmt.Println("json err: ", err, ". content: ", content)
				continue
			}
			
			timestamp := int64(logFormat.Timestamp)
			
			if lastTimestamp == 0 {
				lastTimestamp = timestamp
			}
			
			if timestamp > lastTimestamp {
				
				fmt.Printf("Time: %s, Timestamp: %d, RequestCount: %d, TotalBodyByteSize: %d(KB), AvgRequestTime: %f, AvgResponseTime: %f, RemoteAddrCount: %d.\n", rtimestamp.Time(secondReport.Timestamp*1000).In(time.Local).Format("2006-01-02 15:04:05"),
					secondReport.Timestamp,
					secondReport.RequestCount,
					secondReport.TotalBodyByteSize/1024,
					secondReport.TotalRequestTime/float64(secondReport.RequestCount),
					secondReport.TotalResponseTime/float64(secondReport.RequestCount),
					len(secondIPMap),
				)
				
				if printRemoteAddrCount {
					for ip, count := range secondIPMap {
						fmt.Printf("remote_addr: %s,\tcount: %d.\n", ip, count)
					}
					fmt.Print("\n")
				}
				// reset
				lastTimestamp = timestamp
				secondReport = SecondReport{}
				secondIPMap = make(map[string]int64)
			}
			
			secondReport.Timestamp = timestamp
			secondReport.TotalRequestTime += logFormat.RequestTime
			secondReport.RequestCount += 1
			secondReport.TotalBodyByteSize += logFormat.BodyBytesSent
			
			if logFormat.UpstreamResponseTime != "-" {
				upstreamResponseTime, err := strconv.Atoi(logFormat.UpstreamResponseTime)
				if err == nil {
					secondReport.TotalResponseTime += float64(upstreamResponseTime)
				}
			}
			
			if _, ok := secondIPMap[logFormat.RemoteAddr]; !ok {
				secondIPMap[logFormat.RemoteAddr] = 0
			}
			secondIPMap[logFormat.RemoteAddr] += 1
		}
		last = readStrings[len(readStrings)-1]
	}
}
