package models

/*
Auth User
  - User data for auth database
  - Contains user identifiable data
*/
type AuthUser struct {
	UUID      *string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	FirstName *string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email     *string `json:"email,omitempty" bson:"email,omitempty"`
	Provider  *string `json:"provider,omitempty" bson:"provider,omitempty"`
	CreatedAt *int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

/*
Sign In Request
  - Request coming from client
  - Contains user identifiable data
*/
type AuthRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	UUID      *string `json:"uuid,omitempty"`
}

/*
Log in Response
  - User identifiable data for client
*/
type AuthResponse struct {
	UUID      *string `json:"uuid,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
}
