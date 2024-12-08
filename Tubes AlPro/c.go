package main

import "fmt"

const NMAX int = 1024

type move struct {
	from, to string
	score    float64
	legal    bool
}

type tabMove [NMAX]move

func main() {
	var chess tabMove
	var n int = 3
	for i := 0; i < n; i++ {
		fmt.Scan(&chess[i].from, &chess[i].to, &chess[i].score, &chess[i].legal)
	}
	fmt.Println(bestMove(chess, n))
}

func bestMove(c tabMove, n int) (string, string) {
	var ketemu bool = false
	var i, idx int
	i = 2
	idx = 1
	for !ketemu && i < n {
		ketemu = c[i].score > c[idx].score && c[i].legal
		if ketemu {
			idx = i
		}
		i++
	}
	return c[idx].from, c[idx].to
}
