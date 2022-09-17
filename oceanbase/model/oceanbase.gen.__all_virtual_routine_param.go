package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualRoutineParamMgr struct {
	*_BaseMgr
}

// AllVirtualRoutineParamMgr open func
func AllVirtualRoutineParamMgr(db *gorm.DB) *_AllVirtualRoutineParamMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualRoutineParamMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualRoutineParamMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_routine_param"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualRoutineParamMgr) GetTableName() string {
	return "__all_virtual_routine_param"
}

// Reset 重置gorm会话
func (obj *_AllVirtualRoutineParamMgr) Reset() *_AllVirtualRoutineParamMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualRoutineParamMgr) Get() (result AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualRoutineParamMgr) Gets() (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualRoutineParamMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualRoutineParamMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRoutineID routine_id获取
func (obj *_AllVirtualRoutineParamMgr) WithRoutineID(routineID int64) Option {
	return optionFunc(func(o *options) { o.query["routine_id"] = routineID })
}

// WithSequence sequence获取
func (obj *_AllVirtualRoutineParamMgr) WithSequence(sequence int64) Option {
	return optionFunc(func(o *options) { o.query["sequence"] = sequence })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualRoutineParamMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualRoutineParamMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSubprogramID subprogram_id获取
func (obj *_AllVirtualRoutineParamMgr) WithSubprogramID(subprogramID int64) Option {
	return optionFunc(func(o *options) { o.query["subprogram_id"] = subprogramID })
}

// WithParamPosition param_position获取
func (obj *_AllVirtualRoutineParamMgr) WithParamPosition(paramPosition int64) Option {
	return optionFunc(func(o *options) { o.query["param_position"] = paramPosition })
}

// WithParamLevel param_level获取
func (obj *_AllVirtualRoutineParamMgr) WithParamLevel(paramLevel int64) Option {
	return optionFunc(func(o *options) { o.query["param_level"] = paramLevel })
}

// WithParamName param_name获取
func (obj *_AllVirtualRoutineParamMgr) WithParamName(paramName string) Option {
	return optionFunc(func(o *options) { o.query["param_name"] = paramName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualRoutineParamMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithParamType param_type获取
func (obj *_AllVirtualRoutineParamMgr) WithParamType(paramType int64) Option {
	return optionFunc(func(o *options) { o.query["param_type"] = paramType })
}

// WithParamLength param_length获取
func (obj *_AllVirtualRoutineParamMgr) WithParamLength(paramLength int64) Option {
	return optionFunc(func(o *options) { o.query["param_length"] = paramLength })
}

// WithParamPrecision param_precision获取
func (obj *_AllVirtualRoutineParamMgr) WithParamPrecision(paramPrecision int64) Option {
	return optionFunc(func(o *options) { o.query["param_precision"] = paramPrecision })
}

// WithParamScale param_scale获取
func (obj *_AllVirtualRoutineParamMgr) WithParamScale(paramScale int64) Option {
	return optionFunc(func(o *options) { o.query["param_scale"] = paramScale })
}

// WithParamZeroFill param_zero_fill获取
func (obj *_AllVirtualRoutineParamMgr) WithParamZeroFill(paramZeroFill int64) Option {
	return optionFunc(func(o *options) { o.query["param_zero_fill"] = paramZeroFill })
}

// WithParamCharset param_charset获取
func (obj *_AllVirtualRoutineParamMgr) WithParamCharset(paramCharset int64) Option {
	return optionFunc(func(o *options) { o.query["param_charset"] = paramCharset })
}

// WithParamCollType param_coll_type获取
func (obj *_AllVirtualRoutineParamMgr) WithParamCollType(paramCollType int64) Option {
	return optionFunc(func(o *options) { o.query["param_coll_type"] = paramCollType })
}

// WithFlag flag获取
func (obj *_AllVirtualRoutineParamMgr) WithFlag(flag int64) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithDefaultValue default_value获取
func (obj *_AllVirtualRoutineParamMgr) WithDefaultValue(defaultValue string) Option {
	return optionFunc(func(o *options) { o.query["default_value"] = defaultValue })
}

// WithTypeOwner type_owner获取
func (obj *_AllVirtualRoutineParamMgr) WithTypeOwner(typeOwner int64) Option {
	return optionFunc(func(o *options) { o.query["type_owner"] = typeOwner })
}

// WithTypeName type_name获取
func (obj *_AllVirtualRoutineParamMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithTypeSubname type_subname获取
func (obj *_AllVirtualRoutineParamMgr) WithTypeSubname(typeSubname string) Option {
	return optionFunc(func(o *options) { o.query["type_subname"] = typeSubname })
}

// WithExtendedTypeInfo extended_type_info获取
func (obj *_AllVirtualRoutineParamMgr) WithExtendedTypeInfo(extendedTypeInfo []byte) Option {
	return optionFunc(func(o *options) { o.query["extended_type_info"] = extendedTypeInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualRoutineParamMgr) GetByOption(opts ...Option) (result AllVirtualRoutineParam, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualRoutineParamMgr) GetByOptions(opts ...Option) (results []*AllVirtualRoutineParam, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualRoutineParamMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualRoutineParam, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where(options.query)
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
func (obj *_AllVirtualRoutineParamMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRoutineID 通过routine_id获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromRoutineID(routineID int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`routine_id` = ?", routineID).Find(&results).Error

	return
}

// GetBatchFromRoutineID 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromRoutineID(routineIDs []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`routine_id` IN (?)", routineIDs).Find(&results).Error

	return
}

// GetFromSequence 通过sequence获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromSequence(sequence int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`sequence` = ?", sequence).Find(&results).Error

	return
}

// GetBatchFromSequence 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromSequence(sequences []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`sequence` IN (?)", sequences).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSubprogramID 通过subprogram_id获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromSubprogramID(subprogramID int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`subprogram_id` = ?", subprogramID).Find(&results).Error

	return
}

// GetBatchFromSubprogramID 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromSubprogramID(subprogramIDs []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`subprogram_id` IN (?)", subprogramIDs).Find(&results).Error

	return
}

// GetFromParamPosition 通过param_position获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamPosition(paramPosition int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_position` = ?", paramPosition).Find(&results).Error

	return
}

// GetBatchFromParamPosition 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamPosition(paramPositions []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_position` IN (?)", paramPositions).Find(&results).Error

	return
}

// GetFromParamLevel 通过param_level获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamLevel(paramLevel int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_level` = ?", paramLevel).Find(&results).Error

	return
}

// GetBatchFromParamLevel 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamLevel(paramLevels []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_level` IN (?)", paramLevels).Find(&results).Error

	return
}

// GetFromParamName 通过param_name获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamName(paramName string) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_name` = ?", paramName).Find(&results).Error

	return
}

// GetBatchFromParamName 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamName(paramNames []string) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_name` IN (?)", paramNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromParamType 通过param_type获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamType(paramType int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_type` = ?", paramType).Find(&results).Error

	return
}

// GetBatchFromParamType 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamType(paramTypes []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_type` IN (?)", paramTypes).Find(&results).Error

	return
}

// GetFromParamLength 通过param_length获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamLength(paramLength int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_length` = ?", paramLength).Find(&results).Error

	return
}

// GetBatchFromParamLength 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamLength(paramLengths []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_length` IN (?)", paramLengths).Find(&results).Error

	return
}

// GetFromParamPrecision 通过param_precision获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamPrecision(paramPrecision int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_precision` = ?", paramPrecision).Find(&results).Error

	return
}

// GetBatchFromParamPrecision 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamPrecision(paramPrecisions []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_precision` IN (?)", paramPrecisions).Find(&results).Error

	return
}

// GetFromParamScale 通过param_scale获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamScale(paramScale int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_scale` = ?", paramScale).Find(&results).Error

	return
}

// GetBatchFromParamScale 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamScale(paramScales []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_scale` IN (?)", paramScales).Find(&results).Error

	return
}

// GetFromParamZeroFill 通过param_zero_fill获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamZeroFill(paramZeroFill int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_zero_fill` = ?", paramZeroFill).Find(&results).Error

	return
}

// GetBatchFromParamZeroFill 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamZeroFill(paramZeroFills []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_zero_fill` IN (?)", paramZeroFills).Find(&results).Error

	return
}

// GetFromParamCharset 通过param_charset获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamCharset(paramCharset int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_charset` = ?", paramCharset).Find(&results).Error

	return
}

// GetBatchFromParamCharset 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamCharset(paramCharsets []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_charset` IN (?)", paramCharsets).Find(&results).Error

	return
}

// GetFromParamCollType 通过param_coll_type获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromParamCollType(paramCollType int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_coll_type` = ?", paramCollType).Find(&results).Error

	return
}

// GetBatchFromParamCollType 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromParamCollType(paramCollTypes []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`param_coll_type` IN (?)", paramCollTypes).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromFlag(flag int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromFlag(flags []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromDefaultValue 通过default_value获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromDefaultValue(defaultValue string) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`default_value` = ?", defaultValue).Find(&results).Error

	return
}

// GetBatchFromDefaultValue 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromDefaultValue(defaultValues []string) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`default_value` IN (?)", defaultValues).Find(&results).Error

	return
}

// GetFromTypeOwner 通过type_owner获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromTypeOwner(typeOwner int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`type_owner` = ?", typeOwner).Find(&results).Error

	return
}

// GetBatchFromTypeOwner 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromTypeOwner(typeOwners []int64) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`type_owner` IN (?)", typeOwners).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromTypeName(typeName string) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromTypeName(typeNames []string) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromTypeSubname 通过type_subname获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromTypeSubname(typeSubname string) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`type_subname` = ?", typeSubname).Find(&results).Error

	return
}

// GetBatchFromTypeSubname 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromTypeSubname(typeSubnames []string) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`type_subname` IN (?)", typeSubnames).Find(&results).Error

	return
}

// GetFromExtendedTypeInfo 通过extended_type_info获取内容
func (obj *_AllVirtualRoutineParamMgr) GetFromExtendedTypeInfo(extendedTypeInfo []byte) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`extended_type_info` = ?", extendedTypeInfo).Find(&results).Error

	return
}

// GetBatchFromExtendedTypeInfo 批量查找
func (obj *_AllVirtualRoutineParamMgr) GetBatchFromExtendedTypeInfo(extendedTypeInfos [][]byte) (results []*AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`extended_type_info` IN (?)", extendedTypeInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualRoutineParamMgr) FetchByPrimaryKey(tenantID int64, routineID int64, sequence int64) (result AllVirtualRoutineParam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualRoutineParam{}).Where("`tenant_id` = ? AND `routine_id` = ? AND `sequence` = ?", tenantID, routineID, sequence).First(&result).Error

	return
}
