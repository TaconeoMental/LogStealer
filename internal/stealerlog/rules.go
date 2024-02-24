package stealerlog

import (
    "log"
    "os"
    "bufio"
    "regexp"
    "path/filepath"
    "madoka.pink/logstealer/pkg/common"
)

type extractorFunction = func(string)

type Rule interface {
    setBaseDir(string)
    CheckRoot() bool
    checkSignature() bool
    IsOptional() bool
}

// File rule //

type File struct {
    Path      string
    fullPath  string
    Signature FileSignature
    Extract   extractorFunction
    Optional  bool
}

func (file *File) setBaseDir(path string) {
    expandedPath, _ := common.ExpandPath(path)
    file.fullPath = filepath.Join(expandedPath, file.Path)
}

func (file *File) CheckRoot() bool {
    if !common.FileExists(file.fullPath) {
        // File doesn't exist
        return false
    }

    if !file.checkSignature() {
        // Signature doesn't match
        return false
    }
    return true
}

func (file *File) checkSignature() bool {
    fname, err := os.Open(file.fullPath)
    if err != nil {
        log.Fatal(err)
    }
    defer fname.Close()

    scanner := bufio.NewScanner(fname)
    for scanner.Scan() {
        if file.Signature.Ready() {
            return true
        }

        b := scanner.Text()
        file.Signature.Match(b)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return false
}

func (file *File) IsOptional() bool {
    return file.Optional
}


// File Signatures //

type FileSignatureExpr struct {
    Regex   *regexp.Regexp
    Checked bool
}

type FileSignature struct {
    Signatures []FileSignatureExpr
}

func NewSignature(signatures ...string) FileSignature {
    compiledSignatures := make([]FileSignatureExpr, len(signatures))

    for index, sig := range signatures {
        compiledSignatures[index].Regex = regexp.MustCompile(sig)
    }
    return FileSignature{Signatures: compiledSignatures}
}

func (fs *FileSignature) Match(line string) {
    for index, _ := range fs.Signatures {
        signature := &fs.Signatures[index]

        if signature.Checked {
            continue
        }

        if signature.Regex.MatchString(line) {
            signature.Checked = true
        }
    }
}

func (fs *FileSignature) Ready() bool {
    for _, sig := range fs.Signatures {
        if !sig.Checked {
            return false
        }
    }
    return true
}

