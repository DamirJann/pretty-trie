package drawing

import (
	"bytes"
	"fmt"
	"github.com/DamirJann/pretty-trie/pkg/entity"
)

func Visualize(node []entity.Node, edge []entity.Edge) (string, error) {
	var res bytes.Buffer
	res.WriteString(node[0].Label + "\n")
	dfs(&res, node[0], node, edge, 0, []bool{})
	return res.String(), nil
}

func findByFrom(from int, edge []entity.Edge) (output []entity.Edge) {
	for _, e := range edge {
		if e.From == from {
			output = append(output, e)
		}
	}
	return output
}

func findById(id int, node []entity.Node) *entity.Node {
	for _, n := range node {
		if n.Id == id {
			return &n
		}
	}
	return nil
}

func dfs(out *bytes.Buffer, currentNode entity.Node, node []entity.Node, edge []entity.Edge, depthLevel int, mark []bool) {
	selected := findByFrom(currentNode.Id, edge)

	for ind, n := range selected {
		startSymbol := "├"
		isLastFileInDir := true

		if len(selected) == ind+1 {
			startSymbol = "└"
			isLastFileInDir = false
		}

		for i := 0; i < depthLevel; i++ {
			if mark[i] {
				fmt.Fprintf(out, "│\t")
			} else {
				fmt.Fprintf(out, "\t")
			}
		}

		_, _ = fmt.Fprintf(out, "%s───%s", startSymbol, findById(n.To, node).Label)

		_, _ = fmt.Fprintf(out, "\n")

		dfs(out, *findById(n.To, node), node, edge, depthLevel+1, append(mark, isLastFileInDir))

	}

}
