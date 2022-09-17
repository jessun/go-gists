package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$obTransTableStatusMgr struct {
	*_BaseMgr
}

// Gv$obTransTableStatusMgr open func
func Gv$obTransTableStatusMgr(db *gorm.DB) *_Gv$obTransTableStatusMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$obTransTableStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$obTransTableStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$ob_trans_table_status"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$obTransTableStatusMgr) GetTableName() string {
	return "gv$ob_trans_table_status"
}

// Reset 重置gorm会话
func (obj *_Gv$obTransTableStatusMgr) Reset() *_Gv$obTransTableStatusMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$obTransTableStatusMgr) Get() (result Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$obTransTableStatusMgr) Gets() (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$obTransTableStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_Gv$obTransTableStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$obTransTableStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTenantID TENANT_ID获取 
func (obj *_Gv$obTransTableStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTableID TABLE_ID获取 
func (obj *_Gv$obTransTableStatusMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取 
func (obj *_Gv$obTransTableStatusMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithEndLogID END_LOG_ID获取 
func (obj *_Gv$obTransTableStatusMgr) WithEndLogID(endLogID int64) Option {
	return optionFunc(func(o *options) { o.query["END_LOG_ID"] = endLogID })
}

// WithTransCnt TRANS_CNT获取 
func (obj *_Gv$obTransTableStatusMgr) WithTransCnt(transCnt int64) Option {
	return optionFunc(func(o *options) { o.query["TRANS_CNT"] = transCnt })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$obTransTableStatusMgr) GetByOption(opts ...Option) (result Gv$obTransTableStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$obTransTableStatusMgr) GetByOptions(opts ...Option) (results []*Gv$obTransTableStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$obTransTableStatusMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$obTransTableStatus,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where(options.query)
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
func (obj *_Gv$obTransTableStatusMgr) GetFromSvrIP(svrIP string) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$obTransTableStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$obTransTableStatusMgr) GetFromSvrPort(svrPort int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$obTransTableStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$obTransTableStatusMgr) GetFromTenantID(tenantID int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$obTransTableStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_Gv$obTransTableStatusMgr) GetFromTableID(tableID int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_Gv$obTransTableStatusMgr) GetBatchFromTableID(tableIDs []int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过PARTITION_ID获取内容  
func (obj *_Gv$obTransTableStatusMgr) GetFromPartitionID(partitionID int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_Gv$obTransTableStatusMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromEndLogID 通过END_LOG_ID获取内容  
func (obj *_Gv$obTransTableStatusMgr) GetFromEndLogID(endLogID int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`END_LOG_ID` = ?", endLogID).Find(&results).Error
	
	return
}

// GetBatchFromEndLogID 批量查找 
func (obj *_Gv$obTransTableStatusMgr) GetBatchFromEndLogID(endLogIDs []int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`END_LOG_ID` IN (?)", endLogIDs).Find(&results).Error
	
	return
}
 
// GetFromTransCnt 通过TRANS_CNT获取内容  
func (obj *_Gv$obTransTableStatusMgr) GetFromTransCnt(transCnt int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`TRANS_CNT` = ?", transCnt).Find(&results).Error
	
	return
}

// GetBatchFromTransCnt 批量查找 
func (obj *_Gv$obTransTableStatusMgr) GetBatchFromTransCnt(transCnts []int64) (results []*Gv$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obTransTableStatus{}).Where("`TRANS_CNT` IN (?)", transCnts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

