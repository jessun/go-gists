package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualErrorMgr struct {
	*_BaseMgr
}

// AllVirtualErrorMgr open func
func AllVirtualErrorMgr(db *gorm.DB) *_AllVirtualErrorMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualErrorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualErrorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_error"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualErrorMgr) GetTableName() string {
	return "__all_virtual_error"
}

// Reset 重置gorm会话
func (obj *_AllVirtualErrorMgr) Reset() *_AllVirtualErrorMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualErrorMgr) Get() (result AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualErrorMgr) Gets() (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualErrorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualErrorMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithObjID obj_id获取
func (obj *_AllVirtualErrorMgr) WithObjID(objID int64) Option {
	return optionFunc(func(o *options) { o.query["obj_id"] = objID })
}

// WithObjSeq obj_seq获取
func (obj *_AllVirtualErrorMgr) WithObjSeq(objSeq int64) Option {
	return optionFunc(func(o *options) { o.query["obj_seq"] = objSeq })
}

// WithObjType obj_type获取
func (obj *_AllVirtualErrorMgr) WithObjType(objType int64) Option {
	return optionFunc(func(o *options) { o.query["obj_type"] = objType })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualErrorMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualErrorMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithLine line获取
func (obj *_AllVirtualErrorMgr) WithLine(line int64) Option {
	return optionFunc(func(o *options) { o.query["line"] = line })
}

// WithPosition position获取
func (obj *_AllVirtualErrorMgr) WithPosition(position int64) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithTextLength text_length获取
func (obj *_AllVirtualErrorMgr) WithTextLength(textLength int64) Option {
	return optionFunc(func(o *options) { o.query["text_length"] = textLength })
}

// WithText text获取
func (obj *_AllVirtualErrorMgr) WithText(text string) Option {
	return optionFunc(func(o *options) { o.query["text"] = text })
}

// WithProperty property获取
func (obj *_AllVirtualErrorMgr) WithProperty(property int64) Option {
	return optionFunc(func(o *options) { o.query["property"] = property })
}

// WithErrorNumber error_number获取
func (obj *_AllVirtualErrorMgr) WithErrorNumber(errorNumber int64) Option {
	return optionFunc(func(o *options) { o.query["error_number"] = errorNumber })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualErrorMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualErrorMgr) GetByOption(opts ...Option) (result AllVirtualError, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualErrorMgr) GetByOptions(opts ...Option) (results []*AllVirtualError, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualErrorMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualError, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where(options.query)
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
func (obj *_AllVirtualErrorMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromObjID 通过obj_id获取内容
func (obj *_AllVirtualErrorMgr) GetFromObjID(objID int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`obj_id` = ?", objID).Find(&results).Error

	return
}

// GetBatchFromObjID 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromObjID(objIDs []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`obj_id` IN (?)", objIDs).Find(&results).Error

	return
}

// GetFromObjSeq 通过obj_seq获取内容
func (obj *_AllVirtualErrorMgr) GetFromObjSeq(objSeq int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`obj_seq` = ?", objSeq).Find(&results).Error

	return
}

// GetBatchFromObjSeq 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromObjSeq(objSeqs []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`obj_seq` IN (?)", objSeqs).Find(&results).Error

	return
}

// GetFromObjType 通过obj_type获取内容
func (obj *_AllVirtualErrorMgr) GetFromObjType(objType int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`obj_type` = ?", objType).Find(&results).Error

	return
}

// GetBatchFromObjType 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromObjType(objTypes []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`obj_type` IN (?)", objTypes).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualErrorMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualErrorMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromLine 通过line获取内容
func (obj *_AllVirtualErrorMgr) GetFromLine(line int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`line` = ?", line).Find(&results).Error

	return
}

// GetBatchFromLine 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromLine(lines []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`line` IN (?)", lines).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_AllVirtualErrorMgr) GetFromPosition(position int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`position` = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromPosition(positions []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`position` IN (?)", positions).Find(&results).Error

	return
}

// GetFromTextLength 通过text_length获取内容
func (obj *_AllVirtualErrorMgr) GetFromTextLength(textLength int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`text_length` = ?", textLength).Find(&results).Error

	return
}

// GetBatchFromTextLength 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromTextLength(textLengths []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`text_length` IN (?)", textLengths).Find(&results).Error

	return
}

// GetFromText 通过text获取内容
func (obj *_AllVirtualErrorMgr) GetFromText(text string) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`text` = ?", text).Find(&results).Error

	return
}

// GetBatchFromText 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromText(texts []string) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`text` IN (?)", texts).Find(&results).Error

	return
}

// GetFromProperty 通过property获取内容
func (obj *_AllVirtualErrorMgr) GetFromProperty(property int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`property` = ?", property).Find(&results).Error

	return
}

// GetBatchFromProperty 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromProperty(propertys []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`property` IN (?)", propertys).Find(&results).Error

	return
}

// GetFromErrorNumber 通过error_number获取内容
func (obj *_AllVirtualErrorMgr) GetFromErrorNumber(errorNumber int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`error_number` = ?", errorNumber).Find(&results).Error

	return
}

// GetBatchFromErrorNumber 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromErrorNumber(errorNumbers []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`error_number` IN (?)", errorNumbers).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualErrorMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualErrorMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualErrorMgr) FetchByPrimaryKey(tenantID int64, objID int64, objSeq int64, objType int64) (result AllVirtualError, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualError{}).Where("`tenant_id` = ? AND `obj_id` = ? AND `obj_seq` = ? AND `obj_type` = ?", tenantID, objID, objSeq, objType).First(&result).Error

	return
}
