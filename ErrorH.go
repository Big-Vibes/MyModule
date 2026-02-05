package main

import (
	"errors"
	"fmt"
)

// Custom error type
type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Function that may fail
func riskyOperation(flag int) error {
	if flag == 0 {
		return MyError{Code: 400, Message: "Invalid flag provided"}
	}
	if flag == 1 {
		return fmt.Errorf("wrapped error: %w", MyError{Code: 500, Message: "Internal failure"})
	}
	return nil
}

func main() {
	tests := []int{0, 1, 2}

	for _, t := range tests {
		err := riskyOperation(t)
		if err != nil {
			// Check if it's our custom error type
			var myErr MyError
			if errors.As(err, &myErr) {
				fmt.Println("Caught MyError specifically:")
				fmt.Printf("Code: %d, Message: %s\n", myErr.Code, myErr.Message)

				// Handle based on code
				switch myErr.Code {
				case 400:
					fmt.Println("→ Handle client-side issue")
				case 500:
					fmt.Println("→ Handle server-side issue")
				}
			} else {
				// Generic error handling
				fmt.Println("Generic error:", err)
			}
		} else {
			fmt.Println("Operation succeeded with flag:", t)
		}
	}
}
