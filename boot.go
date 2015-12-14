package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	validExpectations = map[string]struct{}{
		"exitcode": struct{}{},
	}
)

func main() {
	testFolder := flag.String("folder", "./tests", "the folder that holds all of the tests")
	flag.Parse()

	var fileNames []string
	err := filepath.Walk(*testFolder, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "-test.yaml") {
			fileNames = append(fileNames, path)
			log.Printf("Found file %s", path)
		}
		return nil
	})
	if err != nil {
		log.Printf("ERROR walking directory %s (%s)", *testFolder, err)
		os.Exit(1)
	}

	integrationFiles := make([]IntegrationFile, len(fileNames))
	for i, fileName := range fileNames {
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Printf("ERROR reading file %s (%s)", fileName, err)
			os.Exit(1)
		}
		intFile := &IntegrationFile{}
		if err := yaml.Unmarshal(fileBytes, intFile); err != nil {
			log.Printf("ERROR unmarshaling integration file %s (%s)", fileName, err)
			os.Exit(1)
		}
		integrationFiles[i] = *intFile
	}

	for _, intFile := range integrationFiles {
		if err := verifyExpectations(intFile.Expect); err != nil {
			log.Printf("ERROR verifying expectations in file %s (%s)", intFile.Name, err)
			os.Exit(1)
		}
		log.Printf("----- Running %s -----", intFile.Name)
		for j, cmdStr := range intFile.Commands {
			cmd, err := createCmd(cmdStr)
			if err != nil {
				log.Printf("ERROR creating command #%d [%s] in integration file %s (%s)", j, cmdStr, intFile.Name, err)
				os.Exit(1)
			}
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("ERROR running command #%d [%s] in integration file %s (%s)", j, cmdStr, intFile.Name, err)
				os.Exit(1)
			}

			log.Printf("--> %s", string(out))
		}
	}

}

func verifyExpectations(expectations []Expectation) error {
	for _, expectation := range expectations {
		if _, ok := validExpectations[expectation.What]; !ok {
			return fmt.Errorf("invalid expectation %s", expectation.What)
		}
	}
	return nil
}
