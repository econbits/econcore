//Copyright (C) 2020  Germán Fuentes Capella
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const (
	iso4217Url         = "https://www.currency-iso.org/dam/downloads/lists/list_one.xml"
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
	downloadISO4217XML()
	currencies := loadCurrencies()
	writeGoCurrencies(currencies)
}

func loadCurrencies() map[string]Currency {
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		panic(err.Error())
	}
	strdata := string(data)
	strdata = strings.Replace(strdata, ">N.A.<", "><", -1)
	data = []byte(strdata)
	var isoXML ISOXML
	err = xml.Unmarshal(data, &isoXML)
	if err != nil {
		panic(err.Error())
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
	return currencyMap
}

func writeGoCurrencies(currencyMap map[string]Currency) {
	file, err := os.Create(goPath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	tmpl, err := template.New("GoCurrencies").Parse(currenciesTemplate)
	if err != nil {
		panic(err.Error())
	}

	err = tmpl.Execute(file, currencyMap)
	if err != nil {
		panic(err.Error())
	}
}

func downloadISO4217XML() {
	resp, err := http.Get(iso4217Url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(fmt.Errorf("Received response code: %d", resp.StatusCode))
	}

	file, err := os.Create(dataPath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err.Error())
	}
}
