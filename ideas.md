# Todo

- [ ] merge and release

Enhance & expand

- [ ] list tags / secrets
- [ ] obfescate secrets, centralized printing (ensure that is the case in every task / run)
- [ ] locks for tasks
- [ ] more task types
    - [ ] file append
    - [ ] mkdir
    - [ ] chan / mailbox
    - [ ] memory store / load
    - [ ] prevent exit when error in handler pipelines


Rename:
- rename pipeline -> run
- hof/flow cmd

Docs...

---

Build other things cuetils/run

- save all IRC messages to DB
- bookmarks and hn upvotes
- change my lights
- replace helm (need native topo sort)
- OAuth workflow

---

Bookkeeping:
- debug flag to print tasks which run
- stats for tasks & pipelines, chan to central

Exec & Serve

- some way to run in background, and then kill / exit later?
- write to file for stdio

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


### Memory issues

https://github.com/lipence/cue/commit/6ed69100ebd62509577826657536172ab46cf257

### cue/flow

return final value: https://github.com/cue-lang/cue/pull/1390


### streamer

- other social networks / interaction systems
