/*
 * benchyou
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

func TestQuery(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	conf.Max_request = 1
	workers := xworker.CreateWorkers(conf, 1)
	assert.NotNil(t, workers)

	query := NewQuery(conf, workers)
	query.Run()
	time.Sleep(500)
	query.Stop()

	want := 1
	got := int(query.Rows())
	assert.Equal(t, want, got)
}
