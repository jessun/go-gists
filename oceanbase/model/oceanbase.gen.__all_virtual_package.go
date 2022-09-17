package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualPackageMgr struct {
	*_BaseMgr
}

// AllVirtualPackageMgr open func
func AllVirtualPackageMgr(db *gorm.DB) *_AllVirtualPackageMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPackageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPackageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_package"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPackageMgr) GetTableName() string {
	return "__all_virtual_package"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPackageMgr) Reset() *_AllVirtualPackageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPackageMgr) Get() (result AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPackageMgr) Gets() (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPackageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPackageMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPackageID package_id获取
func (obj *_AllVirtualPackageMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualPackageMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualPackageMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualPackageMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithPackageName package_name获取
func (obj *_AllVirtualPackageMgr) WithPackageName(packageName string) Option {
	return optionFunc(func(o *options) { o.query["package_name"] = packageName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualPackageMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithType type获取
func (obj *_AllVirtualPackageMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithFlag flag获取
func (obj *_AllVirtualPackageMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualPackageMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithCompFlag comp_flag获取
func (obj *_AllVirtualPackageMgr) WithCompFlag(compFlag int64) Option {
	return optionFunc(func(o *options) { o.query["comp_flag"] = compFlag })
}

// WithExecEnv exec_env获取
func (obj *_AllVirtualPackageMgr) WithExecEnv(execEnv string) Option {
	return optionFunc(func(o *options) { o.query["exec_env"] = execEnv })
}

// WithSource source获取
func (obj *_AllVirtualPackageMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithComment comment获取
func (obj *_AllVirtualPackageMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRouteSQL route_sql获取
func (obj *_AllVirtualPackageMgr) WithRouteSQL(routeSQL string) Option {
	return optionFunc(func(o *options) { o.query["route_sql"] = routeSQL })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPackageMgr) GetByOption(opts ...Option) (result AllVirtualPackage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPackageMgr) GetByOptions(opts ...Option) (results []*AllVirtualPackage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPackageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPackage, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where(options.query)
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
func (obj *_AllVirtualPackageMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllVirtualPackageMgr) GetFromPackageID(packageID int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualPackageMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualPackageMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualPackageMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromPackageName 通过package_name获取内容
func (obj *_AllVirtualPackageMgr) GetFromPackageName(packageName string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`package_name` = ?", packageName).Find(&results).Error

	return
}

// GetBatchFromPackageName 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromPackageName(packageNames []string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`package_name` IN (?)", packageNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualPackageMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualPackageMgr) GetFromType(_type int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromType(_types []int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllVirtualPackageMgr) GetFromFlag(flag int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromFlag(flags []int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualPackageMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromCompFlag 通过comp_flag获取内容
func (obj *_AllVirtualPackageMgr) GetFromCompFlag(compFlag int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`comp_flag` = ?", compFlag).Find(&results).Error

	return
}

// GetBatchFromCompFlag 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromCompFlag(compFlags []int64) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`comp_flag` IN (?)", compFlags).Find(&results).Error

	return
}

// GetFromExecEnv 通过exec_env获取内容
func (obj *_AllVirtualPackageMgr) GetFromExecEnv(execEnv string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`exec_env` = ?", execEnv).Find(&results).Error

	return
}

// GetBatchFromExecEnv 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromExecEnv(execEnvs []string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`exec_env` IN (?)", execEnvs).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllVirtualPackageMgr) GetFromSource(source string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromSource(sources []string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualPackageMgr) GetFromComment(comment string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromComment(comments []string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromRouteSQL 通过route_sql获取内容
func (obj *_AllVirtualPackageMgr) GetFromRouteSQL(routeSQL string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`route_sql` = ?", routeSQL).Find(&results).Error

	return
}

// GetBatchFromRouteSQL 批量查找
func (obj *_AllVirtualPackageMgr) GetBatchFromRouteSQL(routeSQLs []string) (results []*AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`route_sql` IN (?)", routeSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPackageMgr) FetchByPrimaryKey(tenantID int64, packageID int64) (result AllVirtualPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPackage{}).Where("`tenant_id` = ? AND `package_id` = ?", tenantID, packageID).First(&result).Error

	return
}
