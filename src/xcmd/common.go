/*
 * go-jepsen
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcmd

import (
	"snapshot"
	"time"
	"xcommon"
	"xworker"

	"github.com/spf13/cobra"
)

func parseConf(cmd *cobra.Command) (conf *xcommon.Conf, err error) {
	conf = &xcommon.Conf{}

	if conf.Mysql_host, err = cmd.Flags().GetString("mysql-host"); err != nil {
		return
	}

	if conf.Mysql_user, err = cmd.Flags().GetString("mysql-user"); err != nil {
		return
	}

	if conf.Mysql_password, err = cmd.Flags().GetString("mysql-password"); err != nil {
		return
	}

	if conf.Mysql_port, err = cmd.Flags().GetInt("mysql-port"); err != nil {
		return
	}

	if conf.Mysql_table_engine, err = cmd.Flags().GetString("mysql-table-engine"); err != nil {
		return
	}

	if conf.Tables_size, err = cmd.Flags().GetInt("table-size"); err != nil {
		return
	}

	if conf.Max_time, err = cmd.Flags().GetInt("max-time"); err != nil {
		return
	}

	if conf.Max_request, err = cmd.Flags().GetUint64("max-request"); err != nil {
		return
	}
	return
}

func startSnapshotTest(conf *xcommon.Conf) {
	snapshot := snapshot.NewSnapshot(conf)
	workers := snapshot.Workers()
	snapshot.Run()

	monitor := NewMonitor(conf, workers)
	monitor.Start()

	done := make(chan bool)
	go func(max uint64, ws ...xworker.Handler) {
		if max == 0 {
			return
		}

		var all uint64
		for {
			time.Sleep(time.Millisecond * 10)
			for _, w := range ws {
				all += w.Rows()
			}
			if all >= max {
				done <- true
			}
		}
	}(conf.Max_request, snapshot.Query(), snapshot.Update())

	select {
	case <-time.After(time.Duration(conf.Max_time) * time.Second):
	case <-done:
	}

	snapshot.Stop()
	monitor.Stop()
}
