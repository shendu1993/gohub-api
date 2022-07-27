//处理数据库填充相关逻辑
package seed

import (
	"gohub-api/pkg/console"
	"gohub-api/pkg/database"

	"gorm.io/gorm"
)

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

//GetSeeder 通过名称来获取 Seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

//RunALl 运行所有的Seeder
func RunAll() {
	//先运行ordered 的
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running Odered Seeder:" + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}
	//在运行剩下的
	for _, sdr := range seeders {
		//过滤已运行
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running Seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

//RunSeeder运行单个Seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if name == sdr.Name {
			sdr.Func(database.DB)
		}
	}
}
