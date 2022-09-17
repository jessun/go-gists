package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTenantPlanBaselineMgr struct {
	*_BaseMgr
}

// AllVirtualTenantPlanBaselineMgr open func
func AllVirtualTenantPlanBaselineMgr(db *gorm.DB) *_AllVirtualTenantPlanBaselineMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantPlanBaselineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantPlanBaselineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_plan_baseline"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantPlanBaselineMgr) GetTableName() string {
	return "__all_virtual_tenant_plan_baseline"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantPlanBaselineMgr) Reset() *_AllVirtualTenantPlanBaselineMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantPlanBaselineMgr) Get() (result AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantPlanBaselineMgr) Gets() (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantPlanBaselineMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPlanBaselineID plan_baseline_id获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithPlanBaselineID(planBaselineID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_baseline_id"] = planBaselineID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithDatabaseID(databaseID uint64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithSQLID sql_id获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithSQLID(sqlID []byte) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithPlanHashValue plan_hash_value获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithPlanHashValue(planHashValue uint64) Option {
	return optionFunc(func(o *options) { o.query["plan_hash_value"] = planHashValue })
}

// WithParamsInfo params_info获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithParamsInfo(paramsInfo []byte) Option {
	return optionFunc(func(o *options) { o.query["params_info"] = paramsInfo })
}

// WithOutlineData outline_data获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithOutlineData(outlineData string) Option {
	return optionFunc(func(o *options) { o.query["outline_data"] = outlineData })
}

// WithSQLText sql_text获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithFixed fixed获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithFixed(fixed int8) Option {
	return optionFunc(func(o *options) { o.query["fixed"] = fixed })
}

// WithEnabled enabled获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithEnabled(enabled int8) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithHintsInfo hints_info获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithHintsInfo(hintsInfo string) Option {
	return optionFunc(func(o *options) { o.query["hints_info"] = hintsInfo })
}

// WithHintsAllWorked hints_all_worked获取
func (obj *_AllVirtualTenantPlanBaselineMgr) WithHintsAllWorked(hintsAllWorked int8) Option {
	return optionFunc(func(o *options) { o.query["hints_all_worked"] = hintsAllWorked })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantPlanBaselineMgr) GetByOption(opts ...Option) (result AllVirtualTenantPlanBaseline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantPlanBaselineMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantPlanBaseline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantPlanBaselineMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantPlanBaseline, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where(options.query)
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
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPlanBaselineID 通过plan_baseline_id获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromPlanBaselineID(planBaselineID int64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`plan_baseline_id` = ?", planBaselineID).Find(&results).Error

	return
}

// GetBatchFromPlanBaselineID 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromPlanBaselineID(planBaselineIDs []int64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`plan_baseline_id` IN (?)", planBaselineIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromDatabaseID(databaseID uint64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromDatabaseID(databaseIDs []uint64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromSQLID(sqlID []byte) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromSQLID(sqlIDs [][]byte) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromPlanHashValue 通过plan_hash_value获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromPlanHashValue(planHashValue uint64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`plan_hash_value` = ?", planHashValue).Find(&results).Error

	return
}

// GetBatchFromPlanHashValue 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromPlanHashValue(planHashValues []uint64) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`plan_hash_value` IN (?)", planHashValues).Find(&results).Error

	return
}

// GetFromParamsInfo 通过params_info获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromParamsInfo(paramsInfo []byte) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`params_info` = ?", paramsInfo).Find(&results).Error

	return
}

// GetBatchFromParamsInfo 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromParamsInfo(paramsInfos [][]byte) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`params_info` IN (?)", paramsInfos).Find(&results).Error

	return
}

// GetFromOutlineData 通过outline_data获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromOutlineData(outlineData string) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`outline_data` = ?", outlineData).Find(&results).Error

	return
}

// GetBatchFromOutlineData 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromOutlineData(outlineDatas []string) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`outline_data` IN (?)", outlineDatas).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromSQLText(sqlText string) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromFixed 通过fixed获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromFixed(fixed int8) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`fixed` = ?", fixed).Find(&results).Error

	return
}

// GetBatchFromFixed 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromFixed(fixeds []int8) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`fixed` IN (?)", fixeds).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromEnabled(enabled int8) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`enabled` = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromEnabled(enableds []int8) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`enabled` IN (?)", enableds).Find(&results).Error

	return
}

// GetFromHintsInfo 通过hints_info获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromHintsInfo(hintsInfo string) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`hints_info` = ?", hintsInfo).Find(&results).Error

	return
}

// GetBatchFromHintsInfo 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromHintsInfo(hintsInfos []string) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`hints_info` IN (?)", hintsInfos).Find(&results).Error

	return
}

// GetFromHintsAllWorked 通过hints_all_worked获取内容
func (obj *_AllVirtualTenantPlanBaselineMgr) GetFromHintsAllWorked(hintsAllWorked int8) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`hints_all_worked` = ?", hintsAllWorked).Find(&results).Error

	return
}

// GetBatchFromHintsAllWorked 批量查找
func (obj *_AllVirtualTenantPlanBaselineMgr) GetBatchFromHintsAllWorked(hintsAllWorkeds []int8) (results []*AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`hints_all_worked` IN (?)", hintsAllWorkeds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantPlanBaselineMgr) FetchByPrimaryKey(tenantID int64, planBaselineID int64) (result AllVirtualTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPlanBaseline{}).Where("`tenant_id` = ? AND `plan_baseline_id` = ?", tenantID, planBaselineID).First(&result).Error

	return
}
