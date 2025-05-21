//go:build exercise_2

package main

import (
	"fmt"
	"io"
	"os"
)

/* Write a function called fileLen that has an input parameter of type string and
returns an int and an error. The function takes in a filename and returns the
number of bytes in the file. If there is an error reading the file, return the error.
Use defer to make sure the file is closed properly.*/

func fileLen(input string) (int, error) {
	f, err := os.Open(input)
	if err != nil {
		fmt.Println("Has error opening the file!")
		return 0, err
	}

	defer f.Close()
	data := make([]byte, 2048) /* create a buffer to be filled with bytes */
	total := 0
	for {
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				return total, err
			}
			break
		}
		total += count
	}
	return total, nil

}

func main() {
	count, err := fileLen("file.txt")
	fmt.Println(count)
	fmt.Println(err)
}
