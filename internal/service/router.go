package service

import (
	"github.com/go-chi/chi"
	"github.com/rarimo/zk-biometrics-svc/internal/data/pg"
	"github.com/rarimo/zk-biometrics-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxKVQ(pg.NewKVQ(s.cfg.DB().Clone())),
		),
	)
	r.Route("/integrations/zk-biometrics-svc", func(r chi.Router) {
		r.Route("/value", func(r chi.Router) {
			r.Delete("/", handlers.DeleteData)
			r.Post("/", handlers.AddData)
			r.Get("/", handlers.GetData)
		})
	})

	return r
}
