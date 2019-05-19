//@author Devansh Gupta
//This contains code for lazy function evalution, inspired by haskell's lazy evaluation
//And Elixir's streams

package sugar

import "reflect"

//Lazy, holds lazy evaluation values for faster usage
type Lazy struct {
	fns  []reflect.Value
	iar  reflect.Value
	omap map[int]reflect.Value
}

//MakeLazy, it returns a new Lazy Object for given array or function
func MakeLazy(iar AnyValue, fns ...MapFunction) *Lazy {
	x := new(Lazy)
	x.omap = make(map[int]reflect.Value)
	x.iar = reflect.ValueOf(iar)
	for i := 0; i < len(fns); i++ {
		x.fns = append(x.fns, reflect.ValueOf(fns[i]))
	}
	return x
}

//Take, evaluates first n elements
//n=> no. of elements to be taken
//ar=> slice in which, output will persist
func (x *Lazy) Take(n int, ar AnyValue) {
	oar := reflect.ValueOf(ar)
	for i := 0; i < n; i++ {
		if v, ok := x.omap[i]; ok {
			oar.Index(i).Set(v)
			continue
		}
		var v = []reflect.Value{x.iar.Index(i)}
		for j := 0; j < len(x.fns); j++ {
			v = x.fns[j].Call(v)
		}
		oar.Index(i).Set(v[0])
		x.omap[i] = oar.Index(i)
	}
}

//Slice, evaluates the slice in p<=index<q
//p=> lower index
//q=> upper index
//ar=> output slice
func (x *Lazy) Slice(p, q int, ar AnyValue) {

	oar := reflect.ValueOf(ar)
	var ii = 0
	k := oar.Len()
	for i := p; i < q && k >= 0; i++ {
		var v reflect.Value
		if _, ok := x.omap[i]; ok {
			v = x.omap[i]
		} else {
			var vv = []reflect.Value{x.iar.Index(i)}
			for j := 0; j < len(x.fns); j++ {
				vv = x.fns[j].Call(vv)
			}
			v = vv[0]
			x.omap[p] = v //setting value in v
		}
		oar.Index(ii).Set(v)
		ii++
		k--
	}
}
