package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var actions = []string{
	"log in",
	"log out",
	"create",
	"read",
	"update",
	"delete",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getactivityInfo() string {
	result := fmt.Sprintf("ID: %d | Email: %s\nActivity Logs:\n", u.id, u.email)
	for _, log := range u.logs {
		result += log.action + " " + log.timestamp.String() + "\n"
	}
	return result
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-+=)(*&^%$#@!"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func generateUsers(amount int) []User {
	users := make([]User, amount)
	for i := 0; i < amount; i++ {
		tmpUser := User{
			id:    i,
			email: fmt.Sprintf("%s@%s", randomString(10), "gmail.com"),
			logs: []logItem{
				{action: actions[rand.Intn(len(actions))], timestamp: time.Now()},
				{action: actions[rand.Intn(len(actions))], timestamp: time.Now().Add(time.Hour)},
				{action: actions[rand.Intn(len(actions))], timestamp: time.Now().Add(time.Minute)},
			},
		}
		users[i] = tmpUser
	}
	return users
}

func safeUserInfo(user User, wg *sync.WaitGroup) error {
	printlnRed(fmt.Sprintf("Write to file %d", user.id))
	time.Sleep(time.Second)
	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getactivityInfo())
	if err != nil {
		return err
	}

	// done 1 task
	wg.Done()
	return nil
}

func main() {
	println("Wubba lubba dub dub!")

	// package that is used to wait for a collection of goroutines (concurrently executing functions)
	// to finish their work before proceeding further in the program
	wg := &sync.WaitGroup{}

	users := generateUsers(60)

	t := time.Now()

	for _, user := range users {
		// add 1 task
		wg.Add(1)
		go safeUserInfo(user, wg)

	}

	wg.Wait()

	printlnYellow(fmt.Sprintf("Time Elapsed:", time.Since(t)))

	// ///////////////////////////
	/* GO ROUTINES
	lightweight threads
	*/

	go printlnBlue("Concurrent Blue") // add to queue (and wait for blockers)
	go printlnMagenta("Concurrent Magenta")
	go printlnCyan("Concurrent Cyan")

	// block execution
	time.Sleep(10000)

	printlnGreen("Non concurrent Green")
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
