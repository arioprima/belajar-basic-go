package main_test

import (
	sumarray "belajar-basic-golang/array/1-array-basic/2-array-cp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SumArray", func() {
	Context("dengan array kosong", func() {
		It("harus mengembalikan 0", func() {
			total := sumarray.SumArray([]int{})
			Expect(total).To(Equal(0))
		})
	})

	Context("dengan array satu elemen", func() {
		It("harus mengembalikan nilai elemen tersebut", func() {
			total := sumarray.SumArray([]int{1})
			Expect(total).To(Equal(1))
		})
	})

	Context("dengan array beberapa elemen", func() {
		It("harus mengembalikan jumlah elemen tersebut", func() {
			total := sumarray.SumArray([]int{1, 2, 3})
			Expect(total).To(Equal(6))
		})
	})
})
