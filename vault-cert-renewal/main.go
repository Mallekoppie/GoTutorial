package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/vault/api"
	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
	"log"
	"net"
	"net/http"
	"net/url"
)

func main() {
	generateCertificate()

	//startHttpServerAndRequestCert()

}

func startHttpServerAndRequestCert() {
	issuer := &vault.Issuer{
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:8200",
		},
		Role:       "server",
		AuthMethod: vault.ConstantToken("s.Gs83LqwPCB8NGcMdMSA2HezI"),
	}

	c := certify.Certify{
		CommonName: "local.madsolutions.net",
		Issuer:     issuer,
		Cache:      certify.NewMemCache(),
	}

	tlsConfig := &tls.Config{
		GetCertificate: c.GetCertificate,
	}

	s := &http.Server{
		Addr:      "0.0.0.0:9001",
		TLSConfig: tlsConfig,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Println("Error starting server: ", err.Error())
	}

}

func generateCertificate() {
	config := &api.Config{
		Address:          "http://localhost:8200",
		AgentAddress:     "",
		HttpClient:       nil,
		MaxRetries:       0,
		Timeout:          0,
		Error:            nil,
		Backoff:          nil,
		Limiter:          nil,
		OutputCurlString: false,
	}

	client, err := api.NewClient(config)
	if err != nil {
		log.Println("Unable to create new client: ", err.Error())
		return
	}

	client.SetToken("s.Gs83LqwPCB8NGcMdMSA2HezI")

	logical := client.Logical()

	secret, err := logical.Write("pki/issue/server", map[string]interface{}{
		"common_name": "local.madsolutions.net",
		"ttl":         "24h",
	})

	if err != nil {
		log.Println("Issue generating new certificate: ", err.Error())
		return
	}

	for i := range secret.Data {
		log.Println("Index: ", i)
		log.Println("Data: ", secret.Data[i])
	}
	certString := secret.Data["certificate"].(string)
	keyString := secret.Data["private_key"].(string)
	result := fmt.Sprintf("%s\n%s", certString, keyString)

	log.Println(result)

	block, rest := pem.Decode([]byte(result))

	log.Println(block)
	log.Println(rest)

	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Println("Error making x509 cert: ", err.Error())
		return
	}

	log.Println(certificate)

	//pair, err := tls.X509KeyPair([]byte(certString), []byte(keyString))
	//if err != nil {
	//	log.Println("Error loading key pair: ", err.Error())
	//	return
	//}

	tlsConfig := &tls.Config{
		//Certificates:             []tls.Certificate{pair},
		GetCertificate: func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
			pair, err := tls.X509KeyPair([]byte(certString), []byte(keyString))
			if err != nil {
				log.Println("Error loading key pair: ", err.Error())
			}
			return &pair, nil
		},

		MinVersion:                  0,
		MaxVersion:                  0,
		CurvePreferences:            nil,
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}

	myServer:=http.Server{
		Addr:              "0.0.0.0:9001",
		TLSConfig:         tlsConfig,
	}

	myMux := mux.NewRouter()
	myMux.HandleFunc("/", hello)
	myServer.Handler = myMux

	myServer.ListenAndServeTLS("","")

	conn, err := net.Listen("tcp", myServer.Addr)
	if err != nil {
		log.Fatal(err)
	}

	tlsListener := tls.NewListener(conn, tlsConfig)
	myServer.Serve(tlsListener)

}

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}