package src

import (
	"math"
	"math/rand"
	"time"
	"fmt"
)

const (
	DefaultMaxLevel = 18
	DefaultProbability = 1 / math.E
)

//SetProbability changes the current P value of the list
func (s *SkipList) SetProbability(probability float64) {
	s.probability = probability
	s.probTable = probabilityTable(probability,s.level)
}

//probabilityTable calculates the probability of a new node with a given level
func probabilityTable(probability float64,maxLevel int) []float64 {
	var table []float64
	for i := 1;i<=maxLevel;i++ {
		prob := math.Pow(probability,float64(i-1))
		table = append(table,prob)
	}
	return table
}

//create a new skiplist with maxLevel
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

//create a new skiplist with defaultMaxLevel
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

func (s *SkipList) Set(key float64,value interface{}) *Element {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	prevs := s.getPrevElementNodes(key)
	var element *Element
	element = prevs[0].next[0]
	fmt.Println("len:",len(prevs))

	if element != nil && element.key <= key {
		element.Value = value
		return element
	}

	element = &Element{
		elementNode:elementNode{
			next:make([]*Element,s.randLevel()),
		},
		key:key,
		Value:value,
	}

	for i := range element.next {
		element.next[i] = prevs[i].next[i]
		prevs[i].next[i] = element
	}

	s.length++
	return element
}

func (s *SkipList) Get(key float64) *Element {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var prev = &s.elementNode
	var next *Element

	for i := s.level-1;i>=0;i-- {
		next = prev.next[i]

		for next != nil && key > next.key {
			prev = &next.elementNode
			next = next.next[i]
		}
	}

	if next != nil && next.key <= key {
		return next
	}

	return nil
}

func (s *SkipList) Remove(key float64) *Element {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	prevs := s.getPrevElementNodes(key)
	element := prevs[0].next[0]

	if element != nil && element.key <= key {
		for k,v := range element.next {
			prevs[k].next[k] = v
		}

		s.length--
		return element
	}

	return nil
}

func (s *SkipList) Show() {
	var prev = &s.elementNode
	var next *Element
	for i := s.level -1;i>=0;i-- {
		next = prev.next[i]
		fmt.Printf("Level %v:",i)
		for next != nil {
			fmt.Printf("%v----->",next.key)
			prev = &next.elementNode
			next = next.next[i]
		}
		fmt.Println()
	}
}

func (s *SkipList) getPrevElementNodes(key float64) []*elementNode {
	var prevNode = &s.elementNode
	var next *Element

	prevCache := s.prevNodesCache

	for i := s.level-1;i>=0;i-- {
		next = prevNode.next[i]

		for next != nil && key >next.key {
			prevNode = &next.elementNode
			next = next.next[i]
		}

		prevCache[i] = prevNode
	}

	return prevCache
}

func (s *SkipList) randLevel() int {
	l := 1

	for float64((s.randSource.Int63() >> 32) & 0xFFFF) < s.probTable[l] {
		l++
	}

	if l > s.level {
		l = s.level
	}

	fmt.Println("RandNum:",l)

	return l
}