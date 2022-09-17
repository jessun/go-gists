package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantTriggerMgr struct {
	*_BaseMgr
}

// AllTenantTriggerMgr open func
func AllTenantTriggerMgr(db *gorm.DB) *_AllTenantTriggerMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantTriggerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantTriggerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_trigger"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantTriggerMgr) GetTableName() string {
	return "__all_tenant_trigger"
}

// Reset 重置gorm会话
func (obj *_AllTenantTriggerMgr) Reset() *_AllTenantTriggerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantTriggerMgr) Get() (result AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantTriggerMgr) Gets() (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantTriggerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantTriggerMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantTriggerMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantTriggerMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTriggerID trigger_id获取
func (obj *_AllTenantTriggerMgr) WithTriggerID(triggerID int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_id"] = triggerID })
}

// WithTriggerName trigger_name获取
func (obj *_AllTenantTriggerMgr) WithTriggerName(triggerName string) Option {
	return optionFunc(func(o *options) { o.query["trigger_name"] = triggerName })
}

// WithDatabaseID database_id获取
func (obj *_AllTenantTriggerMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithOwnerID owner_id获取
func (obj *_AllTenantTriggerMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantTriggerMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithTriggerType trigger_type获取
func (obj *_AllTenantTriggerMgr) WithTriggerType(triggerType int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_type"] = triggerType })
}

// WithTriggerEvents trigger_events获取
func (obj *_AllTenantTriggerMgr) WithTriggerEvents(triggerEvents int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_events"] = triggerEvents })
}

// WithTimingPoints timing_points获取
func (obj *_AllTenantTriggerMgr) WithTimingPoints(timingPoints int64) Option {
	return optionFunc(func(o *options) { o.query["timing_points"] = timingPoints })
}

// WithBaseObjectType base_object_type获取
func (obj *_AllTenantTriggerMgr) WithBaseObjectType(baseObjectType int64) Option {
	return optionFunc(func(o *options) { o.query["base_object_type"] = baseObjectType })
}

// WithBaseObjectID base_object_id获取
func (obj *_AllTenantTriggerMgr) WithBaseObjectID(baseObjectID int64) Option {
	return optionFunc(func(o *options) { o.query["base_object_id"] = baseObjectID })
}

// WithTriggerFlags trigger_flags获取
func (obj *_AllTenantTriggerMgr) WithTriggerFlags(triggerFlags int64) Option {
	return optionFunc(func(o *options) { o.query["trigger_flags"] = triggerFlags })
}

// WithUpdateColumns update_columns获取
func (obj *_AllTenantTriggerMgr) WithUpdateColumns(updateColumns string) Option {
	return optionFunc(func(o *options) { o.query["update_columns"] = updateColumns })
}

// WithRefOldName ref_old_name获取
func (obj *_AllTenantTriggerMgr) WithRefOldName(refOldName string) Option {
	return optionFunc(func(o *options) { o.query["ref_old_name"] = refOldName })
}

// WithRefNewName ref_new_name获取
func (obj *_AllTenantTriggerMgr) WithRefNewName(refNewName string) Option {
	return optionFunc(func(o *options) { o.query["ref_new_name"] = refNewName })
}

// WithRefParentName ref_parent_name获取
func (obj *_AllTenantTriggerMgr) WithRefParentName(refParentName string) Option {
	return optionFunc(func(o *options) { o.query["ref_parent_name"] = refParentName })
}

// WithWhenCondition when_condition获取
func (obj *_AllTenantTriggerMgr) WithWhenCondition(whenCondition string) Option {
	return optionFunc(func(o *options) { o.query["when_condition"] = whenCondition })
}

// WithTriggerBody trigger_body获取
func (obj *_AllTenantTriggerMgr) WithTriggerBody(triggerBody string) Option {
	return optionFunc(func(o *options) { o.query["trigger_body"] = triggerBody })
}

// WithPackageSpecSource package_spec_source获取
func (obj *_AllTenantTriggerMgr) WithPackageSpecSource(packageSpecSource string) Option {
	return optionFunc(func(o *options) { o.query["package_spec_source"] = packageSpecSource })
}

// WithPackageBodySource package_body_source获取
func (obj *_AllTenantTriggerMgr) WithPackageBodySource(packageBodySource string) Option {
	return optionFunc(func(o *options) { o.query["package_body_source"] = packageBodySource })
}

// WithPackageFlag package_flag获取
func (obj *_AllTenantTriggerMgr) WithPackageFlag(packageFlag int64) Option {
	return optionFunc(func(o *options) { o.query["package_flag"] = packageFlag })
}

// WithPackageCompFlag package_comp_flag获取
func (obj *_AllTenantTriggerMgr) WithPackageCompFlag(packageCompFlag int64) Option {
	return optionFunc(func(o *options) { o.query["package_comp_flag"] = packageCompFlag })
}

// WithPackageExecEnv package_exec_env获取
func (obj *_AllTenantTriggerMgr) WithPackageExecEnv(packageExecEnv string) Option {
	return optionFunc(func(o *options) { o.query["package_exec_env"] = packageExecEnv })
}

// WithSQLMode sql_mode获取
func (obj *_AllTenantTriggerMgr) WithSQLMode(sqlMode int64) Option {
	return optionFunc(func(o *options) { o.query["sql_mode"] = sqlMode })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantTriggerMgr) GetByOption(opts ...Option) (result AllTenantTrigger, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantTriggerMgr) GetByOptions(opts ...Option) (results []*AllTenantTrigger, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantTriggerMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantTrigger, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where(options.query)
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
func (obj *_AllTenantTriggerMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantTriggerMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantTriggerMgr) GetFromTenantID(tenantID int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTriggerID 通过trigger_id获取内容
func (obj *_AllTenantTriggerMgr) GetFromTriggerID(triggerID int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_id` = ?", triggerID).Find(&results).Error

	return
}

// GetBatchFromTriggerID 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromTriggerID(triggerIDs []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_id` IN (?)", triggerIDs).Find(&results).Error

	return
}

// GetFromTriggerName 通过trigger_name获取内容
func (obj *_AllTenantTriggerMgr) GetFromTriggerName(triggerName string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_name` = ?", triggerName).Find(&results).Error

	return
}

// GetBatchFromTriggerName 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromTriggerName(triggerNames []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_name` IN (?)", triggerNames).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTenantTriggerMgr) GetFromDatabaseID(databaseID int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllTenantTriggerMgr) GetFromOwnerID(ownerID int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantTriggerMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromTriggerType 通过trigger_type获取内容
func (obj *_AllTenantTriggerMgr) GetFromTriggerType(triggerType int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_type` = ?", triggerType).Find(&results).Error

	return
}

// GetBatchFromTriggerType 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromTriggerType(triggerTypes []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_type` IN (?)", triggerTypes).Find(&results).Error

	return
}

// GetFromTriggerEvents 通过trigger_events获取内容
func (obj *_AllTenantTriggerMgr) GetFromTriggerEvents(triggerEvents int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_events` = ?", triggerEvents).Find(&results).Error

	return
}

// GetBatchFromTriggerEvents 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromTriggerEvents(triggerEventss []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_events` IN (?)", triggerEventss).Find(&results).Error

	return
}

// GetFromTimingPoints 通过timing_points获取内容
func (obj *_AllTenantTriggerMgr) GetFromTimingPoints(timingPoints int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`timing_points` = ?", timingPoints).Find(&results).Error

	return
}

// GetBatchFromTimingPoints 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromTimingPoints(timingPointss []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`timing_points` IN (?)", timingPointss).Find(&results).Error

	return
}

// GetFromBaseObjectType 通过base_object_type获取内容
func (obj *_AllTenantTriggerMgr) GetFromBaseObjectType(baseObjectType int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`base_object_type` = ?", baseObjectType).Find(&results).Error

	return
}

// GetBatchFromBaseObjectType 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromBaseObjectType(baseObjectTypes []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`base_object_type` IN (?)", baseObjectTypes).Find(&results).Error

	return
}

// GetFromBaseObjectID 通过base_object_id获取内容
func (obj *_AllTenantTriggerMgr) GetFromBaseObjectID(baseObjectID int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`base_object_id` = ?", baseObjectID).Find(&results).Error

	return
}

// GetBatchFromBaseObjectID 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromBaseObjectID(baseObjectIDs []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`base_object_id` IN (?)", baseObjectIDs).Find(&results).Error

	return
}

// GetFromTriggerFlags 通过trigger_flags获取内容
func (obj *_AllTenantTriggerMgr) GetFromTriggerFlags(triggerFlags int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_flags` = ?", triggerFlags).Find(&results).Error

	return
}

// GetBatchFromTriggerFlags 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromTriggerFlags(triggerFlagss []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_flags` IN (?)", triggerFlagss).Find(&results).Error

	return
}

// GetFromUpdateColumns 通过update_columns获取内容
func (obj *_AllTenantTriggerMgr) GetFromUpdateColumns(updateColumns string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`update_columns` = ?", updateColumns).Find(&results).Error

	return
}

// GetBatchFromUpdateColumns 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromUpdateColumns(updateColumnss []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`update_columns` IN (?)", updateColumnss).Find(&results).Error

	return
}

// GetFromRefOldName 通过ref_old_name获取内容
func (obj *_AllTenantTriggerMgr) GetFromRefOldName(refOldName string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`ref_old_name` = ?", refOldName).Find(&results).Error

	return
}

// GetBatchFromRefOldName 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromRefOldName(refOldNames []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`ref_old_name` IN (?)", refOldNames).Find(&results).Error

	return
}

// GetFromRefNewName 通过ref_new_name获取内容
func (obj *_AllTenantTriggerMgr) GetFromRefNewName(refNewName string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`ref_new_name` = ?", refNewName).Find(&results).Error

	return
}

// GetBatchFromRefNewName 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromRefNewName(refNewNames []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`ref_new_name` IN (?)", refNewNames).Find(&results).Error

	return
}

// GetFromRefParentName 通过ref_parent_name获取内容
func (obj *_AllTenantTriggerMgr) GetFromRefParentName(refParentName string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`ref_parent_name` = ?", refParentName).Find(&results).Error

	return
}

// GetBatchFromRefParentName 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromRefParentName(refParentNames []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`ref_parent_name` IN (?)", refParentNames).Find(&results).Error

	return
}

// GetFromWhenCondition 通过when_condition获取内容
func (obj *_AllTenantTriggerMgr) GetFromWhenCondition(whenCondition string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`when_condition` = ?", whenCondition).Find(&results).Error

	return
}

// GetBatchFromWhenCondition 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromWhenCondition(whenConditions []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`when_condition` IN (?)", whenConditions).Find(&results).Error

	return
}

// GetFromTriggerBody 通过trigger_body获取内容
func (obj *_AllTenantTriggerMgr) GetFromTriggerBody(triggerBody string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_body` = ?", triggerBody).Find(&results).Error

	return
}

// GetBatchFromTriggerBody 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromTriggerBody(triggerBodys []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`trigger_body` IN (?)", triggerBodys).Find(&results).Error

	return
}

// GetFromPackageSpecSource 通过package_spec_source获取内容
func (obj *_AllTenantTriggerMgr) GetFromPackageSpecSource(packageSpecSource string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_spec_source` = ?", packageSpecSource).Find(&results).Error

	return
}

// GetBatchFromPackageSpecSource 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromPackageSpecSource(packageSpecSources []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_spec_source` IN (?)", packageSpecSources).Find(&results).Error

	return
}

// GetFromPackageBodySource 通过package_body_source获取内容
func (obj *_AllTenantTriggerMgr) GetFromPackageBodySource(packageBodySource string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_body_source` = ?", packageBodySource).Find(&results).Error

	return
}

// GetBatchFromPackageBodySource 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromPackageBodySource(packageBodySources []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_body_source` IN (?)", packageBodySources).Find(&results).Error

	return
}

// GetFromPackageFlag 通过package_flag获取内容
func (obj *_AllTenantTriggerMgr) GetFromPackageFlag(packageFlag int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_flag` = ?", packageFlag).Find(&results).Error

	return
}

// GetBatchFromPackageFlag 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromPackageFlag(packageFlags []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_flag` IN (?)", packageFlags).Find(&results).Error

	return
}

// GetFromPackageCompFlag 通过package_comp_flag获取内容
func (obj *_AllTenantTriggerMgr) GetFromPackageCompFlag(packageCompFlag int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_comp_flag` = ?", packageCompFlag).Find(&results).Error

	return
}

// GetBatchFromPackageCompFlag 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromPackageCompFlag(packageCompFlags []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_comp_flag` IN (?)", packageCompFlags).Find(&results).Error

	return
}

// GetFromPackageExecEnv 通过package_exec_env获取内容
func (obj *_AllTenantTriggerMgr) GetFromPackageExecEnv(packageExecEnv string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_exec_env` = ?", packageExecEnv).Find(&results).Error

	return
}

// GetBatchFromPackageExecEnv 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromPackageExecEnv(packageExecEnvs []string) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`package_exec_env` IN (?)", packageExecEnvs).Find(&results).Error

	return
}

// GetFromSQLMode 通过sql_mode获取内容
func (obj *_AllTenantTriggerMgr) GetFromSQLMode(sqlMode int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`sql_mode` = ?", sqlMode).Find(&results).Error

	return
}

// GetBatchFromSQLMode 批量查找
func (obj *_AllTenantTriggerMgr) GetBatchFromSQLMode(sqlModes []int64) (results []*AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`sql_mode` IN (?)", sqlModes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantTriggerMgr) FetchByPrimaryKey(tenantID int64, triggerID int64) (result AllTenantTrigger, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantTrigger{}).Where("`tenant_id` = ? AND `trigger_id` = ?", tenantID, triggerID).First(&result).Error

	return
}
