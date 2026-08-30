package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// https://pkg.go.dev/fmt?utm_source=chatgpt.com#Sprintf
	return fmt.Sprintf("Welcome to the Tech Palace, %v", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	result := ""
	for _ = range numStarsPerLine {
		result += "*"
	}
	result += fmt.Sprintf("\n%v\n",welcomeMsg)
	for _ = range numStarsPerLine {
		result += "*"
	}
	return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
