/*
 * go-jepsen
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
	threads          int
	mysqlHost        string
	mysqlPort        int
	mysqlUser        string
	mysqlPassword    string
	mysqlTableEngine string
	maxTime          int
	maxRequest      uint64
	tableSize       int
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
	rootCmd.PersistentFlags().StringVar(&mysqlHost, "mysql-host", "", "MySQL server host(Default NULL)")
	rootCmd.PersistentFlags().IntVar(&mysqlPort, "mysql-port", 3306, "MySQL server port(Default 3306)")
	rootCmd.PersistentFlags().StringVar(&mysqlUser, "mysql-user", "jepsen", "MySQL user(Default jepsen)")
	rootCmd.PersistentFlags().StringVar(&mysqlPassword, "mysql-password", "jepsen", "MySQL password(Default jepsen)")
	rootCmd.PersistentFlags().StringVar(&mysqlTableEngine, "mysql-table-engine", "innodb", "storage engine to use for the jepsen table {tokudb,innodb,...}(Default innodb)")
	rootCmd.PersistentFlags().IntVar(&maxTime, "max-time", 3600, "limit for total execution time in seconds(Default 3600)")
	rootCmd.PersistentFlags().Uint64Var(&maxRequest, "max-request", 0, "limit for total requests, including write and read(Default 0, means no limits)")
	rootCmd.PersistentFlags().IntVar(&tableSize, "table-size", 10000, "The total number of the jepsen table(Default 10000)")

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
