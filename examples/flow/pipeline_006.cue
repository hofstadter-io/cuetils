p1: {
  @flow(p1)
  t: { text: "p1" } @task(os.Stdout)
}

p2: {
  @flow(p2)
  t: { text: "p2" } @task(os.Stdout)
  p: p1
}
