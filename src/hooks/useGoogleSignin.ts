import { useCallback } from 'react';
import { useNavigate } from 'react-router-dom';

export default function useGoogleSignin() {
  const navigate = useNavigate();
  const signin = useCallback(() => {}, [navigate]);

  return signin;
}
