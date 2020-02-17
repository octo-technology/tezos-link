import styled from 'styled-components/macro'
import { Card, subTextColor, textColor, primaryColor } from 'src/styles'

export const CategoriesStyled = styled(Card)`
  padding: 20px 20px 1px 20px;
  margin-bottom: 20px;
`

export const CategoriesSlider = styled.div``

export const CategoriesCategory = styled.div`
  cursor: pointer;
  display: grid;
  grid-template-columns: 36px auto;
  grid-gap: 20px;
  margin-bottom: 20px;
`

export const CategoriesCategoryIcon = styled.div`
  padding: 6px;

  > svg {
    height: 24px;
    width: 24px;
    stroke: ${subTextColor};

    &.selected {
      stroke: ${primaryColor};
    }
  }
`
export const CategoriesCategoryData = styled.div``

export const CategoriesCategoryTitle = styled.div`
  font-size: 16px;
  font-weight: 700;
  color: ${textColor};

  &.selected {
    color: ${primaryColor};
  }
`

export const CategoriesCategoryDescription = styled.div`
  color: ${subTextColor};
`
