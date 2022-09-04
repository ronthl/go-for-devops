package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

const flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC

type User struct {
	ID   int
	Name string
}

func (user User) String() string {
	return fmt.Sprintf("%d:%s", user.ID, user.Name)
}

func main() {
	const filePath = "./users.txt"

	// Initialize a list of user data we want to write
	users := []User{
		{ID: 1, Name: "Ron"},
		{ID: 2, Name: "Xi"},
		{ID: 3, Name: "Nana"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Open the file to get *os.File object. *os.File implements io.Writer among other interfaces
	file, err := os.OpenFile(filePath, flags, 0644)
	if err != nil {
		panic(err)
	}

	// Write our users to the file
	if err := writeUsers(ctx, file, users); err != nil {
		panic(err)
	}
	defer file.Close() // Close the file to writing

	// Open the file and write its content to stdout to show it worked
	readFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	io.Copy(os.Stdout, readFile)
}

// writeUsers writes a list of users to 'w' with each entry separated by '\n'.
func writeUsers(ctx context.Context, w io.Writer, users []User) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	for i, user := range users {
		if i != 0 {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		}

		if err := writeUser(ctx, w, user); err != nil {
			return err
		}
	}
	return nil
}

func writeUser(ctx context.Context, w io.Writer, user User) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if _, err := w.Write([]byte(user.String())); err != nil {
		return err
	}
	return nil
}
