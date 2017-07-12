package acceptance

import (
	"net/http"
	"net/http/httptest"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = FDescribe("certificate-authorities", func() {
	It("prints a list of certificate authorities", func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			// w.Header().Set("Content-Type", "application/json")

			// switch req.URL.Path {
			// case "/uaa/oauth/token":
			// 	w.Write([]byte(`{
			// 	"access_token": "some-opsman-token",
			// 	"token_type": "bearer",
			// 	"expires_in": 3600
			// }`))
		}))

		command := exec.Command(pathToMain,
			"--target", server.URL,
			"--username", "some-username",
			"--password", "some-password",
			"--skip-ssl-validation",
			"certificate-authorities",
		)

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		Expect(session.Out).To(Equal(`+------------+----------+--------+------------+------------+---------------------------------------------------------+
| ID         | ISSUER   | ACTIVE | CREATED ON | EXPIRES ON | CERTIFICATE PEM                                         |
+------------+----------+--------+------------+------------+---------------------------------------------------------+
| some-guid  | Pivotal  | true   | 2017-01-09 | 2021-01-09 | -----BEGIN CERTIFICATE-----\nMIIC+zCCAeOgAwIBAgI....etc |
| other-guid | Customer | false  | 2017-01-10 | 2021-01-10 | -----BEGIN CERTIFICATE-----\nMIIC+zCCAeOgAwIBAgI....etc |
+------------+----------+--------+------------+------------+---------------------------------------------------------+
`))

	})
})
