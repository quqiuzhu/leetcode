package main

/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */

func formatLine(words []string, spaces, maxWidth, blanks int) string {
	chars := make([]byte, maxWidth)
	for i := 0; i < maxWidth; i++ {
		chars[i] = ' '
	}
	copy(chars, words[0])
	spaces -= blanks
	j := maxWidth - blanks
	for i := len(words) - 1; i > 0; i-- {
		lineSpace := 0
		if i != 0 {
			lineSpace = spaces / i
		}
		copy(chars[j-len(words[i]):], words[i])
		j = j - lineSpace - len(words[i])
		spaces -= lineSpace
	}
	return string(chars)
}

func fullJustify(words []string, maxWidth int) []string {
	var line, lines []string
	var i, lineWidth, newWidth, spaces int
	for i = 0; i < len(words); i++ {
		if lineWidth == 0 {
			lineWidth = len(words[i])
			line = append(line, words[i])
			continue
		}

		newWidth = lineWidth + len(words[i]) + 1
		if newWidth <= maxWidth {
			lineWidth = newWidth
			line = append(line, words[i])
		} else {
			spaces = maxWidth - (lineWidth - (len(line) - 1))
			lines = append(lines, formatLine(line, spaces, maxWidth, 0))
			lineWidth = len(words[i])
			line = nil
			line = append(line, words[i])
		}
	}

	spaces = maxWidth - (lineWidth - (len(line) - 1))
	lines = append(lines, formatLine(line, spaces, maxWidth, spaces-(len(line)-1)))

	return lines
}
