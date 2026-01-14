package main

//121313
//12321 -
//1222222 - 1211111
//1221
//321
//414
//

///1224 - 1212
///444
///123321 1234
//когда мин, то как 1 1223 1212
//когда максимум 2342 1231  23442 11
/*
найти последовательность от мин до мин (2 поинтера) p1, p2, pmax
посчитать в ней сумму,перейти p2
краевые условия
если воозраст
*/
/*рассчет суммы последовательности
min1 .. min2
возр 567 - сумма прогрессии арифметической  n*(a1+an)/2 3*(2+4)/2=9 //2,3,4
убыв  432
возрУбыв 4 8 9 5 - 2,3, 3, 2 делим на 2 по максимуму, считаем сумму, складываем
*/
func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	p1 := 0
	p2 := 0
	pmax := 0
	sum := 0

	for p2 < len(ratings) {

		if p2 == 0 && ratings[0] <= ratings[1] {
			sum += 1
			p1++
			pmax = 0
		} else if p2 == len(ratings)-1 {

			if ratings[p2] <= ratings[p2-1] {
				sum += calcSum(pmax, ratings[p1:p2]) + 1
			} else {
				sum += calcSum(pmax, ratings[p1:p2+1])
			}
			break
		} else if p1 != p2 && ratings[p2-1] > ratings[p2] && ratings[p2] < ratings[p2+1] {
			sum += calcSum(pmax, ratings[p1:p2]) + 1
			p1 = p2 + 1
			pmax = 0
		} else if p2 != 0 && ratings[p2-1] >= ratings[p2] && ratings[p2] <= ratings[p2+1] {
			sum += calcSum(pmax, ratings[p1:p2]) + 1
			p1 = p2 + 1
			pmax = 0
		}

		p2++
		if ratings[pmax+p1] < ratings[p2] {
			pmax = p2 - p1
		}

	}

	return sum
}

func calcSum(pmax int, ratings []int) int {
	if pmax == 0 || pmax == len(ratings)-1 {
		return calcMonoSum(ratings)
	}
	l := calcMonoSum(ratings[0:pmax])
	r := calcMonoSum(ratings[pmax+1:])
	m := 0
	if ratings[pmax+1] == ratings[pmax] {
		m = max(pmax+2, len(ratings[pmax+1:])+1)
	} else {
		m = max(pmax+2, len(ratings[pmax+1:])+2)
	}
	return l + r + m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 567 - сумма прогрессии арифметической  n*(a1+an)/2 3*(2+4)/2=9 //2,3,4
func calcMonoSum(ratings []int) int {
	return len(ratings) * (2 + len(ratings) + 1) / 2
}
