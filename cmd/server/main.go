package main

import (
	"context"
	"github.com/11DingKing/ai-course-cert-go/internal/audit"
	"github.com/11DingKing/ai-course-cert-go/internal/auth"
	"github.com/11DingKing/ai-course-cert-go/internal/config"
	"github.com/11DingKing/ai-course-cert-go/internal/httpapi"
	"github.com/11DingKing/ai-course-cert-go/internal/repository"
	"github.com/11DingKing/ai-course-cert-go/internal/service"
	"github.com/11DingKing/ai-course-cert-go/internal/storage"
	"github.com/11DingKing/ai-course-cert-go/internal/worker"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()
	db, e := storage.Open(cfg.DBPath)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	if e = storage.Migrate(context.Background(), db); e != nil {
		log.Fatal(e)
	}
	u := repository.Users{DB: db}
	s := service.Service{Users: u, Courses: repository.Courses{DB: db}, Submissions: repository.Submissions{DB: db}, Evidences: repository.Evidences{DB: db}, Reviews: repository.Reviews{DB: db}, Audit: audit.Logger{DB: db}}
	ttl, _ := time.ParseDuration(cfg.SessionTTL)
	h := httpapi.New(s, service.CourseService{Courses: repository.Courses{DB: db}}, auth.New(ttl))
	srv := &http.Server{Addr: cfg.Addr, Handler: h.Routes(), ReadHeaderTimeout: 5 * time.Second}
	w := worker.New(db)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	w.Start(ctx)
	go func() {
		<-ctx.Done()
		w.Stop()
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()
		srv.Shutdown(shutdownCtx)
	}()
	log.Printf("listening on %s", cfg.Addr)
	if e := srv.ListenAndServe(); e != nil && e != http.ErrServerClosed {
		log.Fatal(e)
	}
}
