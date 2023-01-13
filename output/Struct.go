package output

type Adapter struct {
	Id             int                      `json:"_id"`
	Description    string                   `json:"description"`
	NounName       string                   `json:"noun_name"`
	MaterialType   string                   `json:"material_type"`
	PressureRating string                   `json:"pressure_rating"`
	Nountype       string                   `json:"noun_type"`
	EndConnection  []EndConnectionAttribute `json:"end_connection"`
	Size           []SizeAttribute          `json:"size"`
}
type Elbow struct {
	Id             int                      `json:"_id"`
	Description    string                   `json:"description"`
	NounName       string                   `json:"noun_name"`
	MaterialType   string                   `json:"material_type"`
	PressureRating string                   `json:"pressure_rating"`
	Nountype       string                   `json:"noun_type"`
	EndConnection  []EndConnectionAttribute `json:"end_connection"`
	Size           []SizeAttribute          `json:"size"`
	Angle          string                   `json:"angle"`
}

var AdapterNouns []Adapter
var ElbowNouns []Elbow

type SizeAttribute struct {
	EndSize1 string `json:"end_size1"`
	EndSize2 string `json:"end_size2"`
}

type EndConnectionAttribute struct {
	EndConn1 string `json:"end_conn1"`
	EndConn2 string `json:"end_conn2"`
}
