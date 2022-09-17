package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualMemLeakCheckerInfoMgr struct {
	*_BaseMgr
}

// AllVirtualMemLeakCheckerInfoMgr open func
func AllVirtualMemLeakCheckerInfoMgr(db *gorm.DB) *_AllVirtualMemLeakCheckerInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualMemLeakCheckerInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualMemLeakCheckerInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_mem_leak_checker_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetTableName() string {
	return "__all_virtual_mem_leak_checker_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualMemLeakCheckerInfoMgr) Reset() *_AllVirtualMemLeakCheckerInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) Get() (result AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualMemLeakCheckerInfoMgr) Gets() (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualMemLeakCheckerInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithModName mod_name获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) WithModName(modName string) Option {
	return optionFunc(func(o *options) { o.query["mod_name"] = modName })
}

// WithModType mod_type获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) WithModType(modType string) Option {
	return optionFunc(func(o *options) { o.query["mod_type"] = modType })
}

// WithAllocCount alloc_count获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) WithAllocCount(allocCount int64) Option {
	return optionFunc(func(o *options) { o.query["alloc_count"] = allocCount })
}

// WithAllocSize alloc_size获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) WithAllocSize(allocSize int64) Option {
	return optionFunc(func(o *options) { o.query["alloc_size"] = allocSize })
}

// WithBackTrace back_trace获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) WithBackTrace(backTrace string) Option {
	return optionFunc(func(o *options) { o.query["back_trace"] = backTrace })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetByOption(opts ...Option) (result AllVirtualMemLeakCheckerInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualMemLeakCheckerInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualMemLeakCheckerInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where(options.query)
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
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromModName 通过mod_name获取内容
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetFromModName(modName string) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`mod_name` = ?", modName).Find(&results).Error

	return
}

// GetBatchFromModName 批量查找
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetBatchFromModName(modNames []string) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`mod_name` IN (?)", modNames).Find(&results).Error

	return
}

// GetFromModType 通过mod_type获取内容
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetFromModType(modType string) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`mod_type` = ?", modType).Find(&results).Error

	return
}

// GetBatchFromModType 批量查找
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetBatchFromModType(modTypes []string) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`mod_type` IN (?)", modTypes).Find(&results).Error

	return
}

// GetFromAllocCount 通过alloc_count获取内容
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetFromAllocCount(allocCount int64) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`alloc_count` = ?", allocCount).Find(&results).Error

	return
}

// GetBatchFromAllocCount 批量查找
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetBatchFromAllocCount(allocCounts []int64) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`alloc_count` IN (?)", allocCounts).Find(&results).Error

	return
}

// GetFromAllocSize 通过alloc_size获取内容
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetFromAllocSize(allocSize int64) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`alloc_size` = ?", allocSize).Find(&results).Error

	return
}

// GetBatchFromAllocSize 批量查找
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetBatchFromAllocSize(allocSizes []int64) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`alloc_size` IN (?)", allocSizes).Find(&results).Error

	return
}

// GetFromBackTrace 通过back_trace获取内容
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetFromBackTrace(backTrace string) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`back_trace` = ?", backTrace).Find(&results).Error

	return
}

// GetBatchFromBackTrace 批量查找
func (obj *_AllVirtualMemLeakCheckerInfoMgr) GetBatchFromBackTrace(backTraces []string) (results []*AllVirtualMemLeakCheckerInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemLeakCheckerInfo{}).Where("`back_trace` IN (?)", backTraces).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
