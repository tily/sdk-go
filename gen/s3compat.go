package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

type S3Compat struct {
	S3ModelSource    S3ModelSource
	S3CompatServices []S3CompatService
}

type S3ModelSource struct {
	Endpoint string
	Version  string
	URLs     map[string]string
}

type S3CompatService struct {
	ID       string
	Metadata Metadata
}

var s3CompatCommand = &cobra.Command{
	Use:   "s3compat",
	Short: "generate s3",
	Long:  "download s3 model jsons and convert them to niftycloud s3-compatible services",
	Run:   s3Compat,
}

func init() {
	RootCommand.AddCommand(s3CompatCommand)
}

func s3Compat(cmd *cobra.Command, args []string) {
	var err error
	var s3Compat = S3Compat{}

	s3CompatJson, err := ioutil.ReadFile("gen/config/s3compat.json")
	err = json.Unmarshal(s3CompatJson, &s3Compat)
	panicIfErr(err)

	endpoint := s3Compat.S3ModelSource.Endpoint
	version := s3Compat.S3ModelSource.Version
	urls := s3Compat.S3ModelSource.URLs
	services := s3Compat.S3CompatServices

	files := downloadS3ModelSource(endpoint, version, urls)

	for file, bytes := range files {
		for _, service := range services {
			generateS3CompatModel(service, file, bytes)
		}
	}
}

func downloadS3ModelSource(endpoint string, version string, urls map[string]string) (files map[string][]byte) {
	files = map[string][]byte{}
	for file, url := range urls {
		if endpoint != "" {
			url = endpoint + url
		}
		response, err := http.Get(url)
		defer response.Body.Close()
		panicIfErr(err)
		bytes, err := ioutil.ReadAll(response.Body)
		panicIfErr(err)
		outPath := fmt.Sprintf("gen/json/s3/%s/%s", version, file)
		prepareDir(outPath)
		err = ioutil.WriteFile(outPath, bytes, 0644)
		panicIfErr(err)
		files[file] = bytes
	}
	return files
}

func generateS3CompatModel(service S3CompatService, file string, bytes []byte) {
	if file == "api-2.json" {
		apiModel := APIModel{}
		err := json.Unmarshal(bytes, &apiModel)
		panicIfErr(err)
		if service.Metadata.APIVersion != "" {
			apiModel.Metadata.APIVersion = service.Metadata.APIVersion
		}
		if service.Metadata.EndpointPrefix != "" {
			apiModel.Metadata.EndpointPrefix = service.Metadata.EndpointPrefix
		}
		if service.Metadata.GlobalEndpoint != "" {
			apiModel.Metadata.GlobalEndpoint = service.Metadata.GlobalEndpoint
		}
		if service.Metadata.ServiceAbbreviation != "" {
			apiModel.Metadata.ServiceAbbreviation = service.Metadata.ServiceAbbreviation
		}
		if service.Metadata.ServiceFullName != "" {
			apiModel.Metadata.ServiceFullName = service.Metadata.ServiceFullName
		}
		if service.Metadata.SignatureVersion != "" {
			apiModel.Metadata.SignatureVersion = service.Metadata.SignatureVersion
		}
		if service.Metadata.XMLNamespace != "" {
			apiModel.Metadata.XMLNamespace = service.Metadata.XMLNamespace
		}
		if service.Metadata.Protocol != "" {
			apiModel.Metadata.Protocol = service.Metadata.Protocol
		}
		bytes, err = json.MarshalIndent(apiModel, "", "  ")
		panicIfErr(err)
	}
	outPath := fmt.Sprintf("models/apis/%s/%s/%s", service.ID, service.Metadata.APIVersion, file)
	prepareDir(outPath)
	err := ioutil.WriteFile(outPath, bytes, 0644)
	panicIfErr(err)
}
