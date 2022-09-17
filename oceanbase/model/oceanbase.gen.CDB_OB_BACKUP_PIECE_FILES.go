package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CdbObBackupPieceFilesMgr struct {
	*_BaseMgr
}

// CdbObBackupPieceFilesMgr open func
func CdbObBackupPieceFilesMgr(db *gorm.DB) *_CdbObBackupPieceFilesMgr {
	if db == nil {
		panic(fmt.Errorf("CdbObBackupPieceFilesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdbObBackupPieceFilesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("CDB_OB_BACKUP_PIECE_FILES"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdbObBackupPieceFilesMgr) GetTableName() string {
	return "CDB_OB_BACKUP_PIECE_FILES"
}

// Reset 重置gorm会话
func (obj *_CdbObBackupPieceFilesMgr) Reset() *_CdbObBackupPieceFilesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdbObBackupPieceFilesMgr) Get() (result CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdbObBackupPieceFilesMgr) Gets() (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdbObBackupPieceFilesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIncarnation INCARNATION获取
func (obj *_CdbObBackupPieceFilesMgr) WithIncarnation(incarnation int64) Option {
	return optionFunc(func(o *options) { o.query["INCARNATION"] = incarnation })
}

// WithTenantID TENANT_ID获取
func (obj *_CdbObBackupPieceFilesMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithRoundID ROUND_ID获取
func (obj *_CdbObBackupPieceFilesMgr) WithRoundID(roundID int64) Option {
	return optionFunc(func(o *options) { o.query["ROUND_ID"] = roundID })
}

// WithBackupPieceID BACKUP_PIECE_ID获取
func (obj *_CdbObBackupPieceFilesMgr) WithBackupPieceID(backupPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["BACKUP_PIECE_ID"] = backupPieceID })
}

// WithCopyID COPY_ID获取
func (obj *_CdbObBackupPieceFilesMgr) WithCopyID(copyID int64) Option {
	return optionFunc(func(o *options) { o.query["COPY_ID"] = copyID })
}

// WithCreateDate CREATE_DATE获取
func (obj *_CdbObBackupPieceFilesMgr) WithCreateDate(createDate int64) Option {
	return optionFunc(func(o *options) { o.query["CREATE_DATE"] = createDate })
}

// WithStartTs START_TS获取
func (obj *_CdbObBackupPieceFilesMgr) WithStartTs(startTs string) Option {
	return optionFunc(func(o *options) { o.query["START_TS"] = startTs })
}

// WithCheckpointTs CHECKPOINT_TS获取
func (obj *_CdbObBackupPieceFilesMgr) WithCheckpointTs(checkpointTs string) Option {
	return optionFunc(func(o *options) { o.query["CHECKPOINT_TS"] = checkpointTs })
}

// WithMaxTs MAX_TS获取
func (obj *_CdbObBackupPieceFilesMgr) WithMaxTs(maxTs string) Option {
	return optionFunc(func(o *options) { o.query["MAX_TS"] = maxTs })
}

// WithStatus STATUS获取
func (obj *_CdbObBackupPieceFilesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["STATUS"] = status })
}

// WithFileStatus FILE_STATUS获取
func (obj *_CdbObBackupPieceFilesMgr) WithFileStatus(fileStatus string) Option {
	return optionFunc(func(o *options) { o.query["FILE_STATUS"] = fileStatus })
}

// WithCompatible COMPATIBLE获取
func (obj *_CdbObBackupPieceFilesMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["COMPATIBLE"] = compatible })
}

// WithStartPieceID START_PIECE_ID获取
func (obj *_CdbObBackupPieceFilesMgr) WithStartPieceID(startPieceID int64) Option {
	return optionFunc(func(o *options) { o.query["START_PIECE_ID"] = startPieceID })
}

// GetByOption 功能选项模式获取
func (obj *_CdbObBackupPieceFilesMgr) GetByOption(opts ...Option) (result CdbObBackupPieceFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdbObBackupPieceFilesMgr) GetByOptions(opts ...Option) (results []*CdbObBackupPieceFiles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdbObBackupPieceFilesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdbObBackupPieceFiles, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where(options.query)
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
func (obj *_CdbObBackupPieceFilesMgr) GetFromIncarnation(incarnation int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`INCARNATION` = ?", incarnation).Find(&results).Error

	return
}

// GetBatchFromIncarnation 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromIncarnation(incarnations []int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`INCARNATION` IN (?)", incarnations).Find(&results).Error

	return
}

// GetFromTenantID 通过TENANT_ID获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromTenantID(tenantID int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoundID 通过ROUND_ID获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromRoundID(roundID int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`ROUND_ID` = ?", roundID).Find(&results).Error

	return
}

// GetBatchFromRoundID 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromRoundID(roundIDs []int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`ROUND_ID` IN (?)", roundIDs).Find(&results).Error

	return
}

// GetFromBackupPieceID 通过BACKUP_PIECE_ID获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromBackupPieceID(backupPieceID int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`BACKUP_PIECE_ID` = ?", backupPieceID).Find(&results).Error

	return
}

// GetBatchFromBackupPieceID 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromBackupPieceID(backupPieceIDs []int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`BACKUP_PIECE_ID` IN (?)", backupPieceIDs).Find(&results).Error

	return
}

// GetFromCopyID 通过COPY_ID获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromCopyID(copyID int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`COPY_ID` = ?", copyID).Find(&results).Error

	return
}

// GetBatchFromCopyID 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromCopyID(copyIDs []int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`COPY_ID` IN (?)", copyIDs).Find(&results).Error

	return
}

// GetFromCreateDate 通过CREATE_DATE获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromCreateDate(createDate int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`CREATE_DATE` = ?", createDate).Find(&results).Error

	return
}

// GetBatchFromCreateDate 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromCreateDate(createDates []int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`CREATE_DATE` IN (?)", createDates).Find(&results).Error

	return
}

// GetFromStartTs 通过START_TS获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromStartTs(startTs string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`START_TS` = ?", startTs).Find(&results).Error

	return
}

// GetBatchFromStartTs 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromStartTs(startTss []string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`START_TS` IN (?)", startTss).Find(&results).Error

	return
}

// GetFromCheckpointTs 通过CHECKPOINT_TS获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromCheckpointTs(checkpointTs string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`CHECKPOINT_TS` = ?", checkpointTs).Find(&results).Error

	return
}

// GetBatchFromCheckpointTs 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromCheckpointTs(checkpointTss []string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`CHECKPOINT_TS` IN (?)", checkpointTss).Find(&results).Error

	return
}

// GetFromMaxTs 通过MAX_TS获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromMaxTs(maxTs string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`MAX_TS` = ?", maxTs).Find(&results).Error

	return
}

// GetBatchFromMaxTs 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromMaxTs(maxTss []string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`MAX_TS` IN (?)", maxTss).Find(&results).Error

	return
}

// GetFromStatus 通过STATUS获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromStatus(status string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`STATUS` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromStatus(statuss []string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`STATUS` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromFileStatus 通过FILE_STATUS获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromFileStatus(fileStatus string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`FILE_STATUS` = ?", fileStatus).Find(&results).Error

	return
}

// GetBatchFromFileStatus 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromFileStatus(fileStatuss []string) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`FILE_STATUS` IN (?)", fileStatuss).Find(&results).Error

	return
}

// GetFromCompatible 通过COMPATIBLE获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromCompatible(compatible int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`COMPATIBLE` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromCompatible(compatibles []int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`COMPATIBLE` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromStartPieceID 通过START_PIECE_ID获取内容
func (obj *_CdbObBackupPieceFilesMgr) GetFromStartPieceID(startPieceID int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`START_PIECE_ID` = ?", startPieceID).Find(&results).Error

	return
}

// GetBatchFromStartPieceID 批量查找
func (obj *_CdbObBackupPieceFilesMgr) GetBatchFromStartPieceID(startPieceIDs []int64) (results []*CdbObBackupPieceFiles, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdbObBackupPieceFiles{}).Where("`START_PIECE_ID` IN (?)", startPieceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
