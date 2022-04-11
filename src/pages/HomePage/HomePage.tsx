import HomeSection from '../../components/home/HomeSection';

interface HomePageProps {}

function HomePage(props: HomePageProps) {
  return (
    <HomeSection>
      <HomeSection.ThinkAbout />
      <HomeSection.Graph />
    </HomeSection>
  );
}

export default HomePage;
