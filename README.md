# extract

A lightweight application used to extract JSON data from a RESTful endpoint and simply generate a local JSON file. Support calls behind a proxy.

###Compile with Go:
```
git clone https://github.com/ericksond/extract.git
cd extract
go get github.com/BurntSushi/toml
go build
```

###Tasks List:

- [x] Appdynamics
- [ ] Encrypt passphrases in config
- [ ] Splunk
- [ ] Keynote
- [ ] TrackJS

###Usage:
```
extract <app> -config <path_to_configuration_file> [<args>]
```

###Configuration
extract uses TOML configuration file syntax. Example config.toml:

```
[proxy]
url = "http://user:pass@proxy.domain.com:port"

[appdynamics]
enabled = true
proxy = false
user = "user"
pass = "pass"
baseurl = "https://app.saas.appdynamics.com/controller/rest/applications/App/metric-data?"
basefilepath = "json/"

[metrics]
  [metrics.first]
  name = "first"
  path = "API-METRIC-PATH-GOES-HERE&output=json"

  [metrics.second]
  name = "second"
  path = "API-METRIC-PATH-GOES-HERE&output=json"
```

###Example Usage:

#####Extract all Appdynamics metrics from the config file
```
extract appdynamics -config config.toml
```

#####Extract a specific Appdynamics metric
```
exctract appdynamics -config config.toml -metric first
```
