package edas

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

// ApplicationInfo is a nested struct in edas response
type ApplicationInfo struct {
	AppId         string `json:"AppId" xml:"AppId"`
	Port          int    `json:"Port" xml:"Port"`
	RegionId      string `json:"RegionId" xml:"RegionId"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	UserId        string `json:"UserId" xml:"UserId"`
	Owner         string `json:"Owner" xml:"Owner"`
	ClusterType   int    `json:"ClusterType" xml:"ClusterType"`
	AppName       string `json:"AppName" xml:"AppName"`
	RegionName    string `json:"RegionName" xml:"RegionName"`
	Dockerize     bool   `json:"Dockerize" xml:"Dockerize"`
	EdasId        string `json:"EdasId" xml:"EdasId"`
}
