// Copyright (C) 2020  Germán Fuentes Capella

package eklark

import (
	"fmt"
	"strings"

	"go.starlark.net/starlark"
)

type MaskFn func(field string, value starlark.Value) string

func NoMaskFn(field string, value starlark.Value) string {
	return value.String()
}

type EKValue struct {
	valueType  string
	attrs      []string
	data       map[string]starlark.Value
	validators map[string]ValidateFn
	frozen     bool
	mask       MaskFn
}

func NewEKValue(
	valueType string,
	attrs []string,
	data map[string]starlark.Value,
	validators map[string]ValidateFn,
	mask MaskFn,
) EKValue {
	return EKValue{
		valueType:  valueType,
		attrs:      attrs,
		data:       data,
		validators: validators,
		frozen:     false,
		mask:       mask,
	}
}

// Implementing starlark Value interface

func (ev *EKValue) String() string {
	fields := make([]string, len(ev.data))
	if len(ev.attrs) > 0 {
		for i, field := range ev.attrs {
			fields[i] = fmt.Sprintf("%s=%v", field, ev.mask(field, ev.data[field]))
		}
	} else {
		i := 0
		for field, value := range ev.data {
			fields[i] = fmt.Sprintf("%s=%v", field, ev.mask(field, value))
			i++
		}
	}
	return fmt.Sprintf("%s{%s}", ev.valueType, strings.Join(fields, ", "))
}

func (ev *EKValue) Type() string {
	return ev.valueType
}

func (ev *EKValue) Freeze() {
	ev.frozen = true
	for _, value := range ev.data {
		value.Freeze()
	}
}

func (ev *EKValue) Truth() starlark.Bool {
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

func (ev *EKValue) Hash() (uint32, error) {
	return 0, fmt.Errorf("Unhashable type: %s", ev.valueType)
}

// Implements starlark HasAttrs interface

func (ev *EKValue) isValidAttr(name string) bool {
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

func (ev *EKValue) Attr(name string) (starlark.Value, error) {
	if !ev.isValidAttr(name) {
		return nil, starlark.NoSuchAttrError(
			fmt.Sprintf("type object '%s' has no attribute '%s'", ev.valueType, name),
		)
	}
	value, ok := ev.data[name]
	if !ok {
		// value not present
		return nil, nil
	}
	return value, nil
}

func (ev *EKValue) AttrNames() []string {
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

func (ev *EKValue) SetField(name string, val starlark.Value) error {
	if ev.frozen {
		return fmt.Errorf("%s is frozen", ev.valueType)
	}
	if !ev.isValidAttr(name) {
		return starlark.NoSuchAttrError(
			fmt.Sprintf("type object '%s' has no attribute '%s'", ev.valueType, name),
		)
	}
	validFn, ok := ev.validators[name]
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

func (ev *EKValue) Get(k starlark.Value) (v starlark.Value, found bool, err error) {
	ks, ok := starlark.AsString(k)
	if !ok {
		return nil, false, fmt.Errorf("%s only accepts string keys; found: '%v'", ev.valueType, k)
	}
	v, found = ev.data[ks]
	if !found {
		return nil, false, nil
	}
	return v, true, nil
}

// Implementing starlark HasSetKey interface

func (ev *EKValue) SetKey(k, v starlark.Value) error {
	if ev.frozen {
		return fmt.Errorf("%s is frozen", ev.valueType)
	}
	ks, ok := starlark.AsString(k)
	if !ok {
		return fmt.Errorf("%s only accepts string keys; found: '%v'", ev.valueType, k)
	}
	ev.data[ks] = v
	return nil
}
