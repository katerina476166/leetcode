package main

import (
	"strings"
	"unicode/utf8"
)

func shortestCommonSupersequence(str1 string, str2 string) string {
	i := 1
	j := 1

	n1 := utf8.RuneCountInString(str1)
	n2 := utf8.RuneCountInString(str2)
	matrix := make([][]A, n1+1)

	for i := 0; i < n1+1; i++ {
		matrix[i] = make([]A, n2+1)
		matrix[i][0] = A{i, 0, 0}
	}

	for i := 0; i < n2+1; i++ {
		matrix[0][i] = A{i, 0, 0}
	}

	for i := 0; i < n1+1; i++ {
	}
	for _, rw1 := range str1 {
		for _, rw2 := range str2 {
			if rw1 == rw2 {
				matrix[i][j].v = matrix[i-1][j-1].v + 1 //то записывае
				matrix[i][j].p1 = i - 1
				matrix[i][j].p2 = j - 1
			} else {
				if matrix[i-1][j].v < matrix[i][j-1].v {
					matrix[i][j].p1 = i - 1
					matrix[i][j].p2 = j
					matrix[i][j].v = matrix[i-1][j].v
				} else {
					matrix[i][j].p1 = i
					matrix[i][j].p2 = j - 1
					matrix[i][j].v = matrix[i][j-1].v
				}
				matrix[i][j].v += 1

			}
			j++
		}
		j = 1
		i++
	}

	r1 := []rune(str1)
	r2 := []rune(str2)
	var sb strings.Builder
	createCommonWord(&r1, &r2, &matrix, n1, n2, &sb)

	return sb.String()
}

type A struct {
	v  int
	p1 int
	p2 int
}

func createCommonWord(str1 *[]rune, str2 *[]rune, m *[][]A, i int, j int, builder *strings.Builder) {
	p1 := (*m)[i][j].p1
	p2 := (*m)[i][j].p2

	if i == 0 && j == 0 {
		return
	}

	if i == 0 {
		for k := 0; k < j; k++ {
			runeValue := (*str2)[k]
			builder.WriteRune(runeValue)
		}
		return
	}

	if j == 0 {
		for k := 0; k < i; k++ {
			runeValue := (*str1)[k]
			builder.WriteRune(runeValue)
		}
		return
	}

	createCommonWord(str1, str2, m, p1, p2, builder)

	if p1 == i {
		r := ((*str2)[j-1])
		builder.WriteRune(r)
	} else {
		r := (*str1)[i-1]
		builder.WriteRune(r)
	}

}
