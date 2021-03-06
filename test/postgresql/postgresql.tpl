{{define "selectTestTable"}}
{{$COLUMNS := "id, username, password"}}
SELECT {{$COLUMNS}} FROM test_table
{{where (ne .Username "") "AND" "\"username\" =" .Username "" | where (ne .Password "pw") "AND" "\"password\"=" .Password}}
{{end}}

{{define "insertTestTable"}}
{{$COLUMNS := "id, username, password"}}
INSERT INTO test_table ({{$COLUMNS}})
  VALUES(
  {{.Id}},
  '{{.Username}}',
  '{{.Password}}'
  )
{{end}}

{{define "insertBatchTestTable"}}
{{$size := len . | add -1}}
INSERT INTO "test_table"("id","username","password")
VALUES {{range $i, $v := .}}
({{$v.Id}}, '{{$v.Username}}', '{{$v.Password}}'){{if lt $i $size}},{{end}}
{{end}}
{{end}}

{{define "updateTestTable"}}
UPDATE test_table
{{set (ne .Username "") "\"username\" =" .Username "" | set (ne .Password "") "\"password\" =" .Password}}
{{where (ne .Id 0) "AND" "\"id\"=" .Id ""}}
{{end}}

{{define "deleteTestTable"}}
DELETE FROM test_table
{{where (ne .Id 0) "AND" "id" .Id "" | where (ne .Username "") "AND" "\"username\"=" .Username | where (ne .Password "") "AND" "\"password\"=" .Password}}
{{end}}
