package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mov(a []int, val int) {
	a[2] = a[1]
	a[1] = a[0]
	a[0] = val

}

func rob2(source []*TreeNode, newchilds *[]*TreeNode, current []int) bool {
	j := 0
	cur := 0
	for i := 0; i < len(source); i++ {
		if source[i] == nil {
			break
		}

		cur += source[i].Val

		if source[i].Left != nil {

			if len(*newchilds) == j {
				*newchilds = append(*newchilds, source[i].Left)
			} else {
				(*newchilds)[j] = source[i].Left
			}
			j++
		}

		if source[i].Right != nil {
			if len(*newchilds) == j {
				*newchilds = append(*newchilds, source[i].Right)
			} else {
				(*newchilds)[j] = source[i].Right
			}
			j++
		}

		source[i] = nil
	}

	if current[1] > current[2] {
		cur += current[1]
	} else {
		cur += current[2]
	}

	mov(current, cur)
	return len(*newchilds) == 0 || (*newchilds)[0] == nil

}

// F[i]=max(sum(child), sum(grandchild)+w[i])
// list chil, list grandchild
// wightsearch
func rob1(root *TreeNode) int {
	l := []int{0, 0, 0}
	isL1 := true
	cur := root
	queue1 := []*TreeNode{cur}
	queue2 := []*TreeNode{}

	for {
		if isL1 {
			if rob2(queue1, &queue2, l) {
				break
			}
			isL1 = false
		} else {
			if rob2(queue2, &queue1, l) {
				break
			}
			isL1 = true

		}

	}
	l1 := maxv(l[1], l[2])
	l1 = maxv(l1, l[0])

	return l1

}

func maxv(a1 int, a2 int) int {
	if a1 > a2 {
		return a1
	}
	return a2
}
