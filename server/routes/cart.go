package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/middleware"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", h.FindCarts).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/carts-id", middleware.Auth(h.FindCartsByID)).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart-id", middleware.Auth(h.UpdateCart)).Methods("PATCH")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
}
