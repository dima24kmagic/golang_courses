package main

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
//_ is used when you don't wan't to, like, using return variable or smthing like that
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

func main() {

}
