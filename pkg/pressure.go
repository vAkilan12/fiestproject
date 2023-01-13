package pkg

import (
	"fmt"
	"log"
	"strings"

	"github.com/hbollon/go-edlib"
)

func GetPRating(descriptionArray []string, attributeList []string, name string) {
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
					splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
					for _, str2 := range splitAftercomma {
						splitAfterspace := strings.Fields(str2)
						for _, str3 := range splitAfterspace {
							convToupper := strings.ToUpper(strings.ReplaceAll(str3, " ", ""))
							for attributeListkey := 0; attributeListkey < len(attributeList); attributeListkey++ {
								if edlib.CosineSimilarity(convToupper, attributeList[attributeListkey], 1) >= 0.7 {
									// to find similarity between description and AttributeListand returing the similar pressure rating in list
									res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.5, edlib.Levenshtein)
									if err != nil {
										log.Println(err)
									}
									fmt.Println("Pressure:", convToupper, res[0])
								}
							}
						}
					}

				}
			}
		}
	}
}

func FindPR(descriptionArray, attributeList []string, descriptionKey int) string {
	var matchedPR string
	splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
	for _, str := range splitAftercomma {
		splitAfterspace := strings.Fields(str)
		for _, str3 := range splitAfterspace {
			convToupper := strings.ToUpper(strings.ReplaceAll(str3, " ", ""))
			for AttributeListkey := 0; AttributeListkey < len(attributeList); AttributeListkey++ {
				if edlib.CosineSimilarity(convToupper, attributeList[AttributeListkey], 1) >= 0.7 {
					// to find similarity between description and AttributeListand returing the similar pressure rating in list
					res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.5, edlib.Levenshtein)
					if err != nil {
						log.Println(err)
					}
					matchedPR = res[0]
				}
			}
		}
	}
	return matchedPR
}
