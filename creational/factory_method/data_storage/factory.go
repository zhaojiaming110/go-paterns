package data_storage

// 工厂类
type IOFactory struct {
}

type handerType string

const (
	XML  = "XMLhander"
	File = "Filehander"
	DB   = "DBhander"
)

// 工厂类方法实现
func (iof *IOFactory) GetIOHander(hander handerType) IOHander {

	var iohander IOHander
	switch hander {
	case XML:
		iohander = new(XMLhander)
		break
	case File:
		iohander = new(Filehander)
		break
	case DB:
		iohander = new(DBhander)
		break
	default:
		panic("no this kind iohander")
	}
	return iohander
}
