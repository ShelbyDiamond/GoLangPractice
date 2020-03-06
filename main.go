// Resources I used in case I need to look back on it:
// https://osquery.readthedocs.io/en/1.4.7/installation/cli-flags/
// https://amehta.github.io/posts/2019/03/wc-implementation-in-go-lang/
// https://golang.org/pkg/os/
// https://pkg.go.dev/os?tab=doc
// https://books.google.com/books?id=HW-5SZ1HKusC&pg=PA345&lpg=PA345&dq=what+is+%258d+in+go&source=bl&ots=Frb016T1jt&sig=ACfU3U0Sl5P4qnELMtqls9rmhizgoGdAOw&hl=en&sa=X&ved=2ahUKEwibuM3YooboAhVqHzQIHYP8AakQ6AEwAnoECAYQAQ#v=onepage&q=what%20is%20%258d%20in%20go&f=false
// https://golang.org/pkg/fmt/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// https://golang.org/pkg/os/
	// Any argument typed AFTER the you list your arguments get passed into the os.Args array with the name of the program at the first index

	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		fmt.Println("I'm sorry, your file is missing. Try running `go run <current_file> <file_name>` for better results")
		os.Exit(1)
	}

	// Once you have the name of the file, you can open it with the os.Open function. If it fails to open, it will give a basic reason why.
	// https://pkg.go.dev/os?tab=doc

	file, error := os.Open(fileName)
	if error != nil {
		fmt.Println("Error ", error)
	}

	//  This code is going to scan the file line by line and print the output to the console using the bufio package.

	scanner := bufio.NewScanner(file)

	// sets the number of lines, words, and characters to 0 to start

	lines, words, characters := 0, 0, 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		// The total number of characters will be the length of each line we iterate through while reading the file and
		lines++

		line := scanner.Text()
		characters += len(line)

		// In order to get the number of words we need to split the line based on when the spacebar was used and count the total number of splits.

		splitLines := strings.Split(line, " ")
		words += len(splitLines)
	}

	// now that we have the answer, we need to display it to the console
	// %8d = 1 for each lines, words, and characters
	// %s = will show the string

	fmt.Printf("lines: %d\nWords: %d\nCharacters: %d\nFile Name Read: %s \n", lines, words, characters, fileName)
}