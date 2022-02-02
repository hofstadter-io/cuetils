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

create: {
  @task(db.Call)
  conn: sqlite: "test.db"
  exec: tQ 
}

insert: {
  @task(db.Call)
  _dep: create.results

  conn: sqlite: "test.db"
  exec: iQ 
  args: ["astaxie", "研发部门", "2012-12-09"]
}

query: {
  @task(db.Call)
  _dep: insert.results

  conn: sqlite: "test.db"
  query: "SELECT * FROM userinfo;"
}

results: {
  @task(os.Stdout)
  text: query.results
}
