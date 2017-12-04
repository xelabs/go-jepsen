/*
 * go-jepsen
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcommon

import (
	"fmt"
	"testing"

	"github.com/XeLabs/go-mysqlstack/driver"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	mysql, cleanup := MockMySQL()
	defer cleanup()
	conf := MockConf(mysql.Addr())

	client, err := driver.NewConn(conf.MysqlUser, conf.MysqlPassword, fmt.Sprintf("%s:%d", conf.MysqlHost, conf.MysqlPort), "", "utf8")
	assert.Nil(t, err)

	_, err = client.Query("drop table t1")
	assert.Nil(t, err)
}
