package main

import (
    "fmt"
	"os"
	"log"
	"path/filepath"
	"io/ioutil"
)

 
func main() {
	
	registryFile := "carpeta/foo.txt"
	// Create directory if it does not already exist.
	registryPath := filepath.Dir(registryFile)
	err := os.MkdirAll(registryPath, 0750)
	if err != nil {
		fmt.Errorf("Failed to created registry file dir %s: %v", registryPath, err)
		log.Fatal(err)
		return
	}

	// Check if files exists
	fileInfo, err := os.Lstat(registryFile)
	if os.IsNotExist(err) {
		log.Print("No registry file found under: %s. Creating a new registry file.", registryFile) 
		
		d1 := []byte("hello\ngo\n")
		err = ioutil.WriteFile(registryFile, d1, 0644)
		return

	}else{
		log.Print("registry file found under: %s. ", registryFile)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	// Check if regular file, no dir, no symlink
	if !fileInfo.Mode().IsRegular() {
		// Special error message for directory
		if fileInfo.IsDir() {
			fmt.Errorf("Registry file path must be a file. %s is a directory.", registryFile)
			return
		}
		fmt.Errorf("Registry file path is not a regular file: %s", registryFile)
		return
	}

	log.Print("Registry file set to: %s", registryFile)
}
