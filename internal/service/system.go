package service

import (
	"io"
	"os"

	pb "github.com/HCH1212/taxin/api/pb/system"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type SystemService struct {
	pb.UnimplementedSystemServiceServer
}

// SendFile 读取一个本地文件以流的形式返回。
func (s *SystemService) SendFile(req *pb.SendFileReq, stream pb.SystemService_SendFileServer) error {
	tracer := otel.Tracer("system-service")
	_, span := tracer.Start(stream.Context(), "SendFile")
	defer span.End()

	// 读取文件
	file, err := os.Open(req.FilePath)
	if err != nil {
		span.SetStatus(codes.Error, "failed to open file")
		return err
	}
	defer file.Close()

	// 读取文件内容并发送给客户端
	buf := make([]byte, 1024*32) // 32KB分块
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			span.SetStatus(codes.Error, "failed to read file")
			return err
		}

		if err := stream.Send(&pb.SendFileResp{
			Content: buf[:n],
		}); err != nil {
			span.SetStatus(codes.Error, "failed to send file")
			return err
		}
	}
	return nil
}
