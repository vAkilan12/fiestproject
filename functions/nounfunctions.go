package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/hbollon/go-edlib"
	"github.com/noun/v2/output"
	"github.com/noun/v2/pkg"
)

// List of Noun Names
var Noun_List = []string{
	"ADAPTER",
	"BUSHING",
	"CLAMP",
	"COUPLING",
	"ELBOW",
	"TEE",
	"CROSS",
	"FALINGE",
	"UNION",
}

// List of Size
var SizeList = []string{
	`1/8"`,
	`1/4"`,
	`3/8"`,
	`1/2"`,
	`3/4"`,
	`0.75`,
	`1"`,
	`1-1/4"`,
	`1-1/2"`,
	`1.25`,
	`2"`,
	`2-1/2"`,
	`3"`,
	`3-1/2"`,
	`5"`,
	`6"`,
	`1.5"`,
	`8"`,
	`10"`,
	`12"`,
	`14"`,
	`16"`,
	`18"`,
	`20"`,
	`22"`,
	`24"`,
	`25"`,
	`26"`,
	`28"`,
	`30"`,
}

// List of material type
var Material_type_List = []string{
	"SS",
	"SS304",
	"SS316",
	"SS304L",
	"SS316L",
	"ABS",
	"ACTL",
	"BRS",
	"BRZ",
	"CI",
	"CPVC",
	"CS",
	"CU",
	"DPLX",
	"HDPE",
	"LDPE",
	"MS",
	"NY",
	"PE",
	"PES",
	"PFA",
	"PP",
	"PPSU",
	"PSU",
	"PTFE",
	"PUR",
	"PVC",
	"PVCU",
	"PVDF",
	"TI",
}

// list of pressure rating
var prating_List = []string{
	"SCH 5",
	"SCH 5S",
	"SCH 10",
	"SCH 10S",
	"SCH 20",
	"SCH 30",
	"SCH 40",
	"SCH 40S",
	"SCH 60",
	"SCH 80",
	"SCH 80S",
	"SCH 100",
	"SCH 120",
	"SCH 140",
	"SCH 160",
	"PN6",
	"PN10",
	"PN16",
	"PN25",
	"PN40",
	"PN64",
	"PN100",
}

// List of end connection
var endconnection_list = []string{
	"BRB",
	"SOC",
	"SPG",
	"T",
	"THD",
	"MT",
	"TU",
	"TB",
	"FT",
	"MALE",
	"FEMALETHD",
}

// List of angle
var angle_List = []string{
	`11.25`,
	`22.5`,
	`45`,
	`60`,
	`90`,
	`180`,
	`120`,
}

// list of type
var Type_List = []string{"STANDARD", "REDUCED"}

func Adapter(description []string) {

	pkg.GetNoun(description, "ADAPTER")
	pkg.GetMaterialType(description, Material_type_List, "ADAPTER")
	pkg.GetPRating(description, prating_List, "ADAPTER")
	pkg.GetEndConnection(description, endconnection_list, "ADAPTER")
	pkg.GetSize(description, SizeList, "ADAPTER")
	pkg.GetType(description, Type_List, "ADAPTER")

}

func FindAdapter(descriptionArray []string, name string) {
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
					Con1, Con2 := pkg.FindEnd(descriptionArray, endconnection_list, descriptionKey)
					Size1, Size2 := pkg.FindSize(descriptionArray, SizeList, descriptionKey)

					output.AdapterNouns = append(output.AdapterNouns, output.Adapter{
						Id:             descriptionKey,
						Description:    descriptionArray[descriptionKey],
						NounName:       name,
						MaterialType:   pkg.FindMaterial(descriptionArray, Material_type_List, descriptionKey),
						PressureRating: pkg.FindPR(descriptionArray, prating_List, descriptionKey),
						Nountype:       pkg.FindType(descriptionArray, Type_List, descriptionKey),
						EndConnection: []output.EndConnectionAttribute{
							{
								EndConn1: Con1,
								EndConn2: Con2,
							},
						},
						Size: []output.SizeAttribute{
							{
								EndSize1: Size1,
								EndSize2: Size2,
							},
						},
					})
				}
			}
		}
	}
	Output, err := json.MarshalIndent(output.AdapterNouns, " ", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nName output:%+v", string(Output))
	jsonFile, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Write(Output)
}

func Bushing(description []string) {

	pkg.GetNoun(description, "BUSHING")
	pkg.GetSize(description, SizeList, "BUSHING")
	pkg.GetMaterialType(description, Material_type_List, "BUSHING")
	pkg.GetType(description, Type_List, "BUSHING")
	pkg.GetPRating(description, prating_List, "BUSHING")
	pkg.GetEndConnection(description, endconnection_list, "BUSHING")

}

func Clamp(description []string) {

	pkg.GetNoun(description, "CLAMP")
	pkg.GetSize(description, SizeList, "CLAMP")
	pkg.GetMaterialType(description, Material_type_List, "CLAMP")
	pkg.GetType(description, Type_List, "CLAMP")
	pkg.GetPRating(description, prating_List, "CLAMP")
	pkg.GetEndConnection(description, endconnection_list, "CLAMP")

}

func Coupling(description []string) {

	pkg.GetNoun(description, "COUPLING")
	pkg.GetSize(description, SizeList, "COUPLING")
	pkg.GetMaterialType(description, Material_type_List, "COUPLING")
	pkg.GetType(description, Type_List, "COUPLING")
	pkg.GetPRating(description, prating_List, "COUPLING")
	pkg.GetEndConnection(description, endconnection_list, "COUPLING")

}

func Elbow(description []string) {

	pkg.GetNoun(description, "ELBOW")
	pkg.GetSize(description, SizeList, "ELBOW")
	pkg.GetMaterialType(description, Material_type_List, "ELBOW")
	pkg.GetType(description, Type_List, "ELBOW")
	pkg.GetPRating(description, prating_List, "ELBOW")
	pkg.GetEndConnection(description, endconnection_list, "ELBOW")
	pkg.GetAngle(description, angle_List, "ELBOW")

}

func FindElbow(descriptionArray []string, Name string) {
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
				if edlib.JaroSimilarity(convToupper, Name) >= 0.85 {
					con1, con2 := pkg.FindEnd(descriptionArray, endconnection_list, descriptionKey)
					size1, size2 := pkg.FindSize(descriptionArray, SizeList, descriptionKey)

					output.ElbowNouns = append(output.ElbowNouns, output.Elbow{
						Id:             descriptionKey,
						Description:    descriptionArray[descriptionKey],
						NounName:       Name,
						MaterialType:   pkg.FindMaterial(descriptionArray, Material_type_List, descriptionKey),
						PressureRating: pkg.FindPR(descriptionArray, prating_List, descriptionKey),
						Nountype:       pkg.FindType(descriptionArray, Type_List, descriptionKey),
						EndConnection: []output.EndConnectionAttribute{
							{
								EndConn1: con1,
								EndConn2: con2,
							},
						},
						Size: []output.SizeAttribute{
							{
								EndSize1: size1,
								EndSize2: size2,
							},
						},
						Angle: pkg.FindAngle(descriptionArray, angle_List, descriptionKey),
					})
				}
			}
		}
	}
	Output, err := json.MarshalIndent(output.ElbowNouns, " ", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nName output:%+v", string(Output))
	jsonFile, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Write(Output)
}

func Tee(description []string) {

	pkg.GetNoun(description, "TEE")
	pkg.GetSize(description, SizeList, "TEE")
	pkg.GetMaterialType(description, Material_type_List, "TEE")
	pkg.GetType(description, Type_List, "TEE")
	pkg.GetPRating(description, prating_List, "TEE")
	pkg.GetEndConnection(description, endconnection_list, "TEE")

}

func Cross(description []string) {

	pkg.GetNoun(description, "CROSS")
	pkg.GetSize(description, SizeList, "CROSS")
	pkg.GetMaterialType(description, Material_type_List, "CROSS")
	pkg.GetType(description, Type_List, "CROSS")
	pkg.GetPRating(description, prating_List, "CROSS")
	pkg.GetEndConnection(description, endconnection_list, "CROSS")

}

func Falinge(description []string) {

	pkg.GetNoun(description, "FALINGE")
	pkg.GetSize(description, SizeList, "FALINGE")
	pkg.GetMaterialType(description, Material_type_List, "FALINGE")
	pkg.GetType(description, Type_List, "FALINGE")
	pkg.GetPRating(description, prating_List, "FALINGE")
	pkg.GetEndConnection(description, endconnection_list, "FALINGE")

}

func Union(description []string) {

	pkg.GetNoun(description, "UNION")
	pkg.GetSize(description, SizeList, "UNION")
	pkg.GetMaterialType(description, Material_type_List, "UNION")
	pkg.GetType(description, Type_List, "UNION")
	pkg.GetPRating(description, prating_List, "UNION")
	pkg.GetEndConnection(description, endconnection_list, "UNION")

}
