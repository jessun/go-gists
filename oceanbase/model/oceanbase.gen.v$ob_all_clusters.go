package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$obAllClustersMgr struct {
	*_BaseMgr
}

// V$obAllClustersMgr open func
func V$obAllClustersMgr(db *gorm.DB) *_V$obAllClustersMgr {
	if db == nil {
		panic(fmt.Errorf("V$obAllClustersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$obAllClustersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$ob_all_clusters"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$obAllClustersMgr) GetTableName() string {
	return "v$ob_all_clusters"
}

// Reset 重置gorm会话
func (obj *_V$obAllClustersMgr) Reset() *_V$obAllClustersMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$obAllClustersMgr) Get() (result V$obAllClusters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obAllClusters{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$obAllClustersMgr) Gets() (results []*V$obAllClusters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obAllClusters{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$obAllClustersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$obAllClusters{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////


// GetByOption 功能选项模式获取
func (obj *_V$obAllClustersMgr) GetByOption(opts ...Option) (result V$obAllClusters, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obAllClusters{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$obAllClustersMgr) GetByOptions(opts ...Option) (results []*V$obAllClusters, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obAllClusters{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$obAllClustersMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$obAllClusters,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$obAllClusters{}).Where(options.query)
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


 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

