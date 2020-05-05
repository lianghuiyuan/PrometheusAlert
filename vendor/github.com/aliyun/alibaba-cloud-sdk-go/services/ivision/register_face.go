package ivision

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

// RegisterFace invokes the ivision.RegisterFace API synchronously
// api document: https://help.aliyun.com/api/ivision/registerface.html
func (client *Client) RegisterFace(request *RegisterFaceRequest) (response *RegisterFaceResponse, err error) {
	response = CreateRegisterFaceResponse()
	err = client.DoAction(request, response)
	return
}

// RegisterFaceWithChan invokes the ivision.RegisterFace API asynchronously
// api document: https://help.aliyun.com/api/ivision/registerface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegisterFaceWithChan(request *RegisterFaceRequest) (<-chan *RegisterFaceResponse, <-chan error) {
	responseChan := make(chan *RegisterFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterFace(request)
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

// RegisterFaceWithCallback invokes the ivision.RegisterFace API asynchronously
// api document: https://help.aliyun.com/api/ivision/registerface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegisterFaceWithCallback(request *RegisterFaceRequest, callback func(response *RegisterFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterFaceResponse
		var err error
		defer close(result)
		response, err = client.RegisterFace(request)
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

// RegisterFaceRequest is the request struct for api RegisterFace
type RegisterFaceRequest struct {
	*requests.RpcRequest
	Content  string           `position:"Query" name:"Content"`
	DataType string           `position:"Query" name:"DataType"`
	ShowLog  string           `position:"Query" name:"ShowLog"`
	GroupId  string           `position:"Query" name:"GroupId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
}

// RegisterFaceResponse is the response struct for api RegisterFace
type RegisterFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	GroupId   string `json:"GroupId" xml:"GroupId"`
	Faces     []Face `json:"Faces" xml:"Faces"`
}

// CreateRegisterFaceRequest creates a request to invoke RegisterFace API
func CreateRegisterFaceRequest() (request *RegisterFaceRequest) {
	request = &RegisterFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "RegisterFace", "ivision", "openAPI")
	return
}

// CreateRegisterFaceResponse creates a response to parse from RegisterFace response
func CreateRegisterFaceResponse() (response *RegisterFaceResponse) {
	response = &RegisterFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}