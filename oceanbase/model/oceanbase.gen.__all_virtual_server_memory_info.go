package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualServerMemoryInfoMgr struct {
	*_BaseMgr
}

// AllVirtualServerMemoryInfoMgr open func
func AllVirtualServerMemoryInfoMgr(db *gorm.DB) *_AllVirtualServerMemoryInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualServerMemoryInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualServerMemoryInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_server_memory_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualServerMemoryInfoMgr) GetTableName() string {
	return "__all_virtual_server_memory_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualServerMemoryInfoMgr) Reset() *_AllVirtualServerMemoryInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualServerMemoryInfoMgr) Get() (result AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualServerMemoryInfoMgr) Gets() (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualServerMemoryInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithServerMemoryHold server_memory_hold获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithServerMemoryHold(serverMemoryHold int64) Option {
	return optionFunc(func(o *options) { o.query["server_memory_hold"] = serverMemoryHold })
}

// WithServerMemoryLimit server_memory_limit获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithServerMemoryLimit(serverMemoryLimit int64) Option {
	return optionFunc(func(o *options) { o.query["server_memory_limit"] = serverMemoryLimit })
}

// WithSystemReserved system_reserved获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithSystemReserved(systemReserved int64) Option {
	return optionFunc(func(o *options) { o.query["system_reserved"] = systemReserved })
}

// WithActiveMemstoreUsed active_memstore_used获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithActiveMemstoreUsed(activeMemstoreUsed int64) Option {
	return optionFunc(func(o *options) { o.query["active_memstore_used"] = activeMemstoreUsed })
}

// WithTotalMemstoreUsed total_memstore_used获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithTotalMemstoreUsed(totalMemstoreUsed int64) Option {
	return optionFunc(func(o *options) { o.query["total_memstore_used"] = totalMemstoreUsed })
}

// WithMajorFreezeTrigger major_freeze_trigger获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithMajorFreezeTrigger(majorFreezeTrigger int64) Option {
	return optionFunc(func(o *options) { o.query["major_freeze_trigger"] = majorFreezeTrigger })
}

// WithMemstoreLimit memstore_limit获取
func (obj *_AllVirtualServerMemoryInfoMgr) WithMemstoreLimit(memstoreLimit int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_limit"] = memstoreLimit })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualServerMemoryInfoMgr) GetByOption(opts ...Option) (result AllVirtualServerMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualServerMemoryInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualServerMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualServerMemoryInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualServerMemoryInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where(options.query)
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
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromServerMemoryHold 通过server_memory_hold获取内容
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromServerMemoryHold(serverMemoryHold int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`server_memory_hold` = ?", serverMemoryHold).Find(&results).Error

	return
}

// GetBatchFromServerMemoryHold 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromServerMemoryHold(serverMemoryHolds []int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`server_memory_hold` IN (?)", serverMemoryHolds).Find(&results).Error

	return
}

// GetFromServerMemoryLimit 通过server_memory_limit获取内容
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromServerMemoryLimit(serverMemoryLimit int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`server_memory_limit` = ?", serverMemoryLimit).Find(&results).Error

	return
}

// GetBatchFromServerMemoryLimit 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromServerMemoryLimit(serverMemoryLimits []int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`server_memory_limit` IN (?)", serverMemoryLimits).Find(&results).Error

	return
}

// GetFromSystemReserved 通过system_reserved获取内容
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromSystemReserved(systemReserved int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`system_reserved` = ?", systemReserved).Find(&results).Error

	return
}

// GetBatchFromSystemReserved 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromSystemReserved(systemReserveds []int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`system_reserved` IN (?)", systemReserveds).Find(&results).Error

	return
}

// GetFromActiveMemstoreUsed 通过active_memstore_used获取内容
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromActiveMemstoreUsed(activeMemstoreUsed int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`active_memstore_used` = ?", activeMemstoreUsed).Find(&results).Error

	return
}

// GetBatchFromActiveMemstoreUsed 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromActiveMemstoreUsed(activeMemstoreUseds []int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`active_memstore_used` IN (?)", activeMemstoreUseds).Find(&results).Error

	return
}

// GetFromTotalMemstoreUsed 通过total_memstore_used获取内容
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromTotalMemstoreUsed(totalMemstoreUsed int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`total_memstore_used` = ?", totalMemstoreUsed).Find(&results).Error

	return
}

// GetBatchFromTotalMemstoreUsed 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromTotalMemstoreUsed(totalMemstoreUseds []int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`total_memstore_used` IN (?)", totalMemstoreUseds).Find(&results).Error

	return
}

// GetFromMajorFreezeTrigger 通过major_freeze_trigger获取内容
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromMajorFreezeTrigger(majorFreezeTrigger int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`major_freeze_trigger` = ?", majorFreezeTrigger).Find(&results).Error

	return
}

// GetBatchFromMajorFreezeTrigger 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromMajorFreezeTrigger(majorFreezeTriggers []int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`major_freeze_trigger` IN (?)", majorFreezeTriggers).Find(&results).Error

	return
}

// GetFromMemstoreLimit 通过memstore_limit获取内容
func (obj *_AllVirtualServerMemoryInfoMgr) GetFromMemstoreLimit(memstoreLimit int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`memstore_limit` = ?", memstoreLimit).Find(&results).Error

	return
}

// GetBatchFromMemstoreLimit 批量查找
func (obj *_AllVirtualServerMemoryInfoMgr) GetBatchFromMemstoreLimit(memstoreLimits []int64) (results []*AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`memstore_limit` IN (?)", memstoreLimits).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualServerMemoryInfoMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllVirtualServerMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerMemoryInfo{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
