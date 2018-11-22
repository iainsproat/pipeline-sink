package sinkchecker

import (
	. "../sinkclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockSinkClient struct {
	Messages []PipelineNotification
}

func (s MockSinkClient) Select(triggeringPipeline string, commitID string) []PipelineNotification {
	return s.Messages
}

var _ = Describe("Sink Checker", func() {
	Context("With no pipeline notifications", func() {
		sinkClient := MockSinkClient {
			Messages: []PipelineNotification{},
		}

		SUT := SinkChecker{
			client: sinkClient,
		}

		It("Does not show success", func() {
			response := SUT.Check("A",
				"GUID1",
				[]string{
				"pipelineA",
				"pipelineB",
			})
			Expect(response).ToNot(Equal(true))
		})
	})
})
