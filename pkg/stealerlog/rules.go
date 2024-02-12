package stealerlog

import (
    "log"
    "os"
    "bufio"
    "errors"
    "regexp"
    "path/filepath"
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
    file.fullPath = filepath.Join(path, file.Path)
}

func (file *File) CheckRoot() bool {
    if _, err := os.Stat(file.fullPath); errors.Is(err, os.ErrNotExist) {
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

        val := signature.Regex.MatchString(line)
        if val {
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


// Rule utils //

func CheckRules(dir string, rules []Rule) bool {
    for _, r := range rules {
        if r.IsOptional() {
            continue
        }
        
        r.setBaseDir(dir)
        if !r.CheckRoot() {
            return false
        }
    }
    return true
}
