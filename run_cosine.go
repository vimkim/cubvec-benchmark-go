package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Command-line flags
	var (
		db    = flag.String("database", "", "Database name")
		dim   = flag.Int("dim", 256, "Vector dimension")
		rows  = flag.Int("rows", 75000, "Number of rows")
		table = flag.String("table", "", "Table name")
		user  = flag.String("u", "", "Database user")
	)
	flag.Parse()

	if *db == "" || *table == "" || *user == "" {
		fmt.Println("Missing required arguments. Usage: --database <db> --table <table> -u <user>")
		os.Exit(1)
	}

	// Generate SQL content
	sqlFilename := fmt.Sprintf("test_%d_%d.sql", *dim, *rows)
	sqlContent := fmt.Sprintf(`SELECT @session_var := vec FROM %s LIMIT 1;
SET TRACE ON;
SELECT count(*) from (SELECT /*+ no_merge */ COSINE_DISTANCE(@session_var, vec) FROM %s);
SHOW TRACE;`, *table, *table)

	// Write SQL to file
	err := os.WriteFile(sqlFilename, []byte(sqlContent), 0644)
	if err != nil {
		fmt.Printf("Error writing SQL file: %v\n", err)
		os.Exit(1)
	}

	// Build csql command
	cmd := exec.Command("csql", "-u", *user, *db, "-i", sqlFilename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run command
	fmt.Printf("Running: csql -u %s %s -i %s\n", *user, *db, sqlFilename)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running csql: %v\n", err)
		os.Exit(1)
	}
}
