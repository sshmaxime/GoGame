package api

import "github.com/rs/cors"

var CorsMiddleware = cors.New(cors.Options{
	AllowedOrigins: []string{"http://localhost:3000"},
	AllowedMethods: []string{"GET", "POST"},
})
