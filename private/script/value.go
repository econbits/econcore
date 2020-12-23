//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"
	"strings"

	"go.starlark.net/starlark"
)

type ekmValueType string

type maskFn func(field string, value starlark.Value) string

func noMask(field string, value starlark.Value) string {
	return value.String()
}

type EKMValue struct {
	valueType    ekmValueType
	attrs        []string
	data         map[string]starlark.Value
	validatorsFn map[string]validatorFunc
	frozen       bool
	mask         maskFn
}

// Implementing starlark Value interface

func (ev *EKMValue) String() string {
	fields := make([]string, len(ev.data))
	i := 0
	for field, value := range ev.data {
		fields[i] = fmt.Sprintf("%s=%v", field, ev.mask(field, value))
		i++
	}
	return fmt.Sprintf("%v{%s}", ev.valueType, strings.Join(fields, ", "))
}

func (ev *EKMValue) Type() string {
	return string(ev.valueType)
}

func (ev *EKMValue) Freeze() {
	ev.frozen = true
	for _, value := range ev.data {
		value.Freeze()
	}
}

func (ev *EKMValue) Truth() starlark.Bool {
	if len(ev.data) == 0 {
		return starlark.False
	}
	for _, v := range ev.data {
		if v.Truth() {
			return starlark.True
		}
	}
	return starlark.False
}

func (ev *EKMValue) Hash() (uint32, error) {
	return 0, fmt.Errorf("Unhashable type: %v", ev.valueType)
}

// Implements starlark HasAttrs interface

func (ev *EKMValue) isValidAttr(name string) bool {
	if len(ev.attrs) == 0 {
		return true
	}
	for _, attr := range ev.attrs {
		if name == attr {
			return true
		}
	}
	return false
}

func (ev *EKMValue) Attr(name string) (starlark.Value, error) {
	if !ev.isValidAttr(name) {
		return nil, fmt.Errorf("Attribute '%s' is not valid. Options: %v", name, ev.attrs)
	}
	value, ok := ev.data[name]
	if ok {
		return value, nil
	}
	return nil, nil
}

func (ev *EKMValue) AttrNames() []string {
	if len(ev.attrs) == 0 {
		attrs := make([]string, len(ev.data))
		i := 0
		for k, _ := range ev.data {
			attrs[i] = k
			i++
		}
		return attrs
	}
	return ev.attrs
}

// Implements starlark HasSetField interface

func (ev *EKMValue) SetField(name string, val starlark.Value) error {
	if ev.frozen {
		return fmt.Errorf("%s is frozen", ev.valueType)
	}
	if !ev.isValidAttr(name) {
		return fmt.Errorf("Attribute '%s' is not valid. Options: %v", name, ev.attrs)
	}
	validFn, ok := ev.validatorsFn[name]
	if ok {
		err := validFn(val)
		if err != nil {
			return err
		}
	}
	ev.data[name] = val
	return nil
}

// Implementing starlark Mapping interface

func (ev *EKMValue) Get(k starlark.Value) (v starlark.Value, found bool, err error) {
	ks, ok := starlark.AsString(k)
	if !ok {
		return nil, false, fmt.Errorf("%v only accepts string keys; found: '%v'", ev.valueType, k)
	}
	v, found = ev.data[ks]
	if !found {
		return nil, false, nil
	}
	return v, true, nil
}

// Implementing starlark HasSetKey interface

func (ev *EKMValue) SetKey(k, v starlark.Value) error {
	if ev.frozen {
		return fmt.Errorf("%s is frozen", ev.valueType)
	}
	ks, ok := starlark.AsString(k)
	if !ok {
		return fmt.Errorf("%v only accepts string keys; found: '%v'", ev.valueType, k)
	}
	ev.data[ks] = v
	return nil
}
