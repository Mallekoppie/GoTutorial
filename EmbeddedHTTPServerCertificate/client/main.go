package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
)

var (
	privateKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCjEWwQGxm0diEE
WkJ9hSZDLM8u473MYAFOUcNhHsvCOH/xvvp+gUJGYb8pdoic/bZBGF+yvpeaf21B
cGq+DNBabxOfa/jEMqSLEx1o2XrdmRzZ2QG6LGiT22p2LKrk7Er5qknf5X4e2k8/
IMgyMlvLmdVlB2qRQW/ld5Ob5U7VCNpr7hApW0UY8rqXDk0KuLRP1Zf8k+uZAv/q
grIpI08LBLlw25O6BY/9Q0KVfqH1FGsqQj8K9aZb1A1C/sct/Keqi25aElRqPYKB
Qz1IiDWjmYjNPBldDV2qRHvV97Z9ooskP2tXTYUtW81M3BMthSIwJerQlrycqy1G
j0CVmf+NAgMBAAECggEAC1hyOBGmoXFpnSupfiG2NozG0niacz6HwLr8GMUDgL0d
G8UBmS0lnw98dSZP3XpihJdtSxqrIh71C01Gw/fQxEX/q8MfoQjz/VAkZu/gtkvJ
n9eTIPCniwgEWXoEnt5Y3hOMxcTqGDvwsQ/3FdT+oYUPvZJ9ReNzZVgDP2C+MiCp
3UPQztVfgxMO1yrcLrmA/tlJJI3Wjh/HHXGCAEVjXcURM4Kc44cw5SWYEk727/iM
Lk/5ALlx4TG2YSJ3zwwLRb+acKGMtFVwozhXDFy1vOSZIAB7t/h35MXs3BnfHIzd
Wk0eIvgQtGdNK/u7xi4azh0p/faNDJLYva9TLrqvyQKBgQDTSojrVoALqmCrP+dY
4jO7xNMjtksEHDcG9OF+KfJLC2NOnvbSFn3FPhrTI0gtkqo3xxG6KvOxQSv+PsFH
nTZ0nH3vmYmCRTHn995Bbc5ViubrvN3w/aIP8atdJOr9z/hGW09+6dt2y+lppPQQ
Tmpubg5N6XMrhWspGnZmr8HBVQKBgQDFkq6zoVBKxOazb13mBPhcgCh6UaeneNkY
vQCNQ2Xd8EA/p8EcajCB3huNXK3/jbdgGXqGxdkF5bGrSU4iS+Flo0LHWK8JJMxx
pCl/siXgLMc3l9X4LRqpQ2/cbUWG0DkE0G8bHrUC8XRSbpM79c6dWDxL83IquGbO
yshkGxqlWQKBgBzljbI540s1IwFovPgf/5lHguTYcov+W9w8s3YwwG1ZAlznSrRS
1SKbqsmDgsSQ0WzNcfQ3GZr9YcnkxURJEQCDwol5QyAo3HlwIdpq7RQv6gMFkoj5
9ycdEyWq/eR2g+sDY+zkiNpHawXCxIpSXyX5MUWgjtq95Xj6oITMtr/9AoGAchD+
tgzDO3e0rhH1B1EFRvpi5EPyZHkb9fFguS4ZSOplXZrvngUc3k2qYJIKjwl1vpDX
MapWYtY7HmYjKx/eEP0SF992pvPM6StbHjNvNJ8FoW4TBcacpu6dTYj05CGhWcsw
6lpvLTTT8KeQb8S5thP0M/Foj4uJJ4jh6xHO73kCgYEAwX6qQjNz8T3tTYHWMtv4
8CttCMD9fO99UZp5FzDKydEh6yS3nbr12LK+ublgYvcfcF8GaajmXDyCtoI5Yjso
Kk4LnC/s6Oqy7B29QBUd96Ms4RSMHacYOa7UwarRojoLcTRrfKzY50+9OVw+3Ysm
ud3f4cRXvGUY09W6izzI6+M=
-----END PRIVATE KEY-----
`
	publicKey = `-----BEGIN CERTIFICATE-----
MIICqDCCAZCgAwIBAgIEX0zR7zANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtj
bGllbnQtY2VydDAeFw0yMDA4MzExMDMzMTlaFw0zMDA4MzExMDMzMTlaMBYxFDAS
BgNVBAMMC2NsaWVudC1jZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC
AQEAoxFsEBsZtHYhBFpCfYUmQyzPLuO9zGABTlHDYR7Lwjh/8b76foFCRmG/KXaI
nP22QRhfsr6Xmn9tQXBqvgzQWm8Tn2v4xDKkixMdaNl63Zkc2dkBuixok9tqdiyq
5OxK+apJ3+V+HtpPPyDIMjJby5nVZQdqkUFv5XeTm+VO1Qjaa+4QKVtFGPK6lw5N
Cri0T9WX/JPrmQL/6oKyKSNPCwS5cNuTugWP/UNClX6h9RRrKkI/CvWmW9QNQv7H
LfynqotuWhJUaj2CgUM9SIg1o5mIzTwZXQ1dqkR71fe2faKLJD9rV02FLVvNTNwT
LYUiMCXq0Ja8nKstRo9AlZn/jQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQB09iOD
j6pKKrf2OKGt/slxh2p+1zsT/y+El7VijwcQTYLHEvZ7A1ycWfbUISZ7fA6is8cy
31h9JeaPs3ogIujn5KIFIox/rTd4CvAY0CQ8gahnSTDs/jpJj9Cxd+ihLjte9EKA
bGXxfcZzqP/mXXsyyNgrrdCSMz9xP3hCn6A08HHiHmVGriilkHbbbdMDPhBQS/DM
6SUPBjElrKT2eWNmUYH/27cCh6RYfv1yqFIOrYZEHTfLIzZFLallBXADSwDlLzlG
40hBMzT91tE8wgKKAITEhlx0jmr5YFJQJgHtOSX/M4mDfqNywcOU95ff6FPGIp33
EdbjNowRU5Sw8qvO
-----END CERTIFICATE-----
`

	serverPublicKey = `-----BEGIN CERTIFICATE-----
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
		log.Fatalln("Unable to load certificate: ", err.Error())
	}

	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM([]byte(serverPublicKey))

	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caPool,
		InsecureSkipVerify: true,
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}

	resp, err := httpClient.Get("https://127.0.0.1:7777/")
	if err != nil {
		log.Fatalln("Error talking to server: ", err.Error())
	}

	log.Println("Response: ", resp)

}
