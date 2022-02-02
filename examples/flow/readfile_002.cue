import "encoding/json"

tasks: {
  @flow(readfile)
	r: { filename: "req.json", contents: string } @task(os.ReadFile)
  j: json.Unmarshal(r.contents)
  final: { out: p.out.tree } @task(os.Stdout)
}


