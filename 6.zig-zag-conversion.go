package main

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// func column(l, numRows int) int {
// 	f := 2*numRows - 2
// 	if f == 0 {
// 		return 1
// 	}

// 	r := l % f
// 	w := l / f

// 	n := 0
// 	if r == 0 {
// 		n = 0
// 	} else if r-numRows <= 0 {
// 		n = 1
// 	} else {
// 		n = r - (numRows - 1)
// 	}
// 	return w*(numRows-1) + n
// }

func row(l, numRows int) int {
	f := 2*numRows - 2
	v := l % f
	var result int
	if v == 0 {
		v = f
	}
	if v <= numRows {
		result = v
	} else {
		result = 2*numRows - v
	}
	return result
}

func convert(s string, numRows int) string {
	// compute numColumns
	l := len(s)
	// numColumns := column(l, numRows)
	if numRows == 1 {
		return s
	}

	rows := make(map[int][]byte, 0)
	buf := make([]byte, 0)
	for i := 0; i < l; i++ {
		r := row(i+1, numRows)
		// fmt.Println(string(s[i]), "r: ", r, "c: ", c)
		if _, ok := rows[r]; !ok {
			// rows[r] = make([]byte, numColumns)
			rows[r] = make([]byte, 0)
		}
		rows[r] = append(rows[r], s[i])

	}
	for i := 1; i <= numRows; i++ {
		buf = append(buf, rows[i]...)
	}
	return string(buf)
}
