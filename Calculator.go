package main
import "fmt"

func add (x int, y int) int {
	return x + y
}

func subtract (x int, y int) int {
    return x - y
}

func multiply (x int, y int) int {
    return x * y
}

func divide (x int, y int) int {
    return x / y
}

func findsqr (x int) int{
	return x * x
}

func findcube (x int) int{
	return x * x * x
}

func main(){
	fmt.Println (divide (33,3))
}
