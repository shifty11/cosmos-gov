// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-gov/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type user.Type `json:"type,omitempty"`
	// LoginToken holds the value of the "login_token" field.
	LoginToken string `json:"login_token,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// TelegramChats holds the value of the telegram_chats edge.
	TelegramChats []*TelegramChat `json:"telegram_chats,omitempty"`
	// DiscordChannels holds the value of the discord_channels edge.
	DiscordChannels []*DiscordChannel `json:"discord_channels,omitempty"`
	// Wallets holds the value of the wallets edge.
	Wallets []*Wallet `json:"wallets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TelegramChatsOrErr returns the TelegramChats value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TelegramChatsOrErr() ([]*TelegramChat, error) {
	if e.loadedTypes[0] {
		return e.TelegramChats, nil
	}
	return nil, &NotLoadedError{edge: "telegram_chats"}
}

// DiscordChannelsOrErr returns the DiscordChannels value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DiscordChannelsOrErr() ([]*DiscordChannel, error) {
	if e.loadedTypes[1] {
		return e.DiscordChannels, nil
	}
	return nil, &NotLoadedError{edge: "discord_channels"}
}

// WalletsOrErr returns the Wallets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WalletsOrErr() ([]*Wallet, error) {
	if e.loadedTypes[2] {
		return e.Wallets, nil
	}
	return nil, &NotLoadedError{edge: "wallets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldUserID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldType, user.FieldLoginToken:
			values[i] = new(sql.NullString)
		case user.FieldCreateTime, user.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = value.Time
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = value.Time
			}
		case user.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				u.UserID = value.Int64
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				u.Type = user.Type(value.String)
			}
		case user.FieldLoginToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login_token", values[i])
			} else if value.Valid {
				u.LoginToken = value.String
			}
		}
	}
	return nil
}

// QueryTelegramChats queries the "telegram_chats" edge of the User entity.
func (u *User) QueryTelegramChats() *TelegramChatQuery {
	return (&UserClient{config: u.config}).QueryTelegramChats(u)
}

// QueryDiscordChannels queries the "discord_channels" edge of the User entity.
func (u *User) QueryDiscordChannels() *DiscordChannelQuery {
	return (&UserClient{config: u.config}).QueryDiscordChannels(u)
}

// QueryWallets queries the "wallets" edge of the User entity.
func (u *User) QueryWallets() *WalletQuery {
	return (&UserClient{config: u.config}).QueryWallets(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(u.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(u.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", u.UserID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", u.Type))
	builder.WriteString(", login_token=")
	builder.WriteString(u.LoginToken)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
