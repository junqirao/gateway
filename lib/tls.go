package security

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// GenerateTlsCert ...
func GenerateTlsCert(noAfter time.Time, org, orgUnit []string, commonName string, ipAddresses []net.IP, dnsNames []string) (cert []byte, privateKey []byte, err error) {
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, err
	}

	subject := pkix.Name{
		Organization:       org,
		OrganizationalUnit: orgUnit,
		CommonName:         commonName,
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     noAfter,
		IsCA:         true,
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  ipAddresses,
		DNSNames:     dnsNames,
	}

	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	certBuf := &bytes.Buffer{}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	if err != nil {
		return
	}
	if err = pem.Encode(certBuf, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return
	}

	keyBuf := &bytes.Buffer{}
	if err = pem.Encode(keyBuf, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}); err != nil {
		return
	}
	return certBuf.Bytes(), keyBuf.Bytes(), nil
}

// GenerateAndSaveTlsCert ...
func GenerateAndSaveTlsCert(noAfter time.Time, org, orgUnit []string, commonName string, addresses []string, path string) error {
	certPath := path + "cert.pem"
	keyPath := path + "key.pem"
	var ipAddresses []net.IP
	var dnsNames []string

	for _, addr := range addresses {
		ip := net.ParseIP(addr)
		if ip == nil {
			// dns names
			dnsNames = append(dnsNames, addr)
		} else {
			// IP
			ipAddresses = append(ipAddresses, ip)
		}
	}

	cert, key, err := GenerateTlsCert(noAfter, org, orgUnit, commonName, ipAddresses, dnsNames)
	if err != nil {
		return err
	}
	certFile, err := os.OpenFile(certPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = certFile.Write(cert)
	if err != nil {
		return err
	}
	keyFile, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	_, err = keyFile.Write(key)
	if err != nil {
		return err
	}
	return nil
}
