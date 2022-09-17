package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllPackageMgr struct {
	*_BaseMgr
}

// AllPackageMgr open func
func AllPackageMgr(db *gorm.DB) *_AllPackageMgr {
	if db == nil {
		panic(fmt.Errorf("AllPackageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllPackageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_package"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllPackageMgr) GetTableName() string {
	return "__all_package"
}

// Reset 重置gorm会话
func (obj *_AllPackageMgr) Reset() *_AllPackageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllPackageMgr) Get() (result AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllPackageMgr) Gets() (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllPackageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllPackageMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllPackageMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllPackageMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPackageID package_id获取
func (obj *_AllPackageMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithDatabaseID database_id获取
func (obj *_AllPackageMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithPackageName package_name获取
func (obj *_AllPackageMgr) WithPackageName(packageName string) Option {
	return optionFunc(func(o *options) { o.query["package_name"] = packageName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllPackageMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithType type获取
func (obj *_AllPackageMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithFlag flag获取
func (obj *_AllPackageMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithOwnerID owner_id获取
func (obj *_AllPackageMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithCompFlag comp_flag获取
func (obj *_AllPackageMgr) WithCompFlag(compFlag int64) Option {
	return optionFunc(func(o *options) { o.query["comp_flag"] = compFlag })
}

// WithExecEnv exec_env获取
func (obj *_AllPackageMgr) WithExecEnv(execEnv string) Option {
	return optionFunc(func(o *options) { o.query["exec_env"] = execEnv })
}

// WithSource source获取
func (obj *_AllPackageMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithComment comment获取
func (obj *_AllPackageMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRouteSQL route_sql获取
func (obj *_AllPackageMgr) WithRouteSQL(routeSQL string) Option {
	return optionFunc(func(o *options) { o.query["route_sql"] = routeSQL })
}

// GetByOption 功能选项模式获取
func (obj *_AllPackageMgr) GetByOption(opts ...Option) (result AllPackage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllPackageMgr) GetByOptions(opts ...Option) (results []*AllPackage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllPackageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllPackage, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where(options.query)
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
func (obj *_AllPackageMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllPackageMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllPackageMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllPackageMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllPackageMgr) GetFromTenantID(tenantID int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllPackageMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllPackageMgr) GetFromPackageID(packageID int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllPackageMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllPackageMgr) GetFromDatabaseID(databaseID int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllPackageMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromPackageName 通过package_name获取内容
func (obj *_AllPackageMgr) GetFromPackageName(packageName string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`package_name` = ?", packageName).Find(&results).Error

	return
}

// GetBatchFromPackageName 批量查找
func (obj *_AllPackageMgr) GetBatchFromPackageName(packageNames []string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`package_name` IN (?)", packageNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllPackageMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllPackageMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllPackageMgr) GetFromType(_type int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllPackageMgr) GetBatchFromType(_types []int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllPackageMgr) GetFromFlag(flag int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllPackageMgr) GetBatchFromFlag(flags []int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllPackageMgr) GetFromOwnerID(ownerID int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllPackageMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromCompFlag 通过comp_flag获取内容
func (obj *_AllPackageMgr) GetFromCompFlag(compFlag int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`comp_flag` = ?", compFlag).Find(&results).Error

	return
}

// GetBatchFromCompFlag 批量查找
func (obj *_AllPackageMgr) GetBatchFromCompFlag(compFlags []int64) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`comp_flag` IN (?)", compFlags).Find(&results).Error

	return
}

// GetFromExecEnv 通过exec_env获取内容
func (obj *_AllPackageMgr) GetFromExecEnv(execEnv string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`exec_env` = ?", execEnv).Find(&results).Error

	return
}

// GetBatchFromExecEnv 批量查找
func (obj *_AllPackageMgr) GetBatchFromExecEnv(execEnvs []string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`exec_env` IN (?)", execEnvs).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllPackageMgr) GetFromSource(source string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllPackageMgr) GetBatchFromSource(sources []string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllPackageMgr) GetFromComment(comment string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllPackageMgr) GetBatchFromComment(comments []string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromRouteSQL 通过route_sql获取内容
func (obj *_AllPackageMgr) GetFromRouteSQL(routeSQL string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`route_sql` = ?", routeSQL).Find(&results).Error

	return
}

// GetBatchFromRouteSQL 批量查找
func (obj *_AllPackageMgr) GetBatchFromRouteSQL(routeSQLs []string) (results []*AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`route_sql` IN (?)", routeSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllPackageMgr) FetchByPrimaryKey(tenantID int64, packageID int64) (result AllPackage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPackage{}).Where("`tenant_id` = ? AND `package_id` = ?", tenantID, packageID).First(&result).Error

	return
}
