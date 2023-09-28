package user

import (
	"testing"

	"github.com/rlawnsxo131/madre-server/domain/domainerr"
)

func Test_UserSocialAccount(t *testing.T) {
	// SetNewSocialAccount Fail
	setNewSocialAccountTests := []struct {
		name string
		args struct {
			user     *User
			socialId string
			provider string
		}
		want func(err error) bool
	}{
		{
			name: "userId 가 없을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				user     *User
				socialId string
				provider string
			}{
				user:     &User{},
				socialId: "socialId",
				provider: "provider",
			},
			want: func(err error) bool {
				return domainerr.IsErrMissingRequiredValue(err)
			},
		},
		{
			name: "socialId 가 없을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				user     *User
				socialId string
				provider string
			}{
				user:     &User{Id: 1},
				socialId: "",
				provider: "provider",
			},
			want: func(err error) bool {
				return domainerr.IsErrMissingRequiredValue(err)
			},
		},
		{
			name: "provider 가 없을때 errMissingRequiredValue 에러를 리턴한다",
			args: struct {
				user     *User
				socialId string
				provider string
			}{
				user:     &User{Id: 1},
				socialId: "socialId",
				provider: "",
			},
			want: func(err error) bool {
				return domainerr.IsErrMissingRequiredValue(err)
			},
		},
		{
			name: "provider 가 GOOGLE 이 아닐때 errNotSupportValue 에러를 리턴한다",
			args: struct {
				user     *User
				socialId string
				provider string
			}{
				user:     &User{Id: 1},
				socialId: "socialId",
				provider: "provider",
			},
			want: func(err error) bool {
				return domainerr.IsErrNotSupportValue(err)
			},
		},
	}

	for _, tt := range setNewSocialAccountTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.user.SetNewSocialAccount(
				tt.args.socialId,
				tt.args.provider,
			)
			if want := tt.want(got); !want {
				t.Errorf("SetNewSocialAccount() = %v, want = %v", got, want)
			}
		})
	}

	// SetNewSocialAccount Success
	t.Run("socialAccount 를 올바른 값으로 만들고, 에러는 nil 을 리턴한다", func(t *testing.T) {
		u := &User{Id: 1}
		socialId := "socialId"
		provider := "GOOGLE"
		err := u.SetNewSocialAccount("socialId", "GOOGLE")

		if err != nil {
			t.Errorf("SetNewSocialAccount() = %+v, want = %v", u, err)
		}
		if u.SocialAccount.SocialId != socialId {
			t.Errorf("SetNewSocialAccount() = %+v, want = %v", u, u.PhotoUrl)
		}
		if u.SocialAccount.Provider != provider {
			t.Errorf("SetNewSocialAccount() = %+v, want = %v", u, u.PhotoUrl)
		}
	})
}
