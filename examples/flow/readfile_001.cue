tasks: {
  @pipeline(readfile)
	r: { filename: "stdata/readfile_001.txt", contents: string } @task(os.ReadFile)
}
