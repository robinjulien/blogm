package cli

import (
	"fmt"
	"os"
)

// prints the help
func help() {
	wd, err := os.Getwd()
	check(err)

	fmt.Printf(`Blogm help page :
You can use one of the following command :
- help : get the help you need
- init : init a blog in the current directory

You are located at %s`, wd)
	fmt.Println()
}
