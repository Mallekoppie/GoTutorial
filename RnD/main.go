package main

import (
	"crypto/cipher"
	base64 "encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"log"
	os "os"
	exec "os/exec"
	"path"
	"io"
	time "time"

	"bufio"

	"strings"

	cpu "github.com/shirou/gopsutil/cpu"

	mathrand "math/rand"

	rand "crypto/rand"

	"crypto/md5"
	"crypto/aes"
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
	EncryptDecryptPassword()
}

func EncryptDecryptPassword() {
	cryptoPassword := "hoopDitWerk"
	password:= "toBeEncypted"

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

func encrypt(data []byte, passPhrase string)[]byte {
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
	err := os.MkdirAll("G:/temp/test/created", 0755)

	if err != nil {
		log.Println("Create directory failed: ", err.Error())
	}
}

func CallJavaProcess(){
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
