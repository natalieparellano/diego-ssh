package handlers_test

import (
	"github.com/cloudfoundry-incubator/diego-ssh/keys"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/ssh"

	"testing"
)

var TestHostKey ssh.Signer

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

var _ = BeforeSuite(func() {
	hostKey, err := keys.NewRSA(1024)
	Ω(err).ShouldNot(HaveOccurred())

	TestHostKey = hostKey.PrivateKey()
})
