// B''H

package main

import (
	"bufio"
	"fmt"
	"jokegetter/query"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

//!+template
const templ = `----------------------------------------
Setup: {{.Setup}}
Punchline: {{.Delivery}}
`

//!-template

//!+exec
var report = template.Must(template.New("joke").
	Parse(templ))

func main() {

	// -- --------------------------------------------------
	// Welcome message
	fmt.Println("I see you're interested in a joke!")
	fmt.Println("Choose a category:")
	fmt.Println("1: Programming\n2: Dark\n3: Misc\n4: Any")

	// Get input
	var jokeCat int

	for true {
		var err error
		jokeCat, err = getOptionInt()

		if err != nil {
			log.Fatal(err)
		}

		if jokeCat > 0 && jokeCat < 5 {
			break
		}
	}

	// -- --------------------------------------------------
	// Offer blacklist options for unwanted joke types
	fmt.Println("Great!")
	fmt.Println("Choose which type (if any) you want to blacklist")
	fmt.Println("Type 'y' to blacklist or 'n' to keep")

	// What you can blacklist
	var blOptions = []string{
		"NSFW",
		"religious",
		"political",
		"racist",
		"sexist",
	}

	// What the user chooses to blacklist
	blacklist := make([]string, 0, len(blOptions))

	// Start the loop through the options
	for _, jokeType := range blOptions {
		for true {
			fmt.Printf("Blacklist %s? ", jokeType)

			option, err := getOptionStr()

			if err != nil {
				log.Fatal(err)
			}

			if option == "y" {
				blacklist = append(blacklist, jokeType)
				break
			} else if option == "n" || option == "" {
				break
			} else {
				fmt.Println("please enter either 'y' or 'n'")
			}
		}
	}

	// -- --------------------------------------------------
	//Run the query to get the joke
	result, err := query.GetJoke(jokeCat, blacklist)

	if err != nil {
		log.Fatal(err)
	}

	// Display the joke according to the template
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func getOptionInt() (int, error) {

	// Initialize input
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Get input
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	// Convert to int
	input = strings.TrimSpace(input)

	option, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	// Return option
	return option, nil

}

func getOptionStr() (string, error) {

	// Initialize input
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// Get input
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	// Convert to int
	option := strings.ToLower(strings.TrimSpace(input))

	// Return option
	return option, nil

}
