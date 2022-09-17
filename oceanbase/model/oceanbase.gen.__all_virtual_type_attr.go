package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualTypeAttrMgr struct {
	*_BaseMgr
}

// AllVirtualTypeAttrMgr open func
func AllVirtualTypeAttrMgr(db *gorm.DB) *_AllVirtualTypeAttrMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTypeAttrMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTypeAttrMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_type_attr"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTypeAttrMgr) GetTableName() string {
	return "__all_virtual_type_attr"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTypeAttrMgr) Reset() *_AllVirtualTypeAttrMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTypeAttrMgr) Get() (result AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTypeAttrMgr) Gets() (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTypeAttrMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTypeAttrMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTypeID type_id获取
func (obj *_AllVirtualTypeAttrMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithAttribute attribute获取
func (obj *_AllVirtualTypeAttrMgr) WithAttribute(attribute int64) Option {
	return optionFunc(func(o *options) { o.query["attribute"] = attribute })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTypeAttrMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTypeAttrMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTypeAttrMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithTypeAttrID type_attr_id获取
func (obj *_AllVirtualTypeAttrMgr) WithTypeAttrID(typeAttrID int64) Option {
	return optionFunc(func(o *options) { o.query["type_attr_id"] = typeAttrID })
}

// WithName name获取
func (obj *_AllVirtualTypeAttrMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithProperties properties获取
func (obj *_AllVirtualTypeAttrMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithCharsetID charset_id获取
func (obj *_AllVirtualTypeAttrMgr) WithCharsetID(charsetID int64) Option {
	return optionFunc(func(o *options) { o.query["charset_id"] = charsetID })
}

// WithCharsetForm charset_form获取
func (obj *_AllVirtualTypeAttrMgr) WithCharsetForm(charsetForm int64) Option {
	return optionFunc(func(o *options) { o.query["charset_form"] = charsetForm })
}

// WithLength length获取
func (obj *_AllVirtualTypeAttrMgr) WithLength(length int64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithNumberPrecision number_precision获取
func (obj *_AllVirtualTypeAttrMgr) WithNumberPrecision(numberPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["number_precision"] = numberPrecision })
}

// WithScale scale获取
func (obj *_AllVirtualTypeAttrMgr) WithScale(scale int64) Option {
	return optionFunc(func(o *options) { o.query["scale"] = scale })
}

// WithZeroFill zero_fill获取
func (obj *_AllVirtualTypeAttrMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithCollType coll_type获取
func (obj *_AllVirtualTypeAttrMgr) WithCollType(collType int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type"] = collType })
}

// WithExternname externname获取
func (obj *_AllVirtualTypeAttrMgr) WithExternname(externname string) Option {
	return optionFunc(func(o *options) { o.query["externname"] = externname })
}

// WithXflags xflags获取
func (obj *_AllVirtualTypeAttrMgr) WithXflags(xflags int64) Option {
	return optionFunc(func(o *options) { o.query["xflags"] = xflags })
}

// WithSetter setter获取
func (obj *_AllVirtualTypeAttrMgr) WithSetter(setter int64) Option {
	return optionFunc(func(o *options) { o.query["setter"] = setter })
}

// WithGetter getter获取
func (obj *_AllVirtualTypeAttrMgr) WithGetter(getter int64) Option {
	return optionFunc(func(o *options) { o.query["getter"] = getter })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTypeAttrMgr) GetByOption(opts ...Option) (result AllVirtualTypeAttr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTypeAttrMgr) GetByOptions(opts ...Option) (results []*AllVirtualTypeAttr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTypeAttrMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTypeAttr, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where(options.query)
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
func (obj *_AllVirtualTypeAttrMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromTypeID(typeID int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromTypeID(typeIDs []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromAttribute 通过attribute获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromAttribute(attribute int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`attribute` = ?", attribute).Find(&results).Error

	return
}

// GetBatchFromAttribute 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromAttribute(attributes []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`attribute` IN (?)", attributes).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromTypeAttrID 通过type_attr_id获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromTypeAttrID(typeAttrID int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`type_attr_id` = ?", typeAttrID).Find(&results).Error

	return
}

// GetBatchFromTypeAttrID 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromTypeAttrID(typeAttrIDs []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`type_attr_id` IN (?)", typeAttrIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromName(name string) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromName(names []string) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromProperties(properties int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromProperties(propertiess []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromCharsetID 通过charset_id获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromCharsetID(charsetID int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`charset_id` = ?", charsetID).Find(&results).Error

	return
}

// GetBatchFromCharsetID 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromCharsetID(charsetIDs []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`charset_id` IN (?)", charsetIDs).Find(&results).Error

	return
}

// GetFromCharsetForm 通过charset_form获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromCharsetForm(charsetForm int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`charset_form` = ?", charsetForm).Find(&results).Error

	return
}

// GetBatchFromCharsetForm 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromCharsetForm(charsetForms []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`charset_form` IN (?)", charsetForms).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromLength(length int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromLength(lengths []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromNumberPrecision 通过number_precision获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromNumberPrecision(numberPrecision int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`number_precision` = ?", numberPrecision).Find(&results).Error

	return
}

// GetBatchFromNumberPrecision 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromNumberPrecision(numberPrecisions []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`number_precision` IN (?)", numberPrecisions).Find(&results).Error

	return
}

// GetFromScale 通过scale获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromScale(scale int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`scale` = ?", scale).Find(&results).Error

	return
}

// GetBatchFromScale 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromScale(scales []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`scale` IN (?)", scales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromZeroFill(zeroFill int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromCollType 通过coll_type获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromCollType(collType int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`coll_type` = ?", collType).Find(&results).Error

	return
}

// GetBatchFromCollType 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromCollType(collTypes []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`coll_type` IN (?)", collTypes).Find(&results).Error

	return
}

// GetFromExternname 通过externname获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromExternname(externname string) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`externname` = ?", externname).Find(&results).Error

	return
}

// GetBatchFromExternname 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromExternname(externnames []string) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`externname` IN (?)", externnames).Find(&results).Error

	return
}

// GetFromXflags 通过xflags获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromXflags(xflags int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`xflags` = ?", xflags).Find(&results).Error

	return
}

// GetBatchFromXflags 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromXflags(xflagss []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`xflags` IN (?)", xflagss).Find(&results).Error

	return
}

// GetFromSetter 通过setter获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromSetter(setter int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`setter` = ?", setter).Find(&results).Error

	return
}

// GetBatchFromSetter 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromSetter(setters []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`setter` IN (?)", setters).Find(&results).Error

	return
}

// GetFromGetter 通过getter获取内容
func (obj *_AllVirtualTypeAttrMgr) GetFromGetter(getter int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`getter` = ?", getter).Find(&results).Error

	return
}

// GetBatchFromGetter 批量查找
func (obj *_AllVirtualTypeAttrMgr) GetBatchFromGetter(getters []int64) (results []*AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`getter` IN (?)", getters).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTypeAttrMgr) FetchByPrimaryKey(tenantID int64, typeID int64, attribute int64) (result AllVirtualTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttr{}).Where("`tenant_id` = ? AND `type_id` = ? AND `attribute` = ?", tenantID, typeID, attribute).First(&result).Error

	return
}
