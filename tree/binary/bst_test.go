package binary

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BSTInsert implements an operation to insert node into a standard Binary Search Tree
func BSTInsert[N numericType](root *Node[N], value N) {
	if root.value > value {
		if root.left == nil {
			root.left = &Node[N]{
				value: value,
			}
		} else {
			BSTInsert(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &Node[N]{
				value: value,
			}
		} else {
			BSTInsert(root.right, value)
		}
	}
}

// Testable represents testable asset
type Testable interface {
	Test(t *testing.T)
}

// BSTInsertBalanceTest determine if the tree is balance
// For example:
//
//	  10
//	 /  \
//	5    11
type BSTInsertBalanceTest[N numericType] struct {
	F      func(*Node[N], N)
	Inputs []N
}

func (b BSTInsertBalanceTest[N]) Test(t *testing.T) {
	var root *Node[N]
	for i, input := range b.Inputs {
		if i == 0 {
			root = &Node[N]{
				value: input,
			}
		} else {
			b.F(root, input)
		}
	}

	assert.NotNil(t, root.left, "There should be something")
	assert.NotNil(t, root.right, "There should be something")
}

// BSTInsertSkewedTest determine if the tree is skewed
// For example:
//
//		  1
//		   \
//		    2
//	         \
//	          3
type BSTInsertSkewedTest[N numericType] struct {
	F      func(*Node[N], N)
	Inputs []N
}

func (b BSTInsertSkewedTest[N]) Test(t *testing.T) {
	var root *Node[N]
	for i, input := range b.Inputs {
		if i == 0 {
			root = &Node[N]{
				value: input,
			}
		} else {
			b.F(root, input)
		}
	}

	assert.NotNil(t, root.right, "There should be something")
	assert.NotNil(t, root.right, "There should be something")
}

func TestBSTInsert(t *testing.T) {
	tcs := []Testable{
		BSTInsertBalanceTest[int16]{
			F:      BSTInsert[int16],
			Inputs: []int16{10, 2, 11},
		},
		BSTInsertSkewedTest[int16]{
			F:      BSTInsert[int16],
			Inputs: []int16{1, 2, 3},
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Case %v", i), func(t *testing.T) {
			tc.Test(t)
		})

	}
}

var ErrSearch = errors.New("search error")

// BSTSearch performs search of Standard Binary Search Tree
func BSTSearch[N numericType](root *Node[N], value N) (*Node[N], error) {
	if root == nil {
		return nil, fmt.Errorf("%w", ErrSearch)
	}
	if value == root.value {
		return root, nil
	}
	if value < root.value {
		return BSTSearch(root.left, value)
	}
	return BSTSearch(root.right, value)
}

type BSTSearchNotFoundTest[N numericType] struct {
	f           func(*Node[N], N) (*Node[N], error)
	searchParam N
	inputs      []N
}

func (b BSTSearchNotFoundTest[N]) Test(t *testing.T) {
	var root *Node[N]
	for i, input := range b.inputs {
		if i == 0 {
			root = &Node[N]{
				value: input,
			}
		} else {
			BSTInsert(root, input)
		}
	}

	got, gotErr := BSTSearch(root, b.searchParam)
	if assert.ErrorIs(t, gotErr, ErrSearch, fmt.Sprintf("Want: %v Got: %v", ErrSearch, gotErr)) {
		assert.Empty(t, got, "Empty result expected")
	}
}

type BSTSearchFoundTest[N numericType] struct {
	f           func(*Node[N], N) (*Node[N], error)
	searchParam N
	inputs      []N
}

func (b BSTSearchFoundTest[N]) Test(t *testing.T) {
	var root *Node[N]
	for i, input := range b.inputs {
		if i == 0 {
			root = &Node[N]{
				value: input,
			}
		} else {
			BSTInsert(root, input)
		}
	}

	got, gotErr := BSTSearch(root, b.searchParam)
	if assert.ErrorIs(t, gotErr, nil, fmt.Sprintf("Want: %v Got: %v", nil, gotErr)) {
		assert.Equal(t, b.searchParam, got.value, fmt.Sprintf("Want: %v Got: %v", b.searchParam, got.value))
	}
}

func TestBSTSearch(t *testing.T) {
	testcases := []Testable{
		BSTSearchNotFoundTest[int16]{
			f:           BSTSearch[int16],
			searchParam: 5,
			inputs:      []int16{1, 2, 3, 7, 8},
		},
		BSTSearchFoundTest[int16]{
			f:           BSTSearch[int16],
			searchParam: 5,
			inputs:      []int16{1, 2, 3, 5, 7, 8},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tc.Test(t)
		})
	}
}

// BSTInorderTraversal collect result of traversal
func BSTInorderTraversal[N numericType](root *Node[N]) []N {
	if root == nil {
		return nil
	}

	leftResult := BSTInorderTraversal(root.left)
	rightResult := BSTInorderTraversal(root.right)

	result := append(leftResult, root.value)
	result = append(result, rightResult...) // Append all elements from rightResult

	return result
}

func Example_balanceBSTInOrderTraversal() {
	var root *Node[int16] = &Node[int16]{
		value: 100,
	}
	BSTInsert(root, 20)
	BSTInsert(root, 10)
	BSTInsert(root, 30)
	BSTInsert(root, 200)
	BSTInsert(root, 150)
	BSTInsert(root, 300)

	result := BSTInorderTraversal(root)
	fmt.Println(result)

	// Output:
	// [10 20 30 100 150 200 300]
}

func Example_skewedBSTInOrderTraversal() {
	var root *Node[int16] = &Node[int16]{
		value: 1,
	}
	BSTInsert(root, 2)
	BSTInsert(root, 3)
	BSTInsert(root, 4)
	BSTInsert(root, 5)
	BSTInsert(root, 6)
	BSTInsert(root, 7)

	result := BSTInorderTraversal(root)
	fmt.Println(result)

	// Output:
	// [1 2 3 4 5 6 7]
}

// BSTPreorderTraversal implements preorder traversal
func BSTPreorderTraversal[N numericType](root *Node[N]) []N {
	if root == nil {
		return []N{}
	}

	leftResult := BSTPreorderTraversal(root.left)
	rightResult := BSTPreorderTraversal(root.right)

	// Logic: Root -> Left -> Right
	result := []N{root.value}               // Add root value first
	result = append(result, leftResult...)  // Then left subtree
	result = append(result, rightResult...) // Then right subtree

	return result
}

func Example_bstPreorderTraversalBalance() {
	var root *Node[int16] = &Node[int16]{
		value: 100,
	}
	BSTInsert(root, 20)
	BSTInsert(root, 10)
	BSTInsert(root, 30)
	BSTInsert(root, 200)
	BSTInsert(root, 150)
	BSTInsert(root, 300)

	result := BSTPreorderTraversal(root)
	fmt.Println(result)

	// Output:
	// [100 20 10 30 200 150 300]
}

func Example_bstPreorderTraversalSkewed() {
	var root *Node[int16] = &Node[int16]{
		value: 1,
	}
	BSTInsert(root, 2)
	BSTInsert(root, 3)
	BSTInsert(root, 4)
	BSTInsert(root, 5)
	BSTInsert(root, 6)
	BSTInsert(root, 7)

	result := BSTPreorderTraversal(root)
	fmt.Println(result)

	// Output:
	// [1 2 3 4 5 6 7]
}

// BSTPostorderTraversal performs post order traversal
func BSTPostorderTraversal[N numericType](root *Node[N]) []N {
	if root == nil {
		return []N{}
	}

	leftResult := BSTPostorderTraversal(root.left)
	rightResult := BSTPostorderTraversal(root.right)

	result := append(leftResult, rightResult...) // Left subtree, then right subtree
	result = append(result, root.value)          // Finally, the root

	return result
}

func Example_bstPostorderTraversalBalance() {
	var root *Node[int16] = &Node[int16]{
		value: 100,
	}
	BSTInsert(root, 20)
	BSTInsert(root, 10)
	BSTInsert(root, 30)
	BSTInsert(root, 200)
	BSTInsert(root, 150)
	BSTInsert(root, 300)

	result := BSTPostorderTraversal(root)
	fmt.Println(result)

	// Output:
	// [10 30 20 150 300 200 100]
}

func Example_bstPostorderTraversalSkewed() {
	var root *Node[int16] = &Node[int16]{
		value: 1,
	}
	BSTInsert(root, 2)
	BSTInsert(root, 3)
	BSTInsert(root, 4)
	BSTInsert(root, 5)
	BSTInsert(root, 6)
	BSTInsert(root, 7)

	result := BSTPostorderTraversal(root)
	fmt.Println(result)

	// Output:
	// [7 6 5 4 3 2 1]
}
