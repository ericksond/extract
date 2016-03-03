package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Proxy       proxy
	Appdynamics appdynamics
	Metrics     map[string]metric
}

type proxy struct {
	URL string
}

type appdynamics struct {
	Enabled      bool
	Proxy        bool
	ProxyURL     string
	User         string
	Pass         string
	BaseURL      string
	BaseFilePath string
	Name         string
	Path         string
}

type metric struct {
	Name string
	Path string
}

func main() {
	// toml configurations
	var config tomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
		return
	}

	// flags
	appdApp := flag.NewFlagSet("appd", flag.ExitOnError)
	metricFlag := appdApp.String("metric", "all", "Appdynamics metric to call")

	splunkApp := flag.NewFlagSet("splunk", flag.ExitOnError)
	//searchFlag := splunkApp.String("search", "s", "Splunk saved search to call")

	if len(os.Args) == 1 {
		fmt.Println("usage: extract <app> [<args>]")
		fmt.Println("The most commonly used extract apps are: ")
		fmt.Println(" appdynamics   AppDynamics")
		fmt.Println(" splunk  Splunk")
		return
	}

	switch os.Args[1] {
	case "appdynamics":
		appdApp.Parse(os.Args[2:])
	case "splunk":
		splunkApp.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not a valid app.\n", os.Args[1])
		os.Exit(2)
	}

	// Parsed appdynamics
	if appdApp.Parsed() {
		if config.Appdynamics.Enabled == true {
			if *metricFlag == "all" || *metricFlag == "" {
				for _, metric := range config.Metrics {
					//fmt.Printf("Metric: %s (%s)\n", metric.Name, metric.Path)
					appd := &appdynamics{
						User:         config.Appdynamics.User,
						Pass:         config.Appdynamics.Pass,
						BaseURL:      config.Appdynamics.BaseURL,
						BaseFilePath: config.Appdynamics.BaseFilePath,
						Path:         metric.Path,
						Name:         metric.Name,
						Proxy:        config.Appdynamics.Proxy,
						ProxyURL:     config.Proxy.URL,
					}
					appd.CreateAppdynamicsFile()
				}
			} else {
				appd := &appdynamics{
					User:         config.Appdynamics.User,
					Pass:         config.Appdynamics.Pass,
					BaseURL:      config.Appdynamics.BaseURL,
					BaseFilePath: config.Appdynamics.BaseFilePath,
					Path:         config.Metrics[*metricFlag].Path,
					Name:         config.Metrics[*metricFlag].Name,
					Proxy:        config.Appdynamics.Proxy,
					ProxyURL:     config.Proxy.URL,
				}
				appd.CreateAppdynamicsFile()

			}
		} else {
			fmt.Println("Appdynamics extraction is disabled.")
		}
	}

	// Static Server for JSON files
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("json"))))
}
