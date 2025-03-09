package ctx

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cli/go-gh"
)

var ghcliExtensionDeps = []string{
	"valeriobelli/gh-milestone",
}

type Current struct {
	RepoName  string `json:"repo_name"`
	RepoOwner string `json:"repo_owner"`
	Branch    string `json:"branch"`
	Path      string `json:"path"`
	Taskfile  string `json:"taskfile"`
}

func (c *Current) String() string {
	return fmt.Sprintf("Current{RepoName: %v, RepoOwner: %v, Branch: %v, Path: %v,  Taskfile: %v}", c.RepoName, c.RepoOwner, c.Branch, c.Path, c.Taskfile)
}

// cachedCurrent stores the current context to avoid repeated expensive operations
var cachedCurrent *Current

func fetchCurrent() (*Current, error) {
	// Return cached current if available
	if cachedCurrent != nil {
		// Verify we're still in the same directory
		currentDir, err := WorkingDir()
		if err == nil && currentDir == cachedCurrent.Path {
			return cachedCurrent, nil
		}
	}

	repo, err := gh.CurrentRepository()
	if err != nil {
		return nil, err
	}

	wrkDir, err := WorkingDir()
	if err != nil {
		return nil, err
	}

	currRepo := repo.Name()
	currOrg := repo.Owner()

	// Check if org directory exists before trying to download
	exists, orgPath, err := orgDirExists(currOrg)
	if err != nil {
		return nil, err
	}

	// Only download if the directory doesn't exist
	if !exists {
		orgPath, err = DownloadOrgData(currOrg)
		if err != nil {
			return nil, err
		}
	}

	// Get branch information in the background
	branch, err := CurrentBranch()
	if err != nil {
		return nil, err
	}

	// Create and cache the current context
	cachedCurrent = &Current{
		RepoName:  currRepo,
		RepoOwner: currOrg,
		Path:      wrkDir,
		Branch:    branch,
		Taskfile:  orgPath,
	}

	return cachedCurrent, nil
}

func WorkingDir() (string, error) {
	return os.Getwd()
}

// CurrentBranch returns the name of the current git branch in the working directory
func CurrentBranch() (string, error) {
	// Create command to run "git branch"
	cmd := exec.Command("git", "branch", "--show-current")

	// Create buffer to capture output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Get output and trim whitespace
	branchName := strings.TrimSpace(out.String())
	return branchName, nil
}
