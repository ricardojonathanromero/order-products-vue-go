package repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var globalTest *testing.T

func TestRepositories(t *testing.T) {
	globalTest = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repositories Suite")
}
