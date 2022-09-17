package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTypeAttrHistoryMgr struct {
	*_BaseMgr
}

// AllTypeAttrHistoryMgr open func
func AllTypeAttrHistoryMgr(db *gorm.DB) *_AllTypeAttrHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTypeAttrHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTypeAttrHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_type_attr_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTypeAttrHistoryMgr) GetTableName() string {
	return "__all_type_attr_history"
}

// Reset 重置gorm会话
func (obj *_AllTypeAttrHistoryMgr) Reset() *_AllTypeAttrHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTypeAttrHistoryMgr) Get() (result AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTypeAttrHistoryMgr) Gets() (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTypeAttrHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTypeAttrHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTypeAttrHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTypeAttrHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTypeID type_id获取
func (obj *_AllTypeAttrHistoryMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithAttribute attribute获取
func (obj *_AllTypeAttrHistoryMgr) WithAttribute(attribute int64) Option {
	return optionFunc(func(o *options) { o.query["attribute"] = attribute })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTypeAttrHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTypeAttrHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithTypeAttrID type_attr_id获取
func (obj *_AllTypeAttrHistoryMgr) WithTypeAttrID(typeAttrID int64) Option {
	return optionFunc(func(o *options) { o.query["type_attr_id"] = typeAttrID })
}

// WithName name获取
func (obj *_AllTypeAttrHistoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithProperties properties获取
func (obj *_AllTypeAttrHistoryMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithCharsetID charset_id获取
func (obj *_AllTypeAttrHistoryMgr) WithCharsetID(charsetID int64) Option {
	return optionFunc(func(o *options) { o.query["charset_id"] = charsetID })
}

// WithCharsetForm charset_form获取
func (obj *_AllTypeAttrHistoryMgr) WithCharsetForm(charsetForm int64) Option {
	return optionFunc(func(o *options) { o.query["charset_form"] = charsetForm })
}

// WithLength length获取
func (obj *_AllTypeAttrHistoryMgr) WithLength(length int64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithNumberPrecision number_precision获取
func (obj *_AllTypeAttrHistoryMgr) WithNumberPrecision(numberPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["number_precision"] = numberPrecision })
}

// WithScale scale获取
func (obj *_AllTypeAttrHistoryMgr) WithScale(scale int64) Option {
	return optionFunc(func(o *options) { o.query["scale"] = scale })
}

// WithZeroFill zero_fill获取
func (obj *_AllTypeAttrHistoryMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithCollType coll_type获取
func (obj *_AllTypeAttrHistoryMgr) WithCollType(collType int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type"] = collType })
}

// WithExternname externname获取
func (obj *_AllTypeAttrHistoryMgr) WithExternname(externname string) Option {
	return optionFunc(func(o *options) { o.query["externname"] = externname })
}

// WithXflags xflags获取
func (obj *_AllTypeAttrHistoryMgr) WithXflags(xflags int64) Option {
	return optionFunc(func(o *options) { o.query["xflags"] = xflags })
}

// WithSetter setter获取
func (obj *_AllTypeAttrHistoryMgr) WithSetter(setter int64) Option {
	return optionFunc(func(o *options) { o.query["setter"] = setter })
}

// WithGetter getter获取
func (obj *_AllTypeAttrHistoryMgr) WithGetter(getter int64) Option {
	return optionFunc(func(o *options) { o.query["getter"] = getter })
}

// GetByOption 功能选项模式获取
func (obj *_AllTypeAttrHistoryMgr) GetByOption(opts ...Option) (result AllTypeAttrHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTypeAttrHistoryMgr) GetByOptions(opts ...Option) (results []*AllTypeAttrHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTypeAttrHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTypeAttrHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where(options.query)
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
func (obj *_AllTypeAttrHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromTypeID(typeID int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromTypeID(typeIDs []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromAttribute 通过attribute获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromAttribute(attribute int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`attribute` = ?", attribute).Find(&results).Error

	return
}

// GetBatchFromAttribute 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromAttribute(attributes []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`attribute` IN (?)", attributes).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromTypeAttrID 通过type_attr_id获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromTypeAttrID(typeAttrID int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`type_attr_id` = ?", typeAttrID).Find(&results).Error

	return
}

// GetBatchFromTypeAttrID 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromTypeAttrID(typeAttrIDs []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`type_attr_id` IN (?)", typeAttrIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromName(name string) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromName(names []string) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromProperties(properties int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromProperties(propertiess []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromCharsetID 通过charset_id获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromCharsetID(charsetID int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`charset_id` = ?", charsetID).Find(&results).Error

	return
}

// GetBatchFromCharsetID 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromCharsetID(charsetIDs []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`charset_id` IN (?)", charsetIDs).Find(&results).Error

	return
}

// GetFromCharsetForm 通过charset_form获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromCharsetForm(charsetForm int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`charset_form` = ?", charsetForm).Find(&results).Error

	return
}

// GetBatchFromCharsetForm 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromCharsetForm(charsetForms []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`charset_form` IN (?)", charsetForms).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromLength(length int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromLength(lengths []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromNumberPrecision 通过number_precision获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromNumberPrecision(numberPrecision int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`number_precision` = ?", numberPrecision).Find(&results).Error

	return
}

// GetBatchFromNumberPrecision 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromNumberPrecision(numberPrecisions []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`number_precision` IN (?)", numberPrecisions).Find(&results).Error

	return
}

// GetFromScale 通过scale获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromScale(scale int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`scale` = ?", scale).Find(&results).Error

	return
}

// GetBatchFromScale 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromScale(scales []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`scale` IN (?)", scales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromZeroFill(zeroFill int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromCollType 通过coll_type获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromCollType(collType int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`coll_type` = ?", collType).Find(&results).Error

	return
}

// GetBatchFromCollType 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromCollType(collTypes []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`coll_type` IN (?)", collTypes).Find(&results).Error

	return
}

// GetFromExternname 通过externname获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromExternname(externname string) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`externname` = ?", externname).Find(&results).Error

	return
}

// GetBatchFromExternname 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromExternname(externnames []string) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`externname` IN (?)", externnames).Find(&results).Error

	return
}

// GetFromXflags 通过xflags获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromXflags(xflags int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`xflags` = ?", xflags).Find(&results).Error

	return
}

// GetBatchFromXflags 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromXflags(xflagss []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`xflags` IN (?)", xflagss).Find(&results).Error

	return
}

// GetFromSetter 通过setter获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromSetter(setter int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`setter` = ?", setter).Find(&results).Error

	return
}

// GetBatchFromSetter 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromSetter(setters []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`setter` IN (?)", setters).Find(&results).Error

	return
}

// GetFromGetter 通过getter获取内容
func (obj *_AllTypeAttrHistoryMgr) GetFromGetter(getter int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`getter` = ?", getter).Find(&results).Error

	return
}

// GetBatchFromGetter 批量查找
func (obj *_AllTypeAttrHistoryMgr) GetBatchFromGetter(getters []int64) (results []*AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`getter` IN (?)", getters).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTypeAttrHistoryMgr) FetchByPrimaryKey(tenantID int64, typeID int64, attribute int64, schemaVersion int64) (result AllTypeAttrHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTypeAttrHistory{}).Where("`tenant_id` = ? AND `type_id` = ? AND `attribute` = ? AND `schema_version` = ?", tenantID, typeID, attribute, schemaVersion).First(&result).Error

	return
}
