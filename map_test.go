package sugar

import (
	"log"
	"testing"
)

func TestMapAr(t *testing.T) {
	var x = []int{1, 2, 3, 4, 5, 6}
	var o = make([]int, len(x))
	MapAr(f, x, o)

	t.Logf("Value of v is %+v", o)
	ForEach(x, f)
}

var (
	f  = func(v int) int { return v * v }
	f1 = func(v int) int { return v / 2 }
	f2 = func(v int) int { return v * 2 }
)

func TestChain(t *testing.T) {
	var x = []int{1, 2, 3, 4, 5, 6}

	var o = make([]bool, len(x))
	Chain(x, o, f, f1, f2)
	log.Printf("%+v", o)
}

func TestLazy(t *testing.T) {
	var ar []int
	for i := 0; i < 1000; i++ {
		ar = append(ar, i)

	}

	l := MakeLazy(ar, f, f1, f2)
	sl := make([]int, 20)
	l.Take(20, sl)
	t.Logf("%+v", sl)
	d := make([]int, 300)
	l.Slice(502, 802, d)
	t.Logf("%+v", d)
	l.Slice(602, 702, d)
	t.Logf("%+v", d)

}
