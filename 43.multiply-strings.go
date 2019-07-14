package main

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

func add(r1 []byte, r2 []byte) []byte {
	i, l1, l2 := 0, len(r1), len(r2)
	var remainder, t1, t2 byte
	var r []byte

	remainder = 0
	for i = 0; i < l1 || i < l2; i++ {
		t1, t2 = 0, 0
		if i < l1 {
			t1 = r1[i]
		}
		if i < l2 {
			t2 = r2[i]
		}
		t1 = t1 + t2 + remainder
		remainder = t1 / 10
		r = append(r, t1%10)
	}

	if remainder != 0 {
		r = append(r, remainder)
	}
	return r
}

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	if l1 == 1 && num1[0] == '0' || l2 == 1 && num2[0] == '0' {
		return "0"
	}

	n1, n2 := make([]byte, 0, l1), make([]byte, 0, l2)
	var i, j int
	for i = 0; i < l1; i++ {
		n1 = append(n1, num1[l1-1-i]-'0')
	}
	for i = 0; i < l2; i++ {
		n2 = append(n2, num2[l2-1-i]-'0')
	}

	var final, r []byte
	var remainder, t byte
	for j = 0; j < l2; j++ {
		r = make([]byte, 0, l1)
		for i = 0; i < j; i++ {
			r = append(r, byte(0))
		}

		remainder = byte(0)
		for i = 0; i < l1; i++ {
			t = n2[j]*n1[i] + remainder
			remainder = t / 10
			r = append(r, t%10)
		}
		if remainder != 0 {
			r = append(r, remainder)
		}

		final = add(final, r)
	}

	l2 = len(final)
	for i = 0; i < (l2+1)/2; i++ {
		final[i], final[l2-1-i] = final[l2-1-i]+'0', final[i]+'0'
	}
	return string(final)
}
