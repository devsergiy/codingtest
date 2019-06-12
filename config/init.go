package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spetrunin/codingtest/models"
)

var (
	DefaultPlans []models.Plan
	Rates        models.RatesMap
)

func init() {
	readJSON("config/plans.json", &DefaultPlans)
	readJSON("config/currency.json", &Rates)
}

func readJSON(file string, model interface{}) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, model)
	if err != nil {
		log.Fatal(err)
	}
}
