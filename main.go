//
// UnityProjectCleaner - Delete Library directory in the Unity project
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Songmu/prompter"
)

var dryRun bool
var automaticYes bool
var targetDir string

func isUnityProject(path string) bool {
	// If path + "ProjectSettings/ProjectVersion.txt" exists, the directory is considered the Unity Project directory.
	projectVersionPath := filepath.Join(path, "ProjectSettings", "ProjectVersion.txt")

	_, err := os.Stat(projectVersionPath)
	if err == nil {
		return true
	}

	return false
}

func findUnityProjectDir(path string, dryRun bool) {
	if isUnityProject(path) == true {
		fmt.Printf("unity project found...path=%s\n", path)

		// Delete [Library, Logs, obj, Temp] directories
		targetDirs := [...]string{"Library", "Logs", "obj", "Temp"}

		for _, dir := range targetDirs {
			d := filepath.Join(path, dir)

			// If in dryRun mode, do not delete files
			if dryRun == true {
				fmt.Printf("   delete... %s (dryRun mode)\n", d)
			} else {
				fmt.Printf("   delete... %s\n", d)
				os.RemoveAll(d)
			}
		}
	} else {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if file.IsDir() {
				subDir := filepath.Join(path, file.Name())
				findUnityProjectDir(subDir, dryRun)
			}
		}
	}
}

func flagUsage() {
	fmt.Fprintf(os.Stderr, "  usage: \n")
	fmt.Fprintf(os.Stderr, "      %s [-d] targetDirectory \n", filepath.Base(os.Args[0]))
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  options: \n")
	fmt.Fprintf(os.Stderr, "      -d  dry run mode\n")
	fmt.Fprintf(os.Stderr, "      -y  automatic yes to prompt\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  examples: \n")
	fmt.Fprintf(os.Stderr, "      %s /home/test/unity_projects/ \n", filepath.Base(os.Args[0]))
	fmt.Fprintf(os.Stderr, "\n")
}

func main() {
	flag.Usage = flagUsage
	flag.BoolVar(&dryRun, "d", false, "dry run mode")
	flag.BoolVar(&automaticYes, "y", false, "automatic yes to prompt")
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	if automaticYes == false {
		if prompter.YN("Do you want to start the clean up process?", false) == false {
			fmt.Printf("Abort clean up process...")
			os.Exit(1)
		}
	}

	fmt.Printf("Start clean up process...")

	targetDir = flag.Args()[0]

	findUnityProjectDir(targetDir, dryRun)
}
