package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllDatabaseMgr struct {
	*_BaseMgr
}

// AllDatabaseMgr open func
func AllDatabaseMgr(db *gorm.DB) *_AllDatabaseMgr {
	if db == nil {
		panic(fmt.Errorf("AllDatabaseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDatabaseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_database"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDatabaseMgr) GetTableName() string {
	return "__all_database"
}

// Reset 重置gorm会话
func (obj *_AllDatabaseMgr) Reset() *_AllDatabaseMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDatabaseMgr) Get() (result AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDatabaseMgr) Gets() (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDatabaseMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDatabaseMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDatabaseMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllDatabaseMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDatabaseID database_id获取
func (obj *_AllDatabaseMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithDatabaseName database_name获取
func (obj *_AllDatabaseMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithReplicaNum replica_num获取
func (obj *_AllDatabaseMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithZoneList zone_list获取
func (obj *_AllDatabaseMgr) WithZoneList(zoneList string) Option {
	return optionFunc(func(o *options) { o.query["zone_list"] = zoneList })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllDatabaseMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithCollationType collation_type获取
func (obj *_AllDatabaseMgr) WithCollationType(collationType int64) Option {
	return optionFunc(func(o *options) { o.query["collation_type"] = collationType })
}

// WithComment comment获取
func (obj *_AllDatabaseMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithReadOnly read_only获取
func (obj *_AllDatabaseMgr) WithReadOnly(readOnly int64) Option {
	return optionFunc(func(o *options) { o.query["read_only"] = readOnly })
}

// WithDefaultTablegroupID default_tablegroup_id获取
func (obj *_AllDatabaseMgr) WithDefaultTablegroupID(defaultTablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["default_tablegroup_id"] = defaultTablegroupID })
}

// WithInRecyclebin in_recyclebin获取
func (obj *_AllDatabaseMgr) WithInRecyclebin(inRecyclebin int64) Option {
	return optionFunc(func(o *options) { o.query["in_recyclebin"] = inRecyclebin })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllDatabaseMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllDatabaseMgr) GetByOption(opts ...Option) (result AllDatabase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDatabaseMgr) GetByOptions(opts ...Option) (results []*AllDatabase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDatabaseMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDatabase, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where(options.query)
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
func (obj *_AllDatabaseMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDatabaseMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDatabaseMgr) GetFromTenantID(tenantID int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllDatabaseMgr) GetFromDatabaseID(databaseID int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllDatabaseMgr) GetFromDatabaseName(databaseName string) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllDatabaseMgr) GetFromReplicaNum(replicaNum int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromZoneList 通过zone_list获取内容
func (obj *_AllDatabaseMgr) GetFromZoneList(zoneList string) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`zone_list` = ?", zoneList).Find(&results).Error

	return
}

// GetBatchFromZoneList 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromZoneList(zoneLists []string) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`zone_list` IN (?)", zoneLists).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllDatabaseMgr) GetFromPrimaryZone(primaryZone string) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromCollationType 通过collation_type获取内容
func (obj *_AllDatabaseMgr) GetFromCollationType(collationType int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`collation_type` = ?", collationType).Find(&results).Error

	return
}

// GetBatchFromCollationType 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromCollationType(collationTypes []int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`collation_type` IN (?)", collationTypes).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllDatabaseMgr) GetFromComment(comment string) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromComment(comments []string) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromReadOnly 通过read_only获取内容
func (obj *_AllDatabaseMgr) GetFromReadOnly(readOnly int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`read_only` = ?", readOnly).Find(&results).Error

	return
}

// GetBatchFromReadOnly 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromReadOnly(readOnlys []int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`read_only` IN (?)", readOnlys).Find(&results).Error

	return
}

// GetFromDefaultTablegroupID 通过default_tablegroup_id获取内容
func (obj *_AllDatabaseMgr) GetFromDefaultTablegroupID(defaultTablegroupID int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`default_tablegroup_id` = ?", defaultTablegroupID).Find(&results).Error

	return
}

// GetBatchFromDefaultTablegroupID 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromDefaultTablegroupID(defaultTablegroupIDs []int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`default_tablegroup_id` IN (?)", defaultTablegroupIDs).Find(&results).Error

	return
}

// GetFromInRecyclebin 通过in_recyclebin获取内容
func (obj *_AllDatabaseMgr) GetFromInRecyclebin(inRecyclebin int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`in_recyclebin` = ?", inRecyclebin).Find(&results).Error

	return
}

// GetBatchFromInRecyclebin 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromInRecyclebin(inRecyclebins []int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`in_recyclebin` IN (?)", inRecyclebins).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllDatabaseMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllDatabaseMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDatabaseMgr) FetchByPrimaryKey(tenantID int64, databaseID int64) (result AllDatabase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatabase{}).Where("`tenant_id` = ? AND `database_id` = ?", tenantID, databaseID).First(&result).Error

	return
}
