package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSequenceV2Mgr struct {
	*_BaseMgr
}

// AllVirtualSequenceV2Mgr open func
func AllVirtualSequenceV2Mgr(db *gorm.DB) *_AllVirtualSequenceV2Mgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSequenceV2Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSequenceV2Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sequence_v2"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSequenceV2Mgr) GetTableName() string {
	return "__all_virtual_sequence_v2"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSequenceV2Mgr) Reset() *_AllVirtualSequenceV2Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSequenceV2Mgr) Get() (result AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSequenceV2Mgr) Gets() (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSequenceV2Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSequenceV2Mgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSequenceKey sequence_key获取
func (obj *_AllVirtualSequenceV2Mgr) WithSequenceKey(sequenceKey int64) Option {
	return optionFunc(func(o *options) { o.query["sequence_key"] = sequenceKey })
}

// WithColumnID column_id获取
func (obj *_AllVirtualSequenceV2Mgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSequenceV2Mgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSequenceV2Mgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSequenceName sequence_name获取
func (obj *_AllVirtualSequenceV2Mgr) WithSequenceName(sequenceName string) Option {
	return optionFunc(func(o *options) { o.query["sequence_name"] = sequenceName })
}

// WithSequenceValue sequence_value获取
func (obj *_AllVirtualSequenceV2Mgr) WithSequenceValue(sequenceValue uint64) Option {
	return optionFunc(func(o *options) { o.query["sequence_value"] = sequenceValue })
}

// WithSyncValue sync_value获取
func (obj *_AllVirtualSequenceV2Mgr) WithSyncValue(syncValue uint64) Option {
	return optionFunc(func(o *options) { o.query["sync_value"] = syncValue })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSequenceV2Mgr) GetByOption(opts ...Option) (result AllVirtualSequenceV2, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSequenceV2Mgr) GetByOptions(opts ...Option) (results []*AllVirtualSequenceV2, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSequenceV2Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSequenceV2, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where(options.query)
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

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualSequenceV2Mgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSequenceV2Mgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSequenceKey 通过sequence_key获取内容
func (obj *_AllVirtualSequenceV2Mgr) GetFromSequenceKey(sequenceKey int64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`sequence_key` = ?", sequenceKey).Find(&results).Error

	return
}

// GetBatchFromSequenceKey 批量查找
func (obj *_AllVirtualSequenceV2Mgr) GetBatchFromSequenceKey(sequenceKeys []int64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`sequence_key` IN (?)", sequenceKeys).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualSequenceV2Mgr) GetFromColumnID(columnID int64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualSequenceV2Mgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSequenceV2Mgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSequenceV2Mgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSequenceV2Mgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSequenceV2Mgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSequenceName 通过sequence_name获取内容
func (obj *_AllVirtualSequenceV2Mgr) GetFromSequenceName(sequenceName string) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`sequence_name` = ?", sequenceName).Find(&results).Error

	return
}

// GetBatchFromSequenceName 批量查找
func (obj *_AllVirtualSequenceV2Mgr) GetBatchFromSequenceName(sequenceNames []string) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`sequence_name` IN (?)", sequenceNames).Find(&results).Error

	return
}

// GetFromSequenceValue 通过sequence_value获取内容
func (obj *_AllVirtualSequenceV2Mgr) GetFromSequenceValue(sequenceValue uint64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`sequence_value` = ?", sequenceValue).Find(&results).Error

	return
}

// GetBatchFromSequenceValue 批量查找
func (obj *_AllVirtualSequenceV2Mgr) GetBatchFromSequenceValue(sequenceValues []uint64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`sequence_value` IN (?)", sequenceValues).Find(&results).Error

	return
}

// GetFromSyncValue 通过sync_value获取内容
func (obj *_AllVirtualSequenceV2Mgr) GetFromSyncValue(syncValue uint64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`sync_value` = ?", syncValue).Find(&results).Error

	return
}

// GetBatchFromSyncValue 批量查找
func (obj *_AllVirtualSequenceV2Mgr) GetBatchFromSyncValue(syncValues []uint64) (results []*AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`sync_value` IN (?)", syncValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSequenceV2Mgr) FetchByPrimaryKey(tenantID int64, sequenceKey int64, columnID int64) (result AllVirtualSequenceV2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceV2{}).Where("`tenant_id` = ? AND `sequence_key` = ? AND `column_id` = ?", tenantID, sequenceKey, columnID).First(&result).Error

	return
}
