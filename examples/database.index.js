const db = sql.open("sqlite", "test.sql");
db.exec("create table if not exists test (id integer primary key, name text)");
console.log(db.exec("insert into test (name) values (?)", "foo").lastInsertId());
db.exec("insert into test (name) values (?)", "bar");

const result = db.query("select * from test");
for (const row of result) {
  console.log(JSON.stringify(row));
}
