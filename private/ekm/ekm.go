// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"errors"
	"fmt"
	"time"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/ekres/account"
	"github.com/econbits/econkit/private/lib/auth/credentials"
	"github.com/econbits/econkit/private/lib/auth/session"
	"github.com/econbits/econkit/private/lib/datetime/datetime"
	"github.com/econbits/econkit/private/lib/fin/transaction"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

var (
	loadErrorClass            = ekerrors.MustRegisterClass("LoadError")
	reservedVarErrorClass     = ekerrors.MustRegisterClass("ReservedVarError")
	loginErrorClass           = ekerrors.MustRegisterClass("LoginFunctionError")
	accountListErrorClass     = ekerrors.MustRegisterClass("AccountsFunctionError")
	transactionListErrorClass = ekerrors.MustRegisterClass("TransactionsFunctionError")
)

type EKM struct {
	tn      *starlark.Thread
	fpath   string
	globals starlark.StringDict
}

func New(fpath string) (*EKM, error) {
	name := slang.ScriptId(fpath)
	thread := &starlark.Thread{Name: name, Load: load}
	globals, err := starlark.ExecFile(thread, fpath, nil, epilogue())
	if err != nil {
		return nil, ekerrors.Wrap(
			loadErrorClass,
			err.Error(),
			err,
		)
	}
	err = validateReservedVars(globals)
	if err != nil {
		return nil, ekerrors.Wrap(
			reservedVarErrorClass,
			err.Error(),
			err,
		)
	}
	return &EKM{tn: thread, fpath: fpath, globals: globals}, nil
}

func (m EKM) stringField(name string, defvalue string) string {
	v, ok := m.globals[name]
	if !ok {
		return defvalue
	}
	sv, ok := starlark.AsString(v)
	if !ok {
		return defvalue
	}
	return sv
}

func (m EKM) stringListField(name string, defvalue []string) []string {
	v, ok := m.globals[name]
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

func (m EKM) Description() string {
	return m.stringField(hDescription, defDescription)
}

func (m EKM) URL() string {
	return m.stringField(hURL, defUrl)
}

func (m EKM) License() string {
	return m.stringField(hLicense, defLicense)
}

func (m EKM) Authors() []string {
	return m.stringListField(hAuthors, defAuthors)
}

func (m EKM) Login(cred *credentials.Credentials) (*session.Session, error) {
	nameFn := "login"
	value, err := m.runFn(nameFn, starlark.Tuple{cred}, loginErrorClass)
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

func (m EKM) Accounts(session *session.Session) ([]*account.Account, error) {
	nameFn := "accounts"
	value, err := m.runFn(nameFn, starlark.Tuple{session}, accountListErrorClass)
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

func (m EKM) runFn(
	nameFn string,
	params starlark.Tuple,
	classOnError *ekerrors.Class,
) (starlark.Value, error) {
	fn := m.globals[nameFn]
	if fn == nil {
		return nil, ekerrors.New(
			classOnError,
			fmt.Sprintf("missing %s Function", nameFn),
		)
	}

	value, err := starlark.Call(m.tn, fn, params, nil)
	if err != nil {
		var kerr *ekerrors.EKError
		if errors.As(err, &kerr) {
			return nil, err
		}
		return nil, ekerrors.Wrap(
			classOnError,
			err.Error(),
			err,
		)
	}
	return value, nil
}

func (m EKM) Transactions(session *session.Session, account *account.Account, since time.Time) ([]*transaction.Transaction, error) {
	nameFn := "transactions"
	value, err := m.runFn(nameFn, starlark.Tuple{session, account, datetime.NewFromTime(since)}, transactionListErrorClass)
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
