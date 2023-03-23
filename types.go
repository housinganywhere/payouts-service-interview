package main

import "time"

type Payout struct {
	Id           int       `json:"id"`
	Amount       int       `json:"amount"`
	Currency     string    `json:"currency"`
	BookingId    int       `json:"bookingId"`
	AdvertiserId int       `json:"advertiserId"`
	TenantId     int       `json:"tenantId"`
	PayinAt      time.Time `json:"payinAt"`
	PayoutDate   time.Time `json:"payoutDate"`
}
