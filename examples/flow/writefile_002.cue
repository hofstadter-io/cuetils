import (
  "github.com/hofstadter-io/cuetils/pipeline/tasks/os"
)

tasks: {
  @pipeline(test)
  words: """
  hello world
  hallo chat!
  """

  t0: os.WriteFile & {  
    filename: "writefile_002.stdout"
    contents: words
  }
}
