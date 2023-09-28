package user

import (
	"testing"

	"github.com/rlawnsxo131/madre-server/domain/domainerr"
)

func Test_User(t *testing.T) {
	// New Fail
	newTests := []struct {
		name string
		args struct {
			email    string
			photoUrl string
		}
		want func(err error) bool
	}{
		{
			name: "email 과 photoUrl 이 없을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				email    string
				photoUrl string
			}{
				email:    "",
				photoUrl: "",
			},
			want: func(err error) bool {
				return domainerr.IsErrMissingRequiredValue(err)
			},
		},
		{
			name: "email 이 없을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				email    string
				photoUrl string
			}{
				email:    "",
				photoUrl: "photoUrl",
			},
			want: func(err error) bool {
				return domainerr.IsErrMissingRequiredValue(err)
			},
		},
		{
			name: "photoUrl 이 없을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				email    string
				photoUrl string
			}{
				email:    "email",
				photoUrl: "",
			},
			want: func(err error) bool {
				return domainerr.IsErrMissingRequiredValue(err)
			},
		},
		{
			name: "email 형식이 잘못 되었을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				email    string
				photoUrl string
			}{
				email:    "email",
				photoUrl: "photo_url",
			},
			want: func(err error) bool {
				return domainerr.IsErrNotSupportValue(err)
			},
		},
	}

	for _, tt := range newTests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := New(tt.args.email, tt.args.photoUrl)
			if want := tt.want(got); !want {
				t.Errorf("New() = %v, want = %v", got, want)
			}
		})
	}

	// New Success
	t.Run("user 를 올바른 값으로 만들고, 에러는 nil 을 리턴한다", func(t *testing.T) {
		email := "rlawnsxo131@gmail.com"
		photoUrl := "photoUrl"
		u, err := New(email, photoUrl)

		if err != nil {
			t.Errorf("New() = %+v, want = %v", u, err)
		}
		if email != u.Email {
			t.Errorf("New() = %+v, want = %v", u, u.PhotoUrl)
		}
		if photoUrl != u.PhotoUrl {
			t.Errorf("New() = %+v, want = %v", u, u.PhotoUrl)
		}
	})

	// SetUsername Fail
	setUsernameTests := []struct {
		name string
		args struct {
			user     *User
			username string
		}
		want func(err error) bool
	}{
		{
			name: "username 이 없을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				user     *User
				username string
			}{
				user:     &User{},
				username: "",
			},
			want: func(err error) bool {
				return domainerr.IsErrMissingRequiredValue(err)
			},
		},
		{
			name: "username 형식이 잘못되었을때 errNotSupportValue 를 리턴한다",
			args: struct {
				user     *User
				username string
			}{
				user:     &User{},
				username: "$@@!",
			},
			want: func(err error) bool {
				return domainerr.IsErrNotSupportValue(err)
			},
		},
	}

	for _, tt := range setUsernameTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.user.SetUsername(tt.args.username)
			if want := tt.want(got); !want {
				t.Errorf("New() = %v, want = %v", got, want)
			}
		})
	}

	// SetUsername Success
	t.Run("username 을 올바른 값으로 만들고, 에러는 nil 을 리턴한다", func(t *testing.T) {
		u := &User{}
		username := "username"
		err := u.SetUsername(username)

		if err != nil {
			t.Errorf("New() = %+v, want = %v", u, err)
		}
		if username != u.Username {
			t.Errorf("New() = %+v, want = %v", u, u.Username)
		}
	})
}
