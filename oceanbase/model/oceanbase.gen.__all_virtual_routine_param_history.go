package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualRoutineParamHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualRoutineParamHistoryMgr open func
func AllVirtualRoutineParamHistoryMgr(db *gorm.DB) *_AllVirtualRoutineParamHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRoutineParamHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRoutineParamHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_routine_param_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRoutineParamHistoryMgr) GetTableName() string {
	return "__all_virtual_routine_param_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRoutineParamHistoryMgr) Reset() *_AllVirtualRoutineParamHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRoutineParamHistoryMgr) Get() (result AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRoutineParamHistoryMgr) Gets() (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRoutineParamHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRoutineID routine_id获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithRoutineID(routineID int64) Option {
	return optionFunc(func(o *options) { o.query["routine_id"] = routineID })
}

// WithSequence sequence获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithSequence(sequence int64) Option {
	return optionFunc(func(o *options) { o.query["sequence"] = sequence })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithSubprogramID subprogram_id获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithSubprogramID(subprogramID int64) Option {
	return optionFunc(func(o *options) { o.query["subprogram_id"] = subprogramID })
}

// WithParamPosition param_position获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamPosition(paramPosition int64) Option {
	return optionFunc(func(o *options) { o.query["param_position"] = paramPosition })
}

// WithParamLevel param_level获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamLevel(paramLevel int64) Option {
	return optionFunc(func(o *options) { o.query["param_level"] = paramLevel })
}

// WithParamName param_name获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamName(paramName string) Option {
	return optionFunc(func(o *options) { o.query["param_name"] = paramName })
}

// WithParamType param_type获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamType(paramType int64) Option {
	return optionFunc(func(o *options) { o.query["param_type"] = paramType })
}

// WithParamLength param_length获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamLength(paramLength int64) Option {
	return optionFunc(func(o *options) { o.query["param_length"] = paramLength })
}

// WithParamPrecision param_precision获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamPrecision(paramPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["param_precision"] = paramPrecision })
}

// WithParamScale param_scale获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamScale(paramScale int64) Option {
	return optionFunc(func(o *options) { o.query["param_scale"] = paramScale })
}

// WithParamZeroFill param_zero_fill获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamZeroFill(paramZeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["param_zero_fill"] = paramZeroFill })
}

// WithParamCharset param_charset获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamCharset(paramCharset int64) Option {
	return optionFunc(func(o *options) { o.query["param_charset"] = paramCharset })
}

// WithParamCollType param_coll_type获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithParamCollType(paramCollType int64) Option {
	return optionFunc(func(o *options) { o.query["param_coll_type"] = paramCollType })
}

// WithFlag flag获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithDefaultValue default_value获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithDefaultValue(defaultValue string) Option {
	return optionFunc(func(o *options) { o.query["default_value"] = defaultValue })
}

// WithTypeOwner type_owner获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithTypeOwner(typeOwner int64) Option {
	return optionFunc(func(o *options) { o.query["type_owner"] = typeOwner })
}

// WithTypeName type_name获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithTypeSubname type_subname获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithTypeSubname(typeSubname string) Option {
	return optionFunc(func(o *options) { o.query["type_subname"] = typeSubname })
}

// WithExtendedTypeInfo extended_type_info获取
func (obj *_AllVirtualRoutineParamHistoryMgr) WithExtendedTypeInfo(extendedTypeInfo []byte) Option {
	return optionFunc(func(o *options) { o.query["extended_type_info"] = extendedTypeInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRoutineParamHistoryMgr) GetByOption(opts ...Option) (result AllVirtualRoutineParamHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRoutineParamHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualRoutineParamHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRoutineParamHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRoutineParamHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where(options.query)
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
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoutineID 通过routine_id获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromRoutineID(routineID int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`routine_id` = ?", routineID).Find(&results).Error

	return
}

// GetBatchFromRoutineID 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromRoutineID(routineIDs []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`routine_id` IN (?)", routineIDs).Find(&results).Error

	return
}

// GetFromSequence 通过sequence获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromSequence(sequence int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`sequence` = ?", sequence).Find(&results).Error

	return
}

// GetBatchFromSequence 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromSequence(sequences []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`sequence` IN (?)", sequences).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromSubprogramID 通过subprogram_id获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromSubprogramID(subprogramID int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`subprogram_id` = ?", subprogramID).Find(&results).Error

	return
}

// GetBatchFromSubprogramID 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromSubprogramID(subprogramIDs []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`subprogram_id` IN (?)", subprogramIDs).Find(&results).Error

	return
}

// GetFromParamPosition 通过param_position获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamPosition(paramPosition int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_position` = ?", paramPosition).Find(&results).Error

	return
}

// GetBatchFromParamPosition 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamPosition(paramPositions []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_position` IN (?)", paramPositions).Find(&results).Error

	return
}

// GetFromParamLevel 通过param_level获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamLevel(paramLevel int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_level` = ?", paramLevel).Find(&results).Error

	return
}

// GetBatchFromParamLevel 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamLevel(paramLevels []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_level` IN (?)", paramLevels).Find(&results).Error

	return
}

// GetFromParamName 通过param_name获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamName(paramName string) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_name` = ?", paramName).Find(&results).Error

	return
}

// GetBatchFromParamName 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamName(paramNames []string) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_name` IN (?)", paramNames).Find(&results).Error

	return
}

// GetFromParamType 通过param_type获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamType(paramType int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_type` = ?", paramType).Find(&results).Error

	return
}

// GetBatchFromParamType 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamType(paramTypes []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_type` IN (?)", paramTypes).Find(&results).Error

	return
}

// GetFromParamLength 通过param_length获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamLength(paramLength int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_length` = ?", paramLength).Find(&results).Error

	return
}

// GetBatchFromParamLength 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamLength(paramLengths []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_length` IN (?)", paramLengths).Find(&results).Error

	return
}

// GetFromParamPrecision 通过param_precision获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamPrecision(paramPrecision int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_precision` = ?", paramPrecision).Find(&results).Error

	return
}

// GetBatchFromParamPrecision 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamPrecision(paramPrecisions []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_precision` IN (?)", paramPrecisions).Find(&results).Error

	return
}

// GetFromParamScale 通过param_scale获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamScale(paramScale int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_scale` = ?", paramScale).Find(&results).Error

	return
}

// GetBatchFromParamScale 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamScale(paramScales []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_scale` IN (?)", paramScales).Find(&results).Error

	return
}

// GetFromParamZeroFill 通过param_zero_fill获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamZeroFill(paramZeroFill int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_zero_fill` = ?", paramZeroFill).Find(&results).Error

	return
}

// GetBatchFromParamZeroFill 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamZeroFill(paramZeroFills []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_zero_fill` IN (?)", paramZeroFills).Find(&results).Error

	return
}

// GetFromParamCharset 通过param_charset获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamCharset(paramCharset int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_charset` = ?", paramCharset).Find(&results).Error

	return
}

// GetBatchFromParamCharset 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamCharset(paramCharsets []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_charset` IN (?)", paramCharsets).Find(&results).Error

	return
}

// GetFromParamCollType 通过param_coll_type获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromParamCollType(paramCollType int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_coll_type` = ?", paramCollType).Find(&results).Error

	return
}

// GetBatchFromParamCollType 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromParamCollType(paramCollTypes []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`param_coll_type` IN (?)", paramCollTypes).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromFlag(flag int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromFlag(flags []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromDefaultValue 通过default_value获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromDefaultValue(defaultValue string) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`default_value` = ?", defaultValue).Find(&results).Error

	return
}

// GetBatchFromDefaultValue 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromDefaultValue(defaultValues []string) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`default_value` IN (?)", defaultValues).Find(&results).Error

	return
}

// GetFromTypeOwner 通过type_owner获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromTypeOwner(typeOwner int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`type_owner` = ?", typeOwner).Find(&results).Error

	return
}

// GetBatchFromTypeOwner 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromTypeOwner(typeOwners []int64) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`type_owner` IN (?)", typeOwners).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromTypeName(typeName string) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromTypeName(typeNames []string) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromTypeSubname 通过type_subname获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromTypeSubname(typeSubname string) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`type_subname` = ?", typeSubname).Find(&results).Error

	return
}

// GetBatchFromTypeSubname 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromTypeSubname(typeSubnames []string) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`type_subname` IN (?)", typeSubnames).Find(&results).Error

	return
}

// GetFromExtendedTypeInfo 通过extended_type_info获取内容
func (obj *_AllVirtualRoutineParamHistoryMgr) GetFromExtendedTypeInfo(extendedTypeInfo []byte) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`extended_type_info` = ?", extendedTypeInfo).Find(&results).Error

	return
}

// GetBatchFromExtendedTypeInfo 批量查找
func (obj *_AllVirtualRoutineParamHistoryMgr) GetBatchFromExtendedTypeInfo(extendedTypeInfos [][]byte) (results []*AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`extended_type_info` IN (?)", extendedTypeInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualRoutineParamHistoryMgr) FetchByPrimaryKey(tenantID int64, routineID int64, sequence int64, schemaVersion int64) (result AllVirtualRoutineParamHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParamHistory{}).Where("`tenant_id` = ? AND `routine_id` = ? AND `sequence` = ? AND `schema_version` = ?", tenantID, routineID, sequence, schemaVersion).First(&result).Error

	return
}
