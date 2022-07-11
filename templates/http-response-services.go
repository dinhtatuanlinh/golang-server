package template

import (
	"server/pkg/libs/alias"
	"server/pkg/utils"

	"encoding/json"
	"net/http"
)

// RawResponse is a func to print file raw response
// func RawResponse(w http.ResponseWriter, r *http.Request, payload string, responseNumber int) {
// 	w.WriteHeader(responseNumber)
// 	_, err := fmt.Fprintf(w, payload)
// 	if err != nil {
// 		errMessage := alias.GetMessageTypeError(alias.ErrMessageBadParam)
// 		HttpResponseError4xx(w, r, errMessage, alias.HTTP_BADREQUEST, int(alias.ErrMessageBadParam))
// 	}
// }

func HttpResponse(w http.ResponseWriter, r *http.Request, res []byte, responseNumber int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseNumber)
	w.Write(res)
}

func HttpResponseError4xx(w http.ResponseWriter, r *http.Request, data interface{}, responseNUmber int, codeNumber int) {
	var responseErr DataErrorMessage4xx
	// tmpCode := findCodeMessage(data.(string))
	// if tmpCode != -1 {
	// 	responseErr.Code = tmpCode
	// } else {
	// 	responseErr.Code = codeNumber
	// }
	responseErr.Code = codeNumber
	responseErr.Message = data
	responseErr.Time = utils.Now(alias.GMT, alias.TIME_FORMAT)
	resp, err := json.Marshal(responseErr)
	if err != nil {
		// utils.Logln(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseNUmber)
	w.Write(resp)
}

// func HttpBadRequestResponse(r *http.Request, w http.ResponseWriter, err error) {
// 	utils.GetErrNotifierFromContext(r.Context())(err, http.StatusBadRequest)
// 	HttpResponseError4xx(w, r, alias.GetMessageTypeError(alias.ErrMessageBadParam), http.StatusBadRequest, int(alias.ErrMessageBadParam))
// }

// func HttpBadRequestWithErrorMessageResponse(r *http.Request, w http.ResponseWriter, err error) {
// 	utils.GetErrNotifierFromContext(r.Context())(err, http.StatusBadRequest)
// 	HttpResponseError4xx(w, r, err.Error(), http.StatusBadRequest, int(alias.ErrMessageBadParam))
// }

// func HttpInternalErrorResponse(r *http.Request, w http.ResponseWriter, err error, info map[string]interface{}) {
// 	monitoring.NotifyError(err, info)
// 	utils.GetErrNotifierFromContext(r.Context())(err, http.StatusInternalServerError)
// 	HttpResponseError5xx(w, r, alias.InternalServerErrMsg, http.StatusInternalServerError)
// }

// func HttpNotFoundResponse(r *http.Request, w http.ResponseWriter, err error, typeErrorMessage alias.TypeErrorMessage) {
// 	utils.GetErrNotifierFromContext(r.Context())(err, http.StatusNotFound)
// 	HttpResponseError4xx(w, r, alias.GetMessageTypeError(typeErrorMessage), http.StatusNotFound, int(typeErrorMessage))
// }

// func HttpResponseWithCustomError(w http.ResponseWriter, r *http.Request, err *e.RequestError) {
// 	utils.GetErrNotifierFromContext(r.Context())(err.Err, err.StatusCode)

// 	if err.StatusCode == http.StatusInternalServerError {
// 		HttpInternalErrorResponse(r, w, err.Err, nil)
// 		return
// 	}

// 	if err.Data != nil {
// 		HttpResponseError4xxWithCustomData(r, w, err.Data, err.StatusCode, err.Code)
// 		return
// 	}

// 	HttpResponseError4xx(w, r, err.Error(), err.StatusCode, err.Code)
// }

// func HttpResponseError4xxWithCustomData(r *http.Request, w http.ResponseWriter, data interface{}, statusCode int, code int) {
// 	var responseErr DataErrorMessage4xxWithCustomerData

// 	responseErr.Data = data
// 	responseErr.Code = code
// 	responseErr.Time = utils.Now(alias.GMT, alias.TIME_FORMAT)
// 	resp, err := json.Marshal(responseErr)

// 	if err != nil {
// 		utils.Logln(err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	w.Write(resp)
// }

// // HttpResponseError4xxOr5xx is a function to return Error 4xx
// // if the parameter data(string) is in alias.ErrMessageArray
// // or return Error 5xx InternalServerError
// func HttpResponseError4xxOr5xx(w http.ResponseWriter, r *http.Request, data string, responseNumber int) {
// 	var responseErr DataErrorMessage4xx
// 	tmpCode := findCodeMessage(data)
// 	if tmpCode == -1 {
// 		HttpResponseError5xx(w, r, data, alias.HTTP_INTERNALSERVERERROR)
// 		return
// 	}
// 	responseErr.Code = tmpCode

// 	responseErr.Message = data
// 	responseErr.Time = utils.Now(alias.GMT, alias.TIME_FORMAT)
// 	resp, err := json.Marshal(responseErr)
// 	if err != nil {
// 		utils.Logln(err)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(responseNumber)
// 	w.Write(resp)
// 	return
// }

// func HttpResponseError5xx(w http.ResponseWriter, r *http.Request, data interface{}, responseNumber int) {
// 	var responseErr DataErrorMessage5xx
// 	responseErr.Message = data
// 	resp, err := json.Marshal(responseErr)
// 	if err != nil {
// 		utils.Logln(err)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(responseNumber)
// 	w.Write(resp)
// }

// // findCodeMessage is a function  to Find error Message in alias.ErrTypeMessageMap ErrType(int)
// // and it will return TypeErrorMessage (int) if found
// // if not found the function will return -1
// func findCodeMessage(message string) int {
// 	return alias.GetTypeID(message)
// }
