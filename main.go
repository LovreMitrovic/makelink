package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
)

func createShortcut(path string, link string) error {
	//Check if file already exists
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		return err
	}

	//Check link validity
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return err
	}

	//Write to file
	data := []byte("<html><head><script>window.location.href=\"" + link + "\"</script></head></html>")
	errFile := os.WriteFile(path, data, 0644)
	if errFile != nil {
		return errFile
	}

	return nil
}

func main() {

	//Parsing the arguments
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Invalid number of arguments\nUsage:\n\tmakelink <name> <link>\n\nLegend:\n\t<name>\tName for file\n\t<link>\tLink to website")
	}

	name := args[0]
	link := args[1]

	//Create path for new file
	path := name + ".html"

	err := createShortcut(path, link)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Shortcut created " + path)

}
