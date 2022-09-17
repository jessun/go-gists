package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _V$obClusterMgr struct {
	*_BaseMgr
}

// V$obClusterMgr open func
func V$obClusterMgr(db *gorm.DB) *_V$obClusterMgr {
	if db == nil {
		panic(fmt.Errorf("V$obClusterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$obClusterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$ob_cluster"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$obClusterMgr) GetTableName() string {
	return "v$ob_cluster"
}

// Reset 重置gorm会话
func (obj *_V$obClusterMgr) Reset() *_V$obClusterMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$obClusterMgr) Get() (result V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$obClusterMgr) Gets() (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$obClusterMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithClusterID cluster_id获取 
func (obj *_V$obClusterMgr) WithClusterID(clusterID int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_id"] = clusterID })
}

// WithClusterName cluster_name获取 
func (obj *_V$obClusterMgr) WithClusterName(clusterName string) Option {
	return optionFunc(func(o *options) { o.query["cluster_name"] = clusterName })
}

// WithCreated created获取 
func (obj *_V$obClusterMgr) WithCreated(created time.Time) Option {
	return optionFunc(func(o *options) { o.query["created"] = created })
}

// WithClusterRole cluster_role获取 
func (obj *_V$obClusterMgr) WithClusterRole(clusterRole string) Option {
	return optionFunc(func(o *options) { o.query["cluster_role"] = clusterRole })
}

// WithClusterStatus cluster_status获取 
func (obj *_V$obClusterMgr) WithClusterStatus(clusterStatus string) Option {
	return optionFunc(func(o *options) { o.query["cluster_status"] = clusterStatus })
}

// WithSwitchover# switchover#获取 
func (obj *_V$obClusterMgr) WithSwitchover#(switchover# int64) Option {
	return optionFunc(func(o *options) { o.query["switchover#"] = switchover# })
}

// WithSwitchoverStatus switchover_status获取 
func (obj *_V$obClusterMgr) WithSwitchoverStatus(switchoverStatus string) Option {
	return optionFunc(func(o *options) { o.query["switchover_status"] = switchoverStatus })
}

// WithSwitchoverInfo switchover_info获取 
func (obj *_V$obClusterMgr) WithSwitchoverInfo(switchoverInfo string) Option {
	return optionFunc(func(o *options) { o.query["switchover_info"] = switchoverInfo })
}

// WithCurrentScn current_scn获取 
func (obj *_V$obClusterMgr) WithCurrentScn(currentScn int64) Option {
	return optionFunc(func(o *options) { o.query["current_scn"] = currentScn })
}

// WithStandbyBecamePrimaryScn standby_became_primary_scn获取 
func (obj *_V$obClusterMgr) WithStandbyBecamePrimaryScn(standbyBecamePrimaryScn int64) Option {
	return optionFunc(func(o *options) { o.query["standby_became_primary_scn"] = standbyBecamePrimaryScn })
}

// WithPrimaryClusterID primary_cluster_id获取 
func (obj *_V$obClusterMgr) WithPrimaryClusterID(primaryClusterID int64) Option {
	return optionFunc(func(o *options) { o.query["primary_cluster_id"] = primaryClusterID })
}

// WithProtectionMode protection_mode获取 
func (obj *_V$obClusterMgr) WithProtectionMode(protectionMode string) Option {
	return optionFunc(func(o *options) { o.query["protection_mode"] = protectionMode })
}

// WithProtectionLevel protection_level获取 
func (obj *_V$obClusterMgr) WithProtectionLevel(protectionLevel string) Option {
	return optionFunc(func(o *options) { o.query["protection_level"] = protectionLevel })
}

// WithRedoTransportOptions redo_transport_options获取 
func (obj *_V$obClusterMgr) WithRedoTransportOptions(redoTransportOptions string) Option {
	return optionFunc(func(o *options) { o.query["redo_transport_options"] = redoTransportOptions })
}


// GetByOption 功能选项模式获取
func (obj *_V$obClusterMgr) GetByOption(opts ...Option) (result V$obCluster, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$obClusterMgr) GetByOptions(opts ...Option) (results []*V$obCluster, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$obClusterMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$obCluster,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where(options.query)
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


// GetFromClusterID 通过cluster_id获取内容  
func (obj *_V$obClusterMgr) GetFromClusterID(clusterID int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`cluster_id` = ?", clusterID).Find(&results).Error
	
	return
}

// GetBatchFromClusterID 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromClusterID(clusterIDs []int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`cluster_id` IN (?)", clusterIDs).Find(&results).Error
	
	return
}
 
// GetFromClusterName 通过cluster_name获取内容  
func (obj *_V$obClusterMgr) GetFromClusterName(clusterName string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`cluster_name` = ?", clusterName).Find(&results).Error
	
	return
}

// GetBatchFromClusterName 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromClusterName(clusterNames []string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`cluster_name` IN (?)", clusterNames).Find(&results).Error
	
	return
}
 
// GetFromCreated 通过created获取内容  
func (obj *_V$obClusterMgr) GetFromCreated(created time.Time) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`created` = ?", created).Find(&results).Error
	
	return
}

// GetBatchFromCreated 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromCreated(createds []time.Time) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`created` IN (?)", createds).Find(&results).Error
	
	return
}
 
// GetFromClusterRole 通过cluster_role获取内容  
func (obj *_V$obClusterMgr) GetFromClusterRole(clusterRole string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`cluster_role` = ?", clusterRole).Find(&results).Error
	
	return
}

// GetBatchFromClusterRole 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromClusterRole(clusterRoles []string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`cluster_role` IN (?)", clusterRoles).Find(&results).Error
	
	return
}
 
// GetFromClusterStatus 通过cluster_status获取内容  
func (obj *_V$obClusterMgr) GetFromClusterStatus(clusterStatus string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`cluster_status` = ?", clusterStatus).Find(&results).Error
	
	return
}

// GetBatchFromClusterStatus 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromClusterStatus(clusterStatuss []string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`cluster_status` IN (?)", clusterStatuss).Find(&results).Error
	
	return
}
 
// GetFromSwitchover# 通过switchover#获取内容  
func (obj *_V$obClusterMgr) GetFromSwitchover#(switchover# int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`switchover#` = ?", switchover#).Find(&results).Error
	
	return
}

// GetBatchFromSwitchover# 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromSwitchover#(switchover#s []int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`switchover#` IN (?)", switchover#s).Find(&results).Error
	
	return
}
 
// GetFromSwitchoverStatus 通过switchover_status获取内容  
func (obj *_V$obClusterMgr) GetFromSwitchoverStatus(switchoverStatus string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`switchover_status` = ?", switchoverStatus).Find(&results).Error
	
	return
}

// GetBatchFromSwitchoverStatus 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromSwitchoverStatus(switchoverStatuss []string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`switchover_status` IN (?)", switchoverStatuss).Find(&results).Error
	
	return
}
 
// GetFromSwitchoverInfo 通过switchover_info获取内容  
func (obj *_V$obClusterMgr) GetFromSwitchoverInfo(switchoverInfo string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`switchover_info` = ?", switchoverInfo).Find(&results).Error
	
	return
}

// GetBatchFromSwitchoverInfo 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromSwitchoverInfo(switchoverInfos []string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`switchover_info` IN (?)", switchoverInfos).Find(&results).Error
	
	return
}
 
// GetFromCurrentScn 通过current_scn获取内容  
func (obj *_V$obClusterMgr) GetFromCurrentScn(currentScn int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`current_scn` = ?", currentScn).Find(&results).Error
	
	return
}

// GetBatchFromCurrentScn 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromCurrentScn(currentScns []int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`current_scn` IN (?)", currentScns).Find(&results).Error
	
	return
}
 
// GetFromStandbyBecamePrimaryScn 通过standby_became_primary_scn获取内容  
func (obj *_V$obClusterMgr) GetFromStandbyBecamePrimaryScn(standbyBecamePrimaryScn int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`standby_became_primary_scn` = ?", standbyBecamePrimaryScn).Find(&results).Error
	
	return
}

// GetBatchFromStandbyBecamePrimaryScn 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromStandbyBecamePrimaryScn(standbyBecamePrimaryScns []int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`standby_became_primary_scn` IN (?)", standbyBecamePrimaryScns).Find(&results).Error
	
	return
}
 
// GetFromPrimaryClusterID 通过primary_cluster_id获取内容  
func (obj *_V$obClusterMgr) GetFromPrimaryClusterID(primaryClusterID int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`primary_cluster_id` = ?", primaryClusterID).Find(&results).Error
	
	return
}

// GetBatchFromPrimaryClusterID 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromPrimaryClusterID(primaryClusterIDs []int64) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`primary_cluster_id` IN (?)", primaryClusterIDs).Find(&results).Error
	
	return
}
 
// GetFromProtectionMode 通过protection_mode获取内容  
func (obj *_V$obClusterMgr) GetFromProtectionMode(protectionMode string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`protection_mode` = ?", protectionMode).Find(&results).Error
	
	return
}

// GetBatchFromProtectionMode 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromProtectionMode(protectionModes []string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`protection_mode` IN (?)", protectionModes).Find(&results).Error
	
	return
}
 
// GetFromProtectionLevel 通过protection_level获取内容  
func (obj *_V$obClusterMgr) GetFromProtectionLevel(protectionLevel string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`protection_level` = ?", protectionLevel).Find(&results).Error
	
	return
}

// GetBatchFromProtectionLevel 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromProtectionLevel(protectionLevels []string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`protection_level` IN (?)", protectionLevels).Find(&results).Error
	
	return
}
 
// GetFromRedoTransportOptions 通过redo_transport_options获取内容  
func (obj *_V$obClusterMgr) GetFromRedoTransportOptions(redoTransportOptions string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`redo_transport_options` = ?", redoTransportOptions).Find(&results).Error
	
	return
}

// GetBatchFromRedoTransportOptions 批量查找 
func (obj *_V$obClusterMgr) GetBatchFromRedoTransportOptions(redoTransportOptionss []string) (results []*V$obCluster, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obCluster{}).Where("`redo_transport_options` IN (?)", redoTransportOptionss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

