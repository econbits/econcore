//Copyright (C) 2020  Germán Fuentes Capella

package script

import (
	"fmt"
)

type ScriptError struct {
	scriptName string
	function   string
	text       string
}

func (se ScriptError) Error() string {
	if len(se.scriptName) > 0 {
		return fmt.Sprintf("[%s][%s] %s", se.scriptName, se.function, se.text)
	}
	return fmt.Sprintf("{%s} %s", se.function, se.text)
}

func newLoginError(scriptName string, text string) error {
	return ScriptError{scriptName: scriptName, function: "login", text: text}
}

func newAccountError(scriptName string, text string) error {
	return ScriptError{scriptName: scriptName, function: "account", text: text}
}

func newBuiltinError(function string, text string) error {
	return ScriptError{scriptName: "", function: function, text: text}
}
