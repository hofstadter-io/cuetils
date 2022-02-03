import "encoding/json"

@flow()

// seed with int value
dummy: { done: _ } @task(gen.Seed)

gen: {
  t: { _, #dep: dummy.done } @task(gen.Now)
  i: { _, #dep: dummy.done } @task(gen.Int)
  s: { _, #dep: dummy.done } @task(gen.Str)
  f: { _, #dep: dummy.done } @task(gen.Float)
  n: { _, #dep: dummy.done } @task(gen.Norm)
  u: { _, #dep: dummy.done } @task(gen.UUID)
  c: { _, #dep: dummy.done } @task(gen.CUID)
  g: { _, #dep: dummy.done } @task(gen.Slug)
}

print: { 
  @task(os.Stdout)
  dep: { for k,v in gen { (k): v } } 
  s: json.Indent(json.Marshal(dep), "", "  ")
  text: s + "\n"
}
