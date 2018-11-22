package sinkchecker

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSinkChecker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SinkChecker Suite")
}
