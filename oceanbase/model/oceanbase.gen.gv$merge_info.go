package	model	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _Gv$mergeInfoMgr struct {
	*_BaseMgr
}

// Gv$mergeInfoMgr open func
func Gv$mergeInfoMgr(db *gorm.DB) *_Gv$mergeInfoMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$mergeInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$mergeInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$merge_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$mergeInfoMgr) GetTableName() string {
	return "gv$merge_info"
}

// Reset 重置gorm会话
func (obj *_Gv$mergeInfoMgr) Reset() *_Gv$mergeInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$mergeInfoMgr) Get() (result Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$mergeInfoMgr) Gets() (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$mergeInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_Gv$mergeInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$mergeInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTenantID TENANT_ID获取 
func (obj *_Gv$mergeInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTableID TABLE_ID获取 
func (obj *_Gv$mergeInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取 
func (obj *_Gv$mergeInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithType TYPE获取 
func (obj *_Gv$mergeInfoMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// WithAction ACTION获取 
func (obj *_Gv$mergeInfoMgr) WithAction(action string) Option {
	return optionFunc(func(o *options) { o.query["ACTION"] = action })
}

// WithVersion VERSION获取 
func (obj *_Gv$mergeInfoMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["VERSION"] = version })
}

// WithStartTime START_TIME获取 
func (obj *_Gv$mergeInfoMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithEndTime END_TIME获取 
func (obj *_Gv$mergeInfoMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["END_TIME"] = endTime })
}

// WithMacroBlockCount MACRO_BLOCK_COUNT获取 
func (obj *_Gv$mergeInfoMgr) WithMacroBlockCount(macroBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["MACRO_BLOCK_COUNT"] = macroBlockCount })
}

// WithReusePct REUSE_PCT获取 
func (obj *_Gv$mergeInfoMgr) WithReusePct(reusePct float64) Option {
	return optionFunc(func(o *options) { o.query["REUSE_PCT"] = reusePct })
}

// WithParallelDegree PARALLEL_DEGREE获取 
func (obj *_Gv$mergeInfoMgr) WithParallelDegree(parallelDegree int64) Option {
	return optionFunc(func(o *options) { o.query["PARALLEL_DEGREE"] = parallelDegree })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$mergeInfoMgr) GetByOption(opts ...Option) (result Gv$mergeInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$mergeInfoMgr) GetByOptions(opts ...Option) (results []*Gv$mergeInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$mergeInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$mergeInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where(options.query)
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


// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromSvrIP(svrIP string) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromSvrPort(svrPort int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromTenantID(tenantID int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromTableID(tableID int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过PARTITION_ID获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromPartitionID(partitionID int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromType 通过TYPE获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromType(_type string) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`TYPE` = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromType(_types []string) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`TYPE` IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromAction 通过ACTION获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromAction(action string) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`ACTION` = ?", action).Find(&results).Error
	
	return
}

// GetBatchFromAction 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromAction(actions []string) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`ACTION` IN (?)", actions).Find(&results).Error
	
	return
}
 
// GetFromVersion 通过VERSION获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromVersion(version string) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`VERSION` = ?", version).Find(&results).Error
	
	return
}

// GetBatchFromVersion 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromVersion(versions []string) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`VERSION` IN (?)", versions).Find(&results).Error
	
	return
}
 
// GetFromStartTime 通过START_TIME获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromStartTime(startTime time.Time) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`START_TIME` = ?", startTime).Find(&results).Error
	
	return
}

// GetBatchFromStartTime 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error
	
	return
}
 
// GetFromEndTime 通过END_TIME获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromEndTime(endTime time.Time) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`END_TIME` = ?", endTime).Find(&results).Error
	
	return
}

// GetBatchFromEndTime 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`END_TIME` IN (?)", endTimes).Find(&results).Error
	
	return
}
 
// GetFromMacroBlockCount 通过MACRO_BLOCK_COUNT获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromMacroBlockCount(macroBlockCount int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`MACRO_BLOCK_COUNT` = ?", macroBlockCount).Find(&results).Error
	
	return
}

// GetBatchFromMacroBlockCount 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromMacroBlockCount(macroBlockCounts []int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`MACRO_BLOCK_COUNT` IN (?)", macroBlockCounts).Find(&results).Error
	
	return
}
 
// GetFromReusePct 通过REUSE_PCT获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromReusePct(reusePct float64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`REUSE_PCT` = ?", reusePct).Find(&results).Error
	
	return
}

// GetBatchFromReusePct 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromReusePct(reusePcts []float64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`REUSE_PCT` IN (?)", reusePcts).Find(&results).Error
	
	return
}
 
// GetFromParallelDegree 通过PARALLEL_DEGREE获取内容  
func (obj *_Gv$mergeInfoMgr) GetFromParallelDegree(parallelDegree int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`PARALLEL_DEGREE` = ?", parallelDegree).Find(&results).Error
	
	return
}

// GetBatchFromParallelDegree 批量查找 
func (obj *_Gv$mergeInfoMgr) GetBatchFromParallelDegree(parallelDegrees []int64) (results []*Gv$mergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$mergeInfo{}).Where("`PARALLEL_DEGREE` IN (?)", parallelDegrees).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

