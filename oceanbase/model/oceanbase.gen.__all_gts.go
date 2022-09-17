package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllGtsMgr struct {
	*_BaseMgr
}

// AllGtsMgr open func
func AllGtsMgr(db *gorm.DB) *_AllGtsMgr {
	if db == nil {
		panic(fmt.Errorf("AllGtsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllGtsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_gts"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllGtsMgr) GetTableName() string {
	return "__all_gts"
}

// Reset 重置gorm会话
func (obj *_AllGtsMgr) Reset() *_AllGtsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllGtsMgr) Get() (result AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllGtsMgr) Gets() (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllGtsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllGts{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllGtsMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllGtsMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithGtsID gts_id获取
func (obj *_AllGtsMgr) WithGtsID(gtsID int64) Option {
	return optionFunc(func(o *options) { o.query["gts_id"] = gtsID })
}

// WithGtsName gts_name获取
func (obj *_AllGtsMgr) WithGtsName(gtsName string) Option {
	return optionFunc(func(o *options) { o.query["gts_name"] = gtsName })
}

// WithRegion region获取
func (obj *_AllGtsMgr) WithRegion(region string) Option {
	return optionFunc(func(o *options) { o.query["region"] = region })
}

// WithEpochID epoch_id获取
func (obj *_AllGtsMgr) WithEpochID(epochID int64) Option {
	return optionFunc(func(o *options) { o.query["epoch_id"] = epochID })
}

// WithMemberList member_list获取
func (obj *_AllGtsMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithStandbyIP standby_ip获取
func (obj *_AllGtsMgr) WithStandbyIP(standbyIP string) Option {
	return optionFunc(func(o *options) { o.query["standby_ip"] = standbyIP })
}

// WithStandbyPort standby_port获取
func (obj *_AllGtsMgr) WithStandbyPort(standbyPort int64) Option {
	return optionFunc(func(o *options) { o.query["standby_port"] = standbyPort })
}

// WithHeartbeatTs heartbeat_ts获取
func (obj *_AllGtsMgr) WithHeartbeatTs(heartbeatTs int64) Option {
	return optionFunc(func(o *options) { o.query["heartbeat_ts"] = heartbeatTs })
}

// GetByOption 功能选项模式获取
func (obj *_AllGtsMgr) GetByOption(opts ...Option) (result AllGts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllGtsMgr) GetByOptions(opts ...Option) (results []*AllGts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllGtsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllGts, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where(options.query)
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
func (obj *_AllGtsMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllGtsMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllGtsMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllGtsMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromGtsID 通过gts_id获取内容
func (obj *_AllGtsMgr) GetFromGtsID(gtsID int64) (result AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gts_id` = ?", gtsID).First(&result).Error

	return
}

// GetBatchFromGtsID 批量查找
func (obj *_AllGtsMgr) GetBatchFromGtsID(gtsIDs []int64) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gts_id` IN (?)", gtsIDs).Find(&results).Error

	return
}

// GetFromGtsName 通过gts_name获取内容
func (obj *_AllGtsMgr) GetFromGtsName(gtsName string) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gts_name` = ?", gtsName).Find(&results).Error

	return
}

// GetBatchFromGtsName 批量查找
func (obj *_AllGtsMgr) GetBatchFromGtsName(gtsNames []string) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gts_name` IN (?)", gtsNames).Find(&results).Error

	return
}

// GetFromRegion 通过region获取内容
func (obj *_AllGtsMgr) GetFromRegion(region string) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`region` = ?", region).Find(&results).Error

	return
}

// GetBatchFromRegion 批量查找
func (obj *_AllGtsMgr) GetBatchFromRegion(regions []string) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`region` IN (?)", regions).Find(&results).Error

	return
}

// GetFromEpochID 通过epoch_id获取内容
func (obj *_AllGtsMgr) GetFromEpochID(epochID int64) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`epoch_id` = ?", epochID).Find(&results).Error

	return
}

// GetBatchFromEpochID 批量查找
func (obj *_AllGtsMgr) GetBatchFromEpochID(epochIDs []int64) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`epoch_id` IN (?)", epochIDs).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllGtsMgr) GetFromMemberList(memberList string) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllGtsMgr) GetBatchFromMemberList(memberLists []string) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromStandbyIP 通过standby_ip获取内容
func (obj *_AllGtsMgr) GetFromStandbyIP(standbyIP string) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`standby_ip` = ?", standbyIP).Find(&results).Error

	return
}

// GetBatchFromStandbyIP 批量查找
func (obj *_AllGtsMgr) GetBatchFromStandbyIP(standbyIPs []string) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`standby_ip` IN (?)", standbyIPs).Find(&results).Error

	return
}

// GetFromStandbyPort 通过standby_port获取内容
func (obj *_AllGtsMgr) GetFromStandbyPort(standbyPort int64) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`standby_port` = ?", standbyPort).Find(&results).Error

	return
}

// GetBatchFromStandbyPort 批量查找
func (obj *_AllGtsMgr) GetBatchFromStandbyPort(standbyPorts []int64) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`standby_port` IN (?)", standbyPorts).Find(&results).Error

	return
}

// GetFromHeartbeatTs 通过heartbeat_ts获取内容
func (obj *_AllGtsMgr) GetFromHeartbeatTs(heartbeatTs int64) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`heartbeat_ts` = ?", heartbeatTs).Find(&results).Error

	return
}

// GetBatchFromHeartbeatTs 批量查找
func (obj *_AllGtsMgr) GetBatchFromHeartbeatTs(heartbeatTss []int64) (results []*AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`heartbeat_ts` IN (?)", heartbeatTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllGtsMgr) FetchByPrimaryKey(gtsID int64) (result AllGts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGts{}).Where("`gts_id` = ?", gtsID).First(&result).Error

	return
}
