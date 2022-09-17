package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllDatabaseHistoryMgr struct {
	*_BaseMgr
}

// AllDatabaseHistoryMgr open func
func AllDatabaseHistoryMgr(db *gorm.DB) *_AllDatabaseHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllDatabaseHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDatabaseHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_database_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDatabaseHistoryMgr) GetTableName() string {
	return "__all_database_history"
}

// Reset 重置gorm会话
func (obj *_AllDatabaseHistoryMgr) Reset() *_AllDatabaseHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDatabaseHistoryMgr) Get() (result AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDatabaseHistoryMgr) Gets() (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDatabaseHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDatabaseHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDatabaseHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllDatabaseHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDatabaseID database_id获取
func (obj *_AllDatabaseHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllDatabaseHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllDatabaseHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseName database_name获取
func (obj *_AllDatabaseHistoryMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithReplicaNum replica_num获取
func (obj *_AllDatabaseHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithZoneList zone_list获取
func (obj *_AllDatabaseHistoryMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllDatabaseHistoryMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithCollationType collation_type获取
func (obj *_AllDatabaseHistoryMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithComment comment获取
func (obj *_AllDatabaseHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithReadOnly read_only获取
func (obj *_AllDatabaseHistoryMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithDefaultTablegroupID default_tablegroup_id获取
func (obj *_AllDatabaseHistoryMgr) WithDefaultTablegroupID(defaultTablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["default_tablegroup_id"] = defaultTablegroupID })
}

// WithInRecyclebin in_recyclebin获取
func (obj *_AllDatabaseHistoryMgr) WithInRecyclebin(inRecyclebin int64) Option {
	return optionFunc(func(o *options) { o.query["in_recyclebin"] = inRecyclebin })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllDatabaseHistoryMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllDatabaseHistoryMgr) GetByOption(opts ...Option) (result AllDatabaseHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDatabaseHistoryMgr) GetByOptions(opts ...Option) (results []*AllDatabaseHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDatabaseHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDatabaseHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where(options.query)
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
func (obj *_AllDatabaseHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromDatabaseName(databaseName string) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromZoneList(zoneList string) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromPrimaryZone(primaryZone string) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromCollationType(collationType int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromComment(comment string) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromComment(comments []string) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromReadOnly(readOnly int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

// GetFromDefaultTablegroupID 通过default_tablegroup_id获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromDefaultTablegroupID(defaultTablegroupID int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`default_tablegroup_id` = ?", defaultTablegroupID).Find(&results).Error

	return
}

// GetBatchFromDefaultTablegroupID 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromDefaultTablegroupID(defaultTablegroupIDs []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`default_tablegroup_id` IN (?)", defaultTablegroupIDs).Find(&results).Error

	return
}

// GetFromInRecyclebin 通过in_recyclebin获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromInRecyclebin(inRecyclebin int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`in_recyclebin` = ?", inRecyclebin).Find(&results).Error

	return
}

// GetBatchFromInRecyclebin 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromInRecyclebin(inRecyclebins []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`in_recyclebin` IN (?)", inRecyclebins).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllDatabaseHistoryMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllDatabaseHistoryMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDatabaseHistoryMgr) FetchByPrimaryKey(tenantID int64, databaseID int64, schemaVersion int64) (result AllDatabaseHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabaseHistory{}).Where("`tenant_id` = ? AND `database_id` = ? AND `schema_version` = ?", tenantID, databaseID, schemaVersion).First(&result).Error

	return
}
