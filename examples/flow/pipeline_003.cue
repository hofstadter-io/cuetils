p0: {
  @flow()
  t: { text: "p0" } @task(os.Stdout)
}

p1: {
  @flow(p1)
  t: { text: "p1" } @task(os.Stdout)
}

nested: {
  p2: {
    @flow(p2,pN)
    t: { text: "p2" } @task(os.Stdout)
  }

  p3: {
    @flow(p3,pN)
    t: { text: "p3" } @task(os.Stdout)
  }
}

