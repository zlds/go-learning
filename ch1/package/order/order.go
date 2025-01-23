package order

func GetOrder(id int) string {
	return "我是测试用户的订单呀...."
}

// init函数在包被导入时执行，且只执行一次
func init() {
	println("我是order init方法，我在包被导入时执行，且只执行一次")
}
