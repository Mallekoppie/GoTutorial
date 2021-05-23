package main

import (
	"bytes"
	"crypto/cipher"
	base64 "encoding/base64"
	"encoding/hex"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	os "os"
	exec "os/exec"
	"path"
	time "time"

	"gopkg.in/square/go-jose.v2/json"

	"bufio"

	"strings"

	cpu "github.com/shirou/gopsutil/cpu"

	mathrand "math/rand"

	rand "crypto/rand"

	"crypto/aes"
	"crypto/md5"
	"net/http"

	//	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/ssh/terminal"
	//	"sync"
	"context"

	"github.com/coreos/go-oidc"
)

func main() {
	//CpuTesting()
	//PlayWithMap()
	//PlayWithLogging()
	//PlayWithGoRoutines()
	//PlayWithNanoSeconds()
	//PlayWithTheConsole()
	//ReadUserInput()
	//ListFilesInFolder()
	//ConvertToBase64()
	//CleanFileNameDisplay()
	//CallJavaProcess()
	//CreateFolder()
	//MoveFile()
	//WriteTextToFile()
	//EncryptDecryptPassword()
	//EncryptStoreAndRead()
	//ReadJsonContractFileToString()
	//GetOAuth2Token()
	//ReadPassword()
	//TagTest()
	//PasswordSplit()
	//PlayWithMutex()
	//GetTokenRoles()
	//GetTokenClaimsTake2()
	//DeleteAllItemsInFolder()
	// CatchPanics()
	// queryEnvironment()
	// getRunningServices()
	// subString()
	// testFlags()
	dateToString()
}

func dateToString() {
	log.Println(time.Now())
}

func testFlags() {
	port := flag.String("port", "7777", "The port the service must listen on")

	flag.Parse()

	log.Println("The value: ", *port)
}

func subString() {
	test := "USERDNSDOMAIN=CAPITECBANK.FIN.SKY"
	result := test[strings.Index(test, "=")+1:]
	log.Println(result)
}

func getRunningServices() {
	command := exec.Command("sc", "query", "state=all")

	buffer := bytes.NewBufferString("")

	command.Stdout = buffer

	err := command.Run()
	if err != nil {
		log.Fatalln("Error executing command: ", err.Error())
	}

	if buffer.Len() > 0 {
		reader := bufio.NewScanner(buffer)

		for reader.Scan() {
			text := reader.Text()

			log.Println(text)
		}
	}
}

func queryEnvironment() {
	variables := os.Environ()

	log.Println("Environment variables")
	log.Println("===================================")
	for index := range variables {
		log.Println(variables[index])
	}

	log.Println("===================================")

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln("Unable to get hostname", err.Error())
	}
	log.Println("Hostname: ", hostname)
	log.Println("===================================")

}

func CatchPanics() {

	go DoWork()

	time.Sleep(time.Second * 2)

	log.Println("Executed")
}

func HandlePanic() {
	if err := recover(); err != nil {
		log.Println("Stil executed. Panic: ", err)
	}

}

func DoWork() {
	defer HandlePanic()
	panic("This must crash")
}

type TokenPayload struct {
	Roles []string `json:"roles,omitempty"`
	Jti   string   `json:"jti,omitempty"`
}

type Claims struct {
	Roles []string `json:"roles,omitempty"`
}

func GetTokenClaimsTake2() {
	token := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJzMTZWYTBzZXZVMHdlN0hmMjU3RzlRSFVzNmQ5WFZPX0ZnSXFGT3lBZC13In0.eyJqdGkiOiI1ZjVmNzM2NC0yYmUxLTRlMmEtYTBlZC03NmJhMGMyYjJjZDkiLCJleHAiOjE1NjU3MjY2MzksIm5iZiI6MCwiaWF0IjoxNTY1NzIzMDM5LCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgxODAvYXV0aC9yZWFsbXMvZ29sYW5nIiwiYXVkIjoiZ290dXRvcmlhbCIsInN1YiI6IjhkNzRmODY1LTNiMzAtNGE4Ni1hNGJkLWQ3YzcxNzdmNDNhYSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImdvdHV0b3JpYWwiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiIxYzQxMGFkMC03ZGYwLTQwNGItYWM1Yy1iMmNjZGE4ZGMyNzkiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInRlc3QiOiJzb21ldmFsdWUiLCJyb2xlcyI6WyJ1c2VyIl0sInByZWZlcnJlZF91c2VybmFtZSI6InRlc3QifQ.A80lrcXnl_LHIwV0vXDK90_1QhJXNr6Mnm3WV4xDUKpaRZrhi5i8oXo7ppHJhcbSU6DCjXEGsk4jPFAEauhYfO_lqa6_4XQXQyHz1UyN1tZeNE8Sfsgmintr0cSOUAUbK74Tp2_cntlppsCocCNmkkTCi6pL1IyijcGZLkm_SmP3dTY2F91ZzDH9P06TDc554bXvDOiT-wX0Vx_GPBLrmawiewPdovqaCUNumHDDMb3ZMyxoHYqqzwO3KYpBZWZQGdGBRtqSYvFHyNsSPU0Wg2j9Sc33HcODmSNB2U9dVuhTBZtsB0OP5biagk6krqZF4RWgnzPWkCzDSApdp482pQ"
	ctx := context.TODO()
	provider, err := oidc.NewProvider(ctx, "http://localhost:8180/auth/realms/golang") // this is bad
	if err != nil {
		panic(err)
	}

	oidcConfig := &oidc.Config{
		ClientID: "gotutorial", // this is bad
	}
	verifier := provider.Verifier(oidcConfig)

	parts := strings.Split(token, " ")

	idToken, err := verifier.Verify(ctx, parts[1])

	claims := Claims{}
	idToken.Claims(&claims)

	log.Println("Claims: ", claims)
}

func GetTokenRoles() {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJzMTZWYTBzZXZVMHdlN0hmMjU3RzlRSFVzNmQ5WFZPX0ZnSXFGT3lBZC13In0.eyJqdGkiOiI1ZjVmNzM2NC0yYmUxLTRlMmEtYTBlZC03NmJhMGMyYjJjZDkiLCJleHAiOjE1NjU3MjY2MzksIm5iZiI6MCwiaWF0IjoxNTY1NzIzMDM5LCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgxODAvYXV0aC9yZWFsbXMvZ29sYW5nIiwiYXVkIjoiZ290dXRvcmlhbCIsInN1YiI6IjhkNzRmODY1LTNiMzAtNGE4Ni1hNGJkLWQ3YzcxNzdmNDNhYSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImdvdHV0b3JpYWwiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiIxYzQxMGFkMC03ZGYwLTQwNGItYWM1Yy1iMmNjZGE4ZGMyNzkiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInRlc3QiOiJzb21ldmFsdWUiLCJyb2xlcyI6WyJ1c2VyIl0sInByZWZlcnJlZF91c2VybmFtZSI6InRlc3QifQ.A80lrcXnl_LHIwV0vXDK90_1QhJXNr6Mnm3WV4xDUKpaRZrhi5i8oXo7ppHJhcbSU6DCjXEGsk4jPFAEauhYfO_lqa6_4XQXQyHz1UyN1tZeNE8Sfsgmintr0cSOUAUbK74Tp2_cntlppsCocCNmkkTCi6pL1IyijcGZLkm_SmP3dTY2F91ZzDH9P06TDc554bXvDOiT-wX0Vx_GPBLrmawiewPdovqaCUNumHDDMb3ZMyxoHYqqzwO3KYpBZWZQGdGBRtqSYvFHyNsSPU0Wg2j9Sc33HcODmSNB2U9dVuhTBZtsB0OP5biagk6krqZF4RWgnzPWkCzDSApdp482pQ"
	split := strings.Split(token, ".")

	payload := split[1]
	decodedData, _ := base64.StdEncoding.DecodeString(payload)
	log.Println("Decoded Data: ", string(decodedData))

	result := TokenPayload{}
	json.Unmarshal(decodedData, &result)

	log.Println("Roles: ", result.Roles)
	log.Println("Jti: ", result.Jti)

}

func PasswordSplit() {
	password := "asdfj=asdfsdf==df"
	result := strings.Index(password, "=")
	value := password[result+1:]

	log.Println("Password: ", value)
}

func TagTest() {
	tags := []string{}

	log.Println("Original Tags: ", tags)

	javaTags := ""
	if len(tags) > 0 {
		for index := range tags {
			if index == 0 {
				javaTags = tags[index]
			} else {
				javaTags = javaTags + "," + tags[index]
			}
		}
	} else {
		// DEV deployment
		javaTags = "test"
	}

	log.Println("Tags: ", javaTags)

}

func ReadPassword() {
	log.Println("Enter password")
	bytes, e := terminal.ReadPassword(int(os.Stdin.Fd()))
	if e != nil {
		log.Println("Unable to read password: ", e.Error())
		return
	}

	log.Println("Password: ", string(bytes))
}

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
}

func GetOAuth2Token() {
	params := url.Values{}
	params.Add("client_id", "gotutorial")
	params.Add("client_secret", "ee297078-16dc-4908-8e3e-b4718506037f")
	params.Add("username", "test")
	params.Add("password", "test")
	params.Add("grant_type", "password")

	resp, err := http.PostForm("http://localhost:8180/auth/realms/golang/protocol/openid-connect/token", params)
	if err != nil {
		log.Println("Unable to get token: ", err.Error())
		return
	}
	defer resp.Body.Close()

	log.Println("Response: ", resp)
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Unable to read body: ", err.Error())
		return
	}

	token := OAuthResponse{}

	err = json.Unmarshal(bytes, &token)
	if err != nil {
		log.Println("Unmarshalling failed: ", err.Error())
		return
	}

	log.Println("Access Token: ", token.AccessToken)

}

func ReadJsonContractFileToString() {
	bytes, _ := ioutil.ReadFile("G:/logs/import-history/DEV/integration.loadtest.v3.contract.json")
	result := string(bytes)

	log.Println(result)
}

func EncryptStoreAndRead() {
	cryptoPassword := "hoopDitWerk"
	password := "toBeEncypted"

	result := encrypt([]byte(password), cryptoPassword)

	file, _ := os.Create("test")
	file.Write(result)

	resultRead, _ := ioutil.ReadFile("test")

	bytes := decrypt(resultRead, cryptoPassword)

	log.Println("Original: ", password)
	log.Println("Read from file: ", string(bytes))
}

func EncryptDecryptPassword() {
	cryptoPassword := "hoopDitWerk"
	password := "toBeEncypted"

	log.Println("Password unencrypted: ", password)
	result := encrypt([]byte(password), cryptoPassword)

	log.Println("Password encrypted binary: ", result)
	log.Println("Password encrypted string: ", string(result))

	decryptResult := decrypt(result, cryptoPassword)

	log.Println("Password decrypted binary: ", decryptResult)
	log.Println("Password decrypted string: ", string(decryptResult))
}

func decrypt(data []byte, passPhrase string) []byte {
	key := []byte(createHash(passPhrase))

	block, err := aes.NewCipher(key)

	if err != nil {
		log.Println("Unable to create cipher: ", err.Error())
		return nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Cannot create new gcm: ", err.Error())
		return nil
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println("Unable to decrypt data: ", err.Error())
		return nil
	}

	return plaintext
}

func encrypt(data []byte, passPhrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passPhrase)))

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		log.Println("Unable to create new GCM: ", err.Error())
		return nil
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println("unable to read nonce: ", err.Error())
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)

	return cipherText
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func WriteTextToFile() {
	// Hierdie werk net. Dit overwrite die existing file
	sampleData := "This is a bunch opf data that will be written in string format. This must overwrite the existing file"

	file, e := os.Create("G:/temp/test/created/newContract.txt")

	if e != nil {
		log.Println("Error opening destination file: ", e.Error())
		return
	}

	n, errWrite := file.WriteString(sampleData)

	if errWrite != nil {
		log.Println("Error writing data to file: ", errWrite.Error())
		return
	}

	log.Println("Number ofg bytes written successfully: ", n)

}

func MoveFile() {
	file, err := os.Create("G:/temp/test/created/result.txt")

	if err != nil {
		log.Println("Error creating destination file: ", err.Error())
		return
	}

	open, errOpen := os.Open("G:/temp/test/test.txt")

	if errOpen != nil {
		log.Println("Error opening source: ", errOpen.Error())
		return
	}

	io.Copy(file, open)

}

func CreateFolder() {
	// Jy kan net create call. As dit klaar exist gooi dit nie n error nie
	err := os.MkdirAll("created", 0755)

	if err != nil {
		log.Println("Create directory failed: ", err.Error())
	}
}

func DeleteAllItemsInFolder() {
	os.RemoveAll("created")
	os.MkdirAll("created", 755)
}

func CallJavaProcess() {
	myCmd := exec.Command("java", "-jar", "G:\\Code\\Source\\java\\java-tutorial\\gravitee-gateway-import\\target\\gravitee-contract-import-0.0.1-SNAPSHOT.jar",
		"admin", "admin", "admin", "9001", "deploy", "none", "test")
	myCmd.Stdout = os.Stdout
	myCmd.Stderr = os.Stderr
	err := myCmd.Run()

	if err != nil {
		log.Println("Error during execution: ", err.Error())
	}
}

func CleanFileNameDisplay() {
	filename := "connectiontest.json"
	result := strings.Index(filename, ".json")

	log.Println("Index: ", result)

	log.Println("Result: ", filename[0:result])
}

func RandomNumber() {
	sleepTime := mathrand.Intn(30) + 30

	log.Println("Random number: ", sleepTime)
}

func ConvertToBase64() {
	files, _ := ioutil.ReadDir("TestFolder")

	for i := 0; i < len(files); i++ {
		if strings.Contains(files[i].Name(), "result") == false {
			log.Println(path.Join("TestFolder", files[i].Name()))
			fileData, _ := ioutil.ReadFile(path.Join("TestFolder", files[i].Name()))
			log.Println(string(fileData))
			encodedData := base64.StdEncoding.EncodeToString(fileData)

			err := ioutil.WriteFile(path.Join("TestFolder", files[i].Name()+".result"), []byte(encodedData), os.ModeExclusive)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func ListFilesInFolder() {
	files, _ := ioutil.ReadDir("TestFolder")

	for i := 0; i < len(files); i++ {
		log.Printf(files[i].Name())
	}
}

func CpuTesting() {
	/*info, err := cpu.PerfInfo()

	if err != nil {
		log.Println("Cannot get CPu stat:", err)
	}

	for i := range info {

		log.Println(info[i].PercentUserTime)
	}
	*/
	data, err2 := cpu.Times(true)

	if err2 != nil {
		log.Println("Couldn't get cpu stats the second time:", err2)
	}

	for i := range data {
		if data[i].CPU == "_Total" {
			log.Println(data[i])
			log.Println(data[i].User)
		}

		//log.Println(data[i])
	}
}

func PlayWithMap() {
	headers := make(map[string]string)

	headers["firstKey"] = "firstValue"
	headers["secondKey"] = "secondValue"

	log.Println(headers)

	log.Println(headers["first"])
	log.Println(headers["second"])

	log.Println("looping over items")
	for i := range headers {

		log.Println(i, headers[i])
	}

	for headerKey, headerValue := range headers {
		log.Printf("HeaderKey: %v HeaderValue: %v", headerKey, headerValue)
	}
}

func PlayWithLogging() {
	value := "testValue"

	log.Printf("The value must be insterted here: %v : in the middle", value)
}

var (
	Users      map[int]bool
	UsersCount int = 0
)

func init() {
	Users = make(map[int]bool)
}

func GoRoutineFunc(id int) {
	for Users[id] == true {
		time.Sleep(time.Second * 2)
		log.Println("Go routine still running: ", id)
	}

	log.Println("Go routine shutting down: ", id)
}

func PlayWithGoRoutines() {
	var mapCount int
	for i := 0; i < 10; i++ {
		mapCount++
		Users[mapCount] = true
		go GoRoutineFunc(i)
		time.Sleep(time.Second * 1)

		log.Println("MapIndex: ", mapCount)
	}

	for i := 0; i < 10; i++ {
		mapCount--
		Users[mapCount] = false
		time.Sleep(time.Second * 1)

		log.Println("MapIndex: ", mapCount)
	}
}

func PlayWithNanoSeconds() {

}

func PlayWithTheConsole() {
	log.Println("bla")
	log.Println("blabla")

	time.Sleep(time.Second * 5)
	ClearOutput()
	log.Println("Cleared")
	time.Sleep(time.Second * 5)
}

func ClearOutput() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func ReadUserInput() {

	scanner := bufio.NewScanner(os.Stdin)
	log.Print("Enter command: ")
	//result, err := reader.ReadString('\lf')
	scanner.Scan()
	log.Println("Received test", scanner.Text())

}
