package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/golangnigeria/MyNneFarm/internal/mailer"
	"github.com/golangnigeria/MyNneFarm/internal/repository"
	"github.com/golangnigeria/MyNneFarm/internal/repository/repo"
	"github.com/joho/godotenv"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string // Data Source Name for the database connection
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	limiter struct {
		enabled bool
		rps     float64
		burst   int
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config config
	logger *log.Logger
	DB     repository.DatabaseRepository
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("DATABASE_URL"), "PostgreSQL data source name")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP server host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 2525, "SMTP server port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "4ede3760dd00c6", "SMTP server username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "1f87de4dec9bdb", "SMTP server password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "stackninja.pro@gmail.com", "SMTP sender email address")

	// Rate limiting configuration
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiting")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Requests per second for rate limiting")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Burst size for rate limiting")

	// Parse command line flags
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config: cfg,
		logger: logger,
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	conn, err := app.connectToDB()
	if err != nil {
		app.logger.Fatal(err)
	}

	app.DB = &repo.NeonDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Printf("Starting %s server on port %d", app.config.env, app.config.port)
	err = srv.ListenAndServe()
	if err != nil {
		app.logger.Fatal(err)
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	shutdownError := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		app.logger.Printf("caught signal", map[string]string{
			"signal": s.String(),
		})
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Call Shutdown() on the server like before, but now we only send on the
		// shutdownError channel if it returns an error
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}
		// Log a message to say that we're waiting for any background goroutines to
		// complete their tasks.
		app.logger.Printf("completing background tasks", map[string]string{
			"addr": srv.Addr,
		})
		// Call Wait() to block until our WaitGroup counter is zero --- essentially
		// blocking until the background goroutines have finished. Then we return nil on
		// the shutdownError channel, to indicate that the shutdown completed without
		// any issues.
		app.wg.Wait()
		shutdownError <- nil
	}()
	app.logger.Printf("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	err = <-shutdownError
	if err != nil {
		return err
	}
	app.logger.Printf("stopped server", map[string]string{
		"addr": srv.Addr,
	})
	return nil
}
