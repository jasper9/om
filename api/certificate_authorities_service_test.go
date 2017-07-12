package api_test

import (
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/api/fakes"
)

var _ = Describe("certificate authorities service", func() {
	var (
		certificateAuthoritiesService api.CertificateAuthoritiesService
		httpClient                    *fakes.HttpClient
	)

	BeforeEach(func() {
		httpClient = &fakes.HttpClient{}
		certificateAuthoritiesService = api.NewCertificateAuthorities(httpClient)
	})
	Describe("Get", func() {
		It("retrieves all the certificate authorities", func() {
			httpClient.DoReturns(&http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(strings.NewReader(`{"certificate_authorities":[
				   {
						 "guid": "some-guid",
						 "issuer": "some-issuer",
						 "created_on": "2017-06-17T03:03:32Z",
						 "expires_on": "2021-06-17T03:03:32Z",
						 "active": true,
						 "cert_pem": "cert perm"
					 }]}`)),
			}, nil)

			service := api.NewCertificateAuthoritiesService(httpClient)

			cas, err := service.Get()
			Expect(err).NotTo(HaveOccurred())
			Expect(cas).To(Equal([]api.CertificateAuthority{
				{
					GUID:      "some-guid",
					Issuer:    "some-issuer",
					CreatedOn: "2017-06-17T03:03:32Z",
					ExpiresOn: "2021-06-17T03:03:32Z",
					Active:    true,
					CertPem:   "cert perm",
				},
			}))

			request := client.DoArgsForCall(0)
			Expect(request.Method).To(Equal("GET"))
			Expect(request.URL.Path).To(Equal("/api/v0/certificate_authorities"))
		})
	})
})
