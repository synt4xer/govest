// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	FullName    string  `json:"fullName"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

type Mutation struct {
}

type Portfolio struct {
	Stock                *Stock  `json:"stock"`
	Quantity             int     `json:"quantity"`
	AveragePrice         float64 `json:"averagePrice"`
	UnrealizedProfitLoss float64 `json:"unrealizedProfitLoss"`
}

type Query struct {
}

type Stock struct {
	ID           string  `json:"id"`
	Ticker       string  `json:"ticker"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"currentPrice"`
}

type StockTransactionInput struct {
	StockID       string  `json:"stockId"`
	Quantity      int     `json:"quantity"`
	PricePerStock float64 `json:"pricePerStock"`
}

type Transaction struct {
	ID              string  `json:"id"`
	Stock           *Stock  `json:"stock"`
	TransactionType string  `json:"transactionType"`
	Quantity        int     `json:"quantity"`
	PricePerStock   float64 `json:"pricePerStock"`
	TransactionDate string  `json:"transactionDate"`
}

type User struct {
	ID          string  `json:"id"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	FullName    string  `json:"fullName"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	CreatedAt   string  `json:"createdAt"`
}