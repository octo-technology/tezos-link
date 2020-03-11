import { createGlobalStyle } from 'styled-components/macro'
import { backgroundColor, textColor, placeholderColor } from './colors'
import { fadeInFromLeft } from './animations'

export const GlobalStyle = createGlobalStyle`
  @font-face {
  font-family: 'Proxima Nova';
  src: url('/fonts/ProximaNova-Thin.woff2') format('woff2'), url('/fonts/ProximaNova-Thin.woff') format('woff');
  font-weight: 100;
  font-style: normal;
}

@font-face {
  font-family: 'Proxima Nova';
  src: url('/fonts/ProximaNova-Light.woff2') format('woff2'), url('/fonts/ProximaNova-Light.woff') format('woff');
  font-weight: 300;
  font-style: normal;
}

@font-face {
  font-family: 'Proxima Nova';
  src: url('/fonts/ProximaNova-Regular.woff2') format('woff2'), url('/fonts/ProximaNova-Regular.woff') format('woff');
  font-weight: normal;
  font-style: normal;
}

@font-face {
  font-family: 'Proxima Nova';
  src: url('/fonts/ProximaNova-Semibold.woff2') format('woff2'), url('/fonts/ProximaNova-Semibold.woff') format('woff');
  font-weight: 600;
  font-style: normal;
}

@font-face {
  font-family: 'Proxima Nova';
  src: url('/fonts/ProximaNova-Bold.woff2') format('woff2'), url('/fonts/ProximaNova-Bold.woff') format('woff');
  font-weight: 700;
  font-style: normal;
}

* {
  box-sizing: border-box;
}

body {
  font-family: 'Proxima Nova', Helvetica, Arial, sans-serif;
  margin: 0;
  padding: 0;
  background-color: ${backgroundColor};
  background-image: url(/background.svg);
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: 100%;
  color: ${textColor};
  font-size: 14px;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

h1 {
  font-family: 'Proxima Nova', Helvetica, Arial, sans-serif;
  font-size: 40px;
  font-weight: 700;
  display: inline-block;
  margin: 20px 0px;
}

input {
  color: ${textColor};
  font-size: 14px;
}

::placeholder {
  color: ${placeholderColor};
  font-size: 14px;
}

*:focus {
  outline: none;
}

a, a:visited {
  color: ${textColor};
  text-decoration: none !important;
  opacity: 1;
  transition: opacity 0.15s ease-in-out;
  will-change: opacity;
}

a:hover {
  opacity: 0.9;
}

code {
  font-family: source-code-pro, Menlo, Monaco, Consolas, 'Courier New', monospace;
}
 
.appear {
  opacity: 0;
  will-change: transform, opacity;
  animation: ${fadeInFromLeft} ease-in 1;
  animation-fill-mode: forwards;
  animation-duration: 0.2s;
}

#confetis {
  z-index: -1;
  position: fixed;
  margin-top: -70px;
}
`
