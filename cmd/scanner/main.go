
package main

import(
	"fmt"
	"Security-Automated-Scanner/pkg/recon" //Repo ismine göre import yolunu kontrol et 
)

func main(){

	fmt.Println("--- Security Automated Scanner Starting ---")

	//Tarayacağımız hedefi belirle 

	target: = "example.com"

	//Scanner nesnesini oluştur

	Scanner := recon.SubdomainScanner{

		Target: target,
	}

	//Tarama işlemini başlat
	Scanner.PerformScan()

	fmt.Println("---Scan Completed ---")
}
import(
	"your-module-path/pkg/reporter" //kendi modül isminle güncelle
	"your-module-path/internal/db"
	)

//tarama kodları

//Tarama bittikten sonra:
allFindings := db.GetAllFindings() //db paketinden verileri çeken fonksiyon
reporter.DisplayFindings(allFindings)

//Otomatik Markdown raporu 
reporter.ExportToMarkdown(allFindings, "scan_report.md")
