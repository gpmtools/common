package ghc

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gpmtools/common/exc"
)

var (
	// QueryProjectList is a command to query the GitHub API for a list of projects
	QueryProjectList = exc.NewCmd("project list --limit 100 --format json -L 100 --jq .items")

	// QueryProjectItemList is a command to query the GitHub API for a list of project items
	QueryProjectItemList = exc.NewCmd("project item-list 4 --owner coindotfi --format json -L 100 --jq .items")

	// QueryUserWhoami is a command to query the GitHub API for the current user
	QueryUserWhoami = exc.NewCmd("api user")
)

func QueryOrgRepos(org string) exc.Command {
	return exc.NewCmdArgs("repo", "list", org, "-L", "100", "--no-archived", "--source", "--visibility", "public", "--json", "nameWithOwner", "--jq", ".[] | .nameWithOwner | split(\"/\")[1]")
}

func QueryDownloadFile(org string, asset string, out string) exc.Command {
	repo := fmt.Sprintf("%s/.github", org)
	outPath := fmt.Sprintf("%s/%s", out, asset)
	return exc.NewCmdArgs("download", repo, asset, "--outfile", outPath)
}

func QueryDownloadFolder(org string, asset string, out string) exc.Command {
	repo := fmt.Sprintf("%s/.github", org)
	outPath := fmt.Sprintf("%s/%s", out, asset)
	return exc.NewCmdArgs("download", repo, asset, "--outdir", outPath)
}

func OrgHasRepo(org, repo string) bool {
	out, err := QueryOrgRepos(org).Exec()
	if err != nil {
		return false
	}
	repos := strings.Split(out, "\n")
	return slices.Contains(repos, repo)
}
