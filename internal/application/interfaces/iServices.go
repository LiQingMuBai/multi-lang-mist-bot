package interfaces

import (
	"ushield_bot/internal/domain"
)

type IUserService interface {
	Create(user domain.User) error
	Update(user domain.User) error
	UpdateAddress(user domain.User) error
	UpdateTimes(_times uint64, _username string) error
	GetByUsername(username string) (domain.User, error)
	GetByUserID(_userID int64) (domain.User, error)
	FetchNewestAddress() ([]domain.User, error)
	NotifyTronAddress() ([]domain.User, error)
	NotifyEthereumAddress() ([]domain.User, error)

	BindEthereumAddress(_address string, _username string) error
	BindTronAddress(_address string, _username string) error
	BindChat(_associates string, _username string) error
	DisableTronAddress(_address string) error
}
