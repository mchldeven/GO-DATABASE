package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	
	ctx := context.Background()

	script := "INSERT INTO pegawai(nik, nama) VALUES('003','Dimas')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses memasukan pegawai baru")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	
	ctx := context.Background()

	script := "SELECT nik, nama FROM pegawai"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var nik, nama string
		err = rows.Scan(&nik, &nama)
		if err != nil {
			panic(err)
		}
		fmt.Println("nik:", nik)
		fmt.Println("nama:", nama)
	}
}