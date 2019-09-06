package filemanagement

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

const (
	// PathDaten defines the path to the server Daten
	PathDaten = "/Volumes/Daten"
)

// ListContentForPaths returns the files of path (excluding directories)
func ListContentForPaths(paths []string) []string {
	files := []string{}
	for _, p := range paths {
		ff, err := ioutil.ReadDir(p)
		if err != nil {
			fmt.Printf("[filemanagement/retrieve.go/ListContentForPaths] could not read directory: %s\n", err)
			return files
		}
		for _, f := range ff {
			if f.IsDir() || f.Name()[0] == '.' {
				continue
			}
			files = append(files, f.Name())
		}
	}
	return files
}

// FindPathsForProject searches for paths for projectnr, starting from basepath
func FindPathsForProject(basepath, projectnr string) []string {
	projectPaths := []string{}
	paths := pathsContaining(basepath, projectnr)
	if len(paths) == 0 {
		return projectPaths
	}
	for {
		for _, p := range paths {
			if strings.Split(path.Base(p), " ")[0] == projectnr {
				projectPaths = append(projectPaths, p)
			}
		}
		if len(projectPaths) > 0 {
			return projectPaths
		}
		if len(paths) == 0 {
			return projectPaths
		}
		tempPaths := paths
		paths = []string{}
		for _, p := range tempPaths {
			paths = append(paths, pathsContaining(p, projectnr)...)
		}
	}
}

func pathsContaining(basepath, str string) []string {
	paths := []string{}
	pp, err := ioutil.ReadDir(basepath)
	if err != nil {
		fmt.Println(err)
		return paths
	}
	if len(pp) == 0 {
		return paths
	}
	for _, p := range pp {
		if strings.Contains(str, strings.Split(p.Name(), " ")[0]) && p.IsDir() {
			paths = append(paths, path.Join(basepath, p.Name()))
		}
	}
	return paths
}
