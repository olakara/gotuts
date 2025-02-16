package main

import (
	"embed"
	"flag"
	"github.com/yaitoo/xun"
	"github.com/yaitoo/xun/ext/htmx"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

//go:embed app/pages
var fsys embed.FS

func main() {

	var dev bool
	flag.BoolVar(&dev, "dev", false, "it is development environment")

	flag.Parse()

	var opts []xun.Option
	if dev {
		// use local filesystem in development, and watch files to reload automatically
		opts = []xun.Option{xun.WithFsys(os.DirFS("./app")), xun.WithWatch()}
	} else {
		// use embed resources in production environment
		views, _ := fs.Sub(fsys, "app")
		opts = []xun.Option{xun.WithFsys(views)}
	}

	opts = append(opts, xun.WithInterceptor(htmx.New()))
	app := xun.New(opts...)

	app.Use(func(next xun.HandleFunc) xun.HandleFunc {
		return func(c *xun.Context) error {
			n := time.Now()
			defer func() {
				duration := time.Since(n)

				log.Println(c.Routing.Pattern, duration)
			}()
			return next(c)
		}
	})

	app.Get("/{$}", func(c *xun.Context) error {
		return c.View(map[string]string{
			"Name": "go-xun",
		})
	})

	app.Start()
	defer app.Close()

	if dev {
		slog.Default().Info("xun-admin is running in development")
	} else {
		slog.Default().Info("xun-admin is running in production")
	}

	err := http.ListenAndServe(":80", http.DefaultServeMux)
	if err != nil {
		panic(err)
	}
}
