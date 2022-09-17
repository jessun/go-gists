package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantMgr struct {
	*_BaseMgr
}

// AllTenantMgr open func
func AllTenantMgr(db *gorm.DB) *_AllTenantMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantMgr) GetTableName() string {
	return "__all_tenant"
}

// Reset 重置gorm会话
func (obj *_AllTenantMgr) Reset() *_AllTenantMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantMgr) Get() (result AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantMgr) Gets() (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取
func (obj *_AllTenantMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithReplicaNum replica_num获取
func (obj *_AllTenantMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithZoneList zone_list获取
func (obj *_AllTenantMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllTenantMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithLocked locked获取
func (obj *_AllTenantMgr) WithLocked(locked int64) Option {
	return optionFunc(func(o *options) { o.query["locked"] = locked })
}

// WithCollationType collation_type获取
func (obj *_AllTenantMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithInfo info获取
func (obj *_AllTenantMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithReadOnly read_only获取
func (obj *_AllTenantMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithRewriteMergeVersion rewrite_merge_version获取
func (obj *_AllTenantMgr) WithRewriteMergeVersion(rewriteMergeVersion int64) Option {
	return optionFunc(func(o *options) { o.query["rewrite_merge_version"] = rewriteMergeVersion })
}

// WithLocality locality获取
func (obj *_AllTenantMgr) WithLocality(locality string) Option {
	return optionFunc(func(o *options) { o.query["locality"] = locality })
}

// WithLogonlyReplicaNum logonly_replica_num获取
func (obj *_AllTenantMgr) WithLogonlyReplicaNum(logonlyReplicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["logonly_replica_num"] = logonlyReplicaNum })
}

// WithPreviousLocality previous_locality获取
func (obj *_AllTenantMgr) WithPreviousLocality(previousLocality string) Option {
	return optionFunc(func(o *options) { o.query["previous_locality"] = previousLocality })
}

// WithStorageFormatVersion storage_format_version获取
func (obj *_AllTenantMgr) WithStorageFormatVersion(storageFormatVersion int64) Option {
	return optionFunc(func(o *options) { o.query["storage_format_version"] = storageFormatVersion })
}

// WithStorageFormatWorkVersion storage_format_work_version获取
func (obj *_AllTenantMgr) WithStorageFormatWorkVersion(storageFormatWorkVersion int64) Option {
	return optionFunc(func(o *options) { o.query["storage_format_work_version"] = storageFormatWorkVersion })
}

// WithDefaultTablegroupID default_tablegroup_id获取
func (obj *_AllTenantMgr) WithDefaultTablegroupID(defaultTablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["default_tablegroup_id"] = defaultTablegroupID })
}

// WithCompatibilityMode compatibility_mode获取
func (obj *_AllTenantMgr) WithCompatibilityMode(compatibilityMode int64) Option {
	return optionFunc(func(o *options) { o.query["compatibility_mode"] = compatibilityMode })
}

// WithDropTenantTime drop_tenant_time获取
func (obj *_AllTenantMgr) WithDropTenantTime(dropTenantTime int64) Option {
	return optionFunc(func(o *options) { o.query["drop_tenant_time"] = dropTenantTime })
}

// WithStatus status获取
func (obj *_AllTenantMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithInRecyclebin in_recyclebin获取
func (obj *_AllTenantMgr) WithInRecyclebin(inRecyclebin int64) Option {
	return optionFunc(func(o *options) { o.query["in_recyclebin"] = inRecyclebin })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantMgr) GetByOption(opts ...Option) (result AllTenant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantMgr) GetByOptions(opts ...Option) (results []*AllTenant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenant, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where(options.query)
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
func (obj *_AllTenantMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantMgr) GetFromTenantID(tenantID int64) (result AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllTenantMgr) GetFromTenantName(tenantName string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllTenantMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllTenantMgr) GetFromReplicaNum(replicaNum int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllTenantMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllTenantMgr) GetFromZoneList(zoneList string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllTenantMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllTenantMgr) GetFromPrimaryZone(primaryZone string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllTenantMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromLocked 通过locked获取内容
func (obj *_AllTenantMgr) GetFromLocked(locked int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`locked` = ?", locked).Find(&results).Error

	return
}

// GetBatchFromLocked 批量查找
func (obj *_AllTenantMgr) GetBatchFromLocked(lockeds []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`locked` IN (?)", lockeds).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllTenantMgr) GetFromCollationType(collationType int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllTenantMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllTenantMgr) GetFromInfo(info string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllTenantMgr) GetBatchFromInfo(infos []string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_AllTenantMgr) GetFromReadOnly(readOnly int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_AllTenantMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

// GetFromRewriteMergeVersion 通过rewrite_merge_version获取内容
func (obj *_AllTenantMgr) GetFromRewriteMergeVersion(rewriteMergeVersion int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`rewrite_merge_version` = ?", rewriteMergeVersion).Find(&results).Error

	return
}

// GetBatchFromRewriteMergeVersion 批量查找
func (obj *_AllTenantMgr) GetBatchFromRewriteMergeVersion(rewriteMergeVersions []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`rewrite_merge_version` IN (?)", rewriteMergeVersions).Find(&results).Error

	return
}

// GetFromLocality 通过locality获取内容
func (obj *_AllTenantMgr) GetFromLocality(locality string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`locality` = ?", locality).Find(&results).Error

	return
}

// GetBatchFromLocality 批量查找
func (obj *_AllTenantMgr) GetBatchFromLocality(localitys []string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`locality` IN (?)", localitys).Find(&results).Error

	return
}

// GetFromLogonlyReplicaNum 通过logonly_replica_num获取内容
func (obj *_AllTenantMgr) GetFromLogonlyReplicaNum(logonlyReplicaNum int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`logonly_replica_num` = ?", logonlyReplicaNum).Find(&results).Error

	return
}

// GetBatchFromLogonlyReplicaNum 批量查找
func (obj *_AllTenantMgr) GetBatchFromLogonlyReplicaNum(logonlyReplicaNums []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`logonly_replica_num` IN (?)", logonlyReplicaNums).Find(&results).Error

	return
}

// GetFromPreviousLocality 通过previous_locality获取内容
func (obj *_AllTenantMgr) GetFromPreviousLocality(previousLocality string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`previous_locality` = ?", previousLocality).Find(&results).Error

	return
}

// GetBatchFromPreviousLocality 批量查找
func (obj *_AllTenantMgr) GetBatchFromPreviousLocality(previousLocalitys []string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`previous_locality` IN (?)", previousLocalitys).Find(&results).Error

	return
}

// GetFromStorageFormatVersion 通过storage_format_version获取内容
func (obj *_AllTenantMgr) GetFromStorageFormatVersion(storageFormatVersion int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`storage_format_version` = ?", storageFormatVersion).Find(&results).Error

	return
}

// GetBatchFromStorageFormatVersion 批量查找
func (obj *_AllTenantMgr) GetBatchFromStorageFormatVersion(storageFormatVersions []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`storage_format_version` IN (?)", storageFormatVersions).Find(&results).Error

	return
}

// GetFromStorageFormatWorkVersion 通过storage_format_work_version获取内容
func (obj *_AllTenantMgr) GetFromStorageFormatWorkVersion(storageFormatWorkVersion int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`storage_format_work_version` = ?", storageFormatWorkVersion).Find(&results).Error

	return
}

// GetBatchFromStorageFormatWorkVersion 批量查找
func (obj *_AllTenantMgr) GetBatchFromStorageFormatWorkVersion(storageFormatWorkVersions []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`storage_format_work_version` IN (?)", storageFormatWorkVersions).Find(&results).Error

	return
}

// GetFromDefaultTablegroupID 通过default_tablegroup_id获取内容
func (obj *_AllTenantMgr) GetFromDefaultTablegroupID(defaultTablegroupID int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`default_tablegroup_id` = ?", defaultTablegroupID).Find(&results).Error

	return
}

// GetBatchFromDefaultTablegroupID 批量查找
func (obj *_AllTenantMgr) GetBatchFromDefaultTablegroupID(defaultTablegroupIDs []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`default_tablegroup_id` IN (?)", defaultTablegroupIDs).Find(&results).Error

	return
}

// GetFromCompatibilityMode 通过compatibility_mode获取内容
func (obj *_AllTenantMgr) GetFromCompatibilityMode(compatibilityMode int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`compatibility_mode` = ?", compatibilityMode).Find(&results).Error

	return
}

// GetBatchFromCompatibilityMode 批量查找
func (obj *_AllTenantMgr) GetBatchFromCompatibilityMode(compatibilityModes []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`compatibility_mode` IN (?)", compatibilityModes).Find(&results).Error

	return
}

// GetFromDropTenantTime 通过drop_tenant_time获取内容
func (obj *_AllTenantMgr) GetFromDropTenantTime(dropTenantTime int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`drop_tenant_time` = ?", dropTenantTime).Find(&results).Error

	return
}

// GetBatchFromDropTenantTime 批量查找
func (obj *_AllTenantMgr) GetBatchFromDropTenantTime(dropTenantTimes []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`drop_tenant_time` IN (?)", dropTenantTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantMgr) GetFromStatus(status string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantMgr) GetBatchFromStatus(statuss []string) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromInRecyclebin 通过in_recyclebin获取内容
func (obj *_AllTenantMgr) GetFromInRecyclebin(inRecyclebin int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`in_recyclebin` = ?", inRecyclebin).Find(&results).Error

	return
}

// GetBatchFromInRecyclebin 批量查找
func (obj *_AllTenantMgr) GetBatchFromInRecyclebin(inRecyclebins []int64) (results []*AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`in_recyclebin` IN (?)", inRecyclebins).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantMgr) FetchByPrimaryKey(tenantID int64) (result AllTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenant{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}
