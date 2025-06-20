package funcvx

import "fmt"

// GoSafe recovers from panics and logs them
// It's useful for running functions in goroutines
// @param fn func() - The function to run in a goroutine
func GoSafe(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("❌ Panic:", r)
			}
		}()
		fn()
	}()
}
