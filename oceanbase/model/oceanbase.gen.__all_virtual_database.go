package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualDatabaseMgr struct {
	*_BaseMgr
}

// AllVirtualDatabaseMgr open func
func AllVirtualDatabaseMgr(db *gorm.DB) *_AllVirtualDatabaseMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDatabaseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDatabaseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_database"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDatabaseMgr) GetTableName() string {
	return "__all_virtual_database"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDatabaseMgr) Reset() *_AllVirtualDatabaseMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDatabaseMgr) Get() (result AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDatabaseMgr) Gets() (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDatabaseMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDatabaseMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualDatabaseMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDatabaseMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDatabaseMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDatabaseName database_name获取
func (obj *_AllVirtualDatabaseMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualDatabaseMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithZoneList zone_list获取
func (obj *_AllVirtualDatabaseMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllVirtualDatabaseMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithCollationType collation_type获取
func (obj *_AllVirtualDatabaseMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithComment comment获取
func (obj *_AllVirtualDatabaseMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithReadOnly read_only获取
func (obj *_AllVirtualDatabaseMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithDefaultTablegroupID default_tablegroup_id获取
func (obj *_AllVirtualDatabaseMgr) WithDefaultTablegroupID(defaultTablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["default_tablegroup_id"] = defaultTablegroupID })
}

// WithInRecyclebin in_recyclebin获取
func (obj *_AllVirtualDatabaseMgr) WithInRecyclebin(inRecyclebin int64) Option {
	return optionFunc(func(o *options) { o.query["in_recyclebin"] = inRecyclebin })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllVirtualDatabaseMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDatabaseMgr) GetByOption(opts ...Option) (result AllVirtualDatabase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDatabaseMgr) GetByOptions(opts ...Option) (results []*AllVirtualDatabase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDatabaseMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDatabase, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where(options.query)
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
func (obj *_AllVirtualDatabaseMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromDatabaseName(databaseName string) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromZoneList(zoneList string) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromPrimaryZone(primaryZone string) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromCollationType(collationType int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromComment(comment string) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromComment(comments []string) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromReadOnly(readOnly int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

// GetFromDefaultTablegroupID 通过default_tablegroup_id获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromDefaultTablegroupID(defaultTablegroupID int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`default_tablegroup_id` = ?", defaultTablegroupID).Find(&results).Error

	return
}

// GetBatchFromDefaultTablegroupID 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromDefaultTablegroupID(defaultTablegroupIDs []int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`default_tablegroup_id` IN (?)", defaultTablegroupIDs).Find(&results).Error

	return
}

// GetFromInRecyclebin 通过in_recyclebin获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromInRecyclebin(inRecyclebin int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`in_recyclebin` = ?", inRecyclebin).Find(&results).Error

	return
}

// GetBatchFromInRecyclebin 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromInRecyclebin(inRecyclebins []int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`in_recyclebin` IN (?)", inRecyclebins).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllVirtualDatabaseMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllVirtualDatabaseMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDatabaseMgr) FetchByPrimaryKey(tenantID int64, databaseID int64) (result AllVirtualDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDatabase{}).Where("`tenant_id` = ? AND `database_id` = ?", tenantID, databaseID).First(&result).Error

	return
}
