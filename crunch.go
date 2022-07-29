package crunch

import (
	"embed"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const projectPath = "crunch-wordlist-code"
const binary = "crunch"

var tmpDir string

//go:embed crunch-wordlist-code/charset.lst
//go:embed crunch-wordlist-code/crunch.c
//go:embed crunch-wordlist-code/crunch.1
//go:embed crunch-wordlist-code/Makefile
var crunchFs embed.FS

func build() {
	var err error

	// create temp directory
	tmpDir, err = ioutil.TempDir("", "crunch")
	if err != nil {
		panic(err)
	}

	// copy embedded files to temp directory
	dir, err := crunchFs.ReadDir(projectPath)
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		src, err := crunchFs.Open(projectPath + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		dst, err := os.Create(tmpDir + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(dst, src)
		if err != nil {
			panic(err)
		}
		src.Close()
		dst.Close()
	}

	// Build the project
	cmd := "cd " + tmpDir + " && make build"
	err = exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		panic(err)
	}

	// check if the binary exists
	_, err = os.Stat(tmpDir + "/" + binary)
	if err != nil {
		panic(err)
	}
}

func init() {
	build()
}

func Run(params ...string) error {
	cmd := fmt.Sprintf("%s/%s %s", tmpDir, binary, strings.Join(params, " "))
	ex := exec.Command("sh", "-c", cmd)
	ex.Stderr = os.Stderr
	ex.Stdout = os.Stdout
	return ex.Run()
}

func GenerateFromCharset(start, end int, charset string) ([]byte, error) {
	out := fmt.Sprintf("%s/%s", tmpDir, "tmp.txt")
	params := fmt.Sprintf("%d %d -f %s/charset.lst %s -o %s", start, end, tmpDir, charset, out)
	err := Run(params)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(out)
	if err != nil {
		return nil, err
	}
	return b, nil
}
