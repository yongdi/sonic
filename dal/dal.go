package dal

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"sonic/config"
	"sonic/consts"
	sonicLog "sonic/log"
	"sonic/model/entity"
	"sonic/util/xerr"
)

// mysqlDsn example  user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
const mysqlDsn = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s&readTimeout=1s&writeTimeout=1s&interpolateParams=true"

var (
	DB     *gorm.DB
	DBType consts.DBType
)

func NewGormDB(conf *config.Config, gormLogger logger.Interface) *gorm.DB {
	var err error
	if conf.SQLite3 != nil && conf.SQLite3.Enable {
		DB, err = initSQLite(conf, gormLogger)
		if err != nil {
			sonicLog.Fatal("open SQLite3 error", zap.Error(err))
		}
		DBType = consts.DBTypeSQLite
	} else if conf.MySQL != nil {
		DB, err = initMySQL(conf, gormLogger)
		if err != nil {
			sonicLog.Fatal("connect to MySQL error", zap.Error(err))
		}
		DBType = consts.DBTypeMySQL
	} else {
		sonicLog.Fatal("No database available")
	}
	if DB == nil {
		sonicLog.Fatal("no available database")
	}
	sonicLog.Info("connect database success")
	sqlDB, err := DB.DB()
	if err != nil {
		sonicLog.Fatal("get database connection error")
	}
	sqlDB.SetMaxIdleConns(200)
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	dbMigrate()
	return DB
}

func initMySQL(conf *config.Config, gormLogger logger.Interface) (*gorm.DB, error) {
	mysqlConfig := conf.MySQL
	if mysqlConfig == nil {
		return nil, xerr.WithMsg(nil, "nil MySQL config")
	}
	dsn := fmt.Sprintf(mysqlDsn, mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DB)
	sonicLog.Info("try connect to MySQL", zap.String("dsn", fmt.Sprintf(mysqlDsn, mysqlConfig.Username, "***", mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DB)))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                   gormLogger,
		PrepareStmt:              true,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})
	return db, err
}

func initSQLite(conf *config.Config, gormLogger logger.Interface) (*gorm.DB, error) {
	sqliteConfig := conf.SQLite3
	if sqliteConfig == nil {
		return nil, xerr.WithMsg(nil, "nil SQLite config")
	}
	sonicLog.Info("try to open SQLite3 db", zap.String("path", sqliteConfig.File))
	db, err := gorm.Open(sqlite.Open(sqliteConfig.File), &gorm.Config{
		Logger:                   gormLogger,
		PrepareStmt:              true,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})
	return db, err
}

func dbMigrate() {
	db := DB.Session(&gorm.Session{
		Logger: DB.Logger.LogMode(logger.Warn),
	})
	err := db.AutoMigrate(&entity.Attachment{}, &entity.Category{}, &entity.Comment{}, &entity.CommentBlack{}, &entity.Journal{},
		&entity.Link{}, &entity.Log{}, &entity.Menu{}, &entity.Meta{}, &entity.Option{}, &entity.Photo{}, &entity.Post{},
		&entity.PostCategory{}, &entity.PostTag{}, &entity.Tag{}, &entity.ThemeSetting{}, &entity.User{})
	if err != nil {
		sonicLog.Fatal("failed auto migrate db", zap.Error(err))
	}
}

type ctxTransaction struct{}

func GetDBByCtx(ctx context.Context) *gorm.DB {
	dbI := ctx.Value(ctxTransaction{})

	if dbI != nil {
		db, ok := dbI.(*gorm.DB)
		if !ok {
			panic("unexpected context db value type")
		}
		if db != nil {
			return db
		}
	}
	return DB.WithContext(ctx)
}

func SetCtxDB(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxTransaction{}, tx)
}

func Transaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	db := GetDBByCtx(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		txCtx := SetCtxDB(ctx, tx)
		return fn(txCtx)
	})
}
