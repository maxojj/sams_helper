package conf

import "errors"

var AuthTokenErr = errors.New("auth-token 可能未设置，请检查")
var ProxyErr = errors.New("网络错误，请检查是否设置错误代理")
var NoValidAddressErr = errors.New("没有有效的收货地址，请前往 APP 添加或者检查 Auth-Token 是否正确")
var NoGoodsErr = errors.New("当前购物车中无有效商品")
var FixCartErr = errors.New("修正购物车列表限购商品数量失败")
var RunModeErr = errors.New("运行模式错误，请检查配置")

var AuthFailErr = errors.New("鉴权失败 auth-token 过期")
var CartGoodChangeErr = errors.New("购物车商品发生变化，请返回购物车页面重新结算")
var LimitedErr = errors.New("服务器正忙,请稍后再试")
var LimitedErr1 = errors.New("当前购物火爆，请稍后再试")
var OutOfSellErr = errors.New("部分商品已缺货")
var CapacityFullErr = errors.New("当前无剩余运力，重新检测是否释放")
var FAILErr = errors.New("未知失败")
var NoMatchDeliverMode = errors.New("当前区域不支持配送，请重新选择地址")
var CloseOrderTimeExceptionErr = errors.New("尊敬的会员，您选择的配送时间已失效，请重新选择")
var NotDeliverCapCityErr = errors.New("当前配送时间段已约满，请重新选择配送时段")
var DecreaseCapacityCountError = errors.New("扣减运力失败")
var StoreHasClosedError = errors.New("门店已打烊")
var DeliveryTypeErr = errors.New("未知设备类型")
var NotCheckShopPendingErr = errors.New("请阅读并勾选《购物须知》")

// 429 429 {"message":"Requests rate limited. stage:service"}
