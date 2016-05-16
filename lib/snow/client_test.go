package snow_test

import (
	"net/url"

	. "github.com/allspark-thd/git-cf-deploy/lib/snow"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var (
		cfg    Config
		client Client
		err    error
	)
	JustBeforeEach(func() {
		client, err = NewSNOWClient(cfg)
	})

	Describe("NewSNOWClient", func() {
		It("requires a MarcosURL", func() {
			cfg := Config{}
			_, err := NewSNOWClient(cfg)
			Ω(err).Should(HaveOccurred())
		})

		Context("given a valid config", func() {
			BeforeEach(func() {
				URL, _ := url.Parse("http://abc.com")
				cfg = Config{
					URL,
				}
			})
			It("returns a valid client", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(client).ShouldNot(BeNil())
			})
		})
	})

})
