package repositories

var _userRepository *Repository

func UserRepository() *Repository {
	if _userRepository != nil {
		return _userRepository
	}

	_userRepository = NewRepository(UsersCollectionName)

	return _userRepository
}
