package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualTenantStatusMgr struct {
	*_BaseMgr
}

// TenantVirtualTenantStatusMgr open func
func TenantVirtualTenantStatusMgr(db *gorm.DB) *_TenantVirtualTenantStatusMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualTenantStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualTenantStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_tenant_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualTenantStatusMgr) GetTableName() string {
	return "__tenant_virtual_tenant_status"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualTenantStatusMgr) Reset() *_TenantVirtualTenantStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualTenantStatusMgr) Get() (result TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualTenantStatusMgr) Gets() (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualTenantStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenant tenant获取
func (obj *_TenantVirtualTenantStatusMgr) WithTenant(tenant string) Option {
	return optionFunc(func(o *options) { o.query["tenant"] = tenant })
}

// WithSvrIP svr_ip获取
func (obj *_TenantVirtualTenantStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_TenantVirtualTenantStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithReadOnly read_only获取
func (obj *_TenantVirtualTenantStatusMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualTenantStatusMgr) GetByOption(opts ...Option) (result TenantVirtualTenantStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualTenantStatusMgr) GetByOptions(opts ...Option) (results []*TenantVirtualTenantStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualTenantStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualTenantStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where(options.query)
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

// GetFromTenant 通过tenant获取内容
func (obj *_TenantVirtualTenantStatusMgr) GetFromTenant(tenant string) (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where("`tenant` = ?", tenant).Find(&results).Error

	return
}

// GetBatchFromTenant 批量查找
func (obj *_TenantVirtualTenantStatusMgr) GetBatchFromTenant(tenants []string) (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where("`tenant` IN (?)", tenants).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_TenantVirtualTenantStatusMgr) GetFromSvrIP(svrIP string) (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_TenantVirtualTenantStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_TenantVirtualTenantStatusMgr) GetFromSvrPort(svrPort int64) (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_TenantVirtualTenantStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_TenantVirtualTenantStatusMgr) GetFromReadOnly(readOnly int64) (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_TenantVirtualTenantStatusMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*TenantVirtualTenantStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualTenantStatus{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
