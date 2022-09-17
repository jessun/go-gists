package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPgBackupLogArchiveStatusMgr struct {
	*_BaseMgr
}

// AllVirtualPgBackupLogArchiveStatusMgr open func
func AllVirtualPgBackupLogArchiveStatusMgr(db *gorm.DB) *_AllVirtualPgBackupLogArchiveStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPgBackupLogArchiveStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPgBackupLogArchiveStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_pg_backup_log_archive_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetTableName() string {
	return "__all_virtual_pg_backup_log_archive_status"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) Reset() *_AllVirtualPgBackupLogArchiveStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) Get() (result AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) Gets() (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithLogArchiveRound log_archive_round获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_round"] = logArchiveRound })
}

// WithLogArchiveStartTs log_archive_start_ts获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithLogArchiveStartTs(logArchiveStartTs int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_start_ts"] = logArchiveStartTs })
}

// WithLogArchiveStatus log_archive_status获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithLogArchiveStatus(logArchiveStatus int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_status"] = logArchiveStatus })
}

// WithLogArchiveCurLogID log_archive_cur_log_id获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithLogArchiveCurLogID(logArchiveCurLogID int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_cur_log_id"] = logArchiveCurLogID })
}

// WithLogArchiveCurTs log_archive_cur_ts获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithLogArchiveCurTs(logArchiveCurTs int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_cur_ts"] = logArchiveCurTs })
}

// WithMaxLogID max_log_id获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithMaxLogID(maxLogID int64) Option {
	return optionFunc(func(o *options) { o.query["max_log_id"] = maxLogID })
}

// WithMaxLogTs max_log_ts获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithMaxLogTs(maxLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["max_log_ts"] = maxLogTs })
}

// WithCurPieceID cur_piece_id获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) WithCurPieceID(curPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["cur_piece_id"] = curPieceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetByOption(opts ...Option) (result AllVirtualPgBackupLogArchiveStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetByOptions(opts ...Option) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPgBackupLogArchiveStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where(options.query)
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
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromTableID(tableID int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过log_archive_round获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_round` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_round` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromLogArchiveStartTs 通过log_archive_start_ts获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromLogArchiveStartTs(logArchiveStartTs int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_start_ts` = ?", logArchiveStartTs).Find(&results).Error

	return
}

// GetBatchFromLogArchiveStartTs 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromLogArchiveStartTs(logArchiveStartTss []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_start_ts` IN (?)", logArchiveStartTss).Find(&results).Error

	return
}

// GetFromLogArchiveStatus 通过log_archive_status获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromLogArchiveStatus(logArchiveStatus int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_status` = ?", logArchiveStatus).Find(&results).Error

	return
}

// GetBatchFromLogArchiveStatus 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromLogArchiveStatus(logArchiveStatuss []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_status` IN (?)", logArchiveStatuss).Find(&results).Error

	return
}

// GetFromLogArchiveCurLogID 通过log_archive_cur_log_id获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromLogArchiveCurLogID(logArchiveCurLogID int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_cur_log_id` = ?", logArchiveCurLogID).Find(&results).Error

	return
}

// GetBatchFromLogArchiveCurLogID 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromLogArchiveCurLogID(logArchiveCurLogIDs []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_cur_log_id` IN (?)", logArchiveCurLogIDs).Find(&results).Error

	return
}

// GetFromLogArchiveCurTs 通过log_archive_cur_ts获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromLogArchiveCurTs(logArchiveCurTs int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_cur_ts` = ?", logArchiveCurTs).Find(&results).Error

	return
}

// GetBatchFromLogArchiveCurTs 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromLogArchiveCurTs(logArchiveCurTss []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`log_archive_cur_ts` IN (?)", logArchiveCurTss).Find(&results).Error

	return
}

// GetFromMaxLogID 通过max_log_id获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromMaxLogID(maxLogID int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`max_log_id` = ?", maxLogID).Find(&results).Error

	return
}

// GetBatchFromMaxLogID 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromMaxLogID(maxLogIDs []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`max_log_id` IN (?)", maxLogIDs).Find(&results).Error

	return
}

// GetFromMaxLogTs 通过max_log_ts获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromMaxLogTs(maxLogTs int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`max_log_ts` = ?", maxLogTs).Find(&results).Error

	return
}

// GetBatchFromMaxLogTs 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromMaxLogTs(maxLogTss []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`max_log_ts` IN (?)", maxLogTss).Find(&results).Error

	return
}

// GetFromCurPieceID 通过cur_piece_id获取内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetFromCurPieceID(curPieceID int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`cur_piece_id` = ?", curPieceID).Find(&results).Error

	return
}

// GetBatchFromCurPieceID 批量查找
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) GetBatchFromCurPieceID(curPieceIDs []int64) (results []*AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`cur_piece_id` IN (?)", curPieceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPgBackupLogArchiveStatusMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64, tableID int64, partitionID int64) (result AllVirtualPgBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPgBackupLogArchiveStatus{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", svrIP, svrPort, tenantID, tableID, partitionID).First(&result).Error

	return
}
