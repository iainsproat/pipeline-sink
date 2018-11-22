package interview_2018_11_22

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _= Describe("Sink", func() {
	Context("For no messages",func(){
		SUT := Sink{
			messages: []PipelineNotification{},
		}

		It("Does not show success", func(){
			response := SUT.Check("A", "GUID1", []string{
				"pipelineA",
				"pipelineB",
			})
			Expect(response).ToNot(Equal(true))
		})
	})
})
