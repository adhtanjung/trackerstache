package services

import (
	"context"
	"net/http"

	"github.com/adhtanjung/auth-svc/pkg/db"
	"github.com/adhtanjung/auth-svc/pkg/models"
	"github.com/adhtanjung/auth-svc/pkg/pb"
	"github.com/adhtanjung/auth-svc/pkg/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Username: req.Username}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Username already exists",
		}, nil
	}

	user.Username = req.Username
	user.Password = utils.HashPassword(req.Password)

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) GetAll(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllResponse, error) {

	var users []models.User

	if result := s.H.DB.Find(&users); result.Error != nil {
		return &pb.GetAllResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	var data []*pb.User
	for _, user := range users {
		createdAt := user.CreatedAt.Format("2006-01-02 15:04:05")
		updatedAt := user.UpdatedAt.Format("2006-01-02 15:04:05")
		data = append(data, &pb.User{
			Id:        user.ID,
			Username:  user.Username,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return &pb.GetAllResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil

}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Username: req.Username}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{Username: claims.Username}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.ID,
	}, nil
}

func (s *Server) CreateRoleForUser(ctx context.Context, req *pb.CreateRoleForUserRequest) (*pb.RoleResponse, error) {
	// var role models.Role

	db := s.H.DB
	baseModel := models.BaseModel{ID: req.RoleId}

	if err := db.First(&models.Role{BaseModel: baseModel}).Error; err != nil {
		return &pb.RoleResponse{
			Status: http.StatusNotFound,
			Error:  "Role not found",
		}, nil
	}

	return &pb.RoleResponse{
		Status: http.StatusOK,
	}, nil

}

func (s *Server) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.RoleResponse, error) {
	var role models.Role

	db := s.H.DB
	role.Name = req.Name

	db.Create(&role)

	return &pb.RoleResponse{
		Status: http.StatusOK,
	}, nil
}
