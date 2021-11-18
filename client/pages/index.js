import buildClient from "../api/build-client";

const LandingPage = (currentUser) => {
  return <h1>Welcome</h1>
};

LandingPage.getInitialProps = async (context) => {
  let data = {};
  await buildClient(context).get('/api/users/currentuser')
  .then((response) => {
    data = response.data;
  })
  .catch((err) => {
    console.log(err.message);
  })
  console.log(data)
  let currentUser = data.currentUser;
  console.log(currentUser)
  return currentUser
};

export default LandingPage;