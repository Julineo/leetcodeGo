/*
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Note: In the string, each word is separated by single space and there will not be any extra space in the string.
*/

func reverseWords(s string) string {
    r := ""
	for _, v := range strings.Split(s, " ") {
		runes := []rune(v)
		for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
			runes[i], runes[j] = runes[j], runes[i]
			
		}
		r = r + " " + string(runes)
	}
	return strings.Trim(r, " ")
}
