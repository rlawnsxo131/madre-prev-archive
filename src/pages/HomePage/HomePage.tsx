import HomeSection from '../../components/home/HomeSection';
import HomeSectionGraph from '../../components/home/HomeSection/HomeSectionGraph';
import HomeSectionThinkAbout from '../../components/home/HomeSection/HomeSectionThinkAbout';

interface HomePageProps {}

function HomePage(props: HomePageProps) {
  return (
    <HomeSection>
      <HomeSectionThinkAbout />
      <HomeSectionGraph />
    </HomeSection>
  );
}

export default HomePage;
