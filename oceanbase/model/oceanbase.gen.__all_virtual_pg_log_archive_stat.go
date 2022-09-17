package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPgLogArchiveStatMgr struct {
	*_BaseMgr
}

// AllVirtualPgLogArchiveStatMgr open func
func AllVirtualPgLogArchiveStatMgr(db *gorm.DB) *_AllVirtualPgLogArchiveStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPgLogArchiveStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPgLogArchiveStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_pg_log_archive_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPgLogArchiveStatMgr) GetTableName() string {
	return "__all_virtual_pg_log_archive_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPgLogArchiveStatMgr) Reset() *_AllVirtualPgLogArchiveStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPgLogArchiveStatMgr) Get() (result AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPgLogArchiveStatMgr) Gets() (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPgLogArchiveStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithLogArchiveRound log_archive_round获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_round"] = logArchiveRound })
}

// WithEpoch epoch获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithEpoch(epoch int64) Option {
	return optionFunc(func(o *options) { o.query["epoch"] = epoch })
}

// WithBeenDeleted been_deleted获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithBeenDeleted(beenDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["been_deleted"] = beenDeleted })
}

// WithIsFirstRecordFinish is_first_record_finish获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithIsFirstRecordFinish(isFirstRecordFinish int8) Option {
	return optionFunc(func(o *options) { o.query["is_first_record_finish"] = isFirstRecordFinish })
}

// WithEncountError encount_error获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithEncountError(encountError int8) Option {
	return optionFunc(func(o *options) { o.query["encount_error"] = encountError })
}

// WithCurrentIlogID current_ilog_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithCurrentIlogID(currentIlogID int64) Option {
	return optionFunc(func(o *options) { o.query["current_ilog_id"] = currentIlogID })
}

// WithMaxLogID max_log_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithMaxLogID(maxLogID int64) Option {
	return optionFunc(func(o *options) { o.query["max_log_id"] = maxLogID })
}

// WithRoundStartLogID round_start_log_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithRoundStartLogID(roundStartLogID int64) Option {
	return optionFunc(func(o *options) { o.query["round_start_log_id"] = roundStartLogID })
}

// WithRoundStartTs round_start_ts获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithRoundStartTs(roundStartTs int64) Option {
	return optionFunc(func(o *options) { o.query["round_start_ts"] = roundStartTs })
}

// WithRoundSnapshotVersion round_snapshot_version获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithRoundSnapshotVersion(roundSnapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["round_snapshot_version"] = roundSnapshotVersion })
}

// WithCurStartLogID cur_start_log_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithCurStartLogID(curStartLogID int64) Option {
	return optionFunc(func(o *options) { o.query["cur_start_log_id"] = curStartLogID })
}

// WithFetcherMaxSplitLogID fetcher_max_split_log_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithFetcherMaxSplitLogID(fetcherMaxSplitLogID int64) Option {
	return optionFunc(func(o *options) { o.query["fetcher_max_split_log_id"] = fetcherMaxSplitLogID })
}

// WithClogSplitMaxLogID clog_split_max_log_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithClogSplitMaxLogID(clogSplitMaxLogID int64) Option {
	return optionFunc(func(o *options) { o.query["clog_split_max_log_id"] = clogSplitMaxLogID })
}

// WithClogSplitMaxLogTs clog_split_max_log_ts获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithClogSplitMaxLogTs(clogSplitMaxLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["clog_split_max_log_ts"] = clogSplitMaxLogTs })
}

// WithClogSplitCheckpointTs clog_split_checkpoint_ts获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithClogSplitCheckpointTs(clogSplitCheckpointTs int64) Option {
	return optionFunc(func(o *options) { o.query["clog_split_checkpoint_ts"] = clogSplitCheckpointTs })
}

// WithMaxArchivedLogID max_archived_log_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithMaxArchivedLogID(maxArchivedLogID int64) Option {
	return optionFunc(func(o *options) { o.query["max_archived_log_id"] = maxArchivedLogID })
}

// WithMaxArchivedLogTs max_archived_log_ts获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithMaxArchivedLogTs(maxArchivedLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["max_archived_log_ts"] = maxArchivedLogTs })
}

// WithMaxArchivedCheckpointTs max_archived_checkpoint_ts获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithMaxArchivedCheckpointTs(maxArchivedCheckpointTs int64) Option {
	return optionFunc(func(o *options) { o.query["max_archived_checkpoint_ts"] = maxArchivedCheckpointTs })
}

// WithClogEpoch clog_epoch获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithClogEpoch(clogEpoch int64) Option {
	return optionFunc(func(o *options) { o.query["clog_epoch"] = clogEpoch })
}

// WithClogAccumChecksum clog_accum_checksum获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithClogAccumChecksum(clogAccumChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["clog_accum_checksum"] = clogAccumChecksum })
}

// WithCurIndexFileID cur_index_file_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithCurIndexFileID(curIndexFileID int64) Option {
	return optionFunc(func(o *options) { o.query["cur_index_file_id"] = curIndexFileID })
}

// WithCurIndexFileOffset cur_index_file_offset获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithCurIndexFileOffset(curIndexFileOffset int64) Option {
	return optionFunc(func(o *options) { o.query["cur_index_file_offset"] = curIndexFileOffset })
}

// WithCurDataFileID cur_data_file_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithCurDataFileID(curDataFileID int64) Option {
	return optionFunc(func(o *options) { o.query["cur_data_file_id"] = curDataFileID })
}

// WithCurDataFileOffset cur_data_file_offset获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithCurDataFileOffset(curDataFileOffset int64) Option {
	return optionFunc(func(o *options) { o.query["cur_data_file_offset"] = curDataFileOffset })
}

// WithClogSplitTaskNum clog_split_task_num获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithClogSplitTaskNum(clogSplitTaskNum int64) Option {
	return optionFunc(func(o *options) { o.query["clog_split_task_num"] = clogSplitTaskNum })
}

// WithSendTaskNum send_task_num获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithSendTaskNum(sendTaskNum int64) Option {
	return optionFunc(func(o *options) { o.query["send_task_num"] = sendTaskNum })
}

// WithCurPieceID cur_piece_id获取
func (obj *_AllVirtualPgLogArchiveStatMgr) WithCurPieceID(curPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["cur_piece_id"] = curPieceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPgLogArchiveStatMgr) GetByOption(opts ...Option) (result AllVirtualPgLogArchiveStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPgLogArchiveStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualPgLogArchiveStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPgLogArchiveStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPgLogArchiveStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过log_archive_round获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`log_archive_round` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`log_archive_round` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromEpoch 通过epoch获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromEpoch(epoch int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`epoch` = ?", epoch).Find(&results).Error

	return
}

// GetBatchFromEpoch 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromEpoch(epochs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`epoch` IN (?)", epochs).Find(&results).Error

	return
}

// GetFromBeenDeleted 通过been_deleted获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromBeenDeleted(beenDeleted int8) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`been_deleted` = ?", beenDeleted).Find(&results).Error

	return
}

// GetBatchFromBeenDeleted 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromBeenDeleted(beenDeleteds []int8) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`been_deleted` IN (?)", beenDeleteds).Find(&results).Error

	return
}

// GetFromIsFirstRecordFinish 通过is_first_record_finish获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromIsFirstRecordFinish(isFirstRecordFinish int8) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`is_first_record_finish` = ?", isFirstRecordFinish).Find(&results).Error

	return
}

// GetBatchFromIsFirstRecordFinish 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromIsFirstRecordFinish(isFirstRecordFinishs []int8) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`is_first_record_finish` IN (?)", isFirstRecordFinishs).Find(&results).Error

	return
}

// GetFromEncountError 通过encount_error获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromEncountError(encountError int8) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`encount_error` = ?", encountError).Find(&results).Error

	return
}

// GetBatchFromEncountError 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromEncountError(encountErrors []int8) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`encount_error` IN (?)", encountErrors).Find(&results).Error

	return
}

// GetFromCurrentIlogID 通过current_ilog_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromCurrentIlogID(currentIlogID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`current_ilog_id` = ?", currentIlogID).Find(&results).Error

	return
}

// GetBatchFromCurrentIlogID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromCurrentIlogID(currentIlogIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`current_ilog_id` IN (?)", currentIlogIDs).Find(&results).Error

	return
}

// GetFromMaxLogID 通过max_log_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromMaxLogID(maxLogID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`max_log_id` = ?", maxLogID).Find(&results).Error

	return
}

// GetBatchFromMaxLogID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromMaxLogID(maxLogIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`max_log_id` IN (?)", maxLogIDs).Find(&results).Error

	return
}

// GetFromRoundStartLogID 通过round_start_log_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromRoundStartLogID(roundStartLogID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`round_start_log_id` = ?", roundStartLogID).Find(&results).Error

	return
}

// GetBatchFromRoundStartLogID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromRoundStartLogID(roundStartLogIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`round_start_log_id` IN (?)", roundStartLogIDs).Find(&results).Error

	return
}

// GetFromRoundStartTs 通过round_start_ts获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromRoundStartTs(roundStartTs int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`round_start_ts` = ?", roundStartTs).Find(&results).Error

	return
}

// GetBatchFromRoundStartTs 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromRoundStartTs(roundStartTss []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`round_start_ts` IN (?)", roundStartTss).Find(&results).Error

	return
}

// GetFromRoundSnapshotVersion 通过round_snapshot_version获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromRoundSnapshotVersion(roundSnapshotVersion int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`round_snapshot_version` = ?", roundSnapshotVersion).Find(&results).Error

	return
}

// GetBatchFromRoundSnapshotVersion 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromRoundSnapshotVersion(roundSnapshotVersions []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`round_snapshot_version` IN (?)", roundSnapshotVersions).Find(&results).Error

	return
}

// GetFromCurStartLogID 通过cur_start_log_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromCurStartLogID(curStartLogID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_start_log_id` = ?", curStartLogID).Find(&results).Error

	return
}

// GetBatchFromCurStartLogID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromCurStartLogID(curStartLogIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_start_log_id` IN (?)", curStartLogIDs).Find(&results).Error

	return
}

// GetFromFetcherMaxSplitLogID 通过fetcher_max_split_log_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromFetcherMaxSplitLogID(fetcherMaxSplitLogID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`fetcher_max_split_log_id` = ?", fetcherMaxSplitLogID).Find(&results).Error

	return
}

// GetBatchFromFetcherMaxSplitLogID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromFetcherMaxSplitLogID(fetcherMaxSplitLogIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`fetcher_max_split_log_id` IN (?)", fetcherMaxSplitLogIDs).Find(&results).Error

	return
}

// GetFromClogSplitMaxLogID 通过clog_split_max_log_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromClogSplitMaxLogID(clogSplitMaxLogID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_split_max_log_id` = ?", clogSplitMaxLogID).Find(&results).Error

	return
}

// GetBatchFromClogSplitMaxLogID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromClogSplitMaxLogID(clogSplitMaxLogIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_split_max_log_id` IN (?)", clogSplitMaxLogIDs).Find(&results).Error

	return
}

// GetFromClogSplitMaxLogTs 通过clog_split_max_log_ts获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromClogSplitMaxLogTs(clogSplitMaxLogTs int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_split_max_log_ts` = ?", clogSplitMaxLogTs).Find(&results).Error

	return
}

// GetBatchFromClogSplitMaxLogTs 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromClogSplitMaxLogTs(clogSplitMaxLogTss []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_split_max_log_ts` IN (?)", clogSplitMaxLogTss).Find(&results).Error

	return
}

// GetFromClogSplitCheckpointTs 通过clog_split_checkpoint_ts获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromClogSplitCheckpointTs(clogSplitCheckpointTs int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_split_checkpoint_ts` = ?", clogSplitCheckpointTs).Find(&results).Error

	return
}

// GetBatchFromClogSplitCheckpointTs 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromClogSplitCheckpointTs(clogSplitCheckpointTss []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_split_checkpoint_ts` IN (?)", clogSplitCheckpointTss).Find(&results).Error

	return
}

// GetFromMaxArchivedLogID 通过max_archived_log_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromMaxArchivedLogID(maxArchivedLogID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`max_archived_log_id` = ?", maxArchivedLogID).Find(&results).Error

	return
}

// GetBatchFromMaxArchivedLogID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromMaxArchivedLogID(maxArchivedLogIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`max_archived_log_id` IN (?)", maxArchivedLogIDs).Find(&results).Error

	return
}

// GetFromMaxArchivedLogTs 通过max_archived_log_ts获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromMaxArchivedLogTs(maxArchivedLogTs int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`max_archived_log_ts` = ?", maxArchivedLogTs).Find(&results).Error

	return
}

// GetBatchFromMaxArchivedLogTs 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromMaxArchivedLogTs(maxArchivedLogTss []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`max_archived_log_ts` IN (?)", maxArchivedLogTss).Find(&results).Error

	return
}

// GetFromMaxArchivedCheckpointTs 通过max_archived_checkpoint_ts获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromMaxArchivedCheckpointTs(maxArchivedCheckpointTs int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`max_archived_checkpoint_ts` = ?", maxArchivedCheckpointTs).Find(&results).Error

	return
}

// GetBatchFromMaxArchivedCheckpointTs 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromMaxArchivedCheckpointTs(maxArchivedCheckpointTss []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`max_archived_checkpoint_ts` IN (?)", maxArchivedCheckpointTss).Find(&results).Error

	return
}

// GetFromClogEpoch 通过clog_epoch获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromClogEpoch(clogEpoch int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_epoch` = ?", clogEpoch).Find(&results).Error

	return
}

// GetBatchFromClogEpoch 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromClogEpoch(clogEpochs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_epoch` IN (?)", clogEpochs).Find(&results).Error

	return
}

// GetFromClogAccumChecksum 通过clog_accum_checksum获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromClogAccumChecksum(clogAccumChecksum int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_accum_checksum` = ?", clogAccumChecksum).Find(&results).Error

	return
}

// GetBatchFromClogAccumChecksum 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromClogAccumChecksum(clogAccumChecksums []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_accum_checksum` IN (?)", clogAccumChecksums).Find(&results).Error

	return
}

// GetFromCurIndexFileID 通过cur_index_file_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromCurIndexFileID(curIndexFileID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_index_file_id` = ?", curIndexFileID).Find(&results).Error

	return
}

// GetBatchFromCurIndexFileID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromCurIndexFileID(curIndexFileIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_index_file_id` IN (?)", curIndexFileIDs).Find(&results).Error

	return
}

// GetFromCurIndexFileOffset 通过cur_index_file_offset获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromCurIndexFileOffset(curIndexFileOffset int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_index_file_offset` = ?", curIndexFileOffset).Find(&results).Error

	return
}

// GetBatchFromCurIndexFileOffset 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromCurIndexFileOffset(curIndexFileOffsets []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_index_file_offset` IN (?)", curIndexFileOffsets).Find(&results).Error

	return
}

// GetFromCurDataFileID 通过cur_data_file_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromCurDataFileID(curDataFileID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_data_file_id` = ?", curDataFileID).Find(&results).Error

	return
}

// GetBatchFromCurDataFileID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromCurDataFileID(curDataFileIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_data_file_id` IN (?)", curDataFileIDs).Find(&results).Error

	return
}

// GetFromCurDataFileOffset 通过cur_data_file_offset获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromCurDataFileOffset(curDataFileOffset int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_data_file_offset` = ?", curDataFileOffset).Find(&results).Error

	return
}

// GetBatchFromCurDataFileOffset 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromCurDataFileOffset(curDataFileOffsets []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_data_file_offset` IN (?)", curDataFileOffsets).Find(&results).Error

	return
}

// GetFromClogSplitTaskNum 通过clog_split_task_num获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromClogSplitTaskNum(clogSplitTaskNum int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_split_task_num` = ?", clogSplitTaskNum).Find(&results).Error

	return
}

// GetBatchFromClogSplitTaskNum 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromClogSplitTaskNum(clogSplitTaskNums []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`clog_split_task_num` IN (?)", clogSplitTaskNums).Find(&results).Error

	return
}

// GetFromSendTaskNum 通过send_task_num获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromSendTaskNum(sendTaskNum int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`send_task_num` = ?", sendTaskNum).Find(&results).Error

	return
}

// GetBatchFromSendTaskNum 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromSendTaskNum(sendTaskNums []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`send_task_num` IN (?)", sendTaskNums).Find(&results).Error

	return
}

// GetFromCurPieceID 通过cur_piece_id获取内容
func (obj *_AllVirtualPgLogArchiveStatMgr) GetFromCurPieceID(curPieceID int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_piece_id` = ?", curPieceID).Find(&results).Error

	return
}

// GetBatchFromCurPieceID 批量查找
func (obj *_AllVirtualPgLogArchiveStatMgr) GetBatchFromCurPieceID(curPieceIDs []int64) (results []*AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`cur_piece_id` IN (?)", curPieceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPgLogArchiveStatMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64, tableID int64, partitionID int64) (result AllVirtualPgLogArchiveStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgLogArchiveStat{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", svrIP, svrPort, tenantID, tableID, partitionID).First(&result).Error

	return
}
