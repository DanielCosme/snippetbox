package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/danielcosme/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

// The responsibilities of main() function are limited to:
//    - Parsing the runtime configuration settings for the application;
//    - Establishing the dependencies for the handlers; and
//    - Running the HTTP server.
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// SQL Data Source Name (dsn)
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
