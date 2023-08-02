package main

import (
	"encoding/json"
	"fmt"
	"github.com/itzg/restify"
	"github.com/rudolfoborges/pdf2go"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

func nnst(y string) string {
	res, _ := http.Get("https://nnst.ru/PDF/Priem/23/rejting-2023.pdf")

	f, _ := os.Create("temp.pdf")

	io.Copy(f, res.Body)

	reader, err := pdf2go.New("temp.pdf", pdf2go.Config{})

	if err != nil {
		log.Printf("Err: %v", err)
		return ""
	}

	text, err := reader.Text()

	if err != nil {
		log.Printf("Err: %v", err)
		return ""
	}

	fmt.Printf("TEXT: %s", text)

	pages, err := reader.Pages()

	if err != nil {
		log.Printf("Err: %v", err)
		return ""
	}
	htmlFile, _ := os.Create("temp.html")
	//33-37
	for k, page := range pages {
		if k >= 32 && k <= 36 {
			fmt.Println(page.Html())
			html, err := page.Html()
			if err != nil {
				return ""
			}
			html = strings.ReplaceAll(html, "ft332", "ft331")
			t := strings.NewReader(html)
			io.Copy(htmlFile, t)
			fmt.Println(page.Number)
		}
	}

	htmlRaw, err := os.Open("temp.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	parsedHtml, err := html.Parse(htmlRaw)
	if err != nil {
		fmt.Println(err.Error())
	}
	classes := []string{"ft331", "ft340", "ft350", "ft360", "ft370"}
	var node []*html.Node
	for _, v := range classes {
		for _, v := range restify.FindSubsetByClass(parsedHtml, v) {
			node = append(node, v)
		}
	}
	jsonTable, _ := restify.ConvertHtmlToJson(node)
	var nnst []NNST
	var nnstOriginal NNSTTable
	var nnstOriginals []NNSTTable
	json.Unmarshal(jsonTable, &nnst)
	var counter int
	for _, v := range nnst {
		switch counter {
		case 0:
			counter++
		case 1:
			nnstOriginal.Code = v.Text
			counter++
		case 2:
			if strings.TrimSpace(v.Text) == "-" {
				nnstOriginal.IsOriginal = false
			} else {
				nnstOriginal.IsOriginal = true
			}
			counter++
		case 3:
			newString := strings.ReplaceAll(v.Text, ",", ".")
			parseFloat, _ := strconv.ParseFloat(newString, 64)
			nnstOriginal.AvgScore = parseFloat
			if nnstOriginal.IsOriginal {
				nnstOriginals = append(nnstOriginals, nnstOriginal)
			}

			counter = 0
		}
	}
	sort.Slice(nnstOriginals[:], func(i, j int) bool {
		return nnstOriginals[i].AvgScore > nnstOriginals[j].AvgScore
	})
	var rank int
	for k, v := range nnstOriginals {
		if v.AvgScore < 4.3158 {
			rank = k + 1
			break
		}
	}
	return fmt.Sprintf("%v %d/%d", y, rank, len(nnstOriginals))

}

type NNSTTable struct {
	AvgScore   float64
	IsOriginal bool
	Code       string
}
type NNST struct {
	Name       string `json:"name"`
	Attributes struct {
		Style string `json:"style"`
	} `json:"attributes"`
	Class string `json:"class"`
	Text  string `json:"text"`
}
