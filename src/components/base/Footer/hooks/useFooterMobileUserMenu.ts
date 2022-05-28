import { useEffect } from 'react';
import { useRef } from 'react';
import { useMemo } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import usePopupAuthActions from '../../../../hooks/popupAuth/usePopupAuthActions';
import useIsUserPath from '../../../../hooks/useIsUserPath';
import useUserProfileState from '../../../../hooks/user/useUserProfileState';

export default function useFooterMobileUserMenu() {
  const isClickRef = useRef(false);
  const navigate = useNavigate();
  const { show } = usePopupAuthActions();
  const { username } = useParams();
  const isUserPath = useIsUserPath();
  const profile = useUserProfileState();

  const isActive = useMemo(() => {
    return isUserPath && profile?.username === username;
  }, [username, profile?.username, isUserPath]);

  const onClick = () => {
    if (!profile) {
      isClickRef.current = true;
      show();
      return;
    }
    navigate(`/@${profile.username}`);
  };

  useEffect(() => {
    if (!profile?.username) return;
    if (!isClickRef.current) return;

    navigate(`/@${profile?.username}`);

    return () => {
      isClickRef.current = false;
    };
  }, [profile?.username, navigate]);

  return {
    isActive,
    onClick,
  };
}
