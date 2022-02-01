p1: {
  @pipeline(p1)
  t: { text: "p1" } @task(os.Stdout)
}

c1: {
  @pipeline(c1)
  t: { text: "c1" } @task(os.Stdout)
  p2: {
    @pipeline(p2,pN)
    t: { text: "p2" } @task(os.Stdout)
  }
  p3: {
    @pipeline(p3,pN)
    t: { text: "p3" } @task(os.Stdout)
  }
}

