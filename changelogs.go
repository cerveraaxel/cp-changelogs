package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

const LogsURL = "https://docs.cpanel.net/changelogs/"

// Struct holding the data about the version build
// Ideally int types should be used, but we require leading 0's
// Since we are simply doing data fetch, string types will sufice
type Version struct {
	Major string
	Minor string
	Build string
	Full  string
}

// Method that populates the struct and performs basic validation
func (verInfo *Version) parser(v string) error {

	m, err := regexp.MatchString("(\\d{2,}\\.)+", v)
	if err != nil {
		return err
	}
	if m == false {
		return fmt.Errorf("Not a valid version")
	}
	versions := strings.Split(v, ".")
	// Pop index 0 if "parent" build is included in version
	if versions[0] == "11" {
		_, versions = versions[0], versions[1:]
	}

	if len(versions) < 3 {
		return fmt.Errorf("This does not appear to be a valid full cPanel version")
	}
	verInfo.Major = versions[0]
	verInfo.Minor = versions[1]
	verInfo.Build = versions[2]
	// Set the full version without the Parent build ID
	verInfo.Full = strings.Join(versions, ".")

	return nil
}

// Method to generate URL from version
func (v *Version) genURL() (url, version string) {
	versionPath := v.Major + "-change-log"
	fullUrl := LogsURL + versionPath
	return fullUrl, v.Full
}

// Struct holding a list of all the results gathered by scraping
type Results struct {
	Version []string
	Date    []string
	Cases   []string
}

// Method to get the index value for the position the
// version is located in.
// Kinda hacky but it works, mostly due in part to the website
// having poor DOM element tagging so its not easy to scrape without
// employing more complex scrapers or using methods like this
func (r *Results) indexPlace(fullVersion string) (int, error) {
	err := fmt.Errorf("Verson not found: %v", fullVersion)
	for i, v := range r.Version {
		if v == fullVersion {
			return i, nil
		}
	}
	return 0, err
}

// Struct holding the final data that is returned in the http response
type Output struct {
	Version string         `json:"Version"`
	Details VersionDetails `json:"Details"`
}

type VersionDetails struct {
	Date  string   `json:"ReleaseDate"`
	Cases []string `json:"CaseList"`
}

// Main runner function called by http router
func GetLogs(v string) (*Output, error) {
	verInfo := Version{}

	err := verInfo.parser(v)
	if err != nil {
		return &Output{}, err
	}

	url, version := verInfo.genURL()
	return fetchData(url, version)
}

// Function that fetches the data from URL
func fetchData(url, version string) (*Output, error) {
	r, err := soup.Get(url)
	if err != nil {
		return &Output{}, err
	}

	doc := soup.HTMLParse(r)
	data := doc.Find("div", "class", "col-md-9")

	versionList := data.FindAll("h3")
	dateList := data.FindAll("h5")
	caseList := data.FindAll("ul")

	// Create a Results Struct
	res := Results{}
	// Populate the struct
	for i, _ := range versionList {

		d := dateList[i].FullText()
		v := versionList[i].FullText()
		c := caseList[i].FullText()

		res.Version = append(res.Version, v)
		res.Date = append(res.Date, d)
		res.Cases = append(res.Cases, c)
	}

	i, err := res.indexPlace(version)
	if err != nil {
		return &Output{}, err
	}

	cList := removeEmpty(strings.Split(res.Cases[i], "\n"))

	return &Output{
		Version: res.Version[i],
		Details: VersionDetails{
			Date:  res.Date[i],
			Cases: cList,
		},
	}, nil
}

// Remove empty entries from the case slice
func removeEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
