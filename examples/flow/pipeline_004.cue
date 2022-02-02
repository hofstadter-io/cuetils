p1: {
  @flow(p1)
  t: { text: "p1" } @task(os.Stdout)
}

// c1 docs string beforehand
// a comment outside of the struct
c1: {
  @flow(c1)
  p2: {
    @flow(p2,pN)
    t: { text: "p2\n" } @task(os.Stdout)
  }

  p3: {
    @flow(p3,pN)
    t: { text: "p3\n" } @task(os.Stdout)
  }
}

