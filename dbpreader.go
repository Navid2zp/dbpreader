package dbpreader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Queries the json version of dbpedia
// Return a map of results returned by dbpedia
// Results might be contain resources that you didn't ask for
// Use FindResource() method to extract the exact resource
func Query(query string) (*DBPediaResult, error) {
	baseURL := "http://dbpedia.org/data/%s.json"
	// Spaces should be replaced by _ in dbpedia page urls
	cleanedQuery := strings.ReplaceAll(query, " ", "_")

	res, err := http.Get(fmt.Sprintf(baseURL, cleanedQuery))
	if err != nil {
		return nil, errors.New("request failed: " + err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("request failed with status code " + string(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result DBPediaResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.New("failed to unmarshal the response: " + err.Error())
	}
	return &result, nil
}

// The json version of dbpedia returns a set of results
// This method will find and return the exact resource
func (r DBPediaResult) FindResource(name string) DBPediaResource {
	baseName := "http://dbpedia.org/resource/"
	return r[baseName+name]
}

// Finds the wikidata item url
func (r DBPediaResource) GetWikiDataItem() string {
	// sameAs is a list of links to other sites/sections for this resource
	sameAs := r["http://www.w3.org/2002/07/owl#sameAs"]

	for _, same := range sameAs {
		// sameAs values are all strings
		// Convert to string in order to do the comparision
		stringValue := same.Value.(string)

		// Remove "https://, http:// and www." for safer comparision
		cleaned := strings.Replace(strings.Replace(strings.Replace(stringValue, "https://", "", 1), "http://", "", 1), "www.", "", 1)
		if len(cleaned) < 20 {
			return ""
		}
		if string(cleaned[0:20]) == "wikidata.org/entity/" {
			return stringValue
		}
	}
	return ""
}
