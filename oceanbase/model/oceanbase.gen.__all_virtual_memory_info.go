package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualMemoryInfoMgr struct {
	*_BaseMgr
}

// AllVirtualMemoryInfoMgr open func
func AllVirtualMemoryInfoMgr(db *gorm.DB) *_AllVirtualMemoryInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualMemoryInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualMemoryInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_memory_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualMemoryInfoMgr) GetTableName() string {
	return "__all_virtual_memory_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualMemoryInfoMgr) Reset() *_AllVirtualMemoryInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualMemoryInfoMgr) Get() (result AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualMemoryInfoMgr) Gets() (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualMemoryInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualMemoryInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualMemoryInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualMemoryInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithCtxID ctx_id获取
func (obj *_AllVirtualMemoryInfoMgr) WithCtxID(ctxID int64) Option {
	return optionFunc(func(o *options) { o.query["ctx_id"] = ctxID })
}

// WithLabel label获取
func (obj *_AllVirtualMemoryInfoMgr) WithLabel(label string) Option {
	return optionFunc(func(o *options) { o.query["label"] = label })
}

// WithCtxName ctx_name获取
func (obj *_AllVirtualMemoryInfoMgr) WithCtxName(ctxName string) Option {
	return optionFunc(func(o *options) { o.query["ctx_name"] = ctxName })
}

// WithModType mod_type获取
func (obj *_AllVirtualMemoryInfoMgr) WithModType(modType string) Option {
	return optionFunc(func(o *options) { o.query["mod_type"] = modType })
}

// WithModID mod_id获取
func (obj *_AllVirtualMemoryInfoMgr) WithModID(modID int64) Option {
	return optionFunc(func(o *options) { o.query["mod_id"] = modID })
}

// WithModName mod_name获取
func (obj *_AllVirtualMemoryInfoMgr) WithModName(modName string) Option {
	return optionFunc(func(o *options) { o.query["mod_name"] = modName })
}

// WithZone zone获取
func (obj *_AllVirtualMemoryInfoMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithHold hold获取
func (obj *_AllVirtualMemoryInfoMgr) WithHold(hold int64) Option {
	return optionFunc(func(o *options) { o.query["hold"] = hold })
}

// WithUsed used获取
func (obj *_AllVirtualMemoryInfoMgr) WithUsed(used int64) Option {
	return optionFunc(func(o *options) { o.query["used"] = used })
}

// WithCount count获取
func (obj *_AllVirtualMemoryInfoMgr) WithCount(count int64) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithAllocCount alloc_count获取
func (obj *_AllVirtualMemoryInfoMgr) WithAllocCount(allocCount int64) Option {
	return optionFunc(func(o *options) { o.query["alloc_count"] = allocCount })
}

// WithFreeCount free_count获取
func (obj *_AllVirtualMemoryInfoMgr) WithFreeCount(freeCount int64) Option {
	return optionFunc(func(o *options) { o.query["free_count"] = freeCount })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualMemoryInfoMgr) GetByOption(opts ...Option) (result AllVirtualMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualMemoryInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualMemoryInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualMemoryInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where(options.query)
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
func (obj *_AllVirtualMemoryInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromCtxID 通过ctx_id获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromCtxID(ctxID int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`ctx_id` = ?", ctxID).Find(&results).Error

	return
}

// GetBatchFromCtxID 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromCtxID(ctxIDs []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`ctx_id` IN (?)", ctxIDs).Find(&results).Error

	return
}

// GetFromLabel 通过label获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromLabel(label string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`label` = ?", label).Find(&results).Error

	return
}

// GetBatchFromLabel 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromLabel(labels []string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`label` IN (?)", labels).Find(&results).Error

	return
}

// GetFromCtxName 通过ctx_name获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromCtxName(ctxName string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`ctx_name` = ?", ctxName).Find(&results).Error

	return
}

// GetBatchFromCtxName 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromCtxName(ctxNames []string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`ctx_name` IN (?)", ctxNames).Find(&results).Error

	return
}

// GetFromModType 通过mod_type获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromModType(modType string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`mod_type` = ?", modType).Find(&results).Error

	return
}

// GetBatchFromModType 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromModType(modTypes []string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`mod_type` IN (?)", modTypes).Find(&results).Error

	return
}

// GetFromModID 通过mod_id获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromModID(modID int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`mod_id` = ?", modID).Find(&results).Error

	return
}

// GetBatchFromModID 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromModID(modIDs []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`mod_id` IN (?)", modIDs).Find(&results).Error

	return
}

// GetFromModName 通过mod_name获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromModName(modName string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`mod_name` = ?", modName).Find(&results).Error

	return
}

// GetBatchFromModName 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromModName(modNames []string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`mod_name` IN (?)", modNames).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromZone(zone string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromZone(zones []string) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromHold 通过hold获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromHold(hold int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`hold` = ?", hold).Find(&results).Error

	return
}

// GetBatchFromHold 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromHold(holds []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`hold` IN (?)", holds).Find(&results).Error

	return
}

// GetFromUsed 通过used获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromUsed(used int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`used` = ?", used).Find(&results).Error

	return
}

// GetBatchFromUsed 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromUsed(useds []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`used` IN (?)", useds).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromCount(count int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromCount(counts []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

// GetFromAllocCount 通过alloc_count获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromAllocCount(allocCount int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`alloc_count` = ?", allocCount).Find(&results).Error

	return
}

// GetBatchFromAllocCount 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromAllocCount(allocCounts []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`alloc_count` IN (?)", allocCounts).Find(&results).Error

	return
}

// GetFromFreeCount 通过free_count获取内容
func (obj *_AllVirtualMemoryInfoMgr) GetFromFreeCount(freeCount int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`free_count` = ?", freeCount).Find(&results).Error

	return
}

// GetBatchFromFreeCount 批量查找
func (obj *_AllVirtualMemoryInfoMgr) GetBatchFromFreeCount(freeCounts []int64) (results []*AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`free_count` IN (?)", freeCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualMemoryInfoMgr) FetchByPrimaryKey(tenantID int64, svrIP string, svrPort int64, ctxID int64, label string) (result AllVirtualMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemoryInfo{}).Where("`tenant_id` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `ctx_id` = ? AND `label` = ?", tenantID, svrIP, svrPort, ctxID, label).First(&result).Error

	return
}
