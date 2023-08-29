package tests

/*
[[0,1], [1, 2], [2, 3], [3, 4], [4, 5], [2, 6], [6, 7], [2, 8]]

=> [0,1,2,3,4,5]

The input is a list of tuple where the first element represents the parent block hash and the second element the child block hash. The list describes the whole structure of a blockchain with forks. Please write a function that takes the input and returns the longest chain.
*/

func InitNodeFromTupleList(tupleList [][2]uint) (map[uint][]uint, uint) {
	Nodes := make(map[uint][]uint, 0)
	//NodeParentNode := make(map[uint]*Node, 0)
	//Nodes[startRoot.CurrentHash] = startRoot

	for _, tuple := range tupleList {
		parent, child := tuple[0], tuple[1]
		_, ok := Nodes[parent]
		if !ok {
			Nodes[parent] = make([]uint, 0)
		}
		Nodes[parent] = append(Nodes[parent], child)
	}

	return Nodes, tupleList[0][0]
}

func GetMaxPath(chains map[uint][]uint, root uint) []uint {
	result := make([]uint, 0)
	result = append(result, root)
	ChildNodeMaxLength := 0

	childNodeMaxPath := make([]uint, 0)
	for _, node := range chains[root] {
		childMaxPath := GetMaxPath(chains, node)
		if len(childMaxPath) > ChildNodeMaxLength {
			ChildNodeMaxLength = len(childMaxPath)
			childNodeMaxPath = childMaxPath
		}
	}
	result = append(result, childNodeMaxPath[:]...)
	return result
}
