tasks: {
  @pipeline(readfile)
	r: { f: "stdata/readfile_001.txt", contents: string } @task(os.ReadFile)
}
