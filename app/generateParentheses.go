package main

func generateParenthesis(n int) []string {

	a := make([]string, 0)
	a = generateParenthesis2(0, 0, n, a, "")
	return a
}

func generateParenthesis2(open int, close int, n int, ans []string, path string) []string {
	if len(path) == 2*n {
		ans = append(ans, path)
		return ans
	}

	if open < n {
		ans = generateParenthesis2(open+1, close, n, ans, path+"(")
	}

	if open > close {
		ans = generateParenthesis2(open, close+1, n, ans, path+")")

	}
	return ans
}

//если возможно, то добавляем еще одну скобку открытую
//добавояем скобку закрытую eсли можно
//усли строка сформировалась,  то к списку
//начинаем с пустой строки
