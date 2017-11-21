/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package main

import (
	"fmt"
	"os"
	"runtime"
	"xcmd"

	"github.com/spf13/cobra"
)

var (
	threads            int
	mysql_host         string
	mysql_port         int
	mysql_user         string
	mysql_password     string
	mysql_table_engine string
	max_time           int
	max_request        uint64
	table_size         int
)

var (
	rootCmd = &cobra.Command{
		Use:        "jepsen",
		Short:      "A tool for distributed systems verification",
		SuggestFor: []string{"jepsen"},
	}
)

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.PersistentFlags().StringVar(&mysql_host, "mysql-host", "", "MySQL server host(Default NULL)")
	rootCmd.PersistentFlags().IntVar(&mysql_port, "mysql-port", 3306, "MySQL server port(Default 3306)")
	rootCmd.PersistentFlags().StringVar(&mysql_user, "mysql-user", "jepsen", "MySQL user(Default jepsen)")
	rootCmd.PersistentFlags().StringVar(&mysql_password, "mysql-password", "jepsen", "MySQL password(Default jepsen)")
	rootCmd.PersistentFlags().StringVar(&mysql_table_engine, "mysql-table-engine", "innodb", "storage engine to use for the jepsen table {tokudb,innodb,...}(Default innodb)")
	rootCmd.PersistentFlags().IntVar(&max_time, "max-time", 3600, "limit for total execution time in seconds(Default 3600)")
	rootCmd.PersistentFlags().Uint64Var(&max_request, "max-request", 0, "limit for total requests, including write and read(Default 0, means no limits)")
	rootCmd.PersistentFlags().IntVar(&table_size, "table-size", 10000, "The total number of the jepsen table(Default 10000)")

	rootCmd.AddCommand(xcmd.NewPrepareCommand())
	rootCmd.AddCommand(xcmd.NewCleanupCommand())
	rootCmd.AddCommand(xcmd.NewSnapshotCommand())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
