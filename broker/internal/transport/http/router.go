package http

import (
	"github.com/anton-uvarenko/cinema/broker-service/internal/core"
	md "github.com/anton-uvarenko/cinema/broker-service/internal/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

type Router struct {
	controllers *Controllers
}

const port = "80"
const ReadTimeout = 15 * time.Second
const WriteTimeout = 15 * time.Second

var verified = true
var notVerified = false

func NewRouter(controllers *Controllers) *Router {
	return &Router{
		controllers: controllers,
	}
}

func (r *Router) InitRoutes() http.Handler {
	app := chi.NewRouter()
	app.Use(middleware.Heartbeat("/ping"))
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	InitCorsPolicy(app)

	app.Route("/auth", func(router chi.Router) {
		router.Post("/signin", r.controllers.AuthController.SignIn)
		router.Post("/signup", r.controllers.AuthController.SignUp)
		router.Post("/google", r.controllers.SocialController.GoogleAuth)
		router.Post("/facebook", r.controllers.SocialController.FacebookAuth)
	})

	app.Route("/verify", func(router chi.Router) {
		mid := md.AuthMiddleware{
			Recovery: false,
			UserType: []core.UserType{
				core.Premium,
				core.Basic,
				core.Admin,
			},
			Verification: &notVerified,
		}
		router.Use(mid.TokenVerify)
		router.Get("/send", r.controllers.VerificationController.SendCode)
		router.Post("/check", r.controllers.VerificationController.VerifyCode)
	})

	app.Route("/recover", func(router chi.Router) {
		router.Post("/send", r.controllers.PassRecoveryController.SendCode)
		router.Post("/check", r.controllers.PassRecoveryController.VerifyCode)

		router.Group(func(rout chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: true,
				UserType: []core.UserType{
					core.Premium,
					core.Basic,
					core.Admin,
				},
			}
			rout.Use(mid.TokenVerify)
			rout.Post("/change", r.controllers.PassRecoveryController.UpdatePassword)
		})
	})

	app.Route("/films", func(router chi.Router) {
		router.Group(func(rt chi.Router) {
			rt.Get("/mane-page", r.controllers.FilmsController.ManePageRequest)
		})

		router.Group(func(rt chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: false,
				UserType: []core.UserType{
					core.Admin,
				},
				Verification: &verified,
			}
			rt.Use(mid.TokenVerify)

			// films
			rt.Post("/", r.controllers.FilmsController.RedirectRequest)
			rt.Put("/{id}/update/", r.controllers.FilmsController.RedirectRequest)
			rt.Patch("/{id}/update/", r.controllers.FilmsController.RedirectRequest)
			rt.Delete("/{id}/delete/", r.controllers.FilmsController.RedirectRequest)
		})

		router.Group(func(rt chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: false,
				UserType: []core.UserType{
					core.Admin,
					core.Basic,
					core.Premium,
				},
				Verification: &verified,
			}
			rt.Use(mid.TokenVerify)

			rt.Get("/", r.controllers.FilmsController.RedirectRequest)
			rt.Get("/{id}/", r.controllers.FilmsController.RedirectRequest)
		})
	})

	app.Route("/films/genres", func(router chi.Router) {
		router.Group(func(rt chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: false,
				UserType: []core.UserType{
					core.Admin,
				},
				Verification: &verified,
			}
			rt.Use(mid.TokenVerify)

			rt.Post("/", r.controllers.FilmsController.RedirectRequest)
			rt.Put("/{id}/update/", r.controllers.FilmsController.RedirectRequest)
			rt.Patch("/{id}/update/", r.controllers.FilmsController.RedirectRequest)
			rt.Delete("/{id}/delete/", r.controllers.FilmsController.RedirectRequest)
		})

		router.Group(func(rt chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: false,
				UserType: []core.UserType{
					core.Admin,
					core.Basic,
					core.Premium,
				},
				Verification: &verified,
			}
			rt.Use(mid.TokenVerify)

			rt.Get("/", r.controllers.FilmsController.RedirectRequest)
			rt.Get("/{id}/", r.controllers.FilmsController.RedirectRequest)
		})
	})

	app.Route("/actors", func(router chi.Router) {
		router.Group(func(rt chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: false,
				UserType: []core.UserType{
					core.Admin,
				},
				Verification: &verified,
			}
			rt.Use(mid.TokenVerify)

			rt.Post("/", r.controllers.FilmsController.RedirectRequest)
			rt.Put("/{id}/update/", r.controllers.FilmsController.RedirectRequest)
			rt.Patch("/{id}/update/", r.controllers.FilmsController.RedirectRequest)
			rt.Delete("/{id}/delete/", r.controllers.FilmsController.RedirectRequest)
		})

		router.Group(func(rt chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: false,
				UserType: []core.UserType{
					core.Admin,
					core.Basic,
					core.Premium,
				},
				Verification: &verified,
			}
			rt.Use(mid.TokenVerify)

			rt.Get("/", r.controllers.FilmsController.RedirectRequest)
			rt.Get("/{id}/", r.controllers.FilmsController.RedirectRequest)
		})
	})

	app.Route("/playlists", func(router chi.Router) {
		router.Group(func(rt chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: false,
				UserType: []core.UserType{
					core.Admin,
					core.Premium,
				},
				Verification: &verified,
			}
			rt.Use(mid.TokenVerify)

			rt.Post("/", r.controllers.FilmsController.RedirectRequest)
			rt.Put("/{id}/update/", r.controllers.FilmsController.RedirectRequest)
			rt.Patch("/{id}/update/", r.controllers.FilmsController.RedirectRequest)
			rt.Delete("/{id}/delete/", r.controllers.FilmsController.RedirectRequest)
		})

		router.Group(func(rt chi.Router) {
			mid := md.AuthMiddleware{
				Recovery: false,
				UserType: []core.UserType{
					core.Admin,
					core.Basic,
					core.Premium,
				},
				Verification: &verified,
			}
			rt.Use(mid.TokenVerify)

			rt.Get("/", r.controllers.FilmsController.RedirectRequest)
			rt.Get("/{id}/", r.controllers.FilmsController.RedirectRequest)
		})
	})

	app.Route("/comment", func(router chi.Router) {
		mid := md.AuthMiddleware{
			Recovery: false,
			UserType: []core.UserType{
				core.Admin,
				core.Basic,
				core.Premium,
			},
			Verification: &verified,
		}
		router.Use(mid.TokenVerify)

		router.Post("/add", r.controllers.CommentsController.AddComment)
		router.Get("/get-public", r.controllers.CommentsController.GetPublicComments)
		router.Get("/get-private", r.controllers.CommentsController.GetPrivateComments)
		router.Post("/like/{id}", r.controllers.CommentsController.LikeComment)
	})

	app.Route("/user-data", func(router chi.Router) {
		mid := md.AuthMiddleware{
			Recovery: false,
			UserType: []core.UserType{
				core.Admin,
				core.Basic,
				core.Premium,
			},
			Verification: &verified,
		}
		router.Use(mid.TokenVerify)

		router.Post("/add", r.controllers.UserDataController.AddData)
		router.Get("/get", r.controllers.UserDataController.GetData)
		router.Delete("/image", r.controllers.UserDataController.DeleteImage)
	})

	app.Route("/admin", func(router chi.Router) {
		mid := md.AuthMiddleware{
			Recovery: false,
			UserType: []core.UserType{
				core.Admin,
			},
			Verification: &verified,
		}
		router.Use(mid.TokenVerify)

		router.Patch("/user-type", r.controllers.AdminContoller.UpdateUserType)
		router.Delete("/user", r.controllers.AdminContoller.DeleteUser)
		router.Get("/", r.controllers.AdminContoller.GetAllUsers)
	})

	app.Route("/premium", func(router chi.Router) {
		mid := md.AuthMiddleware{
			Recovery: false,
			UserType: []core.UserType{
				core.Admin,
				core.Basic,
				core.Premium,
			},
			Verification: &verified,
		}
		router.Use(mid.TokenVerify)

		router.Post("/buy", r.controllers.PremiumController.BuyPremium)
	})
	return app
}

func InitCorsPolicy(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"PUT", "GET", "POST", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func InitHttpServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         "0.0.0.0:" + port,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		Handler:      handler,
	}
}
