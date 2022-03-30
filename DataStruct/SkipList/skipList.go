package main

import (
	"math/rand"
	"time"
)

const (
	MaxLevel = 32
	p        = 0.5
)

type Element uint32

type Node struct {
	value  Element
	levels []*Level
}

type Level struct {
	next *Node
}

type SkipList struct {
	header *Node
	length uint32
	height uint32
}

func NewSkipList() *SkipList {
	return &SkipList{
		header: NewNode(MaxLevel, 0),
		length: 0,
		height: 1,
	}
}

func NewNode(level uint32, value Element) *Node {
	node := new(Node)
	node.value = value
	node.levels = make([]*Level, level)
	for i := 0; i < len(node.levels); i++ {
		node.levels[i] = new(Level)
	}
	return node
}

func (sl *SkipList) randomLevela() int {
	level := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for r.Float64() < p && level < MaxLevel {
		level++
	}
	return level
}

func (sl *SkipList) Add(value Element) bool {
	if value <= 0 {
		return false
	}

	update := make([]*Node, MaxLevel)
	tmp := sl.header

	for i := int(sl.height) - 1; i >= 0; i-- {
		for tmp.levels[i].next != nil && tmp.levels[i].next.value < value {
			tmp = tmp.levels[i].next
		}

		if tmp.levels[i].next != nil && tmp.levels[i].next.value == value {
			return false
		}

		update[i] = tmp
	}

	level := sl.randomLevela()
	node := NewNode(uint32(level), value)

}

func main() {

}
