package resources_test

import (
	"fmt"

	"github.com/midoblgsm/testMe/fakes"
	"github.com/midoblgsm/testMe/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("local-client", func() {
	var (
		myString          string
		fakeTestInterface *fakes.FakeMyTestInterface
		err               error
		result            string
	)
	BeforeEach(func() {
		fakeTestInterface = new(fakes.FakeMyTestInterface)
		myString = "fake-param"
	})

	Context(".Activate", func() {
		It("should succeed when myFakeInterface returns a fake string with no error", func() {
			fakeTestInterface.MyTestMethodReturns("fake-string", nil)
			result, err = fakeTestInterface.MyTestMethod(myString)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("fake-string"))
		})

		It("should fail when myFakeInterface returns an error", func() {
			fakeTestInterface.MyTestMethodReturns("", fmt.Errorf("fake-error"))
			result, err = fakeTestInterface.MyTestMethod(myString)
			Expect(err).To(HaveOccurred())
			Expect(result).To(Equal(""))
		})

		It("should succeed when myFakeInterface returns a string and no error", func() {
			fakeTestInterface.MyTestMethodReturns("fake-string", nil)
			impl := resources.NewMyStruct()
			err = impl.MyFunc("test-fake", fakeTestInterface)
			Expect(err).ToNot(HaveOccurred())
		})

	})
})
