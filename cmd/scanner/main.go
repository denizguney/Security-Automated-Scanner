package main

import(
	"fmt"
	"Security-Automated-Scanner/internal/db"
	"Security-Automated-Scanner/pkg/recon"
	Security-Automated-Scanner/pkg/reporter"
	)
func main(){

	fmt.Println("--Security Automated Scanner Starting --")

	//1.Tarayacğımız hedefi belirle 
	target := "example.com"

	//2.Scanner nesnesini oluştur
	scanner := recon.SubdomainScanner{

		Target: target,
	}

	//3. Tarama işlemini başlat 
	fmt.Println("Scanning target: ", target)
	scanner.PerformScan()
	fmt.Println("----Scan Completed ---")

	//4.Tarama bittikten sonra verileri çek ve raporla 
	allFindings := db.GetAllFindings()

	//5.Raporu görüntüle ve kaydet

	reporter.DisplayFindings(allFindings)
	reporter.ExportToMarkdown(allFindings, "scan_report.md")
	
}
