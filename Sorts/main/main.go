package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	closure "../test_closure"
)

type Options struct {
	par1 int64
	par2 int32
	par3 string
}

func main() {
	var v string = "jaiqjad"
	var customReader *strings.Reader = strings.NewReader(v)
	fmt.Printf("length of the string: %d\n", customReader.Len())
	fmt.Printf("we are going to iterate the whole string's each character\n")
	for s, c := range v {
		fmt.Printf("character on position %d is %c\n", s, c)
	}
	customRune, size, errorA := customReader.ReadRune()
	fmt.Printf("%d -- %d -- The customReader type is %T\n", customRune, size, customRune)
	fmt.Println(errorA)

	var template_byte byte
	for count, err := 0, error(nil); count <= customReader.Len(); {
		template_byte, err = customReader.ReadByte()
		fmt.Printf(" charcter %c on position %d--%c\n", template_byte, count, err)
		if err != nil {
			break
		} else {
			count++
		}
	}
	fmt.Println("...................")

	anotherNewReader := strings.NewReader("we are going to say that hello, mother fucker!")
	ioSectionReader := io.NewSectionReader(anotherNewReader, 0, anotherNewReader.Size())
	bufferForSectionReader := make([]byte, anotherNewReader.Size())

	if _, err := ioSectionReader.Read(bufferForSectionReader); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bufferForSectionReader)

	someExamplesOfLastNames := []string{"Paul", "Json", "Maria", "Luxy"}
	lengtheningArrayTypes(someExamplesOfLastNames...)
	testInitializationTypeStruct(1, "t", Options{par1: 123, par2: 234, par3: "what the fuck..."})
	///////////////////////////////////////////////////
	func_ptr := closure.Add()
	rs := func_ptr(888)
	var ptr_int closure.IntPoint = &rs
	fmt.Printf("the result of the clsoure function Add() is %d\n", *ptr_int)
	func_ptr2 := closure.Adder(20)
	fmt.Printf("the second result of closure functon Adder() is %d\n", func_ptr2(90))
	closure.SayHello()

	var qda *closure.Qda = new(closure.Qda)
	closure.Initialization(33, "Yes", qda)
	closure.Display(qda)
	err := closure.Reset(qda)
	fmt.Printf("%s\n", err)
	closure.Reset(qda)
	err = closure.Display(qda)
	fmt.Printf("%s\n", err)

}

func lengtheningArrayTypes(personLastNames ...string) (size int, err error) {
	if len(personLastNames) == 0 {
		size = 0
		return
	}

	for _, v := range personLastNames {
		fmt.Printf("%s - ", v)
	}
	fmt.Printf("\n")
	size = len(personLastNames)
	err = nil
	return
}

func testInitializationTypeStruct(a int, b string, o Options) {

}
