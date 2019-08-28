package main

/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

type graphNodeII struct {
	Level   int
	Word    []byte
	Visited bool
}

func (n *graphNodeII) expand(m map[string]*graphNodeII) []*graphNodeII {
	var c byte
	var ok bool
	var next *graphNodeII
	var children []*graphNodeII
	var s string
	for i := 0; i < len(n.Word); i++ {
		for c = 'a'; c <= 'z'; c++ {
			if c == n.Word[i] {
				continue
			}
			c, n.Word[i] = n.Word[i], c
			s = string(n.Word)
			if next, ok = m[s]; ok && !next.Visited {
				next.Level = n.Level + 1
				next.Visited = true
				children = append(children, next)
			}
			c, n.Word[i] = n.Word[i], c
		}
	}
	return children
}

func (n *graphNodeII) isTarget(target string) bool {
	return string(n.Word) == target
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	nodes := make(map[string]*graphNodeII)
	for _, word := range wordList {
		nodes[word] = &graphNodeII{Word: []byte(word)}
	}
	if _, ok := nodes[endWord]; !ok {
		return 0
	}

	var it *graphNodeII
	var p, q, children []*graphNodeII
	p = []*graphNodeII{
		&graphNodeII{
			Word:    []byte(beginWord),
			Visited: true,
		},
	}
	for p != nil {
		for _, it = range p {
			if it.isTarget(endWord) {
				return it.Level + 1
			}
			children = it.expand(nodes)
			q = append(q, children...)
		}
		p = q
		q = nil
	}
	return 0
}
