package router

import (
	"database/sql"
	"goauthapi/handlers"
	"goauthapi/middlewares"

	"github.com/gorilla/mux"
)

type Router struct {}

func (r Router) Routes(db *sql.DB, handlers handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/signup", handlers.Signup(db)).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login(db)).Methods("POST")
	router.HandleFunc("/api/protected", middlewares.VerifyTokenMiddleware(handlers.Protected(db))).Methods("GET")

	return router
}
