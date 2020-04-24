package merkletree

//Compare two trees and return all instances of discrepancies 
func CompareTwoTrees(root1 *merkletree.Node, root2 *merkletree.Node) [][]merkletree.Content {
	discrepancies := [][]merkletree.Content{}
	compareTwoNodes(root1, root2, &discrepancies)

	return discrepancies
}

func compareTwoNodes(node1 *merkletree.Node, node2 *merkletree.Node, discrepancies *[][]merkletree.Content) {
	fmt.Println(node1, node2)
	if node1.IsLeaf() || node2.IsLeaf() {
		if bytes.Compare(node1.Hash, node2.Hash) != 0 {
			newEntry := []merkletree.Content{node1.C, node2.C}
			*discrepancies = append(*discrepancies, newEntry)
		}
	}else if bytes.Compare(node1.Hash, node2.Hash) == 0 {
		return
	} else {
		compareTwoNodes(node1.Left, node2.Left, discrepancies)
		compareTwoNodes(node1.Right, node2.Right, discrepancies)
	}
}
