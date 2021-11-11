import 'bootstrap/dist/css/bootstrap.css';
import buildClient from '../api/build-client';
import Header from '../components/header';

const AppComponent =  ({Component, pageProps, currentUser}) => {
    return <div>
        <Header currentUser={currentUser}/>
        <Component {...pageProps} />
    </div>
};

AppComponent.getInitialProps = async (appContext) => {
    const client = buildClient(appContext.ctx);
    let data = {}
    await client.get('/api/users/currentuser')
    .then((response) => {
        data = response.data;
      })
    .catch((err) => {
         console.log(err.message);
     });
     console.log(data)
     let currentUser = data.currentUser;

     let pageProps= {};
     if (appContext.Component.getInitialProps) {
        pageProps = await appContext.Component.getInitialProps(appContext.ctx);
     }
     return {
         pageProps,
         currentUser
     };
};

export default AppComponent;