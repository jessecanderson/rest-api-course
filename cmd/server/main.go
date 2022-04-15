package main

import (
	"fmt"
)

// Run will be responsable for initi and startup of the application
func Run() error {
	fmt.Println("Starting application....")
	return nil
}

func main() {
	fmt.Println("Go REST course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
