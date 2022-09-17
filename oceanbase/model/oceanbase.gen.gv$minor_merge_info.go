package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _Gv$minorMergeInfoMgr struct {
	*_BaseMgr
}

// Gv$minorMergeInfoMgr open func
func Gv$minorMergeInfoMgr(db *gorm.DB) *_Gv$minorMergeInfoMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$minorMergeInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$minorMergeInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$minor_merge_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$minorMergeInfoMgr) GetTableName() string {
	return "gv$minor_merge_info"
}

// Reset 重置gorm会话
func (obj *_Gv$minorMergeInfoMgr) Reset() *_Gv$minorMergeInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$minorMergeInfoMgr) Get() (result Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$minorMergeInfoMgr) Gets() (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$minorMergeInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_Gv$minorMergeInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$minorMergeInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTenantID TENANT_ID获取 
func (obj *_Gv$minorMergeInfoMgr) WithTenantID(tenantID string) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithFreezeSnapshot FREEZE_SNAPSHOT获取 
func (obj *_Gv$minorMergeInfoMgr) WithFreezeSnapshot(freezeSnapshot string) Option {
	return optionFunc(func(o *options) { o.query["FREEZE_SNAPSHOT"] = freezeSnapshot })
}

// WithStartTime START_TIME获取 
func (obj *_Gv$minorMergeInfoMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["START_TIME"] = startTime })
}

// WithFinishTime FINISH_TIME获取 
func (obj *_Gv$minorMergeInfoMgr) WithFinishTime(finishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["FINISH_TIME"] = finishTime })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$minorMergeInfoMgr) GetByOption(opts ...Option) (result Gv$minorMergeInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$minorMergeInfoMgr) GetByOptions(opts ...Option) (results []*Gv$minorMergeInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$minorMergeInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$minorMergeInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where(options.query)
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
func (obj *_Gv$minorMergeInfoMgr) GetFromSvrIP(svrIP string) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$minorMergeInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$minorMergeInfoMgr) GetFromSvrPort(svrPort int64) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$minorMergeInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$minorMergeInfoMgr) GetFromTenantID(tenantID string) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$minorMergeInfoMgr) GetBatchFromTenantID(tenantIDs []string) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromFreezeSnapshot 通过FREEZE_SNAPSHOT获取内容  
func (obj *_Gv$minorMergeInfoMgr) GetFromFreezeSnapshot(freezeSnapshot string) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`FREEZE_SNAPSHOT` = ?", freezeSnapshot).Find(&results).Error
	
	return
}

// GetBatchFromFreezeSnapshot 批量查找 
func (obj *_Gv$minorMergeInfoMgr) GetBatchFromFreezeSnapshot(freezeSnapshots []string) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`FREEZE_SNAPSHOT` IN (?)", freezeSnapshots).Find(&results).Error
	
	return
}
 
// GetFromStartTime 通过START_TIME获取内容  
func (obj *_Gv$minorMergeInfoMgr) GetFromStartTime(startTime time.Time) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`START_TIME` = ?", startTime).Find(&results).Error
	
	return
}

// GetBatchFromStartTime 批量查找 
func (obj *_Gv$minorMergeInfoMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`START_TIME` IN (?)", startTimes).Find(&results).Error
	
	return
}
 
// GetFromFinishTime 通过FINISH_TIME获取内容  
func (obj *_Gv$minorMergeInfoMgr) GetFromFinishTime(finishTime time.Time) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`FINISH_TIME` = ?", finishTime).Find(&results).Error
	
	return
}

// GetBatchFromFinishTime 批量查找 
func (obj *_Gv$minorMergeInfoMgr) GetBatchFromFinishTime(finishTimes []time.Time) (results []*Gv$minorMergeInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$minorMergeInfo{}).Where("`FINISH_TIME` IN (?)", finishTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

