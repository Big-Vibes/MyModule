package main

import "fmt"
//Basic Function
func greet() {
fmt.Println("Hello from Go")
}
//Function with Parameters
func greetUser(name string){
    fmt.Println("name", name)
}
//Function with Return Value
func add( a float64, b float64) float64{
    return (a + b)
}
//Multiple Return Values (Very Go-like)
func calculate(a int, b int) (int,int){
    sum := a + b
    subtract := a - b
    return sum , subtract
}
 //Named Return Values
func rectangle(length, breath int) (area int){
    area = length * breath
    return
}

func moreCal(a, b int) (int , int){
    return a * b, a + b
}

func main() {
    greet()
    greetUser("Abiodun")

    result := add(100.0 , 3.1)
    fmt.Println(result)

    sum, subtract := calculate(5,37)
    fmt.Println(sum, subtract)

    rectangle(3,3)

    mult, add := moreCal(3, 4)
    fmt.Println("Multiplication:", mult)
    fmt.Println("Add:", add)

}
