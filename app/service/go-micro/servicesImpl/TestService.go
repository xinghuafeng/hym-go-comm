package servicesImpl

import (
	"context"
	"hym-go-comm/app/service/go-micro/services/proto"
	"strconv"
)

type TestService struct {

}
 func(this *TestService) Call(c context.Context, req *proto.TestRequerst, rsp *proto.TestRespone) error{
  rsp.Data="test"+strconv.Itoa(int(req.Id))
	 return nil
 }
