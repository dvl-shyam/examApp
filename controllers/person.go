package controllers

import (
	"context"
	"encoding/json"
	"examapp/config"
	"examapp/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection() *mongo.Collection {
    client, _ := config.ConnectDB()
    return client.Database("testdb").Collection("people")
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)

	collection := getCollection()
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	person.Gender = strings.ToLower(person.Gender)
	if person.Gender != "male" && person.Gender != "female" {
			http.Error(w, "Invalid gender", http.StatusBadRequest)
			return
	}

	result, err := collection.InsertOne(ctx, person)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	person.ID = result.InsertedID.(primitive.ObjectID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
	}

	var person models.Person
	collection := getCollection() 

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&person)
	if err != nil {
			fmt.Printf("Error finding person with ID %s: %v\n", id, err)
			
			if err == mongo.ErrNoDocuments {
					http.Error(w, "Person not found", http.StatusNotFound)
			} else {
					http.Error(w, "Error retrieving person", http.StatusInternalServerError)
			}
			return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func calculateAge(dob time.Time) int {
	currentTime := time.Now()
	age := currentTime.Year() - dob.Year()

	if currentTime.YearDay() < dob.YearDay() {
			age--
	}

	return age
}

func GetAge(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
	}

	var person models.Person
	collection := getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, map[string]interface{}{"_id": objID}).Decode(&person)
	if err != nil {
			http.Error(w, "Person not found", http.StatusNotFound)
			return
	}

	dob, err := time.Parse("2006-01-02", person.DOB)
	if err != nil {
			http.Error(w, "Invalid DOB format", http.StatusInternalServerError)
			return
	}

	age := calculateAge(dob)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]int{"Your Current Age is ": age}
	json.NewEncoder(w).Encode(response)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    var updatedPerson models.Person
    _ = json.NewDecoder(r.Body).Decode(&updatedPerson)

    collection := getCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

		updateDetails := bson.M{
			"$set": bson.M{
					"firstName":        updatedPerson.FirstName,
					"middleName":       updatedPerson.MiddleName,
					"lastName":         updatedPerson.LastName,
					"gender":           updatedPerson.Gender,
					"homeDistrict":     updatedPerson.HomeDistrict,
					"dob":              updatedPerson.DOB,
					"stateOfDomicile":  updatedPerson.StateOfDomicile,
					"fatherFirstName":  updatedPerson.FatherFirstName,
					"fatherMiddleName": updatedPerson.FatherMiddleName,
					"fatherLastName":   updatedPerson.FatherLastName,
					"boardName":        updatedPerson.BoardName,
					"yearOfPassing":    updatedPerson.YearOfPassing,
					"rollNumber":       updatedPerson.RollNumber,
					"address":          updatedPerson.Address,
					"houseNoVillage":   updatedPerson.HouseNoVillage,
					"state":            updatedPerson.State,
					"district":         updatedPerson.District,
					"city":             updatedPerson.City,
					"pinCode":          updatedPerson.PinCode,
			},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, updateDetails)
	if err != nil {
			http.Error(w, "Error updating document", http.StatusInternalServerError)
			return
	}

	if result.MatchedCount == 0 {
			http.Error(w, "No matching document found", http.StatusNotFound)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPerson)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeletePerson API hit")

	id := r.URL.Query().Get("id")
	if id == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			fmt.Println("Error: Invalid ID format", err)
			return
	}

	collection := getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Error deleting person:", err)
			return
	}

	if result.DeletedCount == 0 {
			http.Error(w, "Person not found", http.StatusNotFound)
			fmt.Println("No document found to delete with ID:", id)
			return
	}

	response := map[string]interface{}{
			"message": "Person deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}



