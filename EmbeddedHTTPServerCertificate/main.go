package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

var (
	privateKey = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDHP2pJKyOYBSrg
vPTWFGcbAagaWCfpOLHYTtVgGzVdrDxAI4rM1alE9vVqa//jDxIiu8Hr5YibzlrO
YF58Fxgyc9caEZInvP6PGfoK27Ul+WYYXdncbcp+HZdK4sG7VU2SzuVsTFU0pKFV
SKZwVh10IfftDVfkjFW1SECZDjy8zL+G2owwMyyYEKJ9Of+ZZ1lYaLj8ih6hXwST
kTzOouHtewtIXsQAEcB6j2gWkW19h0GWHStToXy6r86QeRpKK3/X/qMXmbN+/mPT
8N12lQV/SbDC/ZGOHqQBw7lon/kdMda9B3V7kobNCk8zIstvfgTz60diZh9zrzcu
w5imeQ0pAgMBAAECggEAN8WD8jwth+l2YZbyr5PhcnlJvSoM7ebNPT9/0Zl8HB7Q
dKGeQgzShzfbZTEa3MtEp81HkEjdLsOZvmfSOsjaIJE1Vhox/4aS+KpsR6rNL3Z6
gfqPN3jR8/BDf5YuZUoQ3zWjmWdaeFRWxoEB+bLZJr81Xzjbb+FqQkaZsl3+WVg1
T6T/vAvFnIP+cHCiVISNl6XhDbG4DlUY243HnH48kKUD53g7yymuWxhaXk/3PJI4
g4xi9gN6hwuneN2J2+4PRaKDwGU3xTm+nSSj63WE9V7a2Ew+pmvQMvdIUPhIrPkQ
a+pJd7vj5n+IDcSk5qf2NFNl3e4+j1gEbTBQxBaCgwKBgQD8lrVo//MHedv1vrT0
Ta7jzyFpWivotdJ1OAvlGljmJ0wo2EF58STlbMkQB/S0r03sLUVP4OWK5C2ZImnW
Gl7ILmfbmsiZBPgFfSQzc/IDr7L4QwC/eUfb6TpVoVdUNaGuv3r+8IuVAOspAGs0
dfYcJmzhmEZ2cs6JjaKkJ2D0swKBgQDJ8EmKhpAgbNBGcm9uSplJ7jzFqm/v6A7i
ZTtSfsQlXLO+ZFy9V9dWK00KGfP7kQDc7JL3sdxugXwDK5PEafc/928HlbMBwYRP
Et6KLWANLcumR4Ya4g6H3QoJyQUQeynKLy1eDvmUomklRdxfhRH5+juD0snLWk7d
PdBJvFw8swKBgBgmaK3jWt8qHYe/dhmtm70YPr+N6YjUaYzfkPOUs08+DqARHSGF
ltuArTClMhZcdHzST0A15nyDGID9s1TPLKTDGxb6E5fdy8DqmB17RnZnxBrMfTPR
ardx8yvgwzK/9fd5Q09AvHfAoLtI2PcpGCDU47AmGQaWL30jV5uOUEzVAoGAULSH
r6U31L2oP39cqWeG+9UK1LTZJ7hGupRasQ2YtpXmACF8Lu/8T6PeljrpF5FMFv81
fWjIACIfveScmdL/zcDFzvI1KG5+wlt4NSENUjcLPNk472WJCSN974s3Su7uNK/G
IeE6PfzqqMrS2BoGpTEst/J2U580BZe+trlg46ECgYAxlLtcTAZWuCePKEsb4BrB
3ek3L2rJRCNbyOmhKt/mN+NS7FDb2wZCbowtPResmQ4+sBToF70KEqSii2Ub4lgI
Kq+ywmGmPrKUUAQ80konDT9ViXlWe7qirX/VY/z13n1asRSihvH1P8dvYTkf4c8C
rVhZKnhwTPtdv1rCIrbbMw==
-----END PRIVATE KEY-----
`
	publicKey = `-----BEGIN CERTIFICATE-----
MIICpjCCAY6gAwIBAgIEX0ycfzANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDDAps
b2NhbC10ZXN0MB4XDTIwMDgzMTA2NDUxOVoXDTMwMDgzMTA2NDUxOVowFTETMBEG
A1UEAwwKbG9jYWwtdGVzdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AMc/akkrI5gFKuC89NYUZxsBqBpYJ+k4sdhO1WAbNV2sPEAjiszVqUT29Wpr/+MP
EiK7wevliJvOWs5gXnwXGDJz1xoRkie8/o8Z+grbtSX5Zhhd2dxtyn4dl0riwbtV
TZLO5WxMVTSkoVVIpnBWHXQh9+0NV+SMVbVIQJkOPLzMv4bajDAzLJgQon05/5ln
WVhouPyKHqFfBJORPM6i4e17C0hexAARwHqPaBaRbX2HQZYdK1OhfLqvzpB5Gkor
f9f+oxeZs37+Y9Pw3XaVBX9JsML9kY4epAHDuWif+R0x1r0HdXuShs0KTzMiy29+
BPPrR2JmH3OvNy7DmKZ5DSkCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAmEgU+Rlg
AuqjK3XrsAOh82vv4Wq8EPcbqp8GgTdXGKV+2dG2Q0WXCTqsutnePHIbyBdqpdar
u2PiigdNssP1BB5cZyHmU6an7O9wXYaWMAGOhOAnqYxZmGxFRYAe5QrCJuYwYPhw
YlTCee0oc/HJMYytGTxQgWlEeqn6t9cxyreGBpYtuHB43Ks0d0EqmajQjQMcekam
8LQb557hkOkUJrKmsNGs57QuHdU4acHc1cwjQHxKf29qn3lAVy/JCJ752vGtvRgr
Zy89fI555aS63JVvXZI6Eb76IenxYViDnhOKwG4FgENNwuSQksnFhRwPrgRcEFIW
Uidj0M9U2XNKhw==
-----END CERTIFICATE-----
`
)

func main() {
	cert, err := tls.X509KeyPair([]byte(publicKey), []byte(privateKey))
	if err != nil {
		log.Fatalln("Unable to load certificate key pair: ", err.Error())
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	server := http.Server{
		TLSConfig: tlsConfig,
		Addr:      "0.0.0.0:7777",
	}

	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalln("Error starting HTTP server: ", err.Error())
	}
}
