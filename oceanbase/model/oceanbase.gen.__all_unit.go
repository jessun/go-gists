package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllUnitMgr struct {
	*_BaseMgr
}

// AllUnitMgr open func
func AllUnitMgr(db *gorm.DB) *_AllUnitMgr {
	if db == nil {
		panic(fmt.Errorf("AllUnitMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllUnitMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_unit"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllUnitMgr) GetTableName() string {
	return "__all_unit"
}

// Reset 重置gorm会话
func (obj *_AllUnitMgr) Reset() *_AllUnitMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllUnitMgr) Get() (result AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllUnitMgr) Gets() (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllUnitMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllUnitMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllUnitMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithUnitID unit_id获取
func (obj *_AllUnitMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithResourcePoolID resource_pool_id获取
func (obj *_AllUnitMgr) WithResourcePoolID(resourcePoolID int64) Option {
	return optionFunc(func(o *options) { o.query["resource_pool_id"] = resourcePoolID })
}

// WithGroupID group_id获取
func (obj *_AllUnitMgr) WithGroupID(groupID int64) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithZone zone获取
func (obj *_AllUnitMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrIP svr_ip获取
func (obj *_AllUnitMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllUnitMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithMigrateFromSvrIP migrate_from_svr_ip获取
func (obj *_AllUnitMgr) WithMigrateFromSvrIP(migrateFromSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["migrate_from_svr_ip"] = migrateFromSvrIP })
}

// WithMigrateFromSvrPort migrate_from_svr_port获取
func (obj *_AllUnitMgr) WithMigrateFromSvrPort(migrateFromSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["migrate_from_svr_port"] = migrateFromSvrPort })
}

// WithManualMigrate manual_migrate获取
func (obj *_AllUnitMgr) WithManualMigrate(manualMigrate int8) Option {
	return optionFunc(func(o *options) { o.query["manual_migrate"] = manualMigrate })
}

// WithStatus status获取
func (obj *_AllUnitMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithReplicaType replica_type获取
func (obj *_AllUnitMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// GetByOption 功能选项模式获取
func (obj *_AllUnitMgr) GetByOption(opts ...Option) (result AllUnit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllUnitMgr) GetByOptions(opts ...Option) (results []*AllUnit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllUnitMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllUnit, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where(options.query)
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
func (obj *_AllUnitMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllUnitMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllUnitMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllUnitMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllUnitMgr) GetFromUnitID(unitID int64) (result AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`unit_id` = ?", unitID).First(&result).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllUnitMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromResourcePoolID 通过resource_pool_id获取内容
func (obj *_AllUnitMgr) GetFromResourcePoolID(resourcePoolID int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`resource_pool_id` = ?", resourcePoolID).Find(&results).Error

	return
}

// GetBatchFromResourcePoolID 批量查找
func (obj *_AllUnitMgr) GetBatchFromResourcePoolID(resourcePoolIDs []int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`resource_pool_id` IN (?)", resourcePoolIDs).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容
func (obj *_AllUnitMgr) GetFromGroupID(groupID int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找
func (obj *_AllUnitMgr) GetBatchFromGroupID(groupIDs []int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllUnitMgr) GetFromZone(zone string) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllUnitMgr) GetBatchFromZone(zones []string) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllUnitMgr) GetFromSvrIP(svrIP string) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllUnitMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllUnitMgr) GetFromSvrPort(svrPort int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllUnitMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromMigrateFromSvrIP 通过migrate_from_svr_ip获取内容
func (obj *_AllUnitMgr) GetFromMigrateFromSvrIP(migrateFromSvrIP string) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`migrate_from_svr_ip` = ?", migrateFromSvrIP).Find(&results).Error

	return
}

// GetBatchFromMigrateFromSvrIP 批量查找
func (obj *_AllUnitMgr) GetBatchFromMigrateFromSvrIP(migrateFromSvrIPs []string) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`migrate_from_svr_ip` IN (?)", migrateFromSvrIPs).Find(&results).Error

	return
}

// GetFromMigrateFromSvrPort 通过migrate_from_svr_port获取内容
func (obj *_AllUnitMgr) GetFromMigrateFromSvrPort(migrateFromSvrPort int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`migrate_from_svr_port` = ?", migrateFromSvrPort).Find(&results).Error

	return
}

// GetBatchFromMigrateFromSvrPort 批量查找
func (obj *_AllUnitMgr) GetBatchFromMigrateFromSvrPort(migrateFromSvrPorts []int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`migrate_from_svr_port` IN (?)", migrateFromSvrPorts).Find(&results).Error

	return
}

// GetFromManualMigrate 通过manual_migrate获取内容
func (obj *_AllUnitMgr) GetFromManualMigrate(manualMigrate int8) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`manual_migrate` = ?", manualMigrate).Find(&results).Error

	return
}

// GetBatchFromManualMigrate 批量查找
func (obj *_AllUnitMgr) GetBatchFromManualMigrate(manualMigrates []int8) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`manual_migrate` IN (?)", manualMigrates).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllUnitMgr) GetFromStatus(status string) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllUnitMgr) GetBatchFromStatus(statuss []string) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllUnitMgr) GetFromReplicaType(replicaType int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllUnitMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllUnitMgr) FetchByPrimaryKey(unitID int64) (result AllUnit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllUnit{}).Where("`unit_id` = ?", unitID).First(&result).Error

	return
}
