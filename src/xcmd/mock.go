package xcmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"
)

// MockInitFlags creates a new mock of init.
func MockInitFlags(cmd *cobra.Command, addr string) {
	host, sport, _ := net.SplitHostPort(addr)
	port, _ := strconv.Atoi(sport)

	cmd.Flags().String("mysql-host", host, "MySQL server host(Default NULL)")
	cmd.Flags().Int("mysql-port", port, "MySQL server port(Default 3306)")
	cmd.Flags().String("mysql-user", "mock", "MySQL user(Default mock)")
	cmd.Flags().String("mysql-password", "mock", "MySQL password(Default mock)")
}
