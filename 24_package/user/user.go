package user

// we must need to make struct name's first char capital
// to have access of it outside this auth package
//same goes for sub feilds in struct
// like email ,password as below
// if we don't want expose it to pacakge outside than we name it with non capital first later

type User struct {
	Email string
	Name  string
}
