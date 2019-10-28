package dbclient

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // 导入数据库驱动
    "myService/crawlads/adsservice/conf"
    "strings"
    "myService/crawlads/adsservice/model"
)

func init() {
    //构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
    path := strings.Join([]string{config.MQ_USERNAME, ":", config.MQ_PASSWD, "@tcp(", config.MQ_HOST, ":", config.MQ_PORT, ")/", config.MQ_DATABASE, "?charset=utf8"}, "")

    // 设置默认数据库
    orm.RegisterDataBase("default", "mysql", path, 30)
    
    // 注册定义的 model
    orm.RegisterModel(new(model.Book), new(model.BookScore))

    // 创建 table
    orm.RunSyncdb("default", false, true)
}

func DbCheck() bool {
    _, err := orm.GetDB()
    if err != nil {
        return false
    } else {
        return true
    }
}