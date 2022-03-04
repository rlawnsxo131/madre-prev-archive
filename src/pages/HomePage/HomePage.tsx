import HomeSection from '../../components/home/HomeSection';
import HomeFooter from '../../components/home/HomeFooter';

interface HomePageProps {}

function HomePage(props: HomePageProps) {
  return (
    <HomeSection>
      <HomeSection.ThinkAbout />
      <HomeSection.Graph />
      <HomeFooter />
    </HomeSection>
  );
}

export default HomePage;
