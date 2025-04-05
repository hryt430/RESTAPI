package entity

type User struct {
	ID               UserUUID         `json:"id,omitempty"`
	Email            UserEmail        `json:"email,omitempty"`
	Password         UserPassWord     `json:"password,omitempty"`
	User_ID          UserID           `json:"user_id,omitempty"`
	FirstName        UserFirstName    `json:"first_name,omitempty"`
	LastName         UserLastName     `json:"last_name,omitempty"`
	Gender           UserGender       `json:"gender,omitempty"`
	BirthDay         UserBirthDay     `json:"birth_day,omitempty"`
	PhoneNumber      UserPhoneNumber  `json:"phone_number,omitempty"`
	PostOfficeNumber PostOfficeNumber `json:"post_office_number,omitempty"`
	Pref             Pref             `json:"pref,omitempty"`
	City             City             `json:"city,omitempty"`
	Extra            Extra            `json:"extra,omitempty"`
}

type UserUUID struct{ value string }
type UserID struct{ value string }
type UserEmail struct{ value string }
type UserPassWord struct{ value string }
type UserFirstName struct{ value string }
type UserLastName struct{ value string }
type UserPhoneNumber struct{ value string }
type UserGender struct{ value string }
type UserBirthDay struct{ value string }
type PostOfficeNumber struct{ value string }
type Pref struct{ value string }
type City struct{ value string }
type Extra struct{ value string }

// func (v *UserID) Valid() error {
// 	r := regexp.MustCompile(USER_ID_PATTERN)
// 	matched := r.MatchString(v.value)

// 	if matched {
// 		return INVALID_USER_ID
// 	}

// 	return nil
// }
