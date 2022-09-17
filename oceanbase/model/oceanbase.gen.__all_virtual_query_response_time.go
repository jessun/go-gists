package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualQueryResponseTimeMgr struct {
	*_BaseMgr
}

// AllVirtualQueryResponseTimeMgr open func
func AllVirtualQueryResponseTimeMgr(db *gorm.DB) *_AllVirtualQueryResponseTimeMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualQueryResponseTimeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualQueryResponseTimeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_query_response_time"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualQueryResponseTimeMgr) GetTableName() string {
	return "__all_virtual_query_response_time"
}

// Reset 重置gorm会话
func (obj *_AllVirtualQueryResponseTimeMgr) Reset() *_AllVirtualQueryResponseTimeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualQueryResponseTimeMgr) Get() (result AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualQueryResponseTimeMgr) Gets() (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualQueryResponseTimeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualQueryResponseTimeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualQueryResponseTimeMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualQueryResponseTimeMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithResponseTime response_time获取
func (obj *_AllVirtualQueryResponseTimeMgr) WithResponseTime(responseTime int64) Option {
	return optionFunc(func(o *options) { o.query["response_time"] = responseTime })
}

// WithCount count获取
func (obj *_AllVirtualQueryResponseTimeMgr) WithCount(count int64) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithTotal total获取
func (obj *_AllVirtualQueryResponseTimeMgr) WithTotal(total int64) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualQueryResponseTimeMgr) GetByOption(opts ...Option) (result AllVirtualQueryResponseTime, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualQueryResponseTimeMgr) GetByOptions(opts ...Option) (results []*AllVirtualQueryResponseTime, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualQueryResponseTimeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualQueryResponseTime, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where(options.query)
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
func (obj *_AllVirtualQueryResponseTimeMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualQueryResponseTimeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualQueryResponseTimeMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualQueryResponseTimeMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualQueryResponseTimeMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualQueryResponseTimeMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromResponseTime 通过response_time获取内容
func (obj *_AllVirtualQueryResponseTimeMgr) GetFromResponseTime(responseTime int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`response_time` = ?", responseTime).Find(&results).Error

	return
}

// GetBatchFromResponseTime 批量查找
func (obj *_AllVirtualQueryResponseTimeMgr) GetBatchFromResponseTime(responseTimes []int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`response_time` IN (?)", responseTimes).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容
func (obj *_AllVirtualQueryResponseTimeMgr) GetFromCount(count int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找
func (obj *_AllVirtualQueryResponseTimeMgr) GetBatchFromCount(counts []int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容
func (obj *_AllVirtualQueryResponseTimeMgr) GetFromTotal(total int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`total` = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量查找
func (obj *_AllVirtualQueryResponseTimeMgr) GetBatchFromTotal(totals []int64) (results []*AllVirtualQueryResponseTime, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualQueryResponseTime{}).Where("`total` IN (?)", totals).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
