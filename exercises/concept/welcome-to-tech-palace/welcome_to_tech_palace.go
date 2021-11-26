package techpalace

import (
	"bytes"
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	ret := fmt.Sprint("Welcome to the Tech Palace, ", strings.ToUpper(customer))
	return ret
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars bytes.Buffer
	for i := 0; i < numStarsPerLine; i++ {
		stars.WriteString("*")
	}
	content := []string{stars.String(), welcomeMsg, stars.String()}
	ret := strings.Join(content, "\n")

	return ret

}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	ret := strings.ReplaceAll(oldMsg, "*", "")

	ret = strings.TrimSpace(ret)
	return ret
}
