package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
)

func main() {
  wd, err := os.Getwd()
  if err != nil {
    panic(err)
  }

  p := ""
  fmt.Print("Module name: ")
  fmt.Scanf("%s", &p)

  err = os.Mkdir(p, fs.ModePerm)
  if err != nil {
    panic(err)
  }

  _, lookErr := exec.LookPath("go")
  if lookErr != nil {
    panic(lookErr)
  }

  modulePath := path.Join(wd, p)

  cmd := exec.Command("go", "mod", "init", fmt.Sprintf("github.com/brunoeduardodev/%s", p))
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Env = os.Environ()

  cmd.Dir = modulePath
  err = cmd.Run()

  if err != nil {
    panic(err)
  }

  mainGoFilePath := modulePath + "/main.go"

  mainGoContent, err := os.ReadFile("templates/main.go.txt")
  if err != nil {
    panic(err)
  }

  os.WriteFile(mainGoFilePath, mainGoContent, fs.ModePerm)
  fmt.Println("Initialized project.")
}