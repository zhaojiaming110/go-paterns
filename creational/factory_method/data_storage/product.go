package data_storage

// 创建一个产品抽象类
type IOHander interface {
	Add(key, value string)
	Remove(key string)
	Update(key, value string)
	Query(key string)
}
