package tbsp

import (
	"fmt"
	"testing"
)

func TestTraverse(t *testing.T) {
	b := New(0, 0, 20, 20)
	b.Split(true, 10)
	b.GetL().Split(false, 5)
	b.GetR().Split(false, 5)

	callback := func(node *Bsp, data *interface{}) error {
		fmt.Printf("Entering level %d\n", node.Level())
		return nil
	}

	if err := b.Traverse(callback, nil); err != nil {
		t.Error(err)
	}
}
