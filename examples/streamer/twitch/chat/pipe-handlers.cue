package chat

import "strings"

pipeHandlers: {

  "!today": {
    @pipeline()
    get: {
      @task(os.Exec)
      cmd: ["date"]
      stdout: string  
    }

    resp: strings.Replace(get.stdout, "\n", "", -1)
  }

  "!docker": {
    @pipeline()

    args?: [...string]

    get: {
      @task(os.Exec)
      cmd: ["docker", "ps"]
      stdout: string  
    }

    resp: strings.Replace(get.stdout, "\n", "", -1)
  }

  "!k8s": {
    @pipeline()

    args?: [...string]

    get: {
      @task(os.Exec)
      cmd: ["kubectl", "get", "pods", "--all-namespaces"]
      stdout: string  
    }

    // chill: { duration: "4s" } @task(os.Sleep)

    lines: strings.Split(get.stdout, "\n")
    count: len(lines) - 1

    resp: "there are \(count) pods running in Dr Verm's cluster" 
  }

  "!so": {
    // @pipeline()

    args?: [...string]
    who: args[1] 

    get: {
      // call twitch api for info about the user
      // eventually also look up custom data in DB
    }
    // chill: { duration: "4s" } @task(os.Sleep)

    resp: "you're the best \(who)"

  }
}
