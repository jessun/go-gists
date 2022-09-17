package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualCollTypeMgr struct {
	*_BaseMgr
}

// AllVirtualCollTypeMgr open func
func AllVirtualCollTypeMgr(db *gorm.DB) *_AllVirtualCollTypeMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualCollTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualCollTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_coll_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualCollTypeMgr) GetTableName() string {
	return "__all_virtual_coll_type"
}

// Reset 重置gorm会话
func (obj *_AllVirtualCollTypeMgr) Reset() *_AllVirtualCollTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualCollTypeMgr) Get() (result AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualCollTypeMgr) Gets() (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualCollTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualCollTypeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithCollTypeID coll_type_id获取
func (obj *_AllVirtualCollTypeMgr) WithCollTypeID(collTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type_id"] = collTypeID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualCollTypeMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualCollTypeMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualCollTypeMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithElemTypeID elem_type_id获取
func (obj *_AllVirtualCollTypeMgr) WithElemTypeID(elemTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["elem_type_id"] = elemTypeID })
}

// WithElemSchemaVersion elem_schema_version获取
func (obj *_AllVirtualCollTypeMgr) WithElemSchemaVersion(elemSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["elem_schema_version"] = elemSchemaVersion })
}

// WithProperties properties获取
func (obj *_AllVirtualCollTypeMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithCharsetID charset_id获取
func (obj *_AllVirtualCollTypeMgr) WithCharsetID(charsetID int64) Option {
	return optionFunc(func(o *options) { o.query["charset_id"] = charsetID })
}

// WithCharsetForm charset_form获取
func (obj *_AllVirtualCollTypeMgr) WithCharsetForm(charsetForm int64) Option {
	return optionFunc(func(o *options) { o.query["charset_form"] = charsetForm })
}

// WithLength length获取
func (obj *_AllVirtualCollTypeMgr) WithLength(length int64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithNumberPrecision number_precision获取
func (obj *_AllVirtualCollTypeMgr) WithNumberPrecision(numberPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["number_precision"] = numberPrecision })
}

// WithScale scale获取
func (obj *_AllVirtualCollTypeMgr) WithScale(scale int64) Option {
	return optionFunc(func(o *options) { o.query["scale"] = scale })
}

// WithZeroFill zero_fill获取
func (obj *_AllVirtualCollTypeMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithCollType coll_type获取
func (obj *_AllVirtualCollTypeMgr) WithCollType(collType int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type"] = collType })
}

// WithUpperBound upper_bound获取
func (obj *_AllVirtualCollTypeMgr) WithUpperBound(upperBound int64) Option {
	return optionFunc(func(o *options) { o.query["upper_bound"] = upperBound })
}

// WithPackageID package_id获取
func (obj *_AllVirtualCollTypeMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithCollName coll_name获取
func (obj *_AllVirtualCollTypeMgr) WithCollName(collName string) Option {
	return optionFunc(func(o *options) { o.query["coll_name"] = collName })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualCollTypeMgr) GetByOption(opts ...Option) (result AllVirtualCollType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualCollTypeMgr) GetByOptions(opts ...Option) (results []*AllVirtualCollType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualCollTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualCollType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where(options.query)
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
func (obj *_AllVirtualCollTypeMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromCollTypeID 通过coll_type_id获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromCollTypeID(collTypeID int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`coll_type_id` = ?", collTypeID).Find(&results).Error

	return
}

// GetBatchFromCollTypeID 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromCollTypeID(collTypeIDs []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`coll_type_id` IN (?)", collTypeIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromElemTypeID 通过elem_type_id获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromElemTypeID(elemTypeID int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`elem_type_id` = ?", elemTypeID).Find(&results).Error

	return
}

// GetBatchFromElemTypeID 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromElemTypeID(elemTypeIDs []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`elem_type_id` IN (?)", elemTypeIDs).Find(&results).Error

	return
}

// GetFromElemSchemaVersion 通过elem_schema_version获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromElemSchemaVersion(elemSchemaVersion int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`elem_schema_version` = ?", elemSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromElemSchemaVersion 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromElemSchemaVersion(elemSchemaVersions []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`elem_schema_version` IN (?)", elemSchemaVersions).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromProperties(properties int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromProperties(propertiess []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromCharsetID 通过charset_id获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromCharsetID(charsetID int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`charset_id` = ?", charsetID).Find(&results).Error

	return
}

// GetBatchFromCharsetID 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromCharsetID(charsetIDs []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`charset_id` IN (?)", charsetIDs).Find(&results).Error

	return
}

// GetFromCharsetForm 通过charset_form获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromCharsetForm(charsetForm int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`charset_form` = ?", charsetForm).Find(&results).Error

	return
}

// GetBatchFromCharsetForm 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromCharsetForm(charsetForms []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`charset_form` IN (?)", charsetForms).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromLength(length int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromLength(lengths []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromNumberPrecision 通过number_precision获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromNumberPrecision(numberPrecision int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`number_precision` = ?", numberPrecision).Find(&results).Error

	return
}

// GetBatchFromNumberPrecision 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromNumberPrecision(numberPrecisions []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`number_precision` IN (?)", numberPrecisions).Find(&results).Error

	return
}

// GetFromScale 通过scale获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromScale(scale int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`scale` = ?", scale).Find(&results).Error

	return
}

// GetBatchFromScale 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromScale(scales []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`scale` IN (?)", scales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromZeroFill(zeroFill int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromCollType 通过coll_type获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromCollType(collType int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`coll_type` = ?", collType).Find(&results).Error

	return
}

// GetBatchFromCollType 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromCollType(collTypes []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`coll_type` IN (?)", collTypes).Find(&results).Error

	return
}

// GetFromUpperBound 通过upper_bound获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromUpperBound(upperBound int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`upper_bound` = ?", upperBound).Find(&results).Error

	return
}

// GetBatchFromUpperBound 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromUpperBound(upperBounds []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`upper_bound` IN (?)", upperBounds).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromPackageID(packageID int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromCollName 通过coll_name获取内容
func (obj *_AllVirtualCollTypeMgr) GetFromCollName(collName string) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`coll_name` = ?", collName).Find(&results).Error

	return
}

// GetBatchFromCollName 批量查找
func (obj *_AllVirtualCollTypeMgr) GetBatchFromCollName(collNames []string) (results []*AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`coll_name` IN (?)", collNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualCollTypeMgr) FetchByPrimaryKey(tenantID int64, collTypeID int64) (result AllVirtualCollType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualCollType{}).Where("`tenant_id` = ? AND `coll_type_id` = ?", tenantID, collTypeID).First(&result).Error

	return
}
