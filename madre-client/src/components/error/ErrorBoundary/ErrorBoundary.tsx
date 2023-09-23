import { Component } from 'react';

interface ErrorBoundaryProps {
  children: React.ReactNode;
}

class ErrorBoundary extends Component<ErrorBoundaryProps> {
  state = {
    hasError: false,
    chunkError: false,
  };

  static getDerivedStateFromError(error: Error) {
    // 다음 렌더링에서 폴백 UI가 보이도록 상태를 업데이트 합니다.
    if (error.name === 'ChunkLoadError') {
      return {
        chunkError: true,
      };
    }
    return { hasError: true };
  }

  componentDidCatch(error: Error, errorMadreStory: any) {
    // 에러 리포팅 서비스에 에러를 기록할 수도 있습니다.(production Sentry)
    // logErrorToMyService(error, errorMadreStory);
  }

  handleResolveError = () => {
    this.setState({
      hasError: false,
    });
  };

  render() {
    if (this.state.hasError) {
      // 폴백 UI를 커스텀하여 렌더링할 수 있습니다.
      return <h1>Something went wrong.</h1>;
    }
    return <>{this.props.children}</>;
  }
}

export default ErrorBoundary;
