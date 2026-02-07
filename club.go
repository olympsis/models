package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Club struct {
	ID          bson.ObjectID   `json:"id" bson:"_id"`
	Parent      *OrganizationDao     `json:"parent,omitempty" bson:"parent,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Tags        []string             `json:"tags" bson:"tags"`
	Sports      []string             `json:"sports" bson:"sports"`
	City        string               `json:"city" bson:"city"`
	State       string               `json:"state" bson:"state"`
	Country     string               `json:"country" bson:"country"`
	Location    GeoJSON              `json:"location" bson:"location"`
	Logo        string               `json:"logo,omitempty" bson:"logo,omitempty"`
	Banner      string               `json:"banner,omitempty" bson:"banner,omitempty"`
	Visibility  string               `json:"visibility" bson:"visibility"`
	Members     []Member             `json:"members" bson:"members"`
	PinnedPosts []bson.ObjectID `json:"pinned_posts" bson:"pinned_posts"`
	BlackList   []string             `json:"black_list" bson:"black_list"`
	Rules       []string             `json:"rules" bson:"rules"`
	SnapshotURL *string              `json:"snapshot_url,omitempty" bson:"snapshot_url,omitempty"`
	IsVerified  bool                 `json:"is_verified,omitempty" bson:"is_verified,omitempty"`
	CreatedAt   bson.DateTime   `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type ClubDao struct {
	ID          *bson.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	ParentID    *bson.ObjectID   `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
	Name        *string               `json:"name,omitempty" bson:"name,omitempty"`
	Description *string               `json:"description,omitempty" bson:"description,omitempty"`
	Tags        *[]string             `json:"tags,omitempty" bson:"tags,omitempty"`
	Sports      *[]string             `json:"sports,omitempty" bson:"sports,omitempty"`
	City        *string               `json:"city,omitempty" bson:"city,omitempty"`
	State       *string               `json:"state,omitempty" bson:"state,omitempty"`
	Country     *string               `json:"country,omitempty" bson:"country,omitempty"`
	Location    *GeoJSON              `json:"location,omitempty" bson:"location,omitempty"`
	Logo        *string               `json:"logo,omitempty" bson:"logo,omitempty"`
	Banner      *string               `json:"banner,omitempty" bson:"banner,omitempty"`
	Visibility  *string               `json:"visibility,omitempty" bson:"visibility,omitempty"`
	PinnedPosts *[]bson.ObjectID `json:"pinned_posts,omitempty" bson:"pinned_posts,omitempty"`
	BlackList   *[]string             `json:"black_list,omitempty" bson:"black_list,omitempty"`
	Rules       *[]string             `json:"rules,omitempty" bson:"rules,omitempty"`
	SnapshotURL *string               `json:"snapshot_url,omitempty" bson:"snapshot_url,omitempty"`
	IsVerified  *bool                 `json:"is_verified,omitempty" bson:"is_verified,omitempty"`
	CreatedAt   *bson.DateTime   `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// Wrapper around the response for club objects
type ClubsResponse struct {
	TotalClubs int    `json:"total_clubs"`
	Clubs      []Club `json:"clubs"`
}

// Object to handle club invitations
type ClubInvite struct {
	ID        bson.ObjectID `json:"id"`
	Club      *Club              `json:"club,omitempty"`
	Sender    *UserData          `json:"sender,omitempty"`
	Status    string             `json:"status"`
	CreatedAt bson.DateTime `json:"created_at"`
}

// Data access object for club invites
type ClubInviteDao struct {
	ID        *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Recipient *string             `json:"recipient,omitempty" bson:"recipient,omitempty"`
	Sender    *string             `json:"sender,omitempty" bson:"sender,omitempty"`
	ClubID    *string             `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// Wrapper around the response for the club invite objects
type ClubInvitesResponse struct {
	TotalInvites int          `json:"total_invites"`
	Invites      []ClubInvite `json:"invites"`
}

// Object to handle club applications
type ClubApplication struct {
	ID        bson.ObjectID `json:"id" bson:"_id"`
	Applicant *UserData          `json:"applicant,omitempty"`
	Status    string             `json:"status"`
	CreatedAt bson.DateTime `json:"created_at" bson:"created_at"`
}

// Data access object for club applications
type ClubApplicationDao struct {
	ID        *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Applicant *string             `json:"applicant,omitempty" bson:"applicant,omitempty"`
	ClubID    *bson.ObjectID `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// Wrapper around the response for club applications
type ClubApplicationsResponse struct {
	TotalApplications int               `json:"total_applications"`
	Applications      []ClubApplication `json:"club_applications"`
}

// Object to handle role change requests
type ChangeRoleRequest struct {
	Role string `json:"role"`
}

type ClubFinancialAccountCreateRequest struct {
	Email *string `json:"string" bson:"string"`
}

type ClubFinancialAccount struct {
	ID                 *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ClubID             *bson.ObjectID `json:"club_id,omitempty" bson:"club_id,omitempty"`
	StripeAccountID    *string             `json:"stripe_account_id,omitempty" bson:"stripe_account_id,omitempty"`
	AccountStatus      *string             `json:"account_status,omitempty" bson:"account_status,omitempty"` // pending, active, restricted
	OnboardingURL      *string             `json:"onboarding_url,omitempty" bson:"onboarding_url,omitempty"`
	AvailableBalance   *float64            `json:"available_balance,omitempty" bson:"available_balance,omitempty"`
	PendingBalance     *float64            `json:"pending_balance,omitempty" bson:"pending_balance,omitempty"`
	Currency           *string             `json:"currency,omitempty" bson:"currency,omitempty"`
	RecentTransactions *[]ClubTransaction  `json:"recent_transactions,omitempty" bson:"recent_transactions,omitempty"`
	CreatedAt          *bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt          *bson.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ClubTransaction struct {
	ID             *string             `json:"id,omitempty" bson:"_id,omitempty"`
	EventID        *bson.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	Type           string              `json:"type" bson:"type"`     // payment_received, payout, fee
	Amount         int64               `json:"amount" bson:"amount"` // cents
	Currency       string              `json:"currency" bson:"currency"`
	StripeChargeID *string             `json:"stripe_charge_id,omitempty" bson:"stripe_charge_id,omitempty"`
	StripePayoutID *string             `json:"stripe_payout_id,omitempty" bson:"stripe_payout_id,omitempty"`
	Status         string              `json:"status" bson:"status"`
	CreatedAt      bson.DateTime  `json:"created_at" bson:"created_at"`
}

type PayoutRequest struct {
	Amount      int64   `json:"amount" bson:"amount"`                               // Amount in cents
	Currency    string  `json:"currency" bson:"currency"`                           // Currency code (e.g., "usd")
	Description *string `json:"description,omitempty" bson:"description,omitempty"` // Optional description
}

type StripeCustomerSheetResponse struct {
	CustomerID              string `json:"customer_id" bson:"customer_id"`
	EphemeralKeySecret      string `json:"ephemeral_key_secret" bson:"ephemeral_key_secret"`
	SetupIntentClientSecret string `json:"setup_intent_client_secret" bson:"setup_intent_client_secret"`
}
