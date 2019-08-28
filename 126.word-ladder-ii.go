package main

/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */

type graphNode struct {
	Level   int
	Word    []byte
	Parents []*graphNode
	Visited bool
	End     bool
}

func (n *graphNode) expand(m map[string]*graphNode, endWord string) []*graphNode {
	var c byte
	var ok bool
	var next *graphNode
	var children []*graphNode
	var s string
	for i := 0; i < len(n.Word); i++ {
		for c = 'a'; c <= 'z'; c++ {
			if c == n.Word[i] {
				continue
			}
			c, n.Word[i] = n.Word[i], c
			s = string(n.Word)
			if s == endWord {
				n.End = true
				c, n.Word[i] = n.Word[i], c
				break
			}
			if next, ok = m[s]; ok {
				if !next.Visited {
					next.Parents = append(next.Parents, n)
					next.Level = n.Level + 1
					next.Visited = true
					children = append(children, next)
				} else if next.Level == n.Level+1 {
					next.Parents = append(next.Parents, n)
				}
			}
			c, n.Word[i] = n.Word[i], c
		}
	}
	return children
}

func (n *graphNode) paths() [][]string {
	if n.Parents == nil {
		return [][]string{
			[]string{string(n.Word)},
		}
	}
	var p *graphNode
	var rs, r [][]string
	var i int
	for _, p = range n.Parents {
		r = p.paths()
		for i = 0; i < len(r); i++ {
			r[i] = append(r[i], string(n.Word))
		}
		rs = append(rs, r...)
	}
	return rs
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	nodes := make(map[string]*graphNode)
	for _, word := range wordList {
		nodes[word] = &graphNode{Word: []byte(word)}
	}
	if _, ok := nodes[endWord]; !ok {
		return nil
	}

	var it *graphNode
	var p, q, ends, children []*graphNode
	var ended bool
	p = []*graphNode{
		&graphNode{
			Word:    []byte(beginWord),
			Visited: true,
		},
	}

	for p != nil && !ended {
		for _, it = range p {
			children = it.expand(nodes, endWord)
			q = append(q, children...)
			if it.End {
				ended = true
				ends = append(ends, it)
			}
		}
		p = q
		q = nil
	}

	if ends == nil {
		return nil
	}

	end := &graphNode{
		Word:    []byte(endWord),
		Parents: ends,
	}
	return end.paths()
}
