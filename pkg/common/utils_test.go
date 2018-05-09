package common

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "All Utils Test")
}

var _ = Describe("GetFromSlice", func() {
	var s []interface{}
	Context("initially", func() {
		It("Creating slice of ints{0..9}", func() {
			for i := 0; i < 10; i++ {
				s = append(s, i)
			}
			Ω(len(s)).Should(BeEquivalentTo(10))
		})
	})

	Context("When searching item", func() {
		It("existing value 5 should be found", func() {
			res, item := GetFromSlice(s, 5)
			Ω(item).Should(BeEquivalentTo(5))
			Ω(res).ShouldNot(BeNil())
		})

		It("existing value 20 shouldn't be found", func() {
			res, item := GetFromSlice(s, 20)
			Ω(item).Should(BeNil())
			Ω(res).Should(BeEquivalentTo(-1))
		})
	})
})

var _ = Describe("IsError", func() {
	Context("Check results", func() {
		It("Error parameter should return true", func() {
			var v interface{} = new(error)
			Ω(IsError(v)).Should(BeTrue())
		})
		It("Non-error parameter should return false", func() {
			var v interface{} = make(chan int)
			Ω(IsError(v)).Should(BeFalse())
		})
	})
})
