package go_training

import (
	"testing"
	"path/filepath"
	"os"
	"log"
	"strings"
	"os/exec"
	"fmt"
)

func TestPath(t *testing.T) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	path:=strings.Replace(dir, "\\", "/", -1)

	t.Log(filepath.Join(path,".."))
	t.Log(path)
}

func TestPath2(t *testing.T) {

	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	s = strings.Replace(s, "\\", "/", -1)
	s = strings.Replace(s, "\\\\", "/", -1)
	i := strings.LastIndex(s, "/")
	path := string(s[0 : i+1])
	t.Log(path)

}

func TestPath3(t *testing.T) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	t.Log(exPath)
}