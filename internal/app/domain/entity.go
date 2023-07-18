package domain

import (
	"go-back/internal/tools/sqlnull"
	"time"

	"github.com/shopspring/decimal"
)

// Product - cтруктура продукта для записи в базу
type Product struct {
	ProductID  int              `json:"product_id"`  // Id продукта
	Name       string           `json:"name"`        // Название продукта
	Descr      string           `json:"description"` // Описание продукта
	Addet_at   time.Time        `json:"added_at"`    //Дата добавления продукта
	Removed_at sqlnull.NullTime `json:"removed_at"`  // дата удаления продукта
	Tags       string           `json:"tags"`        //теги продукта
	Variants   []Variant        `json:"variants"`    //	Список вариантов продукта
}

// Variant - Структура варианта, продукта представляем с собой информацию о продукте который нужно внести в базу
type Variant struct {
	ProductId    int             `json:"product_id" db:"product_id"` //id продука
	VariantId    int             `json:"variant_id" db:"variant_id"` //id конкретного варианта продукта
	Weight       int             `json:"weight" db:"weight"`         // масса или вес продукта
	Unit         string          `json:"unit" db:"unit"`             //единица измерения
	Added_at     time.Time       `json:"added_at" db:"added_at"`     // дата добавления определенного варианта
	CurrentPrice decimal.Decimal `db:"price"`                        //актуальная цена
	InStorages   []int           `json:"in_storages"`                //список id складов в которых есть этот вариант
}

// ProductPrice - Структура для вставки цены продукта
type ProductPrice struct {
	PriceId   int              //id цены продукта
	VariantId int              `json:"variant_id"` //id варианта продука
	StartDate time.Time        `json:"start_date"` // дата начала цены
	EndDate   sqlnull.NullTime `json:"end_date"`   //дата конца цены
	Price     decimal.Decimal  `json:"price"`      //цена продукта
}

// AddProductInStock - Структура для вставки продукта на склад
type AddProductInStock struct {
	VariantId int       `json:"variant_id" db:"variant_id"` //id варианта продукта
	StorageId int       `json:"storage_id" db:"storage_id"` //id склада куда будет помещен этот продукт
	Added_at  time.Time `json:"added_at" db:"added_at" `    //дата добавления продукта на склад
	Quantity  int       `json:"quantity" db:"quantity"`     //кол-во продукта добавленного на склад
}

func (a *AddProductInStock) IsNullFields() {
}

// ProductInfo - Структура информации о продукте о котором нужно получить информацию
type ProductInfo struct {
	ProductId int       `db:"product_id"`  //id продукта
	Name      string    `db:"name"`        //название продукта
	Descr     string    `db:"description"` //описание продукта
	Variants  []Variant //список вариантов продукта
}

// Stock - структура склада
type Stock struct {
	StorageID       int                 `db:"storage_id"` //id склада
	StorageName     string              `db:"name"`       //название склада
	ProductVariants []AddProductInStock //список продуктов на данном складе
}

// Sale - Структура продажи
type Sale struct {
	SaleId      int                `db:"sales_id"`                     //id продажи
	ProductName sqlnull.NullString `db:"name"`                         //id продукта
	VariantId   int                `json:"variant_id" db:"variant_id"` //id варианта продукта
	StorageId   int                `json:"storage_id" db:"storage_id"` //id склада из которого произошла продажа продукта
	SoldAt      time.Time          `db:"sold_at"`                      //дата продажи
	Quantity    int                `json:"quantity" db:"quantity"`     //кол-во проданного продукта
	TotalPrice  decimal.Decimal    `db:"total_price"`                  //общая стоимость с учетом кол-ва продукта
}

// SaleQuery - фильтры продаж по которым нужно вывести информацию
type SaleQuery struct {
	StartDate   time.Time          `json:"start_date"`   //дата начала продаж(обязательные поля)
	EndDate     time.Time          `json:"end_date"`     //дата конца прдаж (обязательные поля)
	Limit       sqlnull.NullInt64  `json:"limit"`        //лимит вывода продаж
	StorageId   sqlnull.NullInt64  `json:"storage_id"`   //id склада
	ProductName sqlnull.NullString `json:"product_name"` //название продукта
}
