// 创建具体产品类，实现业务逻辑
package data_storage

import "log"

type XMLhander struct {
}

func (xml *XMLhander) Add(key, value string) {
	log.Println("XMLhander:添加一条记录")
}
func (xml *XMLhander) Remove(kry string) {
	log.Println("XMLhander:删除一条记录")
}
func (xml *XMLhander) Update(key, value string) {
	log.Println("XMLhander:更新一条记录")
}
func (xml *XMLhander) Query(kry string) {
	log.Println("XMLhander:查询一条记录")
}

type Filehander struct {
}

func (file *Filehander) Add(key, value string) {
	log.Println("Filehander:添加一条记录")
}
func (file *Filehander) Remove(kry string) {
	log.Println("Filehander:删除一条记录")
}
func (file *Filehander) Update(key, value string) {
	log.Println("Filehander:更新一条记录")
}
func (file *Filehander) Query(kry string) {
	log.Println("Filehander:查询一条记录")
}

type DBhander struct {
}

func (db *DBhander) Add(key, value string) {
	log.Println("DBhander:添加一条记录")
}
func (db *DBhander) Remove(kry string) {
	log.Println("DBhander:删除一条记录")
}
func (db *DBhander) Update(key, value string) {
	log.Println("DBhander:更新一条记录")
}
func (db *DBhander) Query(kry string) {
	log.Println("DBhander:查询一条记录")
}
