package client

import (
	model "github.com/eric-chao/go-micro/gin-demo/proto"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestProtobuf(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/protobuf")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	user := &model.User{}
	proto.UnmarshalMerge(body, user)
	t.Log(*user)
}
