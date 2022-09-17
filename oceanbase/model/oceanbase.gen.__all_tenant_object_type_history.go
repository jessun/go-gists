package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantObjectTypeHistoryMgr struct {
	*_BaseMgr
}

// AllTenantObjectTypeHistoryMgr open func
func AllTenantObjectTypeHistoryMgr(db *gorm.DB) *_AllTenantObjectTypeHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantObjectTypeHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantObjectTypeHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_object_type_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantObjectTypeHistoryMgr) GetTableName() string {
	return "__all_tenant_object_type_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantObjectTypeHistoryMgr) Reset() *_AllTenantObjectTypeHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantObjectTypeHistoryMgr) Get() (result AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantObjectTypeHistoryMgr) Gets() (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantObjectTypeHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithObjectTypeID object_type_id获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithObjectTypeID(objectTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["object_type_id"] = objectTypeID })
}

// WithType type获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithProperties properties获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithCharsetID charset_id获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithCharsetID(charsetID int64) Option {
	return optionFunc(func(o *options) { o.query["charset_id"] = charsetID })
}

// WithCharsetForm charset_form获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithCharsetForm(charsetForm int64) Option {
	return optionFunc(func(o *options) { o.query["charset_form"] = charsetForm })
}

// WithLength length获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithLength(length int64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithNumberPrecision number_precision获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithNumberPrecision(numberPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["number_precision"] = numberPrecision })
}

// WithScale scale获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithScale(scale int64) Option {
	return optionFunc(func(o *options) { o.query["scale"] = scale })
}

// WithZeroFill zero_fill获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithZeroFill(zeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["zero_fill"] = zeroFill })
}

// WithCollType coll_type获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithCollType(collType int64) Option {
	return optionFunc(func(o *options) { o.query["coll_type"] = collType })
}

// WithDatabaseID database_id获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithFlag flag获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithOwnerID owner_id获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithCompFlag comp_flag获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithCompFlag(compFlag int64) Option {
	return optionFunc(func(o *options) { o.query["comp_flag"] = compFlag })
}

// WithObjectName object_name获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithExecEnv exec_env获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithExecEnv(execEnv string) Option {
	return optionFunc(func(o *options) { o.query["exec_env"] = execEnv })
}

// WithSource source获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithComment comment获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRouteSQL route_sql获取
func (obj *_AllTenantObjectTypeHistoryMgr) WithRouteSQL(routeSQL string) Option {
	return optionFunc(func(o *options) { o.query["route_sql"] = routeSQL })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantObjectTypeHistoryMgr) GetByOption(opts ...Option) (result AllTenantObjectTypeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantObjectTypeHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantObjectTypeHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantObjectTypeHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantObjectTypeHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where(options.query)
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
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromObjectTypeID 通过object_type_id获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromObjectTypeID(objectTypeID int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`object_type_id` = ?", objectTypeID).Find(&results).Error

	return
}

// GetBatchFromObjectTypeID 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromObjectTypeID(objectTypeIDs []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`object_type_id` IN (?)", objectTypeIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromType(_type int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromType(_types []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromProperties(properties int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromProperties(propertiess []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromCharsetID 通过charset_id获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromCharsetID(charsetID int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`charset_id` = ?", charsetID).Find(&results).Error

	return
}

// GetBatchFromCharsetID 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromCharsetID(charsetIDs []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`charset_id` IN (?)", charsetIDs).Find(&results).Error

	return
}

// GetFromCharsetForm 通过charset_form获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromCharsetForm(charsetForm int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`charset_form` = ?", charsetForm).Find(&results).Error

	return
}

// GetBatchFromCharsetForm 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromCharsetForm(charsetForms []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`charset_form` IN (?)", charsetForms).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromLength(length int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromLength(lengths []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromNumberPrecision 通过number_precision获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromNumberPrecision(numberPrecision int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`number_precision` = ?", numberPrecision).Find(&results).Error

	return
}

// GetBatchFromNumberPrecision 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromNumberPrecision(numberPrecisions []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`number_precision` IN (?)", numberPrecisions).Find(&results).Error

	return
}

// GetFromScale 通过scale获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromScale(scale int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`scale` = ?", scale).Find(&results).Error

	return
}

// GetBatchFromScale 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromScale(scales []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`scale` IN (?)", scales).Find(&results).Error

	return
}

// GetFromZeroFill 通过zero_fill获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromZeroFill(zeroFill int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`zero_fill` = ?", zeroFill).Find(&results).Error

	return
}

// GetBatchFromZeroFill 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromZeroFill(zeroFills []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`zero_fill` IN (?)", zeroFills).Find(&results).Error

	return
}

// GetFromCollType 通过coll_type获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromCollType(collType int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`coll_type` = ?", collType).Find(&results).Error

	return
}

// GetBatchFromCollType 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromCollType(collTypes []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`coll_type` IN (?)", collTypes).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromFlag(flag int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromFlag(flags []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromOwnerID(ownerID int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromCompFlag 通过comp_flag获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromCompFlag(compFlag int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`comp_flag` = ?", compFlag).Find(&results).Error

	return
}

// GetBatchFromCompFlag 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromCompFlag(compFlags []int64) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`comp_flag` IN (?)", compFlags).Find(&results).Error

	return
}

// GetFromObjectName 通过object_name获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromObjectName(objectName string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`object_name` = ?", objectName).Find(&results).Error

	return
}

// GetBatchFromObjectName 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromObjectName(objectNames []string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`object_name` IN (?)", objectNames).Find(&results).Error

	return
}

// GetFromExecEnv 通过exec_env获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromExecEnv(execEnv string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`exec_env` = ?", execEnv).Find(&results).Error

	return
}

// GetBatchFromExecEnv 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromExecEnv(execEnvs []string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`exec_env` IN (?)", execEnvs).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromSource(source string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromSource(sources []string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromComment(comment string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromComment(comments []string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromRouteSQL 通过route_sql获取内容
func (obj *_AllTenantObjectTypeHistoryMgr) GetFromRouteSQL(routeSQL string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`route_sql` = ?", routeSQL).Find(&results).Error

	return
}

// GetBatchFromRouteSQL 批量查找
func (obj *_AllTenantObjectTypeHistoryMgr) GetBatchFromRouteSQL(routeSQLs []string) (results []*AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`route_sql` IN (?)", routeSQLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantObjectTypeHistoryMgr) FetchByPrimaryKey(tenantID int64, objectTypeID int64, _type int64, schemaVersion int64) (result AllTenantObjectTypeHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjectTypeHistory{}).Where("`tenant_id` = ? AND `object_type_id` = ? AND `type` = ? AND `schema_version` = ?", tenantID, objectTypeID, _type, schemaVersion).First(&result).Error

	return
}
