package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type versionDetails struct {
	Name    string
	Version string
}

func discoverVersions() []versionDetails {
	versions := []versionDetails{
		{Name: "Codefresh CLI", Version: "v0.0.1"},
		{Name: "Docker", Version: "v0.0.1"},
	}

	return versions

}

func main() {

	err := os.MkdirAll("docs", 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Creating output directory at docs")

	copy("web/style.css", "docs/style.css")
	copy("web/favicon.png", "docs/favicon.png")

	versionsFound := discoverVersions()

	log.Printf("Found %d version in the Codefresh Artifact hub\n", len(versionsFound))

	tmpl, err := template.ParseFiles("web/index.html.tpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("docs/index.html")
	if err != nil {
		log.Fatal("create file: ", err)
	}
	defer f.Close()

	type templateData struct {
		Now           time.Time
		VersionsFound []versionDetails
	}

	tData := templateData{
		Now:           time.Now(),
		VersionsFound: versionsFound,
	}

	err = tmpl.Execute(f, tData)

	if err != nil {
		log.Fatal("execute: ", err)
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copy(src string, dst string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	checkErr(err)
}
