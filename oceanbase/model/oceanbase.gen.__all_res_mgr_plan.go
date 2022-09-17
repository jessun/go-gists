package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllResMgrPlanMgr struct {
	*_BaseMgr
}

// AllResMgrPlanMgr open func
func AllResMgrPlanMgr(db *gorm.DB) *_AllResMgrPlanMgr {
	if db == nil {
		panic(fmt.Errorf("AllResMgrPlanMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllResMgrPlanMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_res_mgr_plan"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllResMgrPlanMgr) GetTableName() string {
	return "__all_res_mgr_plan"
}

// Reset 重置gorm会话
func (obj *_AllResMgrPlanMgr) Reset() *_AllResMgrPlanMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllResMgrPlanMgr) Get() (result AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllResMgrPlanMgr) Gets() (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllResMgrPlanMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllResMgrPlanMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllResMgrPlanMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllResMgrPlanMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPlan plan获取
func (obj *_AllResMgrPlanMgr) WithPlan(plan string) Option {
	return optionFunc(func(o *options) { o.query["plan"] = plan })
}

// WithComments comments获取
func (obj *_AllResMgrPlanMgr) WithComments(comments string) Option {
	return optionFunc(func(o *options) { o.query["comments"] = comments })
}

// GetByOption 功能选项模式获取
func (obj *_AllResMgrPlanMgr) GetByOption(opts ...Option) (result AllResMgrPlan, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllResMgrPlanMgr) GetByOptions(opts ...Option) (results []*AllResMgrPlan, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllResMgrPlanMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllResMgrPlan, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where(options.query)
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
func (obj *_AllResMgrPlanMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllResMgrPlanMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllResMgrPlanMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllResMgrPlanMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllResMgrPlanMgr) GetFromTenantID(tenantID int64) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllResMgrPlanMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPlan 通过plan获取内容
func (obj *_AllResMgrPlanMgr) GetFromPlan(plan string) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`plan` = ?", plan).Find(&results).Error

	return
}

// GetBatchFromPlan 批量查找
func (obj *_AllResMgrPlanMgr) GetBatchFromPlan(plans []string) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`plan` IN (?)", plans).Find(&results).Error

	return
}

// GetFromComments 通过comments获取内容
func (obj *_AllResMgrPlanMgr) GetFromComments(comments string) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`comments` = ?", comments).Find(&results).Error

	return
}

// GetBatchFromComments 批量查找
func (obj *_AllResMgrPlanMgr) GetBatchFromComments(commentss []string) (results []*AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`comments` IN (?)", commentss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllResMgrPlanMgr) FetchByPrimaryKey(tenantID int64, plan string) (result AllResMgrPlan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResMgrPlan{}).Where("`tenant_id` = ? AND `plan` = ?", tenantID, plan).First(&result).Error

	return
}
