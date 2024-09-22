package helpers

import "fmt"

// Catch catches error occurred
func Catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
