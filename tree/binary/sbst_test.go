package binary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SBSTInsert implements an operation to insert node into a standard Binary Search Tree
func SBSTInsert[N numericType](root *Node[N], value N) {
	if root.value > value {
		if root.left == nil {
			root.left = &Node[N]{
				value: value,
			}
		} else {
			SBSTInsert(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &Node[N]{
				value: value,
			}
		} else {
			SBSTInsert(root.right, value)
		}
	}
}

type Testable interface {
	Test(t *testing.T)
}

// SBSTInsertBalanceTest determine if the tree is balance
// For example:
//
//	  10
//	 /  \
//	5    11
type SBSTInsertBalanceTest[N numericType] struct {
	F      func(*Node[N], N)
	Inputs []N
}

func (s SBSTInsertBalanceTest[N]) Test(t *testing.T) {
	var root *Node[N]
	for i, input := range s.Inputs {
		if i == 0 {
			root = &Node[N]{
				value: input,
			}
		} else {
			s.F(root, input)
		}
	}

	assert.NotNil(t, root.left, "There should be something")
	assert.NotNil(t, root.right, "There should be something")
}

// SBSTInsertSkewedTest determine if the tree is skewed
// For example:
//
//		  1
//		   \
//		    2
//	         \
//	          3
type SBSTInsertSkewedTest[N numericType] struct {
	F      func(*Node[N], N)
	Inputs []N
}

func (s SBSTInsertSkewedTest[N]) Test(t *testing.T) {
	var root *Node[N]
	for i, input := range s.Inputs {
		if i == 0 {
			root = &Node[N]{
				value: input,
			}
		} else {
			s.F(root, input)
		}
	}

	assert.NotNil(t, root.right, "There should be something")
	assert.NotNil(t, root.right, "There should be something")
}

func TestSBSTInsert(t *testing.T) {
	tcs := []Testable{
		SBSTInsertBalanceTest[int16]{
			F:      SBSTInsert[int16],
			Inputs: []int16{10, 2, 11},
		},
		SBSTInsertSkewedTest[int16]{
			F:      SBSTInsert[int16],
			Inputs: []int16{1, 2, 3},
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Case %v", i), func(t *testing.T) {
			tc.Test(t)
		})

	}
}
