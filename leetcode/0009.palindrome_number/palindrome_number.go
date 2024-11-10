package leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		// 負の数は回文になり得ない
		return false
	}

	original := x
	reversed := 0

	// 数値を逆順にする
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 逆順にした数値が元の数値と同じかチェック
	return original == reversed
}

// func isPalindrome(x int) bool {
//     num := x
// var palindrome int
// for i := 0; num > 0; i++ {
//     mod := num % 10
//     num = num / 10
//     palindrome = palindrome*10 + mod
// }

// return palindrome == x

// }


func isPalindrome1(x int) bool {
    if x < 0 {
        return false
    }
    if x%10 == 0 {
        return false
    }

    ss := strconv.Itoa(x)
    length := len(ss)
    for i := 0; i < length/2; i++ {
        if ss[i] != ss[length-i-1] {
            return false
        }
    }

    return true
}

