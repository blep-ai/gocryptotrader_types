package main
import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/common/file"
	"github.com/blep-ai/gocryptotrader_types/core"
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
