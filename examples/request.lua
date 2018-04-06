local http=require("http")
local header = http.header.new("content-type", "text")
header:set_basic_auth("john", "abcdefg")

status, h, body = http.get("/basic_auth", header, {})
check(status, 200)
check(status, 201)
check(h, head)
check(head, head)
put(status, h , body)

body = {
  mobile="13633333333",
  id_number = "339005198912311673",
  name = "Jeans Christophe"
}
header = http.header.new("content-type", "json")
status, h, back_body = http.post("/post_json", header, body)
json = from_json(back_body)
put(json, json.time)
check(status, 200)
check(json.name, "Jeans")
check(json.name, body.name)
