package helpers

import "log"

// CheckErr checks if an error is not nil. 
// If it is not nil, it will log the error and panic.
func CheckErr(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}