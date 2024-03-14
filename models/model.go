package models

type User struct {
	ID       int    `json:"id,omitempty" gorm:"primaryKey"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	Address  string `json:"address,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data,omitempty"`
}

type Product struct {
	ID    int    `json:"id,omitempty" gorm:"primaryKey"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}

type Transaction struct {
	ID        int `json:"ID"`
	UserID    int `json:"UserID,omitempty" gorm:"column:UserID"`
	ProductID int `json:"ProductID,omitempty" gorm:"column:ProductID"`
	Quantity  int `json:"quantity,omitempty"`
}

type TransactionDetail struct {
	ID       int     `json:"transactionID"`
	User     User    `json:"user"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

type TransactionsDetail struct {
	Transaction []TransactionDetail `json:"transactions"`
}

type TransactionDetailResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    TransactionsDetail `json:"data"`
}

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}