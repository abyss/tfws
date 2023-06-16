package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// Execute the Terraform command to get the current workspace
	currentCmd := exec.Command("terraform", "workspace", "show")
	currentOutput, err := currentCmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute 'terraform workspace show': %s", err)
	}

	currentWorkspace := strings.TrimSpace(string(currentOutput))

	// Execute the Terraform command to get the list of workspaces
	cmd := exec.Command("terraform", "workspace", "list", "-no-color")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute 'terraform workspace list': %s", err)
	}

	// Split the output into individual lines
	workspaces := strings.Split(strings.TrimSpace(string(output)), "\n")

	// Remove leading/trailing spaces and asterisks, and count the number of workspaces
	var validWorkspaces []string
	for _, workspace := range workspaces {
		workspaceName := strings.TrimSpace(workspace)
		workspaceName = strings.ReplaceAll(workspaceName, "*", "")
		workspaceName = strings.TrimSpace(workspaceName)
		if workspaceName != "" {
			validWorkspaces = append(validWorkspaces, workspaceName)
		}
	}

	// Display message if there are no workspaces
	if len(validWorkspaces) == 1 {
		fmt.Println("There are no workspaces in use in the current project.")
		fmt.Println("To create a new one, use 'terraform workspace new'")
		return
	}

	// Create items for the prompt
	items := make([]string, len(validWorkspaces))
	defaultIndex := 0
	for i, workspace := range validWorkspaces {
		items[i] = workspace
		if workspace == currentWorkspace {
			defaultIndex = i
		}
	}

	// Create the select prompt
	prompt := &survey.Select{
		Message: "Select a Terraform workspace:",
		Options: items,
		Default: items[defaultIndex],
	}

	// Run the prompt
	var selectedWorkspace string
	err = survey.AskOne(prompt, &selectedWorkspace, survey.WithPageSize(10))
	if err != nil {
		log.Fatalf("Failed to run the prompt: %s", err)
	}

	fmt.Printf("Selected workspace: %s\n", selectedWorkspace)

	// Select the chosen workspace
	cmd = exec.Command("terraform", "workspace", "select", selectedWorkspace)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to execute 'terraform workspace select': %s", err)
	}
}
