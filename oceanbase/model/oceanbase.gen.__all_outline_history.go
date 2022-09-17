package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllOutlineHistoryMgr struct {
	*_BaseMgr
}

// AllOutlineHistoryMgr open func
func AllOutlineHistoryMgr(db *gorm.DB) *_AllOutlineHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllOutlineHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllOutlineHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_outline_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllOutlineHistoryMgr) GetTableName() string {
	return "__all_outline_history"
}

// Reset 重置gorm会话
func (obj *_AllOutlineHistoryMgr) Reset() *_AllOutlineHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllOutlineHistoryMgr) Get() (result AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllOutlineHistoryMgr) Gets() (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllOutlineHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllOutlineHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllOutlineHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllOutlineHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithOutlineID outline_id获取
func (obj *_AllOutlineHistoryMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllOutlineHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllOutlineHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllOutlineHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithName name获取
func (obj *_AllOutlineHistoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSignature signature获取
func (obj *_AllOutlineHistoryMgr) WithSignature(signature []byte) Option {
	return optionFunc(func(o *options) { o.query["signature"] = signature })
}

// WithOutlineContent outline_content获取
func (obj *_AllOutlineHistoryMgr) WithOutlineContent(outlineContent string) Option {
	return optionFunc(func(o *options) { o.query["outline_content"] = outlineContent })
}

// WithSQLText sql_text获取
func (obj *_AllOutlineHistoryMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithOwner owner获取
func (obj *_AllOutlineHistoryMgr) WithOwner(owner string) Option {
	return optionFunc(func(o *options) { o.query["owner"] = owner })
}

// WithUsed used获取
func (obj *_AllOutlineHistoryMgr) WithUsed(used int64) Option {
	return optionFunc(func(o *options) { o.query["used"] = used })
}

// WithVersion version获取
func (obj *_AllOutlineHistoryMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithCompatible compatible获取
func (obj *_AllOutlineHistoryMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithEnabled enabled获取
func (obj *_AllOutlineHistoryMgr) WithEnabled(enabled int64) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithFormat format获取
func (obj *_AllOutlineHistoryMgr) WithFormat(format int64) Option {
	return optionFunc(func(o *options) { o.query["format"] = format })
}

// WithOutlineParams outline_params获取
func (obj *_AllOutlineHistoryMgr) WithOutlineParams(outlineParams []byte) Option {
	return optionFunc(func(o *options) { o.query["outline_params"] = outlineParams })
}

// WithOutlineTarget outline_target获取
func (obj *_AllOutlineHistoryMgr) WithOutlineTarget(outlineTarget string) Option {
	return optionFunc(func(o *options) { o.query["outline_target"] = outlineTarget })
}

// WithSQLID sql_id获取
func (obj *_AllOutlineHistoryMgr) WithSQLID(sqlID []byte) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithOwnerID owner_id获取
func (obj *_AllOutlineHistoryMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// GetByOption 功能选项模式获取
func (obj *_AllOutlineHistoryMgr) GetByOption(opts ...Option) (result AllOutlineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllOutlineHistoryMgr) GetByOptions(opts ...Option) (results []*AllOutlineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllOutlineHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllOutlineHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where(options.query)
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
func (obj *_AllOutlineHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllOutlineHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllOutlineHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromOutlineID 通过outline_id获取内容
func (obj *_AllOutlineHistoryMgr) GetFromOutlineID(outlineID int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`outline_id` = ?", outlineID).Find(&results).Error

	return
}

// GetBatchFromOutlineID 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllOutlineHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllOutlineHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllOutlineHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllOutlineHistoryMgr) GetFromName(name string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromName(names []string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSignature 通过signature获取内容
func (obj *_AllOutlineHistoryMgr) GetFromSignature(signature []byte) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`signature` = ?", signature).Find(&results).Error

	return
}

// GetBatchFromSignature 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromSignature(signatures [][]byte) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`signature` IN (?)", signatures).Find(&results).Error

	return
}

// GetFromOutlineContent 通过outline_content获取内容
func (obj *_AllOutlineHistoryMgr) GetFromOutlineContent(outlineContent string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`outline_content` = ?", outlineContent).Find(&results).Error

	return
}

// GetBatchFromOutlineContent 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromOutlineContent(outlineContents []string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`outline_content` IN (?)", outlineContents).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllOutlineHistoryMgr) GetFromSQLText(sqlText string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromOwner 通过owner获取内容
func (obj *_AllOutlineHistoryMgr) GetFromOwner(owner string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`owner` = ?", owner).Find(&results).Error

	return
}

// GetBatchFromOwner 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromOwner(owners []string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`owner` IN (?)", owners).Find(&results).Error

	return
}

// GetFromUsed 通过used获取内容
func (obj *_AllOutlineHistoryMgr) GetFromUsed(used int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`used` = ?", used).Find(&results).Error

	return
}

// GetBatchFromUsed 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromUsed(useds []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`used` IN (?)", useds).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllOutlineHistoryMgr) GetFromVersion(version string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromVersion(versions []string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllOutlineHistoryMgr) GetFromCompatible(compatible int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容
func (obj *_AllOutlineHistoryMgr) GetFromEnabled(enabled int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`enabled` = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromEnabled(enableds []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`enabled` IN (?)", enableds).Find(&results).Error

	return
}

// GetFromFormat 通过format获取内容
func (obj *_AllOutlineHistoryMgr) GetFromFormat(format int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`format` = ?", format).Find(&results).Error

	return
}

// GetBatchFromFormat 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromFormat(formats []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`format` IN (?)", formats).Find(&results).Error

	return
}

// GetFromOutlineParams 通过outline_params获取内容
func (obj *_AllOutlineHistoryMgr) GetFromOutlineParams(outlineParams []byte) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`outline_params` = ?", outlineParams).Find(&results).Error

	return
}

// GetBatchFromOutlineParams 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromOutlineParams(outlineParamss [][]byte) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`outline_params` IN (?)", outlineParamss).Find(&results).Error

	return
}

// GetFromOutlineTarget 通过outline_target获取内容
func (obj *_AllOutlineHistoryMgr) GetFromOutlineTarget(outlineTarget string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`outline_target` = ?", outlineTarget).Find(&results).Error

	return
}

// GetBatchFromOutlineTarget 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromOutlineTarget(outlineTargets []string) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`outline_target` IN (?)", outlineTargets).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllOutlineHistoryMgr) GetFromSQLID(sqlID []byte) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromSQLID(sqlIDs [][]byte) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllOutlineHistoryMgr) GetFromOwnerID(ownerID int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllOutlineHistoryMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllOutlineHistoryMgr) FetchByPrimaryKey(tenantID int64, outlineID int64, schemaVersion int64) (result AllOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutlineHistory{}).Where("`tenant_id` = ? AND `outline_id` = ? AND `schema_version` = ?", tenantID, outlineID, schemaVersion).First(&result).Error

	return
}
