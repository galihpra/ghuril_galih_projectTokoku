package products

import (
	"fmt"
	"project_tokoku/model"

	"gorm.io/gorm"
)

type ProductSystem struct {
	DB *gorm.DB
}

func (ps *ProductSystem) CreateProduct(userID string) (model.Product, bool) {
	var newProduct = new(model.Product)
	fmt.Print("Masukkan Barcode Produk: ")
	fmt.Scanln(&newProduct.Barcode)
	fmt.Print("Masukkan Nama Produk: ")
	fmt.Scanln(&newProduct.Nama)
	fmt.Print("Masukkan Harga Produk: ")
	fmt.Scanln(&newProduct.Harga)
	fmt.Print("Masukkan Stok Produk: ")
	fmt.Scanln(&newProduct.Stok)
	newProduct.UserID = userID

	err := ps.DB.Create(newProduct).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Product{}, false
	}

	return *newProduct, true
}

func (ps *ProductSystem) ReadProducts() ([]model.Product, bool) {
	var listProduk []model.Product

	err := ps.DB.Model(&model.Product{}).
		Select("products.*, users.nama").
		Joins("JOIN users on products.user_id = users.username").
		Scan(&listProduk).
		Error

	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, false
	}

	return listProduk, true
}