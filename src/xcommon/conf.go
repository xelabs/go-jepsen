/*
 * go-jepsen
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcommon

// Conf tuple.
type Conf struct {
	QThreads           int
	UThreads           int
	MysqlHost         string
	MysqlUser         string
	MysqlPassword     string
	MysqlPort         int
	MysqlTableEngine string
	MaxTime           int
	MaxRequest        uint64
	TablesSize        int
}
