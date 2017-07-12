package fakes

import "github.com/pivotal-cf/om/api"

type CertificateAuthoritiesService struct {
	GetCall struct {
		CallCount int
		Returns   struct {
			CertificateAuthorities []api.CertificateAuthority
		}
	}
}

func (c *CertificateAuthoritiesService) Get() []api.CertificateAuthority {
	c.GetCall.CallCount++
	return c.GetCall.Returns.CertificateAuthorities
}
