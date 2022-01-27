package utils

import "strings"

RepoRoot: { 
  @task(os.Exec)
  cmd: ["bash", "-c", "git rev-parse --show-toplevel"]
  stdout: string
  Out: strings.TrimSpace(stdout)
}
