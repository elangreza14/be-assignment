package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/elangreza14/be-assignment/account/cmd/http/routes"
	"github.com/elangreza14/be-assignment/account/controller"
	"github.com/elangreza14/be-assignment/account/middleware"
	"github.com/elangreza14/be-assignment/account/repository"
	"github.com/elangreza14/be-assignment/account/service"
	genaccount "github.com/elangreza14/be-assignment/gen/go"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// setup env
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	err = godotenv.Load(curDir + "/.env")
	errChecker(err)

	// logger
	logger, err := Logger()
	errChecker(err)
	defer logger.Sync()

	ctx := context.Background()

	// db connection
	db, err := DB(ctx, logger)
	errChecker(err)

	// dependency injection
	userRepository := repository.NewUserRepository(db)
	tokenRepository := repository.NewTokenRepository(db)
	currencyRepository := repository.NewCurrencyRepository(db)
	productRepository := repository.NewProductRepository(db)
	accountRepository := repository.NewAccountRepository(db)

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	paymentClient := genaccount.NewPaymentClient(conn)

	authService := service.NewAuthService(userRepository, tokenRepository)
	currencyService := service.NewCurrencyService(currencyRepository)
	productService := service.NewProductService(productRepository)
	accountService := service.NewAccountService(accountRepository, paymentClient)

	authController := controller.NewAuthController(authService)
	accountController := controller.NewAccountController(accountService)
	currencyController := controller.NewCurrencyController(currencyService)
	productController := controller.NewProductController(productService)

	authMiddleWare := middleware.NewAuthMiddleware(authService)

	// router
	if os.Getenv("ENV") != "DEVELOPMENT" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	// cors middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{fmt.Sprintf("http://localhost%s", os.Getenv("HTTP_PORT"))}
	router.Use(cors.New(config))

	// logger middleware
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// pinger
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// group api
	apiGroup := router.Group("/api")
	routes.AuthRoute(apiGroup, authController)
	routes.CurrencyRoute(apiGroup, currencyController)
	routes.ProductRoute(apiGroup, productController)
	routes.AccountRoute(apiGroup, authMiddleWare, accountController)

	srv := &http.Server{
		Addr:    os.Getenv("HTTP_PORT"),
		Handler: router.Handler(),
	}

	go func() {
		logger.Info("listening http", zap.String("port", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	wait := gracefulShutdown(context.Background(), logger, time.Second*5,
		func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
		func(ctx context.Context) error {
			db.Close()
			return nil
		})

	<-wait
}

func errChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func DB(ctx context.Context, logger *zap.Logger) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOSTNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSL"),
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE SCHEMA IF NOT EXISTS ` + os.Getenv("POSTGRES_SCHEME"))
	if err != nil {
		return nil, err
	}

	newConnString := fmt.Sprintf("%s&search_path=%s", connString, os.Getenv("POSTGRES_SCHEME"))

	db, err = sql.Open("postgres", newConnString)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return nil, err
		}
	}

	config, err := pgxpool.ParseConfig(newConnString)
	if err != nil {
		return nil, err
	}

	logLevel, err := tracelog.LogLevelFromString(logger.Level().String())
	if err != nil {
		return nil, err
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxzap.NewLogger(logger),
		LogLevel: logLevel,
	}

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Logger() (*zap.Logger, error) {
	logger := zap.NewExample(zap.IncreaseLevel(zap.InfoLevel))

	if os.Getenv("ENV") != "DEVELOPMENT" {
		return zap.NewProduction()
	}

	return logger, nil
}

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, logger *zap.Logger, timeout time.Duration, ops ...operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		logger.Info("shutting down")

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		go func() {
			<-ctx.Done()
			logger.Info("force quit the app")
			wait <- struct{}{}
		}()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			go func(key int, op operation) {
				defer wg.Done()
				processName := fmt.Sprintf("process %d", key)

				if err := op(ctx); err != nil {
					logger.Error(processName, zap.Error(err), zap.Bool("success", true))
					return
				}

				logger.Info(processName, zap.Bool("success", true))
			}(key, op)
		}

		wg.Wait()
		cancel()
		close(wait)
	}()

	return wait
}
