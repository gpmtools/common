package iofs

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/gpmtools/common/ghc"
)

func GetAppConfigHome() (string, error) {
	var xdgHome string

	// On macOS (darwin), use ~/.config instead of ~/Library/Application Support
	if runtime.GOOS == "darwin" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		xdgHome = filepath.Join(home, ".config")
	} else {
		// For other platforms, use the standard UserConfigDir
		home, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		xdgHome = home
	}

	return filepath.Join(xdgHome, "gh-task"), nil
}

func GetOrgTaskfilesHome(org string) (string, error) {
	confHome, err := GetAppConfigHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(confHome, "src", org), nil
}

func MkDirOrg(org string) (string, error) {
	taskfilesDir, err := GetOrgTaskfilesHome(org)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(taskfilesDir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return taskfilesDir, nil
}

func DownloadOrgData(org string) (string, error) {
	exists, path, err := OrgDirExists(org)
	if err != nil {
		return "", err
	}
	if exists {
		return path, nil
	}

	if !exists {
		// 1. Create taskfiles directory for org
		dlDir, err := MkDirOrg(org)
		if err != nil {
			return "", err
		}

		// 2. Download Taskfile.yml
		out, err := ghc.QueryDownloadFile(org, "Taskfile.yml", dlDir).Exec()
		if err != nil {
			return "", err
		}

		// 3. Download taskfiles directory
		_, err = ghc.QueryDownloadFolder(org, "taskfiles", dlDir).Exec()
		if err != nil {
			return "", err
		}
		return out, nil
	}
	return path, nil
}

func OrgDirExists(org string) (bool, string, error) {
	home, _ := GetOrgTaskfilesHome(org)
	_, err := os.Stat(home)
	if err != nil {
		if os.IsNotExist(err) {
			return false, home, nil
		}
		return false, "", err
	}
	return true, home, nil
}

func RmDirOrg(org string) error {
	taskfilesDir, err := GetOrgTaskfilesHome(org)
	if err != nil {
		return err
	}
	return os.RemoveAll(taskfilesDir)
}
