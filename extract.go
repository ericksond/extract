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
	// flags
	appdApp := flag.NewFlagSet("appd", flag.ExitOnError)
	appdMetricFlag := appdApp.String("metric", "all", "Appdynamics metric to call")
	appdConfigFlag := appdApp.String("config", "", "configuration file")

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
		// toml configurations
		var config tomlConfig
		if *appdConfigFlag != "" {
			if _, err := toml.DecodeFile(*appdConfigFlag, &config); err != nil {
				log.Fatal(err)
				return
			}
		} else {
			fmt.Println("TOML Configuration file required; -config <path_to_config_file.toml>")
			os.Exit(3)
		}

		if config.Appdynamics.Enabled == true {
			if *appdMetricFlag == "all" || *appdMetricFlag == "" {
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
					Path:         config.Metrics[*appdMetricFlag].Path,
					Name:         config.Metrics[*appdMetricFlag].Name,
					Proxy:        config.Appdynamics.Proxy,
					ProxyURL:     config.Proxy.URL,
				}
				appd.CreateAppdynamicsFile()

			}
		} else {
			fmt.Println("Appdynamics extraction is disabled.")
		}
	}
}
