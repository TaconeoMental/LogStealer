package stealer

import (
	"errors"
	//"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"
	"madoka.pink/logstealer/internal/common"
)

// https://zhwt.github.io/yaml-to-go/ :)
type InfoStealer struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Rules   []Rule `yaml:"rules"`
}

type Rule struct {
	Path       string         `yaml:"path"`
	Signatures []string       `yaml:"signatures,omitempty"`
	Signature  *regexp.Regexp `yaml:"-"`
	Extract    *regexp.Regexp `yaml:"extract,omitempty"`
}

func (rule *Rule) compileSignature() {
	fullSignature := strings.Join(rule.Signatures, "|")
	rule.Signature = regexp.MustCompile(fullSignature)
}

func reSubMatchMap(r *regexp.Regexp, str string) /*map[string]string*/ {
	match := r.FindAllStringSubmatch(str, -1)
	for _, n := range match {
		subMatchMap := make(map[string]string)
		for i, name := range r.SubexpNames() {
			if i != 0 {
				subMatchMap[name] = n[i]
			}
		}
		pp.Print(subMatchMap)
	}
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

func FromConfigFile(filename string) (*InfoStealer, error) {
	if !common.FileExists(filename) {
		return nil, errors.New("Config file doesn't exist")
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var infoStealer = new(InfoStealer)

	err = yaml.Unmarshal(yamlFile, infoStealer)
	if err != nil {
		return nil, err
	}

	for index, _ := range infoStealer.Rules {
		infoStealer.Rules[index].compileSignature()
	}

	return infoStealer, nil
}

func (stealer InfoStealer) ExtractData(path string) {
	// TODO: Error!!!
	expandedPath, _ := common.ExpandPath(path)
	for _, rule := range stealer.Rules {
		rule.ExtractData(expandedPath)
	}
}
