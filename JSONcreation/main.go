package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	//--------------Alias in structExample``
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Handling Json")
	// encodeJson()
	DecodeJson()
}

func encodeJson() {
	TypeCat := []course{
		{"Black React", 222, "catty.vibes", "121212", []string{"web-dev", "js"}},
		{"White React", 2298, "catty.vibes", "12waa", []string{"Full stack", "js"}},
		{"brown React", 93233, "catty.vibes", "12weee", nil},
	}

	// package this data as JSON data

	// finalJson, err := json.Marshal(TypeCat)

	//Indented Output
	finalJson, err := json.MarshalIndent(TypeCat, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

// to Decode Json

func DecodeJson() {
	JsonDataFromWeb := []byte(`
	 {
                "coursename": "BlackReact",
                "Price": 222,
                "website": "catty.vibes",
                "tags": ["web-dev","js"]
        }
	`)

	//to check the data from the web
	var Vibescatty course
	checkValid := json.Valid(JsonDataFromWeb)
	if checkValid {
		fmt.Println("Json Was Valid")
		json.Unmarshal(JsonDataFromWeb, &Vibescatty)
		fmt.Printf("%#v \n", Vibescatty)
	} else {
		fmt.Println("json was not valid")
	}

	//some cases where you just want to add data to key value
	var MyOnlineData map[string]interface{}
	json.Unmarshal(JsonDataFromWeb, &MyOnlineData)
	fmt.Printf("%#v \n", MyOnlineData)

	for k, v := range MyOnlineData {
		fmt.Printf("Key is %v and Values is %v, and Type is: %T\n", k, v, v)
	}
}
