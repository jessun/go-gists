package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualObjectTypeMgr struct {
	*_BaseMgr
}

// AllVirtualObjectTypeMgr open func
func AllVirtualObjectTypeMgr(db *gorm.DB) *_AllVirtualObjectTypeMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualObjectTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualObjectTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_object_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualObjectTypeMgr) GetTableName() string {
	return "__all_virtual_object_type"
}

// Reset 重置gorm会话
func (obj *_AllVirtualObjectTypeMgr) Reset() *_AllVirtualObjectTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualObjectTypeMgr) Get() (result AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualObjectTypeMgr) Gets() (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualObjectTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualObjectTypeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithObjectTypeID object_type_id获取
func (obj *_AllVirtualObjectTypeMgr) WithObjectTypeID(objectTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["object_type_id"] = objectTypeID })
}

// WithType type获取
func (obj *_AllVirtualObjectTypeMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualObjectTypeMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualObjectTypeMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualObjectTypeMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithProperties properties获取
func (obj *_AllVirtualObjectTypeMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithCharsetID charset_id获取
func (obj *_AllVirtualObjectTypeMgr) WithCharsetID(charsetID int64) Option {
	return optionFunc(func(o *options) { o.query["charset_id"] = charsetID })
}

// WithCharsetForm charset_form获取
func (obj *_AllVirtualObjectTypeMgr) WithCharsetForm(charsetForm int64) Option {
	return optionFunc(func(o *options) { o.query["charset_form"] = charsetForm })
}

// WithLength length获取
func (obj *_AllVirtualObjectTypeMgr) WithLength(length int64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithNumberPrecision number_precision获取
func (obj *_AllVirtualObjectTypeMgr) WithNumberPrecision(numberPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["number_precision"] = numberPrecision })
}

// WithScale scale获取
func (obj *_AllVirtualObjectTypeMgr) WithScale(scale int64) Option {
	return optionFunc(func(o *options) { o.query["scale"] = scale })
}

// WithZeroFill zero_fill获取
func (obj *_AllVirtualObjectTypeMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithCollType coll_type获取
func (obj *_AllVirtualObjectTypeMgr) WithCollType(collType int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type"] = collType })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualObjectTypeMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithFlag flag获取
func (obj *_AllVirtualObjectTypeMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualObjectTypeMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithCompFlag comp_flag获取
func (obj *_AllVirtualObjectTypeMgr) WithCompFlag(compFlag int64) Option {
	return optionFunc(func(o *options) { o.query["comp_flag"] = compFlag })
}

// WithObjectName object_name获取
func (obj *_AllVirtualObjectTypeMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithExecEnv exec_env获取
func (obj *_AllVirtualObjectTypeMgr) WithExecEnv(execEnv string) Option {
	return optionFunc(func(o *options) { o.query["exec_env"] = execEnv })
}

// WithSource source获取
func (obj *_AllVirtualObjectTypeMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithComment comment获取
func (obj *_AllVirtualObjectTypeMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRouteSQL route_sql获取
func (obj *_AllVirtualObjectTypeMgr) WithRouteSQL(routeSQL string) Option {
	return optionFunc(func(o *options) { o.query["route_sql"] = routeSQL })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualObjectTypeMgr) GetByOption(opts ...Option) (result AllVirtualObjectType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualObjectTypeMgr) GetByOptions(opts ...Option) (results []*AllVirtualObjectType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualObjectTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualObjectType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where(options.query)
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
func (obj *_AllVirtualObjectTypeMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromObjectTypeID 通过object_type_id获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromObjectTypeID(objectTypeID int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`object_type_id` = ?", objectTypeID).Find(&results).Error

	return
}

// GetBatchFromObjectTypeID 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromObjectTypeID(objectTypeIDs []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`object_type_id` IN (?)", objectTypeIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromType(_type int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromType(_types []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromProperties(properties int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromProperties(propertiess []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromCharsetID 通过charset_id获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromCharsetID(charsetID int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`charset_id` = ?", charsetID).Find(&results).Error

	return
}

// GetBatchFromCharsetID 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromCharsetID(charsetIDs []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`charset_id` IN (?)", charsetIDs).Find(&results).Error

	return
}

// GetFromCharsetForm 通过charset_form获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromCharsetForm(charsetForm int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`charset_form` = ?", charsetForm).Find(&results).Error

	return
}

// GetBatchFromCharsetForm 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromCharsetForm(charsetForms []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`charset_form` IN (?)", charsetForms).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromLength(length int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromLength(lengths []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromNumberPrecision 通过number_precision获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromNumberPrecision(numberPrecision int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`number_precision` = ?", numberPrecision).Find(&results).Error

	return
}

// GetBatchFromNumberPrecision 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromNumberPrecision(numberPrecisions []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`number_precision` IN (?)", numberPrecisions).Find(&results).Error

	return
}

// GetFromScale 通过scale获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromScale(scale int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`scale` = ?", scale).Find(&results).Error

	return
}

// GetBatchFromScale 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromScale(scales []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`scale` IN (?)", scales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromZeroFill(zeroFill int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromCollType 通过coll_type获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromCollType(collType int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`coll_type` = ?", collType).Find(&results).Error

	return
}

// GetBatchFromCollType 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromCollType(collTypes []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`coll_type` IN (?)", collTypes).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromFlag(flag int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromFlag(flags []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromCompFlag 通过comp_flag获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromCompFlag(compFlag int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`comp_flag` = ?", compFlag).Find(&results).Error

	return
}

// GetBatchFromCompFlag 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromCompFlag(compFlags []int64) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`comp_flag` IN (?)", compFlags).Find(&results).Error

	return
}

// GetFromObjectName 通过object_name获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromObjectName(objectName string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`object_name` = ?", objectName).Find(&results).Error

	return
}

// GetBatchFromObjectName 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromObjectName(objectNames []string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`object_name` IN (?)", objectNames).Find(&results).Error

	return
}

// GetFromExecEnv 通过exec_env获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromExecEnv(execEnv string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`exec_env` = ?", execEnv).Find(&results).Error

	return
}

// GetBatchFromExecEnv 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromExecEnv(execEnvs []string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`exec_env` IN (?)", execEnvs).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromSource(source string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromSource(sources []string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromComment(comment string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromComment(comments []string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromRouteSQL 通过route_sql获取内容
func (obj *_AllVirtualObjectTypeMgr) GetFromRouteSQL(routeSQL string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`route_sql` = ?", routeSQL).Find(&results).Error

	return
}

// GetBatchFromRouteSQL 批量查找
func (obj *_AllVirtualObjectTypeMgr) GetBatchFromRouteSQL(routeSQLs []string) (results []*AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`route_sql` IN (?)", routeSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualObjectTypeMgr) FetchByPrimaryKey(tenantID int64, objectTypeID int64, _type int64) (result AllVirtualObjectType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjectType{}).Where("`tenant_id` = ? AND `object_type_id` = ? AND `type` = ?", tenantID, objectTypeID, _type).First(&result).Error

	return
}
