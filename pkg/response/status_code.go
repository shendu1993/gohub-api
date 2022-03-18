package response

//次状态码用于给前端使用，用于告知数据是否正常返回
//是否提示错误信息OR进行其他的下一步的操作
const (
	StatusError   = 0 //接口返回数据失败时，状态码 code =0
	StatusSuccess = 1 //接口返回数据正常时，状态码 code =1
)
