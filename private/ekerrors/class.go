// Copyright (C) 2020  Germán Fuentes Capella

package ekerrors

var (
	registry = map[string]*Class{}
)

type Class struct {
	s string
}

func MustRegisterClass(s string) *Class {
	_, ok := registry[s]
	if ok {
		panic(s + " is already registered")
	}
	class := &Class{s: s}
	registry[s] = class
	return class
}

func MustGetClass(s string) *Class {
	class, ok := registry[s]
	if !ok {
		panic(s + " is not registered")
	}
	return class
}
