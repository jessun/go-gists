package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualZoneStatMgr struct {
	*_BaseMgr
}

// AllVirtualZoneStatMgr open func
func AllVirtualZoneStatMgr(db *gorm.DB) *_AllVirtualZoneStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualZoneStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualZoneStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_zone_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualZoneStatMgr) GetTableName() string {
	return "__all_virtual_zone_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualZoneStatMgr) Reset() *_AllVirtualZoneStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualZoneStatMgr) Get() (result AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualZoneStatMgr) Gets() (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualZoneStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithZone zone获取
func (obj *_AllVirtualZoneStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithIsMerging is_merging获取
func (obj *_AllVirtualZoneStatMgr) WithIsMerging(isMerging int64) Option {
	return optionFunc(func(o *options) { o.query["is_merging"] = isMerging })
}

// WithStatus status获取
func (obj *_AllVirtualZoneStatMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithServerCount server_count获取
func (obj *_AllVirtualZoneStatMgr) WithServerCount(serverCount int64) Option {
	return optionFunc(func(o *options) { o.query["server_count"] = serverCount })
}

// WithResourcePoolCount resource_pool_count获取
func (obj *_AllVirtualZoneStatMgr) WithResourcePoolCount(resourcePoolCount int64) Option {
	return optionFunc(func(o *options) { o.query["resource_pool_count"] = resourcePoolCount })
}

// WithUnitCount unit_count获取
func (obj *_AllVirtualZoneStatMgr) WithUnitCount(unitCount int64) Option {
	return optionFunc(func(o *options) { o.query["unit_count"] = unitCount })
}

// WithCluster cluster获取
func (obj *_AllVirtualZoneStatMgr) WithCluster(cluster string) Option {
	return optionFunc(func(o *options) { o.query["cluster"] = cluster })
}

// WithRegion region获取
func (obj *_AllVirtualZoneStatMgr) WithRegion(region string) Option {
	return optionFunc(func(o *options) { o.query["region"] = region })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualZoneStatMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualZoneStatMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualZoneStatMgr) WithSpare3(spare3 int64) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithSpare4 spare4获取
func (obj *_AllVirtualZoneStatMgr) WithSpare4(spare4 string) Option {
	return optionFunc(func(o *options) { o.query["spare4"] = spare4 })
}

// WithSpare5 spare5获取
func (obj *_AllVirtualZoneStatMgr) WithSpare5(spare5 string) Option {
	return optionFunc(func(o *options) { o.query["spare5"] = spare5 })
}

// WithSpare6 spare6获取
func (obj *_AllVirtualZoneStatMgr) WithSpare6(spare6 string) Option {
	return optionFunc(func(o *options) { o.query["spare6"] = spare6 })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualZoneStatMgr) GetByOption(opts ...Option) (result AllVirtualZoneStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualZoneStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualZoneStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualZoneStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualZoneStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where(options.query)
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

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromZone(zone string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromIsMerging 通过is_merging获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromIsMerging(isMerging int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`is_merging` = ?", isMerging).Find(&results).Error

	return
}

// GetBatchFromIsMerging 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromIsMerging(isMergings []int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`is_merging` IN (?)", isMergings).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromStatus(status string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromServerCount 通过server_count获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromServerCount(serverCount int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`server_count` = ?", serverCount).Find(&results).Error

	return
}

// GetBatchFromServerCount 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromServerCount(serverCounts []int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`server_count` IN (?)", serverCounts).Find(&results).Error

	return
}

// GetFromResourcePoolCount 通过resource_pool_count获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromResourcePoolCount(resourcePoolCount int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`resource_pool_count` = ?", resourcePoolCount).Find(&results).Error

	return
}

// GetBatchFromResourcePoolCount 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromResourcePoolCount(resourcePoolCounts []int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`resource_pool_count` IN (?)", resourcePoolCounts).Find(&results).Error

	return
}

// GetFromUnitCount 通过unit_count获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromUnitCount(unitCount int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`unit_count` = ?", unitCount).Find(&results).Error

	return
}

// GetBatchFromUnitCount 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromUnitCount(unitCounts []int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`unit_count` IN (?)", unitCounts).Find(&results).Error

	return
}

// GetFromCluster 通过cluster获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromCluster(cluster string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`cluster` = ?", cluster).Find(&results).Error

	return
}

// GetBatchFromCluster 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromCluster(clusters []string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`cluster` IN (?)", clusters).Find(&results).Error

	return
}

// GetFromRegion 通过region获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromRegion(region string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`region` = ?", region).Find(&results).Error

	return
}

// GetBatchFromRegion 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromRegion(regions []string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`region` IN (?)", regions).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromSpare3(spare3 int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromSpare3(spare3s []int64) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromSpare4 通过spare4获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromSpare4(spare4 string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare4` = ?", spare4).Find(&results).Error

	return
}

// GetBatchFromSpare4 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromSpare4(spare4s []string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare4` IN (?)", spare4s).Find(&results).Error

	return
}

// GetFromSpare5 通过spare5获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromSpare5(spare5 string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare5` = ?", spare5).Find(&results).Error

	return
}

// GetBatchFromSpare5 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromSpare5(spare5s []string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare5` IN (?)", spare5s).Find(&results).Error

	return
}

// GetFromSpare6 通过spare6获取内容
func (obj *_AllVirtualZoneStatMgr) GetFromSpare6(spare6 string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare6` = ?", spare6).Find(&results).Error

	return
}

// GetBatchFromSpare6 批量查找
func (obj *_AllVirtualZoneStatMgr) GetBatchFromSpare6(spare6s []string) (results []*AllVirtualZoneStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualZoneStat{}).Where("`spare6` IN (?)", spare6s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
