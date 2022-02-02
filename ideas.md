# Todo

Enhance & expand

- [x] list Tags
- [x] list secrets
- [x] locks for tasks
- [ ] more task types
    - [x] file append
    - [x] mkdir
    - [ ] memory store / load
    - [ ] prevent exit when error in handler pipelines?
- [ ] rename pipeline -> run
- [ ] merge and release
- [ ] hof/flow cmd

### Docs...

probably hof/docs


### Other task types:

- async
  - [ ] chan / mailbox

- msg
  - rabbitmq
  - kafka
  - nats
- k/v
  - redis
  - mongo
  - s3/gcs
  - vault
- command line prompt

### Build other things cuetils/run

- save all IRC messages to DB
- bookmarks and hn upvotes
- change my lights
- replace helm (need native topo sort)
- OAuth workflow

### More todo, always...

i/o centralization

- [ ] debug/verbose flag to print tasks which run
- [ ] stats for tasks & pipelines, chan to central
- [ ] obfescate secrets, centralized printing (ensure that is the case in every task / run)

Exec & Serve

- [ ] some way to run in background, and then kill / exit later?
- [ ] write directly to file for stdio, is it a concrete string?

---

Then...


More...

- something like a goroutine, similar to api.Serve/pipeline
- message passing, via chans, websockets, kafka/rabbit

# Ideas

Example Pipeline:
- Exponential retry as an api.Call wrapper with os.Sleep
- replace sleep in example with wait for server ready
- could be tricky, because we need to generate extra tasks after the last one finished, or some conditional to ignore after success
- api req timeout
- retry status codes / message

Helpers:

- canonicalize (sort fields recursively)
- toposort

List processing:

- jsonl
- yaml with `---`
- CUE got streaming json/yaml support
- if extracted value is a list?

Go funcs:

- rename currenct to `*Globs`
- pure Go implementations
- funcs that take values

CLI:

- Support expression on globs, to select out a field on each file
- move implementation?

### Memory issues

https://github.com/lipence/cue/commit/6ed69100ebd62509577826657536172ab46cf257

### cue/flow

return final value: https://github.com/cue-lang/cue/pull/1390


### streamer

- other social networks / interaction systems
