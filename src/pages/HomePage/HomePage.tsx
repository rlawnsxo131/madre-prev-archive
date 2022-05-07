import Button from '../../components/common/Button';
import usePopupCommonActions from '../../hooks/common/usePopupCommonActions';

interface HomePageProps {}

// disabled 랑 outline 스타일 theme 에 맞게 수정하기
function HomePage(props: HomePageProps) {
  const { show } = usePopupCommonActions();
  return (
    <div>
      <button onClick={() => show({ message: 'asdfasdf' })}>popup</button>
      <Button outline>outline1</Button>
      <Button color="blue" outline>
        outline2
      </Button>
      <Button>default1</Button>
      <Button color="blue">default2</Button>
      home
    </div>
  );
}

export default HomePage;
