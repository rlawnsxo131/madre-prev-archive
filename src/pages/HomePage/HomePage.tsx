import HomeSection from '../../components/home/HomeSection';
import HomeFooter from '../../components/home/HomeFooter';
import { useEffect } from 'react';
import apiClient from '../../api/apiClient';

interface HomePageProps {}

function HomePage(props: HomePageProps) {
  useEffect(() => {
    apiClient.get('/temp/get').then((res) => {
      console.log(res);
    });
  }, []);
  return (
    <HomeSection>
      <HomeSection.ThinkAbout />
      <HomeSection.Graph />
      <HomeFooter />
    </HomeSection>
  );
}

export default HomePage;
