package test 

import (
	"context"
  "fmt"
	"testing"
	"github.com/Zumpit/domainfinder"
)

var ctx = context.Background()

func TestSearch(t *testing.T){
	q := "allied infoline"

	opts := domainfinder.SearchOptions{
		Limit: 20,
	}
	ret, err := domainfinder.Search(ctx, q,opts)
	if err != nil {
			t.Errorf("something went wrong : %v", err)
			return 
	}

	if len(ret) == 0 {
		t.Errorf("no results found : %v", ret)
	}	
  fmt.Println(ret)
}
