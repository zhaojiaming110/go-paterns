package main

import (
	data "go-paterns/creational/factory_method/data_storage"
)

func main() {
	iofactory := new(data.IOFactory)
	iohander := iofactory.GetIOHander(data.XML)
	iohander.Add("hhh", "hhh")
	iohander.Remove("sds")
	iohander.Update("ddd", "fff")
	iohander.Query("dsds")
	iohander = iofactory.GetIOHander(data.DB)
	iohander.Add("hhh", "hhh")
	iohander.Remove("sds")
	iohander.Update("ddd", "fff")
	iohander.Query("dsds")
}
