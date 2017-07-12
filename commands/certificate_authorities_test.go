package commands_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"
)

/*
X make a GET request to the Ops Manager server
3. write the data to a table
*/

const (
	ca1         = `-----BEGIN CERTIFICATE-----\nMIIE5DCCAsygAwIBAgIBATANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDEwdvbS10\nZXN0MB4XDTE3MDcxMTIzMDczN1oXDTI3MDcxMTIzMDc0N1owEjEQMA4GA1UEAxMH\nb20tdGVzdDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMh1NifWTBUt\nbWbN8qN7ZCawL/+cNS7v0wonFk04sCYg8Hr7WMfILNl9qwzS+Ae+FfPAygoChDrN\nj9oZBm4F8PgG79jIqkiQaVcDpj4C26P/xMrnLLGVS+jVUn/Uqv/WKJXrdcbpuIMI\nNRSh8ZpHhmo1KVim2lrOt0UiLcm7TdexTtGk5d8IYAMsWo2HFS7R5Ck/94VGj1kV\nlnl9BJ6FHh1AjAKvMODE9U4uqIh5ReOflC1vqIt0ZwfD0hUDnY9Jm/FlzKnH8SYg\nvy6a1aW+68L5OVk3KFJe8LrAe4GH323suSxchcTuH1nPuCmlGPEti/DV4ONsf4P2\n5r0Vz9W1NRIOlOvoJVMQjRldzlrb86QUDIpCq4R52roWHopiKhC8SFIRC8AHXFtY\njKYM/c5j7a6S9T9mtfnLlfKMX6zDUBPMU38h+m1I/VhCTzPNQjQ2a8Vu4wVevOjC\nfSEaaALakbUBUR1kmipI3fqS+S2fyPsqyBhfOixuOBE6lpxyDUqdcCiDRIXEtK5Y\nlXgu0s0QlEjROvqMacxd4l+62bSdXFcM8Qa3j74j8P/KkLdnRAdpzj4DHwBl/f9t\nYdpPSSNvfsEhQC19RwHHQOC3NifQsLUbdR1H7QoxOh7aOCByMjdMZ2LOC0Fkv8JA\nKUXNNdGiFXd1JJK09wPSQSo46rp6ZODpAgMBAAGjRTBDMA4GA1UdDwEB/wQEAwIB\nBjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBQnGz69wH/DMERFA8axLFOX\nvHY51TANBgkqhkiG9w0BAQsFAAOCAgEAxVyIXej8u2eUYX5afQUqYBNNV9mvN8Df\n4J9cM1bxBNKdi8VQWFUh7AWcI75EEBawi/VB4tHTw0pmsIYJuft3ocrbSMf4cKEP\ndC2roNfsUUB5sJP1emjIPAMK2Tnii9qBVvHXSi1IdaYARZnyv8zcluHgiUUMVaTa\nE2x0OBcv4aUrS5z2TY+cL22Y8lW8pw4rZBzZW7+MKEGUUEu4Q8x1y1uJcIscKJbV\n5XkvNcD6u6tzxuMWafg/UNA7lGVdm2PG+Qu46TUH/jZSNku7G7qIKuuD/QpctajT\nvyqXoRwEkl5AkDsscVm0UFpIXsnG7gdVrnNJm9fBNteN1ahbgVKhDSeoOnKgc5J7\nvfHjIqwg3Dmyt6fcLJ8aWilRzkIdB9lY7qRcEfdvXjLLQlz5HCTyyEkMuT72jDJF\neODl69MWzZr317hVUgpkugcl411g/NQqVrM95TM7uGSuKWRPOVVlfa+EymXk4Lhs\nqfctttArgEJbK0gTkZpzyE+PfbIiQl7CQmR8tykx/KXk+Xznb2VOr54QegVVTFB/\nRxNo4aNy5CFUqdSmynqGJCj4Xh5ySxJKWnRlVOb9bxZtQKiuUPA0zmDeuqyWMabb\nu0nfHGIRg7KnIrTxkWBO5iPUl1yOXTbXE2wds7xfs8k4q5Xlp9ZhQKSxnS//p8W6\nASFyLAYUxJY=\n-----END CERTIFICATE-----\n`
	expectedCA1 = `-----BEGIN CERTIFICATE-----
MIIE5DCCAsygAwIBAgIBATANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDEwdvbS10
ZXN0MB4XDTE3MDcxMTIzMDczN1oXDTI3MDcxMTIzMDc0N1owEjEQMA4GA1UEAxMH
b20tdGVzdDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMh1NifWTBUt
bWbN8qN7ZCawL/+cNS7v0wonFk04sCYg8Hr7WMfILNl9qwzS+Ae+FfPAygoChDrN
j9oZBm4F8PgG79jIqkiQaVcDpj4C26P/xMrnLLGVS+jVUn/Uqv/WKJXrdcbpuIMI
NRSh8ZpHhmo1KVim2lrOt0UiLcm7TdexTtGk5d8IYAMsWo2HFS7R5Ck/94VGj1kV
lnl9BJ6FHh1AjAKvMODE9U4uqIh5ReOflC1vqIt0ZwfD0hUDnY9Jm/FlzKnH8SYg
vy6a1aW+68L5OVk3KFJe8LrAe4GH323suSxchcTuH1nPuCmlGPEti/DV4ONsf4P2
5r0Vz9W1NRIOlOvoJVMQjRldzlrb86QUDIpCq4R52roWHopiKhC8SFIRC8AHXFtY
jKYM/c5j7a6S9T9mtfnLlfKMX6zDUBPMU38h+m1I/VhCTzPNQjQ2a8Vu4wVevOjC
fSEaaALakbUBUR1kmipI3fqS+S2fyPsqyBhfOixuOBE6lpxyDUqdcCiDRIXEtK5Y
lXgu0s0QlEjROvqMacxd4l+62bSdXFcM8Qa3j74j8P/KkLdnRAdpzj4DHwBl/f9t
YdpPSSNvfsEhQC19RwHHQOC3NifQsLUbdR1H7QoxOh7aOCByMjdMZ2LOC0Fkv8JA
KUXNNdGiFXd1JJK09wPSQSo46rp6ZODpAgMBAAGjRTBDMA4GA1UdDwEB/wQEAwIB
BjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBQnGz69wH/DMERFA8axLFOX
vHY51TANBgkqhkiG9w0BAQsFAAOCAgEAxVyIXej8u2eUYX5afQUqYBNNV9mvN8Df
4J9cM1bxBNKdi8VQWFUh7AWcI75EEBawi/VB4tHTw0pmsIYJuft3ocrbSMf4cKEP
dC2roNfsUUB5sJP1emjIPAMK2Tnii9qBVvHXSi1IdaYARZnyv8zcluHgiUUMVaTa
E2x0OBcv4aUrS5z2TY+cL22Y8lW8pw4rZBzZW7+MKEGUUEu4Q8x1y1uJcIscKJbV
5XkvNcD6u6tzxuMWafg/UNA7lGVdm2PG+Qu46TUH/jZSNku7G7qIKuuD/QpctajT
vyqXoRwEkl5AkDsscVm0UFpIXsnG7gdVrnNJm9fBNteN1ahbgVKhDSeoOnKgc5J7
vfHjIqwg3Dmyt6fcLJ8aWilRzkIdB9lY7qRcEfdvXjLLQlz5HCTyyEkMuT72jDJF
eODl69MWzZr317hVUgpkugcl411g/NQqVrM95TM7uGSuKWRPOVVlfa+EymXk4Lhs
qfctttArgEJbK0gTkZpzyE+PfbIiQl7CQmR8tykx/KXk+Xznb2VOr54QegVVTFB/
RxNo4aNy5CFUqdSmynqGJCj4Xh5ySxJKWnRlVOb9bxZtQKiuUPA0zmDeuqyWMabb
u0nfHGIRg7KnIrTxkWBO5iPUl1yOXTbXE2wds7xfs8k4q5Xlp9ZhQKSxnS//p8W6
ASFyLAYUxJY=
-----END CERTIFICATE-----
`
	ca2         = `-----BEGIN CERTIFICATE-----\nMIIE5DCCAiygAwIBAgIBATANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDEwdvbS10\nZXN0MB4XDTE3MDcxMTIzMDczN1oXDTI3MDcxMTIzMDc0N1owEjEQMA4GA1UEAxMH\nb20tdGVzdDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMh1NifWTBUt\nbWbN8qN7ZCawL/+cNS7v0wonFk04sCYg8Hr7WMfILNl9qwzS+Ae+FfPAygoChDrN\nj9oZBm4F8PgG79jIqkiQaVcDpj4C26P/xMrnLLGVS+jVUn/Uqv/WKJXrdcbpuIMI\nNRSh8ZpHhmo1KVim2lrOt0UiLcm7TdexTtGk5d8IYAMsWo2HFS7R5Ck/94VGj1kV\nlnl9BJ6FHh1AjAKvMODE9U4uqIh5ReOflC1vqIt0ZwfD0hUDnY9Jm/FlzKnH8SYg\nvy6a1aW+68L5OVk3KFJe8LrAe4GH323suSxchcTuH1nPuCmlGPEti/DV4ONsf4P2\n5r0Vz9W1NRIOlOvoJVMQjRldzlrb86QUDIpCq4R52roWHopiKhC8SFIRC8AHXFtY\njKYM/c5j7a6S9T9mtfnLlfKMX6zDUBPMU38h+m1I/VhCTzPNQjQ2a8Vu4wVevOjC\nfSEaaALakbUBUR1kmipI3fqS+S2fyPsqyBhfOixuOBE6lpxyDUqdcCiDRIXEtK5Y\nlXgu0s0QlEjROvqMacxd4l+62bSdXFcM8Qa3j74j8P/KkLdnRAdpzj4DHwBl/f9t\nYdpPSSNvfsEhQC19RwHHQOC3NifQsLUbdR1H7QoxOh7aOCByMjdMZ2LOC0Fkv8JA\nKUXNNdGiFXd1JJK09wPSQSo46rp6ZODpAgMBAAGjRTBDMA4GA1UdDwEB/wQEAwIB\nBjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBQnGz69wH/DMERFA8axLFOX\nvHY51TANBgkqhkiG9w0BAQsFAAOCAgEAxVyIXej8u2eUYX5afQUqYBNNV9mvN8Df\n4J9cM1bxBNKdi8VQWFUh7AWcI75EEBawi/VB4tHTw0pmsIYJuft3ocrbSMf4cKEP\ndC2roNfsUUB5sJP1emjIPAMK2Tnii9qBVvHXSi1IdaYARZnyv8zcluHgiUUMVaTa\nE2x0OBcv4aUrS5z2TY+cL22Y8lW8pw4rZBzZW7+MKEGUUEu4Q8x1y1uJcIscKJbV\n5XkvNcD6u6tzxuMWafg/UNA7lGVdm2PG+Qu46TUH/jZSNku7G7qIKuuD/QpctajT\nvyqXoRwEkl5AkDsscVm0UFpIXsnG7gdVrnNJm9fBNteN1ahbgVKhDSeoOnKgc5J7\nvfHjIqwg3Dmyt6fcLJ8aWilRzkIdB9lY7qRcEfdvXjLLQlz5HCTyyEkMuT72jDJF\neODl69MWzZr317hVUgpkugcl411g/NQqVrM95TM7uGSuKWRPOVVlfa+EymXk4Lhs\nqfctttArgEJbK0gTkZpzyE+PfbIiQl7CQmR8tykx/KXk+Xznb2VOr54QegVVTFB/\nRxNo4aNy5CFUqdSmynqGJCj4Xh5ySxJKWnRlVOb9bxZtQKiuUPA0zmDeuqyWMabb\nu0nfHGIRg7KnIrTxkWBO5iPUl1yOXTbXE2wds7xfs8k4q5Xlp9ZhQKSxnS//p8W6\nASFyLAYUxJY=\n-----END CERTIFICATE-----\n`
	expectedCA2 = `-----BEGIN CERTIFICATE-----
MIIE5DCCAiygAwIBAgIBATANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDEwdvbS10
ZXN0MB4XDTE3MDcxMTIzMDczN1oXDTI3MDcxMTIzMDc0N1owEjEQMA4GA1UEAxMH
b20tdGVzdDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMh1NifWTBUt
bWbN8qN7ZCawL/+cNS7v0wonFk04sCYg8Hr7WMfILNl9qwzS+Ae+FfPAygoChDrN
j9oZBm4F8PgG79jIqkiQaVcDpj4C26P/xMrnLLGVS+jVUn/Uqv/WKJXrdcbpuIMI
NRSh8ZpHhmo1KVim2lrOt0UiLcm7TdexTtGk5d8IYAMsWo2HFS7R5Ck/94VGj1kV
lnl9BJ6FHh1AjAKvMODE9U4uqIh5ReOflC1vqIt0ZwfD0hUDnY9Jm/FlzKnH8SYg
vy6a1aW+68L5OVk3KFJe8LrAe4GH323suSxchcTuH1nPuCmlGPEti/DV4ONsf4P2
5r0Vz9W1NRIOlOvoJVMQjRldzlrb86QUDIpCq4R52roWHopiKhC8SFIRC8AHXFtY
jKYM/c5j7a6S9T9mtfnLlfKMX6zDUBPMU38h+m1I/VhCTzPNQjQ2a8Vu4wVevOjC
fSEaaALakbUBUR1kmipI3fqS+S2fyPsqyBhfOixuOBE6lpxyDUqdcCiDRIXEtK5Y
lXgu0s0QlEjROvqMacxd4l+62bSdXFcM8Qa3j74j8P/KkLdnRAdpzj4DHwBl/f9t
YdpPSSNvfsEhQC19RwHHQOC3NifQsLUbdR1H7QoxOh7aOCByMjdMZ2LOC0Fkv8JA
KUXNNdGiFXd1JJK09wPSQSo46rp6ZODpAgMBAAGjRTBDMA4GA1UdDwEB/wQEAwIB
BjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBQnGz69wH/DMERFA8axLFOX
vHY51TANBgkqhkiG9w0BAQsFAAOCAgEAxVyIXej8u2eUYX5afQUqYBNNV9mvN8Df
4J9cM1bxBNKdi8VQWFUh7AWcI75EEBawi/VB4tHTw0pmsIYJuft3ocrbSMf4cKEP
dC2roNfsUUB5sJP1emjIPAMK2Tnii9qBVvHXSi1IdaYARZnyv8zcluHgiUUMVaTa
E2x0OBcv4aUrS5z2TY+cL22Y8lW8pw4rZBzZW7+MKEGUUEu4Q8x1y1uJcIscKJbV
5XkvNcD6u6tzxuMWafg/UNA7lGVdm2PG+Qu46TUH/jZSNku7G7qIKuuD/QpctajT
vyqXoRwEkl5AkDsscVm0UFpIXsnG7gdVrnNJm9fBNteN1ahbgVKhDSeoOnKgc5J7
vfHjIqwg3Dmyt6fcLJ8aWilRzkIdB9lY7qRcEfdvXjLLQlz5HCTyyEkMuT72jDJF
eODl69MWzZr317hVUgpkugcl411g/NQqVrM95TM7uGSuKWRPOVVlfa+EymXk4Lhs
qfctttArgEJbK0gTkZpzyE+PfbIiQl7CQmR8tykx/KXk+Xznb2VOr54QegVVTFB/
RxNo4aNy5CFUqdSmynqGJCj4Xh5ySxJKWnRlVOb9bxZtQKiuUPA0zmDeuqyWMabb
u0nfHGIRg7KnIrTxkWBO5iPUl1yOXTbXE2wds7xfs8k4q5Xlp9ZhQKSxnS//p8W6
ASFyLAYUxJY=
-----END CERTIFICATE-----
`
)

var _ = Describe("certificate authorities", func() {
	var (
		certificateAuthorities        commands.CertificateAuthorities
		certificateAuthoritiesService *fakes.CertificateAuthoritiesService
		tableWriter                   *fakes.TableWriter
	)

	BeforeEach(func() {
		tableWriter = &fakes.TableWriter{}
		certificateAuthoritiesService = &fakes.CertificateAuthoritiesService{}
		certificateAuthorities = commands.NewCertificateAuthorities(certificateAuthoritiesService, tableWriter)
	})

	Describe("Execute", func() {
		It("prints certificate authorities to a table", func() {
			certificateAuthoritiesService.GetCall.Returns.CertificateAuthorities = []api.CertificateAuthority{
				{
					GUID:      "some-random-guid",
					Issuer:    "some-issuer",
					CreatedOn: "a-timestamp",
					ExpiresOn: "another-time-stamp",
					CertPem:   ca1,
					Active:    false,
				},
				{
					GUID:      "some-other-random-guid",
					Issuer:    "some-issuer-2",
					CreatedOn: "a-timestamp-2",
					ExpiresOn: "another-time-stamp-2",
					CertPem:   ca2,
					Active:    true,
				},
			}

			err := certificateAuthorities.Execute()
			Expect(err).ToNot(HaveOccurred())

			Expect(certificateAuthoritiesService.GetCall.CallCount).To(Equal(1))

			Expect(tableWriter.SetHeaderCallCount()).To(Equal(1))
			Expect(tableWriter.SetHeaderArgsForCall(0)).To(Equal([]string{"ID", "ISSUER", "ACTIVE", "CREATED ON", "EXPIRES ON", "CERTIFICATE PEM"}))

			Expect(tableWriter.AppendCallCount()).To(Equal(2))
			Expect(tableWriter.AppendArgsForCall(0)).To(Equal([]string{
				"some-random-guid",
				"some-issuer",
				"false",
				"a-timestamp",
				"another-time-stamp",
				expectedCA1,
			}))
			Expect(tableWriter.AppendArgsForCall(1)).To(Equal([]string{
				"some-other-random-guid",
				"some-issuer-2",
				"true",
				"a-timestamp-2",
				"another-time-stamp-2",
				expectedCA2,
			}))

			Expect(tableWriter.RenderCallCount()).To(Equal(1))
		})
	})
})
