p0: {
  @pipeline()
  t: { text: "p0" } @task(os.Stdout)
}

p1: {
  @pipeline(p1)
  t: { text: "p1" } @task(os.Stdout)
}

nested: {
  p2: {
    @pipeline(p2,pN)
    t: { text: "p2" } @task(os.Stdout)
  }

  p3: {
    @pipeline(p3,pN)
    t: { text: "p3" } @task(os.Stdout)
  }
}

