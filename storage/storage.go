// Storage is a package for reading and writing to various storage locations.
package storage

import (
  "io"
  "io/ioutil"
  "path/filepath"
  "os"
  "time"
  "path"
  "errors"
)

// Store provides a standard interface for the different storage locations.
type Store interface {
  Write(content string) (id string, err error)
  List() (ids []string, err error)
  Read(id string) (content string, err error)
}


type FileSystemStore struct {
  dir string
}
func (fss *FileSystemStore) Write(content string) (id string, err error) {
  prefix := time.Now().Format("20060102150405.000000000") + ".";
  f, err := ioutil.TempFile(fss.dir, prefix)
  defer f.Close()
  if err != nil {
    return "", err
  }
  n, err := io.WriteString(f, content)
  if err != nil {
    return "", err
  } else {
    if n != len(content) {
      return "", errors.New("failed to write content to FileSystemStore")
    }
  }
  return filepath.Base(f.Name()), nil
}
func (fss *FileSystemStore) List() (ids []string, err error) {
  f, err := os.Open(fss.dir)
  defer f.Close()
  fis, err := f.Readdir(0)
  for _, v := range fis {
    n := v.Name()
    ids = append(ids, n)
  }
  return ids, err
}
func (fss *FileSystemStore) Read(id string) (content string, err error) {
  data, err := ioutil.ReadFile(path.Join(fss.dir,id))
  return string(data), err
}


// SetupFileSystemWriter checks the specified directory exists and if not tries to create it.
// It returns a FileSystem Store which will write to temp files in the specified directory.
func SetupFileSystem(dir string) (Store, error) {
  fi, err := os.Stat(dir)
  if err != nil && os.IsNotExist(err) {
    err = os.Mkdir(dir, 0777)
    if err != nil {
      return nil, errors.New("failed to create directory " + dir + ". check parent directory exists and has correct permissions")
    } else {
      fi, err = os.Stat(dir)
    }
  }
  if err != nil {
    return nil, err
  }
  if !fi.IsDir() {
    return nil, errors.New(dir + " exists but is not a directory")
  } else {
    fss := FileSystemStore{dir}
    return &fss, nil
  }
}
