import HomeSection from '../../components/home/HomeSection';
import HomeFooter from '../../components/home/HomeFooter';
import { useGetAuthCheckGoogleQuery } from '../../redux/api/authApi';

interface HomePageProps {}

function HomePage(props: HomePageProps) {
  const { isLoading, isError, data } = useGetAuthCheckGoogleQuery({});
  console.log(isLoading);
  console.log(isError);
  console.log(data);

  if (isLoading) return <div>loading</div>;
  console.log(data);
  return (
    <HomeSection>
      <HomeSection.ThinkAbout />
      <HomeSection.Graph />
      <HomeFooter />
    </HomeSection>
  );
}

export default HomePage;
