package snow_test

import (
	"net/url"

	. "github.com/allspark-thd/git-cf-deploy/lib/snow"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Client", func() {
	var (
		server *ghttp.Server
		cfg    Config
		client Client
		err    error
	)
	BeforeEach(func() {
		server = ghttp.NewServer()
		addr, _ := url.Parse(server.URL())
		cfg = Config{
			URL: addr,
		}
	})
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
			It("returns a valid client", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(client).ShouldNot(BeNil())
			})
		})

		Context("provided a valid change and invoked", func() {
			var chg Change

			JustBeforeEach(func() {
				chg = Change{}
				server.AppendHandlers(
					ghttp.VerifyRequest("POST", "/records"),
					ghttp.VerifyJSON(`{}`),
					ghttp.RespondWith(200, "totes"))
				err = client.Connect(chg)
			})

			It("returns a successful response", func() {
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("uses the endpoint provided to the client", func() {
				Ω(server.ReceivedRequests()).Should(HaveLen(1))
			})
		})
	})
})
