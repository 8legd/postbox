package storage

import (
  "testing"
  "io"
  "os"
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


func TestSetupStdoutWriter(t *testing.T) {
  w, _ := SetupStdoutWriter()
  _, err := io.WriteString(w,"")
  if (err != nil) {
    t.Error(err)
  }
}

func TestSetupFileSystemWriter(t *testing.T) {

  dir := "postbox"
  prefix := "postbox"

  // Test succesful creation
  _, err := SetupFileSystemWriter(dir,prefix)
  if (err != nil) {
      t.Error(err)
  }

  // Check exist
  checkDirExists(t,dir)

  // Test re-run of SetupFileSystem
  _, err = SetupFileSystemWriter(dir,prefix)
  if (err != nil) {
    t.Error(err)
  }

  // Check still exist!
  checkDirExists(t,dir)

  // Tidy up (remove dir)
  os.Remove(dir)

  // Create a file with the same name as the dir
  createTestFile(t,dir)

  // Check this returns an appropriate error
  _, err = SetupFileSystemWriter(dir,prefix)
  if err != nil {
    if !strings.Contains(err.Error(),"exists but is not a directory") {
      t.Error(err)
    }
  } else {
    t.Error(errors.New("Expected error because " + dir + " should exist and not be a directory"))
  }

  // Tidy up (remove file)
  os.Remove(dir)


  // Finally try and actually write to the file system!



}
