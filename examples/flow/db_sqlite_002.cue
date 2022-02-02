@flow()

tQ: #"""
CREATE TABLE `userinfo` (
  `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
  `username` VARCHAR(64) NULL,
  `departname` VARCHAR(64) NULL,
  `created` DATE NULL
);
"""#

iQ: "INSERT INTO userinfo(username, departname, created) values(?,?,?)"

prep: {
  @task(db.Call)
  conn: sqlite: "test.db"

  stmts: [
    { exec: tQ },
    { exec: iQ, args: ["astaxie", "研发部门", "2012-12-09"] },
    { query: "SELECT * FROM userinfo;" },
  ]

  results: _
  result: results[2].results
}
results: {
  @task(os.Stdout)
  text: prep.result
}
