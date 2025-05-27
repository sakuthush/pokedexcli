package main
import "fmt"

func main() {
	for {
		fmt.Print("Pokedex > ")
		var input string
		fmt.Scan(&input)
	    // Clean the input (in this case, just a simple format)
			firstword := input[0]

        // Capture the first word only
           fmt.Printf("Your command was: %s\n", firstWord) // Print the first word
       }
   }
