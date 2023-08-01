package product

import (
	"projects_template/internal/entity/product"
	"projects_template/internal/entity/stock"
	"projects_template/internal/transaction"
)

type ProductUseCase interface {
	AddProduct(ts transaction.Session, product product.Product) (productID int, err error)
	AddProductPrice(ts transaction.Session, pr product.ProductPrice) (int, error)
	AddProductInStock(ts transaction.Session, p stock.AddProductInStock) (int, error)
	FindProductInfoById(ts transaction.Session, productID int) (product.ProductInfo, error)
	FindProductList(ts transaction.Session, tag string, limit int) ([]product.ProductInfo, error)
	FindProductsInStock(ts transaction.Session, productID int) ([]stock.Stock, error)
	Buy(ts transaction.Session, p product.Sale) (int, error)
	FindSales(ts transaction.Session, sq product.SaleQuery) ([]product.Sale, error)
}
