package service

import (
	"context"
	"math/rand"
	pb "verifyCodeSystem/api/verifyCodeSystem"
)

type VerifyCodeSystemService struct {
	pb.UnimplementedVerifyCodeSystemServer
}

func NewVerifyCodeSystemService() *VerifyCodeSystemService {
	return &VerifyCodeSystemService{}
}

func (s *VerifyCodeSystemService) GetVerifyCodeSystem(ctx context.Context, req *pb.GetVerifyCodeSystemRequest) (*pb.GetVerifyCodeSystemReply, error) {
	return &pb.GetVerifyCodeSystemReply{
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}

// RandCode 开放的被调用的方法，用于区分类型
func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		return randCode("0123456789", l)
	case pb.TYPE_LETTER:
		return randCode("abcdefghijklmnopqrstuvwxyz", l)
	case pb.TYPE_MIXED:
		return randCode("0123456789abcdefghijklmnopqrstuvwxyz", l)
	default:

	}
	return ""
}

// 随机的核心方法(最简单的实现方法)
func randCode(chars string, l int) string {
	charsLen := len(chars)
	
	result := make([]byte, l)

	for i := 0; i < l; i++ {
		randIndex := rand.Intn(charsLen)
		result[i] = chars[randIndex]
	}

	return string(result)
}
