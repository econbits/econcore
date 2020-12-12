//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"

	"go.starlark.net/starlark"
)

type Script struct {
	tn      *starlark.Thread
	fpath   string
	globals starlark.StringDict
}

func (s Script) stringField(name string, defvalue string) string {
	v, ok := s.globals[name]
	if !ok {
		return defvalue
	}
	sv, ok := starlark.AsString(v)
	if !ok {
		return defvalue
	}
	return sv
}

func (s Script) stringListField(name string, defvalue []string) []string {
	v, ok := s.globals[name]
	if !ok {
		return defvalue
	}
	vlist, ok := v.(*starlark.List)
	if !ok {
		return defvalue
	}
	lv := vlist.Len()
	fields := make([]string, lv)
	for i := 0; i < lv; i++ {
		s, ok := starlark.AsString(vlist.Index(i))
		if ok {
			fields[i] = s
		}
	}
	return fields
}

func (s Script) Description() string {
	return s.stringField(hDescription, defDescription)
}

func (s Script) URL() string {
	return s.stringField(hURL, defUrl)
}

func (s Script) License() string {
	return s.stringField(hLicense, defLicense)
}

func (s Script) Authors() []string {
	return s.stringListField(hAuthors, defAuthors)
}

func (s Script) Login(cred Credentials) (Session, error) {
	login := s.globals["login"]
	if login == nil {
		return Session{}, fmt.Errorf("[%s][login] missing login function", fname(s.fpath))
	}

	vsession, err := starlark.Call(s.tn, login, starlark.Tuple{cred}, nil)
	if err != nil {
		return Session{}, fmt.Errorf("[%s][login] %v", fname(s.fpath), err)
	}
	session, ok := vsession.(Session)
	if !ok {
		return Session{}, fmt.Errorf(
			"[%s][login] login function returned object of type '%T' instead of a session",
			fname(s.fpath),
			vsession,
		)
	}
	return session, nil
}
