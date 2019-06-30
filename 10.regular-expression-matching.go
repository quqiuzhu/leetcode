package main

/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
// Item match item divided by star
type Item struct {
	chars      []byte
	exactMatch bool
}

func patern(p string) []*Item {
	items := make([]*Item, 0)
	var item *Item
	for len(p) > 0 {
		exactMatch := !(len(p) > 1 && p[1] == '*')
		newItem := item == nil || !exactMatch
		if newItem {
			item = &Item{
				chars:      nil,
				exactMatch: exactMatch,
			}
			items = append(items, item)
		}
		c := p[0]
		if c != '*' {
			item.chars = append(item.chars, c)
		}
		p = p[1:]
	}
	return items
}

func charMatch(s, p byte) bool {
	return s == p || p == '.'
}

func match(s string, items []*Item) bool {
	ls, li := len(s), len(items)
	// fmt.Print(s, ":")
	// for i := 0; i < li; i++ {
	// 	fmt.Print(string(items[i].chars), "-", items[i].exactMatch, " ")
	// }
	// fmt.Println()
	if li == 0 {
		return ls == 0
	}

	p := items[0]
	if p.exactMatch {
		lc := len(p.chars)
		if ls < lc {
			return false
		}
		for i := 0; i < lc; i++ {
			if !charMatch(s[i], p.chars[i]) {
				return false
			}
		}
		return match(s[lc:], items[1:])
	}

	c := p.chars[0]
	lc := len(p.chars) - 1

	matchCharCount := 0
	for i := 0; i < ls; i++ {
		if !charMatch(s[i], c) {
			break
		}
		matchCharCount++
	}

	if lc == 0 {
		for i := 0; i <= matchCharCount; i++ {
			if match(s[i:], items[1:]) {
				return true
			}
		}
		return false
	}

	for i := 0; i <= matchCharCount && i <= ls-lc; i++ {
		partialMatch := true
		for j := 0; j < lc; j++ {
			if !charMatch(s[i+j], p.chars[1+j]) {
				partialMatch = false
				break
			}
		}
		if partialMatch {
			if match(s[i+lc:], items[1:]) {
				return true
			}
		}
	}
	return false
}

func isMatch(s string, p string) bool {
	items := patern(p)
	return match(s, items)
}
