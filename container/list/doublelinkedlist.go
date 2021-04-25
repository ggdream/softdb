package list

import (
	"container/list"
	"errors"
	"reflect"
)


type (
	DoubleLinkedList struct {
		data *list.List
	}
)

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		data: list.New(),
	}
}


func (d *DoubleLinkedList) LPush(values ...[]byte) int {

	return d.push(true, values...)
}

func (d *DoubleLinkedList) RPush(values ...[]byte) int {
	
	return d.push(false, values...)
}

func (d *DoubleLinkedList) LPop() ([]byte, error) {
	return d.pop(true)
}

func (d *DoubleLinkedList) RPop() ([]byte, error) {
	return d.pop(false)
}

func (d *DoubleLinkedList) LIndex(idx int) ([]byte, error) {
	newIdx, ok := d.jugdeIdx(idx)
	if !ok {
		return nil, errors.New("idx overflow")
	}
	return d.getValByIdx(newIdx).Value.([]byte), nil
}

func (d *DoubleLinkedList) LRem(value []byte, count int) int {
	var els []*list.Element

	if count == 0 {
		for i := d.data.Front(); i != nil; i = i.Next() {
			if reflect.DeepEqual(i.Value.([]byte), value) {
				els = append(els, i)
			}
		}
	} else if count > 0 {
		for i := d.data.Front(); i != nil && len(els) < count; i =  i.Next() {
			if reflect.DeepEqual(i.Value.([]byte), value) {
				els = append(els, i)
			}
		}
	} else {
		for i := d.data.Back(); i != nil && len(els) < -count; i = i.Prev() {
			if reflect.DeepEqual(i.Value.([]byte), value) {
				els = append(els, i)
			}
		}
	}

	for _, v := range els {
		d.data.Remove(v)
	}
	return len(els)
}

func (d *DoubleLinkedList) LInsert(insertType InsertType, pivot, value []byte) int {
	el := d.find(pivot)
	if el == nil {
		return -1
	}

	switch insertType {
	case Before:
		d.data.InsertBefore(value, el)
	case After:
		d.data.InsertAfter(value, el)
	}

	return d.data.Len()
}

func (d *DoubleLinkedList) LSet(idx int, value []byte) error {
	newIdx, ok := d.jugdeIdx(idx)
	if !ok {
		return errors.New("index error")
	}
	node := d.getValByIdx(newIdx)
	node.Value = value
	return nil
}

func (d *DoubleLinkedList) LRange(sIdx, eIdx int) ([][]byte, error) {
	newSIdx, ok := d.jugdeIdx(sIdx)
	if !ok {
		return nil, errors.New("sInx error")
	}
	newEIdx, ok := d.jugdeIdx(eIdx)
	if !ok {
		return nil, errors.New("eInx error")
	}
	if newSIdx > newEIdx {
		return nil, errors.New("err: sIdx > eIdx")
	}

	var idx int
	var values [][]byte
	for i := d.data.Front(); idx <= newEIdx; i = i.Next() {
		if idx >= newSIdx {
			values = append(values, i.Value.([]byte))
		}
		idx++
	}

	return values, nil
}

func (d *DoubleLinkedList) LLen() int {
	return d.data.Len()
}

func (d *DoubleLinkedList) push(isLeft bool, values ...[]byte) int {
	for _, v := range values {
		if isLeft {
			d.data.PushFront(v)
		} else {
			d.data.PushBack(v)
		}
	}
	return d.data.Len()
}

func (d *DoubleLinkedList) pop(isLeft bool) ([]byte, error) {
	if d.data.Len() <= 0 {
		return nil, errors.New("the list is empty")
	}

	var el *list.Element
	if isLeft {
		el = d.data.Front()
	} else {
		el = d.data.Back()
	}

	return d.data.Remove(el).([]byte), nil
}


func (d *DoubleLinkedList) jugdeIdx(idx int) (int, bool) {
	if idx >= d.data.Len() || idx <= -d.data.Len() {
		return 0, false
	}
	if idx < 0 {
		idx = d.data.Len() + idx
	}
	return idx, true
}

func (d *DoubleLinkedList) getValByIdx(idx int) *list.Element {
	el := d.data.Front()
	for i := 0; i < idx-1; i++ {
		el = el.Next()
	}
	return el
}

func (d *DoubleLinkedList) find(value []byte) *list.Element {
	var el *list.Element
	for i := d.data.Front(); i != nil; i = i.Next() {
		if reflect.DeepEqual(i.Value.([]byte), value) {
			el = i
			break
		}
	}
	return el
}
