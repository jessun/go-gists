package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllRoutineHistoryMgr struct {
	*_BaseMgr
}

// AllRoutineHistoryMgr open func
func AllRoutineHistoryMgr(db *gorm.DB) *_AllRoutineHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllRoutineHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllRoutineHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_routine_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllRoutineHistoryMgr) GetTableName() string {
	return "__all_routine_history"
}

// Reset 重置gorm会话
func (obj *_AllRoutineHistoryMgr) Reset() *_AllRoutineHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllRoutineHistoryMgr) Get() (result AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllRoutineHistoryMgr) Gets() (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllRoutineHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllRoutineHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllRoutineHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllRoutineHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRoutineID routine_id获取
func (obj *_AllRoutineHistoryMgr) WithRoutineID(routineID int64) Option {
	return optionFunc(func(o *options) { o.query["routine_id"] = routineID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllRoutineHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllRoutineHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllRoutineHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithPackageID package_id获取
func (obj *_AllRoutineHistoryMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithRoutineName routine_name获取
func (obj *_AllRoutineHistoryMgr) WithRoutineName(routineName string) Option {
	return optionFunc(func(o *options) { o.query["routine_name"] = routineName })
}

// WithOverload overload获取
func (obj *_AllRoutineHistoryMgr) WithOverload(overload int64) Option {
	return optionFunc(func(o *options) { o.query["overload"] = overload })
}

// WithSubprogramID subprogram_id获取
func (obj *_AllRoutineHistoryMgr) WithSubprogramID(subprogramID int64) Option {
	return optionFunc(func(o *options) { o.query["subprogram_id"] = subprogramID })
}

// WithRoutineType routine_type获取
func (obj *_AllRoutineHistoryMgr) WithRoutineType(routineType int64) Option {
	return optionFunc(func(o *options) { o.query["routine_type"] = routineType })
}

// WithFlag flag获取
func (obj *_AllRoutineHistoryMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithOwnerID owner_id获取
func (obj *_AllRoutineHistoryMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithPrivUser priv_user获取
func (obj *_AllRoutineHistoryMgr) WithPrivUser(privUser string) Option {
	return optionFunc(func(o *options) { o.query["priv_user"] = privUser })
}

// WithCompFlag comp_flag获取
func (obj *_AllRoutineHistoryMgr) WithCompFlag(compFlag int64) Option {
	return optionFunc(func(o *options) { o.query["comp_flag"] = compFlag })
}

// WithExecEnv exec_env获取
func (obj *_AllRoutineHistoryMgr) WithExecEnv(execEnv string) Option {
	return optionFunc(func(o *options) { o.query["exec_env"] = execEnv })
}

// WithRoutineBody routine_body获取
func (obj *_AllRoutineHistoryMgr) WithRoutineBody(routineBody string) Option {
	return optionFunc(func(o *options) { o.query["routine_body"] = routineBody })
}

// WithComment comment获取
func (obj *_AllRoutineHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRouteSQL route_sql获取
func (obj *_AllRoutineHistoryMgr) WithRouteSQL(routeSQL string) Option {
	return optionFunc(func(o *options) { o.query["route_sql"] = routeSQL })
}

// GetByOption 功能选项模式获取
func (obj *_AllRoutineHistoryMgr) GetByOption(opts ...Option) (result AllRoutineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllRoutineHistoryMgr) GetByOptions(opts ...Option) (results []*AllRoutineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllRoutineHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllRoutineHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where(options.query)
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
func (obj *_AllRoutineHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllRoutineHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllRoutineHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoutineID 通过routine_id获取内容
func (obj *_AllRoutineHistoryMgr) GetFromRoutineID(routineID int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`routine_id` = ?", routineID).Find(&results).Error

	return
}

// GetBatchFromRoutineID 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromRoutineID(routineIDs []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`routine_id` IN (?)", routineIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllRoutineHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllRoutineHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllRoutineHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllRoutineHistoryMgr) GetFromPackageID(packageID int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromRoutineName 通过routine_name获取内容
func (obj *_AllRoutineHistoryMgr) GetFromRoutineName(routineName string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`routine_name` = ?", routineName).Find(&results).Error

	return
}

// GetBatchFromRoutineName 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromRoutineName(routineNames []string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`routine_name` IN (?)", routineNames).Find(&results).Error

	return
}

// GetFromOverload 通过overload获取内容
func (obj *_AllRoutineHistoryMgr) GetFromOverload(overload int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`overload` = ?", overload).Find(&results).Error

	return
}

// GetBatchFromOverload 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromOverload(overloads []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`overload` IN (?)", overloads).Find(&results).Error

	return
}

// GetFromSubprogramID 通过subprogram_id获取内容
func (obj *_AllRoutineHistoryMgr) GetFromSubprogramID(subprogramID int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`subprogram_id` = ?", subprogramID).Find(&results).Error

	return
}

// GetBatchFromSubprogramID 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromSubprogramID(subprogramIDs []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`subprogram_id` IN (?)", subprogramIDs).Find(&results).Error

	return
}

// GetFromRoutineType 通过routine_type获取内容
func (obj *_AllRoutineHistoryMgr) GetFromRoutineType(routineType int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`routine_type` = ?", routineType).Find(&results).Error

	return
}

// GetBatchFromRoutineType 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromRoutineType(routineTypes []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`routine_type` IN (?)", routineTypes).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllRoutineHistoryMgr) GetFromFlag(flag int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromFlag(flags []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllRoutineHistoryMgr) GetFromOwnerID(ownerID int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromPrivUser 通过priv_user获取内容
func (obj *_AllRoutineHistoryMgr) GetFromPrivUser(privUser string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`priv_user` = ?", privUser).Find(&results).Error

	return
}

// GetBatchFromPrivUser 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromPrivUser(privUsers []string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`priv_user` IN (?)", privUsers).Find(&results).Error

	return
}

// GetFromCompFlag 通过comp_flag获取内容
func (obj *_AllRoutineHistoryMgr) GetFromCompFlag(compFlag int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`comp_flag` = ?", compFlag).Find(&results).Error

	return
}

// GetBatchFromCompFlag 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromCompFlag(compFlags []int64) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`comp_flag` IN (?)", compFlags).Find(&results).Error

	return
}

// GetFromExecEnv 通过exec_env获取内容
func (obj *_AllRoutineHistoryMgr) GetFromExecEnv(execEnv string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`exec_env` = ?", execEnv).Find(&results).Error

	return
}

// GetBatchFromExecEnv 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromExecEnv(execEnvs []string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`exec_env` IN (?)", execEnvs).Find(&results).Error

	return
}

// GetFromRoutineBody 通过routine_body获取内容
func (obj *_AllRoutineHistoryMgr) GetFromRoutineBody(routineBody string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`routine_body` = ?", routineBody).Find(&results).Error

	return
}

// GetBatchFromRoutineBody 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromRoutineBody(routineBodys []string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`routine_body` IN (?)", routineBodys).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllRoutineHistoryMgr) GetFromComment(comment string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromComment(comments []string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromRouteSQL 通过route_sql获取内容
func (obj *_AllRoutineHistoryMgr) GetFromRouteSQL(routeSQL string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`route_sql` = ?", routeSQL).Find(&results).Error

	return
}

// GetBatchFromRouteSQL 批量查找
func (obj *_AllRoutineHistoryMgr) GetBatchFromRouteSQL(routeSQLs []string) (results []*AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`route_sql` IN (?)", routeSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllRoutineHistoryMgr) FetchByPrimaryKey(tenantID int64, routineID int64, schemaVersion int64) (result AllRoutineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRoutineHistory{}).Where("`tenant_id` = ? AND `routine_id` = ? AND `schema_version` = ?", tenantID, routineID, schemaVersion).First(&result).Error

	return
}
