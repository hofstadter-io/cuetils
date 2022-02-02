@flow()

beg: {
  @task(os.FileLock)
  filename: "foo.lock"
}

mid: {

  wait: {
    @task(os.Sleep)
    duration: "2s"
  }

  unlock: {
    dep: wait.done
    @task(os.FileUnlock)
    filename: "foo.lock"
  }
}

end: {
  wait: {
    @task(os.Sleep)
    duration: "1s"
  }

  lock: {
    dep: wait.done
    @task(os.FileLock)
    filename: "foo.lock"
    retry: "500ms"
  }

  unlock: {
    dep: lock.done
    @task(os.FileUnlock)
    filename: "foo.lock"
  }
}
