package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllSysVariableHistoryMgr struct {
	*_BaseMgr
}

// AllSysVariableHistoryMgr open func
func AllSysVariableHistoryMgr(db *gorm.DB) *_AllSysVariableHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllSysVariableHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSysVariableHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sys_variable_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSysVariableHistoryMgr) GetTableName() string {
	return "__all_sys_variable_history"
}

// Reset 重置gorm会话
func (obj *_AllSysVariableHistoryMgr) Reset() *_AllSysVariableHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSysVariableHistoryMgr) Get() (result AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSysVariableHistoryMgr) Gets() (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSysVariableHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSysVariableHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSysVariableHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSysVariableHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllSysVariableHistoryMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithName name获取
func (obj *_AllSysVariableHistoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSchemaVersion schema_version获取
func (obj *_AllSysVariableHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllSysVariableHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDataType data_type获取
func (obj *_AllSysVariableHistoryMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllSysVariableHistoryMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllSysVariableHistoryMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithFlags flags获取
func (obj *_AllSysVariableHistoryMgr) WithFlags(flags int64) Option {
	return optionFunc(func(o *options) { o.query["flags"] = flags })
}

// WithMinVal min_val获取
func (obj *_AllSysVariableHistoryMgr) WithMinVal(minVal string) Option {
	return optionFunc(func(o *options) { o.query["min_val"] = minVal })
}

// WithMaxVal max_val获取
func (obj *_AllSysVariableHistoryMgr) WithMaxVal(maxVal string) Option {
	return optionFunc(func(o *options) { o.query["max_val"] = maxVal })
}

// GetByOption 功能选项模式获取
func (obj *_AllSysVariableHistoryMgr) GetByOption(opts ...Option) (result AllSysVariableHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSysVariableHistoryMgr) GetByOptions(opts ...Option) (results []*AllSysVariableHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSysVariableHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSysVariableHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where(options.query)
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
func (obj *_AllSysVariableHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromZone(zone string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromZone(zones []string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromName(name string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromName(names []string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromDataType(dataType int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromValue(value string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromValue(values []string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromInfo(info string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromInfo(infos []string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromFlags 通过flags获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromFlags(flags int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`flags` = ?", flags).Find(&results).Error

	return
}

// GetBatchFromFlags 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromFlags(flagss []int64) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`flags` IN (?)", flagss).Find(&results).Error

	return
}

// GetFromMinVal 通过min_val获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromMinVal(minVal string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`min_val` = ?", minVal).Find(&results).Error

	return
}

// GetBatchFromMinVal 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromMinVal(minVals []string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`min_val` IN (?)", minVals).Find(&results).Error

	return
}

// GetFromMaxVal 通过max_val获取内容
func (obj *_AllSysVariableHistoryMgr) GetFromMaxVal(maxVal string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`max_val` = ?", maxVal).Find(&results).Error

	return
}

// GetBatchFromMaxVal 批量查找
func (obj *_AllSysVariableHistoryMgr) GetBatchFromMaxVal(maxVals []string) (results []*AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`max_val` IN (?)", maxVals).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSysVariableHistoryMgr) FetchByPrimaryKey(tenantID int64, zone string, name string, schemaVersion int64) (result AllSysVariableHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariableHistory{}).Where("`tenant_id` = ? AND `zone` = ? AND `name` = ? AND `schema_version` = ?", tenantID, zone, name, schemaVersion).First(&result).Error

	return
}
