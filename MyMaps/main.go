package main

import "fmt"

func main() {
	fmt.Println("Maps in GOlang")

	Languages := make(map[string]string)

	Languages["JS"] = "Javascript"
	Languages["RB"] = "Ruby"
	Languages["PY"] = "Python"

	fmt.Println("List of all Languages: ", Languages)
	fmt.Println("JS is short for: ", Languages["JS"])

	delete(Languages, "RB")
	fmt.Println("List of all Languages: ", Languages)

	//lOOPS IN golang

	for key, value := range Languages {
		fmt.Printf("For key %v, Value is %v\n", key, value)
	}

	for _, value := range Languages {
		fmt.Printf("For key v, Value is %v\n", value)
	}

}
