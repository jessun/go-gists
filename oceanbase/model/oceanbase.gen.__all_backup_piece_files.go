package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllBackupPieceFilesMgr struct {
	*_BaseMgr
}

// AllBackupPieceFilesMgr open func
func AllBackupPieceFilesMgr(db *gorm.DB) *_AllBackupPieceFilesMgr {
	if db == nil {
		panic(fmt.Errorf("AllBackupPieceFilesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllBackupPieceFilesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_backup_piece_files"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllBackupPieceFilesMgr) GetTableName() string {
	return "__all_backup_piece_files"
}

// Reset 重置gorm会话
func (obj *_AllBackupPieceFilesMgr) Reset() *_AllBackupPieceFilesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllBackupPieceFilesMgr) Get() (result AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllBackupPieceFilesMgr) Gets() (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllBackupPieceFilesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllBackupPieceFilesMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllBackupPieceFilesMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIncarnation incarnation获取
func (obj *_AllBackupPieceFilesMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["incarnation"] = incarnation })
}

// WithTenantID tenant_id获取
func (obj *_AllBackupPieceFilesMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRoundID round_id获取
func (obj *_AllBackupPieceFilesMgr) WithRoundID(roundID int64) Option {
	return optionFunc(func(o *options) { o.query["round_id"] = roundID })
}

// WithBackupPieceID backup_piece_id获取
func (obj *_AllBackupPieceFilesMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["backup_piece_id"] = backupPieceID })
}

// WithCopyID copy_id获取
func (obj *_AllBackupPieceFilesMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["copy_id"] = copyID })
}

// WithCreateDate create_date获取
func (obj *_AllBackupPieceFilesMgr) WithCreateDate(createDate int64) Option {
	return optionFunc(func(o *options) { o.query["create_date"] = createDate })
}

// WithStartTs start_ts获取
func (obj *_AllBackupPieceFilesMgr) WithStartTs(startTs int64) Option {
	return optionFunc(func(o *options) { o.query["start_ts"] = startTs })
}

// WithCheckpointTs checkpoint_ts获取
func (obj *_AllBackupPieceFilesMgr) WithCheckpointTs(checkpointTs int64) Option {
	return optionFunc(func(o *options) { o.query["checkpoint_ts"] = checkpointTs })
}

// WithMaxTs max_ts获取
func (obj *_AllBackupPieceFilesMgr) WithMaxTs(maxTs int64) Option {
	return optionFunc(func(o *options) { o.query["max_ts"] = maxTs })
}

// WithStatus status获取
func (obj *_AllBackupPieceFilesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithFileStatus file_status获取
func (obj *_AllBackupPieceFilesMgr) WithFileStatus(fileStatus string) Option {
	return optionFunc(func(o *options) { o.query["file_status"] = fileStatus })
}

// WithBackupDest backup_dest获取
func (obj *_AllBackupPieceFilesMgr) WithBackupDest(backupDest string) Option {
	return optionFunc(func(o *options) { o.query["backup_dest"] = backupDest })
}

// WithCompatible compatible获取
func (obj *_AllBackupPieceFilesMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithStartPieceID start_piece_id获取
func (obj *_AllBackupPieceFilesMgr) WithStartPieceID(startPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["start_piece_id"] = startPieceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllBackupPieceFilesMgr) GetByOption(opts ...Option) (result AllBackupPieceFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllBackupPieceFilesMgr) GetByOptions(opts ...Option) (results []*AllBackupPieceFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllBackupPieceFilesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllBackupPieceFiles, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where(options.query)
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
func (obj *_AllBackupPieceFilesMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIncarnation 通过incarnation获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromIncarnation(incarnation int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`incarnation` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromIncarnation(incarnations []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`incarnation` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromTenantID(tenantID int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoundID 通过round_id获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromRoundID(roundID int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`round_id` = ?", roundID).Find(&results).Error

	return
}

// GetBatchFromRoundID 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromRoundID(roundIDs []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`round_id` IN (?)", roundIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过backup_piece_id获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromBackupPieceID(backupPieceID int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`backup_piece_id` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`backup_piece_id` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过copy_id获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromCopyID(copyID int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`copy_id` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromCopyID(copyIDs []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`copy_id` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromCreateDate 通过create_date获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromCreateDate(createDate int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`create_date` = ?", createDate).Find(&results).Error

	return
}

// GetBatchFromCreateDate 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromCreateDate(createDates []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`create_date` IN (?)", createDates).Find(&results).Error

	return
}

// GetFromStartTs 通过start_ts获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromStartTs(startTs int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`start_ts` = ?", startTs).Find(&results).Error

	return
}

// GetBatchFromStartTs 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromStartTs(startTss []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`start_ts` IN (?)", startTss).Find(&results).Error

	return
}

// GetFromCheckpointTs 通过checkpoint_ts获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromCheckpointTs(checkpointTs int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`checkpoint_ts` = ?", checkpointTs).Find(&results).Error

	return
}

// GetBatchFromCheckpointTs 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromCheckpointTs(checkpointTss []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`checkpoint_ts` IN (?)", checkpointTss).Find(&results).Error

	return
}

// GetFromMaxTs 通过max_ts获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromMaxTs(maxTs int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`max_ts` = ?", maxTs).Find(&results).Error

	return
}

// GetBatchFromMaxTs 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromMaxTs(maxTss []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`max_ts` IN (?)", maxTss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromStatus(status string) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromStatus(statuss []string) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromFileStatus 通过file_status获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromFileStatus(fileStatus string) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`file_status` = ?", fileStatus).Find(&results).Error

	return
}

// GetBatchFromFileStatus 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromFileStatus(fileStatuss []string) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`file_status` IN (?)", fileStatuss).Find(&results).Error

	return
}

// GetFromBackupDest 通过backup_dest获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromBackupDest(backupDest string) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`backup_dest` = ?", backupDest).Find(&results).Error

	return
}

// GetBatchFromBackupDest 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromBackupDest(backupDests []string) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`backup_dest` IN (?)", backupDests).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromCompatible(compatible int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromStartPieceID 通过start_piece_id获取内容
func (obj *_AllBackupPieceFilesMgr) GetFromStartPieceID(startPieceID int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`start_piece_id` = ?", startPieceID).Find(&results).Error

	return
}

// GetBatchFromStartPieceID 批量查找
func (obj *_AllBackupPieceFilesMgr) GetBatchFromStartPieceID(startPieceIDs []int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`start_piece_id` IN (?)", startPieceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllBackupPieceFilesMgr) FetchByPrimaryKey(incarnation int64, tenantID int64, roundID int64, backupPieceID int64, copyID int64) (result AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`incarnation` = ? AND `tenant_id` = ? AND `round_id` = ? AND `backup_piece_id` = ? AND `copy_id` = ?", incarnation, tenantID, roundID, backupPieceID, copyID).First(&result).Error

	return
}

// FetchIndexByIDxDataTableID  获取多个内容
func (obj *_AllBackupPieceFilesMgr) FetchIndexByIDxDataTableID(incarnation int64, backupPieceID int64, copyID int64) (results []*AllBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllBackupPieceFiles{}).Where("`incarnation` = ? AND `backup_piece_id` = ? AND `copy_id` = ?", incarnation, backupPieceID, copyID).Find(&results).Error

	return
}
