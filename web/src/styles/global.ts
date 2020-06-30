import { createGlobalStyle } from 'styled-components/macro'
import { backgroundColor, textColor, placeholderColor, primaryColor, secondaryColor } from './colors'
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
  background-image: url("/bg.svg");
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: inherit;
  background-position: top center;
  color: ${textColor};
  font-size: 14px;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

h1 {
  font-family: 'Proxima Nova', Helvetica, Arial, sans-serif;
  font-size: 40px;
  font-weight: 700;
  display: block;
  margin: 20px 0px;
}

h3 {
  font-family: 'Proxima Nova', Helvetica, Arial, sans-serif;
  font-size: 20px;
  font-weight: 300;
  display: block;
  margin: 0px;
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
  margin-top: -100px;
  margin-left: 70px;
}

blockquote {
  margin: 0 10px 0 10px;
  padding: 2px 10px 2px 10px;
  border-left: 5px solid ${primaryColor};
}

.markdown code {
  background-color: ${backgroundColor};
  border-radius: 2px;
  color: ${primaryColor};
  -moz-osx-font-smoothing: auto;
  -webkit-font-smoothing: auto;
  font-family: monospace;
  padding: 0.25em 0.5em;
  font-size: 0.9em;
}

.markdown pre {
  background-color: ${backgroundColor};
  border-radius: 2px;
  margin-bottom: 1.5rem;
  padding: 1rem;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.markdown pre code {
  background-color: transparent;
  color: #d4d4d4;
}

.markdown th {
  font-size: .875rem;
  text-transform: uppercase;
  background-color: ${primaryColor};
  color: black;
  padding: .5rem;
}

.markdown td {
  padding: .5rem;
}

[data-tooltip]:hover {
  position: absolute;
  overflow-y: visible !important;
}

[data-tooltip]:hover::before {
  all: initial;
  font-family: Arial, Helvetica, sans-serif;
  display: inline-block;
  border-radius: 5px;
  padding: 10px;
  background-color: ${backgroundColor};
  content: attr(data-tooltip);
  color: ${secondaryColor};
  position: absolute;
  bottom: 100%;
  width: auto;
  left: 50%;
  transform: translate(-50%, 0);
  margin-bottom: 15px;
  text-align: center;
  font-size: 14px;
}

[data-tooltip]:hover::after {
  all: initial;
  display: inline-block;
  width: 0; 
  height: 0; 
  border-left: 10px solid transparent;
  border-right: 10px solid transparent;
  border-top: 10px solid ${backgroundColor};
  position: absolute;
  bottom: 100%;
  content: '';
  left: 50%;
  transform: translate(-50%, 0);
  margin-bottom: 5px;
}

`
