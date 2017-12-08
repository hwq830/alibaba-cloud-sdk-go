
package ess

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

import (
"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) CreateScalingConfiguration(request *CreateScalingConfigurationRequest) (response *CreateScalingConfigurationResponse, err error) {
response = CreateCreateScalingConfigurationResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) CreateScalingConfigurationWithChan(request *CreateScalingConfigurationRequest) (<-chan *CreateScalingConfigurationResponse, <-chan error) {
responseChan := make(chan *CreateScalingConfigurationResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateScalingConfiguration(request)
responseChan <- response
errChan <- err
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

func (client *Client) CreateScalingConfigurationWithCallback(request *CreateScalingConfigurationRequest, callback func(response *CreateScalingConfigurationResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateScalingConfigurationResponse
var err error
defer close(result)
response, err = client.CreateScalingConfiguration(request)
callback(response, err)
result <- 1
})
if err != nil {
defer close(result)
callback(nil, err)
result <- 0
}
return result
}

type CreateScalingConfigurationRequest struct {
*requests.RpcRequest
                DataDisk1SnapshotId  string `position:"Query" name:"DataDisk.1.SnapshotId"`
                SystemDiskCategory  string `position:"Query" name:"SystemDisk.Category"`
                DataDisk3Size  string `position:"Query" name:"DataDisk.3.Size"`
                DataDisk2Device  string `position:"Query" name:"DataDisk.2.Device"`
                SecurityEnhancementStrategy  string `position:"Query" name:"SecurityEnhancementStrategy"`
                DataDisk1Device  string `position:"Query" name:"DataDisk.1.Device"`
                DataDisk4Size  string `position:"Query" name:"DataDisk.4.Size"`
                InstanceType  string `position:"Query" name:"InstanceType"`
                UserData  string `position:"Query" name:"UserData"`
                DataDisk3DeleteWithInstance  string `position:"Query" name:"DataDisk.3.DeleteWithInstance"`
                DataDisk4DeleteWithInstance  string `position:"Query" name:"DataDisk.4.DeleteWithInstance"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                DataDisk2Category  string `position:"Query" name:"DataDisk.2.Category"`
                DataDisk1DeleteWithInstance  string `position:"Query" name:"DataDisk.1.DeleteWithInstance"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                LoadBalancerWeight  string `position:"Query" name:"LoadBalancerWeight"`
                DataDisk3Device  string `position:"Query" name:"DataDisk.3.Device"`
                SystemDiskSize  string `position:"Query" name:"SystemDisk.Size"`
                DataDisk4Device  string `position:"Query" name:"DataDisk.4.Device"`
                DataDisk2DeleteWithInstance  string `position:"Query" name:"DataDisk.2.DeleteWithInstance"`
                InstanceTypes  *[]string `position:"Query" name:"InstanceTypes"  type:"Repeated"`
                Tags  string `position:"Query" name:"Tags"`
                DataDisk2Size  string `position:"Query" name:"DataDisk.2.Size"`
                InternetMaxBandwidthOut  string `position:"Query" name:"InternetMaxBandwidthOut"`
                InternetChargeType  string `position:"Query" name:"InternetChargeType"`
                SecurityGroupId  string `position:"Query" name:"SecurityGroupId"`
                DataDisk4SnapshotId  string `position:"Query" name:"DataDisk.4.SnapshotId"`
                DataDisk4Category  string `position:"Query" name:"DataDisk.4.Category"`
                ScalingConfigurationName  string `position:"Query" name:"ScalingConfigurationName"`
                KeyPairName  string `position:"Query" name:"KeyPairName"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                ImageId  string `position:"Query" name:"ImageId"`
                IoOptimized  string `position:"Query" name:"IoOptimized"`
                ScalingGroupId  string `position:"Query" name:"ScalingGroupId"`
                DataDisk2SnapshotId  string `position:"Query" name:"DataDisk.2.SnapshotId"`
                DataDisk1Category  string `position:"Query" name:"DataDisk.1.Category"`
                DataDisk3SnapshotId  string `position:"Query" name:"DataDisk.3.SnapshotId"`
                InstanceName  string `position:"Query" name:"InstanceName"`
                RamRoleName  string `position:"Query" name:"RamRoleName"`
                DataDisk1Size  string `position:"Query" name:"DataDisk.1.Size"`
                InternetMaxBandwidthIn  string `position:"Query" name:"InternetMaxBandwidthIn"`
                DataDisk3Category  string `position:"Query" name:"DataDisk.3.Category"`
}


type CreateScalingConfigurationResponse struct {
*responses.BaseResponse
            ScalingConfigurationId     string `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateScalingConfigurationRequest() (request *CreateScalingConfigurationRequest) {
request = &CreateScalingConfigurationRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingConfiguration", "", "")
return
}

func CreateCreateScalingConfigurationResponse() (response *CreateScalingConfigurationResponse) {
response = &CreateScalingConfigurationResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}
