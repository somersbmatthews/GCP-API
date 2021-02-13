package dashboard

import (
	"github.com/gircapp/api/pg"
	"github.com/jinzhu/gorm"

	// TESTING: change dialects/postgres to dialers/postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/qor/admin"
)

// TODO: initiate qor admin with db from version 1 gorm

var postgrespassword string

func init() {
	// TESTING: uncomment access postgres passsword

	// password, err := accessPostgresPassword()
	// if err != nil {
	// 	panic(err)
	// }

	// postgrespassword = password

}

func SetupAdmin() *admin.Admin {

	DSN := "host=localhost user=gorm password=gorm database=postgres port=5432 sslmode=disable"

	// TESTING: change postgres to cloudsqlpostgres
	db, err := gorm.Open("postgres", DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	adminconfig := &admin.AdminConfig{
		SiteName: "Girc App Dashboard",
		DB:       db,
	}

	Admin := admin.New(adminconfig)

	Admin.AddResource(&pg.User{})
	Admin.AddResource(&pg.Incident{})

	// assetFS := bindatafs.AssetFS

	// assetFS.RegisterPath("api/dashboard/views")
	// assetFS.RegisterPath("api/vender/plugin/views")

	// // NameSpace return namespaced filesystem
	// namespacedFS := assetFS.NameSpace("dashboard")
	// namespacedFS.RegisterPath("api/dashboard/views")
	// // namespacedFS.PrependPath("")
	// // Will lookup file with name "filename.tmpl" from path `/web/app/myspecialviews` but not `/web/app/views`
	// namespacedFS.Asset("application.tmpl")
	// namespacedFS.Glob("*.tmpl")

	// // Generate a sub AssetFS with namespace
	// // adminAssetFS := assetFS.NameSpace("admin_related_files")
	// fmt.Println("this is working")
	// // Compile templates under registered view paths into binary
	// err = assetFS.Compile()
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// fmt.Println("this is working too")
	// // Generate a sub AssetFS with namespace
	// adminAssetFS := assetFS.NameSpace("admin_related_files")

	// // Register view paths into this name space
	// adminAssetFS.RegisterPath("<your_project>/app/admin_views")

	// // Access file that under registered views paths of current name space
	// adminAssetFS.Asset("admin_view.tmpl")

	// assetfs := assetfs.AssetFileSystem{}

	// Register path to AssetFS
	// assetfs.RegisterPath("/home/somersbmatthews/go/src/github.com/gircapp/api/dashboard/views")
	// if err != nil {
	// 	errMsg := errors.Errorf("RegisterPath caused an error: %v", err)
	// 	panic(errMsg)
	// }

	// Get file's content with name from path `/web/app/views`
	// _, err = assetfs.Asset("application.tmpl")
	// if err != nil {
	// 	errMsg := errors.Errorf("Asset() caused an error: %v", err)
	// 	panic(errMsg)
	// }

	// Prepend path to AssetFS
	// err = assetfs.PrependPath("/api/dashboard/views")
	// if err != nil {
	// 	errMsg := errors.Errorf("PrependPath caused an error: %v", err)
	// 	panic(errMsg)
	// }

	// // List matched files from assetfs
	// assetfs.Glob("*.tmpl")

	// Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("dashboard"))

	return Admin

}
