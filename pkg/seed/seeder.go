//处理数据库填充相关逻辑
package seed

import "gorm.io/gorm"

//存放所有 Seeder
var seeders []Seeder

// 按照顺序执行的Seeder数组
// 支持一些必须按顺序执行的 seeder，例如 topic 创建的
// 时必须依赖于 user, 所以 TopicSeeder 应该在 UserSeeder 后执行
var orderedSeederNames []string

type SeederFunc func(db *gorm.DB)

// Seeder 对应每一个 database/seeders目录下的Seeder文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 注册到seeders数组中
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

//SetRunOrder 设置【按顺序执行】
func SetRunOrder(names []string) {
	orderedSeederNames = names
}
