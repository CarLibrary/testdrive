package testdrive

import (
	"CarLibrary/testdrive/model"
	"CarLibrary/testdrive/serializer"
	"context"
	pb "github.com/CarLibrary/proto/testdrive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestDriveServerServer struct {
	pb.UnsafeTestDriveServerServer
}

//试驾
func (s *TestDriveServerServer)TestDrive(ctx context.Context, t *pb.TestDriveRequest) (*pb.TestDriveResponse, error){
	var test=&model.TestDrive{
		Userid:      uint(t.GetUserid()),
		Carband:     t.GetCarband(),
		CarSeries:   t.GetCarSeries(),
		City:        t.GetCity(),
		Username:    t.GetUsername(),
		PhoneNumber: t.GetPhoneNumber(),
	}
	res,err:=test.TestDrive()
	if err!=nil{
		return new(pb.TestDriveResponse),status.Error(codes.Aborted,err.Error())
	}
	return serializer.BuildTestDriveResponse(res),status.Error(codes.OK,"ok")
}
//查看我的试驾
func (s *TestDriveServerServer)FindMyTestDrive(m *pb.MyTestDriveRequest, t pb.TestDriveServer_FindMyTestDriveServer) error{

	res,err:=model.FindMytest(uint(m.GetUserid()))
	if err != nil {
		return err
	}
	for _,v:=range *res{
		v2:=serializer.BuildTestDriveResponse(&v)
		if err:=t.Send(v2);err!=nil{
			//todo 做log
			continue
		}
	}
	return status.Errorf(codes.OK,"ok")
}