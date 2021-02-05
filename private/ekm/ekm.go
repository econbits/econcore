// Copyright (C) 2020  Germán Fuentes Capella

package ekm

import (
	"errors"
	"fmt"
	"time"

	"github.com/econbits/econkit/private/ekerrors"
	"github.com/econbits/econkit/private/lib/account/account"
	"github.com/econbits/econkit/private/lib/auth/credentials"
	"github.com/econbits/econkit/private/lib/auth/session"
	"github.com/econbits/econkit/private/lib/datetime/datetime"
	"github.com/econbits/econkit/private/lib/fin/transaction"
	"github.com/econbits/econkit/private/lib/universe"
	"github.com/econbits/econkit/private/slang"
	"go.starlark.net/starlark"
)

var (
	loadErrorClass        = ekerrors.MustRegisterClass("LoadError")
	reservedVarErrorClass = ekerrors.MustRegisterClass("ReservedVarError")
	missingFuncErrorClass = ekerrors.MustRegisterClass("MissingFunctionError")
	wrongReturnErrorClass = ekerrors.MustRegisterClass("WrongReturnError")
	funcCallErrorClass    = ekerrors.MustRegisterClass("FunctionCallError")
)

type EKM struct {
	tn      *starlark.Thread
	fpath   string
	globals starlark.StringDict
}

func New(fpath string) (*EKM, error) {
	name := slang.ScriptId(fpath)
	thread := &starlark.Thread{Name: name, Load: Load}
	globals, err := starlark.ExecFile(thread, fpath, nil, universe.Lib.Load())
	if err != nil {
		return nil, ekerrors.Wrap(
			loadErrorClass,
			err,
			nil,
		)
	}
	err = validateReservedVars(globals)
	if err != nil {
		return nil, ekerrors.Wrap(
			reservedVarErrorClass,
			err,
			nil,
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

func (m EKM) linkCS(nameFn string, ekerr *ekerrors.EKError) {
	fn := m.globals[nameFn]
	fnv, ok := fn.(*starlark.Function)
	if ok {
		ekerr.LinkCS(
			starlark.CallStack{
				starlark.CallFrame{
					Name: nameFn,
					Pos:  fnv.Position(),
				},
			},
		)
	}
}

func (m EKM) Login(cred *credentials.Credentials) (*session.Session, error) {
	nameFn := "login"
	value, err := m.runFn(nameFn, starlark.Tuple{cred})
	if err != nil {
		return nil, err
	}
	session, ok := value.(*session.Session)
	if !ok {
		ekerr := ekerrors.New(
			wrongReturnErrorClass,
			fmt.Sprintf(
				"%s function returned object of type '%s' instead of 'session'",
				nameFn,
				value.Type(),
			),
		)
		m.linkCS(nameFn, ekerr)
		return nil, ekerr
	}
	return session, nil
}

func (m EKM) Accounts(session *session.Session) ([]*account.Account, error) {
	nameFn := "accounts"
	value, err := m.runFn(nameFn, starlark.Tuple{session})
	if err != nil {
		return nil, err
	}
	vlist, ok := value.(*starlark.List)
	if !ok {
		ekerr := ekerrors.New(
			wrongReturnErrorClass,
			fmt.Sprintf(
				"%s: expected 'list', got '%s'",
				nameFn,
				value.Type(),
			),
		)
		m.linkCS(nameFn, ekerr)
		return nil, ekerr
	}
	accounts, ok := LtoAR(vlist)
	if !ok {
		ekerr := ekerrors.New(
			wrongReturnErrorClass,
			fmt.Sprintf(
				"%s: expected 'account', got: %s",
				nameFn,
				vlist.Index(0).Type(),
			),
		)
		m.linkCS(nameFn, ekerr)
		return nil, ekerr
	}
	return accounts, nil
}

func (m EKM) runFn(
	nameFn string,
	params starlark.Tuple,
) (starlark.Value, error) {
	fn := m.globals[nameFn]
	if fn == nil {
		return nil, ekerrors.New(
			missingFuncErrorClass,
			fmt.Sprintf("missing %s function", nameFn),
		)
	}

	value, err := starlark.Call(m.tn, fn, params, nil)
	if err != nil {
		var kerr *ekerrors.EKError
		if errors.As(err, &kerr) {
			return nil, err
		}

		oerr := ekerrors.Wrap(
			funcCallErrorClass,
			err,
			nil,
		)

		var serr *starlark.EvalError
		if errors.As(err, &serr) {
			oerr.LinkCS(serr.CallStack)
		}

		return nil, oerr
	}
	return value, nil
}

func (m EKM) Transactions(session *session.Session, account *account.Account, since time.Time) ([]*transaction.Transaction, error) {
	nameFn := "transactions"
	value, err := m.runFn(nameFn, starlark.Tuple{session, account, datetime.NewFromTime(since)})
	if err != nil {
		return nil, err
	}
	txList, ok := value.(*starlark.List)
	if !ok {
		ekerr := ekerrors.New(
			wrongReturnErrorClass,
			fmt.Sprintf(
				"%s: expected 'list', got '%s'",
				nameFn,
				value.Type(),
			),
		)
		m.linkCS(nameFn, ekerr)
		return nil, ekerr
	}

	txs, ok := LtoTR(txList)
	if !ok {
		ekerr := ekerrors.New(
			wrongReturnErrorClass,
			fmt.Sprintf(
				"%s: expected transaction, got: %s",
				nameFn,
				txList.Index(0).Type(),
			),
		)
		m.linkCS(nameFn, ekerr)
		return nil, ekerr
	}
	return txs, nil
}
