package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			panic(err)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
			return nil;
	}
}

func main() {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SignatureAlgorithm: 10,
		Subject: pkix.Name{
			CommonName: "localhost",
			Organization: []string{"Chiru Acme"},
		},
		Version: 3,
		SerialNumber: big.NewInt(1),
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(time.Hour * 24 * 365),
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
	}


	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)

	if err != nil {
		panic(err)
	}


	blockCert := &pem.Block{
		Type: "CERTIFICATE",
		Bytes: derBytes,
	}

	fileCert, err := os.Create("cert.pem")

	if err != nil {
		panic(err)
	}

	if err := pem.Encode(fileCert, blockCert); err != nil {
		panic(err)
	}

	privCert, err := os.Create("priv.pem")

	if err != nil {
		panic(err)
	}

	if err := pem.Encode(privCert, pemBlockForKey(priv)); err != nil {
		panic(err)
	}
}