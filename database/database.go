package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// docker run -p 5432:5432 -d  -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=DB -v pgdata:/var/lib/postgresql/data postgres
func Initialization() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/DB"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&KeyValue{})

	return db
}
