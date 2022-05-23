package main

import (
	"flag"
	"github.com/StartOpsTools/nginxLogParse/pkg/statistics"
)


var path string
var tail bool
var printRemoteAddCount bool

func init() {
	flag.StringVar(&path, "path", "", "")
	flag.BoolVar(&tail, "tail", false, "")
	flag.BoolVar(&printRemoteAddCount,"printRemoteAddCount", false,"")
}

func main() {
	flag.Parse()
	statistics.SecondStatistics(path, tail, printRemoteAddCount)
}
