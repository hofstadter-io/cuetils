
code: #"""
/usr/bin/env python3

print("hallo chat!")
"""#

py: {
  @flow()
  f: {
    @task(os.WriteFile)
    filename: "in.py"
    contents: code
  }
  r: {
    @task(os.Exec)
    cmds: ["bash", "-c", "in.py"]
  }
}
