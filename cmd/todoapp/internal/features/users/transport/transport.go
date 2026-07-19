package users_transport_http

type UsersHTTPHandler struct {
	usersService UsersService /*описывыает слой транспорта для юзера те обрабывает запросы и передает их в сервисный слой.под обработкой имеется ввиду принимание запросов и их декодирование*/
}

type UsersService interface { // юзер сервис это зависимость для ххтп хендлера
}

func NewUsersHTTPHandler( 
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
} 

