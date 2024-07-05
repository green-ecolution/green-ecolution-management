package tree

type TreeLocation struct {
	Latitude       float64
	Longitude      float64
	Address        string
	AdditionalInfo string
}

type Tree struct {
	ID       string
	Species  string
	TreeNum  int
	Age      int
	Location TreeLocation
}
