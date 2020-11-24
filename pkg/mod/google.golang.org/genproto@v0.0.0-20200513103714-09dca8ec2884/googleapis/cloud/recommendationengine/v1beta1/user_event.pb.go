// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommendationengine/v1beta1/user_event.proto

package recommendationengine

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// User event source.
type UserEvent_EventSource int32

const (
	// Unspecified event source.
	UserEvent_EVENT_SOURCE_UNSPECIFIED UserEvent_EventSource = 0
	// The event is ingested via a javascript pixel or Recommendations AI Tag
	// through automl datalayer or JS Macros.
	UserEvent_AUTOML UserEvent_EventSource = 1
	// The event is ingested via Recommendations AI Tag through Enhanced
	// Ecommerce datalayer.
	UserEvent_ECOMMERCE UserEvent_EventSource = 2
	// The event is ingested via Import user events API.
	UserEvent_BATCH_UPLOAD UserEvent_EventSource = 3
)

var UserEvent_EventSource_name = map[int32]string{
	0: "EVENT_SOURCE_UNSPECIFIED",
	1: "AUTOML",
	2: "ECOMMERCE",
	3: "BATCH_UPLOAD",
}

var UserEvent_EventSource_value = map[string]int32{
	"EVENT_SOURCE_UNSPECIFIED": 0,
	"AUTOML":                   1,
	"ECOMMERCE":                2,
	"BATCH_UPLOAD":             3,
}

func (x UserEvent_EventSource) String() string {
	return proto.EnumName(UserEvent_EventSource_name, int32(x))
}

func (UserEvent_EventSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f86607756c01e62b, []int{0, 0}
}

// UserEvent captures all metadata information recommendation engine needs to
// know about how end users interact with customers' website.
type UserEvent struct {
	// Required. User event type. Allowed values are:
	//
	// * `add-to-cart` Products being added to cart.
	// * `add-to-list` Items being added to a list (shopping list, favorites
	//   etc).
	// * `category-page-view` Special pages such as sale or promotion pages
	//   viewed.
	// * `checkout-start` User starting a checkout process.
	// * `detail-page-view` Products detail page viewed.
	// * `home-page-view` Homepage viewed.
	// * `page-visit` Generic page visits not included in the event types above.
	// * `purchase-complete` User finishing a purchase.
	// * `refund` Purchased items being refunded or returned.
	// * `remove-from-cart` Products being removed from cart.
	// * `remove-from-list` Items being removed from a list.
	// * `search` Product search.
	// * `shopping-cart-page-view` User viewing a shopping cart.
	// * `impression` List of items displayed. Used by Google Tag Manager.
	EventType string `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	// Required. User information.
	UserInfo *UserInfo `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	// Optional. User event detailed information common across different
	// recommendation types.
	EventDetail *EventDetail `protobuf:"bytes,3,opt,name=event_detail,json=eventDetail,proto3" json:"event_detail,omitempty"`
	// Optional. Retail product specific user event metadata.
	//
	// This field is required for the following event types:
	//
	// * `add-to-cart`
	// * `add-to-list`
	// * `category-page-view`
	// * `checkout-start`
	// * `detail-page-view`
	// * `purchase-complete`
	// * `refund`
	// * `remove-from-cart`
	// * `remove-from-list`
	// * `search`
	//
	// This field is optional for the following event types:
	//
	// * `page-visit`
	// * `shopping-cart-page-view` - note that 'product_event_detail' should be
	//   set for this unless the shopping cart is empty.
	//
	// This field is not allowed for the following event types:
	//
	// * `home-page-view`
	ProductEventDetail *ProductEventDetail `protobuf:"bytes,4,opt,name=product_event_detail,json=productEventDetail,proto3" json:"product_event_detail,omitempty"`
	// Optional. Only required for ImportUserEvents method. Timestamp of user
	// event created.
	EventTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Optional. This field should *not* be set when using JavaScript pixel
	// or the Recommendations AI Tag. Defaults to `EVENT_SOURCE_UNSPECIFIED`.
	EventSource          UserEvent_EventSource `protobuf:"varint,6,opt,name=event_source,json=eventSource,proto3,enum=google.cloud.recommendationengine.v1beta1.UserEvent_EventSource" json:"event_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserEvent) Reset()         { *m = UserEvent{} }
func (m *UserEvent) String() string { return proto.CompactTextString(m) }
func (*UserEvent) ProtoMessage()    {}
func (*UserEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f86607756c01e62b, []int{0}
}

func (m *UserEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEvent.Unmarshal(m, b)
}
func (m *UserEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEvent.Marshal(b, m, deterministic)
}
func (m *UserEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEvent.Merge(m, src)
}
func (m *UserEvent) XXX_Size() int {
	return xxx_messageInfo_UserEvent.Size(m)
}
func (m *UserEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UserEvent proto.InternalMessageInfo

func (m *UserEvent) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *UserEvent) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *UserEvent) GetEventDetail() *EventDetail {
	if m != nil {
		return m.EventDetail
	}
	return nil
}

func (m *UserEvent) GetProductEventDetail() *ProductEventDetail {
	if m != nil {
		return m.ProductEventDetail
	}
	return nil
}

func (m *UserEvent) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *UserEvent) GetEventSource() UserEvent_EventSource {
	if m != nil {
		return m.EventSource
	}
	return UserEvent_EVENT_SOURCE_UNSPECIFIED
}

// Information of end users.
type UserInfo struct {
	// Required. A unique identifier for tracking visitors with a length limit of
	// 128 bytes.
	//
	// For example, this could be implemented with a http cookie, which should be
	// able to uniquely identify a visitor on a single device. This unique
	// identifier should not change if the visitor log in/out of the website.
	// Maximum length 128 bytes. Cannot be empty.
	VisitorId string `protobuf:"bytes,1,opt,name=visitor_id,json=visitorId,proto3" json:"visitor_id,omitempty"`
	// Optional. Unique identifier for logged-in user with a length limit of 128
	// bytes. Required only for logged-in users.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Optional. IP address of the user. This could be either IPv4 (e.g. 104.133.9.80) or
	// IPv6 (e.g. 2001:0db8:85a3:0000:0000:8a2e:0370:7334). This should *not* be
	// set when using the javascript pixel or if `direct_user_request` is set.
	// Used to extract location information for personalization.
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// Optional. User agent as included in the HTTP header. UTF-8 encoded string
	// with a length limit of 1 KiB.
	//
	// This should *not* be set when using the JavaScript pixel or if
	// `directUserRequest` is set.
	UserAgent string `protobuf:"bytes,4,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// Optional. Indicates if the request is made directly from the end user
	// in which case the user_agent and ip_address fields can be populated
	// from the HTTP request. This should *not* be set when using the javascript
	// pixel. This flag should be set only if the API request is made directly
	// from the end user such as a mobile app (and not if a gateway or a server is
	// processing and pushing the user events).
	DirectUserRequest    bool     `protobuf:"varint,5,opt,name=direct_user_request,json=directUserRequest,proto3" json:"direct_user_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f86607756c01e62b, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetVisitorId() string {
	if m != nil {
		return m.VisitorId
	}
	return ""
}

func (m *UserInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserInfo) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *UserInfo) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *UserInfo) GetDirectUserRequest() bool {
	if m != nil {
		return m.DirectUserRequest
	}
	return false
}

// User event details shared by all recommendation types.
type EventDetail struct {
	// Optional. Complete url (window.location.href) of the user's current page.
	// When using the JavaScript pixel, this value is filled in automatically.
	// Maximum length 5KB.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Optional. The referrer url of the current page. When using
	// the JavaScript pixel, this value is filled in automatically.
	ReferrerUri string `protobuf:"bytes,6,opt,name=referrer_uri,json=referrerUri,proto3" json:"referrer_uri,omitempty"`
	// Optional. A unique id of a web page view.
	// This should be kept the same for all user events triggered from the same
	// pageview. For example, an item detail page view could trigger multiple
	// events as the user is browsing the page.
	// The `pageViewId` property should be kept the same for all these events so
	// that they can be grouped together properly. This `pageViewId` will be
	// automatically generated if using the JavaScript pixel.
	PageViewId string `protobuf:"bytes,2,opt,name=page_view_id,json=pageViewId,proto3" json:"page_view_id,omitempty"`
	// Optional. A list of identifiers for the independent experiment groups
	// this user event belongs to. This is used to distinguish between user events
	// associated with different experiment setups (e.g. using Recommendation
	// Engine system, using different recommendation models).
	ExperimentIds []string `protobuf:"bytes,3,rep,name=experiment_ids,json=experimentIds,proto3" json:"experiment_ids,omitempty"`
	// Optional. Recommendation token included in the recommendation prediction
	// response.
	//
	// This field enables accurate attribution of recommendation model
	// performance.
	//
	// This token enables us to accurately attribute page view or purchase back to
	// the event and the particular predict response containing this
	// clicked/purchased item. If user clicks on product K in the recommendation
	// results, pass the `PredictResponse.recommendationToken` property as a url
	// parameter to product K's page. When recording events on product K's page,
	// log the PredictResponse.recommendation_token to this field.
	//
	// Optional, but highly encouraged for user events that are the result of a
	// recommendation prediction query.
	RecommendationToken string `protobuf:"bytes,4,opt,name=recommendation_token,json=recommendationToken,proto3" json:"recommendation_token,omitempty"`
	// Optional. Extra user event features to include in the recommendation
	// model.
	//
	// For product recommendation, an example of extra user information is
	// traffic_channel, i.e. how user arrives at the site. Users can arrive
	// at the site by coming to the site directly, or coming through Google
	// search, and etc.
	EventAttributes      *FeatureMap `protobuf:"bytes,5,opt,name=event_attributes,json=eventAttributes,proto3" json:"event_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EventDetail) Reset()         { *m = EventDetail{} }
func (m *EventDetail) String() string { return proto.CompactTextString(m) }
func (*EventDetail) ProtoMessage()    {}
func (*EventDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_f86607756c01e62b, []int{2}
}

func (m *EventDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventDetail.Unmarshal(m, b)
}
func (m *EventDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventDetail.Marshal(b, m, deterministic)
}
func (m *EventDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDetail.Merge(m, src)
}
func (m *EventDetail) XXX_Size() int {
	return xxx_messageInfo_EventDetail.Size(m)
}
func (m *EventDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDetail.DiscardUnknown(m)
}

var xxx_messageInfo_EventDetail proto.InternalMessageInfo

func (m *EventDetail) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *EventDetail) GetReferrerUri() string {
	if m != nil {
		return m.ReferrerUri
	}
	return ""
}

func (m *EventDetail) GetPageViewId() string {
	if m != nil {
		return m.PageViewId
	}
	return ""
}

func (m *EventDetail) GetExperimentIds() []string {
	if m != nil {
		return m.ExperimentIds
	}
	return nil
}

func (m *EventDetail) GetRecommendationToken() string {
	if m != nil {
		return m.RecommendationToken
	}
	return ""
}

func (m *EventDetail) GetEventAttributes() *FeatureMap {
	if m != nil {
		return m.EventAttributes
	}
	return nil
}

// ProductEventDetail captures user event information specific to retail
// products.
type ProductEventDetail struct {
	// Required for `search` events. Other event types should not set this field.
	// The user's search query as UTF-8 encoded text with a length limit of 5 KiB.
	SearchQuery string `protobuf:"bytes,1,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty"`
	// Required for `category-page-view` events. Other event types should not set
	// this field.
	// The categories associated with a category page.
	// Category pages include special pages such as sales or promotions. For
	// instance, a special sale page may have the category hierarchy:
	// categories : ["Sales", "2017 Black Friday Deals"].
	PageCategories []*CatalogItem_CategoryHierarchy `protobuf:"bytes,2,rep,name=page_categories,json=pageCategories,proto3" json:"page_categories,omitempty"`
	// The main product details related to the event.
	//
	// This field is required for the following event types:
	//
	// * `add-to-cart`
	// * `add-to-list`
	// * `checkout-start`
	// * `detail-page-view`
	// * `purchase-complete`
	// * `refund`
	// * `remove-from-cart`
	// * `remove-from-list`
	//
	// This field is optional for the following event types:
	//
	// * `page-visit`
	// * `shopping-cart-page-view` - note that 'product_details' should be set for
	//   this unless the shopping cart is empty.
	//
	// This field is not allowed for the following event types:
	//
	// * `category-page-view`
	// * `home-page-view`
	// * `search`
	ProductDetails []*ProductDetail `protobuf:"bytes,3,rep,name=product_details,json=productDetails,proto3" json:"product_details,omitempty"`
	// Required for `add-to-list` and `remove-from-list` events. The id or name of
	// the list that the item is being added to or removed from. Other event types
	// should not set this field.
	ListId string `protobuf:"bytes,4,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
	// Optional. The id or name of the associated shopping cart. This id is used
	// to associate multiple items added or present in the cart before purchase.
	//
	// This can only be set for `add-to-cart`, `remove-from-cart`,
	// `checkout-start`, `purchase-complete`, or `shopping-cart-page-view` events.
	CartId string `protobuf:"bytes,5,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	// Optional. A transaction represents the entire purchase transaction.
	// Required for `purchase-complete` events. Optional for `checkout-start`
	// events. Other event types should not set this field.
	PurchaseTransaction  *PurchaseTransaction `protobuf:"bytes,6,opt,name=purchase_transaction,json=purchaseTransaction,proto3" json:"purchase_transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProductEventDetail) Reset()         { *m = ProductEventDetail{} }
func (m *ProductEventDetail) String() string { return proto.CompactTextString(m) }
func (*ProductEventDetail) ProtoMessage()    {}
func (*ProductEventDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_f86607756c01e62b, []int{3}
}

func (m *ProductEventDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductEventDetail.Unmarshal(m, b)
}
func (m *ProductEventDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductEventDetail.Marshal(b, m, deterministic)
}
func (m *ProductEventDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductEventDetail.Merge(m, src)
}
func (m *ProductEventDetail) XXX_Size() int {
	return xxx_messageInfo_ProductEventDetail.Size(m)
}
func (m *ProductEventDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductEventDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ProductEventDetail proto.InternalMessageInfo

func (m *ProductEventDetail) GetSearchQuery() string {
	if m != nil {
		return m.SearchQuery
	}
	return ""
}

func (m *ProductEventDetail) GetPageCategories() []*CatalogItem_CategoryHierarchy {
	if m != nil {
		return m.PageCategories
	}
	return nil
}

func (m *ProductEventDetail) GetProductDetails() []*ProductDetail {
	if m != nil {
		return m.ProductDetails
	}
	return nil
}

func (m *ProductEventDetail) GetListId() string {
	if m != nil {
		return m.ListId
	}
	return ""
}

func (m *ProductEventDetail) GetCartId() string {
	if m != nil {
		return m.CartId
	}
	return ""
}

func (m *ProductEventDetail) GetPurchaseTransaction() *PurchaseTransaction {
	if m != nil {
		return m.PurchaseTransaction
	}
	return nil
}

// A transaction represents the entire purchase transaction.
type PurchaseTransaction struct {
	// Optional. The transaction ID with a length limit of 128 bytes.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Total revenue or grand total associated with the transaction.
	// This value include shipping, tax, or other adjustments to total revenue
	// that you want to include as part of your revenue calculations. This field
	// is not required if the event type is `refund`.
	Revenue float32 `protobuf:"fixed32,2,opt,name=revenue,proto3" json:"revenue,omitempty"`
	// Optional. All the taxes associated with the transaction.
	Taxes map[string]float32 `protobuf:"bytes,3,rep,name=taxes,proto3" json:"taxes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	// Optional. All the costs associated with the product. These can be
	// manufacturing costs, shipping expenses not borne by the end user, or any
	// other costs.
	//
	// Total product cost such that
	//   profit = revenue - (sum(taxes) + sum(costs))
	// If product_cost is not set, then
	//   profit = revenue - tax - shipping - sum(CatalogItem.costs).
	//
	// If CatalogItem.cost is not specified for one of the items, CatalogItem.cost
	// based profit *cannot* be calculated for this Transaction.
	Costs map[string]float32 `protobuf:"bytes,4,rep,name=costs,proto3" json:"costs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	// Required. Currency code. Use three-character ISO-4217 code. This field
	// is not required if the event type is `refund`.
	CurrencyCode         string   `protobuf:"bytes,6,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurchaseTransaction) Reset()         { *m = PurchaseTransaction{} }
func (m *PurchaseTransaction) String() string { return proto.CompactTextString(m) }
func (*PurchaseTransaction) ProtoMessage()    {}
func (*PurchaseTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f86607756c01e62b, []int{4}
}

func (m *PurchaseTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseTransaction.Unmarshal(m, b)
}
func (m *PurchaseTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseTransaction.Marshal(b, m, deterministic)
}
func (m *PurchaseTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseTransaction.Merge(m, src)
}
func (m *PurchaseTransaction) XXX_Size() int {
	return xxx_messageInfo_PurchaseTransaction.Size(m)
}
func (m *PurchaseTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseTransaction proto.InternalMessageInfo

func (m *PurchaseTransaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PurchaseTransaction) GetRevenue() float32 {
	if m != nil {
		return m.Revenue
	}
	return 0
}

func (m *PurchaseTransaction) GetTaxes() map[string]float32 {
	if m != nil {
		return m.Taxes
	}
	return nil
}

func (m *PurchaseTransaction) GetCosts() map[string]float32 {
	if m != nil {
		return m.Costs
	}
	return nil
}

func (m *PurchaseTransaction) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

// Detailed product information associated with a user event.
type ProductDetail struct {
	// Required. Catalog item ID. UTF-8 encoded string with a length limit of 128
	// characters.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional. Currency code for price/costs. Use three-character ISO-4217
	// code. Required only if originalPrice or displayPrice is set.
	CurrencyCode string `protobuf:"bytes,2,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Optional. Original price of the product. If provided, this will override
	// the original price in Catalog for this product.
	OriginalPrice float32 `protobuf:"fixed32,3,opt,name=original_price,json=originalPrice,proto3" json:"original_price,omitempty"`
	// Optional. Display price of the product (e.g. discounted price). If
	// provided, this will override the display price in Catalog for this product.
	DisplayPrice float32 `protobuf:"fixed32,4,opt,name=display_price,json=displayPrice,proto3" json:"display_price,omitempty"`
	// Optional. Item stock state. If provided, this overrides the stock state
	// in Catalog for items in this event.
	StockState ProductCatalogItem_StockState `protobuf:"varint,5,opt,name=stock_state,json=stockState,proto3,enum=google.cloud.recommendationengine.v1beta1.ProductCatalogItem_StockState" json:"stock_state,omitempty"`
	// Optional. Quantity of the product associated with the user event. For
	// example, this field will be 2 if two products are added to the shopping
	// cart for `add-to-cart` event. Required for `add-to-cart`, `add-to-list`,
	// `remove-from-cart`, `checkout-start`, `purchase-complete`, `refund` event
	// types.
	Quantity int32 `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// Optional. Quantity of the products in stock when a user event happens.
	// Optional. If provided, this overrides the available quantity in Catalog for
	// this event. and can only be set if `stock_status` is set to `IN_STOCK`.
	//
	// Note that if an item is out of stock, you must set the `stock_state` field
	// to be `OUT_OF_STOCK`. Leaving this field unspecified / as zero is not
	// sufficient to mark the item out of stock.
	AvailableQuantity int32 `protobuf:"varint,7,opt,name=available_quantity,json=availableQuantity,proto3" json:"available_quantity,omitempty"`
	// Optional. Extra features associated with a product in the user event.
	ItemAttributes       *FeatureMap `protobuf:"bytes,8,opt,name=item_attributes,json=itemAttributes,proto3" json:"item_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ProductDetail) Reset()         { *m = ProductDetail{} }
func (m *ProductDetail) String() string { return proto.CompactTextString(m) }
func (*ProductDetail) ProtoMessage()    {}
func (*ProductDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_f86607756c01e62b, []int{5}
}

func (m *ProductDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductDetail.Unmarshal(m, b)
}
func (m *ProductDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductDetail.Marshal(b, m, deterministic)
}
func (m *ProductDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductDetail.Merge(m, src)
}
func (m *ProductDetail) XXX_Size() int {
	return xxx_messageInfo_ProductDetail.Size(m)
}
func (m *ProductDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ProductDetail proto.InternalMessageInfo

func (m *ProductDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProductDetail) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *ProductDetail) GetOriginalPrice() float32 {
	if m != nil {
		return m.OriginalPrice
	}
	return 0
}

func (m *ProductDetail) GetDisplayPrice() float32 {
	if m != nil {
		return m.DisplayPrice
	}
	return 0
}

func (m *ProductDetail) GetStockState() ProductCatalogItem_StockState {
	if m != nil {
		return m.StockState
	}
	return ProductCatalogItem_STOCK_STATE_UNSPECIFIED
}

func (m *ProductDetail) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *ProductDetail) GetAvailableQuantity() int32 {
	if m != nil {
		return m.AvailableQuantity
	}
	return 0
}

func (m *ProductDetail) GetItemAttributes() *FeatureMap {
	if m != nil {
		return m.ItemAttributes
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.recommendationengine.v1beta1.UserEvent_EventSource", UserEvent_EventSource_name, UserEvent_EventSource_value)
	proto.RegisterType((*UserEvent)(nil), "google.cloud.recommendationengine.v1beta1.UserEvent")
	proto.RegisterType((*UserInfo)(nil), "google.cloud.recommendationengine.v1beta1.UserInfo")
	proto.RegisterType((*EventDetail)(nil), "google.cloud.recommendationengine.v1beta1.EventDetail")
	proto.RegisterType((*ProductEventDetail)(nil), "google.cloud.recommendationengine.v1beta1.ProductEventDetail")
	proto.RegisterType((*PurchaseTransaction)(nil), "google.cloud.recommendationengine.v1beta1.PurchaseTransaction")
	proto.RegisterMapType((map[string]float32)(nil), "google.cloud.recommendationengine.v1beta1.PurchaseTransaction.CostsEntry")
	proto.RegisterMapType((map[string]float32)(nil), "google.cloud.recommendationengine.v1beta1.PurchaseTransaction.TaxesEntry")
	proto.RegisterType((*ProductDetail)(nil), "google.cloud.recommendationengine.v1beta1.ProductDetail")
}

func init() {
	proto.RegisterFile("google/cloud/recommendationengine/v1beta1/user_event.proto", fileDescriptor_f86607756c01e62b)
}

var fileDescriptor_f86607756c01e62b = []byte{
	// 1194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xc7, 0xe7, 0xc4, 0x89, 0xd7, 0x8e, 0xe3, 0x6e, 0x82, 0xb0, 0xa2, 0xa2, 0x06, 0x4b, 0xa0,
	0x14, 0x81, 0xad, 0xba, 0xa2, 0x54, 0x46, 0x54, 0x38, 0xee, 0x95, 0x5a, 0x6a, 0x9a, 0xf4, 0x62,
	0x47, 0x08, 0x15, 0x4e, 0xeb, 0xbb, 0xc9, 0x65, 0xd5, 0xbb, 0xdb, 0xcb, 0xee, 0x9e, 0x5b, 0x7f,
	0x1a, 0x24, 0x1e, 0x79, 0xe5, 0x15, 0xf1, 0x0e, 0x0f, 0x7c, 0x86, 0x3e, 0xf3, 0x29, 0xd0, 0xee,
	0xde, 0xf9, 0x4f, 0x93, 0x87, 0x58, 0xe5, 0xcd, 0x3b, 0xbf, 0xdf, 0xfc, 0x66, 0x76, 0x66, 0x76,
	0x7c, 0xa8, 0x1b, 0x30, 0x16, 0x84, 0xd0, 0xf6, 0x42, 0x96, 0xfa, 0x6d, 0x0e, 0x1e, 0x8b, 0x22,
	0x88, 0x7d, 0x22, 0x29, 0x8b, 0x21, 0x0e, 0x68, 0x0c, 0xed, 0xc9, 0xbd, 0x31, 0x48, 0x72, 0xaf,
	0x9d, 0x0a, 0xe0, 0x2e, 0x4c, 0x20, 0x96, 0xad, 0x84, 0x33, 0xc9, 0xf0, 0x5d, 0xe3, 0xdb, 0xd2,
	0xbe, 0xad, 0xeb, 0x7c, 0x5b, 0x99, 0xef, 0xde, 0x9d, 0x2c, 0x0c, 0x49, 0x68, 0xfb, 0x9c, 0x42,
	0xe8, 0xbb, 0x63, 0xb8, 0x20, 0x13, 0xca, 0xb8, 0xd1, 0xda, 0xfb, 0xfa, 0xe6, 0x79, 0x78, 0x44,
	0x92, 0x90, 0x05, 0x99, 0xe3, 0x83, 0x15, 0x1c, 0x59, 0x14, 0xb1, 0x38, 0xf3, 0xcb, 0x33, 0xd2,
	0xa7, 0x71, 0x7a, 0xde, 0x96, 0x34, 0x02, 0x21, 0x49, 0x94, 0x64, 0x84, 0xdb, 0x0b, 0x29, 0x93,
	0x38, 0x66, 0x52, 0x4b, 0x0a, 0x83, 0x36, 0xdf, 0xae, 0xa1, 0xf2, 0x48, 0x00, 0xb7, 0x55, 0x3d,
	0x70, 0x13, 0x21, 0x5d, 0x18, 0x57, 0x4e, 0x13, 0x68, 0x14, 0xf6, 0x0b, 0x07, 0xe5, 0xc3, 0xe2,
	0xdb, 0x9e, 0xe5, 0x94, 0xb5, 0x79, 0x38, 0x4d, 0x00, 0x8f, 0x50, 0x59, 0x57, 0x90, 0xc6, 0xe7,
	0xac, 0x61, 0xed, 0x17, 0x0e, 0x2a, 0x9d, 0xfb, 0xad, 0x1b, 0x57, 0xb0, 0xa5, 0x82, 0x0d, 0xe2,
	0x73, 0x66, 0x74, 0x37, 0xd3, 0xec, 0x88, 0x7f, 0x46, 0x55, 0x13, 0xda, 0x07, 0x49, 0x68, 0xd8,
	0x28, 0x6a, 0xe5, 0x07, 0x2b, 0x28, 0xeb, 0x2b, 0x3c, 0xd6, 0xde, 0x4a, 0xbc, 0xe0, 0x54, 0x60,
	0x6e, 0xc1, 0x12, 0xed, 0x26, 0x9c, 0xf9, 0xa9, 0x27, 0xdd, 0xa5, 0x38, 0x6b, 0x3a, 0xce, 0xb7,
	0x2b, 0xc4, 0x39, 0x31, 0x32, 0x57, 0xc2, 0xe1, 0xe4, 0x0a, 0x80, 0x1f, 0xcd, 0x0a, 0x4a, 0x23,
	0x68, 0xac, 0xeb, 0x58, 0x7b, 0x79, 0xac, 0xbc, 0x65, 0xad, 0x61, 0xde, 0x32, 0x23, 0x94, 0x15,
	0x9b, 0x46, 0x80, 0x2f, 0xf2, 0xaa, 0x08, 0x96, 0x72, 0x0f, 0x1a, 0xa5, 0xfd, 0xc2, 0x41, 0xad,
	0xf3, 0xdd, 0x8a, 0xf5, 0xd6, 0x19, 0x99, 0xfa, 0x9c, 0x6a, 0x9d, 0xc5, 0xfa, 0x18, 0x4b, 0xf3,
	0x07, 0x54, 0x59, 0x20, 0xe0, 0xdb, 0xa8, 0x61, 0x9f, 0xd9, 0xcf, 0x87, 0xee, 0xe9, 0xf1, 0xc8,
	0xe9, 0xdb, 0xee, 0xe8, 0xf9, 0xe9, 0x89, 0xdd, 0x1f, 0x3c, 0x19, 0xd8, 0x8f, 0xeb, 0x1f, 0x60,
	0x84, 0x4a, 0xbd, 0xd1, 0xf0, 0xf8, 0xe8, 0x59, 0xbd, 0x80, 0xb7, 0x50, 0xd9, 0xee, 0x1f, 0x1f,
	0x1d, 0xd9, 0x4e, 0xdf, 0xae, 0x5b, 0xb8, 0x8e, 0xaa, 0x87, 0xbd, 0x61, 0xff, 0xa9, 0x3b, 0x3a,
	0x79, 0x76, 0xdc, 0x7b, 0x5c, 0x2f, 0x36, 0xff, 0x2a, 0xa0, 0xcd, 0xbc, 0xeb, 0x6a, 0xc2, 0x26,
	0x54, 0x50, 0xc9, 0xb8, 0x4b, 0xfd, 0xa5, 0x09, 0xcb, 0xcc, 0x03, 0x1f, 0xdf, 0x46, 0x1b, 0x66,
	0xc2, 0x7c, 0x3d, 0x5f, 0x65, 0x93, 0x6d, 0x49, 0x8f, 0x8a, 0xaf, 0x14, 0x68, 0xe2, 0x12, 0xdf,
	0xe7, 0x20, 0x84, 0x1e, 0x93, 0x8c, 0x50, 0xa6, 0x49, 0xcf, 0x58, 0x15, 0x47, 0x2b, 0x90, 0x00,
	0x62, 0xa9, 0x5b, 0x9c, 0x73, 0x94, 0xb9, 0xa7, 0xac, 0xf8, 0x3e, 0xda, 0xf1, 0x29, 0x07, 0x4f,
	0xba, 0x9a, 0xca, 0xe1, 0x32, 0x05, 0x21, 0x75, 0x8f, 0x36, 0x0d, 0xf9, 0x96, 0xc1, 0x55, 0xee,
	0x8e, 0x41, 0x9b, 0x7f, 0x5a, 0x59, 0x99, 0xb2, 0xfe, 0x7e, 0x88, 0x8a, 0x29, 0xa7, 0xf3, 0x7b,
	0x14, 0x1c, 0x75, 0xc6, 0x9f, 0xa1, 0x2a, 0x87, 0x73, 0xe0, 0x1c, 0xb8, 0xab, 0xf0, 0xd2, 0x1c,
	0xaf, 0xe4, 0xc0, 0x88, 0x53, 0xfc, 0x29, 0xaa, 0x26, 0x24, 0x00, 0x77, 0x42, 0xe1, 0xf5, 0x3b,
	0xd7, 0x45, 0x0a, 0x38, 0xa3, 0xf0, 0x7a, 0xe0, 0xe3, 0xcf, 0x51, 0x0d, 0xde, 0x24, 0xc0, 0x69,
	0xa4, 0x46, 0x81, 0xfa, 0xea, 0xda, 0xc5, 0x9c, 0xb8, 0x35, 0x87, 0x06, 0xbe, 0xc0, 0x0f, 0xd0,
	0xee, 0xf2, 0x3c, 0xb8, 0x92, 0xbd, 0x82, 0x78, 0xb1, 0x08, 0x3b, 0xcb, 0x84, 0xa1, 0xc2, 0x31,
	0xa0, 0xba, 0x99, 0x34, 0x22, 0x25, 0xa7, 0xe3, 0x54, 0x82, 0xc8, 0xe6, 0xf5, 0xab, 0x15, 0xa6,
	0xed, 0x09, 0x10, 0x99, 0x72, 0x38, 0x22, 0xd9, 0x28, 0x6f, 0x6b, 0xcd, 0xde, 0x4c, 0xb2, 0xf9,
	0x47, 0x11, 0xe1, 0xab, 0x0f, 0x08, 0x7f, 0x82, 0xaa, 0x02, 0x08, 0xf7, 0x2e, 0xdc, 0xcb, 0x14,
	0xf8, 0xd4, 0x14, 0xd4, 0xa9, 0x18, 0xdb, 0x0b, 0x65, 0xc2, 0x97, 0x68, 0x5b, 0xd7, 0xca, 0x23,
	0x12, 0x02, 0xc6, 0x29, 0x88, 0x86, 0xb5, 0x5f, 0x3c, 0xa8, 0x74, 0x9e, 0xae, 0x90, 0x5f, 0xdf,
	0xec, 0xdc, 0x81, 0x84, 0x48, 0xfd, 0x56, 0x42, 0xd3, 0xa7, 0x14, 0xb8, 0x0a, 0x33, 0x75, 0x6a,
	0x2a, 0x40, 0x7f, 0xa6, 0x8f, 0x09, 0xda, 0xce, 0x77, 0x86, 0xd9, 0x16, 0xa6, 0xf0, 0x95, 0xce,
	0xc3, 0xd5, 0xd7, 0x85, 0xb9, 0xa8, 0x53, 0x4b, 0x16, 0x8f, 0x02, 0x7f, 0x84, 0x36, 0x42, 0x2a,
	0x54, 0x53, 0x4d, 0x87, 0x9c, 0x92, 0x3a, 0x9a, 0x47, 0xe0, 0x11, 0xae, 0x81, 0xf5, 0x85, 0x47,
	0xa0, 0x6c, 0x03, 0x1f, 0x4f, 0xd0, 0x6e, 0x92, 0x72, 0xef, 0x82, 0x08, 0x70, 0x25, 0x27, 0xb1,
	0x20, 0x9e, 0x0a, 0xad, 0x07, 0xad, 0xd2, 0x79, 0xb4, 0x4a, 0x7a, 0x99, 0xcc, 0x70, 0xae, 0x92,
	0x4d, 0x49, 0x72, 0x15, 0x69, 0xfe, 0x5e, 0x44, 0x3b, 0xd7, 0x78, 0xe0, 0x1d, 0x64, 0x2d, 0x3e,
	0xe7, 0x82, 0x63, 0x51, 0x1f, 0x7f, 0x8c, 0x36, 0xb8, 0xea, 0x7f, 0x0a, 0x7a, 0xb0, 0x2d, 0xf3,
	0xd0, 0x73, 0x1b, 0xf6, 0xd1, 0xba, 0x24, 0x6f, 0x20, 0xaf, 0xe9, 0xe0, 0xfd, 0x92, 0x6e, 0x0d,
	0x95, 0x96, 0x1d, 0x4b, 0x3e, 0x35, 0x19, 0x18, 0x71, 0x15, 0xc5, 0x63, 0x42, 0x8a, 0xc6, 0xda,
	0xff, 0x12, 0xa5, 0xaf, 0xb4, 0x16, 0xa3, 0x68, 0x71, 0x7c, 0x80, 0xb6, 0xbc, 0x94, 0x73, 0x88,
	0xbd, 0xa9, 0xeb, 0x31, 0x1f, 0xe6, 0x2f, 0xde, 0x72, 0xaa, 0x39, 0xd2, 0x67, 0x3e, 0xec, 0x3d,
	0x44, 0x68, 0x9e, 0x29, 0xae, 0xa3, 0xe2, 0x2b, 0xc8, 0xc7, 0x5d, 0xfd, 0xc4, 0xbb, 0x68, 0x7d,
	0x42, 0xc2, 0xbc, 0x64, 0x8e, 0x39, 0x74, 0xad, 0x87, 0x05, 0xe5, 0x39, 0x8f, 0xbe, 0x8a, 0x67,
	0xf3, 0x9f, 0x22, 0xda, 0x5a, 0x1a, 0xc3, 0x77, 0xfb, 0x65, 0xe9, 0x7e, 0x5d, 0xb9, 0xc4, 0xc2,
	0x3a, 0x5a, 0xba, 0x84, 0x5a, 0x48, 0x8c, 0xd3, 0x80, 0xc6, 0x24, 0x74, 0x13, 0x4e, 0x3d, 0xd0,
	0x7b, 0xd8, 0xca, 0x16, 0x52, 0x0e, 0x9d, 0x28, 0x44, 0xa9, 0xfa, 0x54, 0x24, 0x21, 0x99, 0x66,
	0xd4, 0xb5, 0x39, 0xb5, 0x9a, 0x21, 0x86, 0x19, 0xa3, 0x8a, 0x90, 0xcc, 0x7b, 0xe5, 0x0a, 0x49,
	0xa4, 0xf9, 0xb7, 0xac, 0xad, 0xf4, 0xba, 0xb3, 0x3b, 0x2e, 0x3e, 0xf2, 0x53, 0x25, 0x78, 0xaa,
	0xf4, 0xb2, 0xb5, 0x2a, 0x66, 0x06, 0x7c, 0x07, 0x6d, 0x5e, 0xa6, 0x24, 0x96, 0x54, 0x4e, 0x75,
	0xbf, 0xd6, 0x0d, 0x65, 0x66, 0xc4, 0x1d, 0x84, 0xc9, 0x84, 0xd0, 0x90, 0x8c, 0x43, 0x70, 0x67,
	0xd4, 0x8d, 0x39, 0xf5, 0xd6, 0x0c, 0x7e, 0x91, 0xfb, 0x78, 0x68, 0x9b, 0x4a, 0x88, 0x16, 0xd7,
	0xe8, 0xe6, 0x7b, 0xaf, 0xd1, 0x9a, 0x92, 0x9c, 0x6f, 0xd1, 0xc3, 0x5f, 0x2c, 0xf4, 0xa5, 0xc7,
	0xa2, 0x9b, 0x2b, 0x9e, 0x14, 0x7e, 0xfc, 0x29, 0x23, 0x07, 0x2c, 0x24, 0x71, 0xd0, 0x62, 0x3c,
	0x68, 0x07, 0x10, 0xeb, 0x6f, 0x90, 0xb6, 0x81, 0x48, 0x42, 0xc5, 0x0d, 0xbe, 0x3f, 0xbf, 0xb9,
	0x0e, 0xfc, 0xd5, 0x5a, 0x77, 0xec, 0x7e, 0x6f, 0xf0, 0x9b, 0x75, 0xf7, 0x7b, 0x13, 0xa7, 0xaf,
	0x93, 0x72, 0x96, 0xb8, 0xb6, 0x49, 0xea, 0xec, 0xde, 0xa1, 0x12, 0xfa, 0x3b, 0xe7, 0xbe, 0xd4,
	0xdc, 0x97, 0xd7, 0x71, 0x5f, 0x9e, 0x99, 0xa0, 0xff, 0x5a, 0x5f, 0x18, 0x6e, 0xb7, 0xab, 0xc9,
	0xdd, 0xee, 0x75, 0xec, 0x6e, 0x37, 0xa3, 0x8f, 0x4b, 0xfa, 0x62, 0xf7, 0xff, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x47, 0x3f, 0x66, 0xd8, 0x18, 0x0c, 0x00, 0x00,
}
