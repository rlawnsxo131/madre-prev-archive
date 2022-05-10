import { useMemo } from 'react';
import { useSelector } from 'react-redux';
import { toast, ToastPosition } from 'react-toastify';
import { RootState } from '../store';
import { palette, themePalette } from '../styles';

interface PromisifyParams {
  func: () => Promise<void>;
  message: {
    pending: string;
    success: string;
    error: string;
  };
}

const bodyStyle = {
  color: themePalette.text1,
  fontSize: '0.9rem',
};

export default function useToast() {
  const theme = useSelector((state: RootState) => state.theme.theme);

  return useMemo(
    () => ({
      success(text: string, position: ToastPosition = 'top-right') {
        toast.success(text, {
          bodyStyle,
          theme,
          position,
        });
      },
      error(text: string, position: ToastPosition = 'top-right') {
        toast.error(text, {
          bodyStyle,
          theme,
          position,
        });
      },
      warn(text: string, position: ToastPosition = 'top-right') {
        toast.warn(text, {
          bodyStyle,
          theme,
          position,
        });
      },
      info(text: string, position: ToastPosition = 'top-right') {
        toast.info(text, {
          icon: false,
          bodyStyle,
          theme,
          progressStyle: {
            color: palette.blue['500'],
          },
          position,
        });
      },
      promisify(
        params: PromisifyParams,
        position: ToastPosition = 'top-right',
      ) {
        toast.promise(params.func, params.message, {
          theme,
          position,
        });
      },
    }),
    [theme],
  );
}
