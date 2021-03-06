package baas

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

// GetBcSchema invokes the baas.GetBcSchema API synchronously
// api document: https://help.aliyun.com/api/baas/getbcschema.html
func (client *Client) GetBcSchema(request *GetBcSchemaRequest) (response *GetBcSchemaResponse, err error) {
	response = CreateGetBcSchemaResponse()
	err = client.DoAction(request, response)
	return
}

// GetBcSchemaWithChan invokes the baas.GetBcSchema API asynchronously
// api document: https://help.aliyun.com/api/baas/getbcschema.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBcSchemaWithChan(request *GetBcSchemaRequest) (<-chan *GetBcSchemaResponse, <-chan error) {
	responseChan := make(chan *GetBcSchemaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBcSchema(request)
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

// GetBcSchemaWithCallback invokes the baas.GetBcSchema API asynchronously
// api document: https://help.aliyun.com/api/baas/getbcschema.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBcSchemaWithCallback(request *GetBcSchemaRequest, callback func(response *GetBcSchemaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBcSchemaResponse
		var err error
		defer close(result)
		response, err = client.GetBcSchema(request)
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

// GetBcSchemaRequest is the request struct for api GetBcSchema
type GetBcSchemaRequest struct {
	*requests.RpcRequest
	Bizid string `position:"Body" name:"Bizid"`
}

// GetBcSchemaResponse is the response struct for api GetBcSchema
type GetBcSchemaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateGetBcSchemaRequest creates a request to invoke GetBcSchema API
func CreateGetBcSchemaRequest() (request *GetBcSchemaRequest) {
	request = &GetBcSchemaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "GetBcSchema", "", "")
	return
}

// CreateGetBcSchemaResponse creates a response to parse from GetBcSchema response
func CreateGetBcSchemaResponse() (response *GetBcSchemaResponse) {
	response = &GetBcSchemaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
