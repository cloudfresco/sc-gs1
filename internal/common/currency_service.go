package common

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	"go.uber.org/zap"
)

// CurrencyService - For accessing Currency services
type CurrencyService struct {
	log       *zap.Logger
	DBService *DBService
}

// NewCurrencyService - Create Currency Service
func NewCurrencyService(log *zap.Logger, dbOpt *DBService) *CurrencyService {
	return &CurrencyService{
		log:       log,
		DBService: dbOpt,
	}
}

func (cs *CurrencyService) GetCurrency(ctx context.Context, code string) (*commonproto.Currency, error) {
	const query = `SELECT code, numeric_code, currency_name, minor_unit
                   FROM currencies WHERE code = ?`

	// var currency commonproto.Currency
	// err := cs.DBService.GetContext(ctx, &currency, query, code)
	row := cs.DBService.DB.QueryRowxContext(ctx, query, code)
	currency := commonproto.Currency{}
	err := row.StructScan(&currency)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("currency not found: %s", code)
		}
		return nil, fmt.Errorf("failed to get currency: %w", err)
	}
	return &currency, nil
}

func ParseAmountString(amountStr string, currency *commonproto.Currency) (int64, error) {
	// Check for valid number format
	if !strings.Contains(amountStr, ".") {
		amountStr += "."
	}

	parts := strings.Split(amountStr, ".")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid amount format")
	}

	// Validate decimal places match currency
	if int32(len(parts[1])) != currency.MinorUnit {
		return 0, fmt.Errorf("amount must have exactly %d decimal places", currency.MinorUnit)
	}

	// Combine into integer value
	combined := parts[0] + parts[1]
	amountMinor, err := strconv.ParseInt(combined, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount value")
	}

	return amountMinor, nil
}

// FormatAmountString converts minor units back to properly formatted string
func FormatAmountString(amountMinor int64, currency *commonproto.Currency) string {
	str := strconv.FormatInt(amountMinor, 10)

	if currency.MinorUnit == 0 {
		return str
	}

	// Pad with leading zeros if needed
	if int32(len(str)) <= currency.MinorUnit {
		// str = strings.Repeat("0", currency.MinorUnit-len(str)+1) + str
		// str = strings.Repeat("0", currency.MinorUnit-int32(len(str))+1) + str
		minorUnit := int(currency.MinorUnit)
		str = strings.Repeat("0", minorUnit-len(str)+1) + str
	}

	// Insert decimal point
	pos := int32(len(str)) - currency.MinorUnit
	return str[:pos] + "." + str[pos:]
}
