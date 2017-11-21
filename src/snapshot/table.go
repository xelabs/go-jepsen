/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package snapshot

import (
	"fmt"
	"log"
	"xworker"
)

type Table struct {
	workers []xworker.Worker
}

func NewTable(workers []xworker.Worker) *Table {
	return &Table{workers}
}

func (t *Table) Prepare() {
	session := t.workers[0].S
	engine := t.workers[0].E

	dbQuery := "CREATE DATABASE IF NOT EXISTS jepsen"
	if err := session.Exec(dbQuery); err != nil {
		log.Panicf("create.database.jepsen.error[%v]", err)
	}

	tableQuery := fmt.Sprintf(`CREATE TABLE jepsen.jepsen_si (
							id BIGINT(20) UNSIGNED NOT NULL,
							score BIGINT(20) DEFAULT 0,
							PRIMARY KEY (id)
							) ENGINE=%s PARTITION BY HASH(id)`, engine)

	log.Printf("prepare.create.the.table.jepsen_si(engine=%v) ...\n", engine)
	if err := session.Exec(tableQuery); err != nil {
		log.Panicf("creata.table[%s].error[%v]", tableQuery, err)
	}
}

func (t *Table) Cleanup() {
	session := t.workers[0].S
	sql := "DROP TABLE IF EXISTS jepsen.jepsen_si"
	if err := session.Exec(sql); err != nil {
		log.Panicf("drop.table.error[%v]", err)
	}
}
