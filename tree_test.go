package iavl

import (
	"fmt"
	"math"
	"testing"
	"time"

	"math/rand"
)

func calcIterate(t *Tree, n uint) {
	sumx, sumxx := float64(0), float64(0)

	var iterate func(root *Node, sum float64)
	iterate = func(root *Node, sum float64) {
		x := float64(0)
		if root.parent != nil {
			x = math.Abs(float64(root.parent.path) - float64(root.path))
		}
		sum += x
		if root.height == 0 {
			sumx += sum
			sumxx += sum * sum
		} else {
			iterate(root.left, sum)
			iterate(root.right, sum)
		}
		sum -= x
	}
	iterate(t.root, 0)
	avg := sumx / float64(n)
	std := math.Sqrt(sumxx/float64(n) - avg*avg)
	fmt.Printf("Avg %.3f, Std: %.3f \n", avg, std)
}

func TestAcsInsert(t *testing.T) {
	tree := NewTree()
	n := uint(1000000)
	for i := uint(1); i < n; i++ {
		tree.AddNode(i)
	}
	fmt.Print("Sequenced Insert: ")
	calcIterate(tree, n)
}

func TestDcsInsert(t *testing.T) {
	tree := NewTree()
	n := uint(1000000)
	for i := n; i > 0; i-- {
		tree.AddNode(i)
	}
	fmt.Print("Sequenced Insert: ")
	calcIterate(tree, n)
}

func TestMixInsert(t *testing.T) {
	tree := NewTree()
	n := uint(1000000)
	for i := uint(1); i < n; i += 2 {
		tree.AddNode(i)
	}
	for i := uint(0); i < n; i += 2 {
		tree.AddNode(i)
	}
	fmt.Print("Sequenced Insert: ")
	calcIterate(tree, n)
}

func TestRandomInsert(t *testing.T) {
	tree := NewTree()
	n := uint(1000000)
	rand.Seed(time.Now().UnixNano())
	for i := uint(1); i < n; i++ {
		tree.AddNode(uint(rand.Uint32()))
	}
	fmt.Print("Random Insert: ")
	calcIterate(tree, n)
}
