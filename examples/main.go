package main

import ffmt "gopkg.in/ffmt.v1"

func main() {
	example()
}

func example() {
	m := map[string]interface{}{
		"hello": "w",
		"A": []int{
			1, 2, 3, 4, 5, 6,
		},
	}
	ffmt.Puts(m)
	ffmt.P(m)
	ffmt.Pjson(m)

	m0 := ffmt.ToTable(m, m)
	ffmt.Puts(m0)
	m1 := ffmt.FmtTable(m0)
	ffmt.Puts(m1)

	ffmt.Mark("hello")
}
