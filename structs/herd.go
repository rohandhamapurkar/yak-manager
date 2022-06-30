package structs

type Herd struct {
	Herd []Yak `json:"herd"`
}

// yak struct which contains yak fields
type Yak struct {
	Name          string  `json:"name"`
	Age           float32 `json:"age"`
	Sex           string  `json:"sex"`
	AgeLastShaved float32
}
