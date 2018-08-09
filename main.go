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
//NOTE: _ is used when you don't wan't to, like, using return variable or smthing like that
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
// NOTE: << is a binary shift *:BITWISE OPERATION:*, *basicly when you adding 0 to variable binary representation
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
// 	//NOTE: *int is a type!
// 	fmt.Println("b adress - ",b)
// 	fmt.Println("b value - ", *b)
// }

// Working with referencing/Dereferencing
// func main()  {
// 	a:= 43
// 	fmt.Println(a)
// 	fmt.Println(&a)

// 	//NOTE: we can declare be as b := &a or var b = &a
// 	var b *int = &a

// 	fmt.Println(b)
// 	//Pointer is going to adress and bring a value from it
// 	//NOTE: * operator *:STAR OPERATOR:* this is also called *:DEREFERENCER:*
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
// NOTE: When we passing arguments to the function, it's receive just a value and then assign it to a variable,
// So it got a different memory addresses
// func zero(x int) {
// 	x = 0
// }
// func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x)
// }
// NOTE: This example gonna work, because we passing a memory adress value, and then assign it to a reference
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
// 	//NOTE: Rune is a any character from any language all over the world
// 	//@ASK: What does Rune represent in Golang?
// 	//NOTE: It is a decimal int value, that stands for a character decimal value in UTF-8 Table
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
// NOTE: No fall through in default switch
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
// NOTE: you can specify fallthrough by keyword fallthrough
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
// 	case "Jordan", "Dima":         // NOTE: Multiple cases
// 		fmt.Println("Dima")
// 	case "Rebook":
// 	default:
// 		fmt.Println("Hi")
// 	}
// }

// Switch with no expression
// func main() {
// 	switch { //NOTE: no expresion defined, so it's going through cases finding that evaluate to true
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
// type contact struct {
// 	greeting   string
// 	contactNum string
// }
// func switchontype(x interface{}) { //NOTE: Passing an interface{} is mean, that we don't know, what argument type gonna be, it can be any type
// 	switch x.(type) {
// 	case string:
// 	case int:
// 		fmt.Println("Integer Type!")
// 	case bool:
// 		fmt.Println("It's a boolean!")
// 	case float32:
// 	case contact:
// 		fmt.Println("Thats's Contact type!")
// 	default:
// 		fmt.Println("Hi")
// 	}
// }
// func main() {
// 	x := 5
// 	switchontype(x) // int type
// 	newContact := contact{"Hello World!", "+375447278635"}
// 	switchontype(newContact) //contact type
// 	switchontype(false)      // bool type
// }

// If statements
// func main() {
// 	if true {
// 		fmt.Println("Hi!")
// 	}
// 	if false {
// 		fmt.Println("Bye!")
// 	}

// 	if !true { // *: (! is not)
// 		fmt.Println("FALSE NOT GONNA WORK")
// 	}
// }

//NOTE: If statement with initialization
// func returnString(x *interface{}) string {
// 	// Passing x var of any type and getting reference to it's memory adress
// 	// With pointer change value in memory adress
// 	fmt.Printf("%T \n", *x)
// 	if reflect.TypeOf(*x).String() == "int" {
// 		return "Hola"
// 	}
// 	*x = 7.5
// 	myReturn := "I'm returned"
// 	return myReturn
// }
// func main() {
// b := true
// Initialize something, to keep my scope tight
// if food := "Spaghetti"; b {
// 	fmt.Println(food)
// }
// NOTE:NOTE:
// Effective using:
// if err := file.chmod(7777); err != nil { // we execute function, that return an error, if error is not nil, error gonna be handled, and error variable gonna be in scope of if statement
// 	fmt.Println("Err")
// }
// NOTE: another case
// var x interface{} = 5.5
// 	fmt.Println(x)
// 	if strReturn := returnString(&x); strReturn != "" {
// 		fmt.Println("String returned!", strReturn)
// 	}
// 	fmt.Println(x)
// 	fmt.Printf("%T", x)
// 	// fmt.Println(strReturn) // Can't access strReturn from uorside of if block
// }

// -------------*: FUNCTIONS :* ------------- //

// NOTE: Difference between parameter and argument
// func greet(msg string) { // Func Receive a PARAMETER
// 	fmt.Println(msg)
// }
// func main() {
// 	greetings := "Hola @MoonStorm"
// 	greet(greetings) //We pass a ARGUMENT
// }

// Functions that return value
// func greeting(a ...string) string { // NOTE: ...string is mean UNLIMITED NUMBER OF ARGUMENTS
// 	return fmt.Sprint("Hello ", a, " ")
// }
// func main() {
// 	greeting := greeting("Dima", "Baranov")
// 	fmt.Println(greeting)
// }

//Named return (Don't use it)
// func namedReturn() (s string) {
// 	a := 5
// 	fmt.Println(a)
// 	s = "Wasup my boooi" //NOTE: that i'm not declaring new variable, but assign a value to it
// 	return
// }
// func main() {
// 	fmt.Println(namedReturn())
// }

// NOTE: Multiple value return
// func mulReturn() (string, int) {
// 	a := 42
// 	greet := "My booi is"
// 	return greet, a
// }
// func main() {
// 	greet, age := mulReturn()
// 	fmt.Println(greet, age)
// }

// Variadic functions (My Russian Notice: variadic от слова вариантивный)
// variadic from word (various)
// NOTE: Variadic function is a function, that can receive from 0 to any number of parameters
// func average(allParams ...float64) float64 { //NOTE: Can receive multiple params, it's a *: variadic parameter :*
// 	fmt.Println(allParams)
// 	fmt.Printf("%T \n", allParams)
// 	var total float64
// 	for _, param := range allParams {
// 		total += param
// 	}
// 	fmt.Println(total)
// 	return total / float64(len(allParams))
// }
// func main() {
// 	data := []float64{43,56,48,468,78,45,35}
// 	// average := average(45, 25, 13, 1, 78, 15, 4684) // passing literals
// 	avg := average(data...) //NOTE: It's a variadic argument
// passing slice, data... is telling, that it's should be passed by one
// 	fmt.Println("Average is - ", avg)
// }

// Function Expression
// func makeFunc() func() string {
// 	return func() string{
// 		return "Greetings"
// 	}
// }
// func main()  {
// 	// NOTE: Assigning function to variable is called func expression
// 	funcExp := func(){
// 		fmt.Println("Wasup!")
// 	}
// 	funcExp()

// 	greeter := makeFunc()
// 	fmt.Println(greeter())
// }

// Callbacks
// func visit(numbers []int, cb func(int)){
// 	for _, n := range numbers {
// 		cb(n)
// 	}
// }
// func main()  {
// 	visit([]int{1,2,3,4,5}, func(n int){
// 		fmt.Println(n)
// 	})
// }

// CB example (CB is stands for CallBack)
// func filter(numbers []int, cb func(int) bool) []int {
// 	xs := []int{}
// 	for _, n := range numbers {
// 		if cb(n) {
// 			xs = append(xs, n)
// 		}
// 	}
// 	return xs
// }
// func main() {
// 	data := []int{1, 2, 3, 4, 5, 6}
// 	xs := filter(data, func(n int) bool {
// 		if n%2 == 0 {
// 			return true
// 		}
// 		return false
// 	})
// 	fmt.Println(xs)
// }

// Recursion
// NOTE: Almost every problem we can solve without recursion, so don't use it, 'cause it got big cost
// func factorial(x int) int {
// 	if x == 0 {
// 		return 1
// 	}
// 	return x * factorial(x -1)
// }
// func main()  {
// 	fmt.Println(factorial(4))
// }

// NOTE: Defer
// *: WITHOUT DEFER :*
// func hello()  {
// 	fmt.Print("Hello")
// }
// func world()  {
// 	fmt.Println("World")
// }
// func main()  {
// 	world()
// 	hello()
// 	// out. --> World Hello
// }

// *: WITH DEFER :*
// func hello() {
// 	defer fmt.Print("awesome ")
// 	fmt.Print("Hello ")
// 	// NOTE: Defer execute here
// }
// func world() {
// 	fmt.Println("World")
// }
// func main() {
// 	// NOTE:NOTE:*: Defer is execute line of code right before function, that it's declared exit:*NOTE:NOTE:
// 	defer world()
// 	hello()
// 	// NOTE: You can think, that defer is execute here
// }

// *: Defer Usage :*
// func main() {
// 	file, err := os.Open("./api.hadler.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close() //NOTE: We closing the file, so it's not use the os resources. And not for type this at the end of a func, we just use defer
// }

// Slices is reference type example 
// func changeSlice(x []int) {
// 	// x =[]int{3, 2, 1}
// 	x[0] = 9
// }
// func main() {
// 	// NOTE: Slices maps and channels are data structures, that (reference type) already referencing to memory adress, so, when we change it from another function - it's changed at this scope
// 	mySlice := []int{}
// 	mySlice = []int{1, 2, 3}
// 	fmt.Println(mySlice) // [1 2 3]
// 	changeSlice(mySlice)
// 	fmt.Println(mySlice) // [9 2 3]
// }

// Map is Ref type
// func changeMe(z map[string]int)  {
// 	// NOTE: Maps is smthing like js object
// 	z["Todd"] = 44
// }
// func main()  {
// 	m:=make(map[string]int)
// 	changeMe(m)
// 	fmt.Println(m["Todd"]) // 44
// }

//Anonymus self executing function (Work like in js)
// func main()  {
// 	func(n int){
// 		fmt.Println(n)
// 	}(5)	
// }

// --- *: Bool Expressions :* --- //
// || (or); && (and); ! (not) (work like in js)
// func main()  {
// 	if true || false {
// 		fmt.Println("When use || (or) - true is stronger, so any true can go through any false")
// 	}
// 	if true && false {
// 		fmt.Println("When use && (and) - false is stronger than any true, so if any false is happened - all is evaluate to false")
// 	}
// 	if !true {
// 		fmt.Println("! not is switch true to false / false to true")
// 	}
// 	// Combining
// 	//   |----false-----|  true, false or true? --> true
// 	if ((true && false) || true) {
// 		fmt.Println("Yeaaaah")
// 	}
// }

// NOTE: Switch by type excercise, using function expression
// func main()  {
// 	typeSwitcher := func(a interface{}) {
// 		switch a.(type) {
// 		case int:
			
// 		}
// 	}
// 	typeSwitcher(5)
// 	typeSwitcher("HIII")
// 	typeSwitcher(true)
// }