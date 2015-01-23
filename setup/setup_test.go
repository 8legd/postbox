package setup

import "testing"

func TestSetupLoggingDirs(t *testing.T) {
  // Test succesful creation
  baseDir := "logs"
  logDirs := []string{"sessions","users","errors"}
  err := SetupLoggingDirs(baseDir,logDirs)
  if (err != nil) {
      t.Error(err)
  }
  // TODO Tidy up (remove dirs)
}
