package main_test

import (
	runningsum "belajar-basic-golang/array/1-array-basic/3-array-cp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("runningSum", func() {
	Context("given an array of integers", func() {
		It("should return the running sum of the array", func() {
			nums := []int{1, 2, 3, 4}
			expectedResult := []int{1, 3, 6, 10}

			result := runningsum.RunningSum(nums)
			Expect(result).To(Equal(expectedResult))
		})

		It("should handle an empty array", func() {
			nums := []int{}
			expectedResult := []int{}

			result := runningsum.RunningSum(nums)
			Expect(result).To(Equal(expectedResult))
		})

		It("should handle an array with a single element", func() {
			nums := []int{5}
			expectedResult := []int{5}

			result := runningsum.RunningSum(nums)
			Expect(result).To(Equal(expectedResult))
		})
	})
})
