package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualDatabaseStatusMgr struct {
	*_BaseMgr
}

// TenantVirtualDatabaseStatusMgr open func
func TenantVirtualDatabaseStatusMgr(db *gorm.DB) *_TenantVirtualDatabaseStatusMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualDatabaseStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualDatabaseStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_database_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualDatabaseStatusMgr) GetTableName() string {
	return "__tenant_virtual_database_status"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualDatabaseStatusMgr) Reset() *_TenantVirtualDatabaseStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualDatabaseStatusMgr) Get() (result TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualDatabaseStatusMgr) Gets() (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualDatabaseStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDb db获取
func (obj *_TenantVirtualDatabaseStatusMgr) WithDb(db string) Option {
	return optionFunc(func(o *options) { o.query["db"] = db })
}

// WithSvrIP svr_ip获取
func (obj *_TenantVirtualDatabaseStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_TenantVirtualDatabaseStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithReadOnly read_only获取
func (obj *_TenantVirtualDatabaseStatusMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualDatabaseStatusMgr) GetByOption(opts ...Option) (result TenantVirtualDatabaseStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualDatabaseStatusMgr) GetByOptions(opts ...Option) (results []*TenantVirtualDatabaseStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualDatabaseStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualDatabaseStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where(options.query)
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

// GetFromDb 通过db获取内容
func (obj *_TenantVirtualDatabaseStatusMgr) GetFromDb(db string) (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where("`db` = ?", db).Find(&results).Error

	return
}

// GetBatchFromDb 批量查找
func (obj *_TenantVirtualDatabaseStatusMgr) GetBatchFromDb(dbs []string) (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where("`db` IN (?)", dbs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_TenantVirtualDatabaseStatusMgr) GetFromSvrIP(svrIP string) (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_TenantVirtualDatabaseStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_TenantVirtualDatabaseStatusMgr) GetFromSvrPort(svrPort int64) (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_TenantVirtualDatabaseStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_TenantVirtualDatabaseStatusMgr) GetFromReadOnly(readOnly int64) (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_TenantVirtualDatabaseStatusMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*TenantVirtualDatabaseStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualDatabaseStatus{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
