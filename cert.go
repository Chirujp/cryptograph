package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"
)

type OneproviderServer struct {
	ServerId 				int 		`json:"server_id"`
	IpAddr 					string 	`json:"ip_addr"`
	Hostname 				string 	`json:"hostname"`
	Status 					string 	`json:"status"`
	Location 				string 	`json:"location"`
	NextDueDate 		string 	`json:"next_due_date"`
	BillingCycle 		string 	`json:"billing_cycle"`
	RecurringAmout	float32	`json:"recurring_amount"`
	CancelReason 		string	`json:"cancel_reason"`
	CancelType 			string 	`json:"cancel_type"`
	ManagerType 		string 	`json:"manager_type"`
	Ips 						string 	`json:"ips"`
}

type OneproviderApi struct {
	Result 					string 				`json:"result"`
	Response 				struct{
		Servers []OneproviderServer `json:"servers"`
	} 														`json:"response"`
}

func GetServers() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.oneprovider.com/server/list", nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", "API_nmraw2dzrVChXCVSzaVHSYtAI4yW8kix")
	req.Header.Set("Client-Key", "CK_jzaT274hXwJx7yDiQK1M9MlO032lMwx3")

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      log.Fatalln(err)
   }
   sb := string(body)
   log.Printf(sb)

	var oneproviderApi OneproviderApi

	if err := json.NewDecoder(res.Body).Decode(&oneproviderApi); err != nil {
		panic(err)
	}

	fmt.Println(oneproviderApi)
}

func generateCA() (string, string, error) {
	ca := &x509.Certificate{
		SerialNumber: 				big.NewInt(2019),
		Subject: pkix.Name{
			CommonName: 				"Chiru",
			Organization: 			[]string{"Chiru Acme"},
			Country: 						[]string{"FR"},
			Province: 					[]string{""},
			Locality: 					[]string{"Bordeaux"},
			StreetAddress: 			[]string{""},
			PostalCode: 				[]string{"33400"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)

	if err != nil {
		return "", "", err
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)

	if err != nil {
		return "", "", err
	}

	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type: "CERTIFICATE",
		Bytes: caBytes,
	})

	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})

	// Create Database certificate
	cert := &x509.Certificate{
		SerialNumber: 				big.NewInt(2019),
		Subject: pkix.Name{
			CommonName: 				"",
			Organization: 			[]string{"Chiru Acme"},
	}}

	fmt.Println(cert.Subject)

	return "", "", nil
}