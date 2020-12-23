//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"
	"time"

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

func (s Script) Login(cred *Credentials) (*Session, error) {
	nameFn := "login"
	value, err := s.runFn(nameFn, starlark.Tuple{cred})
	if err != nil {
		return nil, err
	}
	session, ok := value.(*Session)
	if !ok {
		return nil, newLoginError(
			fname(s.fpath),
			SessionError,
			fmt.Sprintf("login function returned object of type '%T' instead of a session", value),
		)
	}
	return session, nil
}

func (s Script) Accounts(session *Session) ([]*Account, error) {
	nameFn := "accounts"
	value, err := s.runFn(nameFn, starlark.Tuple{session})
	if err != nil {
		return []*Account{}, err
	}
	accountList, ok := value.(*starlark.List)
	if !ok {
		return []*Account{}, newAccountError(
			fname(s.fpath),
			AccountListError,
			fmt.Sprintf("account function returned object of type '%T' instead of a list of accounts", value),
		)
	}
	accounts, err := LtoAR(accountList)
	if err != nil {
		return []*Account{}, newAccountError(
			fname(s.fpath),
			AccountListError,
			err.Error(),
		)
	}
	return accounts, nil
}

func (s Script) runFn(nameFn string, params starlark.Tuple) (starlark.Value, error) {
	fn := s.globals[nameFn]
	if fn == nil {
		return nil, ScriptError{
			scriptName: fname(s.fpath),
			function:   nameFn,
			errorType:  MissingFunctionError,
			text:       fmt.Sprintf("missing %s function", nameFn),
		}
	}

	value, err := starlark.Call(s.tn, fn, params, nil)
	if err != nil {
		evalErr, ok := err.(*starlark.EvalError)
		if ok {
			err = evalErr.Unwrap()
			scErr, ok := err.(ScriptError)
			if ok {
				return nil, scErr
			}
		}
		return nil, ScriptError{
			scriptName: fname(s.fpath),
			function:   nameFn,
			errorType:  FunctionCallError,
			text:       err.Error(),
		}
	}
	return value, nil
}

func (s Script) Transactions(session *Session, account *Account, since time.Time) ([]*Transaction, error) {
	nameFn := "transactions"
	value, err := s.runFn(nameFn, starlark.Tuple{session, account, starlark.String("TODO")})
	if err != nil {
		return nil, err
	}
	txList, ok := value.(*starlark.List)
	if !ok {
		return nil, ScriptError{
			scriptName: fname(s.fpath),
			function:   nameFn,
			errorType:  TransactionListError,
			text:       fmt.Sprintf("account function returned object of type '%T' instead of a list of accounts", value),
		}
	}
	txs, err := LtoTR(txList)
	if err != nil {
		return nil, ScriptError{
			scriptName: fname(s.fpath),
			function:   nameFn,
			errorType:  TransactionListError,
			text:       err.Error(),
		}
	}
	return txs, nil
}
