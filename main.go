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

	// Initialize a list of users to write into a stream
	users := []User{
		{ID: 1, Name: "Ron"},
		{ID: 2, Name: "Xi"},
		{ID: 3, Name: "Nana"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	file, err := os.OpenFile(filePath, flags, 0644)
	if err != nil {
		panic(err)
	}

	// We will send our User records to be written via inChan
	inChan := make(chan User, 1)

	errChan := writeUser(ctx, file, inChan)

	for _, user := range users {
		select {
		case err := <-errChan:
			fmt.Println("Had error: ", err)
		case inChan <- user:
		}
	}

	// Let our goroutine started by writeUser() know we have finished
	close(inChan)

	// Block here until all our records are written. If the returned value is not nil, we had some type of error
	if err = <-errChan; err != nil {
		fmt.Println("Had error:", err)
		return
	}

	// Open the file to check the result
	readFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, readFile)
}

// writeUser starts goroutine that will write User records received on 'inChan' and
// write them to 'w'. An error encoutered will write an error to the returned channel.
func writeUser(ctx context.Context, w io.Writer, inChan <-chan User) <-chan error {
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			// Close 'w'. If 'w' also implements io.Closer, we will call .Close() on it
			if closer, ok := w.(io.Closer); ok {
				if err := closer.Close(); err != nil {
					// Try to put an error in our errChan, but if we can't just ignore it
					select {
					case errChan <- err:
					default:
					}
				}
			}
			close(errChan) // Close errChan
		}()

		writttenLine := false
		for {
			select {
			// If our Context is cancelled, return
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			case user, ok := <-inChan:
				if !ok {
					return
				}

				// This puts a carriage return before the next entry unless it is the first entry
				if writttenLine {
					if _, err := w.Write([]byte("\n")); err != nil {
						errChan <- err
						return
					}
				}

				if _, err := w.Write([]byte(user.String())); err != nil {
					errChan <- err
					return
				}
				writttenLine = true
			}

		}
	}()
	return errChan
}
