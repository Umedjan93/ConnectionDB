/* 	данный файл служит для инициализации таблиц в pgx и дальнейшего их создания в зависимости
от структур таблиц в файле ddl.go*/

package database

import (
	"fmt"
	"github.com/jackc/pgx"
)

func DBinitialisation(db *pgx.Conn) {
	DDLs := []string{ClientsDataDDL, MerchantsDataDDL, ManagersDataDDL, ClientAccountsDDL,
		MerchantAccountsDDL, ServicesDDL, TransactionsDDL, ATMsDDL}
	for _, ddl := range DDLs {
		_,  err := db.Exec(ddl)
		if err != nil {
			fmt.Printf("cant create a table %e", err)
		}
	}
}

