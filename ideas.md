# Today

First:

- [ ] other
  - [ ] pipe stdin/out/err to files
  - [ ] some way to run in background, and then kill / exit later?

- [ ] async / client listener
  - [ ] kill chan, also need to catch signals in main(?) and pass down / do right thing
    - [ ] how to tell (server / bg exec'd process) to stop (oauth localhost as example)
  - [ ] chan / mailbox
  - [ ] waitgroup / mutex?
  - [ ] websockets

--- 

Then: 

- [ ] sql
  - [x] sqlite
  - [ ] postgres
  - [ ] mysql

- [ ] msg
  - [ ] rabbitmq
  - [ ] kafka
  - [ ] nats

- [ ] k/v
  - [ ] redis
  - [ ] memcache
  - [ ] gcs
  - [ ] s3

- [ ] obj
  - [ ] elasticsearch
  - [ ] mongo


---

at some point:

- [ ] hof/flow cmd

### Docs...

probably hof/docs


### Other task types:

- temp files / dirs
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
