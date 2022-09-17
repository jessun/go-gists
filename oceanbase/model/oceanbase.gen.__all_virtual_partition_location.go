package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPartitionLocationMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionLocationMgr open func
func AllVirtualPartitionLocationMgr(db *gorm.DB) *_AllVirtualPartitionLocationMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionLocationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionLocationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_location"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionLocationMgr) GetTableName() string {
	return "__all_virtual_partition_location"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionLocationMgr) Reset() *_AllVirtualPartitionLocationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionLocationMgr) Get() (result AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionLocationMgr) Gets() (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionLocationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionLocationMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionLocationMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPartitionLocationMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionLocationMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionLocationMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithUnitID unit_id获取
func (obj *_AllVirtualPartitionLocationMgr) WithUnitID(unitID int64) Option {
	return optionFunc(func(o *options) { o.query["unit_id"] = unitID })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualPartitionLocationMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithZone zone获取
func (obj *_AllVirtualPartitionLocationMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRole role获取
func (obj *_AllVirtualPartitionLocationMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithMemberList member_list获取
func (obj *_AllVirtualPartitionLocationMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualPartitionLocationMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithStatus status获取
func (obj *_AllVirtualPartitionLocationMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDataVersion data_version获取
func (obj *_AllVirtualPartitionLocationMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionLocationMgr) GetByOption(opts ...Option) (result AllVirtualPartitionLocation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionLocationMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionLocation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionLocationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionLocation, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where(options.query)
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
func (obj *_AllVirtualPartitionLocationMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromUnitID 通过unit_id获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromUnitID(unitID int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`unit_id` = ?", unitID).Find(&results).Error

	return
}

// GetBatchFromUnitID 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromUnitID(unitIDs []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`unit_id` IN (?)", unitIDs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromZone(zone string) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromZone(zones []string) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromRole(role int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromMemberList(memberList string) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromMemberList(memberLists []string) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromStatus(status string) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllVirtualPartitionLocationMgr) GetFromDataVersion(dataVersion int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllVirtualPartitionLocationMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPartitionLocationMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, svrIP string, svrPort int64) (result AllVirtualPartitionLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionLocation{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, tableID, partitionID, svrIP, svrPort).First(&result).Error

	return
}
