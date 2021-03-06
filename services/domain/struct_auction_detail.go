package domain

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AuctionDetail is a nested struct in domain response
type AuctionDetail struct {
	DomainName      string  `json:"DomainName" xml:"DomainName"`
	AuctionId       string  `json:"AuctionId" xml:"AuctionId"`
	DomainType      string  `json:"DomainType" xml:"DomainType"`
	BookedPartner   string  `json:"BookedPartner" xml:"BookedPartner"`
	PartnerType     string  `json:"PartnerType" xml:"PartnerType"`
	Currency        string  `json:"Currency" xml:"Currency"`
	YourCurrentBid  float64 `json:"YourCurrentBid" xml:"YourCurrentBid"`
	YourMaxBid      float64 `json:"YourMaxBid" xml:"YourMaxBid"`
	HighBid         float64 `json:"HighBid" xml:"HighBid"`
	NextValidBid    float64 `json:"NextValidBid" xml:"NextValidBid"`
	ReserveMet      bool    `json:"ReserveMet" xml:"ReserveMet"`
	TransferInPrice float64 `json:"TransferInPrice" xml:"TransferInPrice"`
	PayPrice        float64 `json:"PayPrice" xml:"PayPrice"`
	HighBidder      string  `json:"HighBidder" xml:"HighBidder"`
	Status          string  `json:"Status" xml:"Status"`
	PayStatus       string  `json:"PayStatus" xml:"PayStatus"`
	ProduceStatus   string  `json:"ProduceStatus" xml:"ProduceStatus"`
	AuctionEndTime  int     `json:"AuctionEndTime" xml:"AuctionEndTime"`
	BookEndTime     int     `json:"BookEndTime" xml:"BookEndTime"`
	PayEndTime      int     `json:"PayEndTime" xml:"PayEndTime"`
	DeliveryTime    int     `json:"DeliveryTime" xml:"DeliveryTime"`
	FailCode        string  `json:"FailCode" xml:"FailCode"`
}
