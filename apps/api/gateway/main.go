package main

import (
	"course-explorer-monorepo/libs/api"
	"fmt"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(api.Api("Hello"))
}
