package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _AllImmediateEffectIndexSstableMgr struct {
	*_BaseMgr
}

// AllImmediateEffectIndexSstableMgr open func
func AllImmediateEffectIndexSstableMgr(db *gorm.DB) *_AllImmediateEffectIndexSstableMgr {
	if db == nil {
		panic(fmt.Errorf("AllImmediateEffectIndexSstableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllImmediateEffectIndexSstableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_immediate_effect_index_sstable"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllImmediateEffectIndexSstableMgr) GetTableName() string {
	return "__all_immediate_effect_index_sstable"
}

// Reset 重置gorm会话
func (obj *_AllImmediateEffectIndexSstableMgr) Reset() *_AllImmediateEffectIndexSstableMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AllImmediateEffectIndexSstableMgr) Get() (result AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AllImmediateEffectIndexSstableMgr) Gets() (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllImmediateEffectIndexSstableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIndexTableID index_table_id获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithIndexTableID(indexTableID int64) Option {
	return optionFunc(func(o *options) { o.query["index_table_id"] = indexTableID })
}

// WithPartitionID, partition_id,获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithPartitionID,(partitionID, int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id,"] = partitionID, })
}

// WithSvrIP svr_ip获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSnapshot snapshot获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithSnapshot(snapshot int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot"] = snapshot })
}

// WithPartitionCnt partition_cnt获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithDataSize data_size获取 
func (obj *_AllImmediateEffectIndexSstableMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}


// GetByOption 功能选项模式获取
func (obj *_AllImmediateEffectIndexSstableMgr) GetByOption(opts ...Option) (result AllImmediateEffectIndexSstable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllImmediateEffectIndexSstableMgr) GetByOptions(opts ...Option) (results []*AllImmediateEffectIndexSstable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AllImmediateEffectIndexSstableMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllImmediateEffectIndexSstable,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where(options.query)
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
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error
	
	return
}

// GetBatchFromGmtCreate 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error
	
	return
}
 
// GetFromGmtModified 通过gmt_modified获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error
	
	return
}

// GetBatchFromGmtModified 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromTenantID(tenantID int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromIndexTableID 通过index_table_id获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromIndexTableID(indexTableID int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`index_table_id` = ?", indexTableID).Find(&results).Error
	
	return
}

// GetBatchFromIndexTableID 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromIndexTableID(indexTableIDs []int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`index_table_id` IN (?)", indexTableIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID, 通过partition_id,获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromPartitionID,(partitionID, int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`partition_id,` = ?", partitionID,).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID, 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromPartitionID,(partitionID,s []int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`partition_id,` IN (?)", partitionID,s).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromSvrIP(svrIP string) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromSvrPort(svrPort int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromSnapshot 通过snapshot获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromSnapshot(snapshot int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`snapshot` = ?", snapshot).Find(&results).Error
	
	return
}

// GetBatchFromSnapshot 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromSnapshot(snapshots []int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`snapshot` IN (?)", snapshots).Find(&results).Error
	
	return
}
 
// GetFromPartitionCnt 通过partition_cnt获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error
	
	return
}

// GetBatchFromPartitionCnt 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error
	
	return
}
 
// GetFromDataSize 通过data_size获取内容  
func (obj *_AllImmediateEffectIndexSstableMgr) GetFromDataSize(dataSize int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`data_size` = ?", dataSize).Find(&results).Error
	
	return
}

// GetBatchFromDataSize 批量查找 
func (obj *_AllImmediateEffectIndexSstableMgr) GetBatchFromDataSize(dataSizes []int64) (results []*AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AllImmediateEffectIndexSstableMgr) FetchByPrimaryKey(tenantID int64 ,indexTableID int64 ,partitionID, int64 ,svrIP string ,svrPort int64 ) (result AllImmediateEffectIndexSstable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllImmediateEffectIndexSstable{}).Where("`tenant_id` = ? AND `index_table_id` = ? AND `partition_id,` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID , indexTableID , partitionID, , svrIP , svrPort).First(&result).Error
	
	return
}
 

 

	

