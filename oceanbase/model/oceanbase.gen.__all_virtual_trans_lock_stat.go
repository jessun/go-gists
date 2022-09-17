package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTransLockStatMgr struct {
	*_BaseMgr
}

// AllVirtualTransLockStatMgr open func
func AllVirtualTransLockStatMgr(db *gorm.DB) *_AllVirtualTransLockStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTransLockStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTransLockStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trans_lock_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTransLockStatMgr) GetTableName() string {
	return "__all_virtual_trans_lock_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTransLockStatMgr) Reset() *_AllVirtualTransLockStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTransLockStatMgr) Get() (result AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTransLockStatMgr) Gets() (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTransLockStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTransLockStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTransID trans_id获取
func (obj *_AllVirtualTransLockStatMgr) WithTransID(transID string) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTransLockStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTransLockStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithPartition partition获取
func (obj *_AllVirtualTransLockStatMgr) WithPartition(partition string) Option {
	return optionFunc(func(o *options) { o.query["partition"] = partition })
}

// WithTableID table_id获取
func (obj *_AllVirtualTransLockStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithRowkey rowkey获取
func (obj *_AllVirtualTransLockStatMgr) WithRowkey(rowkey string) Option {
	return optionFunc(func(o *options) { o.query["rowkey"] = rowkey })
}

// WithSessionID session_id获取
func (obj *_AllVirtualTransLockStatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithProxyID proxy_id获取
func (obj *_AllVirtualTransLockStatMgr) WithProxyID(proxyID string) Option {
	return optionFunc(func(o *options) { o.query["proxy_id"] = proxyID })
}

// WithCtxCreateTime ctx_create_time获取
func (obj *_AllVirtualTransLockStatMgr) WithCtxCreateTime(ctxCreateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["ctx_create_time"] = ctxCreateTime })
}

// WithExpiredTime expired_time获取
func (obj *_AllVirtualTransLockStatMgr) WithExpiredTime(expiredTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expired_time"] = expiredTime })
}

// WithRowLockAddr row_lock_addr获取
func (obj *_AllVirtualTransLockStatMgr) WithRowLockAddr(rowLockAddr uint64) Option {
	return optionFunc(func(o *options) { o.query["row_lock_addr"] = rowLockAddr })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTransLockStatMgr) GetByOption(opts ...Option) (result AllVirtualTransLockStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTransLockStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTransLockStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTransLockStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTransLockStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where(options.query)
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
func (obj *_AllVirtualTransLockStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTransID 通过trans_id获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromTransID(transID string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromTransID(transIDs []string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromPartition 通过partition获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromPartition(partition string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`partition` = ?", partition).Find(&results).Error

	return
}

// GetBatchFromPartition 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromPartition(partitions []string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`partition` IN (?)", partitions).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromRowkey 通过rowkey获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromRowkey(rowkey string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`rowkey` = ?", rowkey).Find(&results).Error

	return
}

// GetBatchFromRowkey 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromRowkey(rowkeys []string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`rowkey` IN (?)", rowkeys).Find(&results).Error

	return
}

// GetFromSessionID 通过session_id获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromProxyID 通过proxy_id获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromProxyID(proxyID string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`proxy_id` = ?", proxyID).Find(&results).Error

	return
}

// GetBatchFromProxyID 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromProxyID(proxyIDs []string) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`proxy_id` IN (?)", proxyIDs).Find(&results).Error

	return
}

// GetFromCtxCreateTime 通过ctx_create_time获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromCtxCreateTime(ctxCreateTime time.Time) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`ctx_create_time` = ?", ctxCreateTime).Find(&results).Error

	return
}

// GetBatchFromCtxCreateTime 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromCtxCreateTime(ctxCreateTimes []time.Time) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`ctx_create_time` IN (?)", ctxCreateTimes).Find(&results).Error

	return
}

// GetFromExpiredTime 通过expired_time获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromExpiredTime(expiredTime time.Time) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`expired_time` = ?", expiredTime).Find(&results).Error

	return
}

// GetBatchFromExpiredTime 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromExpiredTime(expiredTimes []time.Time) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`expired_time` IN (?)", expiredTimes).Find(&results).Error

	return
}

// GetFromRowLockAddr 通过row_lock_addr获取内容
func (obj *_AllVirtualTransLockStatMgr) GetFromRowLockAddr(rowLockAddr uint64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`row_lock_addr` = ?", rowLockAddr).Find(&results).Error

	return
}

// GetBatchFromRowLockAddr 批量查找
func (obj *_AllVirtualTransLockStatMgr) GetBatchFromRowLockAddr(rowLockAddrs []uint64) (results []*AllVirtualTransLockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransLockStat{}).Where("`row_lock_addr` IN (?)", rowLockAddrs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
