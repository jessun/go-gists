package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllRootserviceJobMgr struct {
	*_BaseMgr
}

// AllRootserviceJobMgr open func
func AllRootserviceJobMgr(db *gorm.DB) *_AllRootserviceJobMgr {
	if db == nil {
		panic(fmt.Errorf("AllRootserviceJobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllRootserviceJobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_rootservice_job"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllRootserviceJobMgr) GetTableName() string {
	return "__all_rootservice_job"
}

// Reset 重置gorm会话
func (obj *_AllRootserviceJobMgr) Reset() *_AllRootserviceJobMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllRootserviceJobMgr) Get() (result AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllRootserviceJobMgr) Gets() (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllRootserviceJobMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllRootserviceJobMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllRootserviceJobMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithJobID job_id获取
func (obj *_AllRootserviceJobMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithJobType job_type获取
func (obj *_AllRootserviceJobMgr) WithJobType(jobType string) Option {
	return optionFunc(func(o *options) { o.query["job_type"] = jobType })
}

// WithJobStatus job_status获取
func (obj *_AllRootserviceJobMgr) WithJobStatus(jobStatus string) Option {
	return optionFunc(func(o *options) { o.query["job_status"] = jobStatus })
}

// WithReturnCode return_code获取
func (obj *_AllRootserviceJobMgr) WithReturnCode(returnCode int64) Option {
	return optionFunc(func(o *options) { o.query["return_code"] = returnCode })
}

// WithProgress progress获取
func (obj *_AllRootserviceJobMgr) WithProgress(progress int64) Option {
	return optionFunc(func(o *options) { o.query["progress"] = progress })
}

// WithTenantID tenant_id获取
func (obj *_AllRootserviceJobMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTenantName tenant_name获取
func (obj *_AllRootserviceJobMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithDatabaseID database_id获取
func (obj *_AllRootserviceJobMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithDatabaseName database_name获取
func (obj *_AllRootserviceJobMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithTableID table_id获取
func (obj *_AllRootserviceJobMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTableName table_name获取
func (obj *_AllRootserviceJobMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithPartitionID partition_id获取
func (obj *_AllRootserviceJobMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllRootserviceJobMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllRootserviceJobMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithUnitID unit_id获取
func (obj *_AllRootserviceJobMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithRsSvrIP rs_svr_ip获取
func (obj *_AllRootserviceJobMgr) WithRsSvrIP(rsSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["rs_svr_ip"] = rsSvrIP })
}

// WithRsSvrPort rs_svr_port获取
func (obj *_AllRootserviceJobMgr) WithRsSvrPort(rsSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["rs_svr_port"] = rsSvrPort })
}

// WithSQLText sql_text获取
func (obj *_AllRootserviceJobMgr) WithSQLText(sqlText string) Option {
	return optionFunc(func(o *options) { o.query["sql_text"] = sqlText })
}

// WithExtraInfo extra_info获取
func (obj *_AllRootserviceJobMgr) WithExtraInfo(extraInfo string) Option {
	return optionFunc(func(o *options) { o.query["extra_info"] = extraInfo })
}

// WithResourcePoolID resource_pool_id获取
func (obj *_AllRootserviceJobMgr) WithResourcePoolID(resourcePoolID int64) Option {
	return optionFunc(func(o *options) { o.query["resource_pool_id"] = resourcePoolID })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllRootserviceJobMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithTablegroupName tablegroup_name获取
func (obj *_AllRootserviceJobMgr) WithTablegroupName(tablegroupName string) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_name"] = tablegroupName })
}

// GetByOption 功能选项模式获取
func (obj *_AllRootserviceJobMgr) GetByOption(opts ...Option) (result AllRootserviceJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllRootserviceJobMgr) GetByOptions(opts ...Option) (results []*AllRootserviceJob, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllRootserviceJobMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllRootserviceJob, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where(options.query)
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
func (obj *_AllRootserviceJobMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllRootserviceJobMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容
func (obj *_AllRootserviceJobMgr) GetFromJobID(jobID int64) (result AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`job_id` = ?", jobID).First(&result).Error

	return
}

// GetBatchFromJobID 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromJobID(jobIDs []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromJobType 通过job_type获取内容
func (obj *_AllRootserviceJobMgr) GetFromJobType(jobType string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`job_type` = ?", jobType).Find(&results).Error

	return
}

// GetBatchFromJobType 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromJobType(jobTypes []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`job_type` IN (?)", jobTypes).Find(&results).Error

	return
}

// GetFromJobStatus 通过job_status获取内容
func (obj *_AllRootserviceJobMgr) GetFromJobStatus(jobStatus string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`job_status` = ?", jobStatus).Find(&results).Error

	return
}

// GetBatchFromJobStatus 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromJobStatus(jobStatuss []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`job_status` IN (?)", jobStatuss).Find(&results).Error

	return
}

// GetFromReturnCode 通过return_code获取内容
func (obj *_AllRootserviceJobMgr) GetFromReturnCode(returnCode int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`return_code` = ?", returnCode).Find(&results).Error

	return
}

// GetBatchFromReturnCode 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromReturnCode(returnCodes []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`return_code` IN (?)", returnCodes).Find(&results).Error

	return
}

// GetFromProgress 通过progress获取内容
func (obj *_AllRootserviceJobMgr) GetFromProgress(progress int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`progress` = ?", progress).Find(&results).Error

	return
}

// GetBatchFromProgress 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromProgress(progresss []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`progress` IN (?)", progresss).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllRootserviceJobMgr) GetFromTenantID(tenantID int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllRootserviceJobMgr) GetFromTenantName(tenantName string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllRootserviceJobMgr) GetFromDatabaseID(databaseID int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllRootserviceJobMgr) GetFromDatabaseName(databaseName string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllRootserviceJobMgr) GetFromTableID(tableID int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllRootserviceJobMgr) GetFromTableName(tableName string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromTableName(tableNames []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllRootserviceJobMgr) GetFromPartitionID(partitionID int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllRootserviceJobMgr) GetFromSvrIP(svrIP string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllRootserviceJobMgr) GetFromSvrPort(svrPort int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllRootserviceJobMgr) GetFromUnitID(unitID int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromRsSvrIP 通过rs_svr_ip获取内容
func (obj *_AllRootserviceJobMgr) GetFromRsSvrIP(rsSvrIP string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`rs_svr_ip` = ?", rsSvrIP).Find(&results).Error

	return
}

// GetBatchFromRsSvrIP 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromRsSvrIP(rsSvrIPs []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`rs_svr_ip` IN (?)", rsSvrIPs).Find(&results).Error

	return
}

// GetFromRsSvrPort 通过rs_svr_port获取内容
func (obj *_AllRootserviceJobMgr) GetFromRsSvrPort(rsSvrPort int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`rs_svr_port` = ?", rsSvrPort).Find(&results).Error

	return
}

// GetBatchFromRsSvrPort 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromRsSvrPort(rsSvrPorts []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`rs_svr_port` IN (?)", rsSvrPorts).Find(&results).Error

	return
}

// GetFromSQLText 通过sql_text获取内容
func (obj *_AllRootserviceJobMgr) GetFromSQLText(sqlText string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`sql_text` = ?", sqlText).Find(&results).Error

	return
}

// GetBatchFromSQLText 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromSQLText(sqlTexts []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`sql_text` IN (?)", sqlTexts).Find(&results).Error

	return
}

// GetFromExtraInfo 通过extra_info获取内容
func (obj *_AllRootserviceJobMgr) GetFromExtraInfo(extraInfo string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`extra_info` = ?", extraInfo).Find(&results).Error

	return
}

// GetBatchFromExtraInfo 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromExtraInfo(extraInfos []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`extra_info` IN (?)", extraInfos).Find(&results).Error

	return
}

// GetFromResourcePoolID 通过resource_pool_id获取内容
func (obj *_AllRootserviceJobMgr) GetFromResourcePoolID(resourcePoolID int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`resource_pool_id` = ?", resourcePoolID).Find(&results).Error

	return
}

// GetBatchFromResourcePoolID 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromResourcePoolID(resourcePoolIDs []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`resource_pool_id` IN (?)", resourcePoolIDs).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllRootserviceJobMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromTablegroupName 通过tablegroup_name获取内容
func (obj *_AllRootserviceJobMgr) GetFromTablegroupName(tablegroupName string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`tablegroup_name` = ?", tablegroupName).Find(&results).Error

	return
}

// GetBatchFromTablegroupName 批量查找
func (obj *_AllRootserviceJobMgr) GetBatchFromTablegroupName(tablegroupNames []string) (results []*AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`tablegroup_name` IN (?)", tablegroupNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllRootserviceJobMgr) FetchByPrimaryKey(jobID int64) (result AllRootserviceJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceJob{}).Where("`job_id` = ?", jobID).First(&result).Error

	return
}
