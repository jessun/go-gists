package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualReplicaTaskMgr struct {
	*_BaseMgr
}

// AllVirtualReplicaTaskMgr open func
func AllVirtualReplicaTaskMgr(db *gorm.DB) *_AllVirtualReplicaTaskMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualReplicaTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualReplicaTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_replica_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualReplicaTaskMgr) GetTableName() string {
	return "__all_virtual_replica_task"
}

// Reset 重置gorm会话
func (obj *_AllVirtualReplicaTaskMgr) Reset() *_AllVirtualReplicaTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualReplicaTaskMgr) Get() (result AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualReplicaTaskMgr) Gets() (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualReplicaTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualReplicaTaskMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualReplicaTaskMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualReplicaTaskMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSrcIP src_ip获取
func (obj *_AllVirtualReplicaTaskMgr) WithSrcIP(srcIP string) Option {
	return optionFunc(func(o *options) { o.query["src_ip"] = srcIP })
}

// WithSrcPort src_port获取
func (obj *_AllVirtualReplicaTaskMgr) WithSrcPort(srcPort int64) Option {
	return optionFunc(func(o *options) { o.query["src_port"] = srcPort })
}

// WithSrcReplicaType src_replica_type获取
func (obj *_AllVirtualReplicaTaskMgr) WithSrcReplicaType(srcReplicaType int64) Option {
	return optionFunc(func(o *options) { o.query["src_replica_type"] = srcReplicaType })
}

// WithZone zone获取
func (obj *_AllVirtualReplicaTaskMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRegion region获取
func (obj *_AllVirtualReplicaTaskMgr) WithRegion(region string) Option {
	return optionFunc(func(o *options) { o.query["region"] = region })
}

// WithDstIP dst_ip获取
func (obj *_AllVirtualReplicaTaskMgr) WithDstIP(dstIP string) Option {
	return optionFunc(func(o *options) { o.query["dst_ip"] = dstIP })
}

// WithDstPort dst_port获取
func (obj *_AllVirtualReplicaTaskMgr) WithDstPort(dstPort int64) Option {
	return optionFunc(func(o *options) { o.query["dst_port"] = dstPort })
}

// WithDstReplicaType dst_replica_type获取
func (obj *_AllVirtualReplicaTaskMgr) WithDstReplicaType(dstReplicaType int64) Option {
	return optionFunc(func(o *options) { o.query["dst_replica_type"] = dstReplicaType })
}

// WithCmdType cmd_type获取
func (obj *_AllVirtualReplicaTaskMgr) WithCmdType(cmdType string) Option {
	return optionFunc(func(o *options) { o.query["cmd_type"] = cmdType })
}

// WithComment comment获取
func (obj *_AllVirtualReplicaTaskMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualReplicaTaskMgr) GetByOption(opts ...Option) (result AllVirtualReplicaTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualReplicaTaskMgr) GetByOptions(opts ...Option) (results []*AllVirtualReplicaTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualReplicaTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualReplicaTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where(options.query)
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
func (obj *_AllVirtualReplicaTaskMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromTableID(tableID int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSrcIP 通过src_ip获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromSrcIP(srcIP string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`src_ip` = ?", srcIP).Find(&results).Error

	return
}

// GetBatchFromSrcIP 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromSrcIP(srcIPs []string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`src_ip` IN (?)", srcIPs).Find(&results).Error

	return
}

// GetFromSrcPort 通过src_port获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromSrcPort(srcPort int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`src_port` = ?", srcPort).Find(&results).Error

	return
}

// GetBatchFromSrcPort 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromSrcPort(srcPorts []int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`src_port` IN (?)", srcPorts).Find(&results).Error

	return
}

// GetFromSrcReplicaType 通过src_replica_type获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromSrcReplicaType(srcReplicaType int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`src_replica_type` = ?", srcReplicaType).Find(&results).Error

	return
}

// GetBatchFromSrcReplicaType 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromSrcReplicaType(srcReplicaTypes []int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`src_replica_type` IN (?)", srcReplicaTypes).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromZone(zone string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromZone(zones []string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromRegion 通过region获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromRegion(region string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`region` = ?", region).Find(&results).Error

	return
}

// GetBatchFromRegion 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromRegion(regions []string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`region` IN (?)", regions).Find(&results).Error

	return
}

// GetFromDstIP 通过dst_ip获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromDstIP(dstIP string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`dst_ip` = ?", dstIP).Find(&results).Error

	return
}

// GetBatchFromDstIP 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromDstIP(dstIPs []string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`dst_ip` IN (?)", dstIPs).Find(&results).Error

	return
}

// GetFromDstPort 通过dst_port获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromDstPort(dstPort int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`dst_port` = ?", dstPort).Find(&results).Error

	return
}

// GetBatchFromDstPort 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromDstPort(dstPorts []int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`dst_port` IN (?)", dstPorts).Find(&results).Error

	return
}

// GetFromDstReplicaType 通过dst_replica_type获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromDstReplicaType(dstReplicaType int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`dst_replica_type` = ?", dstReplicaType).Find(&results).Error

	return
}

// GetBatchFromDstReplicaType 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromDstReplicaType(dstReplicaTypes []int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`dst_replica_type` IN (?)", dstReplicaTypes).Find(&results).Error

	return
}

// GetFromCmdType 通过cmd_type获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromCmdType(cmdType string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`cmd_type` = ?", cmdType).Find(&results).Error

	return
}

// GetBatchFromCmdType 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromCmdType(cmdTypes []string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`cmd_type` IN (?)", cmdTypes).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualReplicaTaskMgr) GetFromComment(comment string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualReplicaTaskMgr) GetBatchFromComment(comments []string) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByAllVirtualReplicaTaskI1  获取多个内容
func (obj *_AllVirtualReplicaTaskMgr) FetchIndexByAllVirtualReplicaTaskI1(tenantID int64) (results []*AllVirtualReplicaTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReplicaTask{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}
