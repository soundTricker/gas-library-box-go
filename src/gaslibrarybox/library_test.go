package gaslibrarybox_test

import (
	. "./"
	"appengine/aetest"
	"appengine/datastore"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var c aetest.Context

var _ = Describe("Library", func() {

	BeforeEach(func() {
		fmt.Printf("test")
		var err error
		c, err = aetest.NewContext(nil)
		Expect(err).NotTo(HaveOccured())
	})

	It("should be got a library by library name", func() {

		l := &Library{}
		err := GetLibrary(c, "hoge", l)
		Expect(err).Should(Equal(datastore.ErrNoSuchEntity))

	})

	AfterEach(func() {
		defer c.Close()
	})

})
