package domainfinder 

import (
	"errors"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	
	
	
)

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

func DomainValidation(domain string) ([]string, error){
	term := ValidateDomainByResolvingIt(domain)
	fmt.Println(term) 
	res, err := http.Get("https://sonar.omnisint.io/subdomains/%s", domain) 
    if err != nil {
		fmt.Println(nil)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	var values []string 

	if err := json.Unmarshal([]byte(body), &values); err != nil {
		panic(err)
	}
	f := make([]string,0)
	for _, item := range values {
		f = append(f, item)
	}
	return f, nil 
}