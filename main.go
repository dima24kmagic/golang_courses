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

