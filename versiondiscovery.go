package main

import (
	"fmt"
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

	GitHubReleaseLimit = 2 // Maximum Number of releases to fetch
)

type VersionDetails struct {
	Name    string
	Version string
	Date    time.Time
	Link    string
}

type ArgoProject struct {
	ArgoHelmChart  VersionDetails
	SourceCodeRepo VersionDetails
}

type GitOpsRuntimeRelease struct {
	GitOpsRuntime VersionDetails
	ArgoCD        ArgoProject
	ArgoRollouts  ArgoProject
	ArgoWorkflows ArgoProject
	ArgoEvents    ArgoProject
}

// Final template that contains all information. Rendered with web/index.html.tpl
type templateData struct {
	Now           time.Time //when the page was generated
	VersionsFound []GitOpsRuntimeRelease
}

func discoverVersions() []GitOpsRuntimeRelease {

	// Helper function to parse time from string
	parseTime := func(dateStr string) time.Time {
		parsedTime, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			log.Printf("Error parsing time: %v\n", err)
			return time.Time{} // Return zero value of time.Time on error
		}
		return parsedTime
	}
	repoURL := GitOpsRuntime
	versions := []GitOpsRuntimeRelease{}

	releases, err := fetchGithubReleases(repoURL, GitHubReleaseLimit)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return versions
	}

	fmt.Printf("Last %d releases:\n", GitHubReleaseLimit)

	//Level 1 - GitOps Runtime
	for _, release := range releases {
		fmt.Printf("Tag: %s, Name: %s, Created At: %s\n", release.TagName, release.Name, release.CreatedAt)

		runtimeRelease := GitOpsRuntimeRelease{
			GitOpsRuntime: VersionDetails{
				Name:    release.Name,
				Version: release.TagName,
				Date:    parseTime(release.CreatedAt),
				Link:    generateReleaseNotesURL(GitOpsRuntime, release.TagName),
			},
		}
		// Level 2 - ArgoCD Helm chart
		findArgoHelmDetails(release.TagName, &runtimeRelease)
		versions = append(versions, runtimeRelease)
	}

	return versions
}

func findArgoHelmDetails(tagName string, gitOpsRuntime *GitOpsRuntimeRelease) {

	yamlContent := fetchFileFromGitHub(GitOpsRuntime, tagName, "charts/gitops-runtime/Chart.yaml")

	fmt.Printf("Argo Helm Chart Content:\n%s\n", yamlContent)

	extractArgoDependencies(yamlContent, gitOpsRuntime)

}

func main() {

	// Discover GitHub releases from all related repositories
	versionsFound := discoverVersions()
	log.Printf("Found %d versions of the GitOps Runtime\n", len(versionsFound))

	log.Println("Creating output directory at ./docs")
	err := os.MkdirAll("docs", 0755)
	if err != nil {
		log.Fatal(err)
	}

	copy("web/style.css", "docs/style.css")
	copy("web/favicon.png", "docs/favicon.png")

	log.Println("Rendering HTML report at ./docs/index.html")

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
		log.Fatal("Could not render HTML template: ", err)
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
