import { keyframes } from '@emotion/react';

const fadeIn = keyframes`
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
`;

const fadeOut = keyframes`
  0% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
`;

const popInFromBottom = keyframes`
  0% {
    opacity: 0;
    transform: translateY(400px) scale(0.75);
  }
  75% {
    opacity: 1;
    transform: translateY(-16px) scale(1.0);
  }
  100% {
    opacity: 1;
    transform: translateY(0px);
  }`;

const popOutToBottom = keyframes`
  0% {
    opacity: 1;
    transform: translateY(0px) scale(1.0);
  }
  100% {
    opacity: 0;
    transform: translateY(400px) scale(0.75);
  }`;

const scaleUp = keyframes`
  0% {
    opacity: 0;
    transform: scaleY(0);
  }
  100% {
    opacity: 1;
    transform: scaleY(1);
  }
`;

const scaleDown = keyframes`
  0% {
    opacity: 1;
    transform: scaleY(1);
  }
  100% {
    opacity: 0;
    transform: scaleY(0);;
  }
`;

const slideUp = keyframes`
  0% {
    transform: translateY(100%);
  }
  100% {
    transform: translateY(0%);
  };
`;

const slideDown = keyframes`
  0% {
    transform: translateY(0%);
  }
  100% {
    transform: translateY(100%);
  };
`;

const shake = keyframes`
  0% {
      transform: translate(-15px);
  }
  25% {
      transform: translate(15px);
  }
  50% {
      transform: translate(-10px);
  }
  75% {
      transform: translate(5px);
  }
  100% {
      transform: translate(0px);
  }
`;

const rotation = keyframes`
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
`;

const transitions = {
  fadeIn,
  fadeOut,
  popInFromBottom,
  popOutToBottom,
  scaleUp,
  scaleDown,
  slideUp,
  slideDown,
  shake,
  rotation,
};

export default transitions;
