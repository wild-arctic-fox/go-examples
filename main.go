package main

import (
	"fmt"
	"go-booking-app/user"
	"strings"
	"time"
)

func Leave(student user.Student) {
	student.LeaveUni()
}

func main() {
	println("Wubba lubba dub dub!")

	// ///////////////////////////
	/* STRUCTS
	Custome complex types (class alternatives)
	could have fields, methods
	*/

	// user1 := user.User{"Rick", 67, 'A', true, 66_000.90}
	user2 := user.UserConstructor("Morty", 14, 'B', true, 800.00)

	// fmt.Println(user1)
	user2.UpdateUser(255)
	user2.PrintUser()
	fmt.Println(user2.Aage.IsAdult())

	Leave(user2)

	fmt.Println(getTypeName(42))      // Output: This is an integer
	fmt.Println(getTypeName(3.14))    // Output: This is a float
	fmt.Println(getTypeName("Hello")) // Output: This is a string
	fmt.Println(getTypeName(true))    // Output: Unknown type

	// Time

	fmt.Println(time.Now())

	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	printlnCyan(formattedTime)

	arrInt := []int{78, 90, 67}
	fmt.Println(sum(arrInt))

	arrFloat32 := []float32{78.00, 90.00, 67.0109}
	printlnYellow((sum(arrFloat32)))

	fmt.Println(find(arrInt, 90))
	fmt.Println(find(arrFloat32, 67.0109))
	printlnGreen(find(arrFloat32, 67.0101))

}

// comparable - every type exept of map|slice or stuct that contains map|slice
func find[C comparable](arr []C, el C) bool {
	res := false
	for _, v := range arr {
		if v == el {
			res = true
		}
	}
	return res
}

func sum[number int | float32](num []number) number {
	var sum number
	for _, v := range num {
		sum += v
	}
	return sum
}

// type switch

func getTypeName(x interface{}) string {
	switch x.(type) {
	case int:
		return "This is an integer"
	case float64:
		return "This is a float"
	case string:
		return "This is a string"
	default:
		return "Unknown type"
	}
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

func printlnCyan[A any](inputString A) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 36, fmt.Sprintf("%v", inputString)))
}

func printlnMagenta[A any](inputString A) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 35, fmt.Sprintf("%v", inputString)))
}

func printlnBlue[A any](inputString A) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, fmt.Sprintf("%v", inputString)))
}

func printlnYellow[A any](inputString A) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, fmt.Sprintf("%v", inputString)))
}

func printlnGreen[A any](inputString A) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, fmt.Sprintf("%v", inputString)))
}

func printlnRed[A any](inputString A) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, fmt.Sprintf("%v", inputString)))
}

func printlnBlack[A any](inputString A) {
	println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 30, fmt.Sprintf("%v", inputString)))
}

// func roundInc() func() uint8 {
// 	var counter uint8 = 0
// 	return func() uint8 {
// 		counter++
// 		return counter
// 	}

// }

// func addSubInt8(a int8, b int8) (int8, int8, error) {
// 	return a + b, a - b, errors.New("test error")
// }

// func switchCaseExample(a uint8) {
// 	switch a {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	default:
// 		fmt.Println("non 1, 2, 3")
// 	}
// }

// func findMin(numbers ...uint) (uint, error) {

// 	if len(numbers) == 0 {
// 		return 0, errors.New("no numbers provided")
// 	}
// 	min := numbers[0]

// 	for _, i := range numbers {
// 		if min > i {
// 			min = i
// 		}
// 	}
// 	return min, nil

// }
