// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/8 上午11:57

package singleton

// 单例类型
type singleton struct {
	count int
}

// 创建唯一的实例
var instance *singleton

// 对外提供该对象的唯一实例
func GetInstance() *singleton {
	if instance == nil {
		// 初始化
		instance = new(singleton)
	}

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return  s.count
}

func (s *singleton) GetCount() int {
	return s.count
}