package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualPackageHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualPackageHistoryMgr open func
func AllVirtualPackageHistoryMgr(db *gorm.DB) *_AllVirtualPackageHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPackageHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPackageHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_package_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPackageHistoryMgr) GetTableName() string {
	return "__all_virtual_package_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPackageHistoryMgr) Reset() *_AllVirtualPackageHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPackageHistoryMgr) Get() (result AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPackageHistoryMgr) Gets() (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPackageHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPackageHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPackageID package_id获取
func (obj *_AllVirtualPackageHistoryMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualPackageHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualPackageHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualPackageHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualPackageHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualPackageHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithPackageName package_name获取
func (obj *_AllVirtualPackageHistoryMgr) WithPackageName(packageName string) Option {
	return optionFunc(func(o *options) { o.query["package_name"] = packageName })
}

// WithType type获取
func (obj *_AllVirtualPackageHistoryMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithFlag flag获取
func (obj *_AllVirtualPackageHistoryMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualPackageHistoryMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithCompFlag comp_flag获取
func (obj *_AllVirtualPackageHistoryMgr) WithCompFlag(compFlag int64) Option {
	return optionFunc(func(o *options) { o.query["comp_flag"] = compFlag })
}

// WithExecEnv exec_env获取
func (obj *_AllVirtualPackageHistoryMgr) WithExecEnv(execEnv string) Option {
	return optionFunc(func(o *options) { o.query["exec_env"] = execEnv })
}

// WithSource source获取
func (obj *_AllVirtualPackageHistoryMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithComment comment获取
func (obj *_AllVirtualPackageHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRouteSQL route_sql获取
func (obj *_AllVirtualPackageHistoryMgr) WithRouteSQL(routeSQL string) Option {
	return optionFunc(func(o *options) { o.query["route_sql"] = routeSQL })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPackageHistoryMgr) GetByOption(opts ...Option) (result AllVirtualPackageHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPackageHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualPackageHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPackageHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPackageHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where(options.query)
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
func (obj *_AllVirtualPackageHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromPackageID(packageID int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromPackageName 通过package_name获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromPackageName(packageName string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`package_name` = ?", packageName).Find(&results).Error

	return
}

// GetBatchFromPackageName 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromPackageName(packageNames []string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`package_name` IN (?)", packageNames).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromType(_type int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromType(_types []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromFlag(flag int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromFlag(flags []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromCompFlag 通过comp_flag获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromCompFlag(compFlag int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`comp_flag` = ?", compFlag).Find(&results).Error

	return
}

// GetBatchFromCompFlag 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromCompFlag(compFlags []int64) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`comp_flag` IN (?)", compFlags).Find(&results).Error

	return
}

// GetFromExecEnv 通过exec_env获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromExecEnv(execEnv string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`exec_env` = ?", execEnv).Find(&results).Error

	return
}

// GetBatchFromExecEnv 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromExecEnv(execEnvs []string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`exec_env` IN (?)", execEnvs).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromSource(source string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromSource(sources []string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromComment(comment string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromComment(comments []string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromRouteSQL 通过route_sql获取内容
func (obj *_AllVirtualPackageHistoryMgr) GetFromRouteSQL(routeSQL string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`route_sql` = ?", routeSQL).Find(&results).Error

	return
}

// GetBatchFromRouteSQL 批量查找
func (obj *_AllVirtualPackageHistoryMgr) GetBatchFromRouteSQL(routeSQLs []string) (results []*AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`route_sql` IN (?)", routeSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPackageHistoryMgr) FetchByPrimaryKey(tenantID int64, packageID int64, schemaVersion int64) (result AllVirtualPackageHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackageHistory{}).Where("`tenant_id` = ? AND `package_id` = ? AND `schema_version` = ?", tenantID, packageID, schemaVersion).First(&result).Error

	return
}
