package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualOutlineMgr struct {
	*_BaseMgr
}

// AllVirtualOutlineMgr open func
func AllVirtualOutlineMgr(db *gorm.DB) *_AllVirtualOutlineMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualOutlineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualOutlineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_outline"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualOutlineMgr) GetTableName() string {
	return "__all_virtual_outline"
}

// Reset 重置gorm会话
func (obj *_AllVirtualOutlineMgr) Reset() *_AllVirtualOutlineMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualOutlineMgr) Get() (result AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualOutlineMgr) Gets() (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualOutlineMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualOutlineMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithOutlineID outline_id获取
func (obj *_AllVirtualOutlineMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualOutlineMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualOutlineMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualOutlineMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualOutlineMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithName name获取
func (obj *_AllVirtualOutlineMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSignature signature获取
func (obj *_AllVirtualOutlineMgr) WithSignature(signature []byte) Option {
	return optionFunc(func(o *options) { o.query["signature"] = signature })
}

// WithOutlineContent outline_content获取
func (obj *_AllVirtualOutlineMgr) WithOutlineContent(outlineContent string) Option {
	return optionFunc(func(o *options) { o.query["outline_content"] = outlineContent })
}

// WithSQLText sql_text获取
func (obj *_AllVirtualOutlineMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithOwner owner获取
func (obj *_AllVirtualOutlineMgr) WithOwner(owner string) Option {
	return optionFunc(func(o *options) { o.query["owner"] = owner })
}

// WithUsed used获取
func (obj *_AllVirtualOutlineMgr) WithUsed(used int64) Option {
	return optionFunc(func(o *options) { o.query["used"] = used })
}

// WithVersion version获取
func (obj *_AllVirtualOutlineMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithCompatible compatible获取
func (obj *_AllVirtualOutlineMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithEnabled enabled获取
func (obj *_AllVirtualOutlineMgr) WithEnabled(enabled int64) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithFormat format获取
func (obj *_AllVirtualOutlineMgr) WithFormat(format int64) Option {
	return optionFunc(func(o *options) { o.query["format"] = format })
}

// WithOutlineParams outline_params获取
func (obj *_AllVirtualOutlineMgr) WithOutlineParams(outlineParams []byte) Option {
	return optionFunc(func(o *options) { o.query["outline_params"] = outlineParams })
}

// WithOutlineTarget outline_target获取
func (obj *_AllVirtualOutlineMgr) WithOutlineTarget(outlineTarget string) Option {
	return optionFunc(func(o *options) { o.query["outline_target"] = outlineTarget })
}

// WithSQLID sql_id获取
func (obj *_AllVirtualOutlineMgr) WithSQLID(sqlID []byte) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualOutlineMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualOutlineMgr) GetByOption(opts ...Option) (result AllVirtualOutline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualOutlineMgr) GetByOptions(opts ...Option) (results []*AllVirtualOutline, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualOutlineMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualOutline, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where(options.query)
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
func (obj *_AllVirtualOutlineMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromOutlineID 通过outline_id获取内容
func (obj *_AllVirtualOutlineMgr) GetFromOutlineID(outlineID int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`outline_id` = ?", outlineID).Find(&results).Error

	return
}

// GetBatchFromOutlineID 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualOutlineMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualOutlineMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualOutlineMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualOutlineMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualOutlineMgr) GetFromName(name string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromName(names []string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSignature 通过signature获取内容
func (obj *_AllVirtualOutlineMgr) GetFromSignature(signature []byte) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`signature` = ?", signature).Find(&results).Error

	return
}

// GetBatchFromSignature 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromSignature(signatures [][]byte) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`signature` IN (?)", signatures).Find(&results).Error

	return
}

// GetFromOutlineContent 通过outline_content获取内容
func (obj *_AllVirtualOutlineMgr) GetFromOutlineContent(outlineContent string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`outline_content` = ?", outlineContent).Find(&results).Error

	return
}

// GetBatchFromOutlineContent 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromOutlineContent(outlineContents []string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`outline_content` IN (?)", outlineContents).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllVirtualOutlineMgr) GetFromSQLText(sqlText string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromOwner 通过owner获取内容
func (obj *_AllVirtualOutlineMgr) GetFromOwner(owner string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`owner` = ?", owner).Find(&results).Error

	return
}

// GetBatchFromOwner 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromOwner(owners []string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`owner` IN (?)", owners).Find(&results).Error

	return
}

// GetFromUsed 通过used获取内容
func (obj *_AllVirtualOutlineMgr) GetFromUsed(used int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`used` = ?", used).Find(&results).Error

	return
}

// GetBatchFromUsed 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromUsed(useds []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`used` IN (?)", useds).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualOutlineMgr) GetFromVersion(version string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromVersion(versions []string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllVirtualOutlineMgr) GetFromCompatible(compatible int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容
func (obj *_AllVirtualOutlineMgr) GetFromEnabled(enabled int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`enabled` = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromEnabled(enableds []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`enabled` IN (?)", enableds).Find(&results).Error

	return
}

// GetFromFormat 通过format获取内容
func (obj *_AllVirtualOutlineMgr) GetFromFormat(format int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`format` = ?", format).Find(&results).Error

	return
}

// GetBatchFromFormat 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromFormat(formats []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`format` IN (?)", formats).Find(&results).Error

	return
}

// GetFromOutlineParams 通过outline_params获取内容
func (obj *_AllVirtualOutlineMgr) GetFromOutlineParams(outlineParams []byte) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`outline_params` = ?", outlineParams).Find(&results).Error

	return
}

// GetBatchFromOutlineParams 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromOutlineParams(outlineParamss [][]byte) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`outline_params` IN (?)", outlineParamss).Find(&results).Error

	return
}

// GetFromOutlineTarget 通过outline_target获取内容
func (obj *_AllVirtualOutlineMgr) GetFromOutlineTarget(outlineTarget string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`outline_target` = ?", outlineTarget).Find(&results).Error

	return
}

// GetBatchFromOutlineTarget 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromOutlineTarget(outlineTargets []string) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`outline_target` IN (?)", outlineTargets).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllVirtualOutlineMgr) GetFromSQLID(sqlID []byte) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromSQLID(sqlIDs [][]byte) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualOutlineMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualOutlineMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualOutlineMgr) FetchByPrimaryKey(tenantID int64, outlineID int64) (result AllVirtualOutline, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutline{}).Where("`tenant_id` = ? AND `outline_id` = ?", tenantID, outlineID).First(&result).Error

	return
}
