package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Name string
	ID   int
	err  error
}

func (user User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", user.ID, user.Name)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	f, err := os.Open("./users.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for userChan := range decodeUsers(ctx, f) {
		if userChan.err != nil {
			panic(err)
		}
		fmt.Println(userChan)
	}
}

func decodeUsers(ctx context.Context, r io.Reader) chan User {
	ch := make(chan User, 1)

	// Start goroutine to feed the channel we will return
	go func() {
		// Close the channel on exit, signaling we're done
		defer close(ch)

		// Wrap the Scanner around our reader so we can read line of content
		scanner := bufio.NewScanner(r)
		for scanner.Scan() { // Scan until nothing to scan
			if ctx.Err() != nil { // Context was cancelled, return error
				ch <- User{err: ctx.Err()}
				return
			}

			// Turn the line of text into a User object
			user, err := getUsers(scanner.Text())
			if err != nil { // line was in incorrect format, return an error
				user.err = err
				ch <- user
				return
			}
			// Everything was fine, return a user record
			ch <- user
		}
	}()

	// Returns the channel we will read off of
	return ch
}

func getUsers(s string) (User, error) {
	sArray := strings.Split(s, ":")
	if len(sArray) != 2 {
		return User{}, fmt.Errorf("record(%s) was not in the correct format", s)
	}

	id, err := strconv.Atoi(sArray[1])
	if err != nil {
		return User{}, fmt.Errorf("record(%s) has non-numeric ID", s)
	}
	return User{Name: strings.TrimSpace(sArray[0]), ID: id}, nil
}
