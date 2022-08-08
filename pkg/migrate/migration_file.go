package migrate

import (
	"database/sql"
	"gorm.io/gorm"
	"path/filepath"
	"strings"
)

type migrationFunc func(migrator gorm.Migrator, db *sql.DB)

// MigrationFile 代表着单个迁移文件
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

var migrationFiles []MigrationFile

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		Up: up,
		Down: down,
		FileName: name,
	})
}

func getMigrationFile(name string) MigrationFile {
	for _, mFile := range migrationFiles{
		if name == mFile.FileName {
			return mFile
		}
	}
	return MigrationFile{}
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func (mFile MigrationFile)isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mFile.FileName {
			return false
		}
	}
	return true
}