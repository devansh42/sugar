//@author Devansh Gupta
//This file cotains code some map utility functions

package sugar

import (
	"reflect"
)

//Map, is just a synctactic sugar over map[string]interface{}
type Map map[string]interface{}

//ContainsKey, returns true if map contains a particular key
func ContainsKey(key interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[key]
	return ok
}

type MapFunction interface{}
type AnyValue interface{}

func getmkf(f MapFunction) reflect.Value {
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	return reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		return fv.Call(in)
	})

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//MapAr is like pythons map function
//f is function which takes a single argument a returns a single value
//iar, is the array to iterate over
//oar, is the output array which contains the result of call to function f of each element in iar
//no. of iteration will be equal to min(len(iar),len(oar))
//This internally use go runtime reflection
func MapAr(f MapFunction, iar, oar AnyValue) {

	viar := reflect.ValueOf(iar)
	voar := reflect.ValueOf(oar)
	l := min(viar.Len(), voar.Len())
	mkf := getmkf(f)
	for x := 0; x < l; x++ {
		ov := mkf.Call([]reflect.Value{viar.Index(x)})[0]
		voar.Index(x).Set(ov)
		//		voar = reflect.Append(voar, ov)
	}

}

//ForEach, is like javascripts foreach method
//iar=> input slice
//fn=> Func to perform
//This internally use go runtime reflection
func ForEach(iar AnyValue, fn MapFunction) {
	v := reflect.ValueOf(iar)
	mkf := getmkf(fn)
	for i := 0; i < v.Len(); i++ {
		mkf.Call([]reflect.Value{v.Index(i)}) //Calling the function
	}
}

//Chain, chain methods and after calling all the functions save output in oar
//iar => input slice
//oar => Output slice
//fns => Functions to Chain Over
//This internally use go runtime reflection
func Chain(iar, oar AnyValue, fns ...MapFunction) {
	var arfn []reflect.Value
	for i := 0; i < len(fns); i++ {
		arfn = append(arfn, getmkf(fns[i]))
	}
	voar := reflect.ValueOf(oar)
	viar := reflect.ValueOf(iar)
	l := min(voar.Len(), viar.Len())
	for i := 0; i < l; i++ {
		v := []reflect.Value{viar.Index(i)}
		for j := 0; j < len(arfn); j++ {
			v = arfn[j].Call(v)
		}
		voar.Index(i).Set(v[0]) //Setting value of output array
	}
}
