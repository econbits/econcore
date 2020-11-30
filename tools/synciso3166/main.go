//Copyright (C) 2020  Germán Fuentes Capella

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/econbits/econkit/private/files"
)

const (
	waybackSnapshot = "https://web.archive.org/web/20161122071627id_/"
	iso3166OrigUrl  = "http://www.iso.org/iso/home/standards/country_codes/country_names_and_code_elements_txt-temp.htm"
	dataPath        = "../../configs/iso3166.csv"
	goPath          = "../../pkg/country/init_iso3166.go"
	templateTxt     = `//Copyright (C) 2020  Germán Fuentes Capella
// This file is auto-generated. DO NOT EDIT

package country

func init() {
{{- range $key, $value := . }}
	countries["{{ $key }}"] = Country{
		name:   "{{$value.Name}}",
		alpha2: "{{$value.Alpha2}}",
	}
{{- end }}
}
`
)

var (
	iso3166Url = waybackSnapshot + iso3166OrigUrl
)

type Country struct {
	Name   string
	Alpha2 string
}

func main() {
	err := files.Download(iso3166Url, dataPath)
	if err != nil {
		panic(err.Error())
	}
	countries, err := load(dataPath)
	if err != nil {
		panic(err)
	}
	err = files.WriteFromTemplate(goPath, templateTxt, countries)
	if err != nil {
		panic(err)
	}
}

func load(path string) (map[string]Country, error) {
	bdata, err := ioutil.ReadFile(path)
	if err != nil {
		return map[string]Country{}, err
	}

	sdata := strings.ReplaceAll(string(bdata), ";", ",")
	r := csv.NewReader(strings.NewReader(sdata))

	countryMap := map[string]Country{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			err, ok := err.(*csv.ParseError)
			if ok && len(record) == 3 {
				name := strings.TrimSpace(record[1]) + " " + record[0]
				countryMap[record[2]] = Country{Name: name, Alpha2: record[2]}
				continue
			}
			return map[string]Country{}, err
		}
		if len(record) != 2 {
			return map[string]Country{}, fmt.Errorf("Expected 2 records, found %d", len(record))
		}
		countryMap[record[1]] = Country{Name: record[0], Alpha2: record[1]}
	}
	return countryMap, nil
}
