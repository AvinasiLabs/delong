package internal

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"
)

type Service interface {
	Name() string
	Init(context.Context) error
	Start(context.Context) error
	Stop(context.Context) error
}

type ServiceManager struct {
	services []Service
}

func NewServiceManager(srvs ...Service) *ServiceManager {
	return &ServiceManager{
		services: srvs,
	}
}

func (sm *ServiceManager) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	for _, srv := range sm.services {
		srv := srv
		g.Go(func() error {
			srvname := srv.Name()
			log.Printf("Initializing %v ...", srvname)
			srv.Init(ctx)
			log.Printf("Starting %v ...", srvname)
			return srv.Start(ctx)
		})
	}

	err := g.Wait()

	for _, srv := range sm.services {
		_ = srv.Stop(ctx)
	}

	return err
}
