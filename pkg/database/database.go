package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	log "github.com/sirupsen/logrus"
)

//Driver list
const (
	DriverMySQL    = "mysql"
	DriverPostgres = "postgres"
)

// SpanName is standarized database query function span name.
const SpanName = "db.query"

//Db object
var (
	Master   *DB
	Slave    *DB
	dbTicker *time.Ticker
)

type (
	// Database is a wrapper around sqlx.DB, used to wrap custom logic such as db tracing.
	Database struct {
		*sqlx.DB
	}

	//DSNConfig for database source name
	DSNConfig struct {
		DSN string
	}

	//DBConfig for databases configuration
	DBConfig struct {
		SlaveDSN      string
		MasterDSN     string
		RetryInterval int
		MaxIdleConn   int
		MaxConn       int
	}

	//DB configuration
	DB struct {
		DBConnection  *Database
		DBString      string
		RetryInterval int
		MaxIdleConn   int
		MaxConn       int
		doneChannel   chan bool
	}

	Store struct {
		Master *Database
		Slave  *Database
	}

	Options struct {
		dbTx *sqlx.Tx
	}
)

func (s *Store) GetMaster() *Database {
	return s.Master
}

func (s *Store) GetSlave() *Database {
	return s.Slave
}

func New(cfg DBConfig, dbDriver string) *Store {
	masterDSN := cfg.MasterDSN
	slaveDSN := cfg.SlaveDSN

	Master = &DB{
		DBString:      masterDSN,
		RetryInterval: cfg.RetryInterval,
		MaxIdleConn:   cfg.MaxIdleConn,
		MaxConn:       cfg.MaxConn,
		doneChannel:   make(chan bool),
	}

	err := Master.ConnectAndMonitor(dbDriver)
	if err != nil {
		log.Fatal("Could not initiate Master DB connection: " + err.Error())
		return &Store{}
	}
	Slave = &DB{
		DBString:      slaveDSN,
		RetryInterval: cfg.RetryInterval,
		MaxIdleConn:   cfg.MaxIdleConn,
		MaxConn:       cfg.MaxConn,
		doneChannel:   make(chan bool),
	}
	err = Slave.ConnectAndMonitor(dbDriver)
	if err != nil {
		log.Fatal("Could not initiate Slave DB connection: " + err.Error())
		return &Store{}
	}

	dbTicker = time.NewTicker(time.Second * 2)

	return &Store{Master: Master.DBConnection, Slave: Slave.DBConnection}
}

// Connect to database
func (d *DB) Connect(driver string) error {
	var db *sqlx.DB
	var err error
	db, err = sqlx.Open(driver, d.DBString)

	if err != nil {
		log.Println("[Error]: DB open connection error", err.Error())
	} else {
		d.DBConnection = &Database{DB: db}
		err = db.Ping()
		if err != nil {
			log.Println("[Error]: DB connection error", err.Error())
		}
		return err
	}

	db.SetMaxOpenConns(d.MaxConn)
	db.SetMaxIdleConns(d.MaxIdleConn)

	return err
}

// ConnectAndMonitor to database
func (d *DB) ConnectAndMonitor(driver string) error {
	err := d.Connect(driver)

	if err != nil {
		log.Printf("Not connected to database %s, trying", d.DBString)
		return err
	}

	ticker := time.NewTicker(time.Duration(d.RetryInterval) * time.Second)
	go func() error {
		for {
			select {
			case <-ticker.C:
				if d.DBConnection == nil {
					d.Connect(driver)
				} else {
					err := d.DBConnection.Ping()
					if err != nil {
						log.Println("[Error]: DB reconnect error", err.Error())
						return err
					}
				}
			case <-d.doneChannel:
				return nil
			}
		}
	}()
	return nil
}

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (db *Database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.ExecContext(ctx, query, args...)
}

// GetContext using this DB.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (db *Database) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.GetContext(ctx, dest, query, args...)
}

// MustExecContext (panic) runs MustExec using this database.
// Any placeholder parameters are replaced with supplied args.
func (db *Database) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.MustExecContext(ctx, query, args...)
}

// NamedExecContext using this DB.
// Any named placeholder parameters are replaced with fields from arg.
func (db *Database) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", arg)

	return db.DB.NamedExecContext(ctx, query, arg)
}

// NamedQueryContext using this DB.
// Any named placeholder parameters are replaced with fields from arg.
func (db *Database) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", arg)

	return db.DB.NamedQueryContext(ctx, query, arg)
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (db *Database) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.QueryContext(ctx, query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (db *Database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.QueryRowContext(ctx, query, args...)
}

// QueryRowxContext queries the database and returns an *sqlx.Row.
// Any placeholder parameters are replaced with supplied args.
func (db *Database) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.QueryRowxContext(ctx, query, args...)
}

// QueryxContext queries the database and returns an *sqlx.Rows.
// Any placeholder parameters are replaced with supplied args.
func (db *Database) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.QueryxContext(ctx, query, args...)
}

// SelectContext using this DB.
// Any placeholder parameters are replaced with supplied args.
func (db *Database) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.SelectContext(ctx, dest, query, args...)
}
