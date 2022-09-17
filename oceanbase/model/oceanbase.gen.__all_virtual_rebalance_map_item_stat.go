package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualRebalanceMapItemStatMgr struct {
	*_BaseMgr
}

// AllVirtualRebalanceMapItemStatMgr open func
func AllVirtualRebalanceMapItemStatMgr(db *gorm.DB) *_AllVirtualRebalanceMapItemStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRebalanceMapItemStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRebalanceMapItemStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_rebalance_map_item_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetTableName() string {
	return "__all_virtual_rebalance_map_item_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRebalanceMapItemStatMgr) Reset() *_AllVirtualRebalanceMapItemStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) Get() (result AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRebalanceMapItemStatMgr) Gets() (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRebalanceMapItemStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithTableID table_id获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithMapType map_type获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithMapType(mapType int64) Option {
	return optionFunc(func(o *options) { o.query["map_type"] = mapType })
}

// WithRowSize row_size获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithRowSize(rowSize int64) Option {
	return optionFunc(func(o *options) { o.query["row_size"] = rowSize })
}

// WithColSize col_size获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithColSize(colSize int64) Option {
	return optionFunc(func(o *options) { o.query["col_size"] = colSize })
}

// WithPartIDx part_idx获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithPartIDx(partIDx int64) Option {
	return optionFunc(func(o *options) { o.query["part_idx"] = partIDx })
}

// WithDesignatedRole designated_role获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithDesignatedRole(designatedRole int64) Option {
	return optionFunc(func(o *options) { o.query["designated_role"] = designatedRole })
}

// WithUnitID unit_id获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithDestUnitID dest_unit_id获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) WithDestUnitID(destUnitID int64) Option {
	return optionFunc(func(o *options) { o.query["dest_unit_id"] = destUnitID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetByOption(opts ...Option) (result AllVirtualRebalanceMapItemStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualRebalanceMapItemStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRebalanceMapItemStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRebalanceMapItemStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where(options.query)
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
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromZone(zone string) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromMapType 通过map_type获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromMapType(mapType int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`map_type` = ?", mapType).Find(&results).Error

	return
}

// GetBatchFromMapType 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromMapType(mapTypes []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`map_type` IN (?)", mapTypes).Find(&results).Error

	return
}

// GetFromRowSize 通过row_size获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromRowSize(rowSize int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`row_size` = ?", rowSize).Find(&results).Error

	return
}

// GetBatchFromRowSize 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromRowSize(rowSizes []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`row_size` IN (?)", rowSizes).Find(&results).Error

	return
}

// GetFromColSize 通过col_size获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromColSize(colSize int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`col_size` = ?", colSize).Find(&results).Error

	return
}

// GetBatchFromColSize 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromColSize(colSizes []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`col_size` IN (?)", colSizes).Find(&results).Error

	return
}

// GetFromPartIDx 通过part_idx获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromPartIDx(partIDx int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`part_idx` = ?", partIDx).Find(&results).Error

	return
}

// GetBatchFromPartIDx 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromPartIDx(partIDxs []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`part_idx` IN (?)", partIDxs).Find(&results).Error

	return
}

// GetFromDesignatedRole 通过designated_role获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromDesignatedRole(designatedRole int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`designated_role` = ?", designatedRole).Find(&results).Error

	return
}

// GetBatchFromDesignatedRole 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromDesignatedRole(designatedRoles []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`designated_role` IN (?)", designatedRoles).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromUnitID(unitID int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromDestUnitID 通过dest_unit_id获取内容
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetFromDestUnitID(destUnitID int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`dest_unit_id` = ?", destUnitID).Find(&results).Error

	return
}

// GetBatchFromDestUnitID 批量查找
func (obj *_AllVirtualRebalanceMapItemStatMgr) GetBatchFromDestUnitID(destUnitIDs []int64) (results []*AllVirtualRebalanceMapItemStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRebalanceMapItemStat{}).Where("`dest_unit_id` IN (?)", destUnitIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
