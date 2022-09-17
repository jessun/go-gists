package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTypeAttrMgr struct {
	*_BaseMgr
}

// AllTypeAttrMgr open func
func AllTypeAttrMgr(db *gorm.DB) *_AllTypeAttrMgr {
	if db == nil {
		panic(fmt.Errorf("AllTypeAttrMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTypeAttrMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_type_attr"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTypeAttrMgr) GetTableName() string {
	return "__all_type_attr"
}

// Reset 重置gorm会话
func (obj *_AllTypeAttrMgr) Reset() *_AllTypeAttrMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTypeAttrMgr) Get() (result AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTypeAttrMgr) Gets() (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTypeAttrMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTypeAttrMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTypeAttrMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTypeAttrMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTypeID type_id获取
func (obj *_AllTypeAttrMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithAttribute attribute获取
func (obj *_AllTypeAttrMgr) WithAttribute(attribute int64) Option {
	return optionFunc(func(o *options) { o.query["attribute"] = attribute })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTypeAttrMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithTypeAttrID type_attr_id获取
func (obj *_AllTypeAttrMgr) WithTypeAttrID(typeAttrID int64) Option {
	return optionFunc(func(o *options) { o.query["type_attr_id"] = typeAttrID })
}

// WithName name获取
func (obj *_AllTypeAttrMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithProperties properties获取
func (obj *_AllTypeAttrMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithCharsetID charset_id获取
func (obj *_AllTypeAttrMgr) WithCharsetID(charsetID int64) Option {
	return optionFunc(func(o *options) { o.query["charset_id"] = charsetID })
}

// WithCharsetForm charset_form获取
func (obj *_AllTypeAttrMgr) WithCharsetForm(charsetForm int64) Option {
	return optionFunc(func(o *options) { o.query["charset_form"] = charsetForm })
}

// WithLength length获取
func (obj *_AllTypeAttrMgr) WithLength(length int64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithNumberPrecision number_precision获取
func (obj *_AllTypeAttrMgr) WithNumberPrecision(numberPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["number_precision"] = numberPrecision })
}

// WithScale scale获取
func (obj *_AllTypeAttrMgr) WithScale(scale int64) Option {
	return optionFunc(func(o *options) { o.query["scale"] = scale })
}

// WithZeroFill zero_fill获取
func (obj *_AllTypeAttrMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithCollType coll_type获取
func (obj *_AllTypeAttrMgr) WithCollType(collType int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type"] = collType })
}

// WithExternname externname获取
func (obj *_AllTypeAttrMgr) WithExternname(externname string) Option {
	return optionFunc(func(o *options) { o.query["externname"] = externname })
}

// WithXflags xflags获取
func (obj *_AllTypeAttrMgr) WithXflags(xflags int64) Option {
	return optionFunc(func(o *options) { o.query["xflags"] = xflags })
}

// WithSetter setter获取
func (obj *_AllTypeAttrMgr) WithSetter(setter int64) Option {
	return optionFunc(func(o *options) { o.query["setter"] = setter })
}

// WithGetter getter获取
func (obj *_AllTypeAttrMgr) WithGetter(getter int64) Option {
	return optionFunc(func(o *options) { o.query["getter"] = getter })
}

// GetByOption 功能选项模式获取
func (obj *_AllTypeAttrMgr) GetByOption(opts ...Option) (result AllTypeAttr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTypeAttrMgr) GetByOptions(opts ...Option) (results []*AllTypeAttr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTypeAttrMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTypeAttr, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where(options.query)
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
func (obj *_AllTypeAttrMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTypeAttrMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTypeAttrMgr) GetFromTenantID(tenantID int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_AllTypeAttrMgr) GetFromTypeID(typeID int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromTypeID(typeIDs []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromAttribute 通过attribute获取内容
func (obj *_AllTypeAttrMgr) GetFromAttribute(attribute int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`attribute` = ?", attribute).Find(&results).Error

	return
}

// GetBatchFromAttribute 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromAttribute(attributes []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`attribute` IN (?)", attributes).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTypeAttrMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromTypeAttrID 通过type_attr_id获取内容
func (obj *_AllTypeAttrMgr) GetFromTypeAttrID(typeAttrID int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`type_attr_id` = ?", typeAttrID).Find(&results).Error

	return
}

// GetBatchFromTypeAttrID 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromTypeAttrID(typeAttrIDs []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`type_attr_id` IN (?)", typeAttrIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllTypeAttrMgr) GetFromName(name string) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromName(names []string) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllTypeAttrMgr) GetFromProperties(properties int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromProperties(propertiess []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromCharsetID 通过charset_id获取内容
func (obj *_AllTypeAttrMgr) GetFromCharsetID(charsetID int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`charset_id` = ?", charsetID).Find(&results).Error

	return
}

// GetBatchFromCharsetID 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromCharsetID(charsetIDs []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`charset_id` IN (?)", charsetIDs).Find(&results).Error

	return
}

// GetFromCharsetForm 通过charset_form获取内容
func (obj *_AllTypeAttrMgr) GetFromCharsetForm(charsetForm int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`charset_form` = ?", charsetForm).Find(&results).Error

	return
}

// GetBatchFromCharsetForm 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromCharsetForm(charsetForms []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`charset_form` IN (?)", charsetForms).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_AllTypeAttrMgr) GetFromLength(length int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromLength(lengths []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromNumberPrecision 通过number_precision获取内容
func (obj *_AllTypeAttrMgr) GetFromNumberPrecision(numberPrecision int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`number_precision` = ?", numberPrecision).Find(&results).Error

	return
}

// GetBatchFromNumberPrecision 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromNumberPrecision(numberPrecisions []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`number_precision` IN (?)", numberPrecisions).Find(&results).Error

	return
}

// GetFromScale 通过scale获取内容
func (obj *_AllTypeAttrMgr) GetFromScale(scale int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`scale` = ?", scale).Find(&results).Error

	return
}

// GetBatchFromScale 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromScale(scales []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`scale` IN (?)", scales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllTypeAttrMgr) GetFromZeroFill(zeroFill int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromCollType 通过coll_type获取内容
func (obj *_AllTypeAttrMgr) GetFromCollType(collType int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`coll_type` = ?", collType).Find(&results).Error

	return
}

// GetBatchFromCollType 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromCollType(collTypes []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`coll_type` IN (?)", collTypes).Find(&results).Error

	return
}

// GetFromExternname 通过externname获取内容
func (obj *_AllTypeAttrMgr) GetFromExternname(externname string) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`externname` = ?", externname).Find(&results).Error

	return
}

// GetBatchFromExternname 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromExternname(externnames []string) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`externname` IN (?)", externnames).Find(&results).Error

	return
}

// GetFromXflags 通过xflags获取内容
func (obj *_AllTypeAttrMgr) GetFromXflags(xflags int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`xflags` = ?", xflags).Find(&results).Error

	return
}

// GetBatchFromXflags 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromXflags(xflagss []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`xflags` IN (?)", xflagss).Find(&results).Error

	return
}

// GetFromSetter 通过setter获取内容
func (obj *_AllTypeAttrMgr) GetFromSetter(setter int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`setter` = ?", setter).Find(&results).Error

	return
}

// GetBatchFromSetter 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromSetter(setters []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`setter` IN (?)", setters).Find(&results).Error

	return
}

// GetFromGetter 通过getter获取内容
func (obj *_AllTypeAttrMgr) GetFromGetter(getter int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`getter` = ?", getter).Find(&results).Error

	return
}

// GetBatchFromGetter 批量查找
func (obj *_AllTypeAttrMgr) GetBatchFromGetter(getters []int64) (results []*AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`getter` IN (?)", getters).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTypeAttrMgr) FetchByPrimaryKey(tenantID int64, typeID int64, attribute int64) (result AllTypeAttr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttr{}).Where("`tenant_id` = ? AND `type_id` = ? AND `attribute` = ?", tenantID, typeID, attribute).First(&result).Error

	return
}
