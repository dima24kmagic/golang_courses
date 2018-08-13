package main

import (
	"fmt"
	"sort"
)

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

// --------- *: DATA STRUCTURES :* --------- //

// 1) SLICES (AKA Array, List) *value of uninitialized is nil					:*Reference Type
// 2) MAPS (*key/value* Store key and values) *value of uninitialized is nil	:*Reference Type
// 3) STRUCT (Structure *similar to object, can have props and methods inside*)

// Arrays
// func main() {
// 	// NOTE: If you got numbers in brackets - its ARRAY, if not - it's a slice
// 	// NOTE: Length is part of array type
// 	// NOTE: Not dynamic, *slices are
// 	var x [58]string
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(x[42])
// 	for i := 65; i <= 122; i++ {
// 		x[i - 65] = string(i)
// 	}
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(x[42])
// }

// Slices
// func main() {
// 	// NOTE: unlike array, slices got dynamic size
// 	mySlice := []int{1, 2, 3, 4, 1, 23}
// 	fmt.Printf("%T\n", mySlice)
// 	fmt.Println(mySlice)
// }

// Slice vs slicing cs index access
// func main()  {
// 	mySlice := []string{"a", "b", "c", "d", "e"}
// 	fmt.Println(mySlice) //[a b c d e]
// 	fmt.Println(mySlice[2:4]) //slicing a slice (basically saying - get all elements including from position to, to excluding position 4) *[c d]
// 	fmt.Println(mySlice[2]) // index access
// 	fmt.Println(string("Some String"[6])) //index access //116, and converst to string
// }

// NOTE: slice got an underlying arary behind, so size of slice is size of slice, :* Capacity is Size of array behind slice *:
// func main() {
// 	// NOTE: make a slice with length 10, and underly array of size 100 *so pc don't need to create new slice and underly array of more size*
// 	mySl := make([]int, 0, 5)
// 	// exact same thing using :* new keyword *:
// 	// mySl := new([100]int)[0:10]
// 	fmt.Println(mySl)
// 	for i := 0; i < 20; i++ {
// 		// :*NOTE:NOTE:*: That's the way to add elements to slice dynamicly
// 		mySl = append(mySl, i)
// 		fmt.Println("Len: ", len(mySl), "Capacity: ", cap(mySl), "Value:", mySl[i])
// 	}
// 	fmt.Println(mySl)
// }

// Appending slice to a slice
// func main()  {
// 	mySlice := []int{1,2,3,4,1,3}
// 	yourSlice := []int{1,3,4,15,2}
// 	mySlice = append(mySlice, yourSlice...) //NOTE: using variadic args ;)
// }

// Deleting from slice
// func main()  {
// 	mySlice := []int{1,2,3}
// 	yourSlice := []int{4,5,6}

// 	mySlice = append(mySlice, yourSlice...)
// 	fmt.Println(mySlice)

// 	// NOTE: like saying - yo, get all elements of a slice except :* second one (it's third one, 'cause 0, 1, 3) *:, then add elements from third position ( but like from fourth position, 'cause 0,1,2,3) to the end
// 	mySlice = append(mySlice[:2], mySlice[3:]...)
// 	fmt.Println(mySlice)
// }

// More slice examples
// Multi dimensional slice
// func main() {
// 	records := make([][]string, 0)
// 	student1 := make([]string, 4)
// 	// First Student
// 	student1[0] = "Dima"
// 	student1[1] = "Baranov"
// 	student1[2] = "100%"
// 	student1[3] = "4.5"
// 	// append student to array
// 	records = append(records, student1)
// 	// Next Student
// 	student2 := make([]string, 4)
// 	student2[0] = "DJ"
// 	student2[1] = "Khalid"
// 	student2[2] = "85%"
// 	student2[3] = "7.5"
// 	// store dj khaled
// 	records = append(records, student2)
// 	fmt.Println(records)
// 	// Slice that contain slice
// 	// NOTE: more convinient way to stire structs *As i think*
// 	type man struct {
// 		name string
// 	}
// 	mans := make([]man, 1)
// 	dima := man{"Dima"}
// 	mans = append(mans, dima)
// }

// TESTING
// WHY? I want to test, how fast it gonna be create a 100 million slice, and delete from it's something or move it higher
// NOTE: It's was super fast
// func main() {
// NOTE: That's an the most idiomatic way to create slices
// 	millionSlice := make([]string, 100000000)
// 	// add values to the slice
// 	chars := [10]string{"a", "b", "d", "c", "k", "s", "q", "e", "o", "m"}

// 	createWord := func(param [10]string) string {
// 		var word string
// 		for i := 0; i < 3; i++ {
// 			randomNum := rand.Intn(len(param))
// 			word += param[randomNum]
// 		}
// 		return word
// 	}
// 	for i := 0; i < 100000000; i++ {
// 		millionSlice[i] = createWord(chars)
// 	}
// 	fmt.Println("done")
// 	fmt.Println(millionSlice[99999997], millionSlice[99999998], millionSlice[99999999])
// 	swap := millionSlice[99999998]
// 	millionSlice[99999998] = millionSlice[99999999]
// 	millionSlice[99999999] = swap
// 	fmt.Println(millionSlice[99999997], millionSlice[99999998], millionSlice[99999999])
// 	fmt.Println("got it")
// 	millionSlice = append(millionSlice[:99999995], millionSlice[99999996:]...)
// 	fmt.Println("Deleted")
// }

// Example using slice
// func main() {
// 	mySlice := make([]int, 1)
// 	fmt.Println(mySlice[0])
// 	mySlice[0] = 5
// 	fmt.Println(mySlice[0])
// 	// :* Cool way to increase slice;) *:
// 	mySlice[0]++
// 	fmt.Println(mySlice[0])
// }

// NOTE: Expirience with slices of interface
// func main() {
// 	// Creating a slice of interfaces *like we can put whatever we want
// 	interSlice := make([]interface{}, 3)
// 	// Put string
// 	interSlice[0] = "A"
// 	// Put int
// 	interSlice[1] = 1
// 	// Put func, but
// 	// NOTE: it's putting not the func, but it's reference to memory adress
// 	interSlice[2] = func(){
// 		fmt.Println("Hi")
// 	}
// 	fmt.Println(interSlice) // [A 1 0x482270]
// 	// So, to use it, we doing assertion, and then self executing this function, btw we can do like this -->
// 	// --> myFunc := interSlice[2].(func()) -->
// 	// --> myFunc()
// 	interSlice[2].(func())()
// }

// --- :* MAPS *: --- //
// some lang called a dictionary

// NOTE: Create a map is pretty instinct make(map[string]int) - is like to saying - yo, make a map, that got key and it's string, and when i use this key - i'll get the int value
// func main() {
// 	myMap := make(map[string]int)
// 	myMap["DIMA"] = 5
// 	myMap["China"] = 175
// 	fmt.Println(myMap)
// }

// func main()  {
// 	myMap := make(map[string]int)
// 	// NOTE: Add to map
// 	myMap["DIMA"] = 5
// 	myMap["China"] = 175
// 	fmt.Println(myMap)
// 	// NOTE: length of a map
// 	fmt.Println(len(myMap))
// 	// NOTE: delete from map
// 	delete(myMap, "China")
// 	// NOTE: create a map with predefined values
// 	constValues := map[string]int{"Drake": 31, "Kendrick": 29}
// 	fmt.Println(constValues, constValues["Drake"], constValues["Kendrick"])
// 	// NOTE: check for a key in map
// 	value, isKeyHere := constValues["Drake"] // checking for key existense
// 	fmt.Println(value,",", isKeyHere) // 31, true
// }

// Using if and map
// func main() {
// 	myMap := map[int]string{1: "1", 2: "2", 3: "3", 4: "4"}
// 	delete(myMap, 1)
// 	if _, exist := myMap[1]; exist { // NOTE: using existence of value in map by it's key
// 		fmt.Println(myMap[1], "Is exist")
// 	} else {
// 		fmt.Println(myMap[1], "Doesn't exist")
// 	}
// }

// More Example
// :* Map in Map *:
// func main() {
// 	elements := map[string]map[string]string{
// 		"H": map[string]string{
// 			"name":  "Helium",
// 			"state": "Gas",
// 		},
// 	}
// 	fmt.Println(elements["H"]["name"])
// }

// :* Hash Tables *:

// :* Range loop *:
// func main() {
// 	// Range loop for map
// 	myMap := map[int]string{
// 		0: "Buy milk",
// 		2: "Meet my friends",
// 	}
// 	// NOTE: keyword range is like a function, that return a values and can be used in for loop
// 	for i, val := range myMap {
// 		fmt.Println(i, val)
// 	}
// 	// Range loop for slice
// 	mySlice := make([]int, 20, 100)
// 	fmt.Println(mySlice)
// 	for i := range mySlice {
// 		mySlice[i] = i
// 	}
// 	fmt.Println(mySlice)
// }

// func main() {
// 	res, _ := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
// 	// bs - bytes slice
// 	bs, _ := ioutil.ReadAll(res.Body)
// 	str := string(bs)
// 	fmt.Println(str)
// }

// Create a map from txt file with words
// func main() {
// 	res, _ := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
// 	words := make(map[string]string)
// 	// :* byfio - buffer input/output
// 	// Buffer - area to stick up some data
// 	// NOTE: we create a buffer and put the response body in it
// 	sc := bufio.NewScanner(res.Body)
// 	// Split data in buffer by words
// 	sc.Split(bufio.ScanWords)
// 	for sc.Scan() {
// 		words[sc.Text()] = ""
// 	}
// 	if err := sc.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "reading input", err)
// 	}
// 	i := 0
// 	for k := range words {
// 		fmt.Println(k)
// 		if i == 200 {
// 			break
// 		}
// 		i++
// 	}
// }

// func hashBucket(word string) int {
// 	return int(word[0])
// }
// func main() {
// 	// Get the book moby dick
// 	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// scan the page
// 	scanner := bufio.NewScanner(res.Body)
// 	defer res.Body.Close()
// 	// Set the split function for the scanning operation
// 	scanner.Split(bufio.ScanWords)
// 	// Create slice to hold counts
// 	buckets := make([]int, 200)
// 	// Loop over the words
// 	for scanner.Scan() {
// 		n := hashBucket(scanner.Text())
// 		buckets[n]++
// 	}
// 	fmt.Println(buckets[65:123])
// }

// :* Go Structs *:
// AGGREGATE TYPE EXAMPLE
// NOTE: Aggregate mean, that you aggregate together bunch of stuff
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }
// func main() {
// 	p1 := person{"Dima", "Baranov", 18}
// 	fmt.Println(p1.first)
// }

// Underlying Type
// NOTE: how my new type reffer too another type
// type myType int

// func main() {
// 	var aka myType
// 	aka = 5
// 	fmt.Println(aka)
// }

// NOTE: Embeded Types
// type person struct {
// 	First string
// 	Last  string
// }
// type human struct {
// 	person
// 	First  string
// 	Height int
// }

// func main() {
// NOTE: Here is overwriting fields
// 	h := human{
// 		person: person{
// 			"Dima",
// 			"Baranov",
// 		},
// 		First: "Cheddr", Height: 173}
// 	fmt.Println(h.person.First, h.First)
// }

// NOTE: struct Methods with receiver in function
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }
// //NOTE: Method for person type, that can access person vars, using receiver
// func (p person) fullName() string {
// 	return p.first + " " + p.last
// }

// func main() {
// 	p1 := person{"Dima", "Baranov", 18}
// 	fmt.Println(p1.fullName())
// }

// Struct Pointers
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }
// func (p person) getAge() {
// 	fmt.Println(p.age)
// }
// func main() {
// 	p1 := person{"Dima", "Baranov", 17}
// 	p1.getAge()
// }

// NOTE: Work With JSON
// :* JSON Marshal *:
// type person struct {
// 	First string `json:"firstName"`
// 	Last  string `json:"lastName"`
// 	Age   int    `json:"age"`
// }
// func main() {
// 	p1 := person{"Dima", "L", 18}
// 	bs, _ := json.Marshal(p1)
// 	fmt.Println(bs)
// 	// fmt.Printf("%T \n", bs)
// 	fmt.Println(string(bs))
// }

// :* JSON Unmarshal *:
// type person struct {
// 	First string
// 	Last  string
// 	Age   int
// }
// NOTE: I use TAGS to set json data from unmarshall here
// type dick struct {
// 	Ty   string `json:"First"`
// 	Pack string `json:"Last"`
// 	Sha  int `json:"Age"`
// 	Koor string
// }
// func main() {
// 	var p1 person
// 	fmt.Println(p1.First)
// 	fmt.Println(p1.Last)
// 	fmt.Println(p1.Age)
// 	bs := []byte(`{"First":"James", "Last":"Bond", "Age":32}`)
// 	//Unmarshal is put values to the interface
// 	var test dick
// 	json.Unmarshal(bs, &test)
// 	fmt.Println(test.Ty)
// 	fmt.Println(test.Pack)
// 	fmt.Println(test.Sha)
// }

// :* JSON ENCODING *:
// type person struct {
// 	First string
// 	Last  string
// 	Age   int
// }
// func main() {
// 	p1 := person{"Dima", "L", 18}
// 	json.NewEncoder(os.Stdout).Encode(p1)
// }

// :* JSON DECODING *:
// type person struct {
// 	First string
// 	Last  string
// 	Age   int
// }
// func main() {
// 	var p1 person
// 	rdr := strings.NewReader(`{"First": "James", "Last":"Bond", "Age": 17}`)
// 	json.NewDecoder(rdr).Decode(&p1)
// 	fmt.Println(p1.First)
// 	fmt.Println(p1.Last)
// 	fmt.Println(p1.Age)
// }

// NOTE::* Interfaces *:NOTE:
// Interfaces let us create substitutability ( взаимозаменяемость )
// This is type of Square but it's implement a method area() on line :* 1095 *:
// type Square struct {
// 	side float64
// }
// // Circle is type of Circle but it's implement a method area() on line :* 1102 *: so it can be treted as shape type
// type Circle struct {
// 	radius float64
// }

// // NOTE: Anything that got this method (area() method) signature - implement this interface
// type shape interface {
// 	area() float64
// }

// func (z Square) area() float64 {
// 	return z.side * z.side
// }

// func (c Circle) area() float64 {
// 	return 2 * math.Pi * c.radius
// }
// // Because Square type is implement Shape interface area method - it's become a type of shape
// func info(z shape) {
// 	fmt.Println(z)
// 	fmt.Println(z.area())
// }
// func main() {
// 	s := Square{10}
// 	info(s)
// 	c := Circle{3.15}
// 	info(c)
// }

//Excersises Sorting
// type people []string
// func main() {
// 	//Sorting
// 	studyGroup := people{"Dima", "Moohamed", "Gena", "Alie"}
// 	fmt.Println(studyGroup)
// 	sort.Strings(studyGroup)
// 	fmt.Println(studyGroup)
// 	s := []string{"Jenkins", "Docker", "AI", "Huncho"}
// 	sort.Strings(s)
// 	fmt.Println(s)
// 	n := []int{1, 3, 5, 2, 4, 5, 1, 2, 42, 512, 13}
// 	sort.Ints(n)
// 	sort.Sort(sort.Reverse(sort.IntSlice(n)))
// 	fmt.Println(n)
// } 
	
// Methods Sets