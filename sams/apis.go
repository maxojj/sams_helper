package sams

// 地址

var AddressListAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/sams-user/receiver_address/address_list"
var SetAddressAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/trade/cart/saveDeliveryAddress"

type SetAddressParam struct {
	AddressId string `json:"addressId"`
	Uid       string `json:"uid"`
	AppId     string `json:"appId"`
	SaasId    string `json:"saasId"`
}

// 商店

var StoreListAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/merchant/storeApi/getRecommendStoreListByLocation"

type StoreListParam struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

// 购物车

var CartAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/trade/cart/getUserCart"

type CartParam struct {
	StoreList         []Store `json:"storeList"`
	AddressId         string  `json:"addressId"`
	Uid               string  `json:"uid"`
	DeliveryType      string  `json:"deliveryType"`
	HomePagelongitude string  `json:"homePagelongitude"`
	HomePagelatitude  string  `json:"homePagelatitude"`
}

var ModifyCartGoodsInfoAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/trade/cart/modifyCartGoodsInfo"

type ModifyCartGoodsInfoParam struct {
	CartGoodsInfo Goods  `json:"cartGoodsInfo"`
	Uid           string `json:"uid"`
}

var AddCartGoodsInfoAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/trade/cart/addCartGoodsInfo"

type AddCartGoodsInfoParam struct {
	CartGoodsInfoList []AddCartGoods `json:"cartGoodsInfoList"`
	Uid               string         `json:"uid"`
}

// 商品

var GoodsInfoAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/trade/settlement/checkGoodsInfo"

type GoodsInfoParam struct {
	FloorId   int64   `json:"floorId"`
	StoreId   string  `json:"storeId"`
	GoodsList []Goods `json:"goodsList"`
}

var SettleInfoAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/trade/settlement/getSettleInfo"

type SettleParam struct {
	Uid              string         `json:"uid"`
	AddressId        string         `json:"addressId"`
	DeliveryInfoVO   DeliveryInfoVO `json:"deliveryInfoVO"`
	CartDeliveryType int64          `json:"cartDeliveryType"`
	StoreInfo        StoreInfo      `json:"storeInfo"`
	CouponList       []string       `json:"couponList"`
	IsSelfPickup     int64          `json:"isSelfPickup"`
	FloorId          int64          `json:"floorId"`
	GoodsList        []Goods        `json:"goodsList"`
}

// 运力

var CapacityDataAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/delivery/portal/getCapacityData"

type CapacityDataParam struct {
	PerDateList             []string `json:"perDateList"`
	StoreDeliveryTemplateId string   `json:"storeDeliveryTemplateId"`
}

type IOSCapacityDataParam struct {
	CapacityDataParam
}

type MiniProgramCapacityDataParam struct {
	CapacityDataParam
	Uid    string `json:"uid"`
	AppId  string `json:"appId"`
	SassId string `json:"sassId"`
}

// 支付

var CommitPayAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/trade/settlement/commitPay"

type CommitPayParam struct {
	GoodsList        []Goods               `json:"goodsList"`
	InvoiceInfo      map[int64]interface{} `json:"invoiceInfo"`
	CartDeliveryType int64                 `json:"cartDeliveryType"`
	FloorId          int64                 `json:"floorId"`

	PurchaserName      string             `json:"purchaserName"`
	SettleDeliveryInfo SettleDeliveryInfo `json:"settleDeliveryInfo"`
	PayType            int64              `json:"payType"`
	Currency           string             `json:"currency"`
	Channel            string             `json:"channel"`
	ShortageId         int64              `json:"shortageId"`
	OrderType          int64              `json:"orderType"`
	Uid                string             `json:"uid"`
	AppId              string             `json:"appId"`
	AddressId          string             `json:"addressId"`
	DeliveryInfoVO     DeliveryInfoVO     `json:"deliveryInfoVO"`
	Remark             string             `json:"remark"`
	StoreInfo          StoreInfo          `json:"storeInfo"`
	ShortageDesc       string             `json:"shortageDesc"`
	PayMethodId        string             `json:"payMethodId"`
	CouponList         []string           `json:"couponList"`
}

type IOSCommitPayParam struct {
	CommitPayParam
	Amount       string `json:"amount"`
	TradeType    string `json:"tradeType"`
	PurchaserId  string `json:"purchaserId"`
	IsSelfPickup int64  `json:"isSelfPickup"`
}

type MiniProgramCommitPayParam struct {
	CommitPayParam
	Amount                int64  `json:"amount"`
	LabelList             string `json:"labelList"`
	IsSelectShoppingNotes bool   `json:"isSelectShoppingNotes"`
	SaasId                string `json:"saasId"`
	AppId                 string `json:"appId"`
}

// Page

var GetPageDataAPI = "https://api-sams.walmartmobile.cn/api/v1/sams/decoration/portal/show/getPageData"

type GetPageDataParam struct {
	Uid           string  `json:"uid"`
	PageContentId string  `json:"pageContentId"`
	Authorize     bool    `json:"authorize"`
	Latitude      string  `json:"latitude"`
	Longitude     string  `json:"longitude"`
	AddressInfo   Address `json:"addressInfo"`
}
