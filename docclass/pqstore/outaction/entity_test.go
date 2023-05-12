package outaction_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-batis/sqlmapper"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/pqstore/outaction"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-batis/sqllks"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

const (
	SQLHOSTENVVAR   = "SQL_HOST"
	SQLDBENVVAR     = "SQL_DB"
	SQLUSERENVVAR   = "SQL_USER"
	SQLPASSWDENVVAR = "SQL_PASSWD"
	SQLPORTENVVAR   = "SQL_PORT"
)

func TestMain(m *testing.M) {

	p, err := strconv.Atoi(os.Getenv(SQLPORTENVVAR))
	if err != nil {
		panic(err)
	}

	cfg := sqllks.Config{
		ServerName:      "default",
		ServerType:      "pq",
		Host:            os.Getenv(SQLHOSTENVVAR),
		Port:            p,
		DbName:          os.Getenv(SQLDBENVVAR),
		UserName:        os.Getenv(SQLUSERENVVAR),
		Password:        os.Getenv(SQLPASSWDENVVAR),
		SslMode:         false,
		EnableMigration: false,
		MaxOpenConns:    0,
		MaxIdleConns:    0,
		ConnMaxLifetime: 0,
		ConnMaxIdleTime: 0,
	}

	_, err = sqllks.Initialize([]sqllks.Config{cfg})
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

const doDDL = false

func TestOutActionEntity(t *testing.T) {
	lks, err := sqllks.GetLinkedService("default")
	require.NoError(t, err)

	sqlDb, err := lks.DB()
	require.NoError(t, err)

	if doDDL {
		t.Log("ddl execution")
		sqlDb.MustExec(outaction.OutActionEntityTableDDL)
		defer sqlDb.MustExec(outaction.OutActionEntityTableDropDDL)
	}

	t.Log("select statement")
	f := sqlmapper.NewFilterBuilder().Limit(2).Offset(0)
	l, err := outaction.Select(sqlDb, f.Build())
	require.NoError(t, err)
	for i, e := range l {
		t.Logf("row: %d - %v", i, e)
	}
	t.Log("select-by-primary-key statement")
	e, err := outaction.SelectByPrimaryKey(sqlDb, outaction.PrimaryKey{ClassId: outaction.MustToMax30Text("pt_contratti"), ActionId: 1})
	require.NoError(t, err)
	t.Log(e)

}
