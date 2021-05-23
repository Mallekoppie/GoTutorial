package main

import (
	"encoding/json"
	"fmt"
	"log"

	"os"

	"github.com/boltdb/bolt"
)

var (
	DB_NAME            = "test.db"
	BUCKET_NAME_TEST   = "TestBucket"
	BUCKET_NAME_PEOPLE = "People"
)

func main() {
	// openDB()
	// createAndReadBucket()
	// storeComplexObjects()
	//returnAllInBucket()
	deleteObject()
}

func deleteObject(){
	db, err := bolt.Open(DB_NAME, os.ModeExclusive, nil)
	if err != nil {
		log.Fatalln("Error opening DB: ", err.Error())
		return
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME_TEST))
		if err != nil {
			log.Println("Error creating bucket: ", err.Error())
			return fmt.Errorf("create bucket: %s", err)
		}

		err = b.Put([]byte("one"), []byte("answer"))
		if err != nil {
			log.Fatalln("Error adding data: ", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalln("Error updating DB:", err.Error())
		return
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_TEST))
		result := b.Get([]byte("one"))
		if len(result) > 0 {
			log.Println("Returned data: ", string(result))
		} else {
			log.Fatalln("No data returned when reading")
		}

		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		buc := tx.Bucket([]byte(BUCKET_NAME_TEST))
		err := buc.Delete([]byte("one"))
		if err != nil {
			return err
		}

		return nil
	})

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_TEST))
		result := b.Get([]byte("one"))
		if len(result) > 0 {
			log.Println("Returned data: ", string(result))
		} else {
			log.Fatalln("No data returned when reading")
		}

		return nil
	})
	if err != nil {
		log.Fatalln("Error reading DB:", err.Error())
		return
	}
}

func returnAllInBucket() {
	results := getAll()

	log.Println("Results")
	for k, v := range results {
		log.Printf("%s %s", k, string(v))

	}
}

func getAll() map[string][]byte {
	results := make(map[string][]byte, 0)

	db, err := bolt.Open(DB_NAME, os.ModeExclusive, nil)
	if err != nil {
		log.Fatalln("Error opening DB: ", err.Error())
		return nil
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_PEOPLE))
		cursor := b.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			log.Printf("Key: %s, Value: %s", string(key), string(value))

			results[string(key)] = make([]byte, 0)
			copy(results[string(key)], value)
		}

		return nil
	})
	if err != nil {
		log.Fatalln("Error reading DB:", err.Error())
		return nil
	}

	return results
}

type Person struct {
	Name     string
	Surname  string
	Age      int
	IDNumber string
}

func storeComplexObjects() {
	db, err := bolt.Open(DB_NAME, os.ModeExclusive, nil)
	if err != nil {
		log.Fatalln("Error opening DB: ", err.Error())
		return
	}
	defer db.Close()

	idNumber := "92834572793"

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME_PEOPLE))
		if err != nil {
			log.Println("Error creating bucket: ", err.Error())
			return fmt.Errorf("create bucket: %s", err)
		}

		person := Person{
			Name:     "Hendrik",
			Surname:  "Bla",
			Age:      32,
			IDNumber: idNumber,
		}

		data, err := json.Marshal(person)
		if err != nil {
			log.Fatalln("Error marshalling person: ", err.Error())
			return err
		}

		err = b.Put([]byte(person.IDNumber), data)
		if err != nil {
			log.Fatalln("Error adding data: ", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalln("Error updating DB:", err.Error())
		return
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_PEOPLE))
		result := b.Get([]byte(idNumber))
		if len(result) > 0 {
			log.Println("Returned data: ", string(result))
			p := Person{}
			err := json.Unmarshal(result, &p)
			if err != nil {
				log.Fatalln("Unable to unmarshall data: ", err.Error())
				return err
			}

			log.Println("Retrieved person: ", p.Name)
		} else {
			log.Fatalln("No data returned when reading")
		}

		return nil
	})
	if err != nil {
		log.Fatalln("Error reading DB:", err.Error())
		return
	}
}

func createAndReadBucket() {
	db, err := bolt.Open(DB_NAME, os.ModeExclusive, nil)
	if err != nil {
		log.Fatalln("Error opening DB: ", err.Error())
		return
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME_TEST))
		if err != nil {
			log.Println("Error creating bucket: ", err.Error())
			return fmt.Errorf("create bucket: %s", err)
		}

		err = b.Put([]byte("one"), []byte("answer"))
		if err != nil {
			log.Fatalln("Error adding data: ", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalln("Error updating DB:", err.Error())
		return
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME_TEST))
		result := b.Get([]byte("one"))
		if len(result) > 0 {
			log.Println("Returned data: ", string(result))
		} else {
			log.Fatalln("No data returned when reading")
		}

		return nil
	})
	if err != nil {
		log.Fatalln("Error reading DB:", err.Error())
		return
	}
}

func openDB() {
	db, err := bolt.Open(DB_NAME, os.ModeExclusive, nil)
	if err != nil {
		log.Fatalln("Error opening DB: ", err.Error())
		return
	}
	defer db.Close()
}
