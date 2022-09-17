package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSysVariableMgr struct {
	*_BaseMgr
}

// AllVirtualSysVariableMgr open func
func AllVirtualSysVariableMgr(db *gorm.DB) *_AllVirtualSysVariableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSysVariableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSysVariableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sys_variable"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSysVariableMgr) GetTableName() string {
	return "__all_virtual_sys_variable"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSysVariableMgr) Reset() *_AllVirtualSysVariableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSysVariableMgr) Get() (result AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSysVariableMgr) Gets() (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSysVariableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSysVariableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllVirtualSysVariableMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithName name获取
func (obj *_AllVirtualSysVariableMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSysVariableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSysVariableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDataType data_type获取
func (obj *_AllVirtualSysVariableMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllVirtualSysVariableMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllVirtualSysVariableMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithFlags flags获取
func (obj *_AllVirtualSysVariableMgr) WithFlags(flags int64) Option {
	return optionFunc(func(o *options) { o.query["flags"] = flags })
}

// WithMinVal min_val获取
func (obj *_AllVirtualSysVariableMgr) WithMinVal(minVal string) Option {
	return optionFunc(func(o *options) { o.query["min_val"] = minVal })
}

// WithMaxVal max_val获取
func (obj *_AllVirtualSysVariableMgr) WithMaxVal(maxVal string) Option {
	return optionFunc(func(o *options) { o.query["max_val"] = maxVal })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSysVariableMgr) GetByOption(opts ...Option) (result AllVirtualSysVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSysVariableMgr) GetByOptions(opts ...Option) (results []*AllVirtualSysVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSysVariableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSysVariable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where(options.query)
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
func (obj *_AllVirtualSysVariableMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromZone(zone string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromZone(zones []string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromName(name string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromName(names []string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromDataType(dataType int64) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromValue(value string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromValue(values []string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromInfo(info string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromFlags 通过flags获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromFlags(flags int64) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`flags` = ?", flags).Find(&results).Error

	return
}

// GetBatchFromFlags 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromFlags(flagss []int64) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`flags` IN (?)", flagss).Find(&results).Error

	return
}

// GetFromMinVal 通过min_val获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromMinVal(minVal string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`min_val` = ?", minVal).Find(&results).Error

	return
}

// GetBatchFromMinVal 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromMinVal(minVals []string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`min_val` IN (?)", minVals).Find(&results).Error

	return
}

// GetFromMaxVal 通过max_val获取内容
func (obj *_AllVirtualSysVariableMgr) GetFromMaxVal(maxVal string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`max_val` = ?", maxVal).Find(&results).Error

	return
}

// GetBatchFromMaxVal 批量查找
func (obj *_AllVirtualSysVariableMgr) GetBatchFromMaxVal(maxVals []string) (results []*AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`max_val` IN (?)", maxVals).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSysVariableMgr) FetchByPrimaryKey(tenantID int64, zone string, name string) (result AllVirtualSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysVariable{}).Where("`tenant_id` = ? AND `zone` = ? AND `name` = ?", tenantID, zone, name).First(&result).Error

	return
}
