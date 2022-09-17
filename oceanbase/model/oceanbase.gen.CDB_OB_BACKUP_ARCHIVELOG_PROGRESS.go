package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CdbObBackupArchivelogProgressMgr struct {
	*_BaseMgr
}

// CdbObBackupArchivelogProgressMgr open func
func CdbObBackupArchivelogProgressMgr(db *gorm.DB) *_CdbObBackupArchivelogProgressMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupArchivelogProgressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupArchivelogProgressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_ARCHIVELOG_PROGRESS"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupArchivelogProgressMgr) GetTableName() string {
	return "CDB_OB_BACKUP_ARCHIVELOG_PROGRESS"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupArchivelogProgressMgr) Reset() *_CdbObBackupArchivelogProgressMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupArchivelogProgressMgr) Get() (result CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupArchivelogProgressMgr) Gets() (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupArchivelogProgressMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithLogArchiveRound LOG_ARCHIVE_ROUND获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["LOG_ARCHIVE_ROUND"] = logArchiveRound })
}

// WithCurPieceID CUR_PIECE_ID获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithCurPieceID(curPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["CUR_PIECE_ID"] = curPieceID })
}

// WithSvrIP SVR_IP获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTableID TABLE_ID获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithMinFirstTime MIN_FIRST_TIME获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithMinFirstTime(minFirstTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["MIN_FIRST_TIME"] = minFirstTime })
}

// WithMaxNextTime MAX_NEXT_TIME获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithMaxNextTime(maxNextTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["MAX_NEXT_TIME"] = maxNextTime })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupArchivelogProgressMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupArchivelogProgressMgr) GetByOption(opts ...Option) (result CdbObBackupArchivelogProgress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupArchivelogProgressMgr) GetByOptions(opts ...Option) (results []*CdbObBackupArchivelogProgress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupArchivelogProgressMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupArchivelogProgress, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where(options.query)
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

// GetFromIncarnation 通过INCARNATION获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过LOG_ARCHIVE_ROUND获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`LOG_ARCHIVE_ROUND` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`LOG_ARCHIVE_ROUND` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromCurPieceID 通过CUR_PIECE_ID获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromCurPieceID(curPieceID int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`CUR_PIECE_ID` = ?", curPieceID).Find(&results).Error

	return
}

// GetBatchFromCurPieceID 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromCurPieceID(curPieceIDs []int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`CUR_PIECE_ID` IN (?)", curPieceIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过SVR_IP获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromSvrIP(svrIP string) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromSvrIP(svrIPs []string) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过SVR_PORT获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromSvrPort(svrPort int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过TABLE_ID获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromTableID(tableID int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromTableID(tableIDs []int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过PARTITION_ID获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromPartitionID(partitionID int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromMinFirstTime 通过MIN_FIRST_TIME获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromMinFirstTime(minFirstTime time.Time) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`MIN_FIRST_TIME` = ?", minFirstTime).Find(&results).Error

	return
}

// GetBatchFromMinFirstTime 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromMinFirstTime(minFirstTimes []time.Time) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`MIN_FIRST_TIME` IN (?)", minFirstTimes).Find(&results).Error

	return
}

// GetFromMaxNextTime 通过MAX_NEXT_TIME获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromMaxNextTime(maxNextTime time.Time) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`MAX_NEXT_TIME` = ?", maxNextTime).Find(&results).Error

	return
}

// GetBatchFromMaxNextTime 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromMaxNextTime(maxNextTimes []time.Time) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`MAX_NEXT_TIME` IN (?)", maxNextTimes).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupArchivelogProgressMgr) GetFromStatus(status string) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupArchivelogProgressMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupArchivelogProgress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupArchivelogProgress{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
