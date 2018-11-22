package interview_2018_11_22
type PipelineNotification struct{
	PipelineID string
	TriggeringPipelineID string
	Commit string
	Status bool
}

type Sink struct{
	messages []PipelineNotification
}

func (s Sink) selectFromDatabase(triggeringPipeline string, commitID string) []PipelineNotification {
	return s.messages
}

func (s Sink) Check(triggeringPipeline string, commitID string, pipelines []string) bool {
  return false
}