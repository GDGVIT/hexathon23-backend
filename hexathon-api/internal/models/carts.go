package models

import (
	"errors"
	"fmt"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Team       Team      `gorm:"foreignKey:TeamID;constraint:OnDelete:CASCADE;"`
	TeamID     uuid.UUID
	Items      []Item `gorm:"many2many:cart_items;"`
	Cost       int
	CheckedOut bool `gorm:"default:false"`
}

// CreateCart creates a new cart
func (cart *Cart) CreateCart() error {
	return database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(cart).Error
}

// UpdateCart updates a cart
func (cart *Cart) UpdateCart() error {
	return database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(cart).Error
}

// CheckItemInCart checks if an item is in the cart
func (cart *Cart) CheckItemInCart(item Item) bool {
	for _, cartItem := range cart.Items {
		if cartItem.ID == item.ID {
			return true
		}
	}
	return false
}

// AddToCart adds an item to the cart
func (cart *Cart) AddToCart(item Item) error {
	if cart.CheckedOut {
		return errors.New("cart already checked out")
	}
	// Loop through the items in the cart and check if the item already exists
	// Check if category max limit is reached
	itemCount := 0
	for _, cartItem := range cart.Items {
		if cartItem.ID == item.ID {
			return errors.New("item already exists in cart")
		}
		if cartItem.CategoryID == item.CategoryID {
			itemCount++
		}
	}
	category, err := GetCategoryByID(item.CategoryID.String())
	if err != nil {
		return err
	}
	if itemCount >= category.MaxItems {
		return errors.New("category max limit reached")
	}
	// Check there is enough amount left
	if cart.Team.Amount-cart.Cost < item.Price {
		return errors.New("not enough amount left")
	}
	cart.Cost += item.Price
	cart.Items = append(cart.Items, item)
	return cart.UpdateCart()
}

// DeleteFromCart deletes an item from the cart
func (cart *Cart) DeleteFromCart(item Item) error {
	if cart.CheckedOut {
		return errors.New("cart already checked out")
	}
	// If item not in cart return error
	found := false
	for _, cartItem := range cart.Items {
		if cartItem.ID == item.ID {
			found = true
			break
		}
	}
	if !found {
		return errors.New("item not in cart")
	}

	cart.Cost -= item.Price
	err := database.DB.Model(cart).Association("Items").Delete(&item)
	if err != nil {
		return err
	}
	return cart.UpdateCart()
}

// CheckoutCart checks out the cart
func (cart *Cart) CheckoutCart() error {
	if cart.CheckedOut {
		return errors.New("cart already checked out")
	}
	// Check if None of the items are over the max limit
	itemCount := make(map[uuid.UUID]int)
	for _, item := range cart.Items {
		itemCount[item.CategoryID]++
	}
	for categoryID, count := range itemCount {
		category, err := GetCategoryByID(categoryID.String())
		if err != nil {
			return err
		}
		if count > category.MaxItems {
			return fmt.Errorf("max limit reached for category: %s", category.Name)
		}
	}
	// Check if the cart is empty
	if len(cart.Items) == 0 {
		return errors.New("Cart is empty")
	}
	cart.CheckedOut = true
	cart.Team.Amount -= cart.Cost
	cart.Team.SetItemsPurchased(cart.Items)
	cart.Team.UpdateTeam()
	return cart.UpdateCart()
}
