package main

import  (
	"fmt"
	"bufio"
	"os"
	"regexp"
)

/*  This version reads a string and then does some regex to parse the input
		If the input is a number (with decimal or not), this returns the part before the decimal
		If the input leads with number, and has junk characters after, it returns the number before any decimal
		If the input leads with not-a-number, this returns "Not a number"
	This is overkill...but just wanted to try some things :)
*/

func main() {
	buf := bufio.NewReader(os.Stdin)	// Described: https://pkg.go.dev/bufio#NewReader

	fmt.Println("Enter number for truncation:")
    s, err := buf.ReadString('\n')		// Described: https://pkg.go.dev/bufio#Reader.ReadString
    if err != nil {
        fmt.Println(err)
    } else {
		rex := regexp.MustCompile("^(\\d+)\\D")   	// Docs: https://pkg.go.dev/regexp; (\\d+) is the submatch for digits \\D is for not-a-digit
		finds := rex.FindStringSubmatch(s)			// returns a list with [full match, submatch...]
		if ( finds == nil ) {
			fmt.Println("Not a number")
		} else {
			fmt.Println(finds[1])
		}
    }
}