/*
 * go-jepsen
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package snapshot

import (
	"testing"
	"xcommon"
	"xworker"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	conf.TablesSize = 10
	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)

	table := NewTable(workers)
	table.Prepare()
	table.Cleanup()
}
