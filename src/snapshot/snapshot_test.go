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

	"github.com/stretchr/testify/assert"
)

func TestSnapshot(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	snapshot := NewSnapshot(conf)
	snapshot.Run()
	snapshot.Stop()
	assert.NotNil(t, snapshot.Query())
	assert.NotNil(t, snapshot.Update())
	assert.NotNil(t, snapshot.Workers())
}
