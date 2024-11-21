package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductID uint32 `gorm:"type:int(11);not null;"`
	Quantity  uint32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var find Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: item.UserID, ProductID: item.ProductID}).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if find.ID != 0 {
		err = db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: item.UserID, ProductID: item.ProductID}).UpdateColumn("quantity", gorm.Expr("quantity+?", item.Quantity)).Error
	} else {
		err = db.WithContext(ctx).Model(&Cart{}).Create(item).Error
	}
	return err
}

func EmptyCart(ctx context.Context, db *gorm.DB, userID uint32) error {
	if userID == 0 {
		return errors.New("user_is is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userID).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) (cartList []*Cart, err error) {
	err = db.Debug().WithContext(ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}
