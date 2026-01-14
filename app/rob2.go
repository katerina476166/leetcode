package main

//0 - with last
//1- miss last
func rob3(root *TreeNode) []int {

	if root == nil {
		return []int{0, 0}
	}

	l1 := rob3(root.Left)
	l2 := rob3(root.Right)
	l3 := maxk(l1[1]+l2[1], l1[0]+l2[0])
	l3 = maxk(l3, l1[1]+l2[0])
	l3 = maxk(l3, l1[0]+l2[1])

	c := l1[1] + l2[1] + root.Val
	return []int{c, l3}

}

// F[i]=max(sum(child), sum(grandchild)+w[i])
// list chil, list grandchild
// wightsearch
func rob(root *TreeNode) int {

	l := rob3(root)

	return maxk(l[0], l[1])

}

func maxk(a1 int, a2 int) int {
	if a1 > a2 {
		return a1
	}
	return a2
}
