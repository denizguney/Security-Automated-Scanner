package db

import(

	"context"
	"database/sql"
	"fmt"
	"github.com/denisenkom/go/mssqldb"
)

func SaveFinding(target string , ip string){

	//Kendi SQL bağlantı bilgilerini buraya gir

	connString := "server=localhost;user id=SA;password=SeninSifren;database=BugBountyDB"
	db, err:= sql.Open("sqlserver", connString)

	if err != nil {

		fmt.Println("Bağlantı hatası:", err)

		return
	}

	defer db.Close()

	query := "INSERT INTO Findings (Target , IP) VALUES (? , ? )"
	_, err = db.ExecContext(context.Background(), query , target ,ip)

	if err != nill{

		fmt.Println("Veri ekleme hatası: ", err)
	}else{

		fmt.Println("Veri tabanına kaydedildi " , target)
	}
}