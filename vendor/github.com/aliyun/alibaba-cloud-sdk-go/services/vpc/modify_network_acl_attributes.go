package vpc

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

// ModifyNetworkAclAttributes invokes the vpc.ModifyNetworkAclAttributes API synchronously
func (client *Client) ModifyNetworkAclAttributes(request *ModifyNetworkAclAttributesRequest) (response *ModifyNetworkAclAttributesResponse, err error) {
	response = CreateModifyNetworkAclAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyNetworkAclAttributesWithChan invokes the vpc.ModifyNetworkAclAttributes API asynchronously
func (client *Client) ModifyNetworkAclAttributesWithChan(request *ModifyNetworkAclAttributesRequest) (<-chan *ModifyNetworkAclAttributesResponse, <-chan error) {
	responseChan := make(chan *ModifyNetworkAclAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNetworkAclAttributes(request)
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

// ModifyNetworkAclAttributesWithCallback invokes the vpc.ModifyNetworkAclAttributes API asynchronously
func (client *Client) ModifyNetworkAclAttributesWithCallback(request *ModifyNetworkAclAttributesRequest, callback func(response *ModifyNetworkAclAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNetworkAclAttributesResponse
		var err error
		defer close(result)
		response, err = client.ModifyNetworkAclAttributes(request)
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

// ModifyNetworkAclAttributesRequest is the request struct for api ModifyNetworkAclAttributes
type ModifyNetworkAclAttributesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	NetworkAclId         string           `position:"Query" name:"NetworkAclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	NetworkAclName       string           `position:"Query" name:"NetworkAclName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyNetworkAclAttributesResponse is the response struct for api ModifyNetworkAclAttributes
type ModifyNetworkAclAttributesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyNetworkAclAttributesRequest creates a request to invoke ModifyNetworkAclAttributes API
func CreateModifyNetworkAclAttributesRequest() (request *ModifyNetworkAclAttributesRequest) {
	request = &ModifyNetworkAclAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNetworkAclAttributes", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyNetworkAclAttributesResponse creates a response to parse from ModifyNetworkAclAttributes response
func CreateModifyNetworkAclAttributesResponse() (response *ModifyNetworkAclAttributesResponse) {
	response = &ModifyNetworkAclAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
