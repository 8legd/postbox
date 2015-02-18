package storage

import (
  "testing"
  "os"
  "bufio"
  "errors"
  "strings"
  "strconv"
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


func TestSetupFileSystem(t *testing.T) {

  dir := "postbox"

  // Test succesful creation
  fs, err := SetupFileSystem(dir)
  if (err != nil) {
      t.Error(err)
  }

  // Check exists
  checkDirExists(t,dir)

  // Test re-run of SetupFileSystem
  fs, err = SetupFileSystem(dir)
  if (err != nil) {
    t.Error(err)
  }

  // Check still exists!
  checkDirExists(t,dir)

  // Tidy up (remove dir)
  os.Remove(dir)

  // Create a file with the same name as the dir
  createTestFile(t,dir)

  // Check this returns an appropriate error
  fs, err = SetupFileSystem(dir)
  if err != nil {
    if !strings.Contains(err.Error(),"exists but is not a directory") {
      t.Error(err)
    }
  } else {
    t.Error(errors.New("Expected error because " + dir + " should exist and not be a directory"))
  }

  // Tidy up (remove file)
  err = os.Remove(dir)
  if (err != nil) {
    t.Error(err)
  }

  // Finally try and actually write to the store!
  fs, err = SetupFileSystem(dir)
  if (err != nil) {
    t.Error(err)
  }
  _, err = fs.Write("testing 1")
  if (err != nil) {
    t.Error(err)
  }
  _, err = fs.Write("testing 2")
  if (err != nil) {
    t.Error(err)
  }
  _, err = fs.Write("testing 3")
  if (err != nil) {
    t.Error(err)
  }

  ids, err := fs.List()
  if (err != nil) {
    t.Error(err)
  }
  for i, v := range ids {
    content, err := fs.Read(v)
    if (err != nil) {
      t.Error(err)
    } else {
      if content != ("testing " + strconv.Itoa(i+1)) {
        t.Error("read content does not match written content for id:" + v)
      }
    }
  }

  // Tidy up (remove dir and contents)
  err = os.RemoveAll(dir)
  if (err != nil) {
    t.Error(err)
  }

}
