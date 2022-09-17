package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualSQLWorkareaActiveMgr struct {
	*_BaseMgr
}

// AllVirtualSQLWorkareaActiveMgr open func
func AllVirtualSQLWorkareaActiveMgr(db *gorm.DB) *_AllVirtualSQLWorkareaActiveMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSQLWorkareaActiveMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSQLWorkareaActiveMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sql_workarea_active"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetTableName() string {
	return "__all_virtual_sql_workarea_active"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSQLWorkareaActiveMgr) Reset() *_AllVirtualSQLWorkareaActiveMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) Get() (result AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSQLWorkareaActiveMgr) Gets() (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSQLWorkareaActiveMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPlanID plan_id获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithPlanID(planID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_id"] = planID })
}

// WithSQLID sql_id获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithSQLExecID sql_exec_id获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithSQLExecID(sqlExecID int64) Option {
	return optionFunc(func(o *options) { o.query["sql_exec_id"] = sqlExecID })
}

// WithOperationType operation_type获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithOperationType(operationType string) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithOperationID operation_id获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithOperationID(operationID int64) Option {
	return optionFunc(func(o *options) { o.query["operation_id"] = operationID })
}

// WithSid sid获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["sid"] = sid })
}

// WithActiveTime active_time获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithActiveTime(activeTime int64) Option {
	return optionFunc(func(o *options) { o.query["active_time"] = activeTime })
}

// WithWorkAreaSize work_area_size获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithWorkAreaSize(workAreaSize int64) Option {
	return optionFunc(func(o *options) { o.query["work_area_size"] = workAreaSize })
}

// WithExpectSize expect_size获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithExpectSize(expectSize int64) Option {
	return optionFunc(func(o *options) { o.query["expect_size"] = expectSize })
}

// WithActualMemUsed actual_mem_used获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithActualMemUsed(actualMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["actual_mem_used"] = actualMemUsed })
}

// WithMaxMemUsed max_mem_used获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithMaxMemUsed(maxMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["max_mem_used"] = maxMemUsed })
}

// WithNumberPasses number_passes获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithNumberPasses(numberPasses int64) Option {
	return optionFunc(func(o *options) { o.query["number_passes"] = numberPasses })
}

// WithTempsegSize tempseg_size获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithTempsegSize(tempsegSize int64) Option {
	return optionFunc(func(o *options) { o.query["tempseg_size"] = tempsegSize })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPolicy policy获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) WithPolicy(policy string) Option {
	return optionFunc(func(o *options) { o.query["policy"] = policy })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetByOption(opts ...Option) (result AllVirtualSQLWorkareaActive, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetByOptions(opts ...Option) (results []*AllVirtualSQLWorkareaActive, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSQLWorkareaActiveMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSQLWorkareaActive, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromPlanID 通过plan_id获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromPlanID(planID int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`plan_id` = ?", planID).Find(&results).Error

	return
}

// GetBatchFromPlanID 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromPlanID(planIDs []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`plan_id` IN (?)", planIDs).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromSQLID(sqlID string) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromSQLID(sqlIDs []string) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromSQLExecID 通过sql_exec_id获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromSQLExecID(sqlExecID int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`sql_exec_id` = ?", sqlExecID).Find(&results).Error

	return
}

// GetBatchFromSQLExecID 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromSQLExecID(sqlExecIDs []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`sql_exec_id` IN (?)", sqlExecIDs).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromOperationType(operationType string) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromOperationType(operationTypes []string) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromOperationID 通过operation_id获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromOperationID(operationID int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`operation_id` = ?", operationID).Find(&results).Error

	return
}

// GetBatchFromOperationID 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromOperationID(operationIDs []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`operation_id` IN (?)", operationIDs).Find(&results).Error

	return
}

// GetFromSid 通过sid获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromSid(sid int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`sid` = ?", sid).Find(&results).Error

	return
}

// GetBatchFromSid 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromSid(sids []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`sid` IN (?)", sids).Find(&results).Error

	return
}

// GetFromActiveTime 通过active_time获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromActiveTime(activeTime int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`active_time` = ?", activeTime).Find(&results).Error

	return
}

// GetBatchFromActiveTime 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromActiveTime(activeTimes []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`active_time` IN (?)", activeTimes).Find(&results).Error

	return
}

// GetFromWorkAreaSize 通过work_area_size获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromWorkAreaSize(workAreaSize int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`work_area_size` = ?", workAreaSize).Find(&results).Error

	return
}

// GetBatchFromWorkAreaSize 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromWorkAreaSize(workAreaSizes []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`work_area_size` IN (?)", workAreaSizes).Find(&results).Error

	return
}

// GetFromExpectSize 通过expect_size获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromExpectSize(expectSize int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`expect_size` = ?", expectSize).Find(&results).Error

	return
}

// GetBatchFromExpectSize 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromExpectSize(expectSizes []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`expect_size` IN (?)", expectSizes).Find(&results).Error

	return
}

// GetFromActualMemUsed 通过actual_mem_used获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromActualMemUsed(actualMemUsed int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`actual_mem_used` = ?", actualMemUsed).Find(&results).Error

	return
}

// GetBatchFromActualMemUsed 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromActualMemUsed(actualMemUseds []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`actual_mem_used` IN (?)", actualMemUseds).Find(&results).Error

	return
}

// GetFromMaxMemUsed 通过max_mem_used获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromMaxMemUsed(maxMemUsed int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`max_mem_used` = ?", maxMemUsed).Find(&results).Error

	return
}

// GetBatchFromMaxMemUsed 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromMaxMemUsed(maxMemUseds []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`max_mem_used` IN (?)", maxMemUseds).Find(&results).Error

	return
}

// GetFromNumberPasses 通过number_passes获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromNumberPasses(numberPasses int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`number_passes` = ?", numberPasses).Find(&results).Error

	return
}

// GetBatchFromNumberPasses 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromNumberPasses(numberPassess []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`number_passes` IN (?)", numberPassess).Find(&results).Error

	return
}

// GetFromTempsegSize 通过tempseg_size获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromTempsegSize(tempsegSize int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`tempseg_size` = ?", tempsegSize).Find(&results).Error

	return
}

// GetBatchFromTempsegSize 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromTempsegSize(tempsegSizes []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`tempseg_size` IN (?)", tempsegSizes).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPolicy 通过policy获取内容
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetFromPolicy(policy string) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`policy` = ?", policy).Find(&results).Error

	return
}

// GetBatchFromPolicy 批量查找
func (obj *_AllVirtualSQLWorkareaActiveMgr) GetBatchFromPolicy(policys []string) (results []*AllVirtualSQLWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSQLWorkareaActive{}).Where("`policy` IN (?)", policys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
