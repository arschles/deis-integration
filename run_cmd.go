package main

import (
  "os/exec"
  "errors"
  "strings"
  "os"
)

var ErrInvalidCmdString = errors.New("invalid command string")

func createCmd(str string) (*exec.Cmd, error) {
  spl := strings.Split(str, " ")
  if len(spl) < 1 {
    return nil, ErrInvalidCmdString
  }
  cmd := exec.Command(spl[0], spl[1:]...)
  cmd.Env = os.Environ()
  return cmd, nil
}
