package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDBName         string = "tutorial"
	MongoCollectionName string = "persons"
)

func main() {

	client, err := ConnectToMongo()
	if err != nil {
		return // we log the errors in the function
	}

	//	err = InsertPerson(client)
	//	if err != nil {
	//		return // we log the errors in the function
	//	}

	//	err = InsertManyPeople(client)
	//	if err != nil {
	//		return // we log the errors in the function
	//	}

	//	err = UpdatePersons(client)
	//	if err != nil {
	//		return // we log the errors in the function
	//	}

	//	err = UpdateOnePerson(client)
	//	if err != nil {
	//		return // we log the errors in the function
	//	}

	//	err = FindPerson(client)
	//	if err != nil {
	//		return // we log the errors in the function
	//	}

	//	err = FindAndDelete(client)
	//	if err != nil {
	//		return // we log the errors in the function
	//	}

	// err = UpdatePersonWithObject(client)
	// if err != nil {
	// 	return // we log the errors in the function
	// }

	// err = StoreAMap(client)
	// if err != nil {
	// 	return // we log the errors in the function
	// }

	err = ReadAMap(client)
	if err != nil {
		return // we log the errors in the function
	}

}

func ReadAMap(client *mongo.Client) error {
	collection := client.Database(MongoDBName).Collection("PlayingWithMaps")

	filter := bson.D{{"name", "Test Data Name"}}

	var store PlayStore

	singleResult := collection.FindOne(context.TODO(), filter)

	if singleResult != nil {
		err := singleResult.Decode(&store)
		if err != nil {
			log.Println("Unable to decode the person that we found: ", err.Error())
			return err
		}

		log.Println("The person that we found: ", store)
	} else {
		log.Println("Unable to find person")
		return nil
	}

	return nil
}

type PlayStore struct {
	Name string
	Data map[string]string
}

func StoreAMap(client *mongo.Client) error {
	collection := client.Database(MongoDBName).Collection("PlayingWithMaps")

	store := PlayStore{}

	store.Name = "Test Data Name"
	store.Data = make(map[string]string, 0)
	store.Data["key one"] = "value one"
	store.Data["key two"] = "value two"

	insertResult, err := collection.InsertOne(context.TODO(), store)
	if err != nil {
		log.Println("Unable to insert document:", err.Error())
		return nil
	}

	log.Println("Person Drikkie inserted:", insertResult.InsertedID)

	return nil
}

func ConnectToMongo() (client *mongo.Client, err error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Unable to connect to mongo:", err.Error())
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Unable to ping mongodb: ", err.Error())
		return nil, err
	}

	log.Println("Mongo connection successfull!")

	return client, nil
}

type Person struct {
	Name string
	Age  int
	City string
}

func InsertPerson(client *mongo.Client) error {
	collection := client.Database(MongoDBName).Collection(MongoCollectionName)

	drikkie := Person{"Drikkie", 31, "Grabouw"}

	insertResult, err := collection.InsertOne(context.TODO(), drikkie)
	if err != nil {
		log.Println("Unable to insert document:", err.Error())
		return nil
	}

	log.Println("Person Drikkie inserted:", insertResult.InsertedID)

	return nil
}

func InsertManyPeople(client *mongo.Client) error {
	collection := client.Database(MongoDBName).Collection(MongoCollectionName)

	drikkie2 := Person{"Drikkie2", 31, "Grabouw"}
	drikkie3 := Person{"Drikkie3", 31, "Grabouw"}

	people := []interface{}{drikkie2, drikkie3}

	insertResult, err := collection.InsertMany(context.TODO(), people)
	if err != nil {
		log.Println("Unable to save multiple people: ", err.Error())
		return err
	}

	log.Println("Added multiple people: ", insertResult)

	return nil

}

func UpdatePersons(client *mongo.Client) error {
	collection := client.Database(MongoDBName).Collection(MongoCollectionName)

	filter := bson.D{}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Println("Unable to update collection: ", err.Error())
		return err
	}

	log.Println("Update succeeded:", updateResult)
	return nil
}

func UpdateOnePerson(client *mongo.Client) error {
	collection := client.Database(MongoDBName).Collection(MongoCollectionName)

	filter := bson.D{{"name", "Drikkie3"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 2},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Unable to update collection: ", err.Error())
		return err
	}

	log.Println("Update succeeded:", updateResult)
	return nil
}

func FindPerson(client *mongo.Client) error {
	collection := client.Database(MongoDBName).Collection(MongoCollectionName)

	InsertPersonToFind := Person{Name: "findMe", Age: 0, City: "Search found the correct one"}

	insertResult, err := collection.InsertOne(context.TODO(), InsertPersonToFind)
	if err != nil {
		log.Println("Unable to insert person to find: ", err.Error())
		return nil
	}

	log.Println("Inserted person to find: ", insertResult)

	filter := bson.D{{"name", "findMe"}}

	var foundPerson Person

	singleResult := collection.FindOne(context.TODO(), filter)

	if singleResult != nil {
		err = singleResult.Decode(&foundPerson)
		if err != nil {
			log.Println("Unable to decode the person that we found: ", err.Error())
			return err
		}

		log.Println("The person that we found: ", foundPerson)
	} else {
		log.Println("Unable to find person")
		return nil
	}

	return nil
}

func FindAndDelete(client *mongo.Client) error {
	var deleteName string = "deleteMe"
	collection := client.Database(MongoDBName).Collection(MongoCollectionName)

	InsertPersonToFind := Person{Name: deleteName, Age: 0, City: "Search found the correct one"}

	insertResult, err := collection.InsertOne(context.TODO(), InsertPersonToFind)
	if err != nil {
		log.Println("Unable to insert person to delete: ", err.Error())
		return nil
	}

	log.Println("Inserted person to find: ", insertResult)

	deleteFilter := bson.D{{"name", deleteName}}

	deleteResult, err := collection.DeleteOne(context.TODO(), deleteFilter)
	if err != nil {
		log.Println("Error deleting: ", err.Error())
		return err
	}

	if deleteResult != nil {
		log.Println("Number of items delete: ", deleteResult.DeletedCount)
	}

	//Confirm that it is deleted
	findResult := collection.FindOne(context.TODO(), deleteFilter)

	if findResult != nil {
		if findResult.Err() == mongo.ErrNoDocuments {
			log.Println("Item no longer exists")
		}
	}

	return nil
}

//Failed to update with object. Doesn't look like it works like that
func UpdatePersonWithObject(client *mongo.Client) error {
	var updatename string = "updateMe"
	collection := client.Database(MongoDBName).Collection(MongoCollectionName)

	InsertPersonToFind := Person{Name: updatename, Age: 0, City: "Not Updated"}

	insertResult, err := collection.InsertOne(context.TODO(), InsertPersonToFind)

	if err != nil {
		log.Println("Unable to insert: ", err.Error())
		return nil
	}

	log.Println("inserted ID: ", insertResult.InsertedID)

	findToUpdateFilter := bson.D{{"name", updatename}}

	findResult := collection.FindOne(context.TODO(), findToUpdateFilter)

	if findResult != nil {
		if findResult.Err() != nil {
			log.Println("Find error: ", findResult.Err().Error())
			return findResult.Err()
		}

		updateFilter := bson.D{
			{"$set", bson.D{
				{"city", "This is updated"},
			}},
		}

		updateResult := collection.FindOneAndUpdate(context.TODO(), findToUpdateFilter, updateFilter)

		if updateResult != nil {
			if updateResult.Err() != nil {
				log.Println("Unabpe to update Person: ", updateResult.Err().Error())
				return updateResult.Err()
			}
			var finalPerson Person

			updateResult.Decode(&finalPerson)

			log.Println("Final Result: ", finalPerson)

			lastResult := collection.FindOne(context.TODO(), findToUpdateFilter)

			if lastResult != nil {
				if lastResult.Err() != nil {
					log.Println("Last find failed: ", lastResult.Err().Error())
					return lastResult.Err()
				}

				var lastLastPerson Person

				lastResult.Decode(&lastLastPerson)

				log.Println("Last Last result: ", lastLastPerson)
			}
		}

	}

	return nil

}
