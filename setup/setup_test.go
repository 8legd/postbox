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
  // TODO
  // Check dirs exist
  // Test re-run of SetupLoggingDirs
  // Check dirs still exist!
  // Tidy up (remove dirs)
  // Create a file with the same name as one of the directories
  // Check this returns an appropriate error
  // Tidy up (remove dirs)
}
