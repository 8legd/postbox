package setup

import (
  "testing"
  "os"
  "path"
  "errors"
  )

func checkDirExists(t *testing.T, dir string) {
  f, err := os.Stat(dir)
  if (err != nil) {
    t.Error(err)
  } else {
    if !f.IsDir() {
      t.Error(errors.New(dir + " exists but is not a directory"))
    }
  }
}

func TestSetupLoggingDirs(t *testing.T) {

  baseDir := "logs/restl"
  hostDir := "localhost"
  logDirs := []string{"sessions","users","errors"}

  // Test succesful creation
  err := SetupLoggingDirs(baseDir,hostDir,logDirs)
  if (err != nil) {
      t.Error(err)
  }

  // Check dirs exist
  for i := range logDirs {
    checkDirExists(t,path.Join(baseDir,hostDir,logDirs[i]))
  }

  // Test re-run of SetupLoggingDirs
  err = SetupLoggingDirs(baseDir,hostDir,logDirs)
  if (err != nil) {
    t.Error(err)
  }

  // Check dirs still exist!
  for i := range logDirs {
    checkDirExists(t,path.Join(baseDir,hostDir,logDirs[i]))
  }

  // Tidy up (remove dirs)
  for i := range logDirs {
    os.Remove(path.Join(baseDir,hostDir,logDirs[i]))
    os.Remove(path.Join(baseDir,hostDir))
  }

  // TODO
  // Create a file with the same name as one of the directories
  // Check this returns an appropriate error
  // Tidy up (remove dirs)
}
