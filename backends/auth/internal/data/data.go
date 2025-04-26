package data

import (
	"auth/internal/conf"
	"auth/internal/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderData is data providers.
var ProviderData = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	db, err := newDB(c, logger)
	if err != nil {
		log.Errorf("[db] newDB error: %v", err)
		return nil, cleanup, err
	}

	data := &Data{
		db: db,
	}

	return data, cleanup, nil
}

func newDB(c *conf.Data, logger log.Logger) (db *gorm.DB, err error) {
	dsn := parseMysqlDSN(c, logger)
	if dsn == "" {
		err = global.ErrInvalidDSN
		return
	}

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("[db] failed to connect database, %s", err.Error())
		err = global.ErrFailedConnect
		return
	}

	return
}

func parseMysqlDSN(c *conf.Data, logger log.Logger) (dsn string) {
	if c == nil || c.Database == nil {
		return
	}

	if c.Database.GetDriver() == "mysql" {
		dsn = c.Database.GetSource()
	}

	return
}
