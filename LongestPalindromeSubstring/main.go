package main

func main() {
    // Palindrome - a word, phrase, or sequence that reads the same backwards as forwards, e.g. madam or nurses run.
}

func longestPalindrome(s string) string {
    res := make(map[rune]string)
    var maxLen string

    for _, v := range s {
        if _, ok := res[v]; ok {
            count := len(res[v])
            if count > len(maxLen) {
                maxLen = res[v] + string(v)
            }

            delete(res, v)
        }

        for k := range res {
            res[k] += string(v)
        }

        res[v] = string(v)
    }

    return maxLen
}


func isPalindrome(s string) bool {

    return false
}