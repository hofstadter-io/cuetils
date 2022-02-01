# TODO

Docs...

rename pipeline -> run,dag

---

Build other things cuetils/run

- bookmarks and hn upvotes

---

- auth / oauth
    - [x] twitch  (oauth code)
- [ ] api calls
    - [x] update stream title 
    - [x] twitch (other)
    - [ ] update go live message
    - [ ] which of my streamers are live
- [x] chat
    - [x] twitch-irc 

- [ ] memory store / load
- [ ] prevent exit when error in handler pipelines
- [ ] api.GraphQL but probably better as a wrapper around api.Call
---

Get server working:
- [x] wired up routes
- [x] route pipelines

CLI work:
- [x] list pipelines that can be run
- [x] enable docs for pipelines to be read/written
- [x] know what inputs / outputs of a pipeline are
- revisit tags and get them working
    - also print
    - find and return path, unify at top?

--- 

Additional examples:
- auth / oauth
    - [ ] twitter
    - [x] youtube (apikey)
- [ ] api calls
    - [ ] twitter
    - [ ] youtube
- [ ] chat
    - [ ] slack
    - [ ] discord

Additional Tasks:
- [ ] os.Getwd

Centralized Printing:
- chan for tasks to write strings to

Secrets:
- `secret: [string]: string` as secrets to be elided from output
- add as a filter to centralized printing

Bookkeeping:
- debug flag to print tasks which run
- stats for tasks & pipelines, chan to central

Exec & Serve

- some way to run in background, and then kill / exit later?
- write to file for stdio

---

Then...

- OAuth workflow
- Twitchbot

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

- extend (add if not present)
- canonicalize (sort fields recursively)

List processing:

- jsonl
- yaml with `---`
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
