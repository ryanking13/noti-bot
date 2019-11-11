package main

import "fmt"

func main() {
	targets := []string{"182360", "069500"}
	stockInfos, err := poll(targets)

	if err != nil {
		fmt.Println(err)
	} else {
		for _, info := range stockInfos {
			fmt.Println(info.name)
		}
	}
}
