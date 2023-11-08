package main_test

import (
	arraycp "belajar-basic-golang/array/1-array-basic/1-array-cp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArrayCp", func() {
	Describe("Display Fruit at Index 2", func() {
		It("should display the fruit at index 2", func() {
			result := arraycp.DisplayFruitAtIndex2()
			Expect(result).To(Equal("apel"))
		})
	})

	Describe("Display All Fruits", func() {
		It("should display all fruits", func() {
			result := arraycp.DisplayAllFruits()
			Expect(result).To(Equal([]string{"mangga", "jeruk", "apel", "anggur", "melon"}))
		})
	})
})
