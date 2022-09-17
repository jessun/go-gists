package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllResMgrConsumerGroupMgr struct {
	*_BaseMgr
}

// AllResMgrConsumerGroupMgr open func
func AllResMgrConsumerGroupMgr(db *gorm.DB) *_AllResMgrConsumerGroupMgr {
	if db == nil {
		panic(fmt.Errorf("AllResMgrConsumerGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllResMgrConsumerGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_res_mgr_consumer_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllResMgrConsumerGroupMgr) GetTableName() string {
	return "__all_res_mgr_consumer_group"
}

// Reset 重置gorm会话
func (obj *_AllResMgrConsumerGroupMgr) Reset() *_AllResMgrConsumerGroupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllResMgrConsumerGroupMgr) Get() (result AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllResMgrConsumerGroupMgr) Gets() (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllResMgrConsumerGroupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllResMgrConsumerGroupMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllResMgrConsumerGroupMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllResMgrConsumerGroupMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithConsumerGroup consumer_group获取
func (obj *_AllResMgrConsumerGroupMgr) WithConsumerGroup(consumerGroup string) Option {
	return optionFunc(func(o *options) { o.query["consumer_group"] = consumerGroup })
}

// WithConsumerGroupID consumer_group_id获取
func (obj *_AllResMgrConsumerGroupMgr) WithConsumerGroupID(consumerGroupID int64) Option {
	return optionFunc(func(o *options) { o.query["consumer_group_id"] = consumerGroupID })
}

// WithComments comments获取
func (obj *_AllResMgrConsumerGroupMgr) WithComments(comments string) Option {
	return optionFunc(func(o *options) { o.query["comments"] = comments })
}

// GetByOption 功能选项模式获取
func (obj *_AllResMgrConsumerGroupMgr) GetByOption(opts ...Option) (result AllResMgrConsumerGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllResMgrConsumerGroupMgr) GetByOptions(opts ...Option) (results []*AllResMgrConsumerGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllResMgrConsumerGroupMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllResMgrConsumerGroup, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where(options.query)
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
func (obj *_AllResMgrConsumerGroupMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllResMgrConsumerGroupMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllResMgrConsumerGroupMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllResMgrConsumerGroupMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllResMgrConsumerGroupMgr) GetFromTenantID(tenantID int64) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllResMgrConsumerGroupMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromConsumerGroup 通过consumer_group获取内容
func (obj *_AllResMgrConsumerGroupMgr) GetFromConsumerGroup(consumerGroup string) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`consumer_group` = ?", consumerGroup).Find(&results).Error

	return
}

// GetBatchFromConsumerGroup 批量查找
func (obj *_AllResMgrConsumerGroupMgr) GetBatchFromConsumerGroup(consumerGroups []string) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`consumer_group` IN (?)", consumerGroups).Find(&results).Error

	return
}

// GetFromConsumerGroupID 通过consumer_group_id获取内容
func (obj *_AllResMgrConsumerGroupMgr) GetFromConsumerGroupID(consumerGroupID int64) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`consumer_group_id` = ?", consumerGroupID).Find(&results).Error

	return
}

// GetBatchFromConsumerGroupID 批量查找
func (obj *_AllResMgrConsumerGroupMgr) GetBatchFromConsumerGroupID(consumerGroupIDs []int64) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`consumer_group_id` IN (?)", consumerGroupIDs).Find(&results).Error

	return
}

// GetFromComments 通过comments获取内容
func (obj *_AllResMgrConsumerGroupMgr) GetFromComments(comments string) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`comments` = ?", comments).Find(&results).Error

	return
}

// GetBatchFromComments 批量查找
func (obj *_AllResMgrConsumerGroupMgr) GetBatchFromComments(commentss []string) (results []*AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`comments` IN (?)", commentss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllResMgrConsumerGroupMgr) FetchByPrimaryKey(tenantID int64, consumerGroup string) (result AllResMgrConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrConsumerGroup{}).Where("`tenant_id` = ? AND `consumer_group` = ?", tenantID, consumerGroup).First(&result).Error

	return
}
