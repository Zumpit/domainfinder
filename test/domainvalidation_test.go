package test 

import (
	"testing"

	"github.com/Zumpit/domainfinder"
)

func TestDomainValidation(t *testing.T){
	domain := "zumpitech.com"
	ret, err := domainfinder.DomainValidation(domain)
	if err != nil {
		t.Errorf("something went wrong : %v", err.Error())
		return 
	}

	if len(ret) == 0 {
		t.Errorf("nothing in the response :%v", ret)
		return
	}
	
}