package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName        string             `json:"firstName"`
	Gender           string             `json:"gender"`
	HomeDistrict     string             `json:"homeDistrict"`
	MiddleName       string             `json:"middleName"`
	DOB              string             `json:"dob"`
	LastName         string             `json:"lastName"`
	StateOfDomicile  string             `json:"stateOfDomicile"`
	FatherFirstName  string             `json:"fatherFirstName"`
	FatherMiddleName string             `json:"fatherMiddleName"`
	FatherLastName   string             `json:"fatherLastName"`
	BoardName        string             `json:"boardName"`
	YearOfPassing    string             `json:"yearOfPassing"`
	RollNumber       string             `json:"rollNumber"`
	Address          string             `json:"address"`
	HouseNoVillage   string             `json:"houseNoVillage"`
	State            string             `json:"state"`
	District         string             `json:"district"`
	City             string             `json:"city"`
	PinCode          int                `json:"pinCode"`
}
