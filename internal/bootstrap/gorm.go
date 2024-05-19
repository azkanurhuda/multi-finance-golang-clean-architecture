package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/config"
	"github.com/golang-migrate/migrate/v4"
	migrateMySQL "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"time"
)

func NewDatabase(config *config.Config, log *logrus.Logger) *gorm.DB {
	username := config.Database.Username
	password := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port
	databaseName := config.Database.Name
	idleConnection := config.Database.Pool.Idle
	maxConnection := config.Database.Pool.Max
	maxLifeTimeConnection := config.Database.Pool.Max
	timezone := url.QueryEscape(config.App.Timezone)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		username, password, host, port, databaseName, timezone)

	log.Logf(log.Level, "DSN: %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.Fatalf("failed to open connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	//err = doAutoMigrateDB(dsn)
	//if err != nil {
	//	log.Fatalf("failed to auto migrate database: %v", err)
	//} else {
	//	log.Info("Database migration successful")
	//}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

func doAutoMigrateDB(dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	driver, err := migrateMySQL.WithInstance(db, &migrateMySQL.Config{})
	if err != nil {
		return fmt.Errorf("failed to get driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///go/bin/database/migration",
		"mysql",
		driver,
	)

	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	return nil
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
