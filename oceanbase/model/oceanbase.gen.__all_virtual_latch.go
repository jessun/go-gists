package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualLatchMgr struct {
	*_BaseMgr
}

// AllVirtualLatchMgr open func
func AllVirtualLatchMgr(db *gorm.DB) *_AllVirtualLatchMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualLatchMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualLatchMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_latch"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualLatchMgr) GetTableName() string {
	return "__all_virtual_latch"
}

// Reset 重置gorm会话
func (obj *_AllVirtualLatchMgr) Reset() *_AllVirtualLatchMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualLatchMgr) Get() (result AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualLatchMgr) Gets() (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualLatchMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualLatchMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualLatchMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualLatchMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithLatchID latch_id获取
func (obj *_AllVirtualLatchMgr) WithLatchID(latchID int64) Option {
	return optionFunc(func(o *options) { o.query["latch_id"] = latchID })
}

// WithName name获取
func (obj *_AllVirtualLatchMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAddr addr获取
func (obj *_AllVirtualLatchMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithLevel level获取
func (obj *_AllVirtualLatchMgr) WithLevel(level int64) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithHash hash获取
func (obj *_AllVirtualLatchMgr) WithHash(hash int64) Option {
	return optionFunc(func(o *options) { o.query["hash"] = hash })
}

// WithGets gets获取
func (obj *_AllVirtualLatchMgr) WithGets(gets int64) Option {
	return optionFunc(func(o *options) { o.query["gets"] = gets })
}

// WithMisses misses获取
func (obj *_AllVirtualLatchMgr) WithMisses(misses int64) Option {
	return optionFunc(func(o *options) { o.query["misses"] = misses })
}

// WithSleeps sleeps获取
func (obj *_AllVirtualLatchMgr) WithSleeps(sleeps int64) Option {
	return optionFunc(func(o *options) { o.query["sleeps"] = sleeps })
}

// WithImmediateGets immediate_gets获取
func (obj *_AllVirtualLatchMgr) WithImmediateGets(immediateGets int64) Option {
	return optionFunc(func(o *options) { o.query["immediate_gets"] = immediateGets })
}

// WithImmediateMisses immediate_misses获取
func (obj *_AllVirtualLatchMgr) WithImmediateMisses(immediateMisses int64) Option {
	return optionFunc(func(o *options) { o.query["immediate_misses"] = immediateMisses })
}

// WithSpinGets spin_gets获取
func (obj *_AllVirtualLatchMgr) WithSpinGets(spinGets int64) Option {
	return optionFunc(func(o *options) { o.query["spin_gets"] = spinGets })
}

// WithWaitTime wait_time获取
func (obj *_AllVirtualLatchMgr) WithWaitTime(waitTime int64) Option {
	return optionFunc(func(o *options) { o.query["wait_time"] = waitTime })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualLatchMgr) GetByOption(opts ...Option) (result AllVirtualLatch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualLatchMgr) GetByOptions(opts ...Option) (results []*AllVirtualLatch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualLatchMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualLatch, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where(options.query)
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
func (obj *_AllVirtualLatchMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualLatchMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualLatchMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromLatchID 通过latch_id获取内容
func (obj *_AllVirtualLatchMgr) GetFromLatchID(latchID int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`latch_id` = ?", latchID).Find(&results).Error

	return
}

// GetBatchFromLatchID 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromLatchID(latchIDs []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`latch_id` IN (?)", latchIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualLatchMgr) GetFromName(name string) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromName(names []string) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容
func (obj *_AllVirtualLatchMgr) GetFromAddr(addr string) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`addr` = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromAddr(addrs []string) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`addr` IN (?)", addrs).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容
func (obj *_AllVirtualLatchMgr) GetFromLevel(level int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromLevel(levels []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromHash 通过hash获取内容
func (obj *_AllVirtualLatchMgr) GetFromHash(hash int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`hash` = ?", hash).Find(&results).Error

	return
}

// GetBatchFromHash 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromHash(hashs []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`hash` IN (?)", hashs).Find(&results).Error

	return
}

// GetFromGets 通过gets获取内容
func (obj *_AllVirtualLatchMgr) GetFromGets(gets int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`gets` = ?", gets).Find(&results).Error

	return
}

// GetBatchFromGets 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromGets(getss []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`gets` IN (?)", getss).Find(&results).Error

	return
}

// GetFromMisses 通过misses获取内容
func (obj *_AllVirtualLatchMgr) GetFromMisses(misses int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`misses` = ?", misses).Find(&results).Error

	return
}

// GetBatchFromMisses 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromMisses(missess []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`misses` IN (?)", missess).Find(&results).Error

	return
}

// GetFromSleeps 通过sleeps获取内容
func (obj *_AllVirtualLatchMgr) GetFromSleeps(sleeps int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`sleeps` = ?", sleeps).Find(&results).Error

	return
}

// GetBatchFromSleeps 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromSleeps(sleepss []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`sleeps` IN (?)", sleepss).Find(&results).Error

	return
}

// GetFromImmediateGets 通过immediate_gets获取内容
func (obj *_AllVirtualLatchMgr) GetFromImmediateGets(immediateGets int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`immediate_gets` = ?", immediateGets).Find(&results).Error

	return
}

// GetBatchFromImmediateGets 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromImmediateGets(immediateGetss []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`immediate_gets` IN (?)", immediateGetss).Find(&results).Error

	return
}

// GetFromImmediateMisses 通过immediate_misses获取内容
func (obj *_AllVirtualLatchMgr) GetFromImmediateMisses(immediateMisses int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`immediate_misses` = ?", immediateMisses).Find(&results).Error

	return
}

// GetBatchFromImmediateMisses 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromImmediateMisses(immediateMissess []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`immediate_misses` IN (?)", immediateMissess).Find(&results).Error

	return
}

// GetFromSpinGets 通过spin_gets获取内容
func (obj *_AllVirtualLatchMgr) GetFromSpinGets(spinGets int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`spin_gets` = ?", spinGets).Find(&results).Error

	return
}

// GetBatchFromSpinGets 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromSpinGets(spinGetss []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`spin_gets` IN (?)", spinGetss).Find(&results).Error

	return
}

// GetFromWaitTime 通过wait_time获取内容
func (obj *_AllVirtualLatchMgr) GetFromWaitTime(waitTime int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`wait_time` = ?", waitTime).Find(&results).Error

	return
}

// GetBatchFromWaitTime 批量查找
func (obj *_AllVirtualLatchMgr) GetBatchFromWaitTime(waitTimes []int64) (results []*AllVirtualLatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLatch{}).Where("`wait_time` IN (?)", waitTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
