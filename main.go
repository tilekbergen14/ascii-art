package main

import (
	"ascii-art/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Oopps, invalid arguments!")
		return
	}
	versions := []string{"shadow", "thinkertoy", "standard"}
	str := os.Args[1]
	version := os.Args[2]
	for _, value := range versions {
		if value == version {
			validFile := functions.FileCheck(value)
			if !validFile {
				fmt.Println("Not a valid file!")
				return
			}
			art, err := os.ReadFile(value + ".txt")
			if err != nil {
				fmt.Println("error occured")
				return
			}
			arr := functions.MakeArt(str, art)
			functions.Write(arr)
			return
		}
	}
	fmt.Println("Invalid format! --> ", version)
}
