package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var RootCommand = &cobra.Command{
	Use:   "sdk-modelgen",
	Short: "utility to generate sdk models",
	Long:  `utility to generate sdk models`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

type Options struct {
	ServiceID string
	OpName    string
}

var options Options

func init() {
	cobra.OnInitialize()
	RootCommand.PersistentFlags().StringVarP(&options.ServiceID, "service", "s", "", "service")
	RootCommand.PersistentFlags().StringVarP(&options.OpName, "operation", "o", "", "operation")
}

type Service struct {
	ID         string
	Downloader Downloader
	Generator  Generator
}

func fetchWithCache(url string, cachePath string) (doc *goquery.Document) {
	prepareDir(cachePath)

	if fileExists(cachePath) {
		fmt.Println("File exists, so won't fetch it and skip ...")
		f, err := os.Open(cachePath)
		panicIfErr(err)
		reader := bufio.NewReader(f)
		doc, err = goquery.NewDocumentFromReader(reader)
		panicIfErr(err)
	} else {
		fmt.Println("File does not exist, so fetching ...")
		res, err := http.Get(url)
		panicIfErr(err)
		defer res.Body.Close()

		utfBody := transform.NewReader(bufio.NewReader(res.Body), japanese.ShiftJIS.NewDecoder())
		doc, err = goquery.NewDocumentFromReader(utfBody)
		panicIfErr(err)

		html, err := goquery.OuterHtml(doc.Selection)
		panicIfErr(err)

		err = ioutil.WriteFile(cachePath, []byte(html), 0644)
		panicIfErr(err)
		if err != nil {
			panic(err)
		}
	}
	return doc
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func prepareDir(path string) {
	err := os.MkdirAll(filepath.Dir(path), 0777)
	panicIfErr(err)
}
