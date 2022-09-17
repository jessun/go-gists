package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPlanCachePlanExplainMgr struct {
	*_BaseMgr
}

// AllVirtualPlanCachePlanExplainMgr open func
func AllVirtualPlanCachePlanExplainMgr(db *gorm.DB) *_AllVirtualPlanCachePlanExplainMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPlanCachePlanExplainMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPlanCachePlanExplainMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_plan_cache_plan_explain"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetTableName() string {
	return "__all_virtual_plan_cache_plan_explain"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPlanCachePlanExplainMgr) Reset() *_AllVirtualPlanCachePlanExplainMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) Get() (result AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPlanCachePlanExplainMgr) Gets() (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPlanCachePlanExplainMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPlanID plan_id获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_id"] = planID })
}

// WithOperator operator获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithOperator(operator string) Option {
	return optionFunc(func(o *options) { o.query["operator"] = operator })
}

// WithName name获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithRows rows获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithRows(rows int64) Option {
	return optionFunc(func(o *options) { o.query["rows"] = rows })
}

// WithCost cost获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithCost(cost int64) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithProperty property获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithProperty(property string) Option {
	return optionFunc(func(o *options) { o.query["property"] = property })
}

// WithPlanDepth plan_depth获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithPlanDepth(planDepth int64) Option {
	return optionFunc(func(o *options) { o.query["plan_depth"] = planDepth })
}

// WithPlanLineID plan_line_id获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) WithPlanLineID(planLineID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_line_id"] = planLineID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetByOption(opts ...Option) (result AllVirtualPlanCachePlanExplain, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetByOptions(opts ...Option) (results []*AllVirtualPlanCachePlanExplain, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPlanCachePlanExplainMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPlanCachePlanExplain, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where(options.query)
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
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromPlanID 通过plan_id获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromPlanID(planID int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`plan_id` = ?", planID).Find(&results).Error

	return
}

// GetBatchFromPlanID 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromPlanID(planIDs []int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`plan_id` IN (?)", planIDs).Find(&results).Error

	return
}

// GetFromOperator 通过operator获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromOperator(operator string) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`operator` = ?", operator).Find(&results).Error

	return
}

// GetBatchFromOperator 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromOperator(operators []string) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`operator` IN (?)", operators).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromName(name string) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromName(names []string) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromRows 通过rows获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromRows(rows int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`rows` = ?", rows).Find(&results).Error

	return
}

// GetBatchFromRows 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromRows(rowss []int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`rows` IN (?)", rowss).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromCost(cost int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromCost(costs []int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromProperty 通过property获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromProperty(property string) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`property` = ?", property).Find(&results).Error

	return
}

// GetBatchFromProperty 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromProperty(propertys []string) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`property` IN (?)", propertys).Find(&results).Error

	return
}

// GetFromPlanDepth 通过plan_depth获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromPlanDepth(planDepth int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`plan_depth` = ?", planDepth).Find(&results).Error

	return
}

// GetBatchFromPlanDepth 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromPlanDepth(planDepths []int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`plan_depth` IN (?)", planDepths).Find(&results).Error

	return
}

// GetFromPlanLineID 通过plan_line_id获取内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetFromPlanLineID(planLineID int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`plan_line_id` = ?", planLineID).Find(&results).Error

	return
}

// GetBatchFromPlanLineID 批量查找
func (obj *_AllVirtualPlanCachePlanExplainMgr) GetBatchFromPlanLineID(planLineIDs []int64) (results []*AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`plan_line_id` IN (?)", planLineIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPlanCachePlanExplainMgr) FetchByPrimaryKey(tenantID int64, svrIP string, svrPort int64, planID int64) (result AllVirtualPlanCachePlanExplain, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPlanCachePlanExplain{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `plan_id` = ?", tenantID, svrIP, svrPort, planID).First(&result).Error

	return
}
