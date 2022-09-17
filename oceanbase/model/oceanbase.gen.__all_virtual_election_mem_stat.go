package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualElectionMemStatMgr struct {
	*_BaseMgr
}

// AllVirtualElectionMemStatMgr open func
func AllVirtualElectionMemStatMgr(db *gorm.DB) *_AllVirtualElectionMemStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualElectionMemStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualElectionMemStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_election_mem_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualElectionMemStatMgr) GetTableName() string {
	return "__all_virtual_election_mem_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualElectionMemStatMgr) Reset() *_AllVirtualElectionMemStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualElectionMemStatMgr) Get() (result AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualElectionMemStatMgr) Gets() (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualElectionMemStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualElectionMemStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualElectionMemStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithType type获取
func (obj *_AllVirtualElectionMemStatMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithAllocCount alloc_count获取
func (obj *_AllVirtualElectionMemStatMgr) WithAllocCount(allocCount int64) Option {
	return optionFunc(func(o *options) { o.query["alloc_count"] = allocCount })
}

// WithReleaseCount release_count获取
func (obj *_AllVirtualElectionMemStatMgr) WithReleaseCount(releaseCount int64) Option {
	return optionFunc(func(o *options) { o.query["release_count"] = releaseCount })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualElectionMemStatMgr) GetByOption(opts ...Option) (result AllVirtualElectionMemStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualElectionMemStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualElectionMemStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualElectionMemStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualElectionMemStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where(options.query)
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
func (obj *_AllVirtualElectionMemStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualElectionMemStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualElectionMemStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualElectionMemStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualElectionMemStatMgr) GetFromType(_type string) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualElectionMemStatMgr) GetBatchFromType(_types []string) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromAllocCount 通过alloc_count获取内容
func (obj *_AllVirtualElectionMemStatMgr) GetFromAllocCount(allocCount int64) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`alloc_count` = ?", allocCount).Find(&results).Error

	return
}

// GetBatchFromAllocCount 批量查找
func (obj *_AllVirtualElectionMemStatMgr) GetBatchFromAllocCount(allocCounts []int64) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`alloc_count` IN (?)", allocCounts).Find(&results).Error

	return
}

// GetFromReleaseCount 通过release_count获取内容
func (obj *_AllVirtualElectionMemStatMgr) GetFromReleaseCount(releaseCount int64) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`release_count` = ?", releaseCount).Find(&results).Error

	return
}

// GetBatchFromReleaseCount 批量查找
func (obj *_AllVirtualElectionMemStatMgr) GetBatchFromReleaseCount(releaseCounts []int64) (results []*AllVirtualElectionMemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualElectionMemStat{}).Where("`release_count` IN (?)", releaseCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
