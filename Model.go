package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName        string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName       string             `json:"middleName,omitempty" bson:"middleName,omitempty"`
	LastName         string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Gender           string             `json:"gender,omitempty" bson:"gender,omitempty"`
	HomeDistrict     string             `json:"homeDistrict,omitempty" bson:"homeDistrict,omitempty"`
	DOB              string             `json:"dob,omitempty" bson:"dob,omitempty"`
	StateOfDomicile  string             `json:"stateOfDomicile,omitempty" bson:"stateOfDomicile,omitempty"`
	FatherFirstName  string             `json:"fatherFirstName,omitempty" bson:"fatherFirstName,omitempty"`
	FatherMiddleName string             `json:"fatherMiddleName,omitempty" bson:"fatherMiddleName,omitempty"`
	FatherLastName   string             `json:"fatherLastName,omitempty" bson:"fatherLastName,omitempty"`
	BoardName        string             `json:"boardName,omitempty" bson:"boardName,omitempty"`
	YearOfPassing    string             `json:"yearOfPassing,omitempty" bson:"yearOfPassing,omitempty"`
	RollNumber       string             `json:"rollNumber,omitempty" bson:"rollNumber,omitempty"`
	Address          string             `json:"address,omitempty" bson:"address,omitempty"`
	HouseNoVillage   string             `json:"houseNoVillage,omitempty" bson:"houseNoVillage,omitempty"`
	State            string             `json:"state,omitempty" bson:"state,omitempty"`
	District         string             `json:"district,omitempty" bson:"district,omitempty"`
	City             string             `json:"city,omitempty" bson:"city,omitempty"`
	PinCode          int                `json:"pinCode,omitempty" bson:"pinCode,omitempty"`
}
