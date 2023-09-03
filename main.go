// ///////////////////////////
// current package - MANDATORY
package main

// ///////////////////////////
// go standard libraries
// https://pkg.go.dev/std
import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var globalVar string = "GLOBAL"

// ///////////////////////////
// invokes in each packege in initialization stage (before main)
func init() {
	println("Executes before main", globalVar)
	globalVar = "FROM INIT()"
}

// ///////////////////////////
// entry point of execution - should be in main package
func main() {
	println("Wubba lubba dub dub!", globalVar)

	// ///////////////////////////
	/* VARIABLES
	:= write value in var and initialize
	= write value in var
	static types of vars
	*/
	msg := "Villanelle"
	msg = "Eve"
	var variableStr string // "strings"
	var variableInt int
	var variableUInt uint
	var variableBool bool
	var variableFloat float32
	var variablebyte byte // int from 0-255
	var variableRune rune // int32 (4byte) 'symbol'

	fmt.Println(reflect.TypeOf(msg))
	fmt.Println("Killing", msg)
	fmt.Println("variableStr", variableStr)     // empty str
	fmt.Println("variableInt", variableInt)     // 0
	fmt.Println("variableUInt", variableUInt)   // 0
	fmt.Println("variableBool", variableBool)   // false
	fmt.Println("variableFloat", variableFloat) // 0
	fmt.Println("variablebyte", variablebyte)   // 0
	fmt.Println("variableRune", variableRune)   // 0

	// multiassign
	a, b, c := 1, 3, false
	// exchange values
	a, b = b, a
	// omited var
	_ = 89

	fmt.Println("a,b,c", a, b, c)

	// ///////////////////////////
	/* FUNCTIONS
	multireturn is allowed
	error should be returned last
	*/
	switchCaseExample(255)
	var sum, sub, err = addSubInt8(127, 1)

	if err == nil {
		fmt.Println("err", err) // err test error
	}

	fmt.Println("sum", sum) // -128
	fmt.Println("sub", sub) // 126

	min, _ := findMin(8, 78, 16, 255, 7)
	fmt.Println("min", min) // 7

	// anonymous self invoking function
	func() {
		println("anon func")
	}()

	// clousure (remembering executing environmants inside functions)
	counter := roundInc()

	for i := 0; i < 4; i++ {
		println((counter()))
	}

	// ///////////////////////////
	/* POINTERS
	pointer = refernce to memory cell
	param *type -> in func params
	&param -> to pass in func when calling = referesing
	*param = dereferesing
	create pointer variable =>  var p *int
	nil => zero/empty pointer that not referenced yet
	*/
	number := 78
	var p *int
	fmt.Println(p) // <nil>

	p = &number
	fmt.Println(p)       // 0x1400000e088
	fmt.Println(&number) // 0x1400000e088

	inputString := "pointers"
	// pass "copy" of var
	toUpperCase(inputString)
	// pass reference to ram
	toUpperCasePointer(&inputString)
	printlnCyan(inputString)

	// ///////////////////////////
	/* ARRAYS & SLICES
	arr => static size arr
	slice => dynamic size arr, arr wrapper what contains reference to origin arr (in func pass pointer)
	slice reallocation creates (recreates) arr, like ArrayList in Java
	*/
	// arr
	msgs := [3]string{}                 // arr with empty values
	fmt.Println(msgs)                   // [  ]
	filled := [3]string{"f", "9", "ck"} // arr with empty values
	filled[1] = "_"                     //
	fmt.Println(filled)                 //[k 9 kk]

	// sclice
	slice := []string{} // empty slice
	fmt.Println(slice)  // []
	uninitslice := make([]string, 6, 17)

	uninitslice[5] = "Patrick"
	uninitslice = append(uninitslice, "Jane")
	fmt.Println(uninitslice)      // []
	fmt.Println(len(uninitslice)) // 6
	fmt.Println(cap(uninitslice)) // 17

	matrix := make([][]int, 10)
	for i := 0; i < 10; i++ {
		for y := 0; y < 10; y++ {
			matrix[y] = make([]int, 10)
			matrix[y][i] = i
			matrix[i][y] = y

		}
		fmt.Println(matrix[i])
	}

	// keyword defer - postpone execution until later (before panic)
	printlnMagenta("Pure Magenta 1")

	defer handlePanic()
	defer printlnMagenta("defer Magenta 1") // second
	defer printlnMagenta("defer Magenta 2") // first

	printlnMagenta("Pure Magenta 2")

	// func panic() - apps stop with exit status 2 (in cases of dynamic errors)
	// panic("help")

	// ///////////////////////////
	/* MAP (like map or obj in js)
	Maps are implemented as hash tables - This means that the average time complexity
	for insertion, deletion, and retrieval (searching for a key) is O(1)
	Iterating Over a Map: Iterating over a map using a for range loop has a time complexity of O(n),
	as you need to visit each key-value pair once.
	*/
	// inilialize with elements
	users := map[string]int{
		"Jane":   78_000,
		"Claire": 89_000,
		"Patric": 2_500,
	}

	fmt.Println(users["Claire"]) // 89000
	fmt.Println(users["No"])     // 0

	// find by key
	_, isExist := users["Claire"]

	fmt.Println(isExist)

	// add to map
	users["Eve"] = 700_000

	// delete from map
	delete(users, "Patric")

	for key, value := range users {
		printlnGreen(key + " " + strconv.FormatInt(int64(value), 10))
	}

	// uninitialized map
	var newMapUsers map[string]int
	// initialize
	newMapUsers = make(map[string]int, 100)

	newMapUsers["Emily"] = 45_000

}

func handlePanic() {
	// recover() built-in Go function used to handle panics
	// r != nil => it means a panic occurred and was successfully recovered
	r := recover()
	if r != nil {
		printlnCyan("Ignore panic")
	}
}

func toUpperCase(inputString string) {
	inputString = strings.ToUpper(inputString)
	printlnCyan(inputString)
}

func toUpperCasePointer(inputString *string) {
	*inputString = strings.ToUpper(*inputString)
	printlnRed(*inputString)
}

func printlnCyan(inputString string) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 36, inputString))
}

func printlnMagenta(inputString string) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 35, inputString))
}

func printlnBlue(inputString string) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, inputString))
}

func printlnYellow(inputString string) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, inputString))
}

func printlnGreen(inputString string) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, inputString))
}

func printlnRed(inputString string) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, inputString))
}

func printlnBlack(inputString string) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 30, inputString))
}

func roundInc() func() uint8 {
	var counter uint8 = 0
	return func() uint8 {
		counter++
		return counter
	}

}

func addSubInt8(a int8, b int8) (int8, int8, error) {
	return a + b, a - b, errors.New("test error")
}

func switchCaseExample(a uint8) {
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("non 1, 2, 3")
	}
}

func findMin(numbers ...uint) (uint, error) {

	if len(numbers) == 0 {
		return 0, errors.New("no numbers provided")
	}
	min := numbers[0]

	for _, i := range numbers {
		if min > i {
			min = i
		}
	}
	return min, nil

}
