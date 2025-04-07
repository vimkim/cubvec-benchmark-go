run-256-75000:
    go run run_cosine.go --database testdb --dim 256 --rows 75000 --table tbl_256_75000 -u dba

run-256-150000:
    go run run_cosine.go --database testdb --dim 256 --rows 150000 --table tbl_256_150000 -u dba

run-256-300000:
    go run run_cosine.go --database testdb --dim 256 --rows 300000 --table tbl_256_300000 -u dba

run-768-75000:
    go run run_cosine.go --database testdb --dim 768 --rows 75000 --table tbl_768_75000 -u dba

run-768-150000:
    go run run_cosine.go --database testdb --dim 768 --rows 150000 --table tbl_768_150000 -u dba

run-768-300000:
    go run run_cosine.go --database testdb --dim 768 --rows 300000 --table tbl_768_300000 -u dba

run-1536-75000:
    go run run_cosine.go --database testdb --dim 1536 --rows 75000 --table tbl_1536_75000 -u dba

run-1536-150000:
    go run run_cosine.go --database testdb --dim 1536 --rows 150000 --table tbl_1536_150000 -u dba

run-1536-300000:
    go run run_cosine.go --database testdb --dim 1536 --rows 300000 --table tbl_1536_300000 -u dba
