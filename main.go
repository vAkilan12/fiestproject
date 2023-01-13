package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/noun/v2/connection"
	"github.com/noun/v2/functions"
	"github.com/noun/v2/pkg"
)

func main() {
	decodeUsingInput()
}

func decodeUsingInput() {
	var search string

	fmt.Printf("Enter the Noun Name:")
	fmt.Scanf("%v", &search)

	Description, _ := connection.ExcelConnect("Text Match.xlsx", "Sheet1", 0)

	switch search {
	case "ADAPTER":
		functions.FindAdapter(Description, "ADAPTER")
	case "ELBOW":
		functions.FindElbow(Description, "ELBOW")
	case "BUSHING":
		functions.Bushing(Description)
	case "CLAMP":
		functions.Clamp(Description)
	case "COUPLING":
		functions.Coupling(Description)
	case "TEE":
		functions.Tee(Description)
	case "CROSS":
		functions.Cross(Description)
	case "FALINGE":
		functions.Falinge(Description)
	case "UNION":
		functions.Union(Description)
	default:
		fmt.Println("Not Found")
	}
}

func decodeUsingConfig() {
	config, _ := ioutil.ReadFile("config.json")
	var nounsToIdentify pkg.Nouns
	json.Unmarshal(config, &nounsToIdentify)

	descriptions, _ := connection.ExcelConnect("Text Match.xlsx", "Sheet1", 0)
	decoded := pkg.Decode(descriptions, nounsToIdentify)
	decodedBytes, _ := json.MarshalIndent(decoded, " ", "")
	jsonFile, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Write(decodedBytes)
}
