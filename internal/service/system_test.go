package service

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/HCH1212/taxin/api/pb/system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

// MockSystemServiceServer 完整实现 SystemService_SendFileServer 接口
type MockSystemServiceServer struct {
	mock.Mock
	ctx context.Context
}

func (m *MockSystemServiceServer) Send(resp *system.SendFileResp) error {
	args := m.Called(resp)
	return args.Error(0)
}

func (m *MockSystemServiceServer) Context() context.Context {
	if m.ctx != nil {
		return m.ctx
	}
	return context.Background()
}

func (m *MockSystemServiceServer) RecvMsg(msg interface{}) error {
	return io.EOF
}

func (m *MockSystemServiceServer) SendMsg(msg interface{}) error {
	return nil
}

func (m *MockSystemServiceServer) SetHeader(metadata.MD) error {
	return nil
}

func (m *MockSystemServiceServer) SendHeader(metadata.MD) error {
	return nil
}

func (m *MockSystemServiceServer) SetTrailer(metadata.MD) {
}

func TestSystemService_SendFile(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		fileContent string
		wantErr     bool
		mockExpect  func(*MockSystemServiceServer)
	}{
		{
			name:        "successful file transfer",
			filePath:    "testfile.txt",
			fileContent: "This is a test file content",
			wantErr:     false,
			mockExpect: func(m *MockSystemServiceServer) {
				// 预期会调用 Send 方法，具体次数取决于文件大小和缓冲区大小
				m.On("Send", mock.AnythingOfType("*system.SendFileResp")).Return(nil)
			},
		},
		{
			name:        "file not found",
			filePath:    "nonexistent.txt",
			fileContent: "",
			wantErr:     true,
			mockExpect:  func(m *MockSystemServiceServer) {},
		},
		{
			name:        "stream send error",
			filePath:    "testfile.txt",
			fileContent: "This is a test file content",
			wantErr:     true,
			mockExpect: func(m *MockSystemServiceServer) {
				m.On("Send", mock.AnythingOfType("*system.SendFileResp")).Return(io.ErrClosedPipe)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建测试文件
			if tt.fileContent != "" {
				err := os.WriteFile(tt.filePath, []byte(tt.fileContent), 0644)
				assert.NoError(t, err)
				defer os.Remove(tt.filePath)
			}

			// 创建模拟流
			mockStream := &MockSystemServiceServer{}
			tt.mockExpect(mockStream)

			// 创建服务实例
			service := &SystemService{}

			// 执行测试
			err := service.SendFile(&system.SendFileReq{FilePath: tt.filePath}, mockStream)

			// 验证结果
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// 验证模拟调用
			mockStream.AssertExpectations(t)
		})
	}
}
