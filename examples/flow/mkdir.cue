@flow()

mkdir: {
  @task(os.Mkdir)
  dir: "foo/bar"
}
