package internal

import "git.fractalqb.de/fractalqb/qbsllm"

//go:generate versioner -pkg internal -p V -bno build_no ../VERSION version.go

var (
	log    = qbsllm.New(qbsllm.Lnormal, "edpcc", nil, nil)
	LogCfg = qbsllm.NewConfig(log)
)
