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
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"xcommon"
	"xworker"
)

// Update tuple.
type Update struct {
	stop     bool
	conf     *xcommon.Conf
	lock     sync.WaitGroup
	workers  []xworker.Worker
	requests uint64
}

// NewUpdate creates the new update handler.
func NewUpdate(conf *xcommon.Conf, workers []xworker.Worker) xworker.Handler {
	return &Update{
		conf:    conf,
		workers: workers,
	}
}

// Run used to start the worker.
func (update *Update) Run() {
	threads := len(update.workers)
	for i := 0; i < threads; i++ {
		update.lock.Add(1)
		go update.Update(&update.workers[i], threads, i)
	}
}

// Stop used to stop the worker.
func (update *Update) Stop() {
	update.stop = true
	update.lock.Wait()
}

// Rows returns the rows number updated.
func (update *Update) Rows() uint64 {
	return atomic.LoadUint64(&update.requests)
}

// Update used to update the whole table under snapshot isolation.
func (update *Update) Update(worker *xworker.Worker, num int, id int) {
	session := worker.S
	for !update.stop {
		t := time.Now()
		rnd := rand.Int31n(20171111)
		sql := fmt.Sprintf("UPDATE jepsen.jepsen_si SET `score`=%v WHERE 1 = 1", rnd)
		if err := session.Exec(sql); err != nil {
			log.Panicf("update.error[%v]", err)
		}
		elapsed := time.Since(t)

		nsec := uint64(elapsed.Nanoseconds())
		worker.M.WCosts += nsec
		if worker.M.WMax == 0 && worker.M.WMin == 0 {
			worker.M.WMax = nsec
			worker.M.WMin = nsec
		}

		if nsec > worker.M.WMax {
			worker.M.WMax = nsec
		}
		if nsec < worker.M.WMin {
			worker.M.WMin = nsec
		}
		worker.M.WNums += uint64(update.conf.TablesSize)
		atomic.AddUint64(&update.requests, 1)
	}
	update.lock.Done()
}
