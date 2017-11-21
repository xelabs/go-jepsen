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

func TestUpdate(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	conf.Max_request = 1
	workers := xworker.CreateWorkers(conf, 1)
	assert.NotNil(t, workers)

	update := NewUpdate(conf, workers)
	update.Run()
	time.Sleep(500)
	update.Stop()

	want := 1
	got := int(update.Rows())
	assert.Equal(t, want, got)
}
