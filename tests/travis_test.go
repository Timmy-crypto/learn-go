package tests

/*func Test_Travis(t *testing.T) {
	a := -5 >> 2
	fmt.Printf("the a is: %d\r\n", a)
	b := -5 / 4
	fmt.Printf("the b is: %d\r\n", b)
	fmt.Printf("test travis\r\n")
	//var c []int
}

type node struct {
	value int
	next  *node
}
type nodeList struct {
	head *node
}

func nodeListAdd(a *nodeList, b *nodeList) *nodeList {
	resultNodeList := nodeList{
		head: &node{},
	}
	aCurNode, bCurNode, rCurNode := a.head, b.head, resultNodeList.head
	increaseNumber := 0

	for {
		if aCurNode == nil && bCurNode != nil {
			aCurNode = &node{}
		} else if bCurNode == nil && aCurNode != nil {
			bCurNode = &node{}
		} else if aCurNode == nil && bCurNode == nil {
			break
		}

		sum := aCurNode.value + bCurNode.value + increaseNumber
		if sum >= 10 {
			increaseNumber = 1
			rCurNode.value = sum - 10
		} else {
			rCurNode.value = sum
			increaseNumber = 0
		}

		rCurNode.next = &node{}
		rCurNode = rCurNode.next
		aCurNode = aCurNode.next
		bCurNode = bCurNode.next
	}
	return &resultNodeList
}*/

/*
Given an interface for a two-dimensional table of integers:

public interface Table {
// Set the cell at (x, y) to value.
void set(int x, int y, int value);

// Compute the sum of values from (0, 0) to (x, y), inclusive.
int sum(int x, int y);
}

Example:

2 3 5
2 2 5
5 5 5

sum(1, 1) = 2 + 3 + 2 + 2 = 9


1. implement the interface:
*/

type TableValue struct {
	x     int
	y     int
	value int
}
type TDTable struct {
	Values   []*TableValue
	SumValue int
}

func (t *TDTable) set(x int, y int, value int) {
	existFlag := false
	for _, val := range t.Values {
		if val.x == x && val.y == y {
			existFlag = true
			val.value = value
		}
	}
	if !existFlag {
		t.Values = append(t.Values, &TableValue{
			x:     x,
			y:     y,
			value: value,
		})
	}
}

func (t *TDTable) sum(x int, y int) {
	for _, val := range t.Values {
		if val.x <= x && val.y <= y {
			t.SumValue += val.value
		}
	}
}
