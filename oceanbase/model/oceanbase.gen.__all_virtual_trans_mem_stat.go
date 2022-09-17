package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTransMemStatMgr struct {
	*_BaseMgr
}

// AllVirtualTransMemStatMgr open func
func AllVirtualTransMemStatMgr(db *gorm.DB) *_AllVirtualTransMemStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTransMemStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTransMemStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trans_mem_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTransMemStatMgr) GetTableName() string {
	return "__all_virtual_trans_mem_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTransMemStatMgr) Reset() *_AllVirtualTransMemStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTransMemStatMgr) Get() (result AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTransMemStatMgr) Gets() (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTransMemStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTransMemStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTransMemStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithType type获取
func (obj *_AllVirtualTransMemStatMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithAllocCount alloc_count获取
func (obj *_AllVirtualTransMemStatMgr) WithAllocCount(allocCount int64) Option {
	return optionFunc(func(o *options) { o.query["alloc_count"] = allocCount })
}

// WithReleaseCount release_count获取
func (obj *_AllVirtualTransMemStatMgr) WithReleaseCount(releaseCount int64) Option {
	return optionFunc(func(o *options) { o.query["release_count"] = releaseCount })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTransMemStatMgr) GetByOption(opts ...Option) (result AllVirtualTransMemStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTransMemStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTransMemStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTransMemStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTransMemStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where(options.query)
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
func (obj *_AllVirtualTransMemStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTransMemStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTransMemStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTransMemStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualTransMemStatMgr) GetFromType(_type string) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualTransMemStatMgr) GetBatchFromType(_types []string) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromAllocCount 通过alloc_count获取内容
func (obj *_AllVirtualTransMemStatMgr) GetFromAllocCount(allocCount int64) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`alloc_count` = ?", allocCount).Find(&results).Error

	return
}

// GetBatchFromAllocCount 批量查找
func (obj *_AllVirtualTransMemStatMgr) GetBatchFromAllocCount(allocCounts []int64) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`alloc_count` IN (?)", allocCounts).Find(&results).Error

	return
}

// GetFromReleaseCount 通过release_count获取内容
func (obj *_AllVirtualTransMemStatMgr) GetFromReleaseCount(releaseCount int64) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`release_count` = ?", releaseCount).Find(&results).Error

	return
}

// GetBatchFromReleaseCount 批量查找
func (obj *_AllVirtualTransMemStatMgr) GetBatchFromReleaseCount(releaseCounts []int64) (results []*AllVirtualTransMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMemStat{}).Where("`release_count` IN (?)", releaseCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
