package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSysVariableHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualSysVariableHistoryMgr open func
func AllVirtualSysVariableHistoryMgr(db *gorm.DB) *_AllVirtualSysVariableHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSysVariableHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSysVariableHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sys_variable_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSysVariableHistoryMgr) GetTableName() string {
	return "__all_virtual_sys_variable_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSysVariableHistoryMgr) Reset() *_AllVirtualSysVariableHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSysVariableHistoryMgr) Get() (result AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSysVariableHistoryMgr) Gets() (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSysVariableHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithName name获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDataType data_type获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithFlags flags获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithFlags(flags int64) Option {
	return optionFunc(func(o *options) { o.query["flags"] = flags })
}

// WithMinVal min_val获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithMinVal(minVal string) Option {
	return optionFunc(func(o *options) { o.query["min_val"] = minVal })
}

// WithMaxVal max_val获取
func (obj *_AllVirtualSysVariableHistoryMgr) WithMaxVal(maxVal string) Option {
	return optionFunc(func(o *options) { o.query["max_val"] = maxVal })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSysVariableHistoryMgr) GetByOption(opts ...Option) (result AllVirtualSysVariableHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSysVariableHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualSysVariableHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSysVariableHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSysVariableHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where(options.query)
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
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromZone(zone string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromZone(zones []string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromName(name string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromName(names []string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromDataType(dataType int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromValue(value string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromValue(values []string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromInfo(info string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromFlags 通过flags获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromFlags(flags int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`flags` = ?", flags).Find(&results).Error

	return
}

// GetBatchFromFlags 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromFlags(flagss []int64) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`flags` IN (?)", flagss).Find(&results).Error

	return
}

// GetFromMinVal 通过min_val获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromMinVal(minVal string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`min_val` = ?", minVal).Find(&results).Error

	return
}

// GetBatchFromMinVal 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromMinVal(minVals []string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`min_val` IN (?)", minVals).Find(&results).Error

	return
}

// GetFromMaxVal 通过max_val获取内容
func (obj *_AllVirtualSysVariableHistoryMgr) GetFromMaxVal(maxVal string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`max_val` = ?", maxVal).Find(&results).Error

	return
}

// GetBatchFromMaxVal 批量查找
func (obj *_AllVirtualSysVariableHistoryMgr) GetBatchFromMaxVal(maxVals []string) (results []*AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`max_val` IN (?)", maxVals).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSysVariableHistoryMgr) FetchByPrimaryKey(tenantID int64, zone string, name string, schemaVersion int64) (result AllVirtualSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariableHistory{}).Where("`tenant_id` = ? AND `zone` = ? AND `name` = ? AND `schema_version` = ?", tenantID, zone, name, schemaVersion).First(&result).Error

	return
}
