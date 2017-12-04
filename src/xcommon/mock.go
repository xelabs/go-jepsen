package xcommon

import (
	"net"
	"strconv"

	"github.com/XeLabs/go-mysqlstack/driver"
	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/XeLabs/go-mysqlstack/xlog"
)

// MockConf creates a new mock of conf.
func MockConf(addr string) *Conf {
	host, sport, _ := net.SplitHostPort(addr)
	port, _ := strconv.Atoi(sport)

	return &Conf{
		MysqlHost:  host,
		MysqlPort:  port,
		MysqlUser:  "mock",
		MaxRequest: 16,
	}
}

// MockMySQL creates a new mysql mock.
func MockMySQL() (*driver.Listener, func()) {
	result1 := &sqltypes.Result{}
	result2 := &sqltypes.Result{
		RowsAffected: 2,
		Fields: []*querypb.Field{
			{
				Name: "id",
				Type: querypb.Type_INT64,
			},
			{
				Name: "score",
				Type: querypb.Type_INT64,
			},
		},
		Rows: [][]sqltypes.Value{
			{
				sqltypes.MakeTrusted(querypb.Type_INT64, []byte("1")),
				sqltypes.MakeTrusted(querypb.Type_INT64, []byte("1")),
			},
			{
				sqltypes.MakeTrusted(querypb.Type_INT64, []byte("2")),
				sqltypes.MakeTrusted(querypb.Type_INT64, []byte("2")),
			},
		},
	}

	log := xlog.NewStdLog(xlog.Level(xlog.ERROR))
	th := driver.NewTestHandler(log)
	svr, err := driver.MockMysqlServer(log, th)
	if err != nil {
		log.Panicf("mock.mysql.error:%+v", err)
	}

	// Querys.
	th.AddQueryPattern("insert .*", result1)
	th.AddQueryPattern("update .*", result1)
	th.AddQueryPattern("select `score` .*", result2)

	th.AddQueryPattern("create .*", result1)
	th.AddQueryPattern("drop .*", result1)

	return svr, func() {
		svr.Close()
	}
}
