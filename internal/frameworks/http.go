package frameworks

import (
	"authService/internal/handlers"
	"log"
	"net/http"
)

type Server struct {
	mux http.Handler
}

func NewServer() *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello World!")) })
	mux.HandleFunc("/auth/get-token", handlers.AuthHandler)
	mux.HandleFunc("/auth/refresh", handlers.RefreshHandler) // refresh token
	return &Server{mux: mux}

}

func (s *Server) Start() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: s.mux,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
