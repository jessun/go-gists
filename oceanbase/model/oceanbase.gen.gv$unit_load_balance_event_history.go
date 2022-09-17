package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _Gv$unitLoadBalanceEventHistoryMgr struct {
	*_BaseMgr
}

// Gv$unitLoadBalanceEventHistoryMgr open func
func Gv$unitLoadBalanceEventHistoryMgr(db *gorm.DB) *_Gv$unitLoadBalanceEventHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$unitLoadBalanceEventHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$unitLoadBalanceEventHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$unit_load_balance_event_history"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetTableName() string {
	return "gv$unit_load_balance_event_history"
}

// Reset 重置gorm会话
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) Reset() *_Gv$unitLoadBalanceEventHistoryMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) Get() (result Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) Gets() (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithZone zone获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithTableID table_id获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithTableID(tableID string) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithPartitionID(partitionID string) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithDataSize data_size获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithDataSize(dataSize string) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithReplicaType replica_type获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithReplicaType(replicaType string) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithSrcIP src_ip获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithSrcIP(srcIP string) Option {
	return optionFunc(func(o *options) { o.query["src_ip"] = srcIP })
}

// WithSrcPort src_port获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithSrcPort(srcPort string) Option {
	return optionFunc(func(o *options) { o.query["src_port"] = srcPort })
}

// WithDestUnitID dest_unit_id获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithDestUnitID(destUnitID string) Option {
	return optionFunc(func(o *options) { o.query["dest_unit_id"] = destUnitID })
}

// WithDestIP dest_ip获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithDestIP(destIP string) Option {
	return optionFunc(func(o *options) { o.query["dest_ip"] = destIP })
}

// WithDestPort dest_port获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithDestPort(destPort string) Option {
	return optionFunc(func(o *options) { o.query["dest_port"] = destPort })
}

// WithDataSrcIP data_src_ip获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithDataSrcIP(dataSrcIP string) Option {
	return optionFunc(func(o *options) { o.query["data_src_ip"] = dataSrcIP })
}

// WithDataSrcPort data_src_port获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithDataSrcPort(dataSrcPort string) Option {
	return optionFunc(func(o *options) { o.query["data_src_port"] = dataSrcPort })
}

// WithResultCode result_code获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithResultCode(resultCode string) Option {
	return optionFunc(func(o *options) { o.query["result_code"] = resultCode })
}

// WithResultStr result_str获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithResultStr(resultStr string) Option {
	return optionFunc(func(o *options) { o.query["result_str"] = resultStr })
}

// WithComment comment获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithRsSvrIP rs_svr_ip获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithRsSvrIP(rsSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["rs_svr_ip"] = rsSvrIP })
}

// WithRsSvrPort rs_svr_port获取 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) WithRsSvrPort(rsSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["rs_svr_port"] = rsSvrPort })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetByOption(opts ...Option) (result Gv$unitLoadBalanceEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetByOptions(opts ...Option) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$unitLoadBalanceEventHistory,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where(options.query)
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
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error
	
	return
}

// GetBatchFromGmtCreate 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error
	
	return
}
 
// GetFromZone 通过zone获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromZone(zone string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`zone` = ?", zone).Find(&results).Error
	
	return
}

// GetBatchFromZone 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromZone(zones []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`zone` IN (?)", zones).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过table_id获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromTableID(tableID string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromTableID(tableIDs []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过partition_id获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromPartitionID(partitionID string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`partition_id` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromPartitionID(partitionIDs []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromDataSize 通过data_size获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromDataSize(dataSize string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`data_size` = ?", dataSize).Find(&results).Error
	
	return
}

// GetBatchFromDataSize 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromDataSize(dataSizes []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error
	
	return
}
 
// GetFromReplicaType 通过replica_type获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromReplicaType(replicaType string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`replica_type` = ?", replicaType).Find(&results).Error
	
	return
}

// GetBatchFromReplicaType 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromReplicaType(replicaTypes []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error
	
	return
}
 
// GetFromSrcIP 通过src_ip获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromSrcIP(srcIP string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`src_ip` = ?", srcIP).Find(&results).Error
	
	return
}

// GetBatchFromSrcIP 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromSrcIP(srcIPs []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`src_ip` IN (?)", srcIPs).Find(&results).Error
	
	return
}
 
// GetFromSrcPort 通过src_port获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromSrcPort(srcPort string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`src_port` = ?", srcPort).Find(&results).Error
	
	return
}

// GetBatchFromSrcPort 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromSrcPort(srcPorts []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`src_port` IN (?)", srcPorts).Find(&results).Error
	
	return
}
 
// GetFromDestUnitID 通过dest_unit_id获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromDestUnitID(destUnitID string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`dest_unit_id` = ?", destUnitID).Find(&results).Error
	
	return
}

// GetBatchFromDestUnitID 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromDestUnitID(destUnitIDs []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`dest_unit_id` IN (?)", destUnitIDs).Find(&results).Error
	
	return
}
 
// GetFromDestIP 通过dest_ip获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromDestIP(destIP string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`dest_ip` = ?", destIP).Find(&results).Error
	
	return
}

// GetBatchFromDestIP 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromDestIP(destIPs []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`dest_ip` IN (?)", destIPs).Find(&results).Error
	
	return
}
 
// GetFromDestPort 通过dest_port获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromDestPort(destPort string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`dest_port` = ?", destPort).Find(&results).Error
	
	return
}

// GetBatchFromDestPort 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromDestPort(destPorts []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`dest_port` IN (?)", destPorts).Find(&results).Error
	
	return
}
 
// GetFromDataSrcIP 通过data_src_ip获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromDataSrcIP(dataSrcIP string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`data_src_ip` = ?", dataSrcIP).Find(&results).Error
	
	return
}

// GetBatchFromDataSrcIP 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromDataSrcIP(dataSrcIPs []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`data_src_ip` IN (?)", dataSrcIPs).Find(&results).Error
	
	return
}
 
// GetFromDataSrcPort 通过data_src_port获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromDataSrcPort(dataSrcPort string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`data_src_port` = ?", dataSrcPort).Find(&results).Error
	
	return
}

// GetBatchFromDataSrcPort 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromDataSrcPort(dataSrcPorts []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`data_src_port` IN (?)", dataSrcPorts).Find(&results).Error
	
	return
}
 
// GetFromResultCode 通过result_code获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromResultCode(resultCode string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`result_code` = ?", resultCode).Find(&results).Error
	
	return
}

// GetBatchFromResultCode 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromResultCode(resultCodes []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`result_code` IN (?)", resultCodes).Find(&results).Error
	
	return
}
 
// GetFromResultStr 通过result_str获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromResultStr(resultStr string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`result_str` = ?", resultStr).Find(&results).Error
	
	return
}

// GetBatchFromResultStr 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromResultStr(resultStrs []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`result_str` IN (?)", resultStrs).Find(&results).Error
	
	return
}
 
// GetFromComment 通过comment获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromComment(comment string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`comment` = ?", comment).Find(&results).Error
	
	return
}

// GetBatchFromComment 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromComment(comments []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error
	
	return
}
 
// GetFromRsSvrIP 通过rs_svr_ip获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromRsSvrIP(rsSvrIP string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`rs_svr_ip` = ?", rsSvrIP).Find(&results).Error
	
	return
}

// GetBatchFromRsSvrIP 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromRsSvrIP(rsSvrIPs []string) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`rs_svr_ip` IN (?)", rsSvrIPs).Find(&results).Error
	
	return
}
 
// GetFromRsSvrPort 通过rs_svr_port获取内容  
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetFromRsSvrPort(rsSvrPort int64) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`rs_svr_port` = ?", rsSvrPort).Find(&results).Error
	
	return
}

// GetBatchFromRsSvrPort 批量查找 
func (obj *_Gv$unitLoadBalanceEventHistoryMgr) GetBatchFromRsSvrPort(rsSvrPorts []int64) (results []*Gv$unitLoadBalanceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$unitLoadBalanceEventHistory{}).Where("`rs_svr_port` IN (?)", rsSvrPorts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

