package api

type CertificateAuthoritiesService struct {
}

func NewCertificateAuthorities(httpClient httpClient) CertificateAuthoritiesService {
	return CertificateAuthoritiesService{}
}

func (c CertificateAuthoritiesService) Get() []CertificateAuthority {
	return nil
}
