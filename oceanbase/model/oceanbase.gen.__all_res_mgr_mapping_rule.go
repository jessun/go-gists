package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllResMgrMappingRuleMgr struct {
	*_BaseMgr
}

// AllResMgrMappingRuleMgr open func
func AllResMgrMappingRuleMgr(db *gorm.DB) *_AllResMgrMappingRuleMgr {
	if db == nil {
		panic(fmt.Errorf("AllResMgrMappingRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllResMgrMappingRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_res_mgr_mapping_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllResMgrMappingRuleMgr) GetTableName() string {
	return "__all_res_mgr_mapping_rule"
}

// Reset 重置gorm会话
func (obj *_AllResMgrMappingRuleMgr) Reset() *_AllResMgrMappingRuleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllResMgrMappingRuleMgr) Get() (result AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllResMgrMappingRuleMgr) Gets() (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllResMgrMappingRuleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllResMgrMappingRuleMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllResMgrMappingRuleMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllResMgrMappingRuleMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithAttribute attribute获取
func (obj *_AllResMgrMappingRuleMgr) WithAttribute(attribute string) Option {
	return optionFunc(func(o *options) { o.query["attribute"] = attribute })
}

// WithValue value获取
func (obj *_AllResMgrMappingRuleMgr) WithValue(value []byte) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithConsumerGroup consumer_group获取
func (obj *_AllResMgrMappingRuleMgr) WithConsumerGroup(consumerGroup string) Option {
	return optionFunc(func(o *options) { o.query["consumer_group"] = consumerGroup })
}

// WithStatus status获取
func (obj *_AllResMgrMappingRuleMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AllResMgrMappingRuleMgr) GetByOption(opts ...Option) (result AllResMgrMappingRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllResMgrMappingRuleMgr) GetByOptions(opts ...Option) (results []*AllResMgrMappingRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllResMgrMappingRuleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllResMgrMappingRule, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where(options.query)
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
func (obj *_AllResMgrMappingRuleMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllResMgrMappingRuleMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllResMgrMappingRuleMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllResMgrMappingRuleMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllResMgrMappingRuleMgr) GetFromTenantID(tenantID int64) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllResMgrMappingRuleMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromAttribute 通过attribute获取内容
func (obj *_AllResMgrMappingRuleMgr) GetFromAttribute(attribute string) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`attribute` = ?", attribute).Find(&results).Error

	return
}

// GetBatchFromAttribute 批量查找
func (obj *_AllResMgrMappingRuleMgr) GetBatchFromAttribute(attributes []string) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`attribute` IN (?)", attributes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllResMgrMappingRuleMgr) GetFromValue(value []byte) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllResMgrMappingRuleMgr) GetBatchFromValue(values [][]byte) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromConsumerGroup 通过consumer_group获取内容
func (obj *_AllResMgrMappingRuleMgr) GetFromConsumerGroup(consumerGroup string) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`consumer_group` = ?", consumerGroup).Find(&results).Error

	return
}

// GetBatchFromConsumerGroup 批量查找
func (obj *_AllResMgrMappingRuleMgr) GetBatchFromConsumerGroup(consumerGroups []string) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`consumer_group` IN (?)", consumerGroups).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllResMgrMappingRuleMgr) GetFromStatus(status int64) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllResMgrMappingRuleMgr) GetBatchFromStatus(statuss []int64) (results []*AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllResMgrMappingRuleMgr) FetchByPrimaryKey(tenantID int64, attribute string, value []byte) (result AllResMgrMappingRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrMappingRule{}).Where("`tenant_id` = ? AND `attribute` = ? AND `value` = ?", tenantID, attribute, value).First(&result).Error

	return
}
