package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualRebalanceMapStatMgr struct {
	*_BaseMgr
}

// AllVirtualRebalanceMapStatMgr open func
func AllVirtualRebalanceMapStatMgr(db *gorm.DB) *_AllVirtualRebalanceMapStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRebalanceMapStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRebalanceMapStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rebalance_map_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRebalanceMapStatMgr) GetTableName() string {
	return "__all_virtual_rebalance_map_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRebalanceMapStatMgr) Reset() *_AllVirtualRebalanceMapStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRebalanceMapStatMgr) Get() (result AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRebalanceMapStatMgr) Gets() (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRebalanceMapStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRebalanceMapStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithMapType map_type获取
func (obj *_AllVirtualRebalanceMapStatMgr) WithMapType(mapType int64) Option {
	return optionFunc(func(o *options) { o.query["map_type"] = mapType })
}

// WithIsValid is_valid获取
func (obj *_AllVirtualRebalanceMapStatMgr) WithIsValid(isValid int8) Option {
	return optionFunc(func(o *options) { o.query["is_valid"] = isValid })
}

// WithRowSize row_size获取
func (obj *_AllVirtualRebalanceMapStatMgr) WithRowSize(rowSize int64) Option {
	return optionFunc(func(o *options) { o.query["row_size"] = rowSize })
}

// WithColSize col_size获取
func (obj *_AllVirtualRebalanceMapStatMgr) WithColSize(colSize int64) Option {
	return optionFunc(func(o *options) { o.query["col_size"] = colSize })
}

// WithTables tables获取
func (obj *_AllVirtualRebalanceMapStatMgr) WithTables(tables string) Option {
	return optionFunc(func(o *options) { o.query["tables"] = tables })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRebalanceMapStatMgr) GetByOption(opts ...Option) (result AllVirtualRebalanceMapStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRebalanceMapStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRebalanceMapStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRebalanceMapStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRebalanceMapStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where(options.query)
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
func (obj *_AllVirtualRebalanceMapStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRebalanceMapStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromMapType 通过map_type获取内容
func (obj *_AllVirtualRebalanceMapStatMgr) GetFromMapType(mapType int64) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`map_type` = ?", mapType).Find(&results).Error

	return
}

// GetBatchFromMapType 批量查找
func (obj *_AllVirtualRebalanceMapStatMgr) GetBatchFromMapType(mapTypes []int64) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`map_type` IN (?)", mapTypes).Find(&results).Error

	return
}

// GetFromIsValid 通过is_valid获取内容
func (obj *_AllVirtualRebalanceMapStatMgr) GetFromIsValid(isValid int8) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`is_valid` = ?", isValid).Find(&results).Error

	return
}

// GetBatchFromIsValid 批量查找
func (obj *_AllVirtualRebalanceMapStatMgr) GetBatchFromIsValid(isValids []int8) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`is_valid` IN (?)", isValids).Find(&results).Error

	return
}

// GetFromRowSize 通过row_size获取内容
func (obj *_AllVirtualRebalanceMapStatMgr) GetFromRowSize(rowSize int64) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`row_size` = ?", rowSize).Find(&results).Error

	return
}

// GetBatchFromRowSize 批量查找
func (obj *_AllVirtualRebalanceMapStatMgr) GetBatchFromRowSize(rowSizes []int64) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`row_size` IN (?)", rowSizes).Find(&results).Error

	return
}

// GetFromColSize 通过col_size获取内容
func (obj *_AllVirtualRebalanceMapStatMgr) GetFromColSize(colSize int64) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`col_size` = ?", colSize).Find(&results).Error

	return
}

// GetBatchFromColSize 批量查找
func (obj *_AllVirtualRebalanceMapStatMgr) GetBatchFromColSize(colSizes []int64) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`col_size` IN (?)", colSizes).Find(&results).Error

	return
}

// GetFromTables 通过tables获取内容
func (obj *_AllVirtualRebalanceMapStatMgr) GetFromTables(tables string) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`tables` = ?", tables).Find(&results).Error

	return
}

// GetBatchFromTables 批量查找
func (obj *_AllVirtualRebalanceMapStatMgr) GetBatchFromTables(tabless []string) (results []*AllVirtualRebalanceMapStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapStat{}).Where("`tables` IN (?)", tabless).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
