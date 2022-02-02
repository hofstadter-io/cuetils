tasks: {
  @flow(test)
  words: """
  hello world
  hallo chat!
  """

  t0: {  
    @task(os.WriteFile)
    filename: "writefile_001.stdout"
    contents: words
    mode: 0o666
  }
}
