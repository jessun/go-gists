package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _V$obTransTableStatusMgr struct {
	*_BaseMgr
}

// V$obTransTableStatusMgr open func
func V$obTransTableStatusMgr(db *gorm.DB) *_V$obTransTableStatusMgr {
	if db == nil {
		panic(fmt.Errorf("V$obTransTableStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$obTransTableStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$ob_trans_table_status"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$obTransTableStatusMgr) GetTableName() string {
	return "v$ob_trans_table_status"
}

// Reset 重置gorm会话
func (obj *_V$obTransTableStatusMgr) Reset() *_V$obTransTableStatusMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$obTransTableStatusMgr) Get() (result V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$obTransTableStatusMgr) Gets() (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$obTransTableStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID TABLE_ID获取 
func (obj *_V$obTransTableStatusMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithPartitionID PARTITION_ID获取 
func (obj *_V$obTransTableStatusMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["PARTITION_ID"] = partitionID })
}

// WithEndLogID END_LOG_ID获取 
func (obj *_V$obTransTableStatusMgr) WithEndLogID(endLogID int64) Option {
	return optionFunc(func(o *options) { o.query["END_LOG_ID"] = endLogID })
}

// WithTransCnt TRANS_CNT获取 
func (obj *_V$obTransTableStatusMgr) WithTransCnt(transCnt int64) Option {
	return optionFunc(func(o *options) { o.query["TRANS_CNT"] = transCnt })
}


// GetByOption 功能选项模式获取
func (obj *_V$obTransTableStatusMgr) GetByOption(opts ...Option) (result V$obTransTableStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$obTransTableStatusMgr) GetByOptions(opts ...Option) (results []*V$obTransTableStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$obTransTableStatusMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$obTransTableStatus,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where(options.query)
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


// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_V$obTransTableStatusMgr) GetFromTableID(tableID int64) (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_V$obTransTableStatusMgr) GetBatchFromTableID(tableIDs []int64) (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过PARTITION_ID获取内容  
func (obj *_V$obTransTableStatusMgr) GetFromPartitionID(partitionID int64) (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where("`PARTITION_ID` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_V$obTransTableStatusMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where("`PARTITION_ID` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromEndLogID 通过END_LOG_ID获取内容  
func (obj *_V$obTransTableStatusMgr) GetFromEndLogID(endLogID int64) (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where("`END_LOG_ID` = ?", endLogID).Find(&results).Error
	
	return
}

// GetBatchFromEndLogID 批量查找 
func (obj *_V$obTransTableStatusMgr) GetBatchFromEndLogID(endLogIDs []int64) (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where("`END_LOG_ID` IN (?)", endLogIDs).Find(&results).Error
	
	return
}
 
// GetFromTransCnt 通过TRANS_CNT获取内容  
func (obj *_V$obTransTableStatusMgr) GetFromTransCnt(transCnt int64) (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where("`TRANS_CNT` = ?", transCnt).Find(&results).Error
	
	return
}

// GetBatchFromTransCnt 批量查找 
func (obj *_V$obTransTableStatusMgr) GetBatchFromTransCnt(transCnts []int64) (results []*V$obTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obTransTableStatus{}).Where("`TRANS_CNT` IN (?)", transCnts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

