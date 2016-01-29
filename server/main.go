package main

import (
	"net/http"

	// Import GoRPC libraries
	"github.com/sergei-svistunov/gorpc"
	"github.com/sergei-svistunov/gorpc/swagger_ui"
	"github.com/sergei-svistunov/gorpc/transport/http_json"
	"github.com/sergei-svistunov/gorpc/transport/http_json/adapter"

	// Import handlers
	handlerCalcSum "github.com/sergei-svistunov/gorpc-tutor/server/handlers/calc/sum"
	handlerHelloWorld "github.com/sergei-svistunov/gorpc-tutor/server/handlers/hello/world"
)

func main() {
	hm := gorpc.NewHandlersManager("github.com/sergei-svistunov/gorpc-tutor/server/handlers", gorpc.HandlersManagerCallbacks{})

	hm.MustRegisterHandlers(
		handlerHelloWorld.NewHandler(),
		handlerCalcSum.NewHandler(),
	)

	// Base API location
	http.Handle("/", http_json.NewAPIHandler(hm, nil, http_json.APIHandlerCallbacks{}))

	// Docs
	http.Handle("/swagger.json", http_json.NewSwaggerJSONHandler(hm, 0, http_json.SwaggerJSONCallbacks{}))
	http.Handle("/docs/", http.StripPrefix("/docs", swagger_ui.NewHTTPHandler()))

	// Client SDK
	http.Handle("/client.go", adapter.NewHandler(hm))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
