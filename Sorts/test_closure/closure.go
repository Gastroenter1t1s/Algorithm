package closure
import(
	"fmt"
)



func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func GetLenForString(s1 StringPoint, s2 StringPoint) (size int) {
	size = len(*s1) + len(*s2)
	return
}

func SayHello(){
	fmt.Printf("Hello \n")
}
