@pipeline()

lock: {
  @task(os.FileLock)
  filename: "foo.lock"
}

after: {

  wait: {
    @task(os.Sleep)
    duration: "1s"
  }

  unlock: {
    dep: wait.done
    @task(os.FileUnlock)
    filename: "foo.lock"
  }
}

