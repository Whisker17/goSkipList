package src

import (
	"math/rand"
	"sync"
)

type Element struct {
	elementNode
	key float64
	Value interface{}
}

type elementNode struct {
	next []*Element
}

type SkipList struct {
	elementNode
	level          int
	length         int
	randSource     rand.Source
	probability    float64
	probTable      []float64
	mutex          sync.RWMutex
	prevNodesCache []*elementNode
}
