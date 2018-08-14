package slb

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

// DescribeMasterSlaveVServerGroups invokes the slb.DescribeMasterSlaveVServerGroups API synchronously
// api document: https://help.aliyun.com/api/slb/describemasterslavevservergroups.html
func (client *Client) DescribeMasterSlaveVServerGroups(request *DescribeMasterSlaveVServerGroupsRequest) (response *DescribeMasterSlaveVServerGroupsResponse, err error) {
	response = CreateDescribeMasterSlaveVServerGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMasterSlaveVServerGroupsWithChan invokes the slb.DescribeMasterSlaveVServerGroups API asynchronously
// api document: https://help.aliyun.com/api/slb/describemasterslavevservergroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMasterSlaveVServerGroupsWithChan(request *DescribeMasterSlaveVServerGroupsRequest) (<-chan *DescribeMasterSlaveVServerGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeMasterSlaveVServerGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMasterSlaveVServerGroups(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeMasterSlaveVServerGroupsWithCallback invokes the slb.DescribeMasterSlaveVServerGroups API asynchronously
// api document: https://help.aliyun.com/api/slb/describemasterslavevservergroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMasterSlaveVServerGroupsWithCallback(request *DescribeMasterSlaveVServerGroupsRequest, callback func(response *DescribeMasterSlaveVServerGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMasterSlaveVServerGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeMasterSlaveVServerGroups(request)
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

// DescribeMasterSlaveVServerGroupsRequest is the request struct for api DescribeMasterSlaveVServerGroups
type DescribeMasterSlaveVServerGroupsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeMasterSlaveVServerGroupsResponse is the response struct for api DescribeMasterSlaveVServerGroups
type DescribeMasterSlaveVServerGroupsResponse struct {
	*responses.BaseResponse
	RequestId                string                                                     `json:"RequestId" xml:"RequestId"`
	MasterSlaveVServerGroups MasterSlaveVServerGroupsInDescribeMasterSlaveVServerGroups `json:"MasterSlaveVServerGroups" xml:"MasterSlaveVServerGroups"`
}

// CreateDescribeMasterSlaveVServerGroupsRequest creates a request to invoke DescribeMasterSlaveVServerGroups API
func CreateDescribeMasterSlaveVServerGroupsRequest() (request *DescribeMasterSlaveVServerGroupsRequest) {
	request = &DescribeMasterSlaveVServerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveVServerGroups", "slb", "openAPI")
	return
}

// CreateDescribeMasterSlaveVServerGroupsResponse creates a response to parse from DescribeMasterSlaveVServerGroups response
func CreateDescribeMasterSlaveVServerGroupsResponse() (response *DescribeMasterSlaveVServerGroupsResponse) {
	response = &DescribeMasterSlaveVServerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}