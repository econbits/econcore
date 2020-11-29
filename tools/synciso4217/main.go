//Copyright (C) 2020  Germán Fuentes Capella

package main

import (
	"encoding/xml"
	"io/ioutil"
	"strings"

	"github.com/econbits/econkit/private/files"
)

// var for testing purposes
var (
	iso4217Url = "https://www.currency-iso.org/dam/downloads/lists/list_one.xml"
)

const (
	dataPath           = "../../configs/iso4217.xml"
	goPath             = "../../pkg/currency/init_iso4217.go"
	currenciesTemplate = `//Copyright (C) 2020  Germán Fuentes Capella
// This file is auto-generated. DO NOT EDIT

package currency

func init() {
{{- range $key, $value := . }}
	currencies["{{ $key }}"] = Currency{
		name:  "{{$value.Name}}",
		code:  "{{$value.Code}}",
		id:    {{$value.Id}},
		units: {{$value.Units}},
	}
{{- end }}
}
`
)

type Currency struct {
	XMLName xml.Name `xml:"CcyNtry"`
	Name    string   `xml:"CcyNm"`
	Code    string   `xml:"Ccy"`
	Id      uint32   `xml:"CcyNbr"`
	Units   uint8    `xml:"CcyMnrUnts"`
}

type Currencies struct {
	XMLName xml.Name   `xml:"CcyTbl"`
	List    []Currency `xml:"CcyNtry"`
}

type ISOXML struct {
	XMLName    xml.Name   `xml:"ISO_4217"`
	Currencies Currencies `xml:"CcyTbl"`
}

func main() {
	err := files.Download(iso4217Url, dataPath)
	if err != nil {
		panic(err)
	}
	currencies, err := load(dataPath)
	if err != nil {
		panic(err)
	}
	err = files.WriteFromTemplate(goPath, currenciesTemplate, currencies)
	if err != nil {
		panic(err)
	}
}

func load(path string) (map[string]Currency, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return map[string]Currency{}, err
	}
	strdata := string(data)
	strdata = strings.Replace(strdata, ">N.A.<", "><", -1)
	data = []byte(strdata)
	var isoXML ISOXML
	err = xml.Unmarshal(data, &isoXML)
	if err != nil {
		return map[string]Currency{}, err
	}
	currencyMap := map[string]Currency{}
	for _, c := range isoXML.Currencies.List {
		if len(c.Code) == 0 {
			continue
		}
		_, exists := currencyMap[c.Code]
		if !exists {
			currencyMap[c.Code] = c
		}
	}
	return currencyMap, nil
}
