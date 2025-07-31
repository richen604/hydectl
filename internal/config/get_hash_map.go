// ! WIP
package config

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Wallpapers struct {
	Hash  string
	Image string
}

func GetHashMap(wallSources []string, skipStrays bool, verbose bool) ([]Wallpapers, error) {
	var wallHash []string
	var wallList []string

	supportedFiles := []string{"gif", "jpg", "jpeg", "png"}

	for _, wallSource := range wallSources {
		if wallSource == "" {
			continue
		}

		var findArgs []string
		for _, ext := range supportedFiles {
			findArgs = append(findArgs, "-iname", fmt.Sprintf("*.%s", ext), "-o")
		}
		findArgs = findArgs[:len(findArgs)-1] // Remove the last "-o"
		findArgs = append([]string{wallSource, "-type", "f", "!", "-path", "*/logo/*"}, findArgs...)

		cmd := exec.Command("find", findArgs...)
		output, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("failed to execute find command: %w", err)
		}

		files := strings.Split(string(output), "\n")
		for _, file := range files {
			if file == "" {
				continue
			}

			f, err := os.Open(file)
			if err != nil {
				return nil, fmt.Errorf("failed to open file: %w", err)
			}

			h := sha1.New()
			if _, err := io.Copy(h, f); err != nil {
				f.Close()
				return nil, fmt.Errorf("failed to calculate sha1sum: %w", err)
			}
			f.Close()

			hash := hex.EncodeToString(h.Sum(nil))
			wallHash = append(wallHash, hash)
			wallList = append(wallList, file)
		}
	}

	if len(wallList) == 0 {
		if skipStrays {
			return nil, fmt.Errorf("no image found in any source")
		} else {
			return nil, fmt.Errorf("no image found in any source")
		}
	}

	wallpapers := make([]Wallpapers, len(wallList))
	for i := range wallList {
		wallpapers[i] = Wallpapers{
			Hash:  wallHash[i],
			Image: wallList[i],
		}
	}

	if verbose {
		fmt.Println("// Hash Map //")
		for _, wp := range wallpapers {
			fmt.Printf(":: %s=\"%s\" :: %s=\"%s\"\n", wp.Hash, wp.Hash, wp.Image, wp.Image)
		}
	}

	return wallpapers, nil
}
