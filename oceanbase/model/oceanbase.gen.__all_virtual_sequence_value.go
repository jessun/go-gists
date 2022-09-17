package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSequenceValueMgr struct {
	*_BaseMgr
}

// AllVirtualSequenceValueMgr open func
func AllVirtualSequenceValueMgr(db *gorm.DB) *_AllVirtualSequenceValueMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSequenceValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSequenceValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sequence_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSequenceValueMgr) GetTableName() string {
	return "__all_virtual_sequence_value"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSequenceValueMgr) Reset() *_AllVirtualSequenceValueMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSequenceValueMgr) Get() (result AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSequenceValueMgr) Gets() (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSequenceValueMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSequenceValueMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSequenceID sequence_id获取
func (obj *_AllVirtualSequenceValueMgr) WithSequenceID(sequenceID int64) Option {
	return optionFunc(func(o *options) { o.query["sequence_id"] = sequenceID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSequenceValueMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSequenceValueMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithNextValue next_value获取
func (obj *_AllVirtualSequenceValueMgr) WithNextValue(nextValue float64) Option {
	return optionFunc(func(o *options) { o.query["next_value"] = nextValue })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSequenceValueMgr) GetByOption(opts ...Option) (result AllVirtualSequenceValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSequenceValueMgr) GetByOptions(opts ...Option) (results []*AllVirtualSequenceValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSequenceValueMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSequenceValue, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where(options.query)
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
func (obj *_AllVirtualSequenceValueMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSequenceValueMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSequenceID 通过sequence_id获取内容
func (obj *_AllVirtualSequenceValueMgr) GetFromSequenceID(sequenceID int64) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`sequence_id` = ?", sequenceID).Find(&results).Error

	return
}

// GetBatchFromSequenceID 批量查找
func (obj *_AllVirtualSequenceValueMgr) GetBatchFromSequenceID(sequenceIDs []int64) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`sequence_id` IN (?)", sequenceIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSequenceValueMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSequenceValueMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSequenceValueMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSequenceValueMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromNextValue 通过next_value获取内容
func (obj *_AllVirtualSequenceValueMgr) GetFromNextValue(nextValue float64) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`next_value` = ?", nextValue).Find(&results).Error

	return
}

// GetBatchFromNextValue 批量查找
func (obj *_AllVirtualSequenceValueMgr) GetBatchFromNextValue(nextValues []float64) (results []*AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`next_value` IN (?)", nextValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSequenceValueMgr) FetchByPrimaryKey(tenantID int64, sequenceID int64) (result AllVirtualSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceValue{}).Where("`tenant_id` = ? AND `sequence_id` = ?", tenantID, sequenceID).First(&result).Error

	return
}
