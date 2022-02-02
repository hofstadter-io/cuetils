# Todo

Enhance & expand

- [x] list Tags
- [x] list secrets
- [x] locks for tasks
- [ ] more task types
    - [x] file append
    - [x] mkdir
    - [x] memory store / load
- [ ] rename pipeline -> run
- [ ] merge and release
- [ ] hof/flow cmd

### Docs...

probably hof/docs


### Other task types:

- gen
  - uuid
  - rand int
  - rand str

- time
  - now

- async / client listener
  - waitgroup / mutex?
  - file locks (https://github.com/gofrs/flock)
  - [ ] chan / mailbox
  - websockets

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

Exec & Serve & async

- [ ] some way to run in background, and then kill / exit later?
- [ ] write directly to file for stdio, is it a concrete string?
- [ ] something like a goroutine, similar to api.Serve/pipeline
- [ ] message passing, via chans, websockets, kafka/rabbit

Bugs?

- [ ] prevent exit when error in handler pipelines?
- [ ] rare & racey structural cycle

Helpers:

- canonicalize (sort fields recursively)
- toposort

List processing:

- jsonl
- yaml with `---`
- CUE got streaming json/yaml support
- if extracted value is a list?

### Other

Go funcs:

- rename currenct to `*Globs`
- pure Go implementations
- funcs that take values

CLI:

- Support expression on globs, to select out a field on each file
- move implementation?

## upstream & references

Memory issues

https://github.com/lipence/cue/commit/6ed69100ebd62509577826657536172ab46cf257

### cue/flow

return final value: https://github.com/cue-lang/cue/pull/1390


### streamer

- other social networks / interaction systems
