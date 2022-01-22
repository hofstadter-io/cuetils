test1: {
  @pipeline()
  t: { FOO: _ } @task(os.Getenv)
  o: { text: t.FOO + "\n" } @task(os.Stdout)
}
