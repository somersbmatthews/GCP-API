package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gircapp/api/emailer"
	"github.com/gircapp/api/models"
	"github.com/gircapp/api/pg"
	"github.com/gircapp/api/restapi"
	"github.com/gircapp/api/restapi/operations"
	"github.com/go-openapi/loads"
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
	//fmt.Println("this is running first")
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	api := operations.NewGircAPI(swaggerSpec)
	server := restapi.NewServer(api)

	// wg := new(sync.WaitGroup)

	//server.Port = 8080

	defer server.Shutdown()
	server.ConfigureFlags()
	server.ConfigureAPI()

	// wg.Add(2)

	// localhost:8080/?key=hello%20golangcode.com
	//	emailConfirmationTemplate := template.Must(template.ParseFiles("email_confirmation_template.html"))

	http.HandleFunc("/directorverifybyemail", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		keys, ok := r.URL.Query()["key"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}
		key := keys[0]

		_, userID, _, _, _, err := emailer.DecodeJWTClaims(key)
		if err != nil {
			panic(err)
		}

		booleanTrue := true
		requestObject := models.Verify{
			UserID:   &userID,
			Verified: &booleanTrue,
		}

		_, ok = pg.VerifyExpert(ctx, requestObject, userID)
		if !ok {
			// TODO: handle cannot find expert here
		}

	})
	http.HandleFunc("/confirmemail", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		keys, ok := r.URL.Query()["key"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}

		key := keys[0]

		log.Println("Url Param 'key' is: " + string(key))

		email, userID, specialty, fullName, verified, err := emailer.DecodeJWTClaims(key)
		if err != nil {
			panic(err)
		}

		if verified == "false" {
			emailer.SendDirectorVerificationEmail(fullName, specialty, email, userID)
		}

		_, ok = pg.ConfirmEmail(ctx, email, userID)
		if !ok {
			// TODO: handle cannot find expert here
		}

		tmpl := template.Must(template.New("email-confirmation-thank-you").Parse(`<!DOCTYPE html>
		<html>
		 <head>
		  <title>
		  </title>
		  <meta content="summary_large_image" name="twitter:card">
		  <meta content="website" property="og:type">
		  <meta content="" property="og:description">
		  <meta content="https://preview.posted-stuff.com/RvJj-hegK-Ckiz-RF27/" property="og:url">
		  <meta content="https://pro-bee-beepro-thumbnails.s3.amazonaws.com/messages/700300/682982/1350708/6593804_large.jpg" property="og:image">
		  <meta content="" property="og:title">
		  <meta content="" name="description">
		  <meta charset="utf-8">
		  <meta content="width=device-width" name="viewport">
		  <link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet" type="text/css">
		  <style>
		   .bee-row-1,
				.bee-row-1 .bee-row-content {
					background-repeat: no-repeat
				}
		
				body {
					background-color: #fff;
					color: #000;
					font-family: Arial, Helvetica Neue, Helvetica, sans-serif
				}
		
				.bee-row-1 .bee-col-1 {
					border-bottom: 0 solid transparent;
					border-left: 0 solid transparent;
					border-right: 0 solid transparent;
					border-top: 0 solid transparent;
					padding-bottom: 5px;
					padding-top: 5px
				}
		
				.bee-row-1 .bee-col-1 .bee-block-1 {
					width: 100%
				}
		
				.bee-row-1 .bee-col-1 .bee-block-2 {
					padding: 10px
				}
		
				* {
					box-sizing: border-box
				}
		
				body,
				p {
					margin: 0
				}
		
				.bee-row-content {
					max-width: 500px;
					margin: 0 auto;
					display: flex
				}
		
				.bee-row-content .bee-col-w12 {
					flex: 12
				}
		
				.bee-image img {
					display: block;
					width: 100%
				}
		
				.bee-image {
					overflow: auto
				}
		
				.bee-image .bee-center {
					margin: 0 auto
				}
		
				.bee-text {
					overflow-wrap: anywhere
				}
		
				@media (max-width:520px) {
					.bee-row-content:not(.no_stack) {
						display: block
					}
				}
		  </style>
		 </head>
		 <body>
		  <div class="bee-page-container">
		   <div class="bee-row bee-row-1">
			<div class="bee-row-content">
			 <div class="bee-col bee-col-1 bee-col-w12">
			  <div class="bee-block bee-block-1 bee-image">
			   <img alt="" class="bee-center bee-autowidth" src="https://d15k2d11r6t6rl.cloudfront.net/public/users/Integrators/BeeProAgency/700300_682982/GRIC%20mark%20suite-01.png" style="max-width:500px;">
			  </div>
			  <div class="bee-block bee-block-2 bee-text">
			   <div class="bee-text-content" style="font-size: 12px; line-height: 120%; font-family: 'Roboto', Tahoma, Verdana, Segoe, sans-serif; color: #393d47;">
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="text-align: center; line-height: 14px;">
				 <span style="font-size: 28px; line-height: 33px;">
				  Thank you for confirming your email. Your submission to the GIRC will be reviewed.
				 </span>
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p aria-expanded="false" aria-label="Show options" role="button" style="line-height: 14px;" tabindex="0">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
				<p style="line-height: 14px;">
				 &nbsp;
				</p>
			   </div>
			  </div>
			 </div>
			</div>
		   </div>
		  </div>
		 </body>
		</html>`))

		tmpl.Execute(w, nil)

	})

	http.Handle("/", server.GetHandler())
	//fmt.Println("this is running")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	//	port := "8080"

	// go func() {
	// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 		fmt.Println(err)
	// 		log.Fatal(err)
	// 		wg.Done()
	// 	}
	// }()

	fmt.Printf("Listening on port %s \n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
		log.Fatal(err)
		// wg.Done()
	}
	// go func() {
	// 	http.ListenAndServe(":80", nil)
	// 	fmt.Println(err)
	// 	log.Fatal(err)
	// 	wg.Done()
	// }()

	// http.ListenAndServe(":8000", api.Serve(nil))

	// if err := server.Serve(); err != nil {
	// 	panic(err)
	// }

	// r := mux.NewRouter()

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

	// r.PathPrefix("/").Handler(api.Serve(nil))

	// http.Handle("/", r)

	// r.Schemes("http")

	// srv := &http.Server{
	// 	// Addr:         "127.0.0.1:8000",
	// 	Handler: r,
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// cert, key, err := accessTLSCertAndKey()
	// if err != nil {
	// 	panic(err)
	// }

	// log.Fatal(srv.ListenAndServeTLS(cert, key))
	// log.Fatal(srv.ListenAndServe())

	// mux := http.NewServeMux()
	// // Admin.MountTo("/admin", mux)
	// mux.Handle("/", api.Serve(nil))

	// err = http.ListenAndServe(":8080", api.Serve(nil)
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
