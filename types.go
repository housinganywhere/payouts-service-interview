package main

import "time"

type Payout struct {
	Id           int       `json:"id"`           // Id of the payout in the database
	Amount       int       `json:"amount"`       // Amount in cents
	Currency     string    `json:"currency"`     // Currency of the payout
	BookingId    int       `json:"bookingId"`    // BookingId from external service.
	AdvertiserId int       `json:"advertiserId"` // AdvertiserId is the id of the landlord (seller)
	TenantId     int       `json:"tenantId"`     // TenantId is the id of the tenant (buyer)
	PayinAt      time.Time `json:"payinAt"`      // PayinAt is the timestamp, which indicates exact time of when the payment from tenant side was made
	PayoutDate   time.Time `json:"payoutDate"`   // PayoutDate is the date, when we should send the money to advertiser
}
