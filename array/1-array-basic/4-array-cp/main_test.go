package main_test

import (
	minArray "belajar-basic-golang/array/1-array-basic/4-array-cp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MinArrayCp", func() {
	Context("When the input array is not empty", func() {
		It("should return the index of the minimum value", func() {
			nums := []int{5, 3, 1, 8, 4}
			expectedIndex := 2
			expectedValue := 1

			index, value := minArray.MinArrayCp(nums)
			Expect(index).To(Equal(expectedIndex))
			Expect(value).To(Equal(expectedValue))
		})
	})

	Context("When the input array is empty", func() {
		It("should return index 0 and value 0", func() {
			nums := []int{}
			expectedIndex := 0
			expectedValue := 0

			index, value := minArray.MinArrayCp(nums)
			Expect(index).To(Equal(expectedIndex))
			Expect(value).To(Equal(expectedValue))
		})
	})
})
