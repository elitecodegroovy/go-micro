package plugin

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/micro/go-micro/v2/apps/json/plugins/db"
	_ "github.com/micro/go-micro/v2/apps/json/plugins/redis"
)
