package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllRootserviceEventHistoryMgr struct {
	*_BaseMgr
}

// AllRootserviceEventHistoryMgr open func
func AllRootserviceEventHistoryMgr(db *gorm.DB) *_AllRootserviceEventHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllRootserviceEventHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllRootserviceEventHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_rootservice_event_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllRootserviceEventHistoryMgr) GetTableName() string {
	return "__all_rootservice_event_history"
}

// Reset 重置gorm会话
func (obj *_AllRootserviceEventHistoryMgr) Reset() *_AllRootserviceEventHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllRootserviceEventHistoryMgr) Get() (result AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllRootserviceEventHistoryMgr) Gets() (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllRootserviceEventHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllRootserviceEventHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithModule module获取
func (obj *_AllRootserviceEventHistoryMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithEvent event获取
func (obj *_AllRootserviceEventHistoryMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithName1 name1获取
func (obj *_AllRootserviceEventHistoryMgr) WithName1(name1 string) Option {
	return optionFunc(func(o *options) { o.query["name1"] = name1 })
}

// WithValue1 value1获取
func (obj *_AllRootserviceEventHistoryMgr) WithValue1(value1 string) Option {
	return optionFunc(func(o *options) { o.query["value1"] = value1 })
}

// WithName2 name2获取
func (obj *_AllRootserviceEventHistoryMgr) WithName2(name2 string) Option {
	return optionFunc(func(o *options) { o.query["name2"] = name2 })
}

// WithValue2 value2获取
func (obj *_AllRootserviceEventHistoryMgr) WithValue2(value2 string) Option {
	return optionFunc(func(o *options) { o.query["value2"] = value2 })
}

// WithName3 name3获取
func (obj *_AllRootserviceEventHistoryMgr) WithName3(name3 string) Option {
	return optionFunc(func(o *options) { o.query["name3"] = name3 })
}

// WithValue3 value3获取
func (obj *_AllRootserviceEventHistoryMgr) WithValue3(value3 string) Option {
	return optionFunc(func(o *options) { o.query["value3"] = value3 })
}

// WithName4 name4获取
func (obj *_AllRootserviceEventHistoryMgr) WithName4(name4 string) Option {
	return optionFunc(func(o *options) { o.query["name4"] = name4 })
}

// WithValue4 value4获取
func (obj *_AllRootserviceEventHistoryMgr) WithValue4(value4 string) Option {
	return optionFunc(func(o *options) { o.query["value4"] = value4 })
}

// WithName5 name5获取
func (obj *_AllRootserviceEventHistoryMgr) WithName5(name5 string) Option {
	return optionFunc(func(o *options) { o.query["name5"] = name5 })
}

// WithValue5 value5获取
func (obj *_AllRootserviceEventHistoryMgr) WithValue5(value5 string) Option {
	return optionFunc(func(o *options) { o.query["value5"] = value5 })
}

// WithName6 name6获取
func (obj *_AllRootserviceEventHistoryMgr) WithName6(name6 string) Option {
	return optionFunc(func(o *options) { o.query["name6"] = name6 })
}

// WithValue6 value6获取
func (obj *_AllRootserviceEventHistoryMgr) WithValue6(value6 string) Option {
	return optionFunc(func(o *options) { o.query["value6"] = value6 })
}

// WithExtraInfo extra_info获取
func (obj *_AllRootserviceEventHistoryMgr) WithExtraInfo(extraInfo string) Option {
	return optionFunc(func(o *options) { o.query["extra_info"] = extraInfo })
}

// WithRsSvrIP rs_svr_ip获取
func (obj *_AllRootserviceEventHistoryMgr) WithRsSvrIP(rsSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["rs_svr_ip"] = rsSvrIP })
}

// WithRsSvrPort rs_svr_port获取
func (obj *_AllRootserviceEventHistoryMgr) WithRsSvrPort(rsSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["rs_svr_port"] = rsSvrPort })
}

// GetByOption 功能选项模式获取
func (obj *_AllRootserviceEventHistoryMgr) GetByOption(opts ...Option) (result AllRootserviceEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllRootserviceEventHistoryMgr) GetByOptions(opts ...Option) (results []*AllRootserviceEventHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllRootserviceEventHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllRootserviceEventHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where(options.query)
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
func (obj *_AllRootserviceEventHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (result AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`gmt_create` = ?", gmtCreate).First(&result).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromModule(module string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`module` = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromModule(modules []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`module` IN (?)", modules).Find(&results).Error

	return
}

// GetFromEvent 通过event获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromEvent(event string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`event` = ?", event).Find(&results).Error

	return
}

// GetBatchFromEvent 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromEvent(events []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`event` IN (?)", events).Find(&results).Error

	return
}

// GetFromName1 通过name1获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromName1(name1 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name1` = ?", name1).Find(&results).Error

	return
}

// GetBatchFromName1 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromName1(name1s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name1` IN (?)", name1s).Find(&results).Error

	return
}

// GetFromValue1 通过value1获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromValue1(value1 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value1` = ?", value1).Find(&results).Error

	return
}

// GetBatchFromValue1 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromValue1(value1s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value1` IN (?)", value1s).Find(&results).Error

	return
}

// GetFromName2 通过name2获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromName2(name2 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name2` = ?", name2).Find(&results).Error

	return
}

// GetBatchFromName2 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromName2(name2s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name2` IN (?)", name2s).Find(&results).Error

	return
}

// GetFromValue2 通过value2获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromValue2(value2 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value2` = ?", value2).Find(&results).Error

	return
}

// GetBatchFromValue2 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromValue2(value2s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value2` IN (?)", value2s).Find(&results).Error

	return
}

// GetFromName3 通过name3获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromName3(name3 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name3` = ?", name3).Find(&results).Error

	return
}

// GetBatchFromName3 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromName3(name3s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name3` IN (?)", name3s).Find(&results).Error

	return
}

// GetFromValue3 通过value3获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromValue3(value3 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value3` = ?", value3).Find(&results).Error

	return
}

// GetBatchFromValue3 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromValue3(value3s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value3` IN (?)", value3s).Find(&results).Error

	return
}

// GetFromName4 通过name4获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromName4(name4 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name4` = ?", name4).Find(&results).Error

	return
}

// GetBatchFromName4 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromName4(name4s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name4` IN (?)", name4s).Find(&results).Error

	return
}

// GetFromValue4 通过value4获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromValue4(value4 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value4` = ?", value4).Find(&results).Error

	return
}

// GetBatchFromValue4 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromValue4(value4s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value4` IN (?)", value4s).Find(&results).Error

	return
}

// GetFromName5 通过name5获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromName5(name5 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name5` = ?", name5).Find(&results).Error

	return
}

// GetBatchFromName5 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromName5(name5s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name5` IN (?)", name5s).Find(&results).Error

	return
}

// GetFromValue5 通过value5获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromValue5(value5 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value5` = ?", value5).Find(&results).Error

	return
}

// GetBatchFromValue5 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromValue5(value5s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value5` IN (?)", value5s).Find(&results).Error

	return
}

// GetFromName6 通过name6获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromName6(name6 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name6` = ?", name6).Find(&results).Error

	return
}

// GetBatchFromName6 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromName6(name6s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`name6` IN (?)", name6s).Find(&results).Error

	return
}

// GetFromValue6 通过value6获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromValue6(value6 string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value6` = ?", value6).Find(&results).Error

	return
}

// GetBatchFromValue6 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromValue6(value6s []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`value6` IN (?)", value6s).Find(&results).Error

	return
}

// GetFromExtraInfo 通过extra_info获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromExtraInfo(extraInfo string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`extra_info` = ?", extraInfo).Find(&results).Error

	return
}

// GetBatchFromExtraInfo 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromExtraInfo(extraInfos []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`extra_info` IN (?)", extraInfos).Find(&results).Error

	return
}

// GetFromRsSvrIP 通过rs_svr_ip获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromRsSvrIP(rsSvrIP string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`rs_svr_ip` = ?", rsSvrIP).Find(&results).Error

	return
}

// GetBatchFromRsSvrIP 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromRsSvrIP(rsSvrIPs []string) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`rs_svr_ip` IN (?)", rsSvrIPs).Find(&results).Error

	return
}

// GetFromRsSvrPort 通过rs_svr_port获取内容
func (obj *_AllRootserviceEventHistoryMgr) GetFromRsSvrPort(rsSvrPort int64) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`rs_svr_port` = ?", rsSvrPort).Find(&results).Error

	return
}

// GetBatchFromRsSvrPort 批量查找
func (obj *_AllRootserviceEventHistoryMgr) GetBatchFromRsSvrPort(rsSvrPorts []int64) (results []*AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`rs_svr_port` IN (?)", rsSvrPorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllRootserviceEventHistoryMgr) FetchByPrimaryKey(gmtCreate time.Time) (result AllRootserviceEventHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllRootserviceEventHistory{}).Where("`gmt_create` = ?", gmtCreate).First(&result).Error

	return
}
