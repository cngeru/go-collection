package main

import "fmt"

func main() {
	profile := map[string]string{
		"name": "Dennis Chege",
		"age":  "22",
		"home": "murang'a",
	}
	// delete(profile, "home")
	// fmt.Println(profile)
	print(profile)

}

func print(profile map[string]string) {
	for key, val := range profile {
		fmt.Println(key, val)
	}
}
