[![Build Status](https://travis-ci.org/XeLabs/go-jepsen.svg?branch=master)](https://travis-ci.org/XeLabs/go-jepsen) [![Go Report Card](https://goreportcard.com/badge/github.com/XeLabs/go-jepsen)](https://goreportcard.com/report/github.com/XeLabs/go-jepsen) [![codecov.io](https://codecov.io/gh/XeLabs/go-jepsen/graphs/badge.svg)](https://codecov.io/gh/XeLabs/go-jepsen/branch/master)

## About

go-jepsen is a toolset for distributed systems verification(MySQL protocol only), now supports:

* Snapshot Isolation Verification (SIV)

## Build

```
$git clone https://github.com/XeLabs/go-jepsen
$cd go-jepsen
$make
$./bin/jepsen -h
```

## Usage

```
$ ./bin/jepsen -h
A tool for distributed systems verification

Usage:
  jepsen [command]

Available Commands:
  prepare     prepare jepsen tables and datas
  cleanup     cleanup jepsen tables and datas
  snapshot    Snapshot Isolation Verification (SIV)

Flags:
      --max-request uint            limit for total requests, including write and read(Default 0, means no limits)
      --max-time int                limit for total execution time in seconds(Default 3600) (default 3600)
      --mysql-host string           MySQL server host(Default NULL)
      --mysql-password string       MySQL password(Default jepsen) (default "jepsen")
      --mysql-port int              MySQL server port(Default 3306) (default 3306)
      --mysql-table-engine string   storage engine to use for the jepsen table {tokudb,innodb,...}(Default innodb) (default "innodb")
      --mysql-user string           MySQL user(Default jepsen) (default "jepsen")
      --table-size int              The total number of the jepsen table(Default 10000) (default 10000)
```

## Examples

### Prepare

```
./bin/jepsen --mysql-host=192.168.0.3 --mysql-user=jepsen --mysql-password=jepsen prepare
prepare.create.the.table.jepsen_si(engine=innodb) ...
prepare.the.datas[10000].for.table.jepsen_si...
```

### Cleanup
```
./bin/jepsen --mysql-host=192.168.0.3 --mysql-user=jepsen --mysql-password=jepsen cleanup
```

### Snapshot Isolation Verification
```
./bin/jepsen --mysql-host=192.168.0.3 --mysql-user=jepsen --mysql-password=jepsen --max-time=60 snapshot

... ...

time       thds                 w-ops           r-ops           error(s)        total-ops
[57s]   [r:16,u:1]              20000           1010000         0               59640000

time       thds                 w-ops           r-ops           error(s)        total-ops
[58s]   [r:16,u:1]              30000           1130000         0               60800000

time       thds                 w-ops           r-ops           error(s)        total-ops
[59s]   [r:16,u:1]              30000           1140000         0               61970000

time       thds                 w-ops           r-ops           error(s)        total-ops
[60s]   [r:16,u:1]              30000           890000          0               62890000

the columns:
time:         verification uptime
thds:         read threads and update threads, here we use 16-threads for table scan and 1-thread for updating the whole table
w-ops:        the number of rows affected by an UPDATE query
r-ops:        the number of rows affected by an SELECT query
errors:       the number of verification failed
total-ops:    the total number of verification

Notes:
If the column of error(s) is not zero, the distributed systems does not satisfy its claims of Snapshot Isolation, reads are inconsistent.
To find out more about Snapshot Isolation, please visit [Jepsen: MariaDB Galera Cluster](https://aphyr.com/posts/327-jepsen-mariadb-galera-cluster)
```
