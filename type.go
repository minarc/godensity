package godensity

import "github.com/PuerkitoBio/goquery"

type Result struct {
	root *Node
	heap *Heap
}

type Node struct {
	goqueryNode *goquery.Selection
	density     float32
	densitySum  float32
	images      []string
	videos      []string
	T           float32
	text        string
	next        *Node
}

type Heap struct {
	nodes []Node
}
