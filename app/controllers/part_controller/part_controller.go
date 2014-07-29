package part_controller

import (
	"github.com/curt-labs/acesfrontend/app/models/part_model"
	"github.com/martini-contrib/render"
	// "io/ioutil"
	"log"
	"net/http"
	"net/url"
	// "strings"
)

func Search(rw http.ResponseWriter, req *http.Request, ren render.Render) {
	data := make(map[string]interface{})
	var err error
	rw.Header().Add("Access-Control-Allow-Origin", "*")
	if err != nil {
		log.Print(err)
	}
	allMakes, err := part_model.GetAllMakes()
	allModels, err := part_model.GetAllModels()
	allYears, err := part_model.GetAllYears()
	allSubmodels, err := part_model.GetAllSubmodels()
	allConfigAttributes, err := part_model.GetAllConfigAttributes()
	data["makes"] = allMakes
	data["models"] = allModels
	data["years"] = allYears
	data["submodels"] = allSubmodels
	data["configAttributes"] = allConfigAttributes

	ren.HTML(200, "search", data)
}

func Get(rw http.ResponseWriter, req *http.Request, ren render.Render) {
	data := make(map[string]interface{})
	var err error
	rw.Header().Add("Access-Control-Allow-Origin", "*")
	// var parameters string

	theMake := req.FormValue("make")
	model := req.FormValue("model")
	year := req.FormValue("year")

	postData := url.Values{}
	postData.Add("make", "make")
	postData.Add("model", "model")
	postData.Add("year", "year")
	postData.Set("make", theMake)
	postData.Set("model", model)
	postData.Set("year", year)

	s := postData.Encode()

	data["output"] = part_model.Get(s)

	if err != nil {
		log.Print(err)
	}

	ren.HTML(200, "test", data)
}
