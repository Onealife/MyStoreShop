package ports

import (
	"github.com/Onealife/MyStoreShop/internal/core/domain"
	"github.com/google/uuid"
)

// UserRepository interface สำหรับจัดการข้อมูลผู้ใช้
type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uuid.UUID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uuid.UUID) error
	UpdateRefreshToken(id uuid.UUID, refreshrToken string) error
	List(limit, offset int) ([]domain.User, error)
}

// RoleRepository interface สำหรับจัดการข้อมูลสิทธิ์การใช้งาน
type RoleRepository interface {
	Create(role *domain.Role) error
	FindByID(id uuid.UUID) (*domain.Role, error)
	FindByName(name string) (*domain.Role, error)
	Update(role *domain.Role) error
	Delete(id uuid.UUID) error
	List() ([]domain.Role, error)
	AddPermission(roleID, permissionID uuid.UUID) error
	RemovePermission(roleID, permissionID uuid.UUID) error
}

// PermissionRepository interface สำหรับจัดการข้อมูลสิทธิ์การเข้าถึง
type PermissionRepository interface {
	Create(permission *domain.Permission) error
	FindByID(id uuid.UUID) (*domain.Permission, error)
	FindByName(name string) (*domain.Permission, error)
	Update(permission *domain.Permission) error
	Delete(id uuid.UUID) error
	List() ([]domain.Permission, error)
}

// CategoryRepository interface สำหรับจัดการข้อมูลหมวดหมู่สินค้า
type CategoryRepository interface {
	Create(category *domain.Category) error
	FindByID(id uuid.UUID) (*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id uuid.UUID) error
	List() ([]domain.Category, error)
}

// ProductRepository interface สำหรับจัดการข้อมูลสินค้า
type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id uuid.UUID) (*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uuid.UUID) error
	List(limit, offset int) ([]domain.Product, error)
	FindByCategoryID(categoryID uuid.UUID, limit, offset int) ([]domain.Product, error)
	Search(keyword string, limit, offset int) ([]domain.Product, error)
	AddImage(productID uuid.UUID, imageURL string) error
	RemoveImage(imageID uuid.UUID) error
}

// CartRepository interface สำหรับจัดการข้อมูลตะกร้าสินค้า
type CartRepository interface {
	Create(cart *domain.Cart) error
	FindByID(id uuid.UUID) (*domain.Cart, error)
	FindByUserID(userID uuid.UUID) (*domain.Cart, error)
	Update(cart *domain.Cart) error
	Delete(id uuid.UUID) error
	AddItem(cartID, productID uuid.UUID, quantity int, price float64) error
	UpdateItem(cartItemID uuid.UUID, quantity int) error
	RemoveItem(cartItemID uuid.UUID) error
	ClearItems(cartID uuid.UUID) error
}

// OrderRepository interface สำหรับจัดการข้อมูลการสั่งซื้อ
type OrderRepository interface {
	Create(order *domain.Order) error
	FindByID(id uuid.UUID) (*domain.Order, error)
	FindByUserID(userID uuid.UUID, limit, offset int) ([]domain.Order, error)
	Update(order *domain.Order) error
	Delete(id uuid.UUID) error
	AddItem(orderID, productID uuid.UUID, quantity int, price float64) error
	UpdateItem(orderItemID uuid.UUID, quantity int) error
	RemoveItem(orderItemID uuid.UUID) error
	UpdateStatus(orderID uuid.UUID, status string) error
	UpdatePaymentStatus(orderID uuid.UUID, status string) error
	UpdateShippingStatus(orderID uuid.UUID, status string) error
	List(limit, offset int) ([]domain.Order, error)
}

// TransactionRepository interface สำหรับจัดการข้อมูลธุรกรรมการชำระเงิน
type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
	FindByID(id uuid.UUID) (*domain.Transaction, error)
	FindByOrderID(orderID uuid.UUID) ([]domain.Transaction, error)
	Update(transaction *domain.Transaction) error
	Delete(id uuid.UUID) error
	UpdateStatus(transactionID uuid.UUID, status string) error
}
