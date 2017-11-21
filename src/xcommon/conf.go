/*
 * go-jepsen
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcommon

type Conf struct {
	QThreads           int
	UThreads           int
	Mysql_host         string
	Mysql_user         string
	Mysql_password     string
	Mysql_port         int
	Mysql_table_engine string
	Max_time           int
	Max_request        uint64
	Tables_size        int
}
