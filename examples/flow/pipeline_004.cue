p1: {
  @pipeline(p1)
  t: { text: "p1" } @task(os.Stdout)
}

// c1 docs string beforehand
// a comment outside of the struct
c1: {
  @pipeline(c1)
  p2: {
    @pipeline(p2,pN)
    t: { text: "p2\n" } @task(os.Stdout)
  }

  p3: {
    @pipeline(p3,pN)
    t: { text: "p3\n" } @task(os.Stdout)
  }
}

