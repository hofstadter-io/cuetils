@pipeline()

set: {
  @task(kv.Mem)
  key: "foo"
  val: "bar"
}

get: {

  wait: {
    @task(os.Sleep)
    duration: "1s"
  }

  load: {
    @task(kv.Mem)
    dep: wait.done
    key: "foo"
  }

  print: {
    @task(os.Stdout)
    text: load.val
  }
}
