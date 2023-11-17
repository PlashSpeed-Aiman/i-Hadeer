package main

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/r3labs/sse/v2"
	"log"
	"net/http"
	"time"
)

type Student struct {
	Name         string `json:"name"`
	MatricNumber string `json:"matricNumber"`
}

func sendPing(server *sse.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Got trigger request. Sending SSE")
		go func() {
			time.Sleep(1 * time.Second)
			server.Publish("messages", &sse.Event{
				Data: []byte(time.Now().String()),
			})
		}()
		w.WriteHeader(200)
	}
}
func sendPingForReal(server *sse.Server) {
	go func() {
		for {
			log.Printf("Sending SSE")
			time.Sleep(1 * time.Second)
			server.Publish("messages", &sse.Event{
				Data: []byte(time.Now().String()),
			})
		}
	}()
}
func main() {
	server := sse.New()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/dist/assets")))
	fs2 := http.FileServer(http.Dir("./web/dist"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	server.CreateStream("messages")
	server.CreateStream("students")
	/*	go sendPingForReal(server)
	 */
	r.Handle("/", fs2)
	r.Handle("/assets/*", fs)

	r.Get("/events", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			// Received Browser Disconnection
			<-r.Context().Done()
			println("The client is disconnected here")
			return
		}()

		server.ServeHTTP(w, r)
	}) // Register the server as an http.Handler
	r.Get("/ping", sendPing(server))
	r.Post("/attendance", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		db, err := sql.Open("sqlite3", "./attendance.db?_busy_timeout=5000")
		var student Student
		json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			log.Printf(err.Error())
		}
		defer db.Close()
		//date time format
		_, err = db.Exec("insert into attendance (Name, MatricNumber, Time) values (?, ?,?)", student.Name, student.MatricNumber, currentTime)
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf("Got attendance request. Sending SSE")
		if err != nil {
			log.Printf(err.Error())
		}
		var x = struct {
			Name         string `json:"name"`
			MatricNumber string `json:"matricNumber"`
			Time         string `json:"time"`
		}{
			Name:         student.Name,
			MatricNumber: student.MatricNumber,
			Time:         currentTime,
		}

		res, err := json.Marshal(x)
		server.Publish("students", &sse.Event{
			Data: res,
		})
		w.WriteHeader(200)
	})
	log.Default().Println("Server started at localhost:8080")
	http.ListenAndServe("localhost:8080", r)

}
