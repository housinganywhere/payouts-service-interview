package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type PayoutRepo struct {
	db *sql.DB
}

// We are dropping the database to make sure we can rerun the project
const initPayoutTableSQL = `
DROP TABLE IF EXISTS payouts;
CREATE TABLE payouts (
    id INTEGER PRIMARY KEY,
	amount         INTEGER,
	currency       VARCHAR(3),
	booking_id      INTEGER,
	advertiser_id   INTEGER,
	tenant_id       INTEGER,
    payin_at DATETIME,
    payout_date     DATE
);`

const insertPayoutSQL = `INSERT INTO payouts (
	amount,	currency, booking_id, advertiser_id, tenant_id, payin_at, payout_date
) VALUES (?, ?, ?, ?, ?, ?, ?)`

const selectAllPayouts = `SELECT * FROM payouts`

func (pr *PayoutRepo) InsertPayout(payout Payout) error {
	_, err := pr.db.Exec(insertPayoutSQL, payout.Amount, payout.Currency, payout.BookingId, payout.AdvertiserId, payout.TenantId, payout.PayinAt, payout.PayoutDate)

	return err
}

func (pr *PayoutRepo) InitPayoutTable() error {
	_, err := pr.db.Exec(initPayoutTableSQL)

	return err
}

func (pr *PayoutRepo) GetAllPayouts() ([]Payout, error) {
	rows, err := pr.db.Query(selectAllPayouts)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	payouts := []Payout{}
	for rows.Next() {
		payout := Payout{}

		err = rows.Scan(&payout.Id, &payout.Amount, &payout.Currency, &payout.BookingId, &payout.AdvertiserId, &payout.TenantId, &payout.PayinAt, &payout.PayoutDate)

		if err != nil {
			return nil, err
		}
		// Add payout to the result array
		payouts = append(payouts, payout)
	}
	return payouts, nil
}
