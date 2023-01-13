package pkg

import (
	"fmt"
	"log"
	"strings"

	"github.com/hbollon/go-edlib"
)

func GetType(descriptionArray []string, attributeList []string, name string) {
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
				//Column_D := "D" + strconv.Itoa(i+1)
				if edlib.JaroSimilarity(convToupper, name) >= 0.85 {
					splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
					for _, str2 := range splitAftercomma {
						splitAfterspace := strings.Fields(str2)
						for _, str3 := range splitAfterspace {
							convToupper := strings.ToUpper(strings.ReplaceAll(str3, " ", ""))

							for j := 0; j < len(attributeList); j++ {
								if edlib.JaroSimilarity(convToupper, attributeList[j]) >= 0.8 {
									// to find similarity between description and List and returing the similar type in list
									res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.5, edlib.Levenshtein)
									if err != nil {
										log.Println(err)
									}
									fmt.Println(" Type:", convToupper, res[0])
								}
							}
						}
					}
				}
			}
		}
	}
}

func FindType(descriptionArray []string, attributeList []string, descriptionKey int) string {
	var matchedType string
	splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
	for _, str2 := range splitAftercomma {
		splitAfterspace := strings.Fields(str2)
		for _, str3 := range splitAfterspace {
			convToupper := strings.ToUpper(strings.ReplaceAll(str3, " ", ""))

			for attributeListKey := 0; attributeListKey < len(attributeList); attributeListKey++ {
				if edlib.JaroSimilarity(convToupper, attributeList[attributeListKey]) >= 0.8 {
					// to find similarity between description and List and returing the similar type in list
					res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.5, edlib.Levenshtein)
					if err != nil {
						log.Println(err)
					}
					matchedType = res[0]
				}
			}
		}
	}
	return matchedType
}
