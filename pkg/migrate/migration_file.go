package migrate

//getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
	for _, mFile := range migrationFiles {
		if name == mFile.FileName {
			return mFile
		}
	}
	return MigrationFile{}
}
