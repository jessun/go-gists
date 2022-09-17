package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTablegroupHistoryMgr struct {
	*_BaseMgr
}

// AllTablegroupHistoryMgr open func
func AllTablegroupHistoryMgr(db *gorm.DB) *_AllTablegroupHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTablegroupHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTablegroupHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tablegroup_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTablegroupHistoryMgr) GetTableName() string {
	return "__all_tablegroup_history"
}

// Reset 重置gorm会话
func (obj *_AllTablegroupHistoryMgr) Reset() *_AllTablegroupHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTablegroupHistoryMgr) Get() (result AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTablegroupHistoryMgr) Gets() (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTablegroupHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTablegroupHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTablegroupHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTablegroupHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllTablegroupHistoryMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTablegroupHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTablegroupHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithTablegroupName tablegroup_name获取
func (obj *_AllTablegroupHistoryMgr) WithTablegroupName(tablegroupName string) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_name"] = tablegroupName })
}

// WithComment comment获取
func (obj *_AllTablegroupHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithPrimaryZone primary_zone获取
func (obj *_AllTablegroupHistoryMgr) WithPrimaryZone(primaryZone string) Option {
	return optionFunc(func(o *options) { o.query["primary_zone"] = primaryZone })
}

// WithLocality locality获取
func (obj *_AllTablegroupHistoryMgr) WithLocality(locality string) Option {
	return optionFunc(func(o *options) { o.query["locality"] = locality })
}

// WithPartLevel part_level获取
func (obj *_AllTablegroupHistoryMgr) WithPartLevel(partLevel int64) Option {
	return optionFunc(func(o *options) { o.query["part_level"] = partLevel })
}

// WithPartFuncType part_func_type获取
func (obj *_AllTablegroupHistoryMgr) WithPartFuncType(partFuncType int64) Option {
	return optionFunc(func(o *options) { o.query["part_func_type"] = partFuncType })
}

// WithPartFuncExprNum part_func_expr_num获取
func (obj *_AllTablegroupHistoryMgr) WithPartFuncExprNum(partFuncExprNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_func_expr_num"] = partFuncExprNum })
}

// WithPartNum part_num获取
func (obj *_AllTablegroupHistoryMgr) WithPartNum(partNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_num"] = partNum })
}

// WithSubPartFuncType sub_part_func_type获取
func (obj *_AllTablegroupHistoryMgr) WithSubPartFuncType(subPartFuncType int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_func_type"] = subPartFuncType })
}

// WithSubPartFuncExprNum sub_part_func_expr_num获取
func (obj *_AllTablegroupHistoryMgr) WithSubPartFuncExprNum(subPartFuncExprNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_func_expr_num"] = subPartFuncExprNum })
}

// WithSubPartNum sub_part_num获取
func (obj *_AllTablegroupHistoryMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithMaxUsedPartID max_used_part_id获取
func (obj *_AllTablegroupHistoryMgr) WithMaxUsedPartID(maxUsedPartID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_part_id"] = maxUsedPartID })
}

// WithPartitionStatus partition_status获取
func (obj *_AllTablegroupHistoryMgr) WithPartitionStatus(partitionStatus int64) Option {
	return optionFunc(func(o *options) { o.query["partition_status"] = partitionStatus })
}

// WithPreviousLocality previous_locality获取
func (obj *_AllTablegroupHistoryMgr) WithPreviousLocality(previousLocality string) Option {
	return optionFunc(func(o *options) { o.query["previous_locality"] = previousLocality })
}

// WithPartitionSchemaVersion partition_schema_version获取
func (obj *_AllTablegroupHistoryMgr) WithPartitionSchemaVersion(partitionSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["partition_schema_version"] = partitionSchemaVersion })
}

// WithBinding binding获取
func (obj *_AllTablegroupHistoryMgr) WithBinding(binding int8) Option {
	return optionFunc(func(o *options) { o.query["binding"] = binding })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllTablegroupHistoryMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// WithIsSubPartTemplate is_sub_part_template获取
func (obj *_AllTablegroupHistoryMgr) WithIsSubPartTemplate(isSubPartTemplate int8) Option {
	return optionFunc(func(o *options) { o.query["is_sub_part_template"] = isSubPartTemplate })
}

// GetByOption 功能选项模式获取
func (obj *_AllTablegroupHistoryMgr) GetByOption(opts ...Option) (result AllTablegroupHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTablegroupHistoryMgr) GetByOptions(opts ...Option) (results []*AllTablegroupHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTablegroupHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTablegroupHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where(options.query)
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
func (obj *_AllTablegroupHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromTablegroupName 通过tablegroup_name获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromTablegroupName(tablegroupName string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`tablegroup_name` = ?", tablegroupName).Find(&results).Error

	return
}

// GetBatchFromTablegroupName 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromTablegroupName(tablegroupNames []string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`tablegroup_name` IN (?)", tablegroupNames).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromComment(comment string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromComment(comments []string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromPrimaryZone 通过primary_zone获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromPrimaryZone(primaryZone string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`primary_zone` = ?", primaryZone).Find(&results).Error

	return
}

// GetBatchFromPrimaryZone 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromPrimaryZone(primaryZones []string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`primary_zone` IN (?)", primaryZones).Find(&results).Error

	return
}

// GetFromLocality 通过locality获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromLocality(locality string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`locality` = ?", locality).Find(&results).Error

	return
}

// GetBatchFromLocality 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromLocality(localitys []string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`locality` IN (?)", localitys).Find(&results).Error

	return
}

// GetFromPartLevel 通过part_level获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromPartLevel(partLevel int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`part_level` = ?", partLevel).Find(&results).Error

	return
}

// GetBatchFromPartLevel 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromPartLevel(partLevels []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`part_level` IN (?)", partLevels).Find(&results).Error

	return
}

// GetFromPartFuncType 通过part_func_type获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromPartFuncType(partFuncType int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`part_func_type` = ?", partFuncType).Find(&results).Error

	return
}

// GetBatchFromPartFuncType 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromPartFuncType(partFuncTypes []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`part_func_type` IN (?)", partFuncTypes).Find(&results).Error

	return
}

// GetFromPartFuncExprNum 通过part_func_expr_num获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromPartFuncExprNum(partFuncExprNum int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`part_func_expr_num` = ?", partFuncExprNum).Find(&results).Error

	return
}

// GetBatchFromPartFuncExprNum 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromPartFuncExprNum(partFuncExprNums []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`part_func_expr_num` IN (?)", partFuncExprNums).Find(&results).Error

	return
}

// GetFromPartNum 通过part_num获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromPartNum(partNum int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`part_num` = ?", partNum).Find(&results).Error

	return
}

// GetBatchFromPartNum 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromPartNum(partNums []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`part_num` IN (?)", partNums).Find(&results).Error

	return
}

// GetFromSubPartFuncType 通过sub_part_func_type获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromSubPartFuncType(subPartFuncType int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`sub_part_func_type` = ?", subPartFuncType).Find(&results).Error

	return
}

// GetBatchFromSubPartFuncType 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromSubPartFuncType(subPartFuncTypes []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`sub_part_func_type` IN (?)", subPartFuncTypes).Find(&results).Error

	return
}

// GetFromSubPartFuncExprNum 通过sub_part_func_expr_num获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromSubPartFuncExprNum(subPartFuncExprNum int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`sub_part_func_expr_num` = ?", subPartFuncExprNum).Find(&results).Error

	return
}

// GetBatchFromSubPartFuncExprNum 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromSubPartFuncExprNum(subPartFuncExprNums []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`sub_part_func_expr_num` IN (?)", subPartFuncExprNums).Find(&results).Error

	return
}

// GetFromSubPartNum 通过sub_part_num获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromSubPartNum(subPartNum int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error

	return
}

// GetBatchFromSubPartNum 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error

	return
}

// GetFromMaxUsedPartID 通过max_used_part_id获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromMaxUsedPartID(maxUsedPartID int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`max_used_part_id` = ?", maxUsedPartID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedPartID 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromMaxUsedPartID(maxUsedPartIDs []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`max_used_part_id` IN (?)", maxUsedPartIDs).Find(&results).Error

	return
}

// GetFromPartitionStatus 通过partition_status获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromPartitionStatus(partitionStatus int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`partition_status` = ?", partitionStatus).Find(&results).Error

	return
}

// GetBatchFromPartitionStatus 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromPartitionStatus(partitionStatuss []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`partition_status` IN (?)", partitionStatuss).Find(&results).Error

	return
}

// GetFromPreviousLocality 通过previous_locality获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromPreviousLocality(previousLocality string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`previous_locality` = ?", previousLocality).Find(&results).Error

	return
}

// GetBatchFromPreviousLocality 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromPreviousLocality(previousLocalitys []string) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`previous_locality` IN (?)", previousLocalitys).Find(&results).Error

	return
}

// GetFromPartitionSchemaVersion 通过partition_schema_version获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromPartitionSchemaVersion(partitionSchemaVersion int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`partition_schema_version` = ?", partitionSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromPartitionSchemaVersion 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromPartitionSchemaVersion(partitionSchemaVersions []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`partition_schema_version` IN (?)", partitionSchemaVersions).Find(&results).Error

	return
}

// GetFromBinding 通过binding获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromBinding(binding int8) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`binding` = ?", binding).Find(&results).Error

	return
}

// GetBatchFromBinding 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromBinding(bindings []int8) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`binding` IN (?)", bindings).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

// GetFromIsSubPartTemplate 通过is_sub_part_template获取内容
func (obj *_AllTablegroupHistoryMgr) GetFromIsSubPartTemplate(isSubPartTemplate int8) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`is_sub_part_template` = ?", isSubPartTemplate).Find(&results).Error

	return
}

// GetBatchFromIsSubPartTemplate 批量查找
func (obj *_AllTablegroupHistoryMgr) GetBatchFromIsSubPartTemplate(isSubPartTemplates []int8) (results []*AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`is_sub_part_template` IN (?)", isSubPartTemplates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTablegroupHistoryMgr) FetchByPrimaryKey(tenantID int64, tablegroupID int64, schemaVersion int64) (result AllTablegroupHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTablegroupHistory{}).Where("`tenant_id` = ? AND `tablegroup_id` = ? AND `schema_version` = ?", tenantID, tablegroupID, schemaVersion).First(&result).Error

	return
}
