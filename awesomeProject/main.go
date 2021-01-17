package main
import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	addr := flag.String("addr", ":5432", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting  server on %v", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		panic(err)
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil { errorLog.Fatal(err)
	}

	defer db.Close()
	app := &application{ errorLog: errorLog, infoLog: infoLog,
	}
	srv := &http.Server{ Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(), }

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)


}

func openDB(dsn string) (*sql.DB, error) { db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err }
	if err = db.Ping(); err != nil { return nil, err
	}
	return db, nil }
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) { id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w) } else {
			app.serverError(w, err) }
		return
	}
	fmt.Fprintf(w, "%v", s) }
