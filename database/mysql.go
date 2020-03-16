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


type Config struct {
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

type MysqlDB struct {
	Config *Config
	DB *gorm.DB
}

func NewMysqlDB(config *Config) *MysqlDB {
	return &MysqlDB{Config: config}
}

// create mysql database connections pool
func (mdb *MysqlDB) Create() (err error) {
	//sample:  "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	if len(mdb.Config.Charset) <= 0 {
		mdb.Config.Charset = "utf8"
	}
	if len(mdb.Config.Parameters) <= 0 {
		mdb.Config.Parameters = "parseTime=True&loc=Local"
	}
	source := fmt.Sprintf("%s:%s@/%s?charset=%s&%s",
		mdb.Config.User, mdb.Config.Password,
		mdb.Config.DBName, mdb.Config.Charset, mdb.Config.Parameters)
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
func (mdb *MysqlDB) Close() error {
	return mdb.DB.Close()
}

// set connections pool
func (mdb *MysqlDB) setConnectionPool()  {
	if mdb.Config.maxIdle > 0 {
		mdb.DB.DB().SetMaxIdleConns(mdb.Config.maxIdle)
	}
	if mdb.Config.MaxOpenConn > 0 {
		mdb.DB.DB().SetMaxOpenConns(mdb.Config.MaxOpenConn)
	}
	if mdb.Config.MaxLifetime > 0 {
		mdb.DB.DB().SetConnMaxLifetime(time.Duration(mdb.Config.MaxLifetime))
	}
}



