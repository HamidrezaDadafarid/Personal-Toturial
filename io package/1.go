package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// This program demonstrates the use of io.Writer, io.Reader, and io.Seeker interfaces.
	// It performs file writing, reading, and seeking operations.

	// ----------------------------------------------
	// Writer interface: wraps the basic Write method.
	// type Writer interface {
	//    Write(p []byte) (n int, err error)
	// }
	// ----------------------------------------------

	// Create a new file named "text.txt" to demonstrate writing operations
	file, _ := os.Create("text.txt")

	// Write "Hello1" to the file using file.Write, which returns the number of bytes written and any error
	n, err := file.Write([]byte("Hello1"))
	fmt.Println(n, err)

	// Wrap the file in an io.Writer interface and write "Hello2" to it
	writer := io.Writer(file)
	n, err = writer.Write([]byte("Hello2"))
	fmt.Println(n, err)

	// Write "Hello3" using io.WriteString, which is a convenience function for writing strings to an io.Writer
	n, err = io.WriteString(file, "Hello3")
	fmt.Println(n, err)

	// Close the file after writing operations are complete
	file.Close()
	// ----------------------------------------------
	// Reader interface: wraps the basic Read method.
	// type Reader interface {
	//    Read(p []byte) (n int, err error)
	// }
	// ----------------------------------------------

	// Open "text.txt" to demonstrate reading operations
	file, _ = os.Open("text.txt")
	reader := io.Reader(file)

	// Read the file content into a buffer of 32 bytes
	buffer := make([]byte, 32)
	n, err = reader.Read(buffer)
	fmt.Printf("buffer={%v}, n={%v}, err={%v}\n", string(buffer), n, err)

	// Close the file after reading
	file.Close()

	// Open the file again for further reading demonstrations
	file, _ = os.Open("text.txt")
	reader = io.Reader(file)

	// Read only one byte using a minimal buffer of size 1
	minBuffer := make([]byte, 1)
	n, err = reader.Read(minBuffer)
	fmt.Printf("buffer={%v}, n={%v}, err={%v}\n", string(minBuffer), n, err)
	file.Close()

	// Demonstrate reading the file byte by byte in a loop
	file, _ = os.Open("text.txt")
	reader = io.Reader(file)
	for {
		n, err = reader.Read(minBuffer)
		fmt.Printf("%v, %v, %v\n", string(minBuffer), n, err)
		// Break the loop when reaching the end of the file (EOF)
		if err != nil {
			break
		}
	}
	file.Close()

	// Use io.ReadAll to read the entire content of the file into a buffer
	file, _ = os.Open("text.txt")
	reader = io.Reader(file)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("Readall buffer={%v}, err={%v}\n", string(buffer), err)
	file.Close()

	// -------------------------------
	// Seeker interface: wraps the basic Seek method.
	// type Seeker interface {
	//    Seek(offset int64, whence int) (int64, error)
	// }
	// -------------------------------

	// Demonstrate the io.Seeker interface to change the reading position within the file
	file, _ = os.Open("text.txt")
	reader = io.Reader(file)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("Readall buffer={%v}, err={%v}\n", string(buffer), err)

	// Cast the reader to an io.Seeker interface to enable seeking
	// The file object returned by os.Open implements both the io.Reader and io.Seeker interfaces. This is because *os.File provides implementations for both Read and Seek methods, making it compatible with both interfaces.
	seeker := reader.(io.Seeker)

	// Reading all the contect of a file and storing it in a buffer is not efficient becuase you have to store that buffer in the memory anf if the file size is massive, it will not be efficient at all.
	// Seek to the beginning of the file and read all contents again
	seeker.Seek(0, io.SeekStart)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("Readall buffer={%v}, err={%v}\n", string(buffer), err)

	// Seek 5 bytes back from the end of the file and read remaining content
	seeker.Seek(-5, io.SeekEnd)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("Readall buffer={%v}, err={%v}\n", string(buffer), err)

	// Seek 5 bytes back from the current position and read remaining content
	seeker.Seek(-5, io.SeekCurrent)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("Readall buffer={%v}, err={%v}\n", string(buffer), err)

	// Close the file after seeking operations are complete
	file.Close()
}
