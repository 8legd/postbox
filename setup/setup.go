package setup

import "os"
import "errors"

func setupDir(dir string) error {
  var err error
  f, err := os.Stat(dir)
  if (err != nil) {
    if (os.IsNotExist(err)) {
      err = os.Mkdir(dir, 0777)
    }
  } else {
    if !f.IsDir() {
      err = errors.New(dir + " exists but is not a directory")
    }
  }
  return err
}


func SetupLoggingDirs(baseDir string, logDirs []string) error {
  var err error
  err = setupDir(baseDir)
  if err != nil {
    // TODO loop through logDirs setting them up too
  }
  return err
}
