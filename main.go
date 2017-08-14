package main

import (
	"encoding/json"
	"fmt"
	"github.com/fagnercarvalho/momus"
	"io/ioutil"
	"log"
)

func main() {
	healthChecker := momus.New(&momus.Config{OnlyDeadLinks: false})
	links := healthChecker.GetLinks("http://fagner.co")

	prettyPrint(links)
	saveJson(links)
	saveHtml(links)
}

func prettyPrint(linkResults []momus.LinkResult) {
	for _, linkResult := range linkResults {
		fmt.Printf("%d | %s \n", linkResult.StatusCode, linkResult.Link)
	}
}

func saveJson(links []momus.LinkResult) {
	entries, err := json.MarshalIndent(links, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("D:/output.json", []byte(string(entries)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func saveHtml(linkResults []momus.LinkResult) {
	var htmlContent string = "<html><head></head><body>"

	htmlContent += "<table border=\"1\"><tr> <th>Link</th> <th>Status code</th> </tr>"
	for _, linkResult := range linkResults {
		htmlContent += "<tr>"
		htmlContent += fmt.Sprintf("<td><a href=\"%s\">%s</a></td> ", linkResult.Link, linkResult.Link)

		if linkResult.StatusCode == 200 {
			htmlContent += fmt.Sprintf("<td style=\"color: green; font-weight: bold\">%d</td>", linkResult.StatusCode)
		} else {
			htmlContent += fmt.Sprintf("<td style=\"color: red; font-weight: bold\">%d</td>", linkResult.StatusCode)
		}
		htmlContent += "</tr>"
	}
	htmlContent += "</table></body></html>"

	err := ioutil.WriteFile("D:/output.html", []byte(htmlContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
