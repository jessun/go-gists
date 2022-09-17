package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantPlanBaselineHistoryMgr struct {
	*_BaseMgr
}

// AllTenantPlanBaselineHistoryMgr open func
func AllTenantPlanBaselineHistoryMgr(db *gorm.DB) *_AllTenantPlanBaselineHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantPlanBaselineHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantPlanBaselineHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_plan_baseline_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantPlanBaselineHistoryMgr) GetTableName() string {
	return "__all_tenant_plan_baseline_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantPlanBaselineHistoryMgr) Reset() *_AllTenantPlanBaselineHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantPlanBaselineHistoryMgr) Get() (result AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantPlanBaselineHistoryMgr) Gets() (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantPlanBaselineHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPlanBaselineID plan_baseline_id获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithPlanBaselineID(planBaselineID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_baseline_id"] = planBaselineID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithDatabaseID(databaseID uint64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSQLID sql_id获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithSQLID(sqlID []byte) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithPlanHashValue plan_hash_value获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithPlanHashValue(planHashValue uint64) Option {
	return optionFunc(func(o *options) { o.query["plan_hash_value"] = planHashValue })
}

// WithParamsInfo params_info获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithParamsInfo(paramsInfo []byte) Option {
	return optionFunc(func(o *options) { o.query["params_info"] = paramsInfo })
}

// WithOutlineData outline_data获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithOutlineData(outlineData string) Option {
	return optionFunc(func(o *options) { o.query["outline_data"] = outlineData })
}

// WithSQLText sql_text获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithFixed fixed获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithFixed(fixed int8) Option {
	return optionFunc(func(o *options) { o.query["fixed"] = fixed })
}

// WithEnabled enabled获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithEnabled(enabled int8) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithHintsInfo hints_info获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithHintsInfo(hintsInfo string) Option {
	return optionFunc(func(o *options) { o.query["hints_info"] = hintsInfo })
}

// WithHintsAllWorked hints_all_worked获取
func (obj *_AllTenantPlanBaselineHistoryMgr) WithHintsAllWorked(hintsAllWorked int8) Option {
	return optionFunc(func(o *options) { o.query["hints_all_worked"] = hintsAllWorked })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantPlanBaselineHistoryMgr) GetByOption(opts ...Option) (result AllTenantPlanBaselineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantPlanBaselineHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantPlanBaselineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantPlanBaselineHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantPlanBaselineHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where(options.query)
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
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPlanBaselineID 通过plan_baseline_id获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromPlanBaselineID(planBaselineID int64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`plan_baseline_id` = ?", planBaselineID).Find(&results).Error

	return
}

// GetBatchFromPlanBaselineID 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromPlanBaselineID(planBaselineIDs []int64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`plan_baseline_id` IN (?)", planBaselineIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromDatabaseID(databaseID uint64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromDatabaseID(databaseIDs []uint64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromSQLID(sqlID []byte) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromSQLID(sqlIDs [][]byte) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromPlanHashValue 通过plan_hash_value获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromPlanHashValue(planHashValue uint64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`plan_hash_value` = ?", planHashValue).Find(&results).Error

	return
}

// GetBatchFromPlanHashValue 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromPlanHashValue(planHashValues []uint64) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`plan_hash_value` IN (?)", planHashValues).Find(&results).Error

	return
}

// GetFromParamsInfo 通过params_info获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromParamsInfo(paramsInfo []byte) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`params_info` = ?", paramsInfo).Find(&results).Error

	return
}

// GetBatchFromParamsInfo 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromParamsInfo(paramsInfos [][]byte) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`params_info` IN (?)", paramsInfos).Find(&results).Error

	return
}

// GetFromOutlineData 通过outline_data获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromOutlineData(outlineData string) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`outline_data` = ?", outlineData).Find(&results).Error

	return
}

// GetBatchFromOutlineData 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromOutlineData(outlineDatas []string) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`outline_data` IN (?)", outlineDatas).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromSQLText(sqlText string) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromFixed 通过fixed获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromFixed(fixed int8) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`fixed` = ?", fixed).Find(&results).Error

	return
}

// GetBatchFromFixed 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromFixed(fixeds []int8) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`fixed` IN (?)", fixeds).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromEnabled(enabled int8) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`enabled` = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromEnabled(enableds []int8) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`enabled` IN (?)", enableds).Find(&results).Error

	return
}

// GetFromHintsInfo 通过hints_info获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromHintsInfo(hintsInfo string) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`hints_info` = ?", hintsInfo).Find(&results).Error

	return
}

// GetBatchFromHintsInfo 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromHintsInfo(hintsInfos []string) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`hints_info` IN (?)", hintsInfos).Find(&results).Error

	return
}

// GetFromHintsAllWorked 通过hints_all_worked获取内容
func (obj *_AllTenantPlanBaselineHistoryMgr) GetFromHintsAllWorked(hintsAllWorked int8) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`hints_all_worked` = ?", hintsAllWorked).Find(&results).Error

	return
}

// GetBatchFromHintsAllWorked 批量查找
func (obj *_AllTenantPlanBaselineHistoryMgr) GetBatchFromHintsAllWorked(hintsAllWorkeds []int8) (results []*AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`hints_all_worked` IN (?)", hintsAllWorkeds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantPlanBaselineHistoryMgr) FetchByPrimaryKey(tenantID int64, planBaselineID int64, schemaVersion int64) (result AllTenantPlanBaselineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaselineHistory{}).Where("`tenant_id` = ? AND `plan_baseline_id` = ? AND `schema_version` = ?", tenantID, planBaselineID, schemaVersion).First(&result).Error

	return
}
