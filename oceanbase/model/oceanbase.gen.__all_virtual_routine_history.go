package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualRoutineHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualRoutineHistoryMgr open func
func AllVirtualRoutineHistoryMgr(db *gorm.DB) *_AllVirtualRoutineHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRoutineHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRoutineHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_routine_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRoutineHistoryMgr) GetTableName() string {
	return "__all_virtual_routine_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRoutineHistoryMgr) Reset() *_AllVirtualRoutineHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRoutineHistoryMgr) Get() (result AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRoutineHistoryMgr) Gets() (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRoutineHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRoutineHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRoutineID routine_id获取
func (obj *_AllVirtualRoutineHistoryMgr) WithRoutineID(routineID int64) Option {
	return optionFunc(func(o *options) { o.query["routine_id"] = routineID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualRoutineHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualRoutineHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualRoutineHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualRoutineHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualRoutineHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithPackageID package_id获取
func (obj *_AllVirtualRoutineHistoryMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithRoutineName routine_name获取
func (obj *_AllVirtualRoutineHistoryMgr) WithRoutineName(routineName string) Option {
	return optionFunc(func(o *options) { o.query["routine_name"] = routineName })
}

// WithOverload overload获取
func (obj *_AllVirtualRoutineHistoryMgr) WithOverload(overload int64) Option {
	return optionFunc(func(o *options) { o.query["overload"] = overload })
}

// WithSubprogramID subprogram_id获取
func (obj *_AllVirtualRoutineHistoryMgr) WithSubprogramID(subprogramID int64) Option {
	return optionFunc(func(o *options) { o.query["subprogram_id"] = subprogramID })
}

// WithRoutineType routine_type获取
func (obj *_AllVirtualRoutineHistoryMgr) WithRoutineType(routineType int64) Option {
	return optionFunc(func(o *options) { o.query["routine_type"] = routineType })
}

// WithFlag flag获取
func (obj *_AllVirtualRoutineHistoryMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualRoutineHistoryMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithPrivUser priv_user获取
func (obj *_AllVirtualRoutineHistoryMgr) WithPrivUser(privUser string) Option {
	return optionFunc(func(o *options) { o.query["priv_user"] = privUser })
}

// WithCompFlag comp_flag获取
func (obj *_AllVirtualRoutineHistoryMgr) WithCompFlag(compFlag int64) Option {
	return optionFunc(func(o *options) { o.query["comp_flag"] = compFlag })
}

// WithExecEnv exec_env获取
func (obj *_AllVirtualRoutineHistoryMgr) WithExecEnv(execEnv string) Option {
	return optionFunc(func(o *options) { o.query["exec_env"] = execEnv })
}

// WithRoutineBody routine_body获取
func (obj *_AllVirtualRoutineHistoryMgr) WithRoutineBody(routineBody string) Option {
	return optionFunc(func(o *options) { o.query["routine_body"] = routineBody })
}

// WithComment comment获取
func (obj *_AllVirtualRoutineHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRouteSQL route_sql获取
func (obj *_AllVirtualRoutineHistoryMgr) WithRouteSQL(routeSQL string) Option {
	return optionFunc(func(o *options) { o.query["route_sql"] = routeSQL })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRoutineHistoryMgr) GetByOption(opts ...Option) (result AllVirtualRoutineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRoutineHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualRoutineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRoutineHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRoutineHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where(options.query)
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
func (obj *_AllVirtualRoutineHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoutineID 通过routine_id获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromRoutineID(routineID int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`routine_id` = ?", routineID).Find(&results).Error

	return
}

// GetBatchFromRoutineID 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromRoutineID(routineIDs []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`routine_id` IN (?)", routineIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromPackageID(packageID int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromRoutineName 通过routine_name获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromRoutineName(routineName string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`routine_name` = ?", routineName).Find(&results).Error

	return
}

// GetBatchFromRoutineName 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromRoutineName(routineNames []string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`routine_name` IN (?)", routineNames).Find(&results).Error

	return
}

// GetFromOverload 通过overload获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromOverload(overload int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`overload` = ?", overload).Find(&results).Error

	return
}

// GetBatchFromOverload 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromOverload(overloads []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`overload` IN (?)", overloads).Find(&results).Error

	return
}

// GetFromSubprogramID 通过subprogram_id获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromSubprogramID(subprogramID int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`subprogram_id` = ?", subprogramID).Find(&results).Error

	return
}

// GetBatchFromSubprogramID 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromSubprogramID(subprogramIDs []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`subprogram_id` IN (?)", subprogramIDs).Find(&results).Error

	return
}

// GetFromRoutineType 通过routine_type获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromRoutineType(routineType int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`routine_type` = ?", routineType).Find(&results).Error

	return
}

// GetBatchFromRoutineType 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromRoutineType(routineTypes []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`routine_type` IN (?)", routineTypes).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromFlag(flag int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromFlag(flags []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromPrivUser 通过priv_user获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromPrivUser(privUser string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`priv_user` = ?", privUser).Find(&results).Error

	return
}

// GetBatchFromPrivUser 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromPrivUser(privUsers []string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`priv_user` IN (?)", privUsers).Find(&results).Error

	return
}

// GetFromCompFlag 通过comp_flag获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromCompFlag(compFlag int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`comp_flag` = ?", compFlag).Find(&results).Error

	return
}

// GetBatchFromCompFlag 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromCompFlag(compFlags []int64) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`comp_flag` IN (?)", compFlags).Find(&results).Error

	return
}

// GetFromExecEnv 通过exec_env获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromExecEnv(execEnv string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`exec_env` = ?", execEnv).Find(&results).Error

	return
}

// GetBatchFromExecEnv 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromExecEnv(execEnvs []string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`exec_env` IN (?)", execEnvs).Find(&results).Error

	return
}

// GetFromRoutineBody 通过routine_body获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromRoutineBody(routineBody string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`routine_body` = ?", routineBody).Find(&results).Error

	return
}

// GetBatchFromRoutineBody 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromRoutineBody(routineBodys []string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`routine_body` IN (?)", routineBodys).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromComment(comment string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromComment(comments []string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromRouteSQL 通过route_sql获取内容
func (obj *_AllVirtualRoutineHistoryMgr) GetFromRouteSQL(routeSQL string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`route_sql` = ?", routeSQL).Find(&results).Error

	return
}

// GetBatchFromRouteSQL 批量查找
func (obj *_AllVirtualRoutineHistoryMgr) GetBatchFromRouteSQL(routeSQLs []string) (results []*AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`route_sql` IN (?)", routeSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualRoutineHistoryMgr) FetchByPrimaryKey(tenantID int64, routineID int64, schemaVersion int64) (result AllVirtualRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineHistory{}).Where("`tenant_id` = ? AND `routine_id` = ? AND `schema_version` = ?", tenantID, routineID, schemaVersion).First(&result).Error

	return
}
