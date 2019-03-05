package src

import (
	"math"
	"math/rand"
	"time"
)

const (
	DefaultMaxLevel = 18
	DefaultProbability = 1 / math.E
)

func probabilityTable(probability float64,maxLevel int) []float64 {
	var table []float64
	for i := 1;i<=maxLevel;i++ {
		prob := math.Pow(probability,float64(i-1))
		table = append(table,prob)
	}
	return table
}

func NewWithMaxLevel(maxLevel int) *SkipList {
	if maxLevel < 1 || maxLevel >64 {
		panic("maxLevel is invalid!Should be <= 64!")
	}
	return &SkipList{
		elementNode:elementNode{next:make([]*Element,maxLevel)},
		level:maxLevel,
		randSource:rand.New(rand.NewSource(time.Now().UnixNano())),
		probability:DefaultProbability,
		probTable:probabilityTable(DefaultProbability, maxLevel),
		prevNodesCache:make([]*elementNode,maxLevel),
	}
}

func New() *SkipList {
	return NewWithMaxLevel(DefaultMaxLevel)
}

// Gets the first element.
func (s *SkipList) Front() *Element {
	return s.next[0]
}

// Gets list length.
func (s *SkipList) Len() int {
	return s.length
}

