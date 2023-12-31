package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"slices"

	"github.com/ogen-go/ogen/middleware"
	"github.com/tiagoposse/connect/ent"

	"github.com/tiagoposse/connect/ent/ogent"
	_ "github.com/tiagoposse/connect/ent/runtime"
	"github.com/tiagoposse/connect/internal/config"
	ctrl "github.com/tiagoposse/connect/internal/controller"
	"github.com/tiagoposse/connect/internal/utils"

	_ "github.com/lib/pq"

	"entgo.io/ent/dialect"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("loading config: %v", err)
		return
	}

	log.SetLevel(log.DebugLevel)
	resolver := utils.NewResolver()

	errs := cfg.Validate(resolver)
	if len(errs) > 0 {
		for _, err := range errs {
			log.Errorf("%v", err)
		}
		return
	}
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open(
		dialect.Postgres,
		fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable",
			cfg.Database.Username,
			*cfg.Database.Password.Value,
			cfg.Database.Host,
			cfg.Database.Database,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Run the migrations.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	c, err := ctrl.NewController(client, cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Init(context.Background()); err != nil {
		log.Fatal(err)
	}

	skipOps := []string{
		"googleAuthStart",
		"googleAuthCallback",
		"userpassLogin",
		"status",
	}

	// Start listening.
	srv, err := ogent.NewServer(
		c,
		c.GetAuthorizationHandler(),
		ogent.WithPathPrefix("/api/v1"),
		ogent.WithMiddleware(ctrl.GetAuthAfterUrl, func(req middleware.Request, next middleware.Next) (resp middleware.Response, err error) {
			resp, err = next(req)
		
			if err != nil {
				return resp, err
			}

			if slices.Index(skipOps, req.OperationID) == -1 {
				err = c.AuditAction(req.Context, req.OperationID)
			}
		
			return resp, err
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	corsOpts := []handlers.CORSOption{
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"POST", "PATCH", "GET", "DELETE"}),
		handlers.AllowedHeaders([]string{"content-type", "x-page", "x-items-per-page"}),
		handlers.ExposedHeaders([]string{"x-total"}),
	}

	var origin string
	if cfg.Web.ServeFrontend {
		origin = cfg.Web.ExternalUrl
	} else {
		origin = cfg.Web.FrontendUrl
	}
	corsOpts = append(corsOpts, handlers.AllowedOrigins([]string{origin}))

	server := http.Server{
		Handler: handlers.CORS(corsOpts...)(srv),
		Addr:    "0.0.0.0:8000",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	log.Debugf("Serving at %s\n", server.Addr)
	err = server.ListenAndServeTLS(cfg.Web.Tls.Certificate, cfg.Web.Tls.Key)
	if err != nil {
		fmt.Println(err)
	}
}
