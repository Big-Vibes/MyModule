package main

import (
	"errors"
	"fmt"
) 

// Basic Error Example
func divide(a, b int) (int, error){
	if b == 0{
		return 0, errors.New("cannot be divide by 0")
	} 
	return a/b, nil
}

func main(){

	//Handling Error
	result , err := divide (4,0)
if err != nil{
	fmt.Println("error:", err)
	return
}
fmt.Println("result", result)


//Using fmt.Errorf
// return 0, fmt.Errorf("invalid value: %d",b)
	
}

