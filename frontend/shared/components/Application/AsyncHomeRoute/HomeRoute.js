import React from 'react';
import Helmet from 'react-helmet';

function HomeRoute() {
  return (
    <div data-screen-id="home">
      <Helmet>
        <title>Dashboard</title>
      </Helmet>

      <h2 className="ms-font-xl">Dashboard</h2>

      <p>
        This starter kit contains all the build tooling and configuration you need to kick off your
        next universal React project, whilst containing a minimal project set up allowing you to
        make your own architecture decisions (Redux/Mobx etc).
      </p>
    </div>
  );
}

export default HomeRoute;
