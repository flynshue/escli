package cmd

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/spf13/viper"
)

// var es *elasticsearch.Client
var es *esapi.API

func esAPI() *esapi.API {
	if es == nil {
		cfg := elasticsearch.Config{
			Addresses: viper.GetStringSlice("addresses"),
			Username:  os.Getenv("USER"),
			Password:  os.Getenv("ES_PASSWORD"),
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: viper.GetBool("insecure-skip-tls-verify"),
				},
			},
		}
		client, err := elasticsearch.NewClient(cfg)
		if err != nil {
			log.Fatalln(err)
		}
		es = client.API
	}
	return es
}
