package main

func lengthOfLongestSubstring(s string) int {
	elements := map[byte]interface{}{}
	start := 0
	end := 0
	max := 0
	for start < len(s) && end < len(s) {
		if _, hasElements := elements[s[end]]; !hasElements {
			elements[s[end]] = nil
			end++
			length := end - start
			if length > max {
				max = length
			}
		} else {
			delete(elements, s[start])
			start++
		}
	}
	return max
}

func main() {

	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("bbbbb"))
	println(lengthOfLongestSubstring("pwwkew"))
	println(lengthOfLongestSubstring("au"))
}
