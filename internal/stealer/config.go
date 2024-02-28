package stealer

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"
	"madoka.pink/logstealer/internal/common"
)

type (
    InfoStealer struct {
        Name    string `yaml:"name"`
        Version string `yaml:"version"`
        Rules   []Rule `yaml:"rules"`
    }

    Rule struct {
        Path       string         `yaml:"path"`
        Signatures []string       `yaml:"signatures,omitempty"`
        Signature  *regexp.Regexp `yaml:"-"`
        Extract    *regexp.Regexp `yaml:"extract,omitempty"`
    }
)

func compileJoint(regexStrigs []string) *regexp.Regexp {
	joinedStrings := strings.Join(regexStrigs, "|")
	return regexp.MustCompile(joinedStrings)
}

func (rule *Rule) compileSignature() {
	rule.Signature = compileJoint(rule.Signatures)
}

func reSubMatchMap(r *regexp.Regexp, str string) []map[string]string {
	match := r.FindAllStringSubmatch(str, -1)
    dataMap := make([]map[string]string, 0, len(match)) // xd
	for _, n := range match {
		subMatchMap := make(map[string]string)
		for i, name := range r.SubexpNames() {
			if i != 0 {
				subMatchMap[name] = n[i]
			}
		}
        dataMap = append(dataMap, subMatchMap)
	}
    return dataMap
}

func (rule Rule) Match(directory string) bool {
	fullPath := filepath.Join(directory, rule.Path)
	return common.FileExists(fullPath)
}

func (rule *Rule) ExtractData(directory string) {
	if !rule.Match(directory) || rule.Extract == nil {
		return
	}

	fullPath := filepath.Join(directory, rule.Path)
	s, err := os.ReadFile(fullPath)
	if err != nil {
		return
	}

    reSubMatchMap(rule.Extract, string(s))
}

func ReadConfigFile(filename string) (*InfoStealer, error) {
	if !common.FileExists(filename) {
		return nil, errors.New("Config file doesn't exist")
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	infoStealer := new(InfoStealer)
	if err := yaml.Unmarshal(yamlFile, infoStealer); err != nil {
		return nil, err
	}

	for index, _ := range infoStealer.Rules {
		infoStealer.Rules[index].compileSignature()
	}

	return infoStealer, nil
}

func (stealer InfoStealer) ExtractData(path string) {
	expandedPath, err := common.ExpandPath(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, rule := range stealer.Rules {
		rule.ExtractData(expandedPath)
	}
}
