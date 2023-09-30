package user

import (
	"testing"

	"github.com/rlawnsxo131/madre-server/domain/errz"
)

// @TODO 테스트 재작성

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
				return errz.IsErrMissingRequiredValue(err)
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
}
