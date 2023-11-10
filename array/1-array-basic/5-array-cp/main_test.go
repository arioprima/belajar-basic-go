// even_numbers_test.go
package main_test

import (
	evenNumbers "belajar-basic-golang/array/1-array-basic/5-array-cp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EvenNumbers", func() {
	var (
		nums []int
	)

	BeforeEach(func() {
		// Inisialisasi data yang dibutuhkan untuk testing
		nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	})

	Context("When the input slice contains even numbers", func() {
		It("should return a slice of even numbers and the count of even numbers", func() {
			evenNumbers, countEven := evenNumbers.EvenNumbers(nums)

			// Assertion menggunakan Gomega
			Expect(evenNumbers).To(ConsistOf(2, 4, 6, 8, 10))
			Expect(countEven).To(Equal(5))
		})
	})

	Context("When the input slice does not contain even numbers", func() {
		It("should return an empty slice and count of zero", func() {
			// Mengubah data agar tidak ada angka genap
			nums = []int{1, 3, 5, 7, 9}

			evenNumbers, countEven := evenNumbers.EvenNumbers(nums)

			// Assertion menggunakan Gomega
			Expect(evenNumbers).To(BeEmpty())
			Expect(countEven).To(Equal(0))
		})
	})
})
