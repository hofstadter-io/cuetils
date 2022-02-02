test1: {
  @flow()
  t: { FOO: _ } @task(os.Getenv)
  o: { text: t.FOO + "\n" } @task(os.Stdout)
}
