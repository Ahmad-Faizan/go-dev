package app

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	DOB  string `json:"dob"`
	Home string `json:"home"`
}

// This is the expected format in which the data is being exchanged.
// An example for the above JSON is defined below
//
// {
// 	"name":"Faizan Ahmad"
// 	"age" :"22"
// 	"dob" :"aaj abhi"
// 	"home":"har jgah"
// }
