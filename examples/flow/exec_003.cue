@pipeline()

run: {
  @task(os.Exec)
  cmd: ["cue", "eval", "-"]
  stdin: "{"
  // stderr: null
  stderr: string
}

err: {
  @task(os.Stdout)
  text: run.stderr
}

code: {
  @task(os.Stdout)
  code: run.exitcode
  text: "exit: \(code)\n" 
}
