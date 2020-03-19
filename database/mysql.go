package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Creater interface {
	Create() error
}

type Closer interface {
	close() error
}


type Driver interface {
	Creater
	Closer
}

var DB *gorm.DB


type config struct {
	//sample:  "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	//see https://github.com/go-sql-driver/mysql#parameters
	// Database user
	User string
	// Database password
	Password string
	// Database db name
	DBName string
	// Database charset
	Charset string
	// parseTime=True&loc=Local......
	Parameters string
	// zero means defaultMaxIdleConn; negative means 0
	maxIdle           int
	// <= 0 means unlimited
	MaxOpenConn           int
	// maximum amount of time a connection may be reused
	MaxLifetime       int
}

type mysqlDB struct {
	config *config
	DB *gorm.DB
}

func NewMysqlDB(config *config) *mysqlDB {
	return &mysqlDB{config: config}
}

func InitMysql() error {
	config := &config{
		User:        "root",
		Password:    "123456",
		DBName:      "ccshop",
		Charset:     "utf8",
		MaxOpenConn: 200,
		MaxLifetime: int(time.Millisecond),
	}
	db := NewMysqlDB(config)
	err := db.Create()
	if err != nil {
		fmt.Println("create connection for mysql database has failed, err2es :", err)
		return err
	}
	DB = db.DB
	return nil
}

// create mysql database connections pool
func (mdb *mysqlDB) Create() (err error) {
	//sample:  "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	if len(mdb.config.Charset) <= 0 {
		mdb.config.Charset = "utf8"
	}
	if len(mdb.config.Parameters) <= 0 {
		mdb.config.Parameters = "parseTime=True&loc=Local"
	}
	source := fmt.Sprintf("%s:%s@/%s?charset=%s&%s",
		mdb.config.User, mdb.config.Password,
		mdb.config.DBName, mdb.config.Charset, mdb.config.Parameters)
	mdb.DB, err = gorm.Open("mysql", source)

	// set mysql db log level
	mdb.DB.LogMode(true)
	if err != nil {
		mdb.DB = nil
		return err
	}
	// SingularTable use singular table by default
	mdb.DB.SingularTable(true)

	mdb.setConnectionPool()
	return nil
}

// close gorm database connection
func (mdb *mysqlDB) Close() error {
	return mdb.DB.Close()
}

// set connections pool
func (mdb *mysqlDB) setConnectionPool()  {
	if mdb.config.maxIdle > 0 {
		mdb.DB.DB().SetMaxIdleConns(mdb.config.maxIdle)
	}
	if mdb.config.MaxOpenConn > 0 {
		mdb.DB.DB().SetMaxOpenConns(mdb.config.MaxOpenConn)
	}
	if mdb.config.MaxLifetime > 0 {
		mdb.DB.DB().SetConnMaxLifetime(time.Duration(mdb.config.MaxLifetime))
	}
}

func Db() *gorm.DB {

	return DB
}



