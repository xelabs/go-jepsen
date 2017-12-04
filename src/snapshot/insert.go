/*
 * go-jepsen
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
	"strings"
	"sync"
	"sync/atomic"
	"xcommon"
	"xworker"
)

// Insert tuple.
type Insert struct {
	stop     bool
	conf     *xcommon.Conf
	lock     sync.WaitGroup
	workers  []xworker.Worker
	requests uint64
}

// NewInsert creates new insert handler.
func NewInsert(conf *xcommon.Conf, workers []xworker.Worker) xworker.Handler {
	return &Insert{
		conf:    conf,
		workers: workers,
	}
}

// Run used to start the worker.
func (insert *Insert) Run() {
	threads := len(insert.workers)
	for i := 0; i < threads; i++ {
		insert.lock.Add(1)
		go insert.Insert(&insert.workers[i], threads, i)
	}
}

// Stop used to stop the worker.
func (insert *Insert) Stop() {
	insert.stop = true
	insert.lock.Wait()
}

// Rows returns the rows number inserted.
func (insert *Insert) Rows() uint64 {
	return atomic.LoadUint64(&insert.requests)
}

// Insert used to execute the insert operation.
func (insert *Insert) Insert(worker *xworker.Worker, num int, id int) {
	session := worker.S
	max := int(insert.conf.TablesSize)
	rows := make([]string, 0, 256)
	for i := 0; i < max; i++ {
		rows = append(rows, fmt.Sprintf("(%v, 0)", i))
		worker.M.WNums++
		atomic.AddUint64(&insert.requests, 1)
	}
	sql := fmt.Sprintf("INSERT INTO jepsen.jepsen_si(`id`, `score`) VALUES%s", strings.Join(rows, ","))
	if err := session.Exec(sql); err != nil {
		log.Panicf("insert.error[%v]", err)
	}
	insert.lock.Done()
}
