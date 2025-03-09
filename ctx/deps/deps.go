package deps

import (
	"strings"

	"github.com/gpmtools/common/exc"
)

var peerExtensions = []string{
	"yuler/gh-download",
}

func CheckPeerDeps() map[string]bool {
	res := make(map[string]bool)
	out, err := exc.NewCmdArgs("extension", "list").Exec()
	if err != nil {
		return res
	}

	for _, ext := range peerExtensions {
		if strings.Contains(out, ext) {
			res[ext] = true
			continue
		} else {
			res[ext] = false
			exc.NewCmdArgs("extension", "install", ext).Exec()
		}
	}
	return res
}
