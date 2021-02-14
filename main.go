package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gircapp/api/restapi"
	"github.com/gircapp/api/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/gorilla/mux"
)

// _ "github.com/GoAdminGroup/go-admin/adapter/gorilla"             // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
// _ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // Import the sql driver
// _ "github.com/GoAdminGroup/go-admin/modules/service"             // Import the theme
// _ "github.com/GoAdminGroup/go-admin/plugins/admin"
// _ "github.com/GoAdminGroup/themes/adminlte"

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

	// eng := engine.Default()

	// GoAdmin global configuration, can also be imported as a json file.
	// cfg := config.Config{
	// 	Databases: config.DatabaseList{"postgres": config.Database{

	// 		Host:       "127.0.0.1",
	// 		Port:       "5432",
	// 		User:       "gorm",
	// 		Pwd:        "gorm",
	// 		Name:       "postgres",
	// 		MaxIdleCon: 50,
	// 		MaxOpenCon: 150,
	// 		Driver:     config.DriverPostgresql,
	// 	},
	// 	},
	// 	UrlPrefix: "dashboard", // The url prefix of the website.
	// 	// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
	// 	Store: config.Store{
	// 		Path:   "./uploads",
	// 		Prefix: "uploads",
	// 	},
	// 	Language: language.EN,
	// }

	// conn := eng.PostgresqlConnection()

	// globalConn = conn

	// eng.ResolvePostgresqlConnection(SetConn)

	// Admin := admin.NewAdmin()
	// // Add configuration and plugins, use the Use method to mount to the web framework.
	// _ = eng.AddConfig(cfg).
	// 	AddPlugins(Admin).
	// 	Use(r)

	// _ = r.Run(":9033")

	r.PathPrefix("/").Handler(api.Serve(nil))

	http.Handle("/", r)

	r.Schemes("http")

	srv := &http.Server{
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// cert, key, err := accessTLSCertAndKey()
	// if err != nil {
	// 	panic(err)
	// }

	// log.Fatal(srv.ListenAndServeTLS(cert, key))
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

// var globalConn db.Connection

// func SetConn(conn db.Connection) {
// 	globalConn = conn
// }

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

// func accessTLSCertAndKey() (string, string, error) {

// 	keyName := "projects/gircapp/secrets/key/versions/latest"
// 	certName := "projects/gircapp/secrets/key/versions/latest"

// 	// Create the client.
// 	ctx := context.Background()
// 	client, err := secretmanager.NewClient(ctx)
// 	if err != nil {
// 		return "", "", fmt.Errorf("failed to create secretmanager client: %v", err)
// 	}

// 	// Build the request.
// 	req := &secretmanagerpb.AccessSecretVersionRequest{
// 		Name: keyName,
// 	}

// 	// Call the API.
// 	result, err := client.AccessSecretVersion(ctx, req)
// 	if err != nil {
// 		return "", "", fmt.Errorf("failed to access secret version: %v", err)
// 	}

// 	key := string(result.Payload.Data)

// 	req = &secretmanagerpb.AccessSecretVersionRequest{
// 		Name: certName,
// 	}

// 	// Call the API.
// 	result, err = client.AccessSecretVersion(ctx, req)
// 	if err != nil {
// 		return "", "", fmt.Errorf("failed to access secret version: %v", err)
// 	}

// 	cert := string(result.Payload.Data)

// 	return cert, key, nil
// }
