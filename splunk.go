package main

import "fmt"

// CreateSplunkFile function
func (s *splunk) CreateSplunkFile() {
	fmt.Printf("Processing saved search: %s.\n", s.SavedSearch)
}
