// Resources I used in case I need to look back on it:

// https://osquery.readthedocs.io/en/1.4.7/installation/cli-flags/
// https://amehta.github.io/posts/2019/03/wc-implementation-in-go-lang/
// https://golang.org/pkg/os/
// https://pkg.go.dev/os?tab=doc
// https://books.google.com/books?id=HW-5SZ1HKusC&pg=PA345&lpg=PA345&dq=what+is+%258d+in+go&source=bl&ots=Frb016T1jt&sig=ACfU3U0Sl5P4qnELMtqls9rmhizgoGdAOw&hl=en&sa=X&ved=2ahUKEwibuM3YooboAhVqHzQIHYP8AakQ6AEwAnoECAYQAQ#v=onepage&q=what%20is%20%258d%20in%20go&f=false
// https://golang.org/pkg/fmt/
// https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
// https://gobyexample.com/command-line-flags
// https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
// https://golang.org/cmd/compile/
// https://en.wikipedia.org/wiki/Wc_(Unix)
// https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
// https://golang.org/pkg/flag/
// https://www.youtube.com/watch?v=v_8584-jm7I
// https://flaviocopes.com/go-command-line-flags/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// set of loolean flags
	countLines := flag.Bool("l", false, "count lines of a file)"
	
	// parse the rest of the non-flag items
	flag.Parse()
	// get the info after the flag ([]string)
	fileName := flag.Args()

	// deference countLinces
	if *countLines {
		if len(fileName) < 1 {
			fmt.Println("Enter a file name")
			os.Exit(1)
		}
		// get teh first item
		fileName := fileName[0]
		file, err := os.Open(fileName)

		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		
		scanner := bufio.NewScanner(file)
		lines := 0

		for scanner.Scan() {
			lines++
		}
		fmt.Printf("%d lines in %s\n", lines, fileName)
	} else {
		// didnt't have -1 set
		flag.Usage()
	}
}
