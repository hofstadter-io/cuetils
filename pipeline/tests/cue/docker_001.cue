
code: #"""
/usr/bin/env python3

print("hallo chat!")
"""#

py: {
  @pipeline()
  f: {
    filename: "in.py"
    contents: code
  }
  r: {
    @task(os.Exec)
    cmds: [
      "base",
      "-c"
    ]
  }
}
