package repository

import (
	"cart/domain/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)
	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartRepository struct {
	mysqlDB *gorm.DB
}

func (c *CartRepository) InitTable() error {
	return c.mysqlDB.CreateTable(&model.Cart{}).Error
}

func (c *CartRepository) FindCartByID(i int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, c.mysqlDB.First(cart, i).Error
}

func (c *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := c.mysqlDB.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}

func (c *CartRepository) DeleteCartByID(i int64) error {
	return c.mysqlDB.Where("id = ?", i).Delete(&model.Cart{}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDB.Model(cart).Update(cart).Error
}

func (c *CartRepository) FindAll(userID int64) (cartAll []model.Cart, err error) {
	return cartAll, c.mysqlDB.Where("user_id = ?", userID).Find(&cartAll).Error
}

func (c *CartRepository) CleanCart(i int64) error {
	return c.mysqlDB.Where("user_id = ?", i).Delete(&model.Cart{}).Error
}

func (c *CartRepository) IncrNum(i int64, i2 int64) error {
	cart := &model.Cart{ID: i}
	return c.mysqlDB.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", i2)).Error
}

func (c *CartRepository) DecrNum(i int64, i2 int64) error {
	cart := &model.Cart{
		ID: i,
	}
	db := c.mysqlDB.Model(cart).Where("num >= ?", i2).UpdateColumn("num", gorm.Expr("num = ?", i2))
	if db.Error  != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDB: db}
}
