package sinkclient

type PipelineNotification struct {
	PipelineID           string
	TriggeringPipelineID string
	Commit               string
	Status               bool
}

type SinkClient interface {
	Select(triggeringPipeline string, commitID string) []PipelineNotification
}