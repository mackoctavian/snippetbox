package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.mackoctavian.com/internal/models"
)

type application struct {
	errLog   *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "root:DatastreamingCowboy2022@/snippetbox?parseTime=true", "mysql data source")
	flag.Parse()

	//Error and info loging using log.New()
	infoLog := log.New(os.Stdout, "INFO/", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR/", log.Ldate|log.Lshortfile|log.Ltime)

	//Opening connection to the database
	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}

	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	app := &application{
		errLog:   errLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	// Using the http.ListenAndServe() function to start a new web server.Passing in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux  created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	infoLog.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}

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
