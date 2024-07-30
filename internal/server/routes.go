package server

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"
	"time"

	"toolbox/cmd/web"
	"toolbox/internal/domain"
	"toolbox/internal/features/todos"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/websocket", s.websocketHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)

	list := domain.NewTodos()
	list.Add("Refactor Authentication Module", "Review and improve the current authentication module for better performance and security. Ensure that password hashing algorithms are up-to-date and that user sessions are managed effectively. Write unit tests to cover edge cases and potential vulnerabilities.")
	list.Add("Implement API Rate Limiting", "Add rate limiting to the API endpoints to prevent abuse and ensure fair usage. Configure limits based on the type of user and endpoint, and implement appropriate error responses for users who exceed their allotted requests. Update documentation to reflect the new rate limiting policies.")
	list.Add("Fix Responsive Design Issues on Checkout Page", "Address responsiveness problems on the checkout page for various devices and screen sizes. Test across multiple browsers and devices to ensure a seamless user experience. Adjust CSS and HTML as needed and perform regression testing to ensure no new issues have been introduced.")

	todos.Mount(r, todos.NewHandler(todos.NewService(list)))

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
