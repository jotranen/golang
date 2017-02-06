package main

import "fmt"

func anagram(s1 string, s2 string) bool {

	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}


func main() {
	fmt.Printf("saippuakauppias vs. saippuakauppias: %t\n", anagram("saippuakauppias", "saippuakauppias"))
	fmt.Printf("saippuakauppias vs. saippuaauppias: %t\n", anagram("saippuakauppias", "saippuaauppias"))
	fmt.Printf("saippuakauppias vs. saippuakauppias: %t\n", anagram("saippuaauppias", "saippuakauppias"))
}
