// Package setup provides functions for setting up the logging directories
package setup

import (
  "os"
  "path"
  "errors"
)

func setupDir(dir string) error {
  f, err := os.Stat(dir)
  if err != nil {
    if os.IsNotExist(err) {
      err = os.Mkdir(dir, 0777)
      if os.IsNotExist(err) {
        return errors.New("failed to create directory " + dir + " check parent directories exist and have correct permissions")
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

// SetupLoggingDirs checks the specified directories exist and if not tries to create them in the baseDir
func SetupLoggingDirs(baseDir string, hostDir string, logDirs []string) error {
  err := setupDir(path.Join(baseDir,hostDir))
  if err != nil {
    return err
  } else {
    for i := range logDirs {
      err = setupDir(path.Join(baseDir,hostDir,logDirs[i]))
      if err != nil {
        return err
      }
    }
    return nil
  }
}
