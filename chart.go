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

func extractComponents() {
	// Example Chart.yaml content
	chartYAMLContent := `
dependencies:
- name: argo-cd
  repository: https://codefresh-io.github.io/argo-helm
  version: 7.8.23-1-cap-v2.14.9-2025-04-20-584fc7f3
- name: argo-events
  repository: https://codefresh-io.github.io/argo-helm
  version: 2.4.7-1-cap-CR-28072
- name: argo-workflows
  repository: https://codefresh-io.github.io/argo-helm
  version: 0.45.2-v3.6.4-cap-CR-27392
- name: argo-rollouts
  repository: https://codefresh-io.github.io/argo-helm
  version: 2.37.3-2-v1.7.2-cap-CR-26082
`

	// Extract dependencies
	dependencies, err := extractHelmDependencies(chartYAMLContent)
	if err != nil {
		log.Fatalf("Error extracting dependencies: %v", err)
	}

	// Print the extracted dependencies
	fmt.Println("Extracted Helm Dependencies:")
	for _, dep := range dependencies {
		fmt.Printf("Name: %s, Version: %s\n", dep.Name, dep.Version)
	}
}
