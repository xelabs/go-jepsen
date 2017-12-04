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
	"log"
	"sync"
	"sync/atomic"
	"time"
	"xcommon"
	"xworker"
)

// Query tuple.
type Query struct {
	stop     bool
	conf     *xcommon.Conf
	lock     sync.WaitGroup
	workers  []xworker.Worker
	requests uint64
}

// NewQuery creates a new query handler.
func NewQuery(conf *xcommon.Conf, workers []xworker.Worker) xworker.Handler {
	return &Query{
		conf:    conf,
		workers: workers,
	}
}

// Run used to start the worker.
func (q *Query) Run() {
	threads := len(q.workers)
	for i := 0; i < threads; i++ {
		q.lock.Add(1)
		go q.Query(&q.workers[i], threads, i)
	}
}

// Stop used to stop the worker.
func (q *Query) Stop() {
	q.stop = true
	q.lock.Wait()
}

// Rows returns the rows number queried.
func (q *Query) Rows() uint64 {
	return atomic.LoadUint64(&q.requests)
}

// Query used to query the whole table under snapshot isolation.
func (q *Query) Query(worker *xworker.Worker, num int, id int) {
	session := worker.S
	for !q.stop {
		t := time.Now()
		sql := "SELECT `score` FROM jepsen.jepsen_si WHERE 1 = 1"
		qr, err := session.FetchAll(sql, -1)
		if err != nil {
			log.Panicf("query.error[%v]", err)
		}

		var want string
		var got string
		for _, row := range qr.Rows {
			got = string(row[0].Raw())
			if want == "" {
				want = got
			}
			if want != got {
				log.Printf("+++++++++++want[%+v]!=got[%+v],sql:%+v\n", want, got, sql)
				worker.M.QErrs++
			}
			want = got
		}
		elapsed := time.Since(t)

		// stats
		nsec := uint64(elapsed.Nanoseconds())
		worker.M.QCosts += nsec
		if worker.M.QMax == 0 && worker.M.QMin == 0 {
			worker.M.QMax = nsec
			worker.M.QMin = nsec
		}
		if nsec > worker.M.QMax {
			worker.M.QMax = nsec
		}
		if nsec < worker.M.QMin {
			worker.M.QMin = nsec
		}

		worker.M.QNums += uint64(len(qr.Rows))
		atomic.AddUint64(&q.requests, 1)
	}
	q.lock.Done()
}
