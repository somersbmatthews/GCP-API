package main

import (
	"log"
	"net/http"

	"github.com/gircapp/api/dashboard"
	"github.com/gircapp/api/restapi"
	"github.com/gircapp/api/restapi/operations"
	"github.com/go-openapi/loads"
)

// type qorHandler func(writer http.ResponseWriter, request *http.Request)

// api := myapi.NewExampleAPI(...)
// globalMiddleware := func(h http.Handler) http.Handler
// authedMiddleware := func(h http.Handler) http.Handler

// http.ListenAndServe(":8080", globalMiddleware(api.Serve(authedMiddleware))

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGircAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer server.Shutdown()

	server.ConfigureAPI()

	// server.Port = 8080

	r := mux.NewRouter()
\
	r.PathPrefix("/").Handler(api.Serve(nil))

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	// mux := http.NewServeMux()
	// // Admin.MountTo("/admin", mux)
	// mux.Handle("/", api.Serve(nil))

	// http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	panic(err)
	// }

	// httpServer := http.Server{
	// 	Handler: mux,
	// 	Addr:    fmt.Sprintf(":%d", server.Port),
	// }

	// fmt.Println(render.Render(mux))

}

// func handler(writer http.ResponseWriter, request *http.Request) {
// 	ctx := context.Background()
// 	config := &render.Config{
// 		ViewPaths:     []string{"app/new_view_path"},
// 		DefaultLayout: "application", // default value is application
// 		// FuncMapMaker: func(*Render, *http.Request, http.ResponseWriter) template.FuncMap {
// 		// 	// genereate FuncMap that could be used when render template based on request info
// 		// },
// 	}
// 	Render := render.New(config)
// 	responder.With("html", func() {
// 		Render.Execute("demo/index", ctx, request, writer)
// 	}).With([]string{"json", "xml"}, func() {
// 		writer.Write([]byte("this is a json or xml request"))
// 	}).Respond(request)

// }

// func (f qorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	f(writer, request)
// }
