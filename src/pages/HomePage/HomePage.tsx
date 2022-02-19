import HomeSection from '../../components/home/HomeSection';
import HomeFooter from '../../components/home/HomeFooter';
import { useGetAuthCheckGoogleQuery } from '../../redux/api/authApi';
import { Suspense } from 'react';

interface HomePageProps {}

function HomePage(props: HomePageProps) {
  const { isLoading, isError, data } = useGetAuthCheckGoogleQuery({});

  console.log(isLoading);
  console.log(isError);
  console.log(data);

  return (
    <Suspense fallback={<div>loading</div>}>
      <HomeSection>
        <HomeSection.ThinkAbout />
        <HomeSection.Graph />
        <HomeFooter />
      </HomeSection>
    </Suspense>
  );
}

export default HomePage;
