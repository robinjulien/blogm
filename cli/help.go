package cli

import "fmt"

// prints the help
func help() {
	fmt.Println(`Blogm help page :
You can use one of the following command :
- help : get the help you need
- init : init a blog in the current directory`)
}
