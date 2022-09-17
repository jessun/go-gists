package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllOutlineMgr struct {
	*_BaseMgr
}

// AllOutlineMgr open func
func AllOutlineMgr(db *gorm.DB) *_AllOutlineMgr {
	if db == nil {
		panic(fmt.Errorf("AllOutlineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllOutlineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_outline"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllOutlineMgr) GetTableName() string {
	return "__all_outline"
}

// Reset 重置gorm会话
func (obj *_AllOutlineMgr) Reset() *_AllOutlineMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllOutlineMgr) Get() (result AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllOutlineMgr) Gets() (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllOutlineMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllOutlineMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllOutlineMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllOutlineMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithOutlineID outline_id获取
func (obj *_AllOutlineMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithDatabaseID database_id获取
func (obj *_AllOutlineMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllOutlineMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithName name获取
func (obj *_AllOutlineMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSignature signature获取
func (obj *_AllOutlineMgr) WithSignature(signature []byte) Option {
	return optionFunc(func(o *options) { o.query["signature"] = signature })
}

// WithOutlineContent outline_content获取
func (obj *_AllOutlineMgr) WithOutlineContent(outlineContent string) Option {
	return optionFunc(func(o *options) { o.query["outline_content"] = outlineContent })
}

// WithSQLText sql_text获取
func (obj *_AllOutlineMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithOwner owner获取
func (obj *_AllOutlineMgr) WithOwner(owner string) Option {
	return optionFunc(func(o *options) { o.query["owner"] = owner })
}

// WithUsed used获取
func (obj *_AllOutlineMgr) WithUsed(used int64) Option {
	return optionFunc(func(o *options) { o.query["used"] = used })
}

// WithVersion version获取
func (obj *_AllOutlineMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithCompatible compatible获取
func (obj *_AllOutlineMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithEnabled enabled获取
func (obj *_AllOutlineMgr) WithEnabled(enabled int64) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithFormat format获取
func (obj *_AllOutlineMgr) WithFormat(format int64) Option {
	return optionFunc(func(o *options) { o.query["format"] = format })
}

// WithOutlineParams outline_params获取
func (obj *_AllOutlineMgr) WithOutlineParams(outlineParams []byte) Option {
	return optionFunc(func(o *options) { o.query["outline_params"] = outlineParams })
}

// WithOutlineTarget outline_target获取
func (obj *_AllOutlineMgr) WithOutlineTarget(outlineTarget string) Option {
	return optionFunc(func(o *options) { o.query["outline_target"] = outlineTarget })
}

// WithSQLID sql_id获取
func (obj *_AllOutlineMgr) WithSQLID(sqlID []byte) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithOwnerID owner_id获取
func (obj *_AllOutlineMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// GetByOption 功能选项模式获取
func (obj *_AllOutlineMgr) GetByOption(opts ...Option) (result AllOutline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllOutlineMgr) GetByOptions(opts ...Option) (results []*AllOutline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllOutlineMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllOutline, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where(options.query)
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
func (obj *_AllOutlineMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllOutlineMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllOutlineMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllOutlineMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllOutlineMgr) GetFromTenantID(tenantID int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllOutlineMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromOutlineID 通过outline_id获取内容
func (obj *_AllOutlineMgr) GetFromOutlineID(outlineID int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`outline_id` = ?", outlineID).Find(&results).Error

	return
}

// GetBatchFromOutlineID 批量查找
func (obj *_AllOutlineMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllOutlineMgr) GetFromDatabaseID(databaseID int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllOutlineMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllOutlineMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllOutlineMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllOutlineMgr) GetFromName(name string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllOutlineMgr) GetBatchFromName(names []string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSignature 通过signature获取内容
func (obj *_AllOutlineMgr) GetFromSignature(signature []byte) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`signature` = ?", signature).Find(&results).Error

	return
}

// GetBatchFromSignature 批量查找
func (obj *_AllOutlineMgr) GetBatchFromSignature(signatures [][]byte) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`signature` IN (?)", signatures).Find(&results).Error

	return
}

// GetFromOutlineContent 通过outline_content获取内容
func (obj *_AllOutlineMgr) GetFromOutlineContent(outlineContent string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`outline_content` = ?", outlineContent).Find(&results).Error

	return
}

// GetBatchFromOutlineContent 批量查找
func (obj *_AllOutlineMgr) GetBatchFromOutlineContent(outlineContents []string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`outline_content` IN (?)", outlineContents).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllOutlineMgr) GetFromSQLText(sqlText string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllOutlineMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromOwner 通过owner获取内容
func (obj *_AllOutlineMgr) GetFromOwner(owner string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`owner` = ?", owner).Find(&results).Error

	return
}

// GetBatchFromOwner 批量查找
func (obj *_AllOutlineMgr) GetBatchFromOwner(owners []string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`owner` IN (?)", owners).Find(&results).Error

	return
}

// GetFromUsed 通过used获取内容
func (obj *_AllOutlineMgr) GetFromUsed(used int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`used` = ?", used).Find(&results).Error

	return
}

// GetBatchFromUsed 批量查找
func (obj *_AllOutlineMgr) GetBatchFromUsed(useds []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`used` IN (?)", useds).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllOutlineMgr) GetFromVersion(version string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllOutlineMgr) GetBatchFromVersion(versions []string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllOutlineMgr) GetFromCompatible(compatible int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllOutlineMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容
func (obj *_AllOutlineMgr) GetFromEnabled(enabled int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`enabled` = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量查找
func (obj *_AllOutlineMgr) GetBatchFromEnabled(enableds []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`enabled` IN (?)", enableds).Find(&results).Error

	return
}

// GetFromFormat 通过format获取内容
func (obj *_AllOutlineMgr) GetFromFormat(format int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`format` = ?", format).Find(&results).Error

	return
}

// GetBatchFromFormat 批量查找
func (obj *_AllOutlineMgr) GetBatchFromFormat(formats []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`format` IN (?)", formats).Find(&results).Error

	return
}

// GetFromOutlineParams 通过outline_params获取内容
func (obj *_AllOutlineMgr) GetFromOutlineParams(outlineParams []byte) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`outline_params` = ?", outlineParams).Find(&results).Error

	return
}

// GetBatchFromOutlineParams 批量查找
func (obj *_AllOutlineMgr) GetBatchFromOutlineParams(outlineParamss [][]byte) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`outline_params` IN (?)", outlineParamss).Find(&results).Error

	return
}

// GetFromOutlineTarget 通过outline_target获取内容
func (obj *_AllOutlineMgr) GetFromOutlineTarget(outlineTarget string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`outline_target` = ?", outlineTarget).Find(&results).Error

	return
}

// GetBatchFromOutlineTarget 批量查找
func (obj *_AllOutlineMgr) GetBatchFromOutlineTarget(outlineTargets []string) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`outline_target` IN (?)", outlineTargets).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllOutlineMgr) GetFromSQLID(sqlID []byte) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllOutlineMgr) GetBatchFromSQLID(sqlIDs [][]byte) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllOutlineMgr) GetFromOwnerID(ownerID int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllOutlineMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllOutlineMgr) FetchByPrimaryKey(tenantID int64, outlineID int64) (result AllOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllOutline{}).Where("`tenant_id` = ? AND `outline_id` = ?", tenantID, outlineID).First(&result).Error

	return
}
