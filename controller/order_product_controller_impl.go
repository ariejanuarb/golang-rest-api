package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"go-rest-api/service"
	"net/http"
	"strconv"
)

type OrderProductControllerImpl struct {
	OrderProductService service.OrderProductService
}

func NewOrderProductController(orderProductService service.OrderProductService) OrderProductController {
	return &OrderProductControllerImpl{
		OrderProductService: orderProductService,
	}
}

func (controller *OrderProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductCreateRequest := web.OrderProductCreateRequest{}
	helper.ReadFromRequestBody(request, &orderProductCreateRequest)

	orderProductResponse := controller.OrderProductService.Create(request.Context(), orderProductCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductUpdateRequest := web.OrderProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &orderProductUpdateRequest)

	orderProductId := params.ByName("orderProductId")
	id, err := strconv.Atoi(orderProductId)
	helper.PanicIfError(err)

	orderProductUpdateRequest.Id = id

	orderProductResponse := controller.OrderProductService.Update(request.Context(), orderProductUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductId := params.ByName("orderProductId")
	id, err := strconv.Atoi(orderProductId)
	helper.PanicIfError(err)

	controller.OrderProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductId := params.ByName("orderProductId")
	id, err := strconv.Atoi(orderProductId)
	helper.PanicIfError(err)

	orderProductResponse := controller.OrderProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductResponses := controller.OrderProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderProductResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
