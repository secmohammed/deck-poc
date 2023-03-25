package tests

import (
    "bytes"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "net/http"
    "net/http/httptest"
)

func MakeRequest(method, url string, body interface{}, router *gin.Engine) *httptest.ResponseRecorder {

    requestBody, _ := json.Marshal(body)
    request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
    request.Header.Add("Content-Type", "application/json")

    writer := httptest.NewRecorder()
    router.ServeHTTP(writer, request)
    return writer
}
