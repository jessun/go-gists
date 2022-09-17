package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllCollTypeHistoryMgr struct {
	*_BaseMgr
}

// AllCollTypeHistoryMgr open func
func AllCollTypeHistoryMgr(db *gorm.DB) *_AllCollTypeHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllCollTypeHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllCollTypeHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_coll_type_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllCollTypeHistoryMgr) GetTableName() string {
	return "__all_coll_type_history"
}

// Reset 重置gorm会话
func (obj *_AllCollTypeHistoryMgr) Reset() *_AllCollTypeHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllCollTypeHistoryMgr) Get() (result AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllCollTypeHistoryMgr) Gets() (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllCollTypeHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllCollTypeHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllCollTypeHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllCollTypeHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithCollTypeID coll_type_id获取
func (obj *_AllCollTypeHistoryMgr) WithCollTypeID(collTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type_id"] = collTypeID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllCollTypeHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllCollTypeHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithElemTypeID elem_type_id获取
func (obj *_AllCollTypeHistoryMgr) WithElemTypeID(elemTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["elem_type_id"] = elemTypeID })
}

// WithElemSchemaVersion elem_schema_version获取
func (obj *_AllCollTypeHistoryMgr) WithElemSchemaVersion(elemSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["elem_schema_version"] = elemSchemaVersion })
}

// WithProperties properties获取
func (obj *_AllCollTypeHistoryMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithCharsetID charset_id获取
func (obj *_AllCollTypeHistoryMgr) WithCharsetID(charsetID int64) Option {
	return optionFunc(func(o *options) { o.query["charset_id"] = charsetID })
}

// WithCharsetForm charset_form获取
func (obj *_AllCollTypeHistoryMgr) WithCharsetForm(charsetForm int64) Option {
	return optionFunc(func(o *options) { o.query["charset_form"] = charsetForm })
}

// WithLength length获取
func (obj *_AllCollTypeHistoryMgr) WithLength(length int64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithNumberPrecision number_precision获取
func (obj *_AllCollTypeHistoryMgr) WithNumberPrecision(numberPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["number_precision"] = numberPrecision })
}

// WithScale scale获取
func (obj *_AllCollTypeHistoryMgr) WithScale(scale int64) Option {
	return optionFunc(func(o *options) { o.query["scale"] = scale })
}

// WithZeroFill zero_fill获取
func (obj *_AllCollTypeHistoryMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithCollType coll_type获取
func (obj *_AllCollTypeHistoryMgr) WithCollType(collType int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type"] = collType })
}

// WithUpperBound upper_bound获取
func (obj *_AllCollTypeHistoryMgr) WithUpperBound(upperBound int64) Option {
	return optionFunc(func(o *options) { o.query["upper_bound"] = upperBound })
}

// WithPackageID package_id获取
func (obj *_AllCollTypeHistoryMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithCollName coll_name获取
func (obj *_AllCollTypeHistoryMgr) WithCollName(collName string) Option {
	return optionFunc(func(o *options) { o.query["coll_name"] = collName })
}

// GetByOption 功能选项模式获取
func (obj *_AllCollTypeHistoryMgr) GetByOption(opts ...Option) (result AllCollTypeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllCollTypeHistoryMgr) GetByOptions(opts ...Option) (results []*AllCollTypeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllCollTypeHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllCollTypeHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where(options.query)
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
func (obj *_AllCollTypeHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromCollTypeID 通过coll_type_id获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromCollTypeID(collTypeID int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`coll_type_id` = ?", collTypeID).Find(&results).Error

	return
}

// GetBatchFromCollTypeID 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromCollTypeID(collTypeIDs []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`coll_type_id` IN (?)", collTypeIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromElemTypeID 通过elem_type_id获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromElemTypeID(elemTypeID int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`elem_type_id` = ?", elemTypeID).Find(&results).Error

	return
}

// GetBatchFromElemTypeID 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromElemTypeID(elemTypeIDs []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`elem_type_id` IN (?)", elemTypeIDs).Find(&results).Error

	return
}

// GetFromElemSchemaVersion 通过elem_schema_version获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromElemSchemaVersion(elemSchemaVersion int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`elem_schema_version` = ?", elemSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromElemSchemaVersion 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromElemSchemaVersion(elemSchemaVersions []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`elem_schema_version` IN (?)", elemSchemaVersions).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromProperties(properties int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromProperties(propertiess []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromCharsetID 通过charset_id获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromCharsetID(charsetID int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`charset_id` = ?", charsetID).Find(&results).Error

	return
}

// GetBatchFromCharsetID 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromCharsetID(charsetIDs []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`charset_id` IN (?)", charsetIDs).Find(&results).Error

	return
}

// GetFromCharsetForm 通过charset_form获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromCharsetForm(charsetForm int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`charset_form` = ?", charsetForm).Find(&results).Error

	return
}

// GetBatchFromCharsetForm 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromCharsetForm(charsetForms []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`charset_form` IN (?)", charsetForms).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromLength(length int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromLength(lengths []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromNumberPrecision 通过number_precision获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromNumberPrecision(numberPrecision int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`number_precision` = ?", numberPrecision).Find(&results).Error

	return
}

// GetBatchFromNumberPrecision 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromNumberPrecision(numberPrecisions []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`number_precision` IN (?)", numberPrecisions).Find(&results).Error

	return
}

// GetFromScale 通过scale获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromScale(scale int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`scale` = ?", scale).Find(&results).Error

	return
}

// GetBatchFromScale 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromScale(scales []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`scale` IN (?)", scales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromZeroFill(zeroFill int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromCollType 通过coll_type获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromCollType(collType int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`coll_type` = ?", collType).Find(&results).Error

	return
}

// GetBatchFromCollType 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromCollType(collTypes []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`coll_type` IN (?)", collTypes).Find(&results).Error

	return
}

// GetFromUpperBound 通过upper_bound获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromUpperBound(upperBound int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`upper_bound` = ?", upperBound).Find(&results).Error

	return
}

// GetBatchFromUpperBound 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromUpperBound(upperBounds []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`upper_bound` IN (?)", upperBounds).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromPackageID(packageID int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromCollName 通过coll_name获取内容
func (obj *_AllCollTypeHistoryMgr) GetFromCollName(collName string) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`coll_name` = ?", collName).Find(&results).Error

	return
}

// GetBatchFromCollName 批量查找
func (obj *_AllCollTypeHistoryMgr) GetBatchFromCollName(collNames []string) (results []*AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`coll_name` IN (?)", collNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllCollTypeHistoryMgr) FetchByPrimaryKey(tenantID int64, collTypeID int64, schemaVersion int64) (result AllCollTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllCollTypeHistory{}).Where("`tenant_id` = ? AND `coll_type_id` = ? AND `schema_version` = ?", tenantID, collTypeID, schemaVersion).First(&result).Error

	return
}
