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

package main

import (
	"fmt"
	"os"
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

}
