package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualDatabaseHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualDatabaseHistoryMgr open func
func AllVirtualDatabaseHistoryMgr(db *gorm.DB) *_AllVirtualDatabaseHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDatabaseHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDatabaseHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_database_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDatabaseHistoryMgr) GetTableName() string {
	return "__all_virtual_database_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDatabaseHistoryMgr) Reset() *_AllVirtualDatabaseHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDatabaseHistoryMgr) Get() (result AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDatabaseHistoryMgr) Gets() (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDatabaseHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseName database_name获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithZoneList zone_list获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithCollationType collation_type获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithComment comment获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithReadOnly read_only获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithDefaultTablegroupID default_tablegroup_id获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithDefaultTablegroupID(defaultTablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["default_tablegroup_id"] = defaultTablegroupID })
}

// WithInRecyclebin in_recyclebin获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithInRecyclebin(inRecyclebin int64) Option {
	return optionFunc(func(o *options) { o.query["in_recyclebin"] = inRecyclebin })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllVirtualDatabaseHistoryMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDatabaseHistoryMgr) GetByOption(opts ...Option) (result AllVirtualDatabaseHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDatabaseHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualDatabaseHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDatabaseHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDatabaseHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where(options.query)
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
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromDatabaseName(databaseName string) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromZoneList(zoneList string) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromPrimaryZone(primaryZone string) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromCollationType(collationType int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromComment(comment string) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromComment(comments []string) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromReadOnly(readOnly int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

// GetFromDefaultTablegroupID 通过default_tablegroup_id获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromDefaultTablegroupID(defaultTablegroupID int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`default_tablegroup_id` = ?", defaultTablegroupID).Find(&results).Error

	return
}

// GetBatchFromDefaultTablegroupID 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromDefaultTablegroupID(defaultTablegroupIDs []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`default_tablegroup_id` IN (?)", defaultTablegroupIDs).Find(&results).Error

	return
}

// GetFromInRecyclebin 通过in_recyclebin获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromInRecyclebin(inRecyclebin int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`in_recyclebin` = ?", inRecyclebin).Find(&results).Error

	return
}

// GetBatchFromInRecyclebin 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromInRecyclebin(inRecyclebins []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`in_recyclebin` IN (?)", inRecyclebins).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllVirtualDatabaseHistoryMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllVirtualDatabaseHistoryMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDatabaseHistoryMgr) FetchByPrimaryKey(tenantID int64, databaseID int64, schemaVersion int64) (result AllVirtualDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabaseHistory{}).Where("`tenant_id` = ? AND `database_id` = ? AND `schema_version` = ?", tenantID, databaseID, schemaVersion).First(&result).Error

	return
}
