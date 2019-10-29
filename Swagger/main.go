package main

import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
)

// Reference: https://github.com/go-openapi/spec/blob/master/spec_test.go

func main() {
	//TestParsePaths()
	TestProcessSwagger()
}

func TestParsePaths() {
	removeParamRegex := "(\\{\\w+\\})|(\\{\\w+\\w+\\})|(\\{\\w+\\w+-\\w+\\})|(\\{\\w+\\w+\\w+\\w+\\})(\\{\\w+\\})"
	testPath := "configure/{test}/bla"
	paramReplaceValue := ":param"

	matched, err := regexp.MatchString(removeParamRegex, testPath)
	if err != nil {
		log.Println("Error during regex match: ", err.Error())
		return
	}

	if matched == true {
		log.Println("Regex matched!")
	}

	compiled := regexp.MustCompile(removeParamRegex)

	cleanPath := compiled.ReplaceAllString(testPath, paramReplaceValue)

	log.Println("New clean path: ", cleanPath)
}

func TestProcessSwagger() {
	rawDoc, err := swag.YAMLDoc("./ChaosAgent.json")
	if err != nil {
		log.Println("Error reading contract: ", err.Error())
		return
	}

	log.Println("Read the document: ", string(rawDoc))

	doc := new(spec.Swagger)

	err = json.Unmarshal(rawDoc, doc)
	if err != nil {
		log.Println("Unable to unMarshal doc:", err.Error())
		return
	}

	log.Println("UnMarshalled doc:", doc)

	// Get Definition
	testCollection := doc.Definitions["TestCollection"]

	result, err := testCollection.MarshalJSON()
	if err != nil {
		log.Println("unable to get test collection", err.Error())
		return
	}

	log.Println("TestCollection:", string(result))

	// Add additional properties to definition "additionalProperties"

	extraProps := make(map[string]interface{})
	extraProps["additionalProperties"] = false

	testCollection.ExtraProps = extraProps

	result, err = testCollection.MarshalJSON()
	if err != nil {
		log.Println("unable to get test collection", err.Error())
		return
	}

	log.Println("TestCollection with extra property:", string(result))

	//Get paths
	if doc.Paths != nil {
		log.Println("All the paths!!")
		for path := range doc.Paths.Paths {
			log.Println("Path: ", path)
			item := doc.Paths.Paths[path]

			if item.Get != nil {
				log.Println("Has Get")

				if item.Get.Responses != nil {
					for response := range item.Get.Responses.StatusCodeResponses {
						log.Println("Get possible response: ", response)

						detail := item.Get.Responses.StatusCodeResponses[response]

						log.Println("Response Detail: ", detail)

						if detail.Schema != nil {

							data, err := detail.Schema.MarshalJSON()
							if err != nil {
								log.Println("Response schema json marshal error: ", err.Error())
								return
							}

							log.Println("Response schema: ", string(data))

							ref, err := detail.Ref.MarshalJSON()
							if err != nil {
								log.Println("unable to marshal ref: ", err.Error())
								return
							}

							// This is empty. Just get the last index of '/' and
							// use the definition name to get it from definitions
							// on the root document
							log.Println("Schema detail ref: ", string(ref))
						}

					}

				}
			}

			if item.Put != nil {
				log.Println("Has Put")

				if item.Put.Responses != nil {
					for response := range item.Put.Responses.StatusCodeResponses {
						log.Println("Put possible response: ", response)
					}
				}
			}
		}
	}

}
