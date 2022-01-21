p1: {
  @pipeline(p1)
  t: { #O: "p1" } @task(os/stdout)
}

p2: {
  @pipeline(p2,pN)
  t: { #O: "p2" } @task(os/stdout)
}

p3: {
  @pipeline(p3,pN)
  t: { #O: "p3" } @task(os/stdout)
}

