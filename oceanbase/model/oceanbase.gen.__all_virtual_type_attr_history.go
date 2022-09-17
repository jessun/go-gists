package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTypeAttrHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualTypeAttrHistoryMgr open func
func AllVirtualTypeAttrHistoryMgr(db *gorm.DB) *_AllVirtualTypeAttrHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTypeAttrHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTypeAttrHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_type_attr_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTypeAttrHistoryMgr) GetTableName() string {
	return "__all_virtual_type_attr_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTypeAttrHistoryMgr) Reset() *_AllVirtualTypeAttrHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTypeAttrHistoryMgr) Get() (result AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTypeAttrHistoryMgr) Gets() (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTypeAttrHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTypeID type_id获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithAttribute attribute获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithAttribute(attribute int64) Option {
	return optionFunc(func(o *options) { o.query["attribute"] = attribute })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithTypeAttrID type_attr_id获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithTypeAttrID(typeAttrID int64) Option {
	return optionFunc(func(o *options) { o.query["type_attr_id"] = typeAttrID })
}

// WithName name获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithProperties properties获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithCharsetID charset_id获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithCharsetID(charsetID int64) Option {
	return optionFunc(func(o *options) { o.query["charset_id"] = charsetID })
}

// WithCharsetForm charset_form获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithCharsetForm(charsetForm int64) Option {
	return optionFunc(func(o *options) { o.query["charset_form"] = charsetForm })
}

// WithLength length获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithLength(length int64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithNumberPrecision number_precision获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithNumberPrecision(numberPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["number_precision"] = numberPrecision })
}

// WithScale scale获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithScale(scale int64) Option {
	return optionFunc(func(o *options) { o.query["scale"] = scale })
}

// WithZeroFill zero_fill获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithCollType coll_type获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithCollType(collType int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type"] = collType })
}

// WithExternname externname获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithExternname(externname string) Option {
	return optionFunc(func(o *options) { o.query["externname"] = externname })
}

// WithXflags xflags获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithXflags(xflags int64) Option {
	return optionFunc(func(o *options) { o.query["xflags"] = xflags })
}

// WithSetter setter获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithSetter(setter int64) Option {
	return optionFunc(func(o *options) { o.query["setter"] = setter })
}

// WithGetter getter获取
func (obj *_AllVirtualTypeAttrHistoryMgr) WithGetter(getter int64) Option {
	return optionFunc(func(o *options) { o.query["getter"] = getter })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTypeAttrHistoryMgr) GetByOption(opts ...Option) (result AllVirtualTypeAttrHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTypeAttrHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualTypeAttrHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTypeAttrHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTypeAttrHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where(options.query)
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
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromTypeID(typeID int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromTypeID(typeIDs []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromAttribute 通过attribute获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromAttribute(attribute int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`attribute` = ?", attribute).Find(&results).Error

	return
}

// GetBatchFromAttribute 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromAttribute(attributes []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`attribute` IN (?)", attributes).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromTypeAttrID 通过type_attr_id获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromTypeAttrID(typeAttrID int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`type_attr_id` = ?", typeAttrID).Find(&results).Error

	return
}

// GetBatchFromTypeAttrID 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromTypeAttrID(typeAttrIDs []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`type_attr_id` IN (?)", typeAttrIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromName(name string) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromName(names []string) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromProperties(properties int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromProperties(propertiess []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromCharsetID 通过charset_id获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromCharsetID(charsetID int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`charset_id` = ?", charsetID).Find(&results).Error

	return
}

// GetBatchFromCharsetID 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromCharsetID(charsetIDs []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`charset_id` IN (?)", charsetIDs).Find(&results).Error

	return
}

// GetFromCharsetForm 通过charset_form获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromCharsetForm(charsetForm int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`charset_form` = ?", charsetForm).Find(&results).Error

	return
}

// GetBatchFromCharsetForm 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromCharsetForm(charsetForms []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`charset_form` IN (?)", charsetForms).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromLength(length int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromLength(lengths []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromNumberPrecision 通过number_precision获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromNumberPrecision(numberPrecision int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`number_precision` = ?", numberPrecision).Find(&results).Error

	return
}

// GetBatchFromNumberPrecision 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromNumberPrecision(numberPrecisions []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`number_precision` IN (?)", numberPrecisions).Find(&results).Error

	return
}

// GetFromScale 通过scale获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromScale(scale int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`scale` = ?", scale).Find(&results).Error

	return
}

// GetBatchFromScale 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromScale(scales []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`scale` IN (?)", scales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromZeroFill(zeroFill int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromCollType 通过coll_type获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromCollType(collType int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`coll_type` = ?", collType).Find(&results).Error

	return
}

// GetBatchFromCollType 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromCollType(collTypes []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`coll_type` IN (?)", collTypes).Find(&results).Error

	return
}

// GetFromExternname 通过externname获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromExternname(externname string) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`externname` = ?", externname).Find(&results).Error

	return
}

// GetBatchFromExternname 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromExternname(externnames []string) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`externname` IN (?)", externnames).Find(&results).Error

	return
}

// GetFromXflags 通过xflags获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromXflags(xflags int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`xflags` = ?", xflags).Find(&results).Error

	return
}

// GetBatchFromXflags 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromXflags(xflagss []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`xflags` IN (?)", xflagss).Find(&results).Error

	return
}

// GetFromSetter 通过setter获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromSetter(setter int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`setter` = ?", setter).Find(&results).Error

	return
}

// GetBatchFromSetter 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromSetter(setters []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`setter` IN (?)", setters).Find(&results).Error

	return
}

// GetFromGetter 通过getter获取内容
func (obj *_AllVirtualTypeAttrHistoryMgr) GetFromGetter(getter int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`getter` = ?", getter).Find(&results).Error

	return
}

// GetBatchFromGetter 批量查找
func (obj *_AllVirtualTypeAttrHistoryMgr) GetBatchFromGetter(getters []int64) (results []*AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`getter` IN (?)", getters).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTypeAttrHistoryMgr) FetchByPrimaryKey(tenantID int64, typeID int64, attribute int64, schemaVersion int64) (result AllVirtualTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTypeAttrHistory{}).Where("`tenant_id` = ? AND `type_id` = ? AND `attribute` = ? AND `schema_version` = ?", tenantID, typeID, attribute, schemaVersion).First(&result).Error

	return
}
