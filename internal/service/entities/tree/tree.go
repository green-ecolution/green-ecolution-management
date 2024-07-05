package tree

type TreeLocationResponse struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Address        string  `json:"address"`
	AdditionalInfo string  `json:"additional_info"`
} //@Name TreeLocation

type TreeResponse struct {
	ID       string               `json:"id"`
	Species  string               `json:"species"`
	TreeNum  int                  `json:"tree_num"`
	Age      int                  `json:"age"`
	Location TreeLocationResponse `json:"location"`
} //@Name Tree
