@pipeline()
append: {
  @task(os.WriteFile)
  filename: "append.txt"
  contents: "Hallo Chat!\n"
  mode: 0o644
  append: true
}
