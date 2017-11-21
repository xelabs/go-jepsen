/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcmd

import (
	"github.com/spf13/cobra"
)

func NewSnapshotCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Run:   snapshotCommandFn,
		Short: "Snapshot Isolation Verification (SIV)",
	}

	return cmd
}

func snapshotCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}
	startSnapshotTest(conf)
}
