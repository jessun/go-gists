package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTenantStatMgr struct {
	*_BaseMgr
}

// AllVirtualTenantStatMgr open func
func AllVirtualTenantStatMgr(db *gorm.DB) *_AllVirtualTenantStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantStatMgr) GetTableName() string {
	return "__all_virtual_tenant_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantStatMgr) Reset() *_AllVirtualTenantStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantStatMgr) Get() (result AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantStatMgr) Gets() (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableCount table_count获取
func (obj *_AllVirtualTenantStatMgr) WithTableCount(tableCount int64) Option {
	return optionFunc(func(o *options) { o.query["table_count"] = tableCount })
}

// WithRowCount row_count获取
func (obj *_AllVirtualTenantStatMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithTotalSize total_size获取
func (obj *_AllVirtualTenantStatMgr) WithTotalSize(totalSize int64) Option {
	return optionFunc(func(o *options) { o.query["total_size"] = totalSize })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantStatMgr) GetByOption(opts ...Option) (result AllVirtualTenantStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where(options.query)
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
func (obj *_AllVirtualTenantStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableCount 通过table_count获取内容
func (obj *_AllVirtualTenantStatMgr) GetFromTableCount(tableCount int64) (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where("`table_count` = ?", tableCount).Find(&results).Error

	return
}

// GetBatchFromTableCount 批量查找
func (obj *_AllVirtualTenantStatMgr) GetBatchFromTableCount(tableCounts []int64) (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where("`table_count` IN (?)", tableCounts).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllVirtualTenantStatMgr) GetFromRowCount(rowCount int64) (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllVirtualTenantStatMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromTotalSize 通过total_size获取内容
func (obj *_AllVirtualTenantStatMgr) GetFromTotalSize(totalSize int64) (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where("`total_size` = ?", totalSize).Find(&results).Error

	return
}

// GetBatchFromTotalSize 批量查找
func (obj *_AllVirtualTenantStatMgr) GetBatchFromTotalSize(totalSizes []int64) (results []*AllVirtualTenantStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantStat{}).Where("`total_size` IN (?)", totalSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
