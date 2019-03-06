package src

import (
	"fmt"
	"testing"
)

var list *SkipList
const size = 100000

func init() {
	//list = New()
	//
	//for i := 0;i<=size;i++ {
	//	list.Set(float64(i),fmt.Sprintf("%v",i))
	//}
}

func check(s *SkipList,t *testing.T) {
	for k,v := range s.next {
		if v == nil {
			continue
		}

		if k > len(v.next) {
			t.Fatal("first node's level must be no less than current level")
		}

		next := v
		cnt := 1

		for next.next[k] != nil {
			if !(next.next[k].key >= next.key) {
				t.Fatalf("next key value must be greater than prev key value. [next:%v] [prev:%v]", next.next[k].key, next.key)
			}

			if k > len(next.next) {
				t.Fatalf("node's level must be no less than current level. [cur:%v] [node:%v]", k, next.next)
			}

			next = next.next[k]
			cnt++
		}
		fmt.Println()

		if k == 0 {
			if cnt != s.length {
				t.Log("CNT:",cnt,"Length:",s.length)
				t.Fatalf("list len must match the level 0 nodes count. [cur:%v] [level0:%v]", cnt, s.length)
			}
		}
	}
}

func TestNew(t *testing.T) {
	var s *SkipList

	s = New()

	s.Set(10, 1)
	s.Set(60, 2)
	s.Set(30, 3)
	s.Set(20, 4)
	s.Set(90, 5)
	t.Log("inserted")
	s.Show()

	s.Set(30, 9)
	t.Log("inserted duplicates")
	check(s, t)

	s.Remove(0)
	s.Remove(20)
	t.Log("removed")
	s.Show()
	check(s, t)

	v1 := s.Get(10)
	v2 := s.Get(60)
	v3 := s.Get(30)
	v4 := s.Get(20)
	v5 := s.Get(90)
	v6 := s.Get(0)

	if v1 == nil || v1.Value.(int) != 1 || v1.key != 10 {
		t.Fatal(`wrong "10" value (expected "1")`, v1)
	}

	if v2 == nil || v2.Value.(int) != 2 {
		t.Fatal(`wrong "60" value (expected "2")`)
	}

	if v3 == nil || v3.Value.(int) != 9 {
		t.Fatal(`wrong "30" value (expected "9")`)
	}

	if v4 != nil {
		t.Fatal(`found value for key "20", which should have been deleted`)
	}

	if v5 == nil || v5.Value.(int) != 5 {
		t.Fatal(`wrong "90" value`)
	}

	if v6 != nil {
		t.Fatal(`found value for key "0", which should have been deleted`)
	}
}