package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllUnitConfigMgr struct {
	*_BaseMgr
}

// AllUnitConfigMgr open func
func AllUnitConfigMgr(db *gorm.DB) *_AllUnitConfigMgr {
	if db == nil {
		panic(fmt.Errorf("AllUnitConfigMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllUnitConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_unit_config"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllUnitConfigMgr) GetTableName() string {
	return "__all_unit_config"
}

// Reset 重置gorm会话
func (obj *_AllUnitConfigMgr) Reset() *_AllUnitConfigMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllUnitConfigMgr) Get() (result AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllUnitConfigMgr) Gets() (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllUnitConfigMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllUnitConfigMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllUnitConfigMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithUnitConfigID unit_config_id获取
func (obj *_AllUnitConfigMgr) WithUnitConfigID(unitConfigID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_config_id"] = unitConfigID })
}

// WithName name获取
func (obj *_AllUnitConfigMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithMaxCPU max_cpu获取
func (obj *_AllUnitConfigMgr) WithMaxCPU(maxCPU float64) Option {
	return optionFunc(func(o *options) { o.query["max_cpu"] = maxCPU })
}

// WithMinCPU min_cpu获取
func (obj *_AllUnitConfigMgr) WithMinCPU(minCPU float64) Option {
	return optionFunc(func(o *options) { o.query["min_cpu"] = minCPU })
}

// WithMaxMemory max_memory获取
func (obj *_AllUnitConfigMgr) WithMaxMemory(maxMemory int64) Option {
	return optionFunc(func(o *options) { o.query["max_memory"] = maxMemory })
}

// WithMinMemory min_memory获取
func (obj *_AllUnitConfigMgr) WithMinMemory(minMemory int64) Option {
	return optionFunc(func(o *options) { o.query["min_memory"] = minMemory })
}

// WithMaxIops max_iops获取
func (obj *_AllUnitConfigMgr) WithMaxIops(maxIops int64) Option {
	return optionFunc(func(o *options) { o.query["max_iops"] = maxIops })
}

// WithMinIops min_iops获取
func (obj *_AllUnitConfigMgr) WithMinIops(minIops int64) Option {
	return optionFunc(func(o *options) { o.query["min_iops"] = minIops })
}

// WithMaxDiskSize max_disk_size获取
func (obj *_AllUnitConfigMgr) WithMaxDiskSize(maxDiskSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_disk_size"] = maxDiskSize })
}

// WithMaxSessionNum max_session_num获取
func (obj *_AllUnitConfigMgr) WithMaxSessionNum(maxSessionNum int64) Option {
	return optionFunc(func(o *options) { o.query["max_session_num"] = maxSessionNum })
}

// GetByOption 功能选项模式获取
func (obj *_AllUnitConfigMgr) GetByOption(opts ...Option) (result AllUnitConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllUnitConfigMgr) GetByOptions(opts ...Option) (results []*AllUnitConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllUnitConfigMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllUnitConfig, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllUnitConfigMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllUnitConfigMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromUnitConfigID 通过unit_config_id获取内容
func (obj *_AllUnitConfigMgr) GetFromUnitConfigID(unitConfigID int64) (result AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`unit_config_id` = ?", unitConfigID).First(&result).Error

	return
}

// GetBatchFromUnitConfigID 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromUnitConfigID(unitConfigIDs []int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`unit_config_id` IN (?)", unitConfigIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllUnitConfigMgr) GetFromName(name string) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromName(names []string) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromMaxCPU 通过max_cpu获取内容
func (obj *_AllUnitConfigMgr) GetFromMaxCPU(maxCPU float64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_cpu` = ?", maxCPU).Find(&results).Error

	return
}

// GetBatchFromMaxCPU 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromMaxCPU(maxCPUs []float64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_cpu` IN (?)", maxCPUs).Find(&results).Error

	return
}

// GetFromMinCPU 通过min_cpu获取内容
func (obj *_AllUnitConfigMgr) GetFromMinCPU(minCPU float64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`min_cpu` = ?", minCPU).Find(&results).Error

	return
}

// GetBatchFromMinCPU 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromMinCPU(minCPUs []float64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`min_cpu` IN (?)", minCPUs).Find(&results).Error

	return
}

// GetFromMaxMemory 通过max_memory获取内容
func (obj *_AllUnitConfigMgr) GetFromMaxMemory(maxMemory int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_memory` = ?", maxMemory).Find(&results).Error

	return
}

// GetBatchFromMaxMemory 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromMaxMemory(maxMemorys []int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_memory` IN (?)", maxMemorys).Find(&results).Error

	return
}

// GetFromMinMemory 通过min_memory获取内容
func (obj *_AllUnitConfigMgr) GetFromMinMemory(minMemory int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`min_memory` = ?", minMemory).Find(&results).Error

	return
}

// GetBatchFromMinMemory 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromMinMemory(minMemorys []int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`min_memory` IN (?)", minMemorys).Find(&results).Error

	return
}

// GetFromMaxIops 通过max_iops获取内容
func (obj *_AllUnitConfigMgr) GetFromMaxIops(maxIops int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_iops` = ?", maxIops).Find(&results).Error

	return
}

// GetBatchFromMaxIops 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromMaxIops(maxIopss []int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_iops` IN (?)", maxIopss).Find(&results).Error

	return
}

// GetFromMinIops 通过min_iops获取内容
func (obj *_AllUnitConfigMgr) GetFromMinIops(minIops int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`min_iops` = ?", minIops).Find(&results).Error

	return
}

// GetBatchFromMinIops 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromMinIops(minIopss []int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`min_iops` IN (?)", minIopss).Find(&results).Error

	return
}

// GetFromMaxDiskSize 通过max_disk_size获取内容
func (obj *_AllUnitConfigMgr) GetFromMaxDiskSize(maxDiskSize int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_disk_size` = ?", maxDiskSize).Find(&results).Error

	return
}

// GetBatchFromMaxDiskSize 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromMaxDiskSize(maxDiskSizes []int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_disk_size` IN (?)", maxDiskSizes).Find(&results).Error

	return
}

// GetFromMaxSessionNum 通过max_session_num获取内容
func (obj *_AllUnitConfigMgr) GetFromMaxSessionNum(maxSessionNum int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_session_num` = ?", maxSessionNum).Find(&results).Error

	return
}

// GetBatchFromMaxSessionNum 批量查找
func (obj *_AllUnitConfigMgr) GetBatchFromMaxSessionNum(maxSessionNums []int64) (results []*AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`max_session_num` IN (?)", maxSessionNums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllUnitConfigMgr) FetchByPrimaryKey(unitConfigID int64) (result AllUnitConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnitConfig{}).Where("`unit_config_id` = ?", unitConfigID).First(&result).Error

	return
}
