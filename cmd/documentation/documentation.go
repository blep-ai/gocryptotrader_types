package main

import (
	"html/template"
)

type Contributor struct {
	Login         string `json:"login"`
	URL           string `json:"html_url"`
	Contributions int    `json:"contributions"`
}
type Config struct {
	GithubRepo      string     `json:"githubRepo"`
	Exclusions      Exclusions `json:"exclusionList"`
	RootReadme      bool       `json:"rootReadmeActive"`
	LicenseFile     bool       `json:"licenseFileActive"`
	ContributorFile bool       `json:"contributorFileActive"`
}
type Exclusions struct {
	Files       []string `json:"Files"`
	Directories []string `json:"Directories"`
}
type DocumentationDetails struct {
	Directories  []string
	Tmpl         *template.Template
	Contributors []Contributor
	Config       *Config
}
type Attributes struct {
	Name            string
	Contributors    []Contributor
	NameURL         string
	Year            int
	CapitalName     string
	DonationAddress string
}
