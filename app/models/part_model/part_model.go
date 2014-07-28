package part_model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Vehicle struct {
	ID               int
	ConfigID         int
	AppID            int
	Region           Region
	Model            Model
	Make             Make
	Year             Year
	Submodel         Submodel
	ConfigAttributes []ConfigAttribute
}

type Model struct {
	ID            int
	AAIAModelID   int
	Name          string
	VehicleTypeID int
}

type Make struct {
	ID         int
	AAIAMakeID int
	Name       string
}
type Year struct {
	ID   int
	Name int
}

type Submodel struct {
	ID             int
	AAIASubmodelID int
	SubmodelName   string
}

type Region struct {
	ID   int
	Name string
}

type ConfigAttribute struct {
	ID                  int
	ConfigAttributeType ConfigAttributeType
	parentID            int
	vcdbID              int
	Name                string
}

type ConfigAttributeType struct {
	ID       int
	Name     string
	AcesType AcesType
	Sort     int
}

type AcesType struct {
	ID   int
	Name string
}

func Get(formData string) []Vehicle {
	client := &http.Client{CheckRedirect: nil}

	log.Print("FROM: ", formData)

	req, err := http.NewRequest("POST", "http://localhost:3000/vehicle/params", bytes.NewBufferString(formData))

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		log.Println(err)
	}
	resp.Body.Close() //close resp.Body when done reading from it

	var data []Vehicle
	err = json.Unmarshal(body, &data)
	if err == nil {
		log.Println(err)
	}

	if err != nil {
		log.Print(err)
	}
	return data
}

func GetAllMakes() ([]Vehicle, error) {
	var vs []Vehicle
	var err error
	resp, err := http.Get("http://localhost:3000/make/all")
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	// log.Print("BODY: ", body)
	if err != nil {
		log.Print("Err Readalling: ", err)
	}
	err = json.Unmarshal(body, &vs)

	return vs, err

}
func GetAllModels() ([]Vehicle, error) {
	var vs []Vehicle
	var err error
	resp, err := http.Get("http://localhost:3000/model/all")
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	// log.Print("BODY: ", body)
	if err != nil {
		log.Print("Err Readalling: ", err)
	}
	err = json.Unmarshal(body, &vs)

	return vs, err

}
