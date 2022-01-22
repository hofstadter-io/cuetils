p1: {
  @pipeline(p1)
  t: { #O: "p1" } @task(os/stdout)
}

p2: {
  @pipeline(p2)
  t: { #O: "p2" } @task(os/stdout)
  p: p1
}
