package main

import (

	"fmt"
        "time"
        "context"
	"encoding/json"
        "goshop/internal/app/repositories"
        "goshop/internal/app/serializers"
	"goshop/internal/app/dbs"
)

func main() {

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

	dbs.Init()

        productRepo := repositories.NewProductRepository()


	products, pagination, err := productRepo.ListProducts(ctx,&serializers.ListProductReq{})
        fmt.Printf("%#v\n",productRepo)
        fmt.Printf("%#v\n",products)
	for k,v := range products {
		fmt.Printf("%#v:%#v\n", k, v)
	}
        fmt.Printf("%#v\n",pagination)
        fmt.Printf("%#v\n",err)


	b, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
