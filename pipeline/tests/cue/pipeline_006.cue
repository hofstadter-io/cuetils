p1: {
  @pipeline(p1)
  t: { text: "p1" } @task(os.Stdout)
}

p2: {
  @pipeline(p2)
  t: { text: "p2" } @task(os.Stdout)
  p: p1
}
