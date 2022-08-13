package migrate

import (
	"go-simple/pkg/console"
	"go-simple/pkg/database"
	"gorm.io/gorm"
	"io/ioutil"
)

type Migrator struct {
	Folder string
	DB *gorm.DB
	Migrator gorm.Migrator
}

type Migration struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch int
}

func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder: "database/migrations/",
		DB: database.DB,
		Migrator: database.DB.Migrator(),
	}
	migrator.createMigrationsTable()
	return migrator
}

func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

func (migrator *Migrator) Up() {
	migrateFiles := migrator.readAllMigrationFiles()
	batch := migrator.getBatch()
	var migrations []Migration
	migrator.DB.Find(&migrations)
	isRun := false
	for _, mFile := range migrateFiles {
		if mFile.isNotMigrated(migrations) {
			migrator.runUpMigration(mFile, batch)
			isRun = true
		}
	}
	if !isRun {
		console.Success("database is up to date")
	}
}

func (migrator *Migrator) Rollback() {
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	var migrations []Migration
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	isRun := false
	for _, _migration := range migrations{
		console.Warning("rollback " + _migration.Migration)
		mFile := getMigrationFile(_migration.Migration)
		if mFile.Down != nil {
			mFile.Down(mFile.DB.DB.Migrator(), mFile.DB.SqlDB)
		}
		isRun = true
		// Delete Migration Log
		migrator.DB.Delete(&_migration)
		console.Success("finsh " + mFile.FileName)
	}
	return isRun
}

func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	files, err := ioutil.ReadDir(migrator.Folder)
	console.ExitIf(err)
	var migrateFiles []MigrationFile
	for _, fi := range files{
		fileName := FileNameWithoutExtension(fi.Name())
		mFile := getMigrationFile(fileName)
		if len(mFile.FileName) > 0{
			migrateFiles = append(migrateFiles, mFile)
		}
	}
	return migrateFiles
}

func (migrator *Migrator) getBatch() int {
	// 默认为 1
	batch := 1

	// 取最后执行的一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	// 如果有值的话，加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

func (migrator *Migrator) runUpMigration(mFile MigrationFile, batch int) {
	if mFile.Up != nil {
		console.Warning("begin migrating " + mFile.FileName)
		mFile.Up(mFile.DB.DB.Migrator(), mFile.DB.SqlDB)
		console.Success("migrated " + mFile.FileName)
	}
	err := migrator.DB.Create(&Migration{Migration: mFile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}

func (migrator *Migrator) Reset() {
	var migrations []Migration
	migrator.DB.Order("id DESC").Find(&migrations)
	// Rollback All Migrations
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to reset.")
	}
}

func (migrator *Migrator) Refresh() {
	migrator.Reset()
	migrator.Up()
}

//func (migrator Migrator) Fresh() {
//	dbname := database.CurrentDatabase()
//	err := database.DeleteAllTables()
//	console.ExitIf(err)
//	console.Success("clearup database " + dbname)
//	migrator.createMigrationsTable()
//	console.Success("[migrations] table created.")
//	migrator.Up()
//}
