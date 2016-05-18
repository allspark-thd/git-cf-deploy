package snow_test

import (
	"fmt"
	"net/url"

	. "github.com/allspark-thd/git-cf-deploy/lib/snow"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v0"
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

		Context("provided a valid change and invoked", func() {
			BeforeEach(func() {
				URL, _ := url.Parse("http://abc.com")
				cfg = Config{
					URL,
				}
			})
			It("returns a succesful response", func() {
				chg := Change{}
				URL := "http://abc.com"
				defer gock.Off()
				gock.New(URL).
					Post("").
					JSON(`{"I'm a change"}`).
					Reply(200).
					BodyString("totes")
				resp, err := client.Connect(chg)
				fmt.Println(resp, err)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(client).ShouldNot(BeNil())
			})
		})
	})
})
