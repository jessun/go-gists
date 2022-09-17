package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualServerBackupLogArchiveStatusMgr struct {
	*_BaseMgr
}

// AllVirtualServerBackupLogArchiveStatusMgr open func
func AllVirtualServerBackupLogArchiveStatusMgr(db *gorm.DB) *_AllVirtualServerBackupLogArchiveStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualServerBackupLogArchiveStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualServerBackupLogArchiveStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_server_backup_log_archive_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetTableName() string {
	return "__all_virtual_server_backup_log_archive_status"
}

// Reset 重置gorm会话
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) Reset() *_AllVirtualServerBackupLogArchiveStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) Get() (result AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) Gets() (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithLogArchiveRound log_archive_round获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithLogArchiveRound(logArchiveRound int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_round"] = logArchiveRound })
}

// WithLogArchiveStartTs log_archive_start_ts获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithLogArchiveStartTs(logArchiveStartTs int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_start_ts"] = logArchiveStartTs })
}

// WithLogArchiveCurTs log_archive_cur_ts获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithLogArchiveCurTs(logArchiveCurTs int64) Option {
	return optionFunc(func(o *options) { o.query["log_archive_cur_ts"] = logArchiveCurTs })
}

// WithPgCount pg_count获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithPgCount(pgCount int64) Option {
	return optionFunc(func(o *options) { o.query["pg_count"] = pgCount })
}

// WithCurPieceID cur_piece_id获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) WithCurPieceID(curPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["cur_piece_id"] = curPieceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetByOption(opts ...Option) (result AllVirtualServerBackupLogArchiveStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetByOptions(opts ...Option) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualServerBackupLogArchiveStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where(options.query)
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
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromIncarnation(incarnation int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromLogArchiveRound 通过log_archive_round获取内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromLogArchiveRound(logArchiveRound int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`log_archive_round` = ?", logArchiveRound).Find(&results).Error

	return
}

// GetBatchFromLogArchiveRound 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromLogArchiveRound(logArchiveRounds []int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`log_archive_round` IN (?)", logArchiveRounds).Find(&results).Error

	return
}

// GetFromLogArchiveStartTs 通过log_archive_start_ts获取内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromLogArchiveStartTs(logArchiveStartTs int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`log_archive_start_ts` = ?", logArchiveStartTs).Find(&results).Error

	return
}

// GetBatchFromLogArchiveStartTs 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromLogArchiveStartTs(logArchiveStartTss []int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`log_archive_start_ts` IN (?)", logArchiveStartTss).Find(&results).Error

	return
}

// GetFromLogArchiveCurTs 通过log_archive_cur_ts获取内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromLogArchiveCurTs(logArchiveCurTs int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`log_archive_cur_ts` = ?", logArchiveCurTs).Find(&results).Error

	return
}

// GetBatchFromLogArchiveCurTs 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromLogArchiveCurTs(logArchiveCurTss []int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`log_archive_cur_ts` IN (?)", logArchiveCurTss).Find(&results).Error

	return
}

// GetFromPgCount 通过pg_count获取内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromPgCount(pgCount int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`pg_count` = ?", pgCount).Find(&results).Error

	return
}

// GetBatchFromPgCount 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromPgCount(pgCounts []int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`pg_count` IN (?)", pgCounts).Find(&results).Error

	return
}

// GetFromCurPieceID 通过cur_piece_id获取内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetFromCurPieceID(curPieceID int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`cur_piece_id` = ?", curPieceID).Find(&results).Error

	return
}

// GetBatchFromCurPieceID 批量查找
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) GetBatchFromCurPieceID(curPieceIDs []int64) (results []*AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`cur_piece_id` IN (?)", curPieceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualServerBackupLogArchiveStatusMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64) (result AllVirtualServerBackupLogArchiveStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerBackupLogArchiveStatus{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ?", svrIP, svrPort, tenantID).First(&result).Error

	return
}
