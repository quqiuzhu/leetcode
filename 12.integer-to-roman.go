package main

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

func intToRoman(num int) string {
	var roman []byte
	chars := [][]string{
		[]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"M", "MM", "MMM"},
	}

	i, j, d := 3, 0, 1000
	for num != 0 {
		j = num / d
		if j != 0 {
			for k := 0; k < len(chars[i][j-1]); k++ {
				roman = append(roman, chars[i][j-1][k])
			}
		}

		num = num - d*j
		i--
		d /= 10
	}
	return string(roman)
}
