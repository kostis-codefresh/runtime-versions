package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"time"
)

// GitHub repositories to check for releases
const (
	GitOpsRuntime     = "https://github.com/codefresh-io/gitops-runtime-helm"
	ArgoHelmRepo      = "https://github.com/codefresh-io/argo-helm"
	ArgoCDRepo        = "https://github.com/codefresh-io/argo-cd"
	ArgoRolloutsRepo  = "https://github.com/codefresh-io/argo-rollouts"
	ArgoWorkflowsRepo = "https://github.com/codefresh-io/argo-workflows"
	ArgoEventsRepo    = "https://github.com/codefresh-io/argo-events"
)

type versionDetails struct {
	Name    string
	Version string
	Date    time.Time
	Link    string
}

type GitOpsRuntimeRelease struct {
	GitOpsRuntime versionDetails
	ArgoHelm      versionDetails
	ArgoCD        versionDetails
	ArgoRollouts  versionDetails
	ArgoWorkflows versionDetails
	ArgoEvents    versionDetails
}

// Final template that contains all information. Rendered with web/index.html.tpl
type templateData struct {
	Now           time.Time //when the page was generated
	VersionsFound []GitOpsRuntimeRelease
}

func discoverVersions() []GitOpsRuntimeRelease {
	versions := []GitOpsRuntimeRelease{}

	// Example: Append a GitOpsRuntimeRelease to the slice
	// You can replace this with actual logic to populate the slice
	versions = append(versions, GitOpsRuntimeRelease{})

	findGitHubReleases()
	readContent()

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
	defer func() {
		if cerr := f.Close(); cerr != nil {
			log.Fatal("close file: ", cerr)
		}
	}()

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
	// Open the source file
	srcFile, err := os.Open(src)
	checkErr(err)
	defer func() {
		if cerr := srcFile.Close(); cerr != nil {
			log.Fatal("close src file: ", cerr)
		}
	}()

	// Create the destination file
	dstFile, err := os.Create(dst)
	checkErr(err)
	defer func() {
		if cerr := dstFile.Close(); cerr != nil {
			log.Fatal("close dst file: ", cerr)
		}
	}()

	// Copy the contents from src to dst
	_, err = io.Copy(dstFile, srcFile)
	checkErr(err)
}
