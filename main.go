package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))
	fmt.Println(TwoToOne("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"))
}
func TwoToOne(s1 string, s2 string) string {
	str := s1 + s2
	splitS1 := strings.Split(str, "")
	for i := 0; i < len(splitS1); i++ {
		if strings.Count(str, splitS1[i]) > 1 {
			str = CheckStringDuplicate(splitS1[i], str, strings.Count(str, splitS1[i]))
		}
	}
	s := []string(strings.Split(str, ""))
	sort.Strings(s)
	result := strings.Join(s, "")
	return result
}

func CheckStringDuplicate(textToCheck string, text string, amount int) string {
	return strings.Replace(text, textToCheck, "", amount-1)
}
