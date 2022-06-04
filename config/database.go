package config

type Database struct {
	User            string //用户名
	Password        string //密码
	Addr            string //ip
	Port            int    //端口
	DbName          string //数据库名称
	Timeout         int
	ConnMaxIdleTime int
	ConnMaxLifeTime int //设置了连接可复用的最大时间
	MaxIdleConns    int //设置空闲连接池中连接的最大数量
	MaxOpenConns    int //设置打开数据库连接的最大数量。

}

var DatabaseConfig = new(Database)
