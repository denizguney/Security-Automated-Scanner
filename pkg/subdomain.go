package recon
import(
	"fmt"
	"net"
)

//SubdomainScanner , hedeflenen alan adı için tarama işlemlerini yönetir 
type SubdomainScanner struct{

	Target string
}

//PerformScan verilen target üzerinde temel bir DNS sorgusu yapar
func (s *SubdomainScanner) PerformScan(){

	fmt.Printf("Scanning target: %s\n" , s.Target)

	//Basit bir test: DNS lookup denemesi

	ips, err:= net.LookupIP(s.Target)
	if err != nil{

		fmt.Printf("Could not get IPs: %v\n", err)
	}
	
	for _, ip := range ips{

		fmt.Printf("Found IP: %s\n", ip.string()
	}
}
