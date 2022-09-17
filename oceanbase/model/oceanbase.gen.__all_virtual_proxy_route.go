package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualProxyRouteMgr struct {
	*_BaseMgr
}

// AllVirtualProxyRouteMgr open func
func AllVirtualProxyRouteMgr(db *gorm.DB) *_AllVirtualProxyRouteMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualProxyRouteMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualProxyRouteMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_proxy_route"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualProxyRouteMgr) GetTableName() string {
	return "__all_virtual_proxy_route"
}

// Reset 重置gorm会话
func (obj *_AllVirtualProxyRouteMgr) Reset() *_AllVirtualProxyRouteMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualProxyRouteMgr) Get() (result AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualProxyRouteMgr) Gets() (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualProxyRouteMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSQLString sql_string获取
func (obj *_AllVirtualProxyRouteMgr) WithSQLString(sqlString string) Option {
	return optionFunc(func(o *options) { o.query["sql_string"] = sqlString })
}

// WithTenantName tenant_name获取
func (obj *_AllVirtualProxyRouteMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithDatabaseName database_name获取
func (obj *_AllVirtualProxyRouteMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithTableName table_name获取
func (obj *_AllVirtualProxyRouteMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithTableID table_id获取
func (obj *_AllVirtualProxyRouteMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithCalculatorBin calculator_bin获取
func (obj *_AllVirtualProxyRouteMgr) WithCalculatorBin(calculatorBin string) Option {
	return optionFunc(func(o *options) { o.query["calculator_bin"] = calculatorBin })
}

// WithResultStatus result_status获取
func (obj *_AllVirtualProxyRouteMgr) WithResultStatus(resultStatus int64) Option {
	return optionFunc(func(o *options) { o.query["result_status"] = resultStatus })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualProxyRouteMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualProxyRouteMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualProxyRouteMgr) WithSpare3(spare3 int64) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithSpare4 spare4获取
func (obj *_AllVirtualProxyRouteMgr) WithSpare4(spare4 string) Option {
	return optionFunc(func(o *options) { o.query["spare4"] = spare4 })
}

// WithSpare5 spare5获取
func (obj *_AllVirtualProxyRouteMgr) WithSpare5(spare5 string) Option {
	return optionFunc(func(o *options) { o.query["spare5"] = spare5 })
}

// WithSpare6 spare6获取
func (obj *_AllVirtualProxyRouteMgr) WithSpare6(spare6 string) Option {
	return optionFunc(func(o *options) { o.query["spare6"] = spare6 })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualProxyRouteMgr) GetByOption(opts ...Option) (result AllVirtualProxyRoute, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualProxyRouteMgr) GetByOptions(opts ...Option) (results []*AllVirtualProxyRoute, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualProxyRouteMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualProxyRoute, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where(options.query)
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

// GetFromSQLString 通过sql_string获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromSQLString(sqlString string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`sql_string` = ?", sqlString).Find(&results).Error

	return
}

// GetBatchFromSQLString 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromSQLString(sqlStrings []string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`sql_string` IN (?)", sqlStrings).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromTenantName(tenantName string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromDatabaseName(databaseName string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromTableName(tableName string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromTableName(tableNames []string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromTableID(tableID int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromCalculatorBin 通过calculator_bin获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromCalculatorBin(calculatorBin string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`calculator_bin` = ?", calculatorBin).Find(&results).Error

	return
}

// GetBatchFromCalculatorBin 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromCalculatorBin(calculatorBins []string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`calculator_bin` IN (?)", calculatorBins).Find(&results).Error

	return
}

// GetFromResultStatus 通过result_status获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromResultStatus(resultStatus int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`result_status` = ?", resultStatus).Find(&results).Error

	return
}

// GetBatchFromResultStatus 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromResultStatus(resultStatuss []int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`result_status` IN (?)", resultStatuss).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromSpare3(spare3 int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromSpare3(spare3s []int64) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromSpare4 通过spare4获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromSpare4(spare4 string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare4` = ?", spare4).Find(&results).Error

	return
}

// GetBatchFromSpare4 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromSpare4(spare4s []string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare4` IN (?)", spare4s).Find(&results).Error

	return
}

// GetFromSpare5 通过spare5获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromSpare5(spare5 string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare5` = ?", spare5).Find(&results).Error

	return
}

// GetBatchFromSpare5 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromSpare5(spare5s []string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare5` IN (?)", spare5s).Find(&results).Error

	return
}

// GetFromSpare6 通过spare6获取内容
func (obj *_AllVirtualProxyRouteMgr) GetFromSpare6(spare6 string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare6` = ?", spare6).Find(&results).Error

	return
}

// GetBatchFromSpare6 批量查找
func (obj *_AllVirtualProxyRouteMgr) GetBatchFromSpare6(spare6s []string) (results []*AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`spare6` IN (?)", spare6s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualProxyRouteMgr) FetchByPrimaryKey(sqlString string, tenantName string, databaseName string) (result AllVirtualProxyRoute, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyRoute{}).Where("`sql_string` = ? AND `tenant_name` = ? AND `database_name` = ?", sqlString, tenantName, databaseName).First(&result).Error

	return
}
