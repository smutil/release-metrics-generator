package main

import (
	"io"
	"log"
	"os"
	"fmt"
	"strings"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/go-git/go-billy/v5/memfs"
	"gopkg.in/yaml.v2"
	"flag"
	"sort"
	"time"
	"math"
	"html/template"
)

type Config struct {
    Git struct {
        URL       string `yaml:"url"`
        Username  string `yaml:"username"`
        Password  string `yaml:"password"`
    }
}

type ChangeDetail struct {
	CommitId string
	Message  string
	Author   string
	Leadtime time.Duration
}

type ReleaseDetail struct {
	TagName 	string
	ReleaseDate time.Time
	Author 		string
	Changes 	[]ChangeDetail
	ChangeVolume int
	LeadTime     float64
}


type ReleaseMetricsPage struct {
	PageTitle   string
	Releases    []ReleaseDetail
}

func main()  {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		log.Fatal(err)
	}
	config := &Config{}
	err := ReadYML(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	generateMetrics(config.Git.URL,config.Git.Username,config.Git.Password)
}

func generateMetrics(scm_repo string, scm_usr string, scm_pwd string) {
	r, _ := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: scm_repo,
        Auth: &http.BasicAuth{
			Username: scm_usr,
			Password: scm_pwd,
		},
	})
	tagrefs, err := r.TagObjects()
	CheckIfError(err)
	var tags[]*object.Tag 
	err = tagrefs.ForEach(func(t *object.Tag) error {
		tags = append(tags, t)	
		return nil
	})
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Tagger.When.After(tags[j].Tagger.When)
	})
	var Releases[]ReleaseDetail
	for i, t := range tags {
		var changes[]ChangeDetail
		var breakme bool = false
		cIter, err := r.Log(&git.LogOptions{From: t.Target})
		var leadtimeMinutes float64
		err = cIter.ForEach(func(c *object.Commit) error {
			hash := c.Hash.String()
			if ( ((len(tags)-1)-i != 0) && c.Hash.String() == tags[i+1].Target.String()) {
				breakme = true
			}
			if !breakme {
				line := strings.Split(c.Message, "\n")
				change := ChangeDetail{CommitId : hash, Message: line[0], Author : c.Author.Name, Leadtime : t.Tagger.When.Sub(c.Author.When)}
				leadtimeMinutes += t.Tagger.When.Sub(c.Author.When).Minutes()
				changes = append(changes, change)
			}
			return nil
		})
		release := ReleaseDetail{TagName : t.Name  , ReleaseDate : t.Tagger.When , Author : t.Tagger.Email, Changes : changes, ChangeVolume: len(changes), LeadTime : math.Trunc(leadtimeMinutes/float64(len(changes)))}
		Releases = append(Releases, release)
		CheckIfError(err)

		tmpl := template.Must(template.ParseFiles("layout.html"))
		data := ReleaseMetricsPage{PageTitle: "Release Metrics", Releases: Releases}

		f, err := os.Create("ReleaseMetrics.html")
		CheckIfError(err)
		err = tmpl.Execute(f, data)
		CheckIfError(err)
        f.Close()
	}
	fmt.Println(Releases)
	CheckIfError(err)
}



func readfile(scm_repo string, scm_usr string, scm_pwd string) {
	fs := memfs.New()
	storer := memory.NewStorage()
	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: scm_repo,
        Auth: &http.BasicAuth{
			Username: scm_usr,
			Password: scm_pwd,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	changelog, err := fs.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, changelog)
}

func CheckIfError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func ReadYML(configPath string, configPointer interface{}) error {
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Init new YAML decoder
	d := yaml.NewDecoder(file)
	if err := d.Decode(configPointer); err != nil {
		return err
	}

	return nil
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}