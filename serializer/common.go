package serializer

import (
	"CarLibrary/testdrive/model"
	pb "github.com/CarLibrary/proto/testdrive"
)

func BuildTestDriveResponse(t *model.TestDrive) *pb.TestDriveResponse {
	return &pb.TestDriveResponse{
		Carband:   t.Carband,
		CarSeries: t.CarSeries,
		City:      t.City,
		Status:    t.Status,
	}
}