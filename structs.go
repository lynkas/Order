package main

import "github.com/go-bongo/bongo"

const UNIT  = "unit"
const ORDER = "order"
const ITEM = "item"
const USER = "user"
const RESTAURANT = "restaurant"
//const NONE = "order"

type Restautant struct {
	bongo.DocumentBase `bson:",inline"`
	Name string `json:"name"`
	Tel string	`json:"tel"`
}

type Order struct {
	bongo.DocumentBase `bson:",inline"`
	Tel string `json:"tel"`
	Name string `json:"name"`
	Item []SingleItem `json:"items"`
	Annotation string `json:"annotation"  binding:"exists"`
	Processed bool `json:"processed" binding:"exists"`
}

type SingleItem struct {
	ItemName 	string `json:"item_name"`
	Quantity 	string	`json:"quantity"`
	Unit		string	`json:"unit"`
	Annotation 	string	`json:"annotation" binding:"exists"`
}

type Item struct {
	bongo.DocumentBase `bson:",inline"`
	Name string `json:"name"`
	DefaultUnit string `json:"default_unit"`
}

type Unit struct {
	bongo.DocumentBase `bson:",inline"`

	Name string `json:"name"`
}
type UserLogin struct {
	bongo.DocumentBase `bson:",inline"`
	Tel string `json:"tel"`
	Password string `json:"password"`
}

type OrderID struct {
	Id string `json:"id"`
} 