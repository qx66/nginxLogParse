package main

import (
	"flag"
	"fmt"
	"github.com/StartOpsTools/nginxLogParse/pkg/statistics"
	"github.com/startopsz/rule/pkg/os/filesystem"
	"go.uber.org/zap"
)

var path string
var tail bool
var printRemoteAddCount bool
var printHttpCodeCount bool
var printUpstreamDistribute bool

func init() {
	flag.StringVar(&path, "path", "", "-path")
	flag.BoolVar(&tail, "tail", false, "-tail")
	flag.BoolVar(&printRemoteAddCount, "printRemoteAddCount", false, "-printRemoteAddCount")
	flag.BoolVar(&printHttpCodeCount, "printHttpCodeCount", false, "-printHttpCodeCount")
	flag.BoolVar(&printUpstreamDistribute, "printUpstreamDistribute", false, "-printUpstreamDistribute")
}

func main() {
	flag.Parse()
	//
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("新建logger失败, err: ", err)
		return
	}
	//
	if !filesystem.IsFile(path) {
		logger.Error(
			"path is not filesystem",
		)
		return
	}
	
	//
	statistics.SecondStatistics(path, tail, printRemoteAddCount)
}
