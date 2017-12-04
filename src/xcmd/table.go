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
	"fmt"
	"snapshot"
	"xworker"

	"github.com/spf13/cobra"
)

// NewPrepareCommand creates new cmd.
func NewPrepareCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare",
		Short: "prepare jepsen tables and datas",
		Run:   prepareCommandFn,
	}

	return cmd
}

func prepareCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}

	workers := xworker.CreateWorkers(conf, 1)
	table := snapshot.NewTable(workers)

	// prepare tables.
	table.Prepare()

	// prepare datas.
	fmt.Printf("prepare.the.datas[%d].for.table.jepsen_si...\n", conf.TablesSize)
	iworkers := xworker.CreateWorkers(conf, 1)
	insert := snapshot.NewInsert(conf, iworkers)
	insert.Run()
	insert.Stop()
}

// NewCleanupCommand creates new cmd.
func NewCleanupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cleanup",
		Short: "cleanup jepsen tables and datas",
		Run:   cleanupCommandFn,
	}

	return cmd
}

func cleanupCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}

	// worker
	workers := xworker.CreateWorkers(conf, 1)
	table := snapshot.NewTable(workers)
	table.Cleanup()
}
