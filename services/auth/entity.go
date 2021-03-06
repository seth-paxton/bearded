package auth

type authEntity struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type sessionEntity struct {
	Token string `json:"token" description:"isn't implemented yet"`
}

type passwordEntity struct {
	Password string `json:"password"`
}

type registerEntity struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:",required"`
}

type resetPasswordEntity struct {
	Email string `json:"email" valid:"email,required"`
}
