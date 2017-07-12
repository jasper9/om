package commands

import (
	"strings"

	"github.com/pivotal-cf/om/api"
)

type CertificateAuthorities struct {
	certificateAuthoritiesService certificateAuthoritiesService
	tableWriter                   tableWriter
}

type certificateAuthoritiesService interface {
	Get() []api.CertificateAuthority
}

func NewCertificateAuthorities(certificateAuthoritiesService certificateAuthoritiesService, tableWriter tableWriter) CertificateAuthorities {
	return CertificateAuthorities{
		certificateAuthoritiesService: certificateAuthoritiesService,
		tableWriter:                   tableWriter,
	}
}

func (c CertificateAuthorities) Execute() error {
	certificateAuthorities := c.certificateAuthoritiesService.Get()
	c.tableWriter.SetHeader([]string{"ID", "ISSUER", "ACTIVE", "CREATED ON", "EXPIRES ON", "CERTIFICATE PEM"})

	for _, ca := range certificateAuthorities {
		var active string
		if ca.Active {
			active = "true"
		} else {
			active = "false"
		}

		formattedCertPem := strings.Replace(ca.CertPem, "\\n", "\n", -1)

		c.tableWriter.Append([]string{ca.GUID, ca.Issuer, active, ca.CreatedOn, ca.ExpiresOn, formattedCertPem})
	}
	c.tableWriter.Render()
	return nil
}
