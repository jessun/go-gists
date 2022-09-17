package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllResMgrDirectiveMgr struct {
	*_BaseMgr
}

// AllResMgrDirectiveMgr open func
func AllResMgrDirectiveMgr(db *gorm.DB) *_AllResMgrDirectiveMgr {
	if db == nil {
		panic(fmt.Errorf("AllResMgrDirectiveMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllResMgrDirectiveMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_res_mgr_directive"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllResMgrDirectiveMgr) GetTableName() string {
	return "__all_res_mgr_directive"
}

// Reset 重置gorm会话
func (obj *_AllResMgrDirectiveMgr) Reset() *_AllResMgrDirectiveMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllResMgrDirectiveMgr) Get() (result AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllResMgrDirectiveMgr) Gets() (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllResMgrDirectiveMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllResMgrDirectiveMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllResMgrDirectiveMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllResMgrDirectiveMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPlan plan获取
func (obj *_AllResMgrDirectiveMgr) WithPlan(plan string) Option {
	return optionFunc(func(o *options) { o.query["plan"] = plan })
}

// WithGroupOrSubplan group_or_subplan获取
func (obj *_AllResMgrDirectiveMgr) WithGroupOrSubplan(groupOrSubplan string) Option {
	return optionFunc(func(o *options) { o.query["group_or_subplan"] = groupOrSubplan })
}

// WithComments comments获取
func (obj *_AllResMgrDirectiveMgr) WithComments(comments string) Option {
	return optionFunc(func(o *options) { o.query["comments"] = comments })
}

// WithMgmtP1 mgmt_p1获取
func (obj *_AllResMgrDirectiveMgr) WithMgmtP1(mgmtP1 int64) Option {
	return optionFunc(func(o *options) { o.query["mgmt_p1"] = mgmtP1 })
}

// WithUtilizationLimit utilization_limit获取
func (obj *_AllResMgrDirectiveMgr) WithUtilizationLimit(utilizationLimit int64) Option {
	return optionFunc(func(o *options) { o.query["utilization_limit"] = utilizationLimit })
}

// GetByOption 功能选项模式获取
func (obj *_AllResMgrDirectiveMgr) GetByOption(opts ...Option) (result AllResMgrDirective, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllResMgrDirectiveMgr) GetByOptions(opts ...Option) (results []*AllResMgrDirective, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllResMgrDirectiveMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllResMgrDirective, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where(options.query)
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
func (obj *_AllResMgrDirectiveMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllResMgrDirectiveMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllResMgrDirectiveMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllResMgrDirectiveMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllResMgrDirectiveMgr) GetFromTenantID(tenantID int64) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllResMgrDirectiveMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPlan 通过plan获取内容
func (obj *_AllResMgrDirectiveMgr) GetFromPlan(plan string) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`plan` = ?", plan).Find(&results).Error

	return
}

// GetBatchFromPlan 批量查找
func (obj *_AllResMgrDirectiveMgr) GetBatchFromPlan(plans []string) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`plan` IN (?)", plans).Find(&results).Error

	return
}

// GetFromGroupOrSubplan 通过group_or_subplan获取内容
func (obj *_AllResMgrDirectiveMgr) GetFromGroupOrSubplan(groupOrSubplan string) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`group_or_subplan` = ?", groupOrSubplan).Find(&results).Error

	return
}

// GetBatchFromGroupOrSubplan 批量查找
func (obj *_AllResMgrDirectiveMgr) GetBatchFromGroupOrSubplan(groupOrSubplans []string) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`group_or_subplan` IN (?)", groupOrSubplans).Find(&results).Error

	return
}

// GetFromComments 通过comments获取内容
func (obj *_AllResMgrDirectiveMgr) GetFromComments(comments string) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`comments` = ?", comments).Find(&results).Error

	return
}

// GetBatchFromComments 批量查找
func (obj *_AllResMgrDirectiveMgr) GetBatchFromComments(commentss []string) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`comments` IN (?)", commentss).Find(&results).Error

	return
}

// GetFromMgmtP1 通过mgmt_p1获取内容
func (obj *_AllResMgrDirectiveMgr) GetFromMgmtP1(mgmtP1 int64) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`mgmt_p1` = ?", mgmtP1).Find(&results).Error

	return
}

// GetBatchFromMgmtP1 批量查找
func (obj *_AllResMgrDirectiveMgr) GetBatchFromMgmtP1(mgmtP1s []int64) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`mgmt_p1` IN (?)", mgmtP1s).Find(&results).Error

	return
}

// GetFromUtilizationLimit 通过utilization_limit获取内容
func (obj *_AllResMgrDirectiveMgr) GetFromUtilizationLimit(utilizationLimit int64) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`utilization_limit` = ?", utilizationLimit).Find(&results).Error

	return
}

// GetBatchFromUtilizationLimit 批量查找
func (obj *_AllResMgrDirectiveMgr) GetBatchFromUtilizationLimit(utilizationLimits []int64) (results []*AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`utilization_limit` IN (?)", utilizationLimits).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllResMgrDirectiveMgr) FetchByPrimaryKey(tenantID int64, plan string, groupOrSubplan string) (result AllResMgrDirective, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrDirective{}).Where("`tenant_id` = ? AND `plan` = ? AND `group_or_subplan` = ?", tenantID, plan, groupOrSubplan).First(&result).Error

	return
}
