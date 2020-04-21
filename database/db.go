package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	c "github.com/sercel98/go-gin/config"
	m "github.com/sercel98/go-gin/models"
)

//TODO: implement singleton
type pgConnection struct {
	Connection *gorm.DB
}

func NewPgConnection(config *c.Configurations) *pgConnection {

	con, err := gorm.Open("postgres", "host="+config.Database.DBHost+" port=5432 user="+
		config.Database.DBUser+" dbname="+config.Database.DBName+" password="+config.Database.DBPassword)
	if err != nil {
		fmt.Printf("Problems connecting to the d	atabase, %v", err)
	}
	return &pgConnection{
		Connection: con,
	}
}

func (pg *pgConnection) Close() {
	err := pg.Connection.Close()
	if err != nil {
		fmt.Printf("Error while closing the db connection, %v", err)
	}
}

func (pg *pgConnection) CreateTables() {
	pg.Connection.AutoMigrate(&m.Article{}, &m.Video{})
	fmt.Println(pg.Connection.HasTable(&m.Article{}))
}
