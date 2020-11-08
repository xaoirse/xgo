package repository

import (
	"database/sql"
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/xaoirse/xgo/graph/model"

	"github.com/jinzhu/gorm"
)

// Repository defines a repository for access the database.
type Repository struct {
	db *gorm.DB
}

var rep *Repository

const (
	// SQLITE represents SQLite3
	SQLITE = "sqlite3"
	// DBname represents database file name
	DBname = "database.db"
)

// InitDB initialize a database connection.
func InitDB() {
	db, err := gorm.Open(SQLITE, DBname)
	if err != nil {
		log.Fatalln("Error in opening db:", err)
	}
	db.AutoMigrate(
		&model.User{})
	db.LogMode(true)
	rep = &Repository{}
	rep.db = db
}

// GetRepository returns the object of repository.
func GetRepository() *Repository {
	return rep
}

// GetDB returns the object of gorm.DB.
func GetDB() *gorm.DB {
	return rep.db
}

// Model specify the model you would like to run db operations
func (rep *Repository) Model(value interface{}) *gorm.DB {
	return rep.db.Model(value)
}

// Select specify fields that you want to retrieve from database when querying, by default, will select all fields;
func (rep *Repository) Select(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Select(query, args...)
}

// Find find records that match given conditions.
func (rep *Repository) Find(out interface{}, where ...interface{}) *gorm.DB {
	return rep.db.Find(out, where...)
}

// Exec exec given SQL using by gorm.DB.
func (rep *Repository) Exec(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Exec(sql, values...)
}

// First returns first record that match given conditions, order by primary key.
func (rep *Repository) First(out interface{}, where ...interface{}) *gorm.DB {
	return rep.db.First(out, where...)
}

// Raw returns the record that executed the given SQL using gorm.DB.
func (rep *Repository) Raw(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Raw(sql, values...)
}

// Create insert the value into database.
func (rep *Repository) Create(value interface{}) *gorm.DB {
	return rep.db.Create(value)
}

// Save update value in database, if the value doesn't have primary key, will insert it.
func (rep *Repository) Save(value interface{}) *gorm.DB {
	return rep.db.Save(value)
}

// Update update value in database
func (rep *Repository) Update(value interface{}) *gorm.DB {
	return rep.db.Update(value)
}

// Delete delete value match given conditions.
func (rep *Repository) Delete(value interface{}) *gorm.DB {
	return rep.db.Delete(value)
}

// Where returns a new relation.
func (rep *Repository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Where(query, args...)
}

// Preload preload associations with given conditions.
func (rep *Repository) Preload(column string, conditions ...interface{}) *gorm.DB {
	return rep.db.Preload(column, conditions...)
}

// Scopes pass current database connection to arguments `func(*DB) *DB`, which could be used to add conditions dynamically
func (rep *Repository) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return rep.db.Scopes(funcs...)
}

// ScanRows scan `*sql.Rows` to give struct
func (rep *Repository) ScanRows(rows *sql.Rows, result interface{}) error {
	return rep.db.ScanRows(rows, result)
}

// Transaction start a transaction as a block.
// If it is failed, will rollback and return error.
// If it is sccuessed, will commit.
// ref: https://github.com/jinzhu/gorm/blob/master/main.go#L533
func (rep *Repository) Transaction(fc func(tx *Repository) error) (err error) {
	panicked := true
	tx := rep.db.Begin()
	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	txrep := &Repository{}
	txrep.db = tx
	err = fc(txrep)

	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}
