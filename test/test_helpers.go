package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"strings"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

// GetUUIDDateValues -- data for tests
func GetUUIDDateValues(log *zap.Logger) ([]byte, string, *timestamp.Timestamp, *timestamp.Timestamp, string, string, string, string, error) {
	uuid4, err := common.GetUUIDBytes()
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, "", nil, nil, "", "", "", "", err
	}
	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)
	uuid4Str, err := common.UUIDBytesToStr(uuid4)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, "", nil, nil, "", "", "", "", err
	}
	uuid4StrJSON, _ := json.Marshal(uuid4Str)
	uuid4JSON, _ := json.Marshal(uuid4)
	createdAtJSON, _ := json.Marshal(tn)
	updatedAtJSON, _ := json.Marshal(tn)
	return uuid4, uuid4Str, tn, tn, string(uuid4JSON), string(uuid4StrJSON), string(createdAtJSON), string(updatedAtJSON), nil
}

// LoadSQL -- drop db, create db, use db, load data
func LoadSQL(log *zap.Logger, dbService *common.DBService) error {
	var err error
	ctx := context.Background()

	if dbService.DBType == common.DBMysql {
		err = execSQLFile(ctx, log, dbService.MySQLTruncateFilePath, dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/animal_identifications.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/debit_credit_advice_line_item_details.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/debit_credit_advice_line_items.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/debit_credit_advices.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/description1000.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/description200.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/description500.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/despatch_advice_line_items.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/despatch_advice_logistic_units.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/despatch_advices.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/despatch_informations.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/ecom_document_references.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/ecom_entity_identification.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/fish_catch_or_production_dates.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/fish_catch_or_productions.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/fish_despatch_advice_line_item_extensions.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/inventory_activity_line_items.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/inventory_activity_quantity_specifications.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/inventory_item_location_informations.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/inventory_reports.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/inventory_status_line_items.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/inventory_sub_locations.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/invoice_line_item_information_after_taxes.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/invoice_line_items.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/invoices.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/invoice_totals.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/levied_duty_fee_taxes.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/logistics_inventory_report_inventory_locations.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/logistics_inventory_reports.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/meat_activity_histories.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/meat_breeding_details.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/meat_cutting_details.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/meat_fattening_details.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/meat_mincing_details.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/meat_slaughtering_details.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/order_line_item_details.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/order_line_items.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/order_logistical_informations.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/order_response_line_items.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/order_responses.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/orders.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/payment_terms.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/receiving_advice_line_items.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/receiving_advices.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/trade_item_inventory_events.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/trade_item_inventory_statuses.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/transactional_item_data.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/transactional_parties.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/object_events.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/association_events.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/aggregation_events.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/transaction_events.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/transformation_events.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/biz_transactions.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/destinations.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/persistent_dispositions.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/epcs.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/quantity_elements.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/sources.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/sensor_elements.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.MySQLTestFilePath+"/sensor_reports.sql", dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

	} else if dbService.DBType == common.DBPgsql {
		err = execSQLFile(ctx, log, dbService.PgSQLTruncateFilePath, dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}

		err = execSQLFile(ctx, log, dbService.PgSQLTestFilePath, dbService.DB)
		if err != nil {
			log.Error("Error", zap.Error(err))
			return err
		}
	}

	return nil
}

func execSQLFile(ctx context.Context, log *zap.Logger, sqlFilePath string, db *sqlx.DB) error {
	content, err := os.ReadFile(sqlFilePath)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	tx, err := db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	sqlLines := strings.Split(string(content), ";\n")
	for _, sqlLine := range sqlLines {
		if sqlLine != "" {
			_, err := tx.ExecContext(ctx, sqlLine)
			if err != nil {
				log.Error("Error", zap.Error(err))
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					log.Error("Error", zap.Error(rollbackErr))
					return err
				}
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}
	return nil
}
