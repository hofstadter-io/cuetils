import (
  "github.com/hofstadter-io/cuetils/flow/tasks/os"
)

tasks: {
  @flow(test)
  words: """
  hello world
  hallo chat!
  """

  t0: os.WriteFile & {  
    filename: "writefile_002.stdout"
    contents: words
  }
}
