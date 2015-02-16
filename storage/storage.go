// Storage is a package for reading and writing to various storage locations
package storage

import (
  "io"
  "io/ioutil"
  "os"
  "errors"
)


// SetupStdoutWriter returns the standard output implementation of the io.Writer interface.
func SetupStdoutWriter() (io.Writer, error) {
  return os.Stdout, nil
}


type fileSystemStore struct {
  dir, prefix string
}
func (fss *fileSystemStore) Write(p []byte) (n int, err error) {
  f, err := ioutil.TempFile(fss.dir, fss.prefix)
  if err != nil {
    return 0, err
  }
  defer f.Close()
  return f.Write(p)
}

// SetupFileSystemWriter checks the specified directory exists and if not tries to create it.
// It returns an io.Writer which will write to a temp file in the specified directory.
func SetupFileSystemWriter(dir, prefix string) (io.Writer, error) {
  fi, err := os.Stat(dir)
  if err != nil {
    if os.IsNotExist(err) {
      err = os.Mkdir(dir, 0777)
      if os.IsNotExist(err) {
        return nil, errors.New("failed to create directory " + dir + ". check parent directory exists and has correct permissions")
      } else {
        return nil, err
      }
    } else {
      return nil, err
    }
  } else {
    if !fi.IsDir() {
      return nil, errors.New(dir + " exists but is not a directory")
    } else {
      fs := fileSystemStore{dir, prefix}
      return &fs, nil
    }
  }
}
