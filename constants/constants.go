package constants

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Operation struct {
	ID   int64  `json:"id"`
	Slug string `json:"slug"`
}

var (
	// Constants : project constants
	Constants = GetConstants()
)

// Default : struct of constants
type Default struct {
	Operations struct {
		Spot         Operation `json:"spot"`
		Installments Operation `json:"installments"`
		Withdraw     Operation `json:"withdraw"`
		Payment      Operation `json:"payment"`
	}
}

// GetConstants : load the constants values from json file
func GetConstants() Default {
	var constants Default

	jsonConstants, _ := ioutil.ReadFile("constants/constants.json")
	if len(jsonConstants) == 0 {
		return constants
	}

	err := json.Unmarshal(jsonConstants, &constants)
	if err != nil {
		log.Fatal(err)
	}
	return constants
}
