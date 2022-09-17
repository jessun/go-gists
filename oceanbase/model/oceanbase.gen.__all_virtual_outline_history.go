package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualOutlineHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualOutlineHistoryMgr open func
func AllVirtualOutlineHistoryMgr(db *gorm.DB) *_AllVirtualOutlineHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualOutlineHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualOutlineHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_outline_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualOutlineHistoryMgr) GetTableName() string {
	return "__all_virtual_outline_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualOutlineHistoryMgr) Reset() *_AllVirtualOutlineHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualOutlineHistoryMgr) Get() (result AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualOutlineHistoryMgr) Gets() (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualOutlineHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualOutlineHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithOutlineID outline_id获取
func (obj *_AllVirtualOutlineHistoryMgr) WithOutlineID(outlineID int64) Option {
	return optionFunc(func(o *options) { o.query["outline_id"] = outlineID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualOutlineHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualOutlineHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualOutlineHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualOutlineHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualOutlineHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithName name获取
func (obj *_AllVirtualOutlineHistoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSignature signature获取
func (obj *_AllVirtualOutlineHistoryMgr) WithSignature(signature []byte) Option {
	return optionFunc(func(o *options) { o.query["signature"] = signature })
}

// WithOutlineContent outline_content获取
func (obj *_AllVirtualOutlineHistoryMgr) WithOutlineContent(outlineContent string) Option {
	return optionFunc(func(o *options) { o.query["outline_content"] = outlineContent })
}

// WithSQLText sql_text获取
func (obj *_AllVirtualOutlineHistoryMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithOwner owner获取
func (obj *_AllVirtualOutlineHistoryMgr) WithOwner(owner string) Option {
	return optionFunc(func(o *options) { o.query["owner"] = owner })
}

// WithUsed used获取
func (obj *_AllVirtualOutlineHistoryMgr) WithUsed(used int64) Option {
	return optionFunc(func(o *options) { o.query["used"] = used })
}

// WithVersion version获取
func (obj *_AllVirtualOutlineHistoryMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithCompatible compatible获取
func (obj *_AllVirtualOutlineHistoryMgr) WithCompatible(compatible int64) Option {
	return optionFunc(func(o *options) { o.query["compatible"] = compatible })
}

// WithEnabled enabled获取
func (obj *_AllVirtualOutlineHistoryMgr) WithEnabled(enabled int64) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithFormat format获取
func (obj *_AllVirtualOutlineHistoryMgr) WithFormat(format int64) Option {
	return optionFunc(func(o *options) { o.query["format"] = format })
}

// WithOutlineParams outline_params获取
func (obj *_AllVirtualOutlineHistoryMgr) WithOutlineParams(outlineParams []byte) Option {
	return optionFunc(func(o *options) { o.query["outline_params"] = outlineParams })
}

// WithOutlineTarget outline_target获取
func (obj *_AllVirtualOutlineHistoryMgr) WithOutlineTarget(outlineTarget string) Option {
	return optionFunc(func(o *options) { o.query["outline_target"] = outlineTarget })
}

// WithSQLID sql_id获取
func (obj *_AllVirtualOutlineHistoryMgr) WithSQLID(sqlID []byte) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualOutlineHistoryMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualOutlineHistoryMgr) GetByOption(opts ...Option) (result AllVirtualOutlineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualOutlineHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualOutlineHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualOutlineHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualOutlineHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where(options.query)
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
func (obj *_AllVirtualOutlineHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromOutlineID 通过outline_id获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromOutlineID(outlineID int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`outline_id` = ?", outlineID).Find(&results).Error

	return
}

// GetBatchFromOutlineID 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromOutlineID(outlineIDs []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`outline_id` IN (?)", outlineIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromName(name string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromName(names []string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSignature 通过signature获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromSignature(signature []byte) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`signature` = ?", signature).Find(&results).Error

	return
}

// GetBatchFromSignature 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromSignature(signatures [][]byte) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`signature` IN (?)", signatures).Find(&results).Error

	return
}

// GetFromOutlineContent 通过outline_content获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromOutlineContent(outlineContent string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`outline_content` = ?", outlineContent).Find(&results).Error

	return
}

// GetBatchFromOutlineContent 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromOutlineContent(outlineContents []string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`outline_content` IN (?)", outlineContents).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromSQLText(sqlText string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromOwner 通过owner获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromOwner(owner string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`owner` = ?", owner).Find(&results).Error

	return
}

// GetBatchFromOwner 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromOwner(owners []string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`owner` IN (?)", owners).Find(&results).Error

	return
}

// GetFromUsed 通过used获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromUsed(used int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`used` = ?", used).Find(&results).Error

	return
}

// GetBatchFromUsed 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromUsed(useds []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`used` IN (?)", useds).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromVersion(version string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromVersion(versions []string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromCompatible 通过compatible获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromCompatible(compatible int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`compatible` = ?", compatible).Find(&results).Error

	return
}

// GetBatchFromCompatible 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromCompatible(compatibles []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`compatible` IN (?)", compatibles).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromEnabled(enabled int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`enabled` = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromEnabled(enableds []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`enabled` IN (?)", enableds).Find(&results).Error

	return
}

// GetFromFormat 通过format获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromFormat(format int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`format` = ?", format).Find(&results).Error

	return
}

// GetBatchFromFormat 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromFormat(formats []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`format` IN (?)", formats).Find(&results).Error

	return
}

// GetFromOutlineParams 通过outline_params获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromOutlineParams(outlineParams []byte) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`outline_params` = ?", outlineParams).Find(&results).Error

	return
}

// GetBatchFromOutlineParams 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromOutlineParams(outlineParamss [][]byte) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`outline_params` IN (?)", outlineParamss).Find(&results).Error

	return
}

// GetFromOutlineTarget 通过outline_target获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromOutlineTarget(outlineTarget string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`outline_target` = ?", outlineTarget).Find(&results).Error

	return
}

// GetBatchFromOutlineTarget 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromOutlineTarget(outlineTargets []string) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`outline_target` IN (?)", outlineTargets).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromSQLID(sqlID []byte) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromSQLID(sqlIDs [][]byte) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualOutlineHistoryMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualOutlineHistoryMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualOutlineHistoryMgr) FetchByPrimaryKey(tenantID int64, outlineID int64, schemaVersion int64) (result AllVirtualOutlineHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualOutlineHistory{}).Where("`tenant_id` = ? AND `outline_id` = ? AND `schema_version` = ?", tenantID, outlineID, schemaVersion).First(&result).Error

	return
}
