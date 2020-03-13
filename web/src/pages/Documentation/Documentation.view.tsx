import * as React from 'react'
import {
  DocumentationViewContent,
  DocumentationViewLeftSide,
  DocumentationViewMenu,
  DocumentationViewRightSide,
  DocumentationViewStyled
} from './Documentation.style'

import ReactMarkdown from 'react-markdown'
import * as PropTypes from 'prop-types'

function flatten(text: string, child: any): boolean {
  return typeof child === 'string' ? text + child : React.Children.toArray(child.props.children).reduce(flatten, text)
}

function HeadingRenderer(props: any) {
  const children = React.Children.toArray(props.children)
  const text = children.reduce(flatten, '')
  const slug = text.toLowerCase().replace(/\W/g, '-')
  return React.createElement('h' + props.level, { id: slug }, props.children)
}

type DocumentationViewProps = {
  menu: string
  content: string
}

export const DocumentationView = ({ menu, content }: DocumentationViewProps) => (
  <DocumentationViewStyled>
    <DocumentationViewLeftSide>
      <DocumentationViewMenu>
        <h2>Getting started</h2>
        <ReactMarkdown className={'markdown'} source={menu} renderers={{ heading: HeadingRenderer }} />
      </DocumentationViewMenu>
    </DocumentationViewLeftSide>
    <DocumentationViewRightSide>
      <DocumentationViewContent>
        <ReactMarkdown className={'markdown'} source={content} renderers={{ heading: HeadingRenderer }} />
      </DocumentationViewContent>
    </DocumentationViewRightSide>
  </DocumentationViewStyled>
)

DocumentationView.propTypes = {
  menu: PropTypes.string,
  content: PropTypes.string
}

DocumentationView.defaultProps = {
  menu: '',
  content: ''
}
