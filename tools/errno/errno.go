package errno

import "errors"


var (
	NotFoundTheKeyErr = errors.New("not find the key") // 无法找到这个KEY
	NotMatchOperateTypeErr = errors.New("cannot match the type of value") // 不能对实例类型做本操作
	UnImplementErr = errors.New("the method isn't implemented")
)
