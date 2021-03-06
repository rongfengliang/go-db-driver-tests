package sqltest

import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

var (
	oracle = &oracleDB{driver: "goracle", connectionString: "oracle://user:pass@db/"}
)

type oracleDB database

func (o *oracleDB) DB() *sql.DB {
	return o.db
}

func (o *oracleDB) T() *testing.T {
	return o.t
}

func (o *oracleDB) RunTest(t *testing.T, fn func(Tester)) {
	//TODO
}

func (o *oracleDB) SQLBlobParam(size int) string {
	return fmt.Sprintf("RAW(%d)", size)
}

func (o *oracleDB) q(sql string) string {
	n := 0
	return qrx.ReplaceAllStringFunc(sql, func(string) string {
		n++
		return ":" + strconv.Itoa(n)
	})
}
