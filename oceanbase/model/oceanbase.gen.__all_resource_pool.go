package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllResourcePoolMgr struct {
	*_BaseMgr
}

// AllResourcePoolMgr open func
func AllResourcePoolMgr(db *gorm.DB) *_AllResourcePoolMgr {
	if db == nil {
		panic(fmt.Errorf("AllResourcePoolMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllResourcePoolMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_resource_pool"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllResourcePoolMgr) GetTableName() string {
	return "__all_resource_pool"
}

// Reset 重置gorm会话
func (obj *_AllResourcePoolMgr) Reset() *_AllResourcePoolMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllResourcePoolMgr) Get() (result AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllResourcePoolMgr) Gets() (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllResourcePoolMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllResourcePoolMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllResourcePoolMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithResourcePoolID resource_pool_id获取
func (obj *_AllResourcePoolMgr) WithResourcePoolID(resourcePoolID int64) Option {
	return optionFunc(func(o *options) { o.query["resource_pool_id"] = resourcePoolID })
}

// WithName name获取
func (obj *_AllResourcePoolMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithUnitCount unit_count获取
func (obj *_AllResourcePoolMgr) WithUnitCount(unitCount int64) Option {
	return optionFunc(func(o *options) { o.query["unit_count"] = unitCount })
}

// WithUnitConfigID unit_config_id获取
func (obj *_AllResourcePoolMgr) WithUnitConfigID(unitConfigID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_config_id"] = unitConfigID })
}

// WithZoneList zone_list获取
func (obj *_AllResourcePoolMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithTenantID tenant_id获取
func (obj *_AllResourcePoolMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithReplicaType replica_type获取
func (obj *_AllResourcePoolMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithIsTenantSysPool is_tenant_sys_pool获取
func (obj *_AllResourcePoolMgr) WithIsTenantSysPool(isTenantSysPool int8) Option {
	return optionFunc(func(o *options) { o.query["is_tenant_sys_pool"] = isTenantSysPool })
}

// GetByOption 功能选项模式获取
func (obj *_AllResourcePoolMgr) GetByOption(opts ...Option) (result AllResourcePool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllResourcePoolMgr) GetByOptions(opts ...Option) (results []*AllResourcePool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllResourcePoolMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllResourcePool, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where(options.query)
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
func (obj *_AllResourcePoolMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllResourcePoolMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromResourcePoolID 通过resource_pool_id获取内容
func (obj *_AllResourcePoolMgr) GetFromResourcePoolID(resourcePoolID int64) (result AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`resource_pool_id` = ?", resourcePoolID).First(&result).Error

	return
}

// GetBatchFromResourcePoolID 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromResourcePoolID(resourcePoolIDs []int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`resource_pool_id` IN (?)", resourcePoolIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllResourcePoolMgr) GetFromName(name string) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromName(names []string) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromUnitCount 通过unit_count获取内容
func (obj *_AllResourcePoolMgr) GetFromUnitCount(unitCount int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`unit_count` = ?", unitCount).Find(&results).Error

	return
}

// GetBatchFromUnitCount 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromUnitCount(unitCounts []int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`unit_count` IN (?)", unitCounts).Find(&results).Error

	return
}

// GetFromUnitConfigID 通过unit_config_id获取内容
func (obj *_AllResourcePoolMgr) GetFromUnitConfigID(unitConfigID int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`unit_config_id` = ?", unitConfigID).Find(&results).Error

	return
}

// GetBatchFromUnitConfigID 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromUnitConfigID(unitConfigIDs []int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`unit_config_id` IN (?)", unitConfigIDs).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllResourcePoolMgr) GetFromZoneList(zoneList string) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllResourcePoolMgr) GetFromTenantID(tenantID int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllResourcePoolMgr) GetFromReplicaType(replicaType int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromIsTenantSysPool 通过is_tenant_sys_pool获取内容
func (obj *_AllResourcePoolMgr) GetFromIsTenantSysPool(isTenantSysPool int8) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`is_tenant_sys_pool` = ?", isTenantSysPool).Find(&results).Error

	return
}

// GetBatchFromIsTenantSysPool 批量查找
func (obj *_AllResourcePoolMgr) GetBatchFromIsTenantSysPool(isTenantSysPools []int8) (results []*AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`is_tenant_sys_pool` IN (?)", isTenantSysPools).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllResourcePoolMgr) FetchByPrimaryKey(resourcePoolID int64) (result AllResourcePool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllResourcePool{}).Where("`resource_pool_id` = ?", resourcePoolID).First(&result).Error

	return
}
