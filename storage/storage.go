// Package storage provides functions for working with various storage locations
package storage

import (
  "os"
  "errors"
)

// Function SetupFileSystem checks the specified directory exists and if not tries to create it
func SetupFileSystem(dir string) error {
  f, err := os.Stat(dir)
  if err != nil {
    if os.IsNotExist(err) {
      err = os.Mkdir(dir, 0777)
      if os.IsNotExist(err) {
        return errors.New("failed to create directory " + dir + ". check parent directory exists and has correct permissions")
      } else {
        return err
      }
    } else {
      return err
    }
  } else {
    if !f.IsDir() {
      return errors.New(dir + " exists but is not a directory")
    } else {
      return nil
    }
  }
}
