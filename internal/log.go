package internal

import (
	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qblog"
)

//go:generate versioner -pkg internal -p V -bno build_no ../VERSION version.go

var (
	log                     = qblog.New("edpcc")
	LogCfg c4hgol.LogConfig = log
)
