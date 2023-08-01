package main

import (
	"encoding/json"
	"fmt"
	"github.com/itzg/restify"
	url2 "net/url"
	"sort"
	"strconv"
	"strings"
)

type LogisticsStruct struct {
	Name     string `json:"name"`
	Class    string `json:"class"`
	Elements []struct {
		Name     string `json:"name"`
		Class    string `json:"class,omitempty"`
		Text     string `json:"text,omitempty"`
		Elements []struct {
			Name string `json:"name"`
		} `json:"elements,omitempty"`
	} `json:"elements"`
}
type LogisticsTable struct {
	AvgScore   float64
	IsOriginal bool
	Agreement  bool
}

func LogisticsGetOriginals(x, y string) string {
	var logisticsOriginal []LogisticsTable
	var logistics []LogisticsStruct
	var tempOriginal LogisticsTable
	url, err := url2.Parse(x)
	content, err := restify.LoadContent(url, "")
	if err != nil {
		fmt.Println(err)
	}
	node := restify.FindSubsetByClass(content, "R0")
	jsonTable, _ := restify.ConvertHtmlToJson(node)
	json.Unmarshal(jsonTable, &logistics)
	for _, v := range logistics[8:] {
		v.Elements[2].Text = strings.Replace(v.Elements[2].Text, ",", ".", -1)
		avgScore, err := strconv.ParseFloat(v.Elements[2].Text, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tempOriginal.AvgScore = avgScore
		if v.Elements[3].Text == "✓" {
			tempOriginal.IsOriginal = true
		} else {
			tempOriginal.IsOriginal = false
		}
		if v.Elements[4].Text == "✓" {
			tempOriginal.Agreement = true
		} else {
			tempOriginal.Agreement = false
		}
		if tempOriginal.IsOriginal {
			logisticsOriginal = append(logisticsOriginal, tempOriginal)
		}

	}
	sort.Slice(logisticsOriginal[:], func(i, j int) bool {
		return logisticsOriginal[i].AvgScore > logisticsOriginal[j].AvgScore
	})
	var rank int
	for k, v := range logisticsOriginal {
		if v.AvgScore < 4.3158 {
			rank = k + 1
			break
		}
	}
	return fmt.Sprintf("%v %d/%d", y, rank, len(logisticsOriginal))
}
