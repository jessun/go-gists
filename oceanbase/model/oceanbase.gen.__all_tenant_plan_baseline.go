package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantPlanBaselineMgr struct {
	*_BaseMgr
}

// AllTenantPlanBaselineMgr open func
func AllTenantPlanBaselineMgr(db *gorm.DB) *_AllTenantPlanBaselineMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantPlanBaselineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantPlanBaselineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_plan_baseline"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantPlanBaselineMgr) GetTableName() string {
	return "__all_tenant_plan_baseline"
}

// Reset 重置gorm会话
func (obj *_AllTenantPlanBaselineMgr) Reset() *_AllTenantPlanBaselineMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantPlanBaselineMgr) Get() (result AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantPlanBaselineMgr) Gets() (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantPlanBaselineMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantPlanBaselineMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantPlanBaselineMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantPlanBaselineMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPlanBaselineID plan_baseline_id获取
func (obj *_AllTenantPlanBaselineMgr) WithPlanBaselineID(planBaselineID int64) Option {
	return optionFunc(func(o *options) { o.query["plan_baseline_id"] = planBaselineID })
}

// WithDatabaseID database_id获取
func (obj *_AllTenantPlanBaselineMgr) WithDatabaseID(databaseID uint64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantPlanBaselineMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithSQLID sql_id获取
func (obj *_AllTenantPlanBaselineMgr) WithSQLID(sqlID []byte) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithPlanHashValue plan_hash_value获取
func (obj *_AllTenantPlanBaselineMgr) WithPlanHashValue(planHashValue uint64) Option {
	return optionFunc(func(o *options) { o.query["plan_hash_value"] = planHashValue })
}

// WithParamsInfo params_info获取
func (obj *_AllTenantPlanBaselineMgr) WithParamsInfo(paramsInfo []byte) Option {
	return optionFunc(func(o *options) { o.query["params_info"] = paramsInfo })
}

// WithOutlineData outline_data获取
func (obj *_AllTenantPlanBaselineMgr) WithOutlineData(outlineData string) Option {
	return optionFunc(func(o *options) { o.query["outline_data"] = outlineData })
}

// WithSQLText sql_text获取
func (obj *_AllTenantPlanBaselineMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithFixed fixed获取
func (obj *_AllTenantPlanBaselineMgr) WithFixed(fixed int8) Option {
	return optionFunc(func(o *options) { o.query["fixed"] = fixed })
}

// WithEnabled enabled获取
func (obj *_AllTenantPlanBaselineMgr) WithEnabled(enabled int8) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithHintsInfo hints_info获取
func (obj *_AllTenantPlanBaselineMgr) WithHintsInfo(hintsInfo string) Option {
	return optionFunc(func(o *options) { o.query["hints_info"] = hintsInfo })
}

// WithHintsAllWorked hints_all_worked获取
func (obj *_AllTenantPlanBaselineMgr) WithHintsAllWorked(hintsAllWorked int8) Option {
	return optionFunc(func(o *options) { o.query["hints_all_worked"] = hintsAllWorked })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantPlanBaselineMgr) GetByOption(opts ...Option) (result AllTenantPlanBaseline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantPlanBaselineMgr) GetByOptions(opts ...Option) (results []*AllTenantPlanBaseline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantPlanBaselineMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantPlanBaseline, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where(options.query)
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
func (obj *_AllTenantPlanBaselineMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromTenantID(tenantID int64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPlanBaselineID 通过plan_baseline_id获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromPlanBaselineID(planBaselineID int64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`plan_baseline_id` = ?", planBaselineID).Find(&results).Error

	return
}

// GetBatchFromPlanBaselineID 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromPlanBaselineID(planBaselineIDs []int64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`plan_baseline_id` IN (?)", planBaselineIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromDatabaseID(databaseID uint64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromDatabaseID(databaseIDs []uint64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromSQLID(sqlID []byte) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromSQLID(sqlIDs [][]byte) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromPlanHashValue 通过plan_hash_value获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromPlanHashValue(planHashValue uint64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`plan_hash_value` = ?", planHashValue).Find(&results).Error

	return
}

// GetBatchFromPlanHashValue 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromPlanHashValue(planHashValues []uint64) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`plan_hash_value` IN (?)", planHashValues).Find(&results).Error

	return
}

// GetFromParamsInfo 通过params_info获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromParamsInfo(paramsInfo []byte) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`params_info` = ?", paramsInfo).Find(&results).Error

	return
}

// GetBatchFromParamsInfo 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromParamsInfo(paramsInfos [][]byte) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`params_info` IN (?)", paramsInfos).Find(&results).Error

	return
}

// GetFromOutlineData 通过outline_data获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromOutlineData(outlineData string) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`outline_data` = ?", outlineData).Find(&results).Error

	return
}

// GetBatchFromOutlineData 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromOutlineData(outlineDatas []string) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`outline_data` IN (?)", outlineDatas).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromSQLText(sqlText string) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromFixed 通过fixed获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromFixed(fixed int8) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`fixed` = ?", fixed).Find(&results).Error

	return
}

// GetBatchFromFixed 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromFixed(fixeds []int8) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`fixed` IN (?)", fixeds).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromEnabled(enabled int8) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`enabled` = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromEnabled(enableds []int8) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`enabled` IN (?)", enableds).Find(&results).Error

	return
}

// GetFromHintsInfo 通过hints_info获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromHintsInfo(hintsInfo string) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`hints_info` = ?", hintsInfo).Find(&results).Error

	return
}

// GetBatchFromHintsInfo 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromHintsInfo(hintsInfos []string) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`hints_info` IN (?)", hintsInfos).Find(&results).Error

	return
}

// GetFromHintsAllWorked 通过hints_all_worked获取内容
func (obj *_AllTenantPlanBaselineMgr) GetFromHintsAllWorked(hintsAllWorked int8) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`hints_all_worked` = ?", hintsAllWorked).Find(&results).Error

	return
}

// GetBatchFromHintsAllWorked 批量查找
func (obj *_AllTenantPlanBaselineMgr) GetBatchFromHintsAllWorked(hintsAllWorkeds []int8) (results []*AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`hints_all_worked` IN (?)", hintsAllWorkeds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantPlanBaselineMgr) FetchByPrimaryKey(tenantID int64, planBaselineID int64) (result AllTenantPlanBaseline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPlanBaseline{}).Where("`tenant_id` = ? AND `plan_baseline_id` = ?", tenantID, planBaselineID).First(&result).Error

	return
}
