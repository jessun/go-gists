package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllSysVariableMgr struct {
	*_BaseMgr
}

// AllSysVariableMgr open func
func AllSysVariableMgr(db *gorm.DB) *_AllSysVariableMgr {
	if db == nil {
		panic(fmt.Errorf("AllSysVariableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSysVariableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sys_variable"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSysVariableMgr) GetTableName() string {
	return "__all_sys_variable"
}

// Reset 重置gorm会话
func (obj *_AllSysVariableMgr) Reset() *_AllSysVariableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSysVariableMgr) Get() (result AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSysVariableMgr) Gets() (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSysVariableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSysVariableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSysVariableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSysVariableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllSysVariableMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithName name获取
func (obj *_AllSysVariableMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDataType data_type获取
func (obj *_AllSysVariableMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllSysVariableMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllSysVariableMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithFlags flags获取
func (obj *_AllSysVariableMgr) WithFlags(flags int64) Option {
	return optionFunc(func(o *options) { o.query["flags"] = flags })
}

// WithMinVal min_val获取
func (obj *_AllSysVariableMgr) WithMinVal(minVal string) Option {
	return optionFunc(func(o *options) { o.query["min_val"] = minVal })
}

// WithMaxVal max_val获取
func (obj *_AllSysVariableMgr) WithMaxVal(maxVal string) Option {
	return optionFunc(func(o *options) { o.query["max_val"] = maxVal })
}

// GetByOption 功能选项模式获取
func (obj *_AllSysVariableMgr) GetByOption(opts ...Option) (result AllSysVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSysVariableMgr) GetByOptions(opts ...Option) (results []*AllSysVariable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSysVariableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSysVariable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where(options.query)
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
func (obj *_AllSysVariableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSysVariableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSysVariableMgr) GetFromTenantID(tenantID int64) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllSysVariableMgr) GetFromZone(zone string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromZone(zones []string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllSysVariableMgr) GetFromName(name string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromName(names []string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllSysVariableMgr) GetFromDataType(dataType int64) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllSysVariableMgr) GetFromValue(value string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromValue(values []string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllSysVariableMgr) GetFromInfo(info string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromInfo(infos []string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromFlags 通过flags获取内容
func (obj *_AllSysVariableMgr) GetFromFlags(flags int64) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`flags` = ?", flags).Find(&results).Error

	return
}

// GetBatchFromFlags 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromFlags(flagss []int64) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`flags` IN (?)", flagss).Find(&results).Error

	return
}

// GetFromMinVal 通过min_val获取内容
func (obj *_AllSysVariableMgr) GetFromMinVal(minVal string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`min_val` = ?", minVal).Find(&results).Error

	return
}

// GetBatchFromMinVal 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromMinVal(minVals []string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`min_val` IN (?)", minVals).Find(&results).Error

	return
}

// GetFromMaxVal 通过max_val获取内容
func (obj *_AllSysVariableMgr) GetFromMaxVal(maxVal string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`max_val` = ?", maxVal).Find(&results).Error

	return
}

// GetBatchFromMaxVal 批量查找
func (obj *_AllSysVariableMgr) GetBatchFromMaxVal(maxVals []string) (results []*AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`max_val` IN (?)", maxVals).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSysVariableMgr) FetchByPrimaryKey(tenantID int64, zone string, name string) (result AllSysVariable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysVariable{}).Where("`tenant_id` = ? AND `zone` = ? AND `name` = ?", tenantID, zone, name).First(&result).Error

	return
}
