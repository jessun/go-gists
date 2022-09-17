package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllSequenceValueMgr struct {
	*_BaseMgr
}

// AllSequenceValueMgr open func
func AllSequenceValueMgr(db *gorm.DB) *_AllSequenceValueMgr {
	if db == nil {
		panic(fmt.Errorf("AllSequenceValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSequenceValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sequence_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSequenceValueMgr) GetTableName() string {
	return "__all_sequence_value"
}

// Reset 重置gorm会话
func (obj *_AllSequenceValueMgr) Reset() *_AllSequenceValueMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSequenceValueMgr) Get() (result AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSequenceValueMgr) Gets() (results []*AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSequenceValueMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSequenceValueMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSequenceValueMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSequenceID sequence_id获取
func (obj *_AllSequenceValueMgr) WithSequenceID(sequenceID int64) Option {
	return optionFunc(func(o *options) { o.query["sequence_id"] = sequenceID })
}

// WithNextValue next_value获取
func (obj *_AllSequenceValueMgr) WithNextValue(nextValue float64) Option {
	return optionFunc(func(o *options) { o.query["next_value"] = nextValue })
}

// GetByOption 功能选项模式获取
func (obj *_AllSequenceValueMgr) GetByOption(opts ...Option) (result AllSequenceValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSequenceValueMgr) GetByOptions(opts ...Option) (results []*AllSequenceValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSequenceValueMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSequenceValue, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where(options.query)
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
func (obj *_AllSequenceValueMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSequenceValueMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSequenceValueMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSequenceValueMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSequenceID 通过sequence_id获取内容
func (obj *_AllSequenceValueMgr) GetFromSequenceID(sequenceID int64) (result AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`sequence_id` = ?", sequenceID).First(&result).Error

	return
}

// GetBatchFromSequenceID 批量查找
func (obj *_AllSequenceValueMgr) GetBatchFromSequenceID(sequenceIDs []int64) (results []*AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`sequence_id` IN (?)", sequenceIDs).Find(&results).Error

	return
}

// GetFromNextValue 通过next_value获取内容
func (obj *_AllSequenceValueMgr) GetFromNextValue(nextValue float64) (results []*AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`next_value` = ?", nextValue).Find(&results).Error

	return
}

// GetBatchFromNextValue 批量查找
func (obj *_AllSequenceValueMgr) GetBatchFromNextValue(nextValues []float64) (results []*AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`next_value` IN (?)", nextValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSequenceValueMgr) FetchByPrimaryKey(sequenceID int64) (result AllSequenceValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceValue{}).Where("`sequence_id` = ?", sequenceID).First(&result).Error

	return
}
