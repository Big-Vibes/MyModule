package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in GOlang")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "friday", "saturday"}
	days = append(days, "thursday")
	// sort.Strings(days)
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index+1, day)
	// fmt.Println(index+1, day)
	// }

	// for _, day := range days {
	// fmt.Printf("index is and value is %v \n", day)
	// fmt.Println(day)
	// }

	roughValue := 1
	for roughValue < 10 {
		// if roughValue == 5 {
		// 	break
		// }
		if roughValue == 2 {
			goto lco
		}
		if roughValue == 5 {
			roughValue++
			continue
		}

		fmt.Println("value is ", roughValue)
		roughValue++
	}

	//
lco:
	fmt.Println("Jumping at afafewf")
}
