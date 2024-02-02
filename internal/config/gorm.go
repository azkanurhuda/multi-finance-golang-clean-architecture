package config

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migrateMySQL "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	database := viper.GetString("database.name")
	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.max")
	maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(database)

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
		fmt.Println("err2", err)
		log.Fatalf("failed to open connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		fmt.Println("err3", err)
		log.Fatalf("failed to connect database: %v", err)
	}

	err = doAutoMigrateDB(dsn)
	if err != nil {
		log.Fatalf("failed to auto migrate database: %v", err)
	} else {
		log.Info("Database migration successful")
	}

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
		"file://database/migration",
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
