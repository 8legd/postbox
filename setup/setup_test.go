package setup

import (
  "testing"
  "os"
  "path"
  "bufio"
  "errors"
  "strings"
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

func createTestFile(t *testing.T, path string) {
  f, err := os.Create(path)
  defer f.Close()
  if err != nil {
    t.Error(err)
  } else {
    w := bufio.NewWriter(f)
    _, err := w.WriteString("testing, testing, 123\n")
    if err != nil {
      t.Error(err)
    }
    w.Flush()
  }
}

func removeTestFile(t *testing.T, path string) {
  f, err := os.Stat(path)
  if (err != nil) {
    t.Error(err)
  } else {
    if !f.IsDir() {
      t.Error(errors.New(path + " exists but is not a directory"))
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

  // Create a file with the same name as the hostDir
  createTestFile(t,path.Join(baseDir,hostDir))

  // Check this returns an appropriate error
  err = SetupLoggingDirs(baseDir,hostDir,logDirs)
  if err != nil {
    if !strings.Contains(err.Error(),"exists but is not a directory") {
      t.Error(err)
    }
  } else {
    t.Error(errors.New("Expected error because " + path.Join(baseDir,hostDir) + " should exist and not be a directory"))
  }

  // Tidy up (remove file)
  os.Remove(path.Join(baseDir,hostDir))
}
