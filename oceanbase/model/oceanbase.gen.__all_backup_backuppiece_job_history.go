package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllBackupBackuppieceJobHistoryMgr struct {
	*_BaseMgr
}

// AllBackupBackuppieceJobHistoryMgr open func
func AllBackupBackuppieceJobHistoryMgr(db *gorm.DB) *_AllBackupBackuppieceJobHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupBackuppieceJobHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupBackuppieceJobHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_backuppiece_job_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetTableName() string {
	return "__all_backup_backuppiece_job_history"
}

// Reset 重置gorm会话
func (obj *_AllBackupBackuppieceJobHistoryMgr) Reset() *_AllBackupBackuppieceJobHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) Get() (result AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupBackuppieceJobHistoryMgr) Gets() (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupBackuppieceJobHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithBackupPieceID backup_piece_id获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_piece_id"] = backupPieceID })
}

// WithMaxBackupTimes max_backup_times获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithMaxBackupTimes(maxBackupTimes int64) Option {
	return optionFunc(func(o *options) { o.query["max_backup_times"] = maxBackupTimes })
}

// WithResult result获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithResult(result int64) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithStatus status获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithComment comment获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithType type获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetByOption(opts ...Option) (result AllBackupBackuppieceJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetByOptions(opts ...Option) (results []*AllBackupBackuppieceJobHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupBackuppieceJobHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupBackuppieceJobHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromJobID(jobID int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过backup_piece_id获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromBackupPieceID(backupPieceID int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`backup_piece_id` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`backup_piece_id` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromMaxBackupTimes 通过max_backup_times获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromMaxBackupTimes(maxBackupTimes int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`max_backup_times` = ?", maxBackupTimes).Find(&results).Error

	return
}

// GetBatchFromMaxBackupTimes 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromMaxBackupTimes(maxBackupTimess []int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`max_backup_times` IN (?)", maxBackupTimess).Find(&results).Error

	return
}

// GetFromResult 通过result获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromResult(result int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`result` = ?", result).Find(&results).Error

	return
}

// GetBatchFromResult 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromResult(results []int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`result` IN (?)", results).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromStatus(status string) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromBackupDest(backupDest string) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromComment(comment string) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromComment(comments []string) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetFromType(_type int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllBackupBackuppieceJobHistoryMgr) GetBatchFromType(_types []int64) (results []*AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupBackuppieceJobHistoryMgr) FetchByPrimaryKey(jobID int64, tenantID int64, incarnation int64, backupPieceID int64) (result AllBackupBackuppieceJobHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupBackuppieceJobHistory{}).Where("`job_id` = ? AND `tenant_id` = ? AND `incarnation` = ? AND `backup_piece_id` = ?", jobID, tenantID, incarnation, backupPieceID).First(&result).Error

	return
}
