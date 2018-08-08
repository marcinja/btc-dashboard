package main

import (
	"github.com/btcsuite/btcd/btcjson"
)

// The Data type encapsulates all future tables added to the dashboard.
// Each field should correspond to a different table and have its own struct type.
// This makes using go-pg easy because of ORM.
type Data struct {
	Version          int64           `json:"version"`
	DashboardDataRow DashboardDataV2 `json:"dashboard_data"`

	// Future tables below:
}

// dataBatch is used internally for collecting data for batch insertions.
type dataBatch struct {
	versions          []int64
	dashboardDataRows []DashboardDataV2
}

type BlockStats struct {
	*btcjson.GetBlockStatsResult
}

func (metrics BlockStats) transformToDashboardData() DashboardDataV2 {
	data := DashboardDataV2{}
	data.Id = metrics.Height

	data.Mto_consolidations = metrics.Mto_consolidations
	data.Mto_output_count = metrics.Mto_output_count
	data.Mto_avg_fee = metrics.Mto_avg_fee
	data.Mto_avg_value = metrics.Mto_avg_value

	data.Per_90th_fee = metrics.Per_90th_fee
	data.Per_90th_fee_rate = metrics.Per_90th_fee_rate
	data.Per_90th_tx_size = metrics.Per_90th_tx_size

	data.Value_of_native_P2WPKH_outputs_spent = metrics.Value_of_native_P2WPKH_outputs_spent
	data.Value_of_native_P2WSH_outputs_spent = metrics.Value_of_native_P2WSH_outputs_spent
	data.Value_of_nested_P2WPKH_outputs_spent = metrics.Value_of_nested_P2WPKH_outputs_spent
	data.Value_of_nested_P2WSH_outputs_spent = metrics.Value_of_nested_P2WSH_outputs_spent
	data.Value_of_new_P2WPKH_outputs = metrics.Value_of_new_P2WPKH_outputs
	data.Value_of_new_P2WSH_outputs = metrics.Value_of_new_P2WSH_outputs

	data.Time = metrics.Time

	data.Avg_fee = metrics.AverageFee
	data.Avg_fee_rate = metrics.AverageFeeRate
	data.Avg_tx_size = metrics.AverageTxSize

	data.Max_fee = metrics.MaxFee
	data.Max_fee_rate = metrics.MaxFeeRate
	data.Max_tx_size = metrics.MaxTxSize

	data.Min_fee = metrics.MinFee
	data.Min_fee_rate = metrics.MinFeeRate
	data.Min_tx_size = metrics.MinTxSize

	data.Median_fee = metrics.MedianFee
	data.Median_fee_rate = metrics.MedianFeeRate
	data.Median_tx_size = metrics.MedianTxSize

	data.Block_size = metrics.TotalSize
	data.Volume_btc = metrics.TotalOut
	data.Num_txs = metrics.Txs

	data.Hash = metrics.Hash
	data.Height = metrics.Height
	data.Num_inputs = metrics.Ins
	data.Num_outputs = metrics.Outs

	data.Subsidy = metrics.Subsidy
	data.Segwit_total_size = metrics.SegWitTotalSize
	data.Segwit_total_weight = metrics.SegWitTotalWeight
	data.Num_segwit_txs = metrics.SegWitTxs

	data.Total_amount_out = metrics.TotalOut
	data.Total_size = metrics.TotalSize
	data.Total_weight = metrics.TotalWeight
	data.Total_fee = metrics.TotalFee

	data.Utxo_increase = metrics.UTXOIncrease
	data.Utxo_size_increase = metrics.UTXOSizeIncrease

	data.Nested_P2WPKH_outputs_spent = metrics.NestedP2WPKHOutputsSpent
	data.Native_P2WPKH_outputs_spent = metrics.NativeP2WPKHOutputsSpent
	data.Nested_P2WSH_outputs_spent = metrics.NestedP2WSHOutputsSpent
	data.Native_P2WSH_outputs_spent = metrics.NativeP2WSHOutputsSpent

	data.Txs_spending_nested_p2wpkh_outputs = metrics.TxsSpendingNestedP2WPKHOutputs
	data.Txs_spending_nested_p2wsh_outputs = metrics.TxsSpendingNestedP2WSHOutputs
	data.Txs_spending_native_p2wpkh_outputs = metrics.TxsSpendingNativeP2WPKHOutputs
	data.Txs_spending_native_p2wsh_outputs = metrics.TxsSpendingNativeP2WSHOutputs

	data.Txs_spending_native_sw_outputs = metrics.TxsSpendingNativeP2WPKHOutputs + metrics.TxsSpendingNativeP2WSHOutputs
	data.Txs_spending_nested_sw_outputs = metrics.TxsSpendingNestedP2WPKHOutputs + metrics.TxsSpendingNestedP2WSHOutputs

	data.New_P2WPKH_outputs = metrics.NewP2WPKHOutputs
	data.New_P2WSH_outputs = metrics.NewP2WSHOutputs
	data.Num_txs_creating_P2WPKH = metrics.TxsCreatingP2WPKHOutputs
	data.Num_txs_creating_P2WSH = metrics.TxsCreatingP2WSHOutputs

	data.TxsSignallingOptInRBF = metrics.TxsSignallingOptInRBF
	data.Num_consolidating_txs = metrics.TxsConsolidating
	data.Num_outputs_consolidated = metrics.OutputsConsolidated
	data.Num_batching_txs = metrics.TxsBatching

	// Derived added below /////////////////////////////////////////////////////
	data.Num_txs_creating_native_segwit_outputs = metrics.TxsCreatingP2WPKHOutputs + metrics.TxsCreatingP2WSHOutputs

	data.Dust_bin_0 = metrics.DustBins[0]
	data.Dust_bin_1 = metrics.DustBins[1]
	data.Dust_bin_2 = metrics.DustBins[2]
	data.Dust_bin_3 = metrics.DustBins[3]
	data.Dust_bin_4 = metrics.DustBins[4]
	data.Dust_bin_5 = metrics.DustBins[5]
	data.Dust_bin_6 = metrics.DustBins[6]
	data.Dust_bin_7 = metrics.DustBins[7]
	data.Dust_bin_8 = metrics.DustBins[8]
	data.Dust_bin_9 = metrics.DustBins[9]
	data.Dust_bin_10 = metrics.DustBins[10]
	data.Dust_bin_11 = metrics.DustBins[11]
	data.Dust_bin_12 = metrics.DustBins[12]
	data.Dust_bin_13 = metrics.DustBins[13]
	data.Dust_bin_14 = metrics.DustBins[14]
	data.Dust_bin_15 = metrics.DustBins[15]
	data.Dust_bin_16 = metrics.DustBins[16]
	data.Dust_bin_17 = metrics.DustBins[17]
	data.Dust_bin_18 = metrics.DustBins[18]
	data.Dust_bin_19 = metrics.DustBins[19]
	data.Dust_bin_20 = metrics.DustBins[20]
	data.Dust_bin_21 = metrics.DustBins[21]
	data.Output_count_bin_0 = metrics.OutputCountBins[0]
	data.Output_count_bin_1 = metrics.OutputCountBins[1]
	data.Output_count_bin_2 = metrics.OutputCountBins[2]
	data.Output_count_bin_3 = metrics.OutputCountBins[3]
	data.Output_count_bin_4 = metrics.OutputCountBins[4]
	data.Output_count_bin_5 = metrics.OutputCountBins[5]
	data.Output_count_bin_6 = metrics.OutputCountBins[6]

	if metrics.Txs != 0 {
		data.Batch_range_0 = float64(metrics.OutputCountBins[0]) / float64(metrics.Txs)
		data.Batch_range_1 = float64(metrics.OutputCountBins[1]) / float64(metrics.Txs)
		data.Batch_range_2 = float64(metrics.OutputCountBins[2]) / float64(metrics.Txs)
		data.Batch_range_3 = float64(metrics.OutputCountBins[3]) / float64(metrics.Txs)
		data.Batch_range_4 = float64(metrics.OutputCountBins[4]) / float64(metrics.Txs)
		data.Batch_range_5 = float64(metrics.OutputCountBins[5]) / float64(metrics.Txs)
		data.Batch_range_6 = float64(metrics.OutputCountBins[6]) / float64(metrics.Txs)

		data.Percent_txs_signalling_opt_in_RBF = float64(metrics.TxsSignallingOptInRBF) / float64(metrics.Txs)
		data.Percent_txs_batching = float64(metrics.TxsBatching) / float64(metrics.Txs)
		data.Percent_txs_consolidating = float64(metrics.TxsConsolidating) / float64(metrics.Txs)
		data.Percent_txs_creating_native_segwit_outputs = float64(metrics.TxsCreatingP2WPKHOutputs+metrics.TxsCreatingP2WSHOutputs) / float64(metrics.Txs)
		data.Percent_txs_creating_P2WSH_outputs = float64(metrics.TxsCreatingP2WSHOutputs) / float64(metrics.Txs)
		data.Percent_txs_creating_P2WPKH_outputs = float64(metrics.TxsCreatingP2WPKHOutputs) / float64(metrics.Txs)

		data.Percent_txs_spending_native_segwit_outputs = float64(metrics.TxsSpendingNativeP2WPKHOutputs+metrics.TxsSpendingNativeP2WSHOutputs) / float64(metrics.Txs)
		data.Percent_txs_spending_nested_segwit_outputs = float64(metrics.TxsSpendingNestedP2WPKHOutputs+metrics.TxsSpendingNestedP2WSHOutputs) / float64(metrics.Txs)
		data.Percent_txs_spending_native_P2WPKH_outputs = float64(metrics.TxsSpendingNativeP2WPKHOutputs) / float64(metrics.Txs)
		data.Percent_txs_spending_native_P2WSH_outputs = float64(metrics.TxsSpendingNativeP2WSHOutputs) / float64(metrics.Txs)
		data.Percent_txs_spending_nested_P2WPKH_outputs = float64(metrics.TxsSpendingNestedP2WPKHOutputs) / float64(metrics.Txs)
		data.Percent_txs_spending_nested_P2WSH_outputs = float64(metrics.TxsSpendingNestedP2WSHOutputs) / float64(metrics.Txs)
		data.Percent_txs_spending_P2WSH_outputs = float64(metrics.TxsSpendingNativeP2WSHOutputs+metrics.TxsSpendingNestedP2WSHOutputs) / float64(metrics.Txs)
		data.Percent_txs_spending_P2WPKH_outputs = float64(metrics.TxsSpendingNativeP2WPKHOutputs+metrics.TxsSpendingNestedP2WPKHOutputs) / float64(metrics.Txs)
		data.Percent_txs_that_are_segwit_txs = float64(metrics.SegWitTxs) / float64(metrics.Txs)

	}

	if metrics.Outs != 0 {
		data.Percent_new_outs_in_dust_bin_0 = float64(metrics.DustBins[0]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_1 = float64(metrics.DustBins[1]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_2 = float64(metrics.DustBins[2]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_3 = float64(metrics.DustBins[3]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_4 = float64(metrics.DustBins[4]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_5 = float64(metrics.DustBins[5]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_6 = float64(metrics.DustBins[6]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_7 = float64(metrics.DustBins[7]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_8 = float64(metrics.DustBins[8]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_9 = float64(metrics.DustBins[9]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_10 = float64(metrics.DustBins[10]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_11 = float64(metrics.DustBins[11]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_12 = float64(metrics.DustBins[12]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_13 = float64(metrics.DustBins[13]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_14 = float64(metrics.DustBins[14]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_15 = float64(metrics.DustBins[15]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_16 = float64(metrics.DustBins[16]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_17 = float64(metrics.DustBins[17]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_18 = float64(metrics.DustBins[18]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_19 = float64(metrics.DustBins[19]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_20 = float64(metrics.DustBins[20]) / float64(metrics.Outs)
		data.Percent_new_outs_in_dust_bin_21 = float64(metrics.DustBins[21]) / float64(metrics.Outs)

		data.Percent_new_outs_P2WPKH_outputs = float64(metrics.NewP2WPKHOutputs) / float64(metrics.Outs)
		data.Percent_new_outs_P2WSH_outputs = float64(metrics.NewP2WSHOutputs) / float64(metrics.Outs)
	}

	if metrics.Ins != 0 {
		data.Percent_of_inputs_spending_nested_P2WPKH_output = float64(metrics.NestedP2WPKHOutputsSpent) / float64(metrics.Ins)
		data.Percent_of_inputs_spending_native_P2WPKH_outputs = float64(metrics.NativeP2WPKHOutputsSpent) / float64(metrics.Ins)
		data.Percent_of_inputs_spending_P2WPKH_outputs = float64(metrics.NativeP2WPKHOutputsSpent+metrics.NestedP2WPKHOutputsSpent) / float64(metrics.Ins)
		data.Percent_of_inputs_spending_nested_P2WSH_outputs = float64(metrics.NestedP2WSHOutputsSpent) / float64(metrics.Ins)
		data.Percent_of_inputs_spending_native_P2WSH_outputs = float64(metrics.NativeP2WSHOutputsSpent) / float64(metrics.Ins)
		data.Percent_of_inputs_spending_P2WSH_outputs = float64(metrics.NativeP2WSHOutputsSpent+metrics.NestedP2WSHOutputsSpent) / float64(metrics.Ins)
		data.Percent_of_inputs_spending_native_sw_outputs = float64(metrics.NativeP2WSHOutputsSpent+metrics.NativeP2WSHOutputsSpent) / float64(metrics.Ins)
		data.Percent_inputs_consolidated = float64(metrics.OutputsConsolidated) / float64(metrics.Ins)
	}

	if metrics.SegWitTxs != 0 {
		data.Percent_txs_native_segwit_over_total_sw_txs = float64(metrics.TxsSpendingNativeP2WSHOutputs+metrics.TxsSpendingNativeP2WPKHOutputs) / float64(metrics.SegWitTxs)
	}

	return data
}

type DashboardData struct {
	Id int64 `json:"id,omit_empty" sql:",notnull;primary_key"`

	Mto_consolidations int64 `json:"mto_consolidations,omit_empty" sql:",notnull"`
	Mto_output_count   int64 `json:"mto_output_count,omit_empty" sql:",notnull"`

	Time int64 `json:"time" sql:",notnull"`

	Avg_fee      int64 `json:"avg_fee" sql:",notnull"`
	Avg_fee_rate int64 `json:"avg_fee_rate" sql:",notnull"`
	Avg_tx_size  int64 `json:"avg_tx_size" sql:",notnull"`

	Max_fee      int64 `json:"max_fee" sql:",notnull"`
	Max_fee_rate int64 `json:"max_fee_rate" sql:",notnull"`
	Max_tx_size  int64 `json:"max_tx_size" sql:",notnull"`

	Median_fee      int64 `json:"median_fee" sql:",notnull"`
	Median_fee_rate int64 `json:"median_fee_rate" sql:",notnull"`
	Median_tx_size  int64 `json:"median_tx_size" sql:",notnull"`

	Min_fee      int64 `json:"min_fee" sql:",notnull"`
	Min_fee_rate int64 `json:"min_fee_rate" sql:",notnull"`
	Min_tx_size  int64 `json:"min_tx_size" sql:",notnull"`

	Batch_range_0 float64 `json:"batch_range_0" sql:",notnull"`
	Batch_range_1 float64 `json:"batch_range_1" sql:",notnull"`
	Batch_range_2 float64 `json:"batch_range_2" sql:",notnull"`
	Batch_range_3 float64 `json:"batch_range_3" sql:",notnull"`
	Batch_range_4 float64 `json:"batch_range_4" sql:",notnull"`
	Batch_range_5 float64 `json:"batch_range_5" sql:",notnull"`
	Batch_range_6 float64 `json:"batch_range_6" sql:",notnull"`

	Block_size int64 `json:"block_size" sql:",notnull"`

	Dust_bin_0  int64 `json:"dust_bin_0" sql:",notnull"`
	Dust_bin_1  int64 `json:"dust_bin_1" sql:",notnull"`
	Dust_bin_2  int64 `json:"dust_bin_2" sql:",notnull"`
	Dust_bin_3  int64 `json:"dust_bin_3" sql:",notnull"`
	Dust_bin_4  int64 `json:"dust_bin_4" sql:",notnull"`
	Dust_bin_5  int64 `json:"dust_bin_5" sql:",notnull"`
	Dust_bin_6  int64 `json:"dust_bin_6" sql:",notnull"`
	Dust_bin_7  int64 `json:"dust_bin_7" sql:",notnull"`
	Dust_bin_8  int64 `json:"dust_bin_8" sql:",notnull"`
	Dust_bin_9  int64 `json:"dust_bin_9" sql:",notnull"`
	Dust_bin_10 int64 `json:"dust_bin_10" sql:",notnull"`
	Dust_bin_11 int64 `json:"dust_bin_11" sql:",notnull"`
	Dust_bin_12 int64 `json:"dust_bin_12" sql:",notnull"`
	Dust_bin_13 int64 `json:"dust_bin_13" sql:",notnull"`
	Dust_bin_14 int64 `json:"dust_bin_14" sql:",notnull"`
	Dust_bin_15 int64 `json:"dust_bin_15" sql:",notnull"`
	Dust_bin_16 int64 `json:"dust_bin_16" sql:",notnull"`
	Dust_bin_17 int64 `json:"dust_bin_17" sql:",notnull"`
	Dust_bin_18 int64 `json:"dust_bin_18" sql:",notnull"`
	Dust_bin_19 int64 `json:"dust_bin_19" sql:",notnull"`
	Dust_bin_20 int64 `json:"dust_bin_20" sql:",notnull"`
	Dust_bin_21 int64 `json:"dust_bin_21" sql:",notnull"`

	Hash   string `json:"hash" sql:",notnull"`
	Height int64  `json:"height" sql:",notnull"`

	Native_P2WPKH_outputs_spent            int64 `json:"native_P2WPKH_outputs_spent" sql:",notnull"`
	Native_P2WSH_outputs_spent             int64 `json:"native_P2WSH_outputs_spent" sql:",notnull"`
	Nested_P2WPKH_outputs_spent            int64 `json:"nested_P2WPKH_outputs_spent" sql:",notnull"`
	Nested_P2WSH_outputs_spent             int64 `json:"nested_P2WSH_outputs_spent" sql:",notnull"`
	New_P2WPKH_outputs                     int64 `json:"new_P2WPKH_outputs" sql:",notnull"`
	New_P2WSH_outputs                      int64 `json:"new_P2WSH_outputs" sql:",notnull"`
	Num_batching_txs                       int64 `json:"num_batching_txs" sql:",notnull"`
	Num_consolidating_txs                  int64 `json:"num_consolidating_txs" sql:",notnull"`
	Num_inputs                             int64 `json:"num_inputs" sql:",notnull"`
	Num_outputs                            int64 `json:"num_outputs" sql:",notnull"`
	Num_outputs_consolidated               int64 `json:"num_outputs_consolidated" sql:",notnull"`
	Num_segwit_txs                         int64 `json:"num_segwit_txs" sql:",notnull"`
	Num_txs                                int64 `json:"num_txs" sql:",notnull"`
	Num_txs_creating_P2WPKH                int64 `json:"num_txs_creating_P2WPKH" sql:",notnull"`
	Num_txs_creating_P2WSH                 int64 `json:"num_txs_creating_P2WSH" sql:",notnull"`
	Num_txs_creating_native_segwit_outputs int64 `json:"num_txs_creating_native_segwit_outputs" sql:",notnull"`
	Num_txs_signalling_rbf                 int64 `json:"num_txs_signalling_rbf" sql:",notnull"`
	Output_count_bin_0                     int64 `json:"output_count_bin_0" sql:",notnull"`
	Output_count_bin_1                     int64 `json:"output_count_bin_1" sql:",notnull"`
	Output_count_bin_2                     int64 `json:"output_count_bin_2" sql:",notnull"`
	Output_count_bin_3                     int64 `json:"output_count_bin_3" sql:",notnull"`
	Output_count_bin_4                     int64 `json:"output_count_bin_4" sql:",notnull"`
	Output_count_bin_5                     int64 `json:"output_count_bin_5" sql:",notnull"`
	Output_count_bin_6                     int64 `json:"output_count_bin_6" sql:",notnull"`

	Percent_inputs_consolidated     float64 `json:"percent_inputs_consolidated" sql:",notnull"`
	Percent_new_outs_P2WPKH_outputs float64 `json:"percent_new_outs_P2WPKH_outputs" sql:",notnull"`
	Percent_new_outs_P2WSH_outputs  float64 `json:"percent_new_outs_P2WSH_outputs" sql:",notnull"`
	Percent_new_outs_in_dust_bin_0  float64 `json:"percent_new_outs_in_dust_bin_0" sql:",notnull"`
	Percent_new_outs_in_dust_bin_1  float64 `json:"percent_new_outs_in_dust_bin_1" sql:",notnull"`
	Percent_new_outs_in_dust_bin_2  float64 `json:"percent_new_outs_in_dust_bin_2" sql:",notnull"`
	Percent_new_outs_in_dust_bin_3  float64 `json:"percent_new_outs_in_dust_bin_3" sql:",notnull"`
	Percent_new_outs_in_dust_bin_4  float64 `json:"percent_new_outs_in_dust_bin_4" sql:",notnull"`
	Percent_new_outs_in_dust_bin_5  float64 `json:"percent_new_outs_in_dust_bin_5" sql:",notnull"`
	Percent_new_outs_in_dust_bin_6  float64 `json:"percent_new_outs_in_dust_bin_6" sql:",notnull"`
	Percent_new_outs_in_dust_bin_7  float64 `json:"percent_new_outs_in_dust_bin_7" sql:",notnull"`
	Percent_new_outs_in_dust_bin_8  float64 `json:"percent_new_outs_in_dust_bin_8" sql:",notnull"`
	Percent_new_outs_in_dust_bin_9  float64 `json:"percent_new_outs_in_dust_bin_9" sql:",notnull"`
	Percent_new_outs_in_dust_bin_10 float64 `json:"percent_new_outs_in_dust_bin_10" sql:",notnull"`
	Percent_new_outs_in_dust_bin_11 float64 `json:"percent_new_outs_in_dust_bin_11" sql:",notnull"`
	Percent_new_outs_in_dust_bin_12 float64 `json:"percent_new_outs_in_dust_bin_12" sql:",notnull"`
	Percent_new_outs_in_dust_bin_13 float64 `json:"percent_new_outs_in_dust_bin_13" sql:",notnull"`
	Percent_new_outs_in_dust_bin_14 float64 `json:"percent_new_outs_in_dust_bin_14" sql:",notnull"`
	Percent_new_outs_in_dust_bin_15 float64 `json:"percent_new_outs_in_dust_bin_15" sql:",notnull"`
	Percent_new_outs_in_dust_bin_16 float64 `json:"percent_new_outs_in_dust_bin_16" sql:",notnull"`
	Percent_new_outs_in_dust_bin_17 float64 `json:"percent_new_outs_in_dust_bin_17" sql:",notnull"`
	Percent_new_outs_in_dust_bin_18 float64 `json:"percent_new_outs_in_dust_bin_18" sql:",notnull"`
	Percent_new_outs_in_dust_bin_19 float64 `json:"percent_new_outs_in_dust_bin_19" sql:",notnull"`
	Percent_new_outs_in_dust_bin_20 float64 `json:"percent_new_outs_in_dust_bin_20" sql:",notnull"`
	Percent_new_outs_in_dust_bin_21 float64 `json:"percent_new_outs_in_dust_bin_21" sql:",notnull"`

	Percent_of_inputs_spending_P2WPKH_outputs        float64 `json:"percent_of_inputs_spending_P2WPKH_outputs" sql:",notnull"`
	Percent_of_inputs_spending_P2WSH_outputs         float64 `json:"percent_of_inputs_spending_P2WSH_outputs" sql:",notnull"`
	Percent_of_inputs_spending_native_P2WPKH_outputs float64 `json:"percent_of_inputs_spending_native_P2WPKH_outputs" sql:",notnull"`
	Percent_of_inputs_spending_native_P2WSH_outputs  float64 `json:"percent_of_inputs_spending_native_P2WSH_outputs" sql:",notnull"`
	Percent_of_inputs_spending_native_sw_outputs     float64 `json:"percent_of_inputs_spending_native_sw_outputs" sql:",notnull"`
	Percent_of_inputs_spending_nested_P2WPKH_output  float64 `json:"percent_of_inputs_spending_nested_P2WPKH_output" sql:",notnull"`
	Percent_of_inputs_spending_nested_P2WSH_outputs  float64 `json:"percent_of_inputs_spending_nested_P2WSH_outputs" sql:",notnull"`
	Percent_txs_batching                             float64 `json:"percent_txs_batching" sql:",notnull"`
	Percent_txs_consolidating                        float64 `json:"percent_txs_consolidating" sql:",notnull"`
	Percent_txs_creating_P2WPKH_outputs              float64 `json:"percent_txs_creating_P2WPKH_outputs" sql:",notnull"`
	Percent_txs_creating_P2WSH_outputs               float64 `json:"percent_txs_creating_P2WSH_outputs" sql:",notnull"`
	Percent_txs_creating_native_segwit_outputs       float64 `json:"percent_txs_creating_native_segwit_outputs" sql:",notnull"`
	Percent_txs_native_segwit_over_total_sw_txs      float64 `json:"percent_txs_native_segwit_over_total_sw_txs" sql:",notnull"`
	Percent_txs_signalling_RBF                       float64 `json:"percent_txs_signalling_RBF" sql:",notnull"`
	Percent_txs_spending_P2WPKH_outputs              float64 `json:"percent_txs_spending_P2WPKH_outputs" sql:",notnull"`
	Percent_txs_spending_P2WSH_outputs               float64 `json:"percent_txs_spending_P2WSH_outputs" sql:",notnull"`
	Percent_txs_spending_native_P2WPKH_outputs       float64 `json:"percent_txs_spending_native_P2WPKH_outputs" sql:",notnull"`
	Percent_txs_spending_native_P2WSH_outputs        float64 `json:"percent_txs_spending_native_P2WSH_outputs" sql:",notnull"`
	Percent_txs_spending_native_segwit_outputs       float64 `json:"percent_txs_spending_native_segwit_outputs" sql:",notnull"`
	Percent_txs_spending_nested_segwit_outputs       float64 `json:"percent_txs_spending_nested_segwit_outputs" sql:",notnull"`
	Percent_txs_spending_nested_P2WPKH_outputs       float64 `json:"percent_txs_spending_nested_P2WPKH_outputs" sql:",notnull"`
	Percent_txs_spending_nested_P2WSH_outputs        float64 `json:"percent_txs_spending_nested_P2WSH_outputs" sql:",notnull"`
	Percent_txs_that_are_segwit_txs                  float64 `json:"percent_txs_that_are_segwit_txs" sql:",notnull"`

	Segwit_total_size                  int64 `json:"segwit_total_size" sql:",notnull"`
	Segwit_total_weight                int64 `json:"segwit_total_weight" sql:",notnull"`
	Subsidy                            int64 `json:"subsidy" sql:",notnull"`
	Total_amount_out                   int64 `json:"total_amount_out" sql:",notnull"`
	Total_fee                          int64 `json:"total_fee" sql:",notnull"`
	Total_size                         int64 `json:"total_size" sql:",notnull"`
	Total_weight                       int64 `json:"total_weight" sql:",notnull"`
	Txs_spending_native_p2wpkh_outputs int64 `json:"txs_spending_native_p2wpkh_outputs" sql:",notnull"`
	Txs_spending_native_p2wsh_outputs  int64 `json:"txs_spending_native_p2wsh_outputs" sql:",notnull"`
	Txs_spending_nested_p2wpkh_outputs int64 `json:"txs_spending_nested_p2wpkh_outputs" sql:",notnull"`
	Txs_spending_nested_p2wsh_outputs  int64 `json:"txs_spending_nested_p2wsh_outputs" sql:",notnull"`
	Txs_spending_native_sw_outputs     int64 `json:"txs_spending_native_sw_outputs"`
	Txs_spending_nested_sw_outputs     int64 `json:"txs_spending_nested_sw_outputs"`
	Utxo_increase                      int64 `json:"utxo_increase" sql:",notnull"`
	Utxo_size_increase                 int64 `json:"utxo_size_increase" sql:",notnull"`
	Volume_btc                         int64 `json:"volume_btc" sql:",notnull"`
}

type DashboardDataV2 struct {
	Id int64 `json:"id,omit_empty" sql:",notnull;primary_key"`

	Mto_consolidations int64 `json:"mto_consolidations" sql:",notnull"`
	Mto_output_count   int64 `json:"mto_output_count" sql:",notnull"`
	Mto_avg_fee        int64 `json:"mto_avg_fee" sql:",notnull"`
	Mto_avg_value      int64 `json:"mto_avg_value" sql:",notnull"`

	Time int64 `json:"time" sql:",notnull"`

	Avg_fee      int64 `json:"avg_fee" sql:",notnull"`
	Avg_fee_rate int64 `json:"avg_fee_rate" sql:",notnull"`
	Avg_tx_size  int64 `json:"avg_tx_size" sql:",notnull"`

	Max_fee      int64 `json:"max_fee" sql:",notnull"`
	Max_fee_rate int64 `json:"max_fee_rate" sql:",notnull"`
	Max_tx_size  int64 `json:"max_tx_size" sql:",notnull"`

	Median_fee      int64 `json:"median_fee" sql:",notnull"`
	Median_fee_rate int64 `json:"median_fee_rate" sql:",notnull"`
	Median_tx_size  int64 `json:"median_tx_size" sql:",notnull"`

	Min_fee      int64 `json:"min_fee" sql:",notnull"`
	Min_fee_rate int64 `json:"min_fee_rate" sql:",notnull"`
	Min_tx_size  int64 `json:"min_tx_size" sql:",notnull"`

	// Ninetieth Percentile fee stats
	Per_90th_fee      int64 `json:"per_90th_fee" sql:",notnull"`
	Per_90th_fee_rate int64 `json:"per_90th_fee_rate" sql:",notnull"`
	Per_90th_tx_size  int64 `json:"per_90th_tx_size" sql:",notnull"`

	Batch_range_0 float64 `json:"batch_range_0" sql:",notnull"`
	Batch_range_1 float64 `json:"batch_range_1" sql:",notnull"`
	Batch_range_2 float64 `json:"batch_range_2" sql:",notnull"`
	Batch_range_3 float64 `json:"batch_range_3" sql:",notnull"`
	Batch_range_4 float64 `json:"batch_range_4" sql:",notnull"`
	Batch_range_5 float64 `json:"batch_range_5" sql:",notnull"`
	Batch_range_6 float64 `json:"batch_range_6" sql:",notnull"`

	Block_size int64 `json:"block_size" sql:",notnull"`

	Dust_bin_0  int64 `json:"dust_bin_0" sql:",notnull"`
	Dust_bin_1  int64 `json:"dust_bin_1" sql:",notnull"`
	Dust_bin_2  int64 `json:"dust_bin_2" sql:",notnull"`
	Dust_bin_3  int64 `json:"dust_bin_3" sql:",notnull"`
	Dust_bin_4  int64 `json:"dust_bin_4" sql:",notnull"`
	Dust_bin_5  int64 `json:"dust_bin_5" sql:",notnull"`
	Dust_bin_6  int64 `json:"dust_bin_6" sql:",notnull"`
	Dust_bin_7  int64 `json:"dust_bin_7" sql:",notnull"`
	Dust_bin_8  int64 `json:"dust_bin_8" sql:",notnull"`
	Dust_bin_9  int64 `json:"dust_bin_9" sql:",notnull"`
	Dust_bin_10 int64 `json:"dust_bin_10" sql:",notnull"`
	Dust_bin_11 int64 `json:"dust_bin_11" sql:",notnull"`
	Dust_bin_12 int64 `json:"dust_bin_12" sql:",notnull"`
	Dust_bin_13 int64 `json:"dust_bin_13" sql:",notnull"`
	Dust_bin_14 int64 `json:"dust_bin_14" sql:",notnull"`
	Dust_bin_15 int64 `json:"dust_bin_15" sql:",notnull"`
	Dust_bin_16 int64 `json:"dust_bin_16" sql:",notnull"`
	Dust_bin_17 int64 `json:"dust_bin_17" sql:",notnull"`
	Dust_bin_18 int64 `json:"dust_bin_18" sql:",notnull"`
	Dust_bin_19 int64 `json:"dust_bin_19" sql:",notnull"`
	Dust_bin_20 int64 `json:"dust_bin_20" sql:",notnull"`
	Dust_bin_21 int64 `json:"dust_bin_21" sql:",notnull"`

	Hash   string `json:"hash" sql:",notnull"`
	Height int64  `json:"height" sql:",notnull"`

	Native_P2WPKH_outputs_spent int64 `json:"native_P2WPKH_outputs_spent" sql:",notnull"`
	Native_P2WSH_outputs_spent  int64 `json:"native_P2WSH_outputs_spent" sql:",notnull"`
	Nested_P2WPKH_outputs_spent int64 `json:"nested_P2WPKH_outputs_spent" sql:",notnull"`
	Nested_P2WSH_outputs_spent  int64 `json:"nested_P2WSH_outputs_spent" sql:",notnull"`
	New_P2WPKH_outputs          int64 `json:"new_P2WPKH_outputs" sql:",notnull"`
	New_P2WSH_outputs           int64 `json:"new_P2WSH_outputs" sql:",notnull"`

	Value_of_native_P2WPKH_outputs_spent int64 `json:"value_of_native_P2WPKH_outputs_spent" sql:",notnull"`
	Value_of_native_P2WSH_outputs_spent  int64 `json:"value_of_native_P2WSH_outputs_spent" sql:",notnull"`
	Value_of_nested_P2WPKH_outputs_spent int64 `json:"value_of_nested_P2WPKH_outputs_spent" sql:",notnull"`
	Value_of_nested_P2WSH_outputs_spent  int64 `json:"value_of_nested_P2WSH_outputs_spent" sql:",notnull"`
	Value_of_new_P2WPKH_outputs          int64 `json:"value_of_new_P2WPKH_outputs" sql:",notnull"`
	Value_of_new_P2WSH_outputs           int64 `json:"value_of_new_P2WSH_outputs" sql:",notnull"`

	Num_batching_txs                       int64 `json:"num_batching_txs" sql:",notnull"`
	Num_consolidating_txs                  int64 `json:"num_consolidating_txs" sql:",notnull"`
	Num_inputs                             int64 `json:"num_inputs" sql:",notnull"`
	Num_outputs                            int64 `json:"num_outputs" sql:",notnull"`
	Num_outputs_consolidated               int64 `json:"num_outputs_consolidated" sql:",notnull"`
	Num_segwit_txs                         int64 `json:"num_segwit_txs" sql:",notnull"`
	Num_txs                                int64 `json:"num_txs" sql:",notnull"`
	Num_txs_creating_P2WPKH                int64 `json:"num_txs_creating_P2WPKH" sql:",notnull"`
	Num_txs_creating_P2WSH                 int64 `json:"num_txs_creating_P2WSH" sql:",notnull"`
	Num_txs_creating_native_segwit_outputs int64 `json:"num_txs_creating_native_segwit_outputs" sql:",notnull"`
	TxsSignallingOptInRBF                  int64 `json:"txs_signalling_opt_in_rbf" sql:",notnull"`
	Output_count_bin_0                     int64 `json:"output_count_bin_0" sql:",notnull"`
	Output_count_bin_1                     int64 `json:"output_count_bin_1" sql:",notnull"`
	Output_count_bin_2                     int64 `json:"output_count_bin_2" sql:",notnull"`
	Output_count_bin_3                     int64 `json:"output_count_bin_3" sql:",notnull"`
	Output_count_bin_4                     int64 `json:"output_count_bin_4" sql:",notnull"`
	Output_count_bin_5                     int64 `json:"output_count_bin_5" sql:",notnull"`
	Output_count_bin_6                     int64 `json:"output_count_bin_6" sql:",notnull"`

	Percent_inputs_consolidated     float64 `json:"percent_inputs_consolidated" sql:",notnull"`
	Percent_new_outs_P2WPKH_outputs float64 `json:"percent_new_outs_P2WPKH_outputs" sql:",notnull"`
	Percent_new_outs_P2WSH_outputs  float64 `json:"percent_new_outs_P2WSH_outputs" sql:",notnull"`
	Percent_new_outs_in_dust_bin_0  float64 `json:"percent_new_outs_in_dust_bin_0" sql:",notnull"`
	Percent_new_outs_in_dust_bin_1  float64 `json:"percent_new_outs_in_dust_bin_1" sql:",notnull"`
	Percent_new_outs_in_dust_bin_2  float64 `json:"percent_new_outs_in_dust_bin_2" sql:",notnull"`
	Percent_new_outs_in_dust_bin_3  float64 `json:"percent_new_outs_in_dust_bin_3" sql:",notnull"`
	Percent_new_outs_in_dust_bin_4  float64 `json:"percent_new_outs_in_dust_bin_4" sql:",notnull"`
	Percent_new_outs_in_dust_bin_5  float64 `json:"percent_new_outs_in_dust_bin_5" sql:",notnull"`
	Percent_new_outs_in_dust_bin_6  float64 `json:"percent_new_outs_in_dust_bin_6" sql:",notnull"`
	Percent_new_outs_in_dust_bin_7  float64 `json:"percent_new_outs_in_dust_bin_7" sql:",notnull"`
	Percent_new_outs_in_dust_bin_8  float64 `json:"percent_new_outs_in_dust_bin_8" sql:",notnull"`
	Percent_new_outs_in_dust_bin_9  float64 `json:"percent_new_outs_in_dust_bin_9" sql:",notnull"`
	Percent_new_outs_in_dust_bin_10 float64 `json:"percent_new_outs_in_dust_bin_10" sql:",notnull"`
	Percent_new_outs_in_dust_bin_11 float64 `json:"percent_new_outs_in_dust_bin_11" sql:",notnull"`
	Percent_new_outs_in_dust_bin_12 float64 `json:"percent_new_outs_in_dust_bin_12" sql:",notnull"`
	Percent_new_outs_in_dust_bin_13 float64 `json:"percent_new_outs_in_dust_bin_13" sql:",notnull"`
	Percent_new_outs_in_dust_bin_14 float64 `json:"percent_new_outs_in_dust_bin_14" sql:",notnull"`
	Percent_new_outs_in_dust_bin_15 float64 `json:"percent_new_outs_in_dust_bin_15" sql:",notnull"`
	Percent_new_outs_in_dust_bin_16 float64 `json:"percent_new_outs_in_dust_bin_16" sql:",notnull"`
	Percent_new_outs_in_dust_bin_17 float64 `json:"percent_new_outs_in_dust_bin_17" sql:",notnull"`
	Percent_new_outs_in_dust_bin_18 float64 `json:"percent_new_outs_in_dust_bin_18" sql:",notnull"`
	Percent_new_outs_in_dust_bin_19 float64 `json:"percent_new_outs_in_dust_bin_19" sql:",notnull"`
	Percent_new_outs_in_dust_bin_20 float64 `json:"percent_new_outs_in_dust_bin_20" sql:",notnull"`
	Percent_new_outs_in_dust_bin_21 float64 `json:"percent_new_outs_in_dust_bin_21" sql:",notnull"`

	Percent_of_inputs_spending_P2WPKH_outputs        float64 `json:"percent_of_inputs_spending_P2WPKH_outputs" sql:",notnull"`
	Percent_of_inputs_spending_P2WSH_outputs         float64 `json:"percent_of_inputs_spending_P2WSH_outputs" sql:",notnull"`
	Percent_of_inputs_spending_native_P2WPKH_outputs float64 `json:"percent_of_inputs_spending_native_P2WPKH_outputs" sql:",notnull"`
	Percent_of_inputs_spending_native_P2WSH_outputs  float64 `json:"percent_of_inputs_spending_native_P2WSH_outputs" sql:",notnull"`
	Percent_of_inputs_spending_native_sw_outputs     float64 `json:"percent_of_inputs_spending_native_sw_outputs" sql:",notnull"`
	Percent_of_inputs_spending_nested_P2WPKH_output  float64 `json:"percent_of_inputs_spending_nested_P2WPKH_output" sql:",notnull"`
	Percent_of_inputs_spending_nested_P2WSH_outputs  float64 `json:"percent_of_inputs_spending_nested_P2WSH_outputs" sql:",notnull"`
	Percent_txs_batching                             float64 `json:"percent_txs_batching" sql:",notnull"`
	Percent_txs_consolidating                        float64 `json:"percent_txs_consolidating" sql:",notnull"`
	Percent_txs_creating_P2WPKH_outputs              float64 `json:"percent_txs_creating_P2WPKH_outputs" sql:",notnull"`
	Percent_txs_creating_P2WSH_outputs               float64 `json:"percent_txs_creating_P2WSH_outputs" sql:",notnull"`
	Percent_txs_creating_native_segwit_outputs       float64 `json:"percent_txs_creating_native_segwit_outputs" sql:",notnull"`
	Percent_txs_native_segwit_over_total_sw_txs      float64 `json:"percent_txs_native_segwit_over_total_sw_txs" sql:",notnull"`
	Percent_txs_signalling_opt_in_RBF                float64 `json:"percent_txs_signalling_opt_in_RBF" sql:",notnull"`
	Percent_txs_spending_P2WPKH_outputs              float64 `json:"percent_txs_spending_P2WPKH_outputs" sql:",notnull"`
	Percent_txs_spending_P2WSH_outputs               float64 `json:"percent_txs_spending_P2WSH_outputs" sql:",notnull"`
	Percent_txs_spending_native_P2WPKH_outputs       float64 `json:"percent_txs_spending_native_P2WPKH_outputs" sql:",notnull"`
	Percent_txs_spending_native_P2WSH_outputs        float64 `json:"percent_txs_spending_native_P2WSH_outputs" sql:",notnull"`
	Percent_txs_spending_native_segwit_outputs       float64 `json:"percent_txs_spending_native_segwit_outputs" sql:",notnull"`
	Percent_txs_spending_nested_segwit_outputs       float64 `json:"percent_txs_spending_nested_segwit_outputs" sql:",notnull"`
	Percent_txs_spending_nested_P2WPKH_outputs       float64 `json:"percent_txs_spending_nested_P2WPKH_outputs" sql:",notnull"`
	Percent_txs_spending_nested_P2WSH_outputs        float64 `json:"percent_txs_spending_nested_P2WSH_outputs" sql:",notnull"`
	Percent_txs_that_are_segwit_txs                  float64 `json:"percent_txs_that_are_segwit_txs" sql:",notnull"`

	Segwit_total_size                  int64 `json:"segwit_total_size" sql:",notnull"`
	Segwit_total_weight                int64 `json:"segwit_total_weight" sql:",notnull"`
	Subsidy                            int64 `json:"subsidy" sql:",notnull"`
	Total_amount_out                   int64 `json:"total_amount_out" sql:",notnull"`
	Total_fee                          int64 `json:"total_fee" sql:",notnull"`
	Total_size                         int64 `json:"total_size" sql:",notnull"`
	Total_weight                       int64 `json:"total_weight" sql:",notnull"`
	Txs_spending_native_p2wpkh_outputs int64 `json:"txs_spending_native_p2wpkh_outputs" sql:",notnull"`
	Txs_spending_native_p2wsh_outputs  int64 `json:"txs_spending_native_p2wsh_outputs" sql:",notnull"`
	Txs_spending_nested_p2wpkh_outputs int64 `json:"txs_spending_nested_p2wpkh_outputs" sql:",notnull"`
	Txs_spending_nested_p2wsh_outputs  int64 `json:"txs_spending_nested_p2wsh_outputs" sql:",notnull"`
	Txs_spending_native_sw_outputs     int64 `json:"txs_spending_native_sw_outputs" sql:",notnull"`
	Txs_spending_nested_sw_outputs     int64 `json:"txs_spending_nested_sw_outputs" sql:",notnull"`
	Utxo_increase                      int64 `json:"utxo_increase" sql:",notnull"`
	Utxo_size_increase                 int64 `json:"utxo_size_increase" sql:",notnull"`
	Volume_btc                         int64 `json:"volume_btc" sql:",notnull"`
}
