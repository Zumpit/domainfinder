package domainfinder 

import (
	"errors"
	//"encoding/json"
	"fmt"
	"net"
	"regexp"
	"github.com/likexian/whois"
)


type ValidationResult struct {
	Dns_Exist   bool      `json:"dns_exist"`
    Syntax      bool      `json:"syntax"`
	  
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

// func DomainLookup(domain string) {
     
// }