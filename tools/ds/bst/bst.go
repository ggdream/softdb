package bst

import "errors"


type (
	Node struct {
		lNode *Node
		rNode *Node
		index int
		value interface{}
	}

	Tree struct {
		node   *Node
		length int
	}
)

func New(index int, value interface{}) *Tree {
	return &Tree{
		node: &Node{
			index: index,
			value: value,
		},
		length: 1,
	}
}

// Insert 插入
func (t *Tree) Insert(index int, value interface{}) {
	parPtr :=  t.node
	var isLeft bool
	for i := t.node; i != nil; {
		parPtr = i
		if i.index > index {
			i = i.lNode
			isLeft = true
		} else if i.index < index {
			i = i.rNode
			isLeft = false
		}
	}

	node := &Node{
		lNode: nil,
		rNode: nil,
		index: index,
		value: value,
	}
	if isLeft {
		parPtr.lNode = node
	} else {
		parPtr.rNode = node
	}
}

// Delete 删除
func (t *Tree) Delete(index int) error {
	parPtr := t.node
	var isLeft bool
	for i := t.node; i != nil; {
		if i.index > index {
			i = i.lNode
			isLeft = true
		} else if i.index < index {
			i = i.rNode
			isLeft = false
		}

		if i.index == index {
			if isLeft {
				parPtr.lNode = nil
			} else {
				parPtr.rNode = nil

			}
			return nil
		}
		parPtr = i
	}

	return errors.New("no the index")
}

// Update 更新
func (t *Tree) Update(index int, value interface{}) error {
	for i := t.node; i != nil; {
		if i.index > index {
			i = i.lNode
		} else if i.index < index {
			i = i.rNode
		}

		if i.index == index {
			i.value = value
			return nil
		}
	}

	return errors.New("no the index")
}

// Search 查找
func (t *Tree) Search(index int) (interface{}, error) {
	for i := t.node; i != nil; {
		if i.index > index {
			i = i.lNode
		} else if i.index < index {
			i = i.rNode
		}

		if i.index == index {
			return i.value, nil
		}
	}

	return nil, errors.New("no the index")
}
