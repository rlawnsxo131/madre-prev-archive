import HomeSection from '../../components/home/HomeSection';
import HomeFooter from '../../components/home/HomeFooter';
import { useGetAuthCheckGoogleQuery } from '../../store/api/authApi';

interface HomePageProps {}

function HomePage(props: HomePageProps) {
  const { data } = useGetAuthCheckGoogleQuery({});

  console.log('data: ', data);

  return (
    <HomeSection>
      <HomeSection.ThinkAbout />
      <HomeSection.Graph />
      <HomeFooter />
    </HomeSection>
  );
}

export default HomePage;
