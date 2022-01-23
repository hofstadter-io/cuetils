import "encoding/json"

tasks: {
  @pipeline(readfile)
	r: { #F: "../tree.json", Contents: string } @task(os/readfile)
  j: json.Unmarshal(r.Contents)
  p: { #X: j, #P: { tree: cow: _ } } @task(st/pick)

  final: { #O: p.Out.tree } @task(os/stdout)
}


