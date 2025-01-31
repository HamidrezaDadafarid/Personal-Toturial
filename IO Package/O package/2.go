package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// countAlphabets reads data from an io.Reader and counts the number of alphabetic characters (both uppercase and lowercase).
func countAlphabets(r io.Reader) (int, error) {
	count := 0                   // Initialize a counter for alphabetic characters
	buffer := make([]byte, 1024) // Create a buffer to read data in chunks

	// Loop to read data in chunks until the end of the file or reader is reached
	for {
		n, err := r.Read(buffer) // Read up to 1024 bytes into buffer

		// Loop through the bytes read and count alphabetic characters
		for _, l := range buffer[:n] {
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
				count++
			}
		}

		// Check if the end of the file or reader has been reached
		if err == io.EOF {
			return count, nil // Return the count if at the end
		}

		// If another error occurs, return it
		if err != nil {
			return 0, err
		}
	}
}

// writeString writes a string to an io.Writer and returns the number of bytes written.
func writeString(s string, w io.Writer) (int, error) {
	// Convert the string to a byte slice and write it to the writer
	n, err := w.Write([]byte(s))
	if err != nil {
		return 0, err
	}
	return n, nil // Return the number of bytes written
}

func main() {
	// Open a file in read mode and count alphabetic characters in it using countAlphabets
	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	n, err := countAlphabets(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Count of letters:", n)
	file.Close()

	// Use a strings.Reader as an io.Reader and count alphabetic characters in a string
	reader := strings.NewReader("Hello, World!")
	n, err = countAlphabets(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Count of letters:", n)

	// Open or create a file in write mode and write a string to it using writeString
	file, err = os.Create("text.txt")
	if err != nil {
		panic(err)
	}
	n, err = writeString("Hello1Hello2Hello3", file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of written bytes:", n)
	file.Close()
}
