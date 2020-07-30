import * as React from 'react'
import { Link } from 'react-router-dom'
import { Button } from 'src/App/App.components/Button/Button.controller'

// prettier-ignore
import {
  HomeStyled,
  HomeLeft,
  HomeRight,
  HomeHead,
  HomeLeftInside,
  HomeBrought,
  HomeBuilt,
  HomeTrusted,
  HomeMetrics,
  HomePanels,
  HomeGetStarted,
  HomeFooter,
  HomeSection,
  HomeH1,
  HomeH3,
  HomeBuiltArg,
  HomeBuiltArgTitle,
  HomeBuiltArgText,
  HomeBuiltArgImg,
  HomeBuiltArgRev,
  HomeTrustedGrid,
  HomePanel,
  HomePanel2,
  HomeFooterGrid
} from './Home.style'

export const Home = () => (
  <HomeStyled>
    <HomeHead>
      <HomeLeft>
        <HomeLeftInside>
          <h1>Your gateway to the tezos network</h1>
          <h3>Free and scalable API access to the Tezos network</h3>
          <h3>and usage analytics for your projects</h3>
          <Link to="/new-project">
            <Button text="NEW PROJECT" icon="plus-card" />
          </Link>
        </HomeLeftInside>
      </HomeLeft>
      <HomeRight>
        <img alt="torus" className="torus-cables" src="/images/torus-cables.svg" />
        <img alt="torus" className="torus-bg" src="/images/torus-bg.svg" />
        <div className="meteor meteor1" />
        <div className="meteor meteor2" />
        <div className="meteor meteor3" />
        <div className="meteor meteor4" />
        <img alt="torus" className="torus-fg" src="/images/torus-fg.svg" />
        <img alt="torus" className="torus-logo" src="/images/torus-logo.svg" />
      </HomeRight>
      <HomeBrought alt="octo" src="/images/brought.svg" />
    </HomeHead>
    <HomeBuilt>
      <HomeSection>
        <HomeH1>Develop now with our Tezos APIs</HomeH1>
        <HomeH3>
          Tezos Link's infrastructure will ensure your decentralized application scales to meet your user demand.
        </HomeH3>
        <HomeBuiltArg>
          <div>
            <HomeBuiltArgTitle>BUILT FOR DEVELOPERS</HomeBuiltArgTitle>
            <HomeBuiltArgText>
              Connect your app immediately with our instant access APIs. We support RPC over HTTPS interfaces, providing
              high speed connections to the Tezos network.
            </HomeBuiltArgText>
          </div>
          <HomeBuiltArgImg alt="tezos" src="/images/arg1.svg" />
        </HomeBuiltArg>
        <HomeBuiltArgRev>
          <HomeBuiltArgImg alt="tezos" src="/images/arg2.svg" />
          <div>
            <HomeBuiltArgTitle>BUILT FOR EASE</HomeBuiltArgTitle>
            <HomeBuiltArgText>
              Start using Tezos Link with a single URL. Our 24/7 team of experts is ready to handle all network changes
              and upgrades so you can focus on building your applications.
            </HomeBuiltArgText>
          </div>
        </HomeBuiltArgRev>
        <HomeBuiltArg>
          <div>
            <HomeBuiltArgTitle>BUILT FOR BUILDERS</HomeBuiltArgTitle>
            <HomeBuiltArgText>
              We believe in a future powered by decentralized networks and protocols. We provide world-class
              infrastructure for developers so you can spend your time building and creating.
            </HomeBuiltArgText>
          </div>
          <HomeBuiltArgImg alt="tezos" src="/images/arg3.svg" />
        </HomeBuiltArg>
        <HomeBuiltArgRev>
          <HomeBuiltArgImg alt="tezos" src="/images/architecture.png" />
          <div>
            <HomeBuiltArgTitle>SCALABLE</HomeBuiltArgTitle>
            <HomeBuiltArgText>
              Our architecture supports the workload required by your project, by scaling up Tezos nodes when we see an
              increase of requests. The infrastructure is open-sourced in our{' '}
              <a href="https://github.com/octo-technology/tezos-link" target="_blank">
                Github project
              </a>
            </HomeBuiltArgText>
          </div>
        </HomeBuiltArgRev>
      </HomeSection>
    </HomeBuilt>
    <HomeTrusted>
      <HomeH1>Trusted by hundreds of developers</HomeH1>
      <HomeH3>
        Used worldwide by dozens of production applications without having to install or manage a single node.
      </HomeH3>
      <HomeTrustedGrid>
        <img alt="accenture" src="/images/accenture.svg" />
        <img alt="nomadic" src="/images/nomadic.svg" />
        <img alt="tq" src="/images/tq.svg" />
      </HomeTrustedGrid>
      <a href="mailto:beta@octo.com" target="_blank">
        <Button text="CONTACT US FOR ENTERPRISE INFOS" icon="plus-card" />
      </a>
    </HomeTrusted>
    <HomeMetrics>
      <HomeH1>Insights from your app</HomeH1>
      <HomeH3>The Tezos Link dashboard allows you to get valuable statistics from your utilization of the APIs</HomeH3>
      <img alt="metrics" src="/images/metrics.png" />
    </HomeMetrics>
    <HomePanels>
      <Link to="/documentation">
        <HomePanel>
          <img className="logo" alt="books" src="/icons/books.svg" />
          <div>
            <h3>DOCUMENTATION</h3>
            <p>Learn how to use Tezos Link</p>
          </div>
          <img className="arrow" alt="arrow" src="/icons/arrow-black.svg" />
        </HomePanel>
      </Link>
      <a href="mailto:beta@octo.com" target="_blank">
        <HomePanel2>
          <img className="logo" alt="books" src="/icons/support.svg" />
          <div>
            <h3>SUPPORT</h3>
            <p>Ask us your questions</p>
          </div>
          <img className="arrow" alt="arrow" src="/icons/arrow-white.svg" />
        </HomePanel2>
      </a>
    </HomePanels>
    <HomeGetStarted>
      <HomeH1>Connect to TEZOS now</HomeH1>
      <Link to="/new-project">
        <Button text="GET STARTED" icon="plus-card" />
      </Link>
      <img className="left" alt="torus" src="/images/mini-torus.png" />
      <img className="right" alt="torus" src="/images/mini-torus.png" />
    </HomeGetStarted>
    <HomeFooter>
      <HomeFooterGrid>
        <img alt="logo" src="/images/logo.svg" />
        <div>
          <p>About Tezos Link</p>
          <a href="https://github.com/octo-technology/tezos-link" target="_blank">
            Github
          </a>
          <a href="mailto:beta@octo.com" target="_blank">
            Support
          </a>
          <a href="https://www.reddit.com/r/tezos/" target="_blank">
            Reddit
          </a>
        </div>
        <div>
          <p>About OCTO</p>
          <a href="https://octo.com" target="_blank">
            Homepage
          </a>
          <a href="https://blog.octo.com" target="_blank">
            Our blog
          </a>
        </div>
        <div>
          <p>About the devs</p>
          <a href="https://www.linkedin.com/in/aymeric-bethencourt-96665046/" target="_blank">
            Aymeric Bethencourt
          </a>
          <a href="https://www.linkedin.com/in/adrien-boulay-2679aa113/" target="_blank">
            Adrien Boulay
          </a>
          <a href="https://www.linkedin.com/in/loup-theron-b1785397/" target="_blank">
            Loup Theron
          </a>
        </div>
        <img alt="octo" src="/images/brought.svg" />
      </HomeFooterGrid>
    </HomeFooter>
  </HomeStyled>
)
