package main

import (
	"ChatApp/database"
	"ChatApp/routs"
	"flag"
	"github.com/gofiber/fiber/v2"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once sync.Once

	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	app := fiber.New()
	database.Connect()
	routs.Setup(app)

	var addr = flag.String("addr", ":8080", "Uygulamanın adresi")
	flag.Parse()

	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/register", &templateHandler{filename: "register.html"})
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/room", r)
	// get the room going
	go r.run()
	//start the web server
	log.Println("Web Sunucuları başlıyor", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
