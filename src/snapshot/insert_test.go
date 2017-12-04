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
	"time"
	"xcommon"
	"xworker"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	conf.TablesSize = 10
	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)

	insert := NewInsert(conf, workers)
	insert.Run()
	time.Sleep(500)
	insert.Stop()

	want := 20
	got := int(insert.Rows())
	assert.Equal(t, want, got)
}
