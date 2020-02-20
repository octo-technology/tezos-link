import * as PropTypes from 'prop-types'
import * as React from 'react'

//prettier-ignore
import { CategoriesCategory, CategoriesCategoryData, CategoriesCategoryDescription, CategoriesCategoryIcon, CategoriesCategoryTitle, CategoriesSlider, CategoriesStyled } from './Categories.style'

type Category = {
  title: string
  slug: string
  description: string
  icon: string
}

type CategoriesViewProps = {
  categories: Category[]
  selectedCategory: number
  categoryCallback: (e: any) => void
}

export const CategoriesView = ({ categories, selectedCategory, categoryCallback }: CategoriesViewProps) => {
  return (
    <CategoriesStyled>
      <CategoriesSlider className={`goto${selectedCategory}`} />
      {categories.map((category, i) => (
        <CategoriesCategory key={category.slug} onClick={(e: any) => categoryCallback(e)}>
          <CategoriesCategoryIcon>
            <svg className={selectedCategory === i ? 'selected' : ''}>
              <use xlinkHref={`/icons/sprites.svg#${category.icon}`} />
            </svg>
          </CategoriesCategoryIcon>
          <CategoriesCategoryData>
            <CategoriesCategoryTitle className={selectedCategory === i ? 'selected' : ''}>
              {category.title}
            </CategoriesCategoryTitle>
            <CategoriesCategoryDescription>{category.description}</CategoriesCategoryDescription>
          </CategoriesCategoryData>
        </CategoriesCategory>
      ))}
    </CategoriesStyled>
  )
}

CategoriesView.propTypes = {
  categories: PropTypes.array,
  selectedCategory: PropTypes.number,
  categoryCallback: PropTypes.func.isRequired
}

CategoriesView.defaultProps = {
  categories: [],
  selectedCategory: 0
}
