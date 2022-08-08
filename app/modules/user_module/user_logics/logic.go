package user_logics

import (
    "errors"
    "go-simple/app/modules/user_module/user"
    "go-simple/app/requests"
)

type UserLogic struct {

}

func (logic *UserLogic) IndexLogic () (userModels []user.User , err error) {
	userModels = user.All()
	return userModels, err
}

func (logic *UserLogic) ShowLogic (ID string) (userModel user.User, err error) {
	userModel = user.Get(ID)
	if userModel.ID == 0 {
		return userModel, errors.New("not found")
	}
	return
}

func (logic *UserLogic) StoreLogic (request requests.UserRequest) (userModel user.User, err error) {
	userModel = user.User{
		Name:      request.Name,
		Username:      request.Username,
	}
	userModel.Create()
	if userModel.ID > 0 {
		return userModel, nil
	}
	return user.User{}, errors.New("create fail")
}

func (logic *UserLogic) UpdateLogic (ID string, request requests.UserRequest) (userModel user.User, err error) {
	userModels := user.Get(ID)
	if userModels.ID == 0 {
	    return userModel, errors.New("not found")
	}
	userModel.Name = request.Name
    rowsAffected := userModel.Save()
    if rowsAffected > 0 {
        return userModel, nil
    } else {
        return user.User{}, errors.New("update fail")
    }
}

func (logic *UserLogic) DeleteLogic (ID string) (rowsAffected int64, err error) {
	userModel := user.Get(ID)
	if userModel.ID == 0 {
		return 0, errors.New("not found")
	}
    rowsAffected = userModel.Delete()
    if rowsAffected > 0 {
        return rowsAffected, nil
    }
    return 0, errors.New("deleted fail")
}



