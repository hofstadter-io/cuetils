
code: #"""
/usr/bin/env python3

print("hallo chat!")
"""#

py: {
  @pipeline()
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
