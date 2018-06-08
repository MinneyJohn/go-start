package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) int{
	if nil == t {
		return 0
	}
	
	//fmt.Println("In Walk")
	if nil != t.Left {
		//fmt.Println("Go to Left")
		Walk(t.Left, ch) 
	}
	
	//fmt.Println("Show Value", t.Value)
	
	ch <- t.Value
	//fmt.Println(t.Value)

	
	if nil != t.Right {
		//fmt.Println("Go to Right")
		Walk(t.Right, ch)
	}
	return 0
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch_1 := make(chan int)
	ch_2 := make(chan int)
	
	go Walk(t1, ch_1)
	go Walk(t2, ch_2)
	
	for counter := 0; counter < 10; counter++{
		value_1 := <-ch_1
		value_2 := <-ch_2
		fmt.Println("Verify ", value_1, " ", value_2)
		if value_1 != value_2 {
			return false
		}
	}
	
	return true
}

func main() {
	tree_1 := tree.New(1)
	tree_2 := tree.New(2)
	fmt.Println(Same(tree_1, tree_2))
}

