package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantHistoryMgr struct {
	*_BaseMgr
}

// AllTenantHistoryMgr open func
func AllTenantHistoryMgr(db *gorm.DB) *_AllTenantHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantHistoryMgr) GetTableName() string {
	return "__all_tenant_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantHistoryMgr) Reset() *_AllTenantHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantHistoryMgr) Get() (result AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantHistoryMgr) Gets() (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithTenantName tenant_name获取
func (obj *_AllTenantHistoryMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithReplicaNum replica_num获取
func (obj *_AllTenantHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithZoneList zone_list获取
func (obj *_AllTenantHistoryMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllTenantHistoryMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithLocked locked获取
func (obj *_AllTenantHistoryMgr) WithLocked(locked int64) Option {
	return optionFunc(func(o *options) { o.query["locked"] = locked })
}

// WithCollationType collation_type获取
func (obj *_AllTenantHistoryMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithInfo info获取
func (obj *_AllTenantHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithReadOnly read_only获取
func (obj *_AllTenantHistoryMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithRewriteMergeVersion rewrite_merge_version获取
func (obj *_AllTenantHistoryMgr) WithRewriteMergeVersion(rewriteMergeVersion int64) Option {
	return optionFunc(func(o *options) { o.query["rewrite_merge_version"] = rewriteMergeVersion })
}

// WithLocality locality获取
func (obj *_AllTenantHistoryMgr) WithLocality(locality string) Option {
	return optionFunc(func(o *options) { o.query["locality"] = locality })
}

// WithLogonlyReplicaNum logonly_replica_num获取
func (obj *_AllTenantHistoryMgr) WithLogonlyReplicaNum(logonlyReplicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["logonly_replica_num"] = logonlyReplicaNum })
}

// WithPreviousLocality previous_locality获取
func (obj *_AllTenantHistoryMgr) WithPreviousLocality(previousLocality string) Option {
	return optionFunc(func(o *options) { o.query["previous_locality"] = previousLocality })
}

// WithStorageFormatVersion storage_format_version获取
func (obj *_AllTenantHistoryMgr) WithStorageFormatVersion(storageFormatVersion int64) Option {
	return optionFunc(func(o *options) { o.query["storage_format_version"] = storageFormatVersion })
}

// WithStorageFormatWorkVersion storage_format_work_version获取
func (obj *_AllTenantHistoryMgr) WithStorageFormatWorkVersion(storageFormatWorkVersion int64) Option {
	return optionFunc(func(o *options) { o.query["storage_format_work_version"] = storageFormatWorkVersion })
}

// WithDefaultTablegroupID default_tablegroup_id获取
func (obj *_AllTenantHistoryMgr) WithDefaultTablegroupID(defaultTablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["default_tablegroup_id"] = defaultTablegroupID })
}

// WithCompatibilityMode compatibility_mode获取
func (obj *_AllTenantHistoryMgr) WithCompatibilityMode(compatibilityMode int64) Option {
	return optionFunc(func(o *options) { o.query["compatibility_mode"] = compatibilityMode })
}

// WithDropTenantTime drop_tenant_time获取
func (obj *_AllTenantHistoryMgr) WithDropTenantTime(dropTenantTime int64) Option {
	return optionFunc(func(o *options) { o.query["drop_tenant_time"] = dropTenantTime })
}

// WithStatus status获取
func (obj *_AllTenantHistoryMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithInRecyclebin in_recyclebin获取
func (obj *_AllTenantHistoryMgr) WithInRecyclebin(inRecyclebin int64) Option {
	return optionFunc(func(o *options) { o.query["in_recyclebin"] = inRecyclebin })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantHistoryMgr) GetByOption(opts ...Option) (result AllTenantHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where(options.query)
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
func (obj *_AllTenantHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllTenantHistoryMgr) GetFromTenantName(tenantName string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllTenantHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllTenantHistoryMgr) GetFromZoneList(zoneList string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllTenantHistoryMgr) GetFromPrimaryZone(primaryZone string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromLocked 通过locked获取内容
func (obj *_AllTenantHistoryMgr) GetFromLocked(locked int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`locked` = ?", locked).Find(&results).Error

	return
}

// GetBatchFromLocked 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromLocked(lockeds []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`locked` IN (?)", lockeds).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllTenantHistoryMgr) GetFromCollationType(collationType int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllTenantHistoryMgr) GetFromInfo(info string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromInfo(infos []string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_AllTenantHistoryMgr) GetFromReadOnly(readOnly int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

// GetFromRewriteMergeVersion 通过rewrite_merge_version获取内容
func (obj *_AllTenantHistoryMgr) GetFromRewriteMergeVersion(rewriteMergeVersion int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`rewrite_merge_version` = ?", rewriteMergeVersion).Find(&results).Error

	return
}

// GetBatchFromRewriteMergeVersion 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromRewriteMergeVersion(rewriteMergeVersions []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`rewrite_merge_version` IN (?)", rewriteMergeVersions).Find(&results).Error

	return
}

// GetFromLocality 通过locality获取内容
func (obj *_AllTenantHistoryMgr) GetFromLocality(locality string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`locality` = ?", locality).Find(&results).Error

	return
}

// GetBatchFromLocality 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromLocality(localitys []string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`locality` IN (?)", localitys).Find(&results).Error

	return
}

// GetFromLogonlyReplicaNum 通过logonly_replica_num获取内容
func (obj *_AllTenantHistoryMgr) GetFromLogonlyReplicaNum(logonlyReplicaNum int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`logonly_replica_num` = ?", logonlyReplicaNum).Find(&results).Error

	return
}

// GetBatchFromLogonlyReplicaNum 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromLogonlyReplicaNum(logonlyReplicaNums []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`logonly_replica_num` IN (?)", logonlyReplicaNums).Find(&results).Error

	return
}

// GetFromPreviousLocality 通过previous_locality获取内容
func (obj *_AllTenantHistoryMgr) GetFromPreviousLocality(previousLocality string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`previous_locality` = ?", previousLocality).Find(&results).Error

	return
}

// GetBatchFromPreviousLocality 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromPreviousLocality(previousLocalitys []string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`previous_locality` IN (?)", previousLocalitys).Find(&results).Error

	return
}

// GetFromStorageFormatVersion 通过storage_format_version获取内容
func (obj *_AllTenantHistoryMgr) GetFromStorageFormatVersion(storageFormatVersion int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`storage_format_version` = ?", storageFormatVersion).Find(&results).Error

	return
}

// GetBatchFromStorageFormatVersion 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromStorageFormatVersion(storageFormatVersions []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`storage_format_version` IN (?)", storageFormatVersions).Find(&results).Error

	return
}

// GetFromStorageFormatWorkVersion 通过storage_format_work_version获取内容
func (obj *_AllTenantHistoryMgr) GetFromStorageFormatWorkVersion(storageFormatWorkVersion int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`storage_format_work_version` = ?", storageFormatWorkVersion).Find(&results).Error

	return
}

// GetBatchFromStorageFormatWorkVersion 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromStorageFormatWorkVersion(storageFormatWorkVersions []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`storage_format_work_version` IN (?)", storageFormatWorkVersions).Find(&results).Error

	return
}

// GetFromDefaultTablegroupID 通过default_tablegroup_id获取内容
func (obj *_AllTenantHistoryMgr) GetFromDefaultTablegroupID(defaultTablegroupID int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`default_tablegroup_id` = ?", defaultTablegroupID).Find(&results).Error

	return
}

// GetBatchFromDefaultTablegroupID 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromDefaultTablegroupID(defaultTablegroupIDs []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`default_tablegroup_id` IN (?)", defaultTablegroupIDs).Find(&results).Error

	return
}

// GetFromCompatibilityMode 通过compatibility_mode获取内容
func (obj *_AllTenantHistoryMgr) GetFromCompatibilityMode(compatibilityMode int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`compatibility_mode` = ?", compatibilityMode).Find(&results).Error

	return
}

// GetBatchFromCompatibilityMode 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromCompatibilityMode(compatibilityModes []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`compatibility_mode` IN (?)", compatibilityModes).Find(&results).Error

	return
}

// GetFromDropTenantTime 通过drop_tenant_time获取内容
func (obj *_AllTenantHistoryMgr) GetFromDropTenantTime(dropTenantTime int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`drop_tenant_time` = ?", dropTenantTime).Find(&results).Error

	return
}

// GetBatchFromDropTenantTime 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromDropTenantTime(dropTenantTimes []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`drop_tenant_time` IN (?)", dropTenantTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantHistoryMgr) GetFromStatus(status string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromStatus(statuss []string) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromInRecyclebin 通过in_recyclebin获取内容
func (obj *_AllTenantHistoryMgr) GetFromInRecyclebin(inRecyclebin int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`in_recyclebin` = ?", inRecyclebin).Find(&results).Error

	return
}

// GetBatchFromInRecyclebin 批量查找
func (obj *_AllTenantHistoryMgr) GetBatchFromInRecyclebin(inRecyclebins []int64) (results []*AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`in_recyclebin` IN (?)", inRecyclebins).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantHistoryMgr) FetchByPrimaryKey(tenantID int64, schemaVersion int64) (result AllTenantHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantHistory{}).Where("`tenant_id` = ? AND `schema_version` = ?", tenantID, schemaVersion).First(&result).Error

	return
}
