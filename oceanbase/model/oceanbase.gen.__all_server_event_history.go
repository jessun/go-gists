package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllServerEventHistoryMgr struct {
	*_BaseMgr
}

// AllServerEventHistoryMgr open func
func AllServerEventHistoryMgr(db *gorm.DB) *_AllServerEventHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllServerEventHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllServerEventHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_server_event_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllServerEventHistoryMgr) GetTableName() string {
	return "__all_server_event_history"
}

// Reset 重置gorm会话
func (obj *_AllServerEventHistoryMgr) Reset() *_AllServerEventHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllServerEventHistoryMgr) Get() (result AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllServerEventHistoryMgr) Gets() (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllServerEventHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllServerEventHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithSvrIP svr_ip获取
func (obj *_AllServerEventHistoryMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllServerEventHistoryMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithModule module获取
func (obj *_AllServerEventHistoryMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithEvent event获取
func (obj *_AllServerEventHistoryMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithName1 name1获取
func (obj *_AllServerEventHistoryMgr) WithName1(name1 string) Option {
	return optionFunc(func(o *options) { o.query["name1"] = name1 })
}

// WithValue1 value1获取
func (obj *_AllServerEventHistoryMgr) WithValue1(value1 string) Option {
	return optionFunc(func(o *options) { o.query["value1"] = value1 })
}

// WithName2 name2获取
func (obj *_AllServerEventHistoryMgr) WithName2(name2 string) Option {
	return optionFunc(func(o *options) { o.query["name2"] = name2 })
}

// WithValue2 value2获取
func (obj *_AllServerEventHistoryMgr) WithValue2(value2 string) Option {
	return optionFunc(func(o *options) { o.query["value2"] = value2 })
}

// WithName3 name3获取
func (obj *_AllServerEventHistoryMgr) WithName3(name3 string) Option {
	return optionFunc(func(o *options) { o.query["name3"] = name3 })
}

// WithValue3 value3获取
func (obj *_AllServerEventHistoryMgr) WithValue3(value3 string) Option {
	return optionFunc(func(o *options) { o.query["value3"] = value3 })
}

// WithName4 name4获取
func (obj *_AllServerEventHistoryMgr) WithName4(name4 string) Option {
	return optionFunc(func(o *options) { o.query["name4"] = name4 })
}

// WithValue4 value4获取
func (obj *_AllServerEventHistoryMgr) WithValue4(value4 string) Option {
	return optionFunc(func(o *options) { o.query["value4"] = value4 })
}

// WithName5 name5获取
func (obj *_AllServerEventHistoryMgr) WithName5(name5 string) Option {
	return optionFunc(func(o *options) { o.query["name5"] = name5 })
}

// WithValue5 value5获取
func (obj *_AllServerEventHistoryMgr) WithValue5(value5 string) Option {
	return optionFunc(func(o *options) { o.query["value5"] = value5 })
}

// WithName6 name6获取
func (obj *_AllServerEventHistoryMgr) WithName6(name6 string) Option {
	return optionFunc(func(o *options) { o.query["name6"] = name6 })
}

// WithValue6 value6获取
func (obj *_AllServerEventHistoryMgr) WithValue6(value6 string) Option {
	return optionFunc(func(o *options) { o.query["value6"] = value6 })
}

// WithExtraInfo extra_info获取
func (obj *_AllServerEventHistoryMgr) WithExtraInfo(extraInfo string) Option {
	return optionFunc(func(o *options) { o.query["extra_info"] = extraInfo })
}

// GetByOption 功能选项模式获取
func (obj *_AllServerEventHistoryMgr) GetByOption(opts ...Option) (result AllServerEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllServerEventHistoryMgr) GetByOptions(opts ...Option) (results []*AllServerEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllServerEventHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllServerEventHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where(options.query)
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
func (obj *_AllServerEventHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllServerEventHistoryMgr) GetFromSvrIP(svrIP string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllServerEventHistoryMgr) GetFromSvrPort(svrPort int64) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容
func (obj *_AllServerEventHistoryMgr) GetFromModule(module string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`module` = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromModule(modules []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`module` IN (?)", modules).Find(&results).Error

	return
}

// GetFromEvent 通过event获取内容
func (obj *_AllServerEventHistoryMgr) GetFromEvent(event string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`event` = ?", event).Find(&results).Error

	return
}

// GetBatchFromEvent 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromEvent(events []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`event` IN (?)", events).Find(&results).Error

	return
}

// GetFromName1 通过name1获取内容
func (obj *_AllServerEventHistoryMgr) GetFromName1(name1 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name1` = ?", name1).Find(&results).Error

	return
}

// GetBatchFromName1 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromName1(name1s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name1` IN (?)", name1s).Find(&results).Error

	return
}

// GetFromValue1 通过value1获取内容
func (obj *_AllServerEventHistoryMgr) GetFromValue1(value1 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value1` = ?", value1).Find(&results).Error

	return
}

// GetBatchFromValue1 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromValue1(value1s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value1` IN (?)", value1s).Find(&results).Error

	return
}

// GetFromName2 通过name2获取内容
func (obj *_AllServerEventHistoryMgr) GetFromName2(name2 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name2` = ?", name2).Find(&results).Error

	return
}

// GetBatchFromName2 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromName2(name2s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name2` IN (?)", name2s).Find(&results).Error

	return
}

// GetFromValue2 通过value2获取内容
func (obj *_AllServerEventHistoryMgr) GetFromValue2(value2 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value2` = ?", value2).Find(&results).Error

	return
}

// GetBatchFromValue2 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromValue2(value2s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value2` IN (?)", value2s).Find(&results).Error

	return
}

// GetFromName3 通过name3获取内容
func (obj *_AllServerEventHistoryMgr) GetFromName3(name3 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name3` = ?", name3).Find(&results).Error

	return
}

// GetBatchFromName3 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromName3(name3s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name3` IN (?)", name3s).Find(&results).Error

	return
}

// GetFromValue3 通过value3获取内容
func (obj *_AllServerEventHistoryMgr) GetFromValue3(value3 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value3` = ?", value3).Find(&results).Error

	return
}

// GetBatchFromValue3 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromValue3(value3s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value3` IN (?)", value3s).Find(&results).Error

	return
}

// GetFromName4 通过name4获取内容
func (obj *_AllServerEventHistoryMgr) GetFromName4(name4 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name4` = ?", name4).Find(&results).Error

	return
}

// GetBatchFromName4 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromName4(name4s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name4` IN (?)", name4s).Find(&results).Error

	return
}

// GetFromValue4 通过value4获取内容
func (obj *_AllServerEventHistoryMgr) GetFromValue4(value4 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value4` = ?", value4).Find(&results).Error

	return
}

// GetBatchFromValue4 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromValue4(value4s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value4` IN (?)", value4s).Find(&results).Error

	return
}

// GetFromName5 通过name5获取内容
func (obj *_AllServerEventHistoryMgr) GetFromName5(name5 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name5` = ?", name5).Find(&results).Error

	return
}

// GetBatchFromName5 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromName5(name5s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name5` IN (?)", name5s).Find(&results).Error

	return
}

// GetFromValue5 通过value5获取内容
func (obj *_AllServerEventHistoryMgr) GetFromValue5(value5 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value5` = ?", value5).Find(&results).Error

	return
}

// GetBatchFromValue5 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromValue5(value5s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value5` IN (?)", value5s).Find(&results).Error

	return
}

// GetFromName6 通过name6获取内容
func (obj *_AllServerEventHistoryMgr) GetFromName6(name6 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name6` = ?", name6).Find(&results).Error

	return
}

// GetBatchFromName6 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromName6(name6s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`name6` IN (?)", name6s).Find(&results).Error

	return
}

// GetFromValue6 通过value6获取内容
func (obj *_AllServerEventHistoryMgr) GetFromValue6(value6 string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value6` = ?", value6).Find(&results).Error

	return
}

// GetBatchFromValue6 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromValue6(value6s []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`value6` IN (?)", value6s).Find(&results).Error

	return
}

// GetFromExtraInfo 通过extra_info获取内容
func (obj *_AllServerEventHistoryMgr) GetFromExtraInfo(extraInfo string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`extra_info` = ?", extraInfo).Find(&results).Error

	return
}

// GetBatchFromExtraInfo 批量查找
func (obj *_AllServerEventHistoryMgr) GetBatchFromExtraInfo(extraInfos []string) (results []*AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`extra_info` IN (?)", extraInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllServerEventHistoryMgr) FetchByPrimaryKey(gmtCreate time.Time, svrIP string, svrPort int64) (result AllServerEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerEventHistory{}).Where("`gmt_create` = ? AND `svr_ip` = ? AND `svr_port` = ?", gmtCreate, svrIP, svrPort).First(&result).Error

	return
}
