@pipeline()

run: {
  @task(os.Exec)
  cmd: ["jq", "."]
  stdin: """
  { "foo": "bar" }
  """
  // stdout: string
  // stdout: null
}
