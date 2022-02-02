tasks: {
  @flow(readfile)
	r: { filename: "stdata/readfile_001.txt", contents: string } @task(os.ReadFile)
}
