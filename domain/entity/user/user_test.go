package user

import (
	"log"
	"testing"

	"github.com/rlawnsxo131/madre-server/domain/errz"
)

// @TODO 테스트 재작성

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
			name: "email 이 없을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				email    string
				photoUrl string
			}{
				email:    "",
				photoUrl: "photoUrl",
			},
			want: func(err error) bool {
				return errz.IsErrMissingRequiredValue(err)
			},
		},
	}

	for _, tt := range newTests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := New(tt.args.email, tt.args.photoUrl)
			log.Println(got)
			if want := tt.want(got); !want {
				t.Errorf("New() = %v, want = %v", got, want)
			}
		})
	}

	// New Success
	t.Run("user 를 올바른 값으로 만들고, 에러는 nil 을 리턴한다", func(t *testing.T) {
		email := "rlawnsxo131@gmail.com"
		photoUrl := "https://google.com"
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
}
