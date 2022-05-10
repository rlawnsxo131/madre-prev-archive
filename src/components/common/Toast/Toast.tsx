import { ToastContainer, Flip } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

interface ToastProps {}

function Toast(props: ToastProps) {
  return (
    <ToastContainer
      autoClose={2000}
      hideProgressBar={false}
      newestOnTop={true}
      transition={Flip}
      closeOnClick
      pauseOnHover
    />
  );
}

export default Toast;
