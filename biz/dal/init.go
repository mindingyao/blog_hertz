package dal

import (
	"blog_hertz/biz/dal/goredis"
	"blog_hertz/biz/dal/mysql"
)

func Init() {
	mysql.Init()
	goredis.Init()
}
