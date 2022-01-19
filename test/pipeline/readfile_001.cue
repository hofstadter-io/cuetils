tasks: {
  @pipeline(readfile)
	r: { #F: "readfile_001.txt", Contents: string } @task(os/readfile) @print(Contents)
}
