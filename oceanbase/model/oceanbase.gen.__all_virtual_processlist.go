package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualProcesslistMgr struct {
	*_BaseMgr
}

// AllVirtualProcesslistMgr open func
func AllVirtualProcesslistMgr(db *gorm.DB) *_AllVirtualProcesslistMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualProcesslistMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualProcesslistMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_processlist"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualProcesslistMgr) GetTableName() string {
	return "__all_virtual_processlist"
}

// Reset 重置gorm会话
func (obj *_AllVirtualProcesslistMgr) Reset() *_AllVirtualProcesslistMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualProcesslistMgr) Get() (result AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualProcesslistMgr) Gets() (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualProcesslistMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AllVirtualProcesslistMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUser user获取
func (obj *_AllVirtualProcesslistMgr) WithUser(user string) Option {
	return optionFunc(func(o *options) { o.query["user"] = user })
}

// WithTenant tenant获取
func (obj *_AllVirtualProcesslistMgr) WithTenant(tenant string) Option {
	return optionFunc(func(o *options) { o.query["tenant"] = tenant })
}

// WithHost host获取
func (obj *_AllVirtualProcesslistMgr) WithHost(host string) Option {
	return optionFunc(func(o *options) { o.query["host"] = host })
}

// WithDb db获取
func (obj *_AllVirtualProcesslistMgr) WithDb(db string) Option {
	return optionFunc(func(o *options) { o.query["db"] = db })
}

// WithCommand command获取
func (obj *_AllVirtualProcesslistMgr) WithCommand(command string) Option {
	return optionFunc(func(o *options) { o.query["command"] = command })
}

// WithSQLID sql_id获取
func (obj *_AllVirtualProcesslistMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithTime time获取
func (obj *_AllVirtualProcesslistMgr) WithTime(time int64) Option {
	return optionFunc(func(o *options) { o.query["time"] = time })
}

// WithState state获取
func (obj *_AllVirtualProcesslistMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithInfo info获取
func (obj *_AllVirtualProcesslistMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualProcesslistMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualProcesslistMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSQLPort sql_port获取
func (obj *_AllVirtualProcesslistMgr) WithSQLPort(sqlPort int64) Option {
	return optionFunc(func(o *options) { o.query["sql_port"] = sqlPort })
}

// WithProxySessid proxy_sessid获取
func (obj *_AllVirtualProcesslistMgr) WithProxySessid(proxySessid uint64) Option {
	return optionFunc(func(o *options) { o.query["proxy_sessid"] = proxySessid })
}

// WithMasterSessid master_sessid获取
func (obj *_AllVirtualProcesslistMgr) WithMasterSessid(masterSessid uint64) Option {
	return optionFunc(func(o *options) { o.query["master_sessid"] = masterSessid })
}

// WithUserClientIP user_client_ip获取
func (obj *_AllVirtualProcesslistMgr) WithUserClientIP(userClientIP string) Option {
	return optionFunc(func(o *options) { o.query["user_client_ip"] = userClientIP })
}

// WithUserHost user_host获取
func (obj *_AllVirtualProcesslistMgr) WithUserHost(userHost string) Option {
	return optionFunc(func(o *options) { o.query["user_host"] = userHost })
}

// WithTransID trans_id获取
func (obj *_AllVirtualProcesslistMgr) WithTransID(transID uint64) Option {
	return optionFunc(func(o *options) { o.query["trans_id"] = transID })
}

// WithThreadID thread_id获取
func (obj *_AllVirtualProcesslistMgr) WithThreadID(threadID uint64) Option {
	return optionFunc(func(o *options) { o.query["thread_id"] = threadID })
}

// WithSslCipher ssl_cipher获取
func (obj *_AllVirtualProcesslistMgr) WithSslCipher(sslCipher string) Option {
	return optionFunc(func(o *options) { o.query["ssl_cipher"] = sslCipher })
}

// WithTraceID trace_id获取
func (obj *_AllVirtualProcesslistMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualProcesslistMgr) GetByOption(opts ...Option) (result AllVirtualProcesslist, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualProcesslistMgr) GetByOptions(opts ...Option) (results []*AllVirtualProcesslist, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualProcesslistMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualProcesslist, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromID(id uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromID(ids []uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUser 通过user获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromUser(user string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`user` = ?", user).Find(&results).Error

	return
}

// GetBatchFromUser 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromUser(users []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`user` IN (?)", users).Find(&results).Error

	return
}

// GetFromTenant 通过tenant获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromTenant(tenant string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`tenant` = ?", tenant).Find(&results).Error

	return
}

// GetBatchFromTenant 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromTenant(tenants []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`tenant` IN (?)", tenants).Find(&results).Error

	return
}

// GetFromHost 通过host获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromHost(host string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`host` = ?", host).Find(&results).Error

	return
}

// GetBatchFromHost 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromHost(hosts []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`host` IN (?)", hosts).Find(&results).Error

	return
}

// GetFromDb 通过db获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromDb(db string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`db` = ?", db).Find(&results).Error

	return
}

// GetBatchFromDb 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromDb(dbs []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`db` IN (?)", dbs).Find(&results).Error

	return
}

// GetFromCommand 通过command获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromCommand(command string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`command` = ?", command).Find(&results).Error

	return
}

// GetBatchFromCommand 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromCommand(commands []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`command` IN (?)", commands).Find(&results).Error

	return
}

// GetFromSQLID 通过sql_id获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromSQLID(sqlID string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`sql_id` = ?", sqlID).Find(&results).Error

	return
}

// GetBatchFromSQLID 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromSQLID(sqlIDs []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error

	return
}

// GetFromTime 通过time获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromTime(time int64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`time` = ?", time).Find(&results).Error

	return
}

// GetBatchFromTime 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromTime(times []int64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`time` IN (?)", times).Find(&results).Error

	return
}

// GetFromState 通过state获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromState(state string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromState(states []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromInfo(info string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSQLPort 通过sql_port获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromSQLPort(sqlPort int64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`sql_port` = ?", sqlPort).Find(&results).Error

	return
}

// GetBatchFromSQLPort 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromSQLPort(sqlPorts []int64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`sql_port` IN (?)", sqlPorts).Find(&results).Error

	return
}

// GetFromProxySessid 通过proxy_sessid获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromProxySessid(proxySessid uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`proxy_sessid` = ?", proxySessid).Find(&results).Error

	return
}

// GetBatchFromProxySessid 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromProxySessid(proxySessids []uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`proxy_sessid` IN (?)", proxySessids).Find(&results).Error

	return
}

// GetFromMasterSessid 通过master_sessid获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromMasterSessid(masterSessid uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`master_sessid` = ?", masterSessid).Find(&results).Error

	return
}

// GetBatchFromMasterSessid 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromMasterSessid(masterSessids []uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`master_sessid` IN (?)", masterSessids).Find(&results).Error

	return
}

// GetFromUserClientIP 通过user_client_ip获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromUserClientIP(userClientIP string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`user_client_ip` = ?", userClientIP).Find(&results).Error

	return
}

// GetBatchFromUserClientIP 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromUserClientIP(userClientIPs []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`user_client_ip` IN (?)", userClientIPs).Find(&results).Error

	return
}

// GetFromUserHost 通过user_host获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromUserHost(userHost string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`user_host` = ?", userHost).Find(&results).Error

	return
}

// GetBatchFromUserHost 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromUserHost(userHosts []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`user_host` IN (?)", userHosts).Find(&results).Error

	return
}

// GetFromTransID 通过trans_id获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromTransID(transID uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`trans_id` = ?", transID).Find(&results).Error

	return
}

// GetBatchFromTransID 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromTransID(transIDs []uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`trans_id` IN (?)", transIDs).Find(&results).Error

	return
}

// GetFromThreadID 通过thread_id获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromThreadID(threadID uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`thread_id` = ?", threadID).Find(&results).Error

	return
}

// GetBatchFromThreadID 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromThreadID(threadIDs []uint64) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`thread_id` IN (?)", threadIDs).Find(&results).Error

	return
}

// GetFromSslCipher 通过ssl_cipher获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromSslCipher(sslCipher string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`ssl_cipher` = ?", sslCipher).Find(&results).Error

	return
}

// GetBatchFromSslCipher 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromSslCipher(sslCiphers []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`ssl_cipher` IN (?)", sslCiphers).Find(&results).Error

	return
}

// GetFromTraceID 通过trace_id获取内容
func (obj *_AllVirtualProcesslistMgr) GetFromTraceID(traceID string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`trace_id` = ?", traceID).Find(&results).Error

	return
}

// GetBatchFromTraceID 批量查找
func (obj *_AllVirtualProcesslistMgr) GetBatchFromTraceID(traceIDs []string) (results []*AllVirtualProcesslist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProcesslist{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
