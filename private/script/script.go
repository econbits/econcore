// Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"errors"
	"fmt"
	"time"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/ekres/credentials"
	"github.com/econbits/econkit/private/ekres/datetime"
	"github.com/econbits/econkit/private/ekres/session"
	"github.com/econbits/econkit/private/ekres/transaction"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

var (
	loadErrorClass            = ekerrors.MustRegisterClass("LoadError")
	reservedVarErrorClass     = ekerrors.MustRegisterClass("ReservedVarError")
	loginErrorClass           = ekerrors.MustRegisterClass("LoginError")
	accountListErrorClass     = ekerrors.MustRegisterClass("AccountListError")
	transactionListErrorClass = ekerrors.MustRegisterClass("TransactionListError")
	missingFunctionErrorClass = ekerrors.MustRegisterClass("MissingFunctionError")
	functionCallErrorClass    = ekerrors.MustRegisterClass("FunctionCallError")
)

type Script struct {
	tn      *starlark.Thread
	fpath   string
	globals starlark.StringDict
}

func New(fpath string) (Script, error) {
	name := slang.ScriptId(fpath)
	thread := &starlark.Thread{Name: name}
	globals, err := starlark.ExecFile(thread, fpath, nil, epilogue())
	if err != nil {
		return Script{}, ekerrors.Wrap(
			loadErrorClass,
			err.Error(),
			err,
		)
	}
	err = validateGlobalVars(globals)
	if err != nil {
		return Script{}, ekerrors.Wrap(
			reservedVarErrorClass,
			err.Error(),
			err,
		)
	}
	return Script{tn: thread, fpath: fpath, globals: globals}, nil
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

func (s Script) Login(cred *credentials.Credentials) (*session.Session, error) {
	nameFn := "login"
	value, err := s.runFn(nameFn, starlark.Tuple{cred})
	if err != nil {
		return nil, err
	}
	session, ok := value.(*session.Session)
	if !ok {
		return nil, ekerrors.New(
			loginErrorClass,
			fmt.Sprintf("login Function returned object of type '%T' instead of a session", value),
		)
	}
	return session, nil
}

func (s Script) Accounts(session *session.Session) ([]*account.Account, error) {
	nameFn := "accounts"
	value, err := s.runFn(nameFn, starlark.Tuple{session})
	if err != nil {
		return nil, err
	}
	accountList, ok := value.(*starlark.List)
	if !ok {
		return nil, ekerrors.New(
			accountListErrorClass,
			fmt.Sprintf(
				"account Function returned object of type '%T' instead of a list of accounts",
				value,
			),
		)
	}
	accounts, err := LtoAR(accountList)
	if err != nil {
		return nil, ekerrors.Wrap(
			accountListErrorClass,
			err.Error(),
			err,
		)
	}
	return accounts, nil
}

func (s Script) runFn(nameFn string, params starlark.Tuple) (starlark.Value, error) {
	fn := s.globals[nameFn]
	if fn == nil {
		return nil, ekerrors.New(
			missingFunctionErrorClass,
			fmt.Sprintf("missing %s Function", nameFn),
		)
	}

	value, err := starlark.Call(s.tn, fn, params, nil)
	if err != nil {
		var kerr *ekerrors.EKError
		if errors.As(err, &kerr) {
			return nil, err
		}
		return nil, ekerrors.Wrap(
			functionCallErrorClass,
			err.Error(),
			err,
		)
	}
	return value, nil
}

func (s Script) Transactions(session *session.Session, account *account.Account, since time.Time) ([]*transaction.Transaction, error) {
	nameFn := "transactions"
	value, err := s.runFn(nameFn, starlark.Tuple{session, account, datetime.NewFromTime(since)})
	if err != nil {
		return nil, err
	}
	txList, ok := value.(*starlark.List)
	if !ok {
		return nil, ekerrors.New(
			transactionListErrorClass,
			fmt.Sprintf("account Function returned object of type '%T' instead of a list of accounts", value),
		)
	}
	txs, err := LtoTR(txList)
	if err != nil {
		return nil, ekerrors.Wrap(
			transactionListErrorClass,
			err.Error(),
			err,
		)
	}
	return txs, nil
}
