package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestArrayCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ArrayCp Suite")
}
