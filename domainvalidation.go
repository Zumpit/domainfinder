package domainfinder 

import (
	"errors"
	//"encoding/json"
	"fmt"
	"net"
	"net/http"
	"io/ioutil"
	"regexp"
)

type Subdomains struct {
	Domain     string     `json:"domain"`     
}

type ValidationResult struct {
	Dns_Exist   bool      `json:"dns_exist"`
    Syntax      bool      `json:"syntax"`
	Subdomains   
}

var domainRegexp = regexp.MustCompile(`^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+\.?$`)


func IsValidDomain(domain string) bool {
	return domainRegexp.MatchString(domain)
}
  
var ErrInvalidaDomain = errors.New("invalid doamin")
  
  //dns lookup  
func ValidateDomainByResolvingIt(domain string) bool {
	if !IsValidDomain(domain){
	  fmt.Println("Invlid domain reson : ",ErrInvalidaDomain)
	  return false     
	}
  
	res, err := net.LookupHost(domain)
	if err!=nil{
	  fmt.Println("Error : ",err)
	  return false 
	}
  
	if len(res) == 0 {
	  fmt.Println(ErrInvalidaDomain)
	  return false 
	}
	fmt.Println("Response :", res)
	return true
}

func DomainValidation(domain string) ([]ValidationResult, error){
	dns_exist := ValidateDomainByResolvingIt(domain)
	syntax := IsValidDomain(domain)
	fmt.Println(dns_exist) 
	fmt.Println(syntax)
	res, err := http.Get(fmt.Sprintf("https://sonar.omnisint.io/subdomains/%s" ,domain)) 
    if err != nil {
		fmt.Println(nil)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Body : ",body)
	
	result := []ValidationResult{}
	
	// result.Dns_Exist = dns_exist
	// result.Syntax = syntax 
	// result.Objects = f 

    con := ValidationResult {
		Dns_Exist: dns_exist,
		Syntax : syntax,
	}
    bytes := append(result, con) 

	return bytes, nil 
}