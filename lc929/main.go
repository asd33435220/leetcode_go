package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}

	res := numUniqueEmails(emails)
	fmt.Println(res)
}
func handleEmail(email string) string {
	arr2 := strings.Split(email, "@")
	arr := strings.Split(arr2[0], ".")
	email = strings.Join(arr, "")
	arr = strings.Split(email, "+")
	name := arr[0]
	return name + "@" + arr2[1]
}
func numUniqueEmails(emails []string) int {
	hash := make(map[string]bool)
	for _, email := range emails {
		res := handleEmail(email)
		hash[res] = true
	}
	return len(hash)
}
