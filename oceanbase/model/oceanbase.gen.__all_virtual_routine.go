package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualRoutineMgr struct {
	*_BaseMgr
}

// AllVirtualRoutineMgr open func
func AllVirtualRoutineMgr(db *gorm.DB) *_AllVirtualRoutineMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRoutineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRoutineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_routine"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRoutineMgr) GetTableName() string {
	return "__all_virtual_routine"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRoutineMgr) Reset() *_AllVirtualRoutineMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRoutineMgr) Get() (result AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRoutineMgr) Gets() (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRoutineMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRoutineMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRoutineID routine_id获取
func (obj *_AllVirtualRoutineMgr) WithRoutineID(routineID int64) Option {
	return optionFunc(func(o *options) { o.query["routine_id"] = routineID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualRoutineMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualRoutineMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualRoutineMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithPackageID package_id获取
func (obj *_AllVirtualRoutineMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithRoutineName routine_name获取
func (obj *_AllVirtualRoutineMgr) WithRoutineName(routineName string) Option {
	return optionFunc(func(o *options) { o.query["routine_name"] = routineName })
}

// WithOverload overload获取
func (obj *_AllVirtualRoutineMgr) WithOverload(overload int64) Option {
	return optionFunc(func(o *options) { o.query["overload"] = overload })
}

// WithSubprogramID subprogram_id获取
func (obj *_AllVirtualRoutineMgr) WithSubprogramID(subprogramID int64) Option {
	return optionFunc(func(o *options) { o.query["subprogram_id"] = subprogramID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualRoutineMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithRoutineType routine_type获取
func (obj *_AllVirtualRoutineMgr) WithRoutineType(routineType int64) Option {
	return optionFunc(func(o *options) { o.query["routine_type"] = routineType })
}

// WithFlag flag获取
func (obj *_AllVirtualRoutineMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualRoutineMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithPrivUser priv_user获取
func (obj *_AllVirtualRoutineMgr) WithPrivUser(privUser string) Option {
	return optionFunc(func(o *options) { o.query["priv_user"] = privUser })
}

// WithCompFlag comp_flag获取
func (obj *_AllVirtualRoutineMgr) WithCompFlag(compFlag int64) Option {
	return optionFunc(func(o *options) { o.query["comp_flag"] = compFlag })
}

// WithExecEnv exec_env获取
func (obj *_AllVirtualRoutineMgr) WithExecEnv(execEnv string) Option {
	return optionFunc(func(o *options) { o.query["exec_env"] = execEnv })
}

// WithRoutineBody routine_body获取
func (obj *_AllVirtualRoutineMgr) WithRoutineBody(routineBody string) Option {
	return optionFunc(func(o *options) { o.query["routine_body"] = routineBody })
}

// WithComment comment获取
func (obj *_AllVirtualRoutineMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRouteSQL route_sql获取
func (obj *_AllVirtualRoutineMgr) WithRouteSQL(routeSQL string) Option {
	return optionFunc(func(o *options) { o.query["route_sql"] = routeSQL })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRoutineMgr) GetByOption(opts ...Option) (result AllVirtualRoutine, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRoutineMgr) GetByOptions(opts ...Option) (results []*AllVirtualRoutine, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRoutineMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRoutine, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where(options.query)
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
func (obj *_AllVirtualRoutineMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoutineID 通过routine_id获取内容
func (obj *_AllVirtualRoutineMgr) GetFromRoutineID(routineID int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`routine_id` = ?", routineID).Find(&results).Error

	return
}

// GetBatchFromRoutineID 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromRoutineID(routineIDs []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`routine_id` IN (?)", routineIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualRoutineMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualRoutineMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualRoutineMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllVirtualRoutineMgr) GetFromPackageID(packageID int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromRoutineName 通过routine_name获取内容
func (obj *_AllVirtualRoutineMgr) GetFromRoutineName(routineName string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`routine_name` = ?", routineName).Find(&results).Error

	return
}

// GetBatchFromRoutineName 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromRoutineName(routineNames []string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`routine_name` IN (?)", routineNames).Find(&results).Error

	return
}

// GetFromOverload 通过overload获取内容
func (obj *_AllVirtualRoutineMgr) GetFromOverload(overload int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`overload` = ?", overload).Find(&results).Error

	return
}

// GetBatchFromOverload 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromOverload(overloads []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`overload` IN (?)", overloads).Find(&results).Error

	return
}

// GetFromSubprogramID 通过subprogram_id获取内容
func (obj *_AllVirtualRoutineMgr) GetFromSubprogramID(subprogramID int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`subprogram_id` = ?", subprogramID).Find(&results).Error

	return
}

// GetBatchFromSubprogramID 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromSubprogramID(subprogramIDs []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`subprogram_id` IN (?)", subprogramIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualRoutineMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromRoutineType 通过routine_type获取内容
func (obj *_AllVirtualRoutineMgr) GetFromRoutineType(routineType int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`routine_type` = ?", routineType).Find(&results).Error

	return
}

// GetBatchFromRoutineType 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromRoutineType(routineTypes []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`routine_type` IN (?)", routineTypes).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllVirtualRoutineMgr) GetFromFlag(flag int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromFlag(flags []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualRoutineMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromPrivUser 通过priv_user获取内容
func (obj *_AllVirtualRoutineMgr) GetFromPrivUser(privUser string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`priv_user` = ?", privUser).Find(&results).Error

	return
}

// GetBatchFromPrivUser 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromPrivUser(privUsers []string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`priv_user` IN (?)", privUsers).Find(&results).Error

	return
}

// GetFromCompFlag 通过comp_flag获取内容
func (obj *_AllVirtualRoutineMgr) GetFromCompFlag(compFlag int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`comp_flag` = ?", compFlag).Find(&results).Error

	return
}

// GetBatchFromCompFlag 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromCompFlag(compFlags []int64) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`comp_flag` IN (?)", compFlags).Find(&results).Error

	return
}

// GetFromExecEnv 通过exec_env获取内容
func (obj *_AllVirtualRoutineMgr) GetFromExecEnv(execEnv string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`exec_env` = ?", execEnv).Find(&results).Error

	return
}

// GetBatchFromExecEnv 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromExecEnv(execEnvs []string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`exec_env` IN (?)", execEnvs).Find(&results).Error

	return
}

// GetFromRoutineBody 通过routine_body获取内容
func (obj *_AllVirtualRoutineMgr) GetFromRoutineBody(routineBody string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`routine_body` = ?", routineBody).Find(&results).Error

	return
}

// GetBatchFromRoutineBody 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromRoutineBody(routineBodys []string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`routine_body` IN (?)", routineBodys).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualRoutineMgr) GetFromComment(comment string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromComment(comments []string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromRouteSQL 通过route_sql获取内容
func (obj *_AllVirtualRoutineMgr) GetFromRouteSQL(routeSQL string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`route_sql` = ?", routeSQL).Find(&results).Error

	return
}

// GetBatchFromRouteSQL 批量查找
func (obj *_AllVirtualRoutineMgr) GetBatchFromRouteSQL(routeSQLs []string) (results []*AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`route_sql` IN (?)", routeSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualRoutineMgr) FetchByPrimaryKey(tenantID int64, routineID int64) (result AllVirtualRoutine, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutine{}).Where("`tenant_id` = ? AND `routine_id` = ?", tenantID, routineID).First(&result).Error

	return
}
