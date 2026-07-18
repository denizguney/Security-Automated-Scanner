
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
