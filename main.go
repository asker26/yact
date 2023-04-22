package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var dirs = map[string]string{
	"php": "C:\\laragon\\www",
	"go":  "C:\\Users\\rade\\GolandProjects",
	"js":  "C:\\Users\\rade\\WebStormProjects",
	"c#":  "C:\\Users\\rade\\RiderProjects",
}

var ides = map[string]string{
	"php": "phpstorm",
	"go":  "goland",
	"js":  "webstorm",
	"c#":  "rider",
}

var existingDirs []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// prompt the user for input
	fmt.Print("Please, enter a project you would like to open:")

	// read a line of input from the user
	scanner.Scan()
	project := scanner.Text()

	for key, d := range dirs {
		if folderExists(d, project) {
			existingDirs = append(existingDirs, key)
		}
	}

	for len(existingDirs) == 0 {
		fmt.Println("Sorry, but there are no existing projects with the given name...")
		fmt.Println("Would you like to try again? Y\\N")

		scanner.Scan()
		answer := scanner.Text()

		if answer == "N" {
			return
		}

		if answer != "Y" {
			for answer != "Y" && answer != "N" {
				fmt.Println("Sorry I couldn't understand what you wanted...")
				fmt.Println("Would you like to try again? Y\\N")
				scanner.Scan()
				answer = scanner.Text()
			}

			if answer == "N" {
				return
			}
		}

		fmt.Print("Please, enter a project you would like to open:")
		scanner.Scan()
		project = scanner.Text()

		for key, d := range dirs {
			if folderExists(d, project) {
				existingDirs = append(existingDirs, key)
			}
		}
	}

	if len(existingDirs) == 1 {
		key := existingDirs[0]
		cmd := exec.Command("cmd", "/C", fmt.Sprintf("%s %s\\%s", ides[key], dirs[key], project))
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	fmt.Println("there are multiple projects with the given name in different directories")
	fmt.Println("which type of project would you like to open?")

	fmt.Println(strings.Join(existingDirs, ", "))
	fmt.Print("Your choice:")

	scanner.Scan()
	lang := scanner.Text()

	path, exists := dirs[lang]

	for !exists {
		fmt.Println("Please use a valid type")

		fmt.Print("Your choice:")

		scanner.Scan()
		lang = scanner.Text()

		path, exists = dirs[lang]
	}

	cmd := exec.Command("cmd", "/C", fmt.Sprintf("%s %s\\%s", ides[lang], path, project))
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// check if the specified folder exists within the given directory
func folderExists(directory string, folder string) bool {
	cmd := exec.Command("cmd", "/C", fmt.Sprintf("dir %s\\%s", directory, folder))
	_, err := cmd.Output()
	if err != nil {
		return false
	}

	// if we've gone through all the folders and haven't found the specified folder, return false
	return true
}
