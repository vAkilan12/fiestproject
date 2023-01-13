package pkg

import (
	"fmt"
	"log"
	"strings"

	"github.com/hbollon/go-edlib"
)

func GetAngle(descriptionArray []string, attributeList []string, name string) {
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
							for attributeListKey := 0; attributeListKey < len(attributeList); attributeListKey++ {
								if convToupper == attributeList[attributeListKey] {
									// to find similarity between description andAttributeListand returing the similar angle in list
									res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.5, edlib.Levenshtein)
									if err != nil {
										log.Println(err)
									}
									//adding degree"°" symbol
									res1 := res[0] + "°"
									fmt.Println("Angle:", convToupper, res1)
								}

							}

						}
					}

				}
			}
		}
	}
}

func FindAngle(descriptionArray, attributeList []string, descriptionKey int) string {
	var matchedAngle string
	splitAftercomma := strings.Split(descriptionArray[descriptionKey], ",")
	for _, str2 := range splitAftercomma {
		splitAfterspace := strings.Fields(str2)
		for _, str3 := range splitAfterspace {
			convToupper := strings.ToUpper(strings.ReplaceAll(str3, " ", ""))
			for attributeListKey := 0; attributeListKey < len(attributeList); attributeListKey++ {
				if convToupper == attributeList[attributeListKey] {
					// to find similarity between description andAttributeListand returing the similar angle in list
					res, err := edlib.FuzzySearchSetThreshold(convToupper, attributeList, 1, 0.5, edlib.Levenshtein)
					if err != nil {
						log.Println(err)
					}
					//adding degree"°" symbol
					res1 := res[0] + "°"
					matchedAngle = res1
					fmt.Println("Angle:", convToupper, res1)
				}

			}
		}
	}
	return matchedAngle
}
