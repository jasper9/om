package api

type CertificateAuthority struct {
	GUID      string
	Issuer    string
	CreatedOn string
	ExpiresOn string
	Active    bool
	CertPem   string
}

/*
{
      "guid": "6bd3d87d2f10d183ae8b",
      "issuer": "Pivotal",
      "created_on": "2017-06-17T03:03:32Z",
      "expires_on": "2021-06-17T03:03:32Z",
      "active": true,
      "cert_pem": "-----BEGIN CERTIFICATE-----\nMIIDDzCCAfegAwIBAgIVAPZmAyIT9VSqvsgX75L/8j5YGS0xMA0GCSqGSIb3DQEB\nCwUAMB8xCzAJBgNVBAYTAlVTMRAwDgYDVQQKDAdQaXZvdGFsMB4XDTE3MDYxNjAz\nMDMzMloXDTIxMDYxNzAzMDMzMlowHzELMAkGA1UEBhMCVVMxEDAOBgNVBAoMB1Bp\ndm90YWwwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDcwQK2yTlkV5r4\nDVraXqTdL342bmTg3BCyPNPyNdTysNAZJ5Dccqx7ZAWZIHX62SL6GLeRmCS/c6VI\niG/Ij/QbM/5I/FW2xCMP8dR+CW8n041fadB/jQN8ywFLgR8ZhwWw5LMcv1j5Ksvl\nV9QMEfiLxeYjBbUsptPO7xCEjLnMriS6/B8WZfxXFx7veqmedZJG8sQMwAMRBoq6\n4J7TwiPWtiiNK2ij3y8aurxxoznqtgs9uqf+kbIZnk7/U0jlw0OYNXjOKudl1Ete\noGPKSmys4HwHUh0R8OYyCvxopKKvpyg3LEB53wW1nLgybPWvuQfcWDdg2Ns8mHv3\n3+S+9sVlAgMBAAGjQjBAMB0GA1UdDgQWBBQWmNQOntIFD5SGAIZ9kVdSPRp3fDAP\nBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBBjANBgkqhkiG9w0BAQsFAAOC\nAQEAj2YNv/82oHGC7Y5gLbPtI9zw78gjOplnPEmQ//U6DfjwtJo4L+w+fF8rN7rk\nGj94KYKZ/lGp/lgOmd+VCBAYpD/rK99WuOkTNZOV4RPb4ex/aQWEyvCRpeflltKo\nmwX010/g//dsVmDtLxtMLUUkSgMR+qF+tTeGQSTQqbn1e1Z1x1b53345Cgp6s7kU\n51GF4R9gixDKirfxnx5ortYvT6hjzoHG8PdHjEPsq8u7/+bGiaNDh2+ccgMBAXsU\nCxfl28I7ZhAH986IXT5ueDv70V5l6bgCp76okAdsKhUJ+GLPS9/51P0Vs1//UvjP\nX0b/wISH+d98PATgFQeDGe9vEg==\n-----END CERTIFICATE-----\n"
    }
*/
