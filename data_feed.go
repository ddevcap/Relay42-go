package relay42

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DataFeedService holds the R42 service
type DataFeedService service

// GetEntries returns datafeed entries by feedPrefix and keys
func (service *DataFeedService) GetEntries(feedPrefix string, keys ...string) (map[string]interface{}, error) {
	method := http.MethodGet
	path := fmt.Sprintf("v1/site-%s/datafeeds/%s/entries", service.r.siteID, feedPrefix)
	query := url.Values{}
	query.Set("key", strings.Join(keys, ","))

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var entries map[string]interface{}
	err = service.r.do(req, &entries)

	return entries, err
}

// AddEntries adds entries to a datafeed by feedPrefix
func (service *ProfileService) AddEntries(feedPrefix string, entries []*DataFeedEntry) error {
	method := http.MethodPost
	path := fmt.Sprintf("v1/site-%s/datafeeds/%s/entries", service.r.siteID, feedPrefix)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(entries)

	req, err := service.r.newRequest(method, path, nil, b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	return service.r.do(req, nil)
}

// DeleteEntries deletes entries from datafeed by feedPrefix and keys
func (service *DataFeedService) DeleteEntries(feedPrefix string, keys ...string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("v1/site-%s/datafeeds/%s/entries", service.r.siteID, feedPrefix)
	query := url.Values{}
	query.Set("key", strings.Join(keys, ","))

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}

// GetEntry returns an entry by feedPrefix and key
func (service *DataFeedService) GetEntry(feedPrefix, key string) (map[string]interface{}, error) {
	method := http.MethodGet
	path := fmt.Sprintf("v1/site-%s/datafeeds/%s/entries/%s", service.r.siteID, feedPrefix, key)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var entry map[string]interface{}
	err = service.r.do(req, &entry)

	return entry, err
}

// AddEntry adds an entry by feedPrefix
func (service *ProfileService) AddEntry(feedPrefix, entry *DataFeedEntry) error {
	method := http.MethodPost
	path := fmt.Sprintf("v1/site-%s/datafeeds/%s/entries/%s", service.r.siteID, feedPrefix, entry.Key)
	query := url.Values{}
	query.Set("ttl", strconv.Itoa(entry.TTL))

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(entry.Values)

	req, err := service.r.newRequest(method, path, nil, b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	return service.r.do(req, nil)
}

// DeleteEntry deletes an entry by feedPrefix and key
func (service *DataFeedService) DeleteEntry(feedPrefix, key string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("v1/site-%s/datafeeds/%s/entries/%s", service.r.siteID, feedPrefix, key)

	req, err := service.r.newRequest(method, path, nil, nil)

	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}
