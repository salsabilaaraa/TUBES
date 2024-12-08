package main

import "fmt"

const NMAX int = 20

type klub struct {
	nama      string
	point, gd int
}
type tabKlub [NMAX]klub

func main() {
	var k tabKlub
	var n int = 3
	for i := 0; i < n; i++ {
		fmt.Scan(&k[i].nama, &k[i].point, &k[i].gd)
	}
	fmt.Println(juara(k, n))
}

func juara(k tabKlub, n int) string {
	var i, idx int
	idx = 1
	for i = 2; i < n; i++ {
		if k[i].point > k[idx].point {
			idx = i
		} else if k[i].point == k[idx].point && k[i].gd > k[idx].gd {
			idx = i
		}
	}
	return k[idx].nama
}
