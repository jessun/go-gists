package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllSequenceMgr struct {
	*_BaseMgr
}

// AllSequenceMgr open func
func AllSequenceMgr(db *gorm.DB) *_AllSequenceMgr {
	if db == nil {
		panic(fmt.Errorf("AllSequenceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSequenceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sequence"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSequenceMgr) GetTableName() string {
	return "__all_sequence"
}

// Reset 重置gorm会话
func (obj *_AllSequenceMgr) Reset() *_AllSequenceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSequenceMgr) Get() (result AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSequenceMgr) Gets() (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSequenceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSequenceMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSequenceMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSequenceMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSequenceKey sequence_key获取
func (obj *_AllSequenceMgr) WithSequenceKey(sequenceKey int64) Option {
	return optionFunc(func(o *options) { o.query["sequence_key"] = sequenceKey })
}

// WithColumnID column_id获取
func (obj *_AllSequenceMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithSequenceName sequence_name获取
func (obj *_AllSequenceMgr) WithSequenceName(sequenceName string) Option {
	return optionFunc(func(o *options) { o.query["sequence_name"] = sequenceName })
}

// WithSequenceValue sequence_value获取
func (obj *_AllSequenceMgr) WithSequenceValue(sequenceValue uint64) Option {
	return optionFunc(func(o *options) { o.query["sequence_value"] = sequenceValue })
}

// WithSyncValue sync_value获取
func (obj *_AllSequenceMgr) WithSyncValue(syncValue uint64) Option {
	return optionFunc(func(o *options) { o.query["sync_value"] = syncValue })
}

// WithMigrated migrated获取
func (obj *_AllSequenceMgr) WithMigrated(migrated int64) Option {
	return optionFunc(func(o *options) { o.query["migrated"] = migrated })
}

// GetByOption 功能选项模式获取
func (obj *_AllSequenceMgr) GetByOption(opts ...Option) (result AllSequence, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSequenceMgr) GetByOptions(opts ...Option) (results []*AllSequence, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSequenceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSequence, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where(options.query)
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
func (obj *_AllSequenceMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSequenceMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSequenceMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSequenceMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSequenceMgr) GetFromTenantID(tenantID int64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSequenceMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSequenceKey 通过sequence_key获取内容
func (obj *_AllSequenceMgr) GetFromSequenceKey(sequenceKey int64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`sequence_key` = ?", sequenceKey).Find(&results).Error

	return
}

// GetBatchFromSequenceKey 批量查找
func (obj *_AllSequenceMgr) GetBatchFromSequenceKey(sequenceKeys []int64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`sequence_key` IN (?)", sequenceKeys).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllSequenceMgr) GetFromColumnID(columnID int64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllSequenceMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromSequenceName 通过sequence_name获取内容
func (obj *_AllSequenceMgr) GetFromSequenceName(sequenceName string) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`sequence_name` = ?", sequenceName).Find(&results).Error

	return
}

// GetBatchFromSequenceName 批量查找
func (obj *_AllSequenceMgr) GetBatchFromSequenceName(sequenceNames []string) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`sequence_name` IN (?)", sequenceNames).Find(&results).Error

	return
}

// GetFromSequenceValue 通过sequence_value获取内容
func (obj *_AllSequenceMgr) GetFromSequenceValue(sequenceValue uint64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`sequence_value` = ?", sequenceValue).Find(&results).Error

	return
}

// GetBatchFromSequenceValue 批量查找
func (obj *_AllSequenceMgr) GetBatchFromSequenceValue(sequenceValues []uint64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`sequence_value` IN (?)", sequenceValues).Find(&results).Error

	return
}

// GetFromSyncValue 通过sync_value获取内容
func (obj *_AllSequenceMgr) GetFromSyncValue(syncValue uint64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`sync_value` = ?", syncValue).Find(&results).Error

	return
}

// GetBatchFromSyncValue 批量查找
func (obj *_AllSequenceMgr) GetBatchFromSyncValue(syncValues []uint64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`sync_value` IN (?)", syncValues).Find(&results).Error

	return
}

// GetFromMigrated 通过migrated获取内容
func (obj *_AllSequenceMgr) GetFromMigrated(migrated int64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`migrated` = ?", migrated).Find(&results).Error

	return
}

// GetBatchFromMigrated 批量查找
func (obj *_AllSequenceMgr) GetBatchFromMigrated(migrateds []int64) (results []*AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`migrated` IN (?)", migrateds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSequenceMgr) FetchByPrimaryKey(tenantID int64, sequenceKey int64, columnID int64) (result AllSequence, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequence{}).Where("`tenant_id` = ? AND `sequence_key` = ? AND `column_id` = ?", tenantID, sequenceKey, columnID).First(&result).Error

	return
}
