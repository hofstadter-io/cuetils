@pipeline()

run: {
  @task(os.Exec)
  cmd: ["bash", "-c", "echo $FOO"]
  dir: "../"
  env: {
    FOO: "BAR"
  }
  stdout: string
}

print: {
  @task(os.Stdout)
  text: run.stdout
}
