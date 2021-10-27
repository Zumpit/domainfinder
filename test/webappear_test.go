package test

import (
     "context"
     "testing"
     "github.com/Zumpit/domainfinder"

)

var ctx = context.Background()

func WebAppear_Test(t *testing.T){
     q := "postman inc"

     opts := domainfinder.SearchOptions{
        Limit : 20,
     }
     ret, err := domainfinder.WebAppear(ctx,q,opts)
     if err != nil {
         t.Errorf("testing returning %v", err)
         return
     }

     if len(ret) == 0 {
         t.Errorf("testing retruning %v", ret)
     }
}
