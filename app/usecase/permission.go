package usecase

import "automatic-trade/app"

type PermissionUseCase struct {
	permissionRequest IPermissionRequest
}

type IPermissionRequest interface {
	GetPermission() (*[]app.Permission, error)
}

func NewPermissionUseCase(permissionRequest IPermissionRequest) PermissionUseCase {
	return PermissionUseCase{permissionRequest: permissionRequest}
}

func (useCase PermissionUseCase) GetPermission() (*[]app.Permission, error) {
	permissions, err := useCase.permissionRequest.GetPermission()
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
