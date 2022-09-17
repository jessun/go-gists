package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTenantMemstoreInfoMgr struct {
	*_BaseMgr
}

// AllVirtualTenantMemstoreInfoMgr open func
func AllVirtualTenantMemstoreInfoMgr(db *gorm.DB) *_AllVirtualTenantMemstoreInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantMemstoreInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantMemstoreInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_memstore_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetTableName() string {
	return "__all_virtual_tenant_memstore_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantMemstoreInfoMgr) Reset() *_AllVirtualTenantMemstoreInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) Get() (result AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantMemstoreInfoMgr) Gets() (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantMemstoreInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithActiveMemstoreUsed active_memstore_used获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) WithActiveMemstoreUsed(activeMemstoreUsed int64) Option {
	return optionFunc(func(o *options) { o.query["active_memstore_used"] = activeMemstoreUsed })
}

// WithTotalMemstoreUsed total_memstore_used获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) WithTotalMemstoreUsed(totalMemstoreUsed int64) Option {
	return optionFunc(func(o *options) { o.query["total_memstore_used"] = totalMemstoreUsed })
}

// WithMajorFreezeTrigger major_freeze_trigger获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) WithMajorFreezeTrigger(majorFreezeTrigger int64) Option {
	return optionFunc(func(o *options) { o.query["major_freeze_trigger"] = majorFreezeTrigger })
}

// WithMemstoreLimit memstore_limit获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) WithMemstoreLimit(memstoreLimit int64) Option {
	return optionFunc(func(o *options) { o.query["memstore_limit"] = memstoreLimit })
}

// WithFreezeCnt freeze_cnt获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) WithFreezeCnt(freezeCnt int64) Option {
	return optionFunc(func(o *options) { o.query["freeze_cnt"] = freezeCnt })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetByOption(opts ...Option) (result AllVirtualTenantMemstoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantMemstoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantMemstoreInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantMemstoreInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where(options.query)
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
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromActiveMemstoreUsed 通过active_memstore_used获取内容
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetFromActiveMemstoreUsed(activeMemstoreUsed int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`active_memstore_used` = ?", activeMemstoreUsed).Find(&results).Error

	return
}

// GetBatchFromActiveMemstoreUsed 批量查找
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetBatchFromActiveMemstoreUsed(activeMemstoreUseds []int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`active_memstore_used` IN (?)", activeMemstoreUseds).Find(&results).Error

	return
}

// GetFromTotalMemstoreUsed 通过total_memstore_used获取内容
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetFromTotalMemstoreUsed(totalMemstoreUsed int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`total_memstore_used` = ?", totalMemstoreUsed).Find(&results).Error

	return
}

// GetBatchFromTotalMemstoreUsed 批量查找
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetBatchFromTotalMemstoreUsed(totalMemstoreUseds []int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`total_memstore_used` IN (?)", totalMemstoreUseds).Find(&results).Error

	return
}

// GetFromMajorFreezeTrigger 通过major_freeze_trigger获取内容
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetFromMajorFreezeTrigger(majorFreezeTrigger int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`major_freeze_trigger` = ?", majorFreezeTrigger).Find(&results).Error

	return
}

// GetBatchFromMajorFreezeTrigger 批量查找
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetBatchFromMajorFreezeTrigger(majorFreezeTriggers []int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`major_freeze_trigger` IN (?)", majorFreezeTriggers).Find(&results).Error

	return
}

// GetFromMemstoreLimit 通过memstore_limit获取内容
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetFromMemstoreLimit(memstoreLimit int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`memstore_limit` = ?", memstoreLimit).Find(&results).Error

	return
}

// GetBatchFromMemstoreLimit 批量查找
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetBatchFromMemstoreLimit(memstoreLimits []int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`memstore_limit` IN (?)", memstoreLimits).Find(&results).Error

	return
}

// GetFromFreezeCnt 通过freeze_cnt获取内容
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetFromFreezeCnt(freezeCnt int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`freeze_cnt` = ?", freezeCnt).Find(&results).Error

	return
}

// GetBatchFromFreezeCnt 批量查找
func (obj *_AllVirtualTenantMemstoreInfoMgr) GetBatchFromFreezeCnt(freezeCnts []int64) (results []*AllVirtualTenantMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantMemstoreInfo{}).Where("`freeze_cnt` IN (?)", freezeCnts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
