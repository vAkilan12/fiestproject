package pkg

import (
	"log"
	"strings"

	"github.com/hbollon/go-edlib"
	"github.com/noun/v2/output"
)

func GetNoun(descriptionArray []string, name string) {
	// Using for Loop
	for descriptionKey := 1; descriptionKey < len(descriptionArray); descriptionKey++ {
		// Spliting the description after "Comma"
		splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
		for _, str := range splitAftercomma {
			//Spliting the description after "Space"
			splitAfterspace := strings.Fields(str)
			for _, str1 := range splitAfterspace {
				// removing all the spaces
				// Converting description to uppercase
				convToupper := strings.ToUpper(strings.ReplaceAll(str1, " ", ""))

				if edlib.JaroSimilarity(convToupper, name) >= 0.85 {
					output.AdapterNouns = append(output.AdapterNouns, output.Adapter{
						Id:          descriptionKey,
						Description: descriptionArray[descriptionKey],
						NounName:    name,
					})
				}
			}
		}
	}
}

func FindName(descriptionArray, attributeList []string, descriptionKey int) string {
	var matchedName string
	// Spliting the description after "Comma"
	splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
	for _, str := range splitAftercomma {
		//Spliting the description after "Space"
		splitAfterspace := strings.Fields(str)
		for _, str1 := range splitAfterspace {
			// removing all the spaces
			// Converting description to uppercase
			convToupper := strings.ToUpper(strings.ReplaceAll(str1, " ", ""))
			for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {
				if edlib.JaroSimilarity(convToupper, attributeList[attributeListkey]) >= 0.85 {
					res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.5, edlib.Levenshtein)
					if err != nil {
						log.Println(err)
					}
					matchedName = res[0]
				}
			}
		}
	}
	return matchedName
}
