package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
)

var Info struct{
	Admin string`json:"admin"`
	DatabaseUrl string `json:"database_url"`
	Database string `json:"database"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func basicInfo()  {
	jsonFile, err := os.Open("config.json")
	defer jsonFile.Close()
	if err!=nil {
		println("no config.json")
		os.Exit(1)
	}
	byteValue,_:=ioutil.ReadAll(jsonFile)
	err=json.Unmarshal(byteValue,&Info)
	if err!=nil {
		println("config err")
		os.Exit(1)
	}

}