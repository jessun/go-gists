package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantTriggerHistoryMgr struct {
	*_BaseMgr
}

// AllTenantTriggerHistoryMgr open func
func AllTenantTriggerHistoryMgr(db *gorm.DB) *_AllTenantTriggerHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantTriggerHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantTriggerHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_trigger_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantTriggerHistoryMgr) GetTableName() string {
	return "__all_tenant_trigger_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantTriggerHistoryMgr) Reset() *_AllTenantTriggerHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantTriggerHistoryMgr) Get() (result AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantTriggerHistoryMgr) Gets() (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantTriggerHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantTriggerHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantTriggerHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantTriggerHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTriggerID trigger_id获取
func (obj *_AllTenantTriggerHistoryMgr) WithTriggerID(triggerID int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_id"] = triggerID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantTriggerHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantTriggerHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithTriggerName trigger_name获取
func (obj *_AllTenantTriggerHistoryMgr) WithTriggerName(triggerName string) Option {
	return optionFunc(func(o *options) { o.query["trigger_name"] = triggerName })
}

// WithDatabaseID database_id获取
func (obj *_AllTenantTriggerHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithOwnerID owner_id获取
func (obj *_AllTenantTriggerHistoryMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithTriggerType trigger_type获取
func (obj *_AllTenantTriggerHistoryMgr) WithTriggerType(triggerType int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_type"] = triggerType })
}

// WithTriggerEvents trigger_events获取
func (obj *_AllTenantTriggerHistoryMgr) WithTriggerEvents(triggerEvents int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_events"] = triggerEvents })
}

// WithTimingPoints timing_points获取
func (obj *_AllTenantTriggerHistoryMgr) WithTimingPoints(timingPoints int64) Option {
	return optionFunc(func(o *options) { o.query["timing_points"] = timingPoints })
}

// WithBaseObjectType base_object_type获取
func (obj *_AllTenantTriggerHistoryMgr) WithBaseObjectType(baseObjectType int64) Option {
	return optionFunc(func(o *options) { o.query["base_object_type"] = baseObjectType })
}

// WithBaseObjectID base_object_id获取
func (obj *_AllTenantTriggerHistoryMgr) WithBaseObjectID(baseObjectID int64) Option {
	return optionFunc(func(o *options) { o.query["base_object_id"] = baseObjectID })
}

// WithTriggerFlags trigger_flags获取
func (obj *_AllTenantTriggerHistoryMgr) WithTriggerFlags(triggerFlags int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_flags"] = triggerFlags })
}

// WithUpdateColumns update_columns获取
func (obj *_AllTenantTriggerHistoryMgr) WithUpdateColumns(updateColumns string) Option {
	return optionFunc(func(o *options) { o.query["update_columns"] = updateColumns })
}

// WithRefOldName ref_old_name获取
func (obj *_AllTenantTriggerHistoryMgr) WithRefOldName(refOldName string) Option {
	return optionFunc(func(o *options) { o.query["ref_old_name"] = refOldName })
}

// WithRefNewName ref_new_name获取
func (obj *_AllTenantTriggerHistoryMgr) WithRefNewName(refNewName string) Option {
	return optionFunc(func(o *options) { o.query["ref_new_name"] = refNewName })
}

// WithRefParentName ref_parent_name获取
func (obj *_AllTenantTriggerHistoryMgr) WithRefParentName(refParentName string) Option {
	return optionFunc(func(o *options) { o.query["ref_parent_name"] = refParentName })
}

// WithWhenCondition when_condition获取
func (obj *_AllTenantTriggerHistoryMgr) WithWhenCondition(whenCondition string) Option {
	return optionFunc(func(o *options) { o.query["when_condition"] = whenCondition })
}

// WithTriggerBody trigger_body获取
func (obj *_AllTenantTriggerHistoryMgr) WithTriggerBody(triggerBody string) Option {
	return optionFunc(func(o *options) { o.query["trigger_body"] = triggerBody })
}

// WithPackageSpecSource package_spec_source获取
func (obj *_AllTenantTriggerHistoryMgr) WithPackageSpecSource(packageSpecSource string) Option {
	return optionFunc(func(o *options) { o.query["package_spec_source"] = packageSpecSource })
}

// WithPackageBodySource package_body_source获取
func (obj *_AllTenantTriggerHistoryMgr) WithPackageBodySource(packageBodySource string) Option {
	return optionFunc(func(o *options) { o.query["package_body_source"] = packageBodySource })
}

// WithPackageFlag package_flag获取
func (obj *_AllTenantTriggerHistoryMgr) WithPackageFlag(packageFlag int64) Option {
	return optionFunc(func(o *options) { o.query["package_flag"] = packageFlag })
}

// WithPackageCompFlag package_comp_flag获取
func (obj *_AllTenantTriggerHistoryMgr) WithPackageCompFlag(packageCompFlag int64) Option {
	return optionFunc(func(o *options) { o.query["package_comp_flag"] = packageCompFlag })
}

// WithPackageExecEnv package_exec_env获取
func (obj *_AllTenantTriggerHistoryMgr) WithPackageExecEnv(packageExecEnv string) Option {
	return optionFunc(func(o *options) { o.query["package_exec_env"] = packageExecEnv })
}

// WithSQLMode sql_mode获取
func (obj *_AllTenantTriggerHistoryMgr) WithSQLMode(sqlMode int64) Option {
	return optionFunc(func(o *options) { o.query["sql_mode"] = sqlMode })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantTriggerHistoryMgr) GetByOption(opts ...Option) (result AllTenantTriggerHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantTriggerHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantTriggerHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantTriggerHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantTriggerHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where(options.query)
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
func (obj *_AllTenantTriggerHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTriggerID 通过trigger_id获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromTriggerID(triggerID int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_id` = ?", triggerID).Find(&results).Error

	return
}

// GetBatchFromTriggerID 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromTriggerID(triggerIDs []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_id` IN (?)", triggerIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromTriggerName 通过trigger_name获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromTriggerName(triggerName string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_name` = ?", triggerName).Find(&results).Error

	return
}

// GetBatchFromTriggerName 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromTriggerName(triggerNames []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_name` IN (?)", triggerNames).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromOwnerID(ownerID int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromTriggerType 通过trigger_type获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromTriggerType(triggerType int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_type` = ?", triggerType).Find(&results).Error

	return
}

// GetBatchFromTriggerType 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromTriggerType(triggerTypes []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_type` IN (?)", triggerTypes).Find(&results).Error

	return
}

// GetFromTriggerEvents 通过trigger_events获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromTriggerEvents(triggerEvents int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_events` = ?", triggerEvents).Find(&results).Error

	return
}

// GetBatchFromTriggerEvents 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromTriggerEvents(triggerEventss []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_events` IN (?)", triggerEventss).Find(&results).Error

	return
}

// GetFromTimingPoints 通过timing_points获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromTimingPoints(timingPoints int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`timing_points` = ?", timingPoints).Find(&results).Error

	return
}

// GetBatchFromTimingPoints 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromTimingPoints(timingPointss []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`timing_points` IN (?)", timingPointss).Find(&results).Error

	return
}

// GetFromBaseObjectType 通过base_object_type获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromBaseObjectType(baseObjectType int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`base_object_type` = ?", baseObjectType).Find(&results).Error

	return
}

// GetBatchFromBaseObjectType 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromBaseObjectType(baseObjectTypes []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`base_object_type` IN (?)", baseObjectTypes).Find(&results).Error

	return
}

// GetFromBaseObjectID 通过base_object_id获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromBaseObjectID(baseObjectID int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`base_object_id` = ?", baseObjectID).Find(&results).Error

	return
}

// GetBatchFromBaseObjectID 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromBaseObjectID(baseObjectIDs []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`base_object_id` IN (?)", baseObjectIDs).Find(&results).Error

	return
}

// GetFromTriggerFlags 通过trigger_flags获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromTriggerFlags(triggerFlags int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_flags` = ?", triggerFlags).Find(&results).Error

	return
}

// GetBatchFromTriggerFlags 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromTriggerFlags(triggerFlagss []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_flags` IN (?)", triggerFlagss).Find(&results).Error

	return
}

// GetFromUpdateColumns 通过update_columns获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromUpdateColumns(updateColumns string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`update_columns` = ?", updateColumns).Find(&results).Error

	return
}

// GetBatchFromUpdateColumns 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromUpdateColumns(updateColumnss []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`update_columns` IN (?)", updateColumnss).Find(&results).Error

	return
}

// GetFromRefOldName 通过ref_old_name获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromRefOldName(refOldName string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`ref_old_name` = ?", refOldName).Find(&results).Error

	return
}

// GetBatchFromRefOldName 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromRefOldName(refOldNames []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`ref_old_name` IN (?)", refOldNames).Find(&results).Error

	return
}

// GetFromRefNewName 通过ref_new_name获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromRefNewName(refNewName string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`ref_new_name` = ?", refNewName).Find(&results).Error

	return
}

// GetBatchFromRefNewName 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromRefNewName(refNewNames []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`ref_new_name` IN (?)", refNewNames).Find(&results).Error

	return
}

// GetFromRefParentName 通过ref_parent_name获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromRefParentName(refParentName string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`ref_parent_name` = ?", refParentName).Find(&results).Error

	return
}

// GetBatchFromRefParentName 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromRefParentName(refParentNames []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`ref_parent_name` IN (?)", refParentNames).Find(&results).Error

	return
}

// GetFromWhenCondition 通过when_condition获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromWhenCondition(whenCondition string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`when_condition` = ?", whenCondition).Find(&results).Error

	return
}

// GetBatchFromWhenCondition 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromWhenCondition(whenConditions []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`when_condition` IN (?)", whenConditions).Find(&results).Error

	return
}

// GetFromTriggerBody 通过trigger_body获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromTriggerBody(triggerBody string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_body` = ?", triggerBody).Find(&results).Error

	return
}

// GetBatchFromTriggerBody 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromTriggerBody(triggerBodys []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`trigger_body` IN (?)", triggerBodys).Find(&results).Error

	return
}

// GetFromPackageSpecSource 通过package_spec_source获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromPackageSpecSource(packageSpecSource string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_spec_source` = ?", packageSpecSource).Find(&results).Error

	return
}

// GetBatchFromPackageSpecSource 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromPackageSpecSource(packageSpecSources []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_spec_source` IN (?)", packageSpecSources).Find(&results).Error

	return
}

// GetFromPackageBodySource 通过package_body_source获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromPackageBodySource(packageBodySource string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_body_source` = ?", packageBodySource).Find(&results).Error

	return
}

// GetBatchFromPackageBodySource 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromPackageBodySource(packageBodySources []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_body_source` IN (?)", packageBodySources).Find(&results).Error

	return
}

// GetFromPackageFlag 通过package_flag获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromPackageFlag(packageFlag int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_flag` = ?", packageFlag).Find(&results).Error

	return
}

// GetBatchFromPackageFlag 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromPackageFlag(packageFlags []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_flag` IN (?)", packageFlags).Find(&results).Error

	return
}

// GetFromPackageCompFlag 通过package_comp_flag获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromPackageCompFlag(packageCompFlag int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_comp_flag` = ?", packageCompFlag).Find(&results).Error

	return
}

// GetBatchFromPackageCompFlag 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromPackageCompFlag(packageCompFlags []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_comp_flag` IN (?)", packageCompFlags).Find(&results).Error

	return
}

// GetFromPackageExecEnv 通过package_exec_env获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromPackageExecEnv(packageExecEnv string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_exec_env` = ?", packageExecEnv).Find(&results).Error

	return
}

// GetBatchFromPackageExecEnv 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromPackageExecEnv(packageExecEnvs []string) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`package_exec_env` IN (?)", packageExecEnvs).Find(&results).Error

	return
}

// GetFromSQLMode 通过sql_mode获取内容
func (obj *_AllTenantTriggerHistoryMgr) GetFromSQLMode(sqlMode int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`sql_mode` = ?", sqlMode).Find(&results).Error

	return
}

// GetBatchFromSQLMode 批量查找
func (obj *_AllTenantTriggerHistoryMgr) GetBatchFromSQLMode(sqlModes []int64) (results []*AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`sql_mode` IN (?)", sqlModes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantTriggerHistoryMgr) FetchByPrimaryKey(tenantID int64, triggerID int64, schemaVersion int64) (result AllTenantTriggerHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTriggerHistory{}).Where("`tenant_id` = ? AND `trigger_id` = ? AND `schema_version` = ?", tenantID, triggerID, schemaVersion).First(&result).Error

	return
}
