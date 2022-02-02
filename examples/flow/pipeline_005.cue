p1: {
  @flow(p1)
  t: { text: "p1" } @task(os.Stdout)
}

c1: {
  @flow(c1)
  t: { text: "c1" } @task(os.Stdout)
  p2: {
    @flow(p2,pN)
    t: { text: "p2" } @task(os.Stdout)
  }
  p3: {
    @flow(p3,pN)
    t: { text: "p3" } @task(os.Stdout)
  }
}

