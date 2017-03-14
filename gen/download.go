package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/url"
	"regexp"
)

var downloadCommand = &cobra.Command{
	Use:   "download",
	Short: "download source html files",
	Long:  `download source html files`,
	Run:   download,
}

func init() {
	RootCommand.AddCommand(downloadCommand)
}

func download(cmd *cobra.Command, args []string) {
	var err error
	var services []Service

	servicesJson, err := ioutil.ReadFile("gen/config/services.json")
	err = json.Unmarshal(servicesJson, &services)
	panicIfErr(err)

	for _, service := range services {
		if options.ServiceID != "" && options.ServiceID != service.ID {
			continue
		}

		downloader := service.Downloader

		fmt.Printf("Downloading list\n")
		downloader.downloadOperationList(service.ID)

		fmt.Printf("Downloading %d items\n", len(downloader.OperationList))
		downloader.downloadOperationHTMLs(service.ID)
	}
}

type Downloader struct {
	OperationListURL      string
	OperationListSelector string
	OperationListRegexp   string
	OperationList         []OperationHTML
}

type OperationHTML struct {
	URL    string
	OpName string
	Path   string
}

func (d *Downloader) downloadOperationList(serviceID string) {
	var doc = fetchWithCache(d.OperationListURL, fmt.Sprintf("gen/html/%s/_list.html", serviceID))
	var err error
	var list []OperationHTML

	listURL, err := url.Parse(d.OperationListURL)
	panicIfErr(err)

	doc.Find(d.OperationListSelector).Each(func(_ int, s *goquery.Selection) {
		path, _ := s.Attr("href")
		r := regexp.MustCompile(d.OperationListRegexp)
		group := r.FindSubmatch([]byte(path))
		if group != nil {
			url := fmt.Sprintf("%s://%s%s", listURL.Scheme, listURL.Host, path)
			opHTML := OperationHTML{OpName: string(group[1]), URL: url}
			list = append(list, opHTML)
		}
	})
	d.OperationList = list
}

func (d *Downloader) downloadOperationHTMLs(serviceID string) {
	for _, opHTML := range d.OperationList {
		if options.OpName != "" && options.OpName != opHTML.OpName {
			continue
		}

		fmt.Printf("Downloading %s\n", opHTML.OpName)
		path := fmt.Sprintf("gen/html/%s/%s.html", serviceID, opHTML.OpName)
		fetchWithCache(opHTML.URL, path)
	}
}
