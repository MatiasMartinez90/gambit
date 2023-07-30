package bd

import (
	"database/sql"
	"fmt"

	//"strconv"
	"github.com/gambit/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Inicializando funcion  db.InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_Name, Categ_Path ) Values('" + c.CategName + "','" + c.CategPath + "')"

	fmt.Println("ejecutnado sentencia")
	fmt.Println(sentencia)
	var result sql.Result

	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println(err.Error())
		return 0, err2
	}

	fmt.Println("Insert Category > Ejecucion Exitosa")
	return LastInsertId, nil

}
