package main

import (

	"fmt"
        "time"
        "context"

        "goshop/internal/app/repositories"
        "goshop/internal/app/serializers"

)

func main() {

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        productRepo := repositories.NewProductRepository()
        orderRepo := repositories.NewOrderRepository()

	fmt.Printf("%#v\n",2345)

/*
app/serializers/user.go
type User struct {
        ID        string    `json:"id"`
        Email     string    `json:"email"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
}


app/serializers/product.go:20:type ListProductReq struct {
type Product struct {
        ID          string    `json:"id"`
        Code        string    `json:"code"`
        Name        string    `json:"name"`
        Description string    `json:"description"`
        Price       float64   `json:"price"`
        Active      bool      `json:"active"`
        CreatedAt   time.Time `json:"created_at"`
        UpdatedAt   time.Time `json:"updated_at"`
}

type ListProductReq struct {
        Name      string `json:"name,omitempty" form:"name"`
        Code      string `json:"code,omitempty" form:"code"`
        Page      int64  `json:"-" form:"page"`
        Limit     int64  `json:"-" form:"limit"`
        OrderBy   string `json:"-" form:"order_by"`
        OrderDesc bool   `json:"-" form:"order_desc"`
}

type ListProductRes struct {
        Products   []*Product         `json:"products"`
        Pagination *paging.Pagination `json:"pagination"`
}


func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
        ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)

func (u *UserRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
        ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)

*/


	products, pagination, err := productRepo.ListProducts(ctx,&serializers.ListProductReq{})
        fmt.Printf("%#v\n",productRepo)
        fmt.Printf("%#v\n",products)
        fmt.Printf("%#v\n",pagination)
        fmt.Printf("%#v\n",err)

/*
func (r *ProductRepo) ListProducts(ctx context.Context, req *serializers.ListProductReq) ([]*models.Product, *paging.Pagination, error) {
        ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)

*/

        order,pagination, err := orderRepo.GetMyOrders(ctx,&serializers.ListOrderReq{})
        fmt.Printf("%#v\n", orderRepo)
        fmt.Printf("%#v\n", order)
        fmt.Printf("%#v\n",pagination)
        fmt.Printf("%#v\n",err)

/*

type Order struct {
        ID         string       `json:"id"`
        Code       string       `json:"code"`
        Lines      []*OrderLine `json:"lines"`
        TotalPrice float64      `json:"total_price"`
        Status     string       `json:"status"`
}


app/serializers/order.go

type ListOrderReq struct {
        UserID    string `json:"-"`
        Code      string `json:"code,omitempty" form:"code"`
        Status    string `json:"status,omitempty" form:"status"`
        Page      int64  `json:"-" form:"page"`
        Limit     int64  `json:"-" form:"limit"`
        OrderBy   string `json:"-" form:"order_by"`
        OrderDesc bool   `json:"-" form:"order_desc"`
}

func (r *OrderRepo) GetMyOrders(ctx context.Context, req *serializers.ListOrderReq) ([]*models.Order, *paging.Pagination, error) {
        ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)


*/
}
