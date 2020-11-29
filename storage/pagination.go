package storage

type PageQuery struct {
	Page  	int64 `form:"page"`
	Length 	int64 `form:"length"`
}

var pageMin = int64(1)
var LengthDefault = int64(10)
var LengthMax = int64(500)

func (page *PageQuery)Offset()int64{
	if page.Page <= 0 {
		page.Page = pageMin
	}
	return (page.Page-1)*page.Limit()
}

func (page *PageQuery)Limit()int64{
	if page.Length <= 0 {
		return LengthDefault
	}
	if page.Length > LengthMax {

		return LengthMax
	}
	return page.Length
}

func (page *PageQuery)Pagination()int64{
	if page.Page <= 0 {
		page.Page = pageMin
	}
	return page.Page
}

type PageResult map[string]interface{}

/**
	分页接口
 */
type Pagination interface {
	Pagination(PageQuery)
}


///*
//	根据分页接口的实现对象，动态创建DB分页模型
//	@param pagination 接收实现了分页接口的模型（对象）作为参数
//*/
//func PageDB(pagination Pagination, page PageQuery) (*gorm.DB){
//
//	if page.Page <= 0 {
//		page.Page = 1
//	}
//	if page.Length <= 0 {
//		page.Length = 10
//	}
//	if page.Length > 200 {
//		page.Length = 200
//	}
//
//	return DB.Model(pagination).Offset(int((page.Page-1)*page.Length)).Limit(page.Length)
//}
