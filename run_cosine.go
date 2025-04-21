package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func main() {
	// Command-line flags
	var (
		db       = flag.String("database", "", "Database name")
		dim      = flag.Int("dim", 256, "Vector dimension")
		rows     = flag.Int("rows", 75000, "Number of rows")
		table    = flag.String("table", "", "Table name")
		user     = flag.String("u", "", "Database user")
		tmplname = flag.String("tmpl", "", "Database user")
	)
	flag.Parse()

	if *db == "" || *table == "" || *user == "" || *tmplname == "" {
		fmt.Println("Missing required arguments. Usage: --database <db> --table <table> -u <user> --tmpl <template>")
		os.Exit(1)
	}

	// Generate SQL content
	sqlFilename := fmt.Sprintf("test_%d_%d.sql", *dim, *rows)
	// 	sqlContent := fmt.Sprintf(`SELECT @session_var := vec FROM %s LIMIT 1;
	// SET TRACE ON;
	// SELECT /*+ recompile */ count(*) from (SELECT /*+ no_merge */ COSINE_DISTANCE(@session_var, vec) FROM %s);
	// SHOW TRACE;
	// SELECT /*+ recompile */ count(*) from (SELECT /*+ no_merge */ COSINE_DISTANCE(@session_var, vec) FROM %s);
	// SHOW TRACE;`, *table, *table, *table)
	//
	// // Step 1: Read the SQL template from file
	content, err := os.ReadFile(*tmplname)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	// Step 2: Parse the SQL as a text/template
	tmpl, err := template.New("cosine_distance").Parse(string(content))
	if err != nil {
		log.Fatalf("Failed to parse SQL template: %v", err)
	}

	// Step 3: Define substitution variables
	data := struct {
		Table string
	}{
		Table: *table,
	}

	// Step 4: Execute the template with the data
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalf("Failed to execute SQL template: %v", err)
	}

	// Step 5: Write rendered SQL to file
	if err := os.WriteFile(sqlFilename, buf.Bytes(), 0644); err != nil {
		log.Fatalf("Error writing SQL file: %v", err)
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
