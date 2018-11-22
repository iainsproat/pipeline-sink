package sinkchecker

import . "../sinkclient"

type SinkChecker struct {
	client SinkClient
}

func (s SinkChecker) Check(triggeringPipeline string, commitID string, pipelines []string) bool {
	return false
}
