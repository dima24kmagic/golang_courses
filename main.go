package main

import "fmt"

// Passign the callback and returning function
// func wrapper(cb func()) func(someVal int) {
// 	cb()
// 	return func(someVal int) {
// 		fmt.Println(someVal)
// 	}
// }
// passing anonymus func as a callback and receiving anonymus func and assign it's to the variable
// func main() {
// 	var x int
// 	exec := wrapper(func(){
// 		x = 5
// 	})
// 	exec(x)
// }

//Need to think about, neet to reflect on this
// func wrapper() func() int {
// 	var x = 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func main() {
// 	// we creating x variable and ->
// 	// -> we creating var *increment* and putting in it function, that's increment it's x
// 	var increment = wrapper()
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

// Showing, that every var should be used
// func main() {
// 	a := "a"
// 	b := "b"
// 	fmt.Println(a)
// 	// fmt.Println(b)
// }

//Example of using _*underscore* variable
//@NOTE: _ is used when you don't wan't to, like, using return variable or smthing like that
// func main() {
// 	res, _ := http.Get("https://www.google.com")
// 	page, _ := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	fmt.Printf("%s", page)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// fmt.Printf("%s", res)
// }

//Example of creating functions with two return
// func main() {
// 	booName, age := returnTwoShits()
// 	fmt.Println(booName, age)
// }
// func returnTwoShits() (string, int) {
// 	booAge := 18
// 	myBooName := "Leela"
// 	return myBooName, booAge
// }

//Example with iota
// @NOTE: << is a binary shift *:BITWISE OPERATION:*, *basicly when you adding 0 to variable binary representation
// const (
// 	_ = iota
// 	a = 1 << iota * 10
// 	b = 1 << iota * 10
// )
// func main() {
// 	const myAge = 18
// 	const MeAGEE int = 15
// 	fmt.Printf("%d \t %b \t %d \t %b \t \n", a, a, b, b)

// }

//Example with UNTYPED and TYPED variables
// func main() {
// 	// Creating UNTYPED const value and sum up them, give us an TYPED *float64* variable
// 	const untyped = 5
// 	const anotherUntuped = 5.4
// 	var test = untyped + anotherUntuped
// 	fmt.Println(untyped, anotherUntuped, test)
// 	fmt.Printf("%T \n", test)

// 	//Creating two untyped vars, but go compiler make them int and float64 types, so we can't add them together
// 	var unt1 = 5
// 	var unt2 = 2.7
// 	fmt.Println(unt1, unt2)
// 	//summ them up into new variables
// 	typed1 := float64(unt1) + unt2
// 	fmt.Printf("%T %f\n", typed1, typed1)
// }

//Memory adress & Referencing & umpersand symbol is reference to memory adress
// func main() {
// 	a := 43
// 	fmt.Println("a - ", a)
// 	//Memory adress, where 0x stands for hexidecimal number
// 	fmt.Println("a's memory adress - ", &a)
// 	//Transofrm memory adress hex to decimal value
// 	fmt.Printf("%d \n", &a)
// }

// Scanf method is put a value inside memory adress using &
// const metersToYards float64 = 1.09361
// func main() {
// 	var meters float64
// 	fmt.Println("Enter meters swam: ")
// 	fmt.Scan(&meters)
// 	yards := meters * metersToYards
// 	fmt.Println(meters, " meters is ", yards, " yards")
// }

//Working with Pointers
// func main() {
// 	//var a is type int
// 	a := 43
// 	fmt.Println(a)
// 	//reference to memory adress of a
// 	fmt.Println(&a)

// 	//b is type of pointer to int, that's point to memory adress of a
// 	//b is pointer to reference of a memory adress
// 	var b *int = &a
// 	//@NOTE: *int is a type!
// 	fmt.Println("b adress - ",b)
// 	fmt.Println("b value - ", *b)
// }

// Working with referencing/Dereferencing
// func main()  {
// 	a:= 43
// 	fmt.Println(a)
// 	fmt.Println(&a)

// 	//@NOTE: we can declare be as b := &a or var b = &a
// 	var b *int = &a

// 	fmt.Println(b)
// 	//Pointer is going to adress and bring a value from it
// 	//@NOTE: * operator *:STAR OPERATOR:* this is also called *:DEREFERENCER:*
// 	fmt.Println(*b)
// }

//Workink with Pointers

// Func that took a reference to mem adress and change a value
// func referencer(ref *int, newValue int) {
// 	// Change the value of reference value
// 	*ref = newValue
// }
// func main() {
// 	a := 15
// 	fmt.Println("a - ", a)
// 	fmt.Println("a adress - ", &a)
// 	// Using referencefunction
// 	referencer(&a, 25)
// 	fmt.Println("a adress - ", &a)
// 	fmt.Println("a - ", a)
// }

// *: EXERCISES ON POINTERS :*
// @NOTE: When we passing arguments to the function, it's receive just a value and then assign it to a variable,
// So it got a different memory addresses
// func zero(x int) {
// 	x = 0
// }
// func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x)
// }
// @NOTE: This example gonna work, because we passing a memory adress value, and then assign it to a reference
// func zero(z *int) {
// 	*z = 0
// }
// func main() {
// 	x := 106
// 	fmt.Println("x - ", x)
// 	fmt.Println("x adress - ", &x)
// 	zero(&x)
// 	fmt.Println("x - ", x)
// 	fmt.Println("x adress - ", &x)
// }

// Remainders
// func main() {
// 	x := 15 % 4 // Return 2 WHY?
// 	fmt.Println(x)
// 	// @BECAUSE: 10 % 4 is like ---> 10 / 4 = 2, then 2 * 4 = 8, then 10 - 8 = 2 >> Bigger than
// 	//Examples
// 	// 17 % 3
// 	z:= 17%3
// 	fmt.Println("z is by remainder operator - ",z)
// 	divided := 17
// 	divider := 3
// 	a := divided / divider
// 	fmt.Println("By hands - ", divided - a * divider)
// }

// --------*: Exercises :*-------- //
// func main() {
// 	fmt.Println("Dima")
// }

// func main() {
// 	const name string = "Dima"
// 	fmt.Println(name)
// }

// func main()  {
// 	var name string
// 	fmt.Println("Enter Your name")
// 	fmt.Scan(&name)
// 	fmt.Println("Hello ", name)
// }

// func main() {
// 	fmt.Println("Enter small number")
// 	var smallNum int
// 	fmt.Scan(&smallNum)
// 	fmt.Println("Enter Large Number")
// 	var largeNum int
// 	fmt.Scan(&largeNum)
// 	fmt.Println("Remainder is - ", largeNum%smallNum)
// }

// --------*: CONTROL FLOW :*-------- //

//For loop
// func main() {
// 	for i := 1; i <= 100; i++ {
// 		fmt.Println(i)
// 	}
// }

// For loop in for loop
// func main() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			fmt.Println(i, " - ", j)
// 		}
// 	}
// }

// For loop with condition (like C While Loop)
// func main()  {
// 	i := 0
// 	for i<10 {
// 		fmt.Println("Sugar")
// 		i++
// 	}
// }

// Using *: continue :* keyword
// func main() {
// 	i := 1
// 	for {
// 		i++
// 		if i%2 != 0 {
// 			continue // if continue call - stop current loop iteration and start it from beginnig
// 		}
// 		fmt.Println(i)
// 		if i > 50 {
// 			break
// 		}
// 	}
// }

//*: Runes :*
// func main() {
// 	//@NOTE: Rune is a any character from any language all over the world
// 	//@ASK: What does Rune represent in Golang?
// 	//@NOTE: It is a decimal int value, that stands for a character decimal value in UTF-8 Table
// 	a := 'a' //97 in UTF-8 table
// 	fmt.Println(a) //97
// 	fmt.Printf("%b \n", a) // binary representation of a rune (from ascii table)
// }

// func main()  {
// 	for i := 33; i < 256; i++ {
// 		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
// 	}
// }

// Switch
// @NOTE: No fall through in default switch
// func main() {
// 	switch "Dima" {
// 	case "Kate":
// 	case "Nikol":
// 		fmt.Println("Nikol")
// 	case "Dima":
// 		fmt.Println("Dima")
// 	case "Rebook":
// 	default:
// 		fmt.Println("Hi")
// 	}
// }

// Switch with *: fallthrough :*
// @NOTE: you can specify fallthrough by keyword fallthrough
// func main() {
// 	switch "Dima" {
// 	case "Kate":
// 	case "Nikol":
// 		fmt.Println("Nikol")
// 	case "Dima":
// 		fmt.Println("Dima")
// 		fallthrough
// 	case "Rebook":
// 		fmt.Println("I'm fallthrough consequence!")
// 	default:
// 		fmt.Println("Hi")
// 	}
// }

// Switch with cases variation
// func main() {
// 	switch "Dima" {
// 	case "Kate":
// 	case "Nikol":
// 		fmt.Println("Nikol")
// 	case "Jordan", "Dima":         // @NOTE: Multiple cases
// 		fmt.Println("Dima")
// 	case "Rebook":
// 	default:
// 		fmt.Println("Hi")
// 	}
// }

// Switch with no expression
// func main() {
// 	switch { //@NOTE: no expresion defined, so it's going through cases finding that evaluate to true
// 	case 5 == 2:
// 	case 1 > 3:
// 		fmt.Println("Nikol")
// 	case 2 < 7:
// 		fmt.Println("Dima")
// 	case 7 > 10:
// 	default:
// 		fmt.Println("Hi")
// 	}
// }

// Switch by type
type Contact struct {
	greeting stringname string
}
func switchontype(x interface{}){
	switch x, (type) {
	case string:
	case int:
		fmt.Println("Nikol")
	case bool:
		fmt.Println("Dima")
	case float32:
	default:
		fmt.Println("Hi")
	}
})
func main() {
	x := 5
	
}
