package model

import "time"

type PortfolioContactModel struct {
	UUID        string    `json:"uuid"`
	RealName    string    `json:"real_name"`
	Email       string    `json:"email"`
	TextContent string    `json:"text_content"`
	CreatedAt   time.Time `json:"created_at"`
}
