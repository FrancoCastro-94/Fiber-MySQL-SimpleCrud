package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // sql driver
)

func dbConnection() (db *sql.DB) {
	dbDriver := "mysql"  // Driver sql
	dbUser := "root"     // User db
	dbPass := "root"     // Pasword db
	dbName := "products" // Name of db
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db.Exec("create table if not exists `all_products` (`id` int(6) unsigned NOT NULL AUTO_INCREMENT, `name_product` varchar(30) NOT NULL , `product_description` varchar(300) NOT NULL, `image` varchar(300) NOT NULL, PRIMARY KEY (`id`))")
	if err != nil {
		panic(err.Error())
	}
	return db
}
