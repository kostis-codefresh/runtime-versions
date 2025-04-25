package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

// HelmDependency represents a single dependency in the Helm Chart.yaml file
type HelmDependency struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

// ChartMetadata represents the structure of the Chart.yaml file
type ChartMetadata struct {
	Dependencies []HelmDependency `yaml:"dependencies"`
}

type AppVersion struct {
	Name               string `yaml:"name"`
	Version            string `yaml:"version"`
	ApplicationVersion string `yaml:"appVersion"`
}

// extractHelmDependencies reads a Helm Chart.yaml file and extracts all dependency names and versions
func extractHelmDependencies(chartYAMLContent string) ([]HelmDependency, error) {
	var metadata ChartMetadata

	// Parse the YAML content
	err := yaml.Unmarshal([]byte(chartYAMLContent), &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Chart.yaml: %w", err)
	}

	return metadata.Dependencies, nil
}

func extractAppVersion(chartYAMLContent string) string {
	var appVersion AppVersion

	// Parse the YAML content
	err := yaml.Unmarshal([]byte(chartYAMLContent), &appVersion)
	if err != nil {
		log.Printf("failed to parse Chart.yamls: %v", err)
		return ""
	}

	return appVersion.ApplicationVersion
}

func extractArgoDependencies(chartYAMLContent string, gitOpsRuntime *GitOpsRuntimeRelease) {
	// Extract dependencies from Helm projects
	dependencies, err := extractHelmDependencies(chartYAMLContent)
	if err != nil {
		log.Printf("Error extracting dependencies: %v", err)
	}

	// Print the extracted dependencies
	fmt.Println("Extracted Helm Dependencies:")
	for _, dep := range dependencies {
		fmt.Printf("Name: %s, Version: %s\n", dep.Name, dep.Version)

		if dep.Name == "argo-cd" {
			gitOpsRuntime.ArgoCD.ArgoHelmChart = VersionDetails{
				Name:    dep.Name,
				Version: dep.Version,
				GitTag:  dep.Name + "-" + dep.Version,
				// Date:    gitOpsRuntime.GitOpsRuntime.Date,
				Link: generateReleaseNotesURL(ArgoHelmRepo, dep.Name+"-"+dep.Version),
			}
		} else if dep.Name == "argo-rollouts" {
			gitOpsRuntime.ArgoRollouts.ArgoHelmChart = VersionDetails{
				Name:    dep.Name,
				Version: dep.Version,
				GitTag:  dep.Name + "-" + dep.Version,
				// Date:    gitOpsRuntime.GitOpsRuntime.Date,
				Link: generateReleaseNotesURL(ArgoHelmRepo, dep.Name+"-"+dep.Version),
			}
		} else if dep.Name == "argo-workflows" {
			gitOpsRuntime.ArgoWorkflows.ArgoHelmChart = VersionDetails{
				Name:    dep.Name,
				Version: dep.Version,
				GitTag:  dep.Name + "-" + dep.Version,
				// Date:    gitOpsRuntime.GitOpsRuntime.Date,
				Link: generateReleaseNotesURL(ArgoHelmRepo, dep.Name+"-"+dep.Version),
			}
		} else if dep.Name == "argo-events" {
			gitOpsRuntime.ArgoEvents.ArgoHelmChart = VersionDetails{
				Name:    dep.Name,
				Version: dep.Version,
				GitTag:  dep.Name + "-" + dep.Version,
				// Date:    gitOpsRuntime.GitOpsRuntime.Date,
				Link: generateReleaseNotesURL(ArgoHelmRepo, dep.Name+"-"+dep.Version),
			}
		}
	}

	fmt.Printf("Argo CD Version: %s\n", gitOpsRuntime.ArgoCD.ArgoHelmChart.Version)
	fmt.Printf("Argo Rollouts Version: %s\n", gitOpsRuntime.ArgoRollouts.ArgoHelmChart.Version)
	fmt.Printf("Argo Workflows Version: %s\n", gitOpsRuntime.ArgoWorkflows.ArgoHelmChart.Version)
	fmt.Printf("Argo Events Version: %s\n", gitOpsRuntime.ArgoEvents.ArgoHelmChart.Version)
}
